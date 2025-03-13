package doitpay

import (
	"context"

	"github.com/automotechnologies/doitpay-go/client/payment"
	"github.com/automotechnologies/doitpay-go/models"
)

// Parameter structs
type CreatePaymentParams struct {
	Request    *models.CreateEwalletRequest
	ExternalID string
	ChannelID  string
	RequestID  string
}

type CancelPaymentParams struct {
	Request    *models.CancelHostToHostPaymentRequest
	ExternalID string
	ChannelID  string
}

type CheckPaymentStatusParams struct {
	Request    *models.CheckPaymentStatusEwalletRequest
	ExternalID string
	ChannelID  string
	RequestID  string
}

type PaymentClient struct {
	paymentClient payment.ClientService
}

func NewPaymentClient(clientService payment.ClientService) *PaymentClient {
	return &PaymentClient{
		paymentClient: clientService,
	}
}

// Create creates a new payment transaction
func (c *PaymentClient) Create(ctx context.Context, params *CreatePaymentParams) (*models.CreateEwalletResponse, error) {
	result, err := c.paymentClient.PostPaymentV10PaymentHostToHost(
		&payment.PostPaymentV10PaymentHostToHostParams{
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
func (c *PaymentClient) Cancel(ctx context.Context, params *CancelPaymentParams) (*models.CancelHostToHostPaymentResponse, error) {
	result, err := c.paymentClient.PostPaymentV10PaymentHostToHostCancel(
		&payment.PostPaymentV10PaymentHostToHostCancelParams{
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

// GetStatus checks the status of a payment transaction
func (c *PaymentClient) GetStatus(ctx context.Context, params *CheckPaymentStatusParams) (*models.CheckPaymentStatusEwalletResponse, error) {
	result, err := c.paymentClient.PostPaymentV10PaymentHostToHostStatus(
		&payment.PostPaymentV10PaymentHostToHostStatusParams{
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

