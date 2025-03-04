package doitpay

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"os"

	"github.com/automotechnologies/doitpay-go/v2/client"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// ClientOption allows for functional options to configure the client
type ClientOption func(*ClientConfig)

// WithHost sets a custom host
func WithHost(host string) ClientOption {
    return func(c *ClientConfig) {
        c.Host = host
    }
}

// WithBasePath sets a custom base path
func WithBasePath(basePath string) ClientOption {
    return func(c *ClientConfig) {
        c.BasePath = basePath
    }
}

// DefaultConfig provides the default configuration for the client
func DefaultConfig() ClientConfig {
    return ClientConfig{
        Host:     "api.doitpay.co", // example default
        BasePath: "",             // example default
    }
}

// DoitpayClient wraps the generated client with custom auth
type DoitpayClient struct {
    doitpay *client.Doitpay
    config ClientConfig

    // Services
    qris *QrisClient
    disbursement *DisbursementClient
    simulate *SimulateClient
}

type ClientConfig struct {
    Host      string
    BasePath  string
    ClientSecret string
    PrivateKey *rsa.PrivateKey
}
// NewClient creates a new authenticated client with optional configurations
func NewClient(clientSecret, privateKeyPath string, opts ...ClientOption) (*DoitpayClient, error) {
    // Start with default config
    cfg := DefaultConfig()
    
    // Apply required credentials
    cfg.ClientSecret = clientSecret
    
    // Apply any custom options
    for _, opt := range opts {
        opt(&cfg)
    }


    // validate privateKeyPath
    if _, err := os.Stat(privateKeyPath); os.IsNotExist(err) {
        return nil, fmt.Errorf("private key file does not exist: %s", privateKeyPath)
    }

    // read private key
    privateKeyBytes, err := os.ReadFile(privateKeyPath)
    if err != nil {
        return nil, fmt.Errorf("failed to read private key: %s", err)
    }

    // parse private key
    privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyBytes)
    if err != nil {
        return nil, fmt.Errorf("failed to parse private key: %s", err)
    }


    rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
    if !ok {
        return nil, fmt.Errorf("private key is not an RSA private key")
    }
    cfg.PrivateKey = rsaPrivateKey

    // Create transport with auth
    transport := httptransport.New(cfg.Host, cfg.BasePath, []string{"https"})

    // Create base client
    baseClient := client.New(transport, strfmt.Default)
    transport.DefaultAuthentication = &doitpayAuth{config: cfg, authService: baseClient.Authentication}


    qrisClient := NewQrisClient(baseClient.Qris)
    disbursementClient := NewDisbursementClient(baseClient.Disbursement)
    simulateClient := NewSimulateClient(baseClient.PublicSimulate)
    return &DoitpayClient{
        doitpay: baseClient,
        config:  cfg,
        qris: qrisClient,
        disbursement: disbursementClient,
        simulate: simulateClient,
    }, nil
}

// Qris returns the QRIS client
func (c *DoitpayClient) Qris() *QrisClient {
    return c.qris
}
