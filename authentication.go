package doitpay

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/automotechnologies/doitpay-go/client/authentication"
	"github.com/automotechnologies/doitpay-go/models"
	"github.com/automotechnologies/doitpay-go/pkg/signature"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AccessToken represents an access token with expiration time
type AccessToken struct {
    AccessToken string
    ExpiresAt   time.Time
    TokenType   string
}

// Custom auth implementation
type doitpayAuth struct {
    config ClientConfig
    accessToken *AccessToken
    authService authentication.ClientService
}


// GetAccessToken retrieves an access token from the authentication service
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
            Request: &models.PartnerAccessTokenRequest{
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

// AuthenticateRequest authenticates a request with the access token
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

