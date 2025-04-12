package doitpay

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/automotechnologies/doitpay-go/client"
	"github.com/automotechnologies/doitpay-go/client/authentication"
	"github.com/automotechnologies/doitpay-go/models"
	"github.com/automotechnologies/doitpay-go/pkg/signature"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// AccessToken represents an access token with expiration time
type AccessToken struct {
	AccessToken string
	ExpiresAt   time.Time
	TokenType   string
}

var (
	ErrorBearerTokenNotFound    = errors.New("mandatory header bearer token not found")
	ErrorAccessTokenInvalid     = errors.New("access token invalid")
	ErrorTimestampNotFound      = errors.New("mandatory header timestamp not found")
	ErrorExternalIdNotFound     = errors.New("mandatory header external id not found")
	ErrorPartnerIdNotFound      = errors.New("mandatory header partner id not found")
	ErrorSignatureNotFound      = errors.New("mandatory header signature not found")
	ErrorPartnerIdNotRegistered = errors.New("ValidateSignature: partner id not registered")
	ErrorSignatureInvalid       = errors.New("ValidateSignature: invalid signature")
)

// Custom auth implementation
type DoitpayAuth struct {
	config      ClientConfig
	accessToken *AccessToken
	mu          sync.Mutex
}

func NewDoitpayAuth(config ClientConfig) *DoitpayAuth {
	return &DoitpayAuth{config: config}
}

// GetAccessToken retrieves an access token from the authentication service
func (a *DoitpayAuth) GetAccessToken(ctx context.Context) (string, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.accessToken == nil || a.accessToken.ExpiresAt.Before(time.Now()) {

		timestamp := time.Now().UTC()
		signature, err := signature.GenerateSNAPAsymmetricSignature(a.config.ClientSecret, timestamp, a.config.PrivateKey)
		if err != nil {
			return "", err
		}

		transport := httptransport.New(a.config.Host, a.config.BasePath, []string{a.config.Scheme})
		authService := client.New(transport, strfmt.Default).Authentication

		resp, err := authService.AccessToken(&authentication.AccessTokenParams{
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
func (a *DoitpayAuth) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
	timestamp := time.Now().UTC()

	// Set required headers
	headers := map[string]string{
		"X-CLIENT-KEY": a.config.ClientSecret,
		"X-TIMESTAMP":  timestamp.Format(time.RFC3339),
		"X-PARTNER-ID": a.config.PartnerID,
		"X-SIGNATURE":  "",
	}

	accessToken, err := a.GetAccessToken(context.Background())
	if err != nil {
		return err
	}

	endpoint := req.GetPath()
	if params := req.GetQueryParams(); len(params) > 0 {
		endpoint = endpoint + "?" + params.Encode()
	}

	// Generate signature
	signature, err := signature.GenerateSNAPSymmetricSignature(
		req.GetMethod(),
		endpoint,
		accessToken,
		a.config.ClientSecret,
		timestamp,
		req.GetBody(),
	)

	if err != nil {
		return err
	}

	headers["X-SIGNATURE"] = signature
	headers["Authorization"] = "Bearer " + accessToken

	// Set all headers
	for key, value := range headers {
		if err := req.SetHeaderParam(key, value); err != nil {
			return err
		}
	}

	return nil
}

func (a *DoitpayAuth) ValidateSignature(ctx context.Context, req *http.Request) error {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return errors.Wrap(err, "ValidateSignature: failed to read request body")
	}

	req.Body = io.NopCloser(bytes.NewBuffer(body))

	bearerToken := req.Header.Get("Authorization")
	if bearerToken == "" {
		return ErrorBearerTokenNotFound
	}

	if len(strings.Split(bearerToken, " ")) < 2 {
		return ErrorAccessTokenInvalid
	}

	if strings.Split(bearerToken, " ")[0] != "Bearer" {
		return ErrorAccessTokenInvalid
	}

	headerTimestamp := req.Header.Get("X-TIMESTAMP")
	if headerTimestamp == "" {
		return ErrorTimestampNotFound
	}

	externalId := req.Header.Get("X-EXTERNAL-ID")
	if externalId == "" {
		return ErrorExternalIdNotFound
	}

	partnerId := req.Header.Get("X-PARTNER-ID")
	if partnerId == "" {
		return ErrorPartnerIdNotFound
	}

	headerSignatureKey := req.Header.Get("X-SIGNATURE")
	if headerSignatureKey == "" {
		return ErrorSignatureNotFound
	}

	timestamp, err := time.Parse(time.RFC3339, headerTimestamp)
	if err != nil {
		return ErrorTimestampNotFound
	}

	token := strings.Split(bearerToken, " ")[1]
	expectedSignature, err := signature.GenerateSNAPSymmetricSignature(req.Method, req.URL.Path, token,
		a.config.ClientSecret, timestamp, body)
	if err != nil {
		return ErrorSignatureInvalid
	}

	if expectedSignature != headerSignatureKey {
		return ErrorSignatureInvalid
	}
	return nil
}
