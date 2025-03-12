package doitpay

import (
	"context"

	"github.com/automotechnologies/doitpay-go/v2/client/disbursement"
	"github.com/automotechnologies/doitpay-go/v2/models"
)

// Parameter structs
type DisbursementParams struct {
	Request    *models.CreateDisbursementRequest
	ExternalID string
	ChannelID  string
}

type EMoneyDisbursementParams struct {
	Request    *models.CreateEwalletTopupRequest
	ExternalID string
	ChannelID  string
}

type BankAccountValidationParams struct {
	Request    *models.BankAccountValidationRequest
	ExternalID string
	ChannelID  string
}

type EwalletAccountValidationParams struct {
	Request    *models.EmoneyAccountValidationRequest
	ExternalID string
	ChannelID  string
}

type GetStatusParams struct {
	Request    *models.InquiryDisbursementStatusRequest
	ExternalID string
	ChannelID  string
}

type GetStatusEMoneyParams struct {
	Request    *models.InquiryEwalletTopupStatusRequest
	ExternalID string
	ChannelID  string
}

type DisbursementClient struct {
	disbursementClient disbursement.ClientService
}

func NewDisbursementClient(clientService disbursement.ClientService) *DisbursementClient {
	return &DisbursementClient{
		disbursementClient: clientService,
	}
}

func (c *DisbursementClient) Create(ctx context.Context, params *DisbursementParams) (*models.CreateDisbursementResponse, error) {
	result, err := c.disbursementClient.PostDisbursementV10BankDisbursement(
		&disbursement.PostDisbursementV10BankDisbursementParams{
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

func (c *DisbursementClient) CreateEMoney(ctx context.Context, params *EMoneyDisbursementParams) (*models.CreateEwalletTopupResponse, error) {
	result, err := c.disbursementClient.PostDisbursementV10EmoneyDisbursement(
		&disbursement.PostDisbursementV10EmoneyDisbursementParams{
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

func (c *DisbursementClient) ValidateBankAccount(ctx context.Context, params *BankAccountValidationParams) (*models.BankAccountValidationResponse, error) {
	result, err := c.disbursementClient.PostDisbursementV10BankAccountValidation(
		&disbursement.PostDisbursementV10BankAccountValidationParams{
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

func (c *DisbursementClient) ValidateEwalletAccount(ctx context.Context, params *EwalletAccountValidationParams) (*models.EmoneyAccountValidationResponse, error) {
	result, err := c.disbursementClient.PostDisbursementV10EmoneyAccountValidation(
		&disbursement.PostDisbursementV10EmoneyAccountValidationParams{
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

func (c *DisbursementClient) GetStatus(ctx context.Context, params *GetStatusParams) (*models.InquiryDisbursementStatusResponse, error) {
	result, err := c.disbursementClient.PostDisbursementV10BankDisbursementStatus(
		&disbursement.PostDisbursementV10BankDisbursementStatusParams{
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

func (c *DisbursementClient) GetStatusEMoney(ctx context.Context, params *GetStatusEMoneyParams) (*models.InquiryEwalletTopupStatusResponse, error) {
	result, err := c.disbursementClient.PostDisbursementV10EmoneyDisbursementStatus(
		&disbursement.PostDisbursementV10EmoneyDisbursementStatusParams{
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

