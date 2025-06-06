package doitpay

import (
	"context"

	"github.com/automotechnologies/doitpay-go/client/q_r_i_s"
	"github.com/automotechnologies/doitpay-go/models"
)

// Parameter structs
type GenerateQRISParams struct {
	Request    *models.QrisRequestScheme
	ExternalID string
	ChannelID  string
}

type QueryQRISParams struct {
	Request    *models.QrisQueryPaymentRequest
	ExternalID string
	ChannelID  string
}

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
func (c *QrisClient) GenerateQRISCode(ctx context.Context, params *GenerateQRISParams) (*models.QrisResponseScheme, error) {
	result, err := c.qrisClient.PostQrisV10QrQrMpmGenerate(
		&q_r_i_s.PostQrisV10QrQrMpmGenerateParams{
			Request:     params.Request,
			XEXTERNALID: params.ExternalID,
			CHANNELID:  params.ChannelID,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}

// QueryQRISCode queries a QRIS code
func (c *QrisClient) QueryQRISCode(ctx context.Context, params *QueryQRISParams) (*models.QrisQueryPaymentResponse, error) {
	result, err := c.qrisClient.PostQrisV10QrQrMpmQuery(
		&q_r_i_s.PostQrisV10QrQrMpmQueryParams{
			Request:     params.Request,
			XEXTERNALID: params.ExternalID,
			CHANNELID:   params.ChannelID,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}

	return result.Payload, nil
}
