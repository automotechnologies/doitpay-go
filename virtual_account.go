package doitpay

import (
	"context"

	"github.com/automotechnologies/doitpay-go/client/virtual_account"
	"github.com/automotechnologies/doitpay-go/models"
)

// Parameter structs
type CreateVAParams struct {
	Request    *models.CreateVirtualAccountRequest
	ExternalID string
	ChannelID  string
}

type UpdateVAParams struct {
	Request    *models.UpdateVirtualAccountRequest
	ExternalID string
	ChannelID  string
}

type DeleteVAParams struct {
	Request    *models.DeleteVirtualAccountRequest
	ExternalID string
	ChannelID  string
}

type InquiryVAParams struct {
	Request    *models.InquiryVirtualAccountRequest
	ExternalID string
	ChannelID  string
}

type PaymentStatusParams struct {
	Request    *models.CheckVirtualAccountPaymentStatusRequest
	ExternalID string
	ChannelID  string
}

type VirtualAccountClient struct {
	virtualAccountClient virtual_account.ClientService
}

func NewVirtualAccountClient(clientService virtual_account.ClientService) *VirtualAccountClient {
	return &VirtualAccountClient{
		virtualAccountClient: clientService,
	}
}

// Create creates a new virtual account
func (c *VirtualAccountClient) Create(ctx context.Context, params *CreateVAParams) (*models.CreateVirtualAccountResponse, error) {
	result, err := c.virtualAccountClient.PostV1VirtualAccount(
		&virtual_account.PostV1VirtualAccountParams{
			Body:        params.Request,
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

// Update updates an existing virtual account
func (c *VirtualAccountClient) Update(ctx context.Context, params *UpdateVAParams) (*models.UpdateVirtualAccountResponse, error) {
	result, err := c.virtualAccountClient.PutV1VirtualAccount(
		&virtual_account.PutV1VirtualAccountParams{
			Body:        params.Request,
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

// Delete deletes an existing virtual account
func (c *VirtualAccountClient) Delete(ctx context.Context, params *DeleteVAParams) (*models.DeleteVirtualAccountResponse, error) {
	result, err := c.virtualAccountClient.DeleteV1VirtualAccount(
		&virtual_account.DeleteV1VirtualAccountParams{
			Body:        params.Request,
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

// Inquiry gets details of an existing virtual account
func (c *VirtualAccountClient) Inquiry(ctx context.Context, params *InquiryVAParams) (*models.InquiryVirtualAccountResponse, error) {
	result, err := c.virtualAccountClient.PostV1VirtualAccountInquiry(
		&virtual_account.PostV1VirtualAccountInquiryParams{
			Body:        params.Request,
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

// GetPaymentStatus checks the payment status of a virtual account
func (c *VirtualAccountClient) GetPaymentStatus(ctx context.Context, params *PaymentStatusParams) (*models.CheckVirtualAccountPaymentStatusResponse, error) {
	result, err := c.virtualAccountClient.PostV1VirtualAccountPaymentStatus(
		&virtual_account.PostV1VirtualAccountPaymentStatusParams{
			Body:        params.Request,
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
