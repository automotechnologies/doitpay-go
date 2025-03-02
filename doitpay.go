package doitpay

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/automotechnologies/doitpay-go/client"
	"github.com/automotechnologies/doitpay-go/client/authentication"
	"github.com/automotechnologies/doitpay-go/models"
	"github.com/automotechnologies/doitpay-go/pkg/signature"

	"github.com/go-openapi/runtime"
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
}

type ClientConfig struct {
    Host      string
    BasePath  string
    ClientSecret string
    PrivateKey *rsa.PrivateKey
}

// Custom auth implementation
type doitpayAuth struct {
    config ClientConfig
    accessToken *AccessToken
    authService authentication.ClientService
}

func (a *doitpayAuth) GetAccessToken(ctx context.Context) (string, error) {
    if a.accessToken == nil || a.accessToken.ExpiresAt.Before(time.Now()) {
    
        timestamp := time.Now().UTC()
        signature, err := signature.GenerateSNAPAsymmetricSignature(a.config.ClientSecret, timestamp, a.config.PrivateKey)
        if err != nil {
            return "", err
        }

        resp, err := a.authService.AccessToken(&authentication.AccessTokenParams{
            XTIMESTAMP: timestamp.Format(time.RFC3339),
            XSIGNATURE: signature,
            XCLIENTKEY: a.config.ClientSecret,
            Request: &models.Apiv1PartnerAccessTokenRequest{
                GrantType: "client_credentials",
            },
        })

        if err != nil {
            return "", err
        }

        expiresIn, err := strconv.Atoi(resp.GetPayload().ExpiresIn)
        if err != nil {
            return "", err
        }

        expiresAt := time.Now().Add(time.Duration(expiresIn) * time.Second)
        a.accessToken = &AccessToken{
            AccessToken: resp.GetPayload().AccessToken,
            ExpiresAt:   expiresAt,
            TokenType:   resp.GetPayload().TokenType,
        }
        return a.accessToken.AccessToken, nil
    }
    return a.accessToken.AccessToken, nil
}

func (a *doitpayAuth) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
    timestamp := time.Now().UTC()
    
    // Set required headers
    headers := map[string]string{
        "X-CLIENT-KEY": a.config.ClientSecret,
        "X-TIMESTAMP":  timestamp.Format(time.RFC3339),
        "X-SIGNATURE":  "",
    }

    accessToken, err := a.GetAccessToken(context.Background())
    if err != nil {
        return err
    }

    endpointUrl, err := url.JoinPath(a.config.Host, a.config.BasePath, req.GetPath())
    if err != nil {
        return err
    }

    // Generate signature
    signature, err := signature.GenerateSNAPSymmetricSignature(
        req.GetMethod(),
        endpointUrl,
        accessToken,
        timestamp,
        req.GetBody(),
    )

    if err != nil {
        return err
    }

    headers["X-SIGNATURE"] = signature

    // Set all headers
    for key, value := range headers {
        if err := req.SetHeaderParam(key, value); err != nil {
            return err
        }
    }

    return nil
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
    baseClient := client.New(transport, strfmt.Default)
    transport.DefaultAuthentication = &doitpayAuth{config: cfg, authService: baseClient.Authentication}
    // Create base client

    return &DoitpayClient{
        doitpay: baseClient,
        config:  cfg,
    }, nil
}

type AccessToken struct {
    AccessToken string
    ExpiresAt   time.Time
    TokenType   string
}
