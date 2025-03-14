package doitpay

import (
	"context"
	"fmt"
	"github.com/automotechnologies/doitpay-go/pkg/signature"
	"github.com/stretchr/testify/require"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestAuthenticate_ValidateSignature(t *testing.T) {
	cc := ClientConfig{
		ClientSecret: "Secret",
		PrivateKey:   nil,
	}
	auth := NewDoitpayAuth(cc)

	testCases := []struct {
		name        string
		method      string
		endpoint    string
		accessToken string
		timestamp   time.Time
		body        string
	}{
		{
			name:        "success",
			method:      "POST",
			endpoint:    "/api",
			accessToken: "AccessToken",
			timestamp:   time.Now(),
			body:        `{"grant_type": "credentials"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Generate signature
			sign, err := signature.GenerateSNAPSymmetricSignature(
				tc.method,
				tc.endpoint,
				tc.accessToken,
				cc.ClientSecret,
				tc.timestamp,
				[]byte(tc.body),
			)
			require.NoError(t, err)

			req, err := http.NewRequest(tc.method, tc.endpoint, strings.NewReader(tc.body))
			require.NoError(t, err)

			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tc.accessToken))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-External-ID", "123123123")
			req.Header.Set("X-Partner-ID", "123123123")
			req.Header.Set("X-Signature", sign)
			req.Header.Set("X-Timestamp", tc.timestamp.Format(time.RFC3339))

			err = auth.ValidateSignature(context.Background(), req)
			require.NoError(t, err)
		})
	}

}
