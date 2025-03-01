package doitpay

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/automotechnologies/doitpay-go/client"
	"github.com/automotechnologies/doitpay-go/client/authentication"
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
        Host:     "api.doitpay.com", // example default
        BasePath: "/v1",             // example default
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
    PrivateKeyPath string
}

// Custom auth implementation
type doitpayAuth struct {
    config ClientConfig
}

func (a *doitpayAuth) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
    timestamp := time.Now().UTC().Format(time.RFC3339)
    
    // Set required headers
    headers := map[string]string{
        "X-CLIENT-KEY": a.config.ClientKey,
        "X-TIMESTAMP":  timestamp,
    }

    // Generate signature
    signature, err := generateSignature(
        req.GetMethod(),
        req.GetPath(),
        timestamp,
        a.config.ClientKey,
        a.config.SecretKey,
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

func generateSignature(method, path, timestamp, clientKey, secretKey string, body interface{}) (string, error) {
    // Example signature generation - adjust according to your needs
    data := method + ":" + path + ":" + timestamp + ":" + clientKey
    
    h := hmac.New(sha256.New, []byte(secretKey))
    h.Write([]byte(data))
    
    return hex.EncodeToString(h.Sum(nil)), nil
}

// NewClient creates a new authenticated client with optional configurations
func NewClient(clientSecret, privateKeyPath string, opts ...ClientOption) *DoitpayClient {
    // Start with default config
    cfg := DefaultConfig()
    
    // Apply required credentials
    cfg.ClientSecret = clientSecret
    cfg.PrivateKeyPath = privateKeyPath
    
    // Apply any custom options
    for _, opt := range opts {
        opt(&cfg)
    }

    // Create transport with auth
    transport := httptransport.New(cfg.Host, cfg.BasePath, []string{"https"})
    transport.DefaultAuthentication = &doitpayAuth{config: cfg}

    // Create base client
    baseClient := client.New(transport, strfmt.Default)

    return &DoitpayClient{
        Doitpay: baseClient,
        config:  cfg,
    }
}


func (c *DoitpayClient) GetAccessToken(ctx context.Context) (string, error) {
    timestamp := time.Now().UTC()
    signature, err := signature.GenerateSNAPAsymmetricSignature(c.config.ClientSecret, timestamp, c.config.PrivateKey)
    if err != nil {
        return "", err
    }

    resp, err := c.doitpay.Authentication.PostAuthV10AccessTokenB2b(&authentication.PostAuthV10AccessTokenB2bParams{
        XTIMESTAMP: timestamp.Format(time.RFC3339),
        XSIGNATURE: signature,
        XCLIENTKEY: c.config.ClientSecret,
    })

    if err != nil {
        return "", err
    }

    return resp.Payload.Validate(strfmt.Default), nil


}