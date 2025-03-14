package doitpay

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/automotechnologies/doitpay-go/client"

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
		Host:     client.DefaultHost,     // example default
		BasePath: client.DefaultBasePath, // example default
	}
}

// DoitpayClient wraps the generated client with custom auth
type DoitpayClient struct {
	doitpay *client.Doitpay
	config  ClientConfig

	// Services
	qris           *QrisClient
	disbursement   *DisbursementClient
	simulate       *SimulateClient
	merchant       *MerchantClient
	virtualAccount *VirtualAccountClient
}

type ClientConfig struct {
	Host         string
	BasePath     string
	ClientSecret string
	PartnerID    string
	PrivateKey   *rsa.PrivateKey
}

// NewClient creates a new authenticated client with optional configurations
func NewClient(partnerID, clientSecret, privateKeyPath string, opts ...ClientOption) (*DoitpayClient, error) {
	// Start with default config
	cfg := DefaultConfig()

	// Apply required credentials
	cfg.PartnerID = partnerID
	cfg.ClientSecret = clientSecret

	// Apply any custom options
	for _, opt := range opts {
		opt(&cfg)
	}

	// validate privateKeyPath
	privateKeyBytes := []byte(privateKeyPath)
	if _, err := os.Stat(privateKeyPath); os.IsExist(err) {
		// read private key
		privateKeyBytes, err = os.ReadFile(privateKeyPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read private key: %s", err)
		}
	}

	// parse private key
	privatePem, _ := pem.Decode(privateKeyBytes)
	if privatePem == nil {
		return nil, fmt.Errorf("failed to decode private key")
	}

	if privatePem.Type != "RSA PRIVATE KEY" && privatePem.Type != "ENCRYPTED PRIVATE KEY" && privatePem.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("private key is not a valid PEM file")
	}

	parsedKey, err := x509.ParsePKCS8PrivateKey(privatePem.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %s", err)
	}

	rsaPrivateKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("private key is not an RSA private key")
	}
	cfg.PrivateKey = rsaPrivateKey

	// Create transport with auth
	transport := httptransport.New(cfg.Host, cfg.BasePath, []string{"https", "http"})

	// Need to use default authentication that follows SNAP flows.
	transport.DefaultAuthentication = NewDoitpayAuth(cfg)

	// Create base client
	baseClient := client.New(transport, strfmt.Default)

	qrisClient := NewQrisClient(baseClient.Qris)
	disbursementClient := NewDisbursementClient(baseClient.Disbursement)
	simulateClient := NewSimulateClient(baseClient.PublicSimulate)
	merchant := NewMerchantClient(baseClient.Merchants)
	virtualAccount := NewVirtualAccountClient(baseClient.VirtualAccount)
	return &DoitpayClient{
		doitpay:        baseClient,
		config:         cfg,
		qris:           qrisClient,
		disbursement:   disbursementClient,
		simulate:       simulateClient,
		merchant:       merchant,
		virtualAccount: virtualAccount,
	}, nil
}

// Qris returns the QRIS client
func (c *DoitpayClient) Qris() *QrisClient {
	return c.qris
}

// Disbursement returns the Disbursement client
func (c *DoitpayClient) Disbursement() *DisbursementClient {
	return c.disbursement
}

// Simulate returns the Simulate client
func (c *DoitpayClient) Simulate() *SimulateClient {
	return c.simulate
}

// Merchant return the Merchant client
func (c *DoitpayClient) Merchant() *MerchantClient {
	return c.merchant
}

// VirtualAccount return the VA client
func (c *DoitpayClient) VirtualAccount() *VirtualAccountClient {
	return c.virtualAccount
}
