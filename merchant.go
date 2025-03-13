package doitpay

import (
	"context"

	"github.com/automotechnologies/doitpay-go/v2/client/merchants"
	"github.com/automotechnologies/doitpay-go/v2/models"
)

type MerchantClient struct {
	merchantClient merchants.ClientService
}


func NewMerchantClient(clientService merchants.ClientService) *MerchantClient {
	return &MerchantClient{
		merchantClient: clientService,
	}
}

type GetMerchantParams struct {
	MerchantID string
	ChannelID  string
	ExternalID string
	Limit      int64
	Page       int64
	Search     string
}

func (c *MerchantClient) GetMerchants(ctx context.Context, params *GetMerchantParams) (*models.PaginatedMerchantResponse, error) {
	defaultLimit := int64(50)
	defaultPage := int64(1)

	if params.Limit == 0 {
		params.Limit = defaultLimit
	}
	if params.Page == 0 {
		params.Page = defaultPage
	}

	result, err := c.merchantClient.GetV1Merchant(
		&merchants.GetV1MerchantParams{
			CHANNELID:   params.ChannelID,
			XEXTERNALID: params.ExternalID,
			Limit:       &params.Limit,
			Page:        &params.Page,
			Search:      &params.Search,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}

type GetMerchantByRefParams struct {
	MerchantRef string
	ChannelID   string
	ExternalID  string
}

func (c *MerchantClient) GetMerchantByRef(ctx context.Context, params *GetMerchantByRefParams) (*models.MerchantResponse, error) {
	result, err := c.merchantClient.GetV1MerchantMerchantRef(
		&merchants.GetV1MerchantMerchantRefParams{
			XEXTERNALID: params.ExternalID,
			CHANNELID:   params.ChannelID,
			MerchantRef: params.MerchantRef,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}

type CreateMerchantParams struct {
	MerchantRef string
	ChannelID   string
	ExternalID  string
	BusinessName string
	Name         string
}

func (c *MerchantClient) CreateMerchant(ctx context.Context, params *CreateMerchantParams) (*models.MerchantResponse, error) {
	result, err := c.merchantClient.PostV1Merchant(
		&merchants.PostV1MerchantParams{
			Request: &models.CreateMerchantRequest{
				BusinessName: &params.BusinessName,
				Name:         &params.Name,
			},
			CHANNELID:   params.ChannelID,
			XEXTERNALID: params.ExternalID,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}
