package doitpay

import (
	"context"

	"github.com/automotechnologies/doitpay-go/v2/client/public_simulate"
	"github.com/automotechnologies/doitpay-go/v2/models"
)

type SimulateClient struct {
	simulate public_simulate.ClientService
}

func NewSimulateClient(client public_simulate.ClientService) *SimulateClient {
	return &SimulateClient{
		simulate: client,
	}
}

func (c *SimulateClient) SimulateQrisPayment(ctx context.Context, request *models.QrisSimulateRequest) (*public_simulate.SimulatePaymentOK, error) {
	res, err := c.simulate.SimulatePayment(&public_simulate.SimulatePaymentParams{
		Request: request,
	},
		nil,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
