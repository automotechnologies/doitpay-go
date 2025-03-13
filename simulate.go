package doitpay

import (
	"context"

	"github.com/automotechnologies/doitpay-go/client/public_simulate"
	"github.com/automotechnologies/doitpay-go/models"
)

// Parameter struct
type SimulateQRISParams struct {
	Request *models.QrisSimulateRequest
}

type SimulateClient struct {
	simulate public_simulate.ClientService
}

func NewSimulateClient(client public_simulate.ClientService) *SimulateClient {
	return &SimulateClient{
		simulate: client,
	}
}

func (c *SimulateClient) SimulateQrisPayment(ctx context.Context, params *SimulateQRISParams) (*public_simulate.SimulatePaymentOK, error) {
	res, err := c.simulate.SimulatePayment(&public_simulate.SimulatePaymentParams{
		Request: params.Request,
	},
		nil,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
