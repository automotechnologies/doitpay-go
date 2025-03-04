package doitpay

import (
	"context"

	"github.com/automotechnologies/doitpay-go/v2/client/q_r_i_s"
	"github.com/automotechnologies/doitpay-go/v2/models"
)

// QrisClient is a client for the QRIS API
type QrisClient struct {
	qrisClient q_r_i_s.ClientService
}

// NewQrisClient creates a new QRIS client
func NewQrisClient(clientService q_r_i_s.ClientService) *QrisClient {
	return &QrisClient{
		qrisClient: clientService,
	}
}

// GenerateQRISCode generates a QRIS code
func (c *QrisClient) GenerateQRISCode(ctx context.Context, request *models.QrisRequestScheme) (*models.QrisResponseScheme, error) {
	result, err := c.qrisClient.PostQrisV10QrQrMpmGenerate(
		&q_r_i_s.PostQrisV10QrQrMpmGenerateParams{
			Request: request,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}

// QueryQRISCode queries a QRIS code
func (c *QrisClient) QueryQRISCode(ctx context.Context, request *models.QrisQueryPaymentRequest) (*models.QrisQueryPaymentResponse, error) {
	result, err := c.qrisClient.PostQrisV10QrQrMpmQuery(
		&q_r_i_s.PostQrisV10QrQrMpmQueryParams{
			Request: request,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}

	return result.Payload, nil
}