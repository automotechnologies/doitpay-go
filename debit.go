package doitpay

import (
	"context"

	"github.com/automotechnologies/doitpay-go/client/debit"
	"github.com/automotechnologies/doitpay-go/models"
)

// Parameter structs
type PaymentHostToHostParams struct {
	Request    *models.CreateEwalletRequest
	ExternalID string
	ChannelID  string
	RequestID  string
}

type CancelParams struct {
	Request    *models.CancelHostToHostPaymentRequest
	ExternalID string
	ChannelID  string
}

type StatusParams struct {
	Request    *models.CheckPaymentStatusEwalletRequest
	ExternalID string
	ChannelID  string
	RequestID  string
}

type DebitClient struct {
	debitClient debit.ClientService
}

func NewDebitClient(clientService debit.ClientService) *DebitClient {
	return &DebitClient{
		debitClient: clientService,
	}
}

// Create creates a new payment transaction
func (c *DebitClient) PaymentHostToHost(ctx context.Context, params *PaymentHostToHostParams) (*models.CreateEwalletResponse, error) {
	result, err := c.debitClient.PostAPIDebitV10PaymentHostToHost(
		&debit.PostAPIDebitV10PaymentHostToHostParams{
			Request:     params.Request,
			XEXTERNALID: params.ExternalID,
			CHANNELID:   params.ChannelID,
			XREQUESTID:  params.RequestID,
			Context:     ctx,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}

// Cancel cancels an existing payment transaction
func (c *DebitClient) Cancel(ctx context.Context, params *CancelParams) (*models.CancelHostToHostPaymentResponse, error) {
	result, err := c.debitClient.PostAPIV10DebitCancel(
		&debit.PostAPIV10DebitCancelParams{
			Request:     params.Request,
			XEXTERNALID: params.ExternalID,
			CHANNELID:   params.ChannelID,
			Context:     ctx,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}

// Status checks the status of a payment transaction
func (c *DebitClient) Status(ctx context.Context, params *StatusParams) (*models.CheckPaymentStatusEwalletResponse, error) {
	result, err := c.debitClient.PostAPIV10DebitStatus(
		&debit.PostAPIV10DebitStatusParams{
			Request:     params.Request,
			XEXTERNALID: params.ExternalID,
			CHANNELID:   params.ChannelID,
			XREQUESTID:  params.RequestID,
			Context:     ctx,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}
