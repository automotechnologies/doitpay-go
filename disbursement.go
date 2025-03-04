package doitpay

import (
	"context"

	"github.com/automotechnologies/doitpay-go/v2/client/disbursement"
	"github.com/automotechnologies/doitpay-go/v2/models"
)


type DisbursementClient struct {
	disbursementClient disbursement.ClientService
}



func NewDisbursementClient(clientService disbursement.ClientService) *DisbursementClient {
	return &DisbursementClient{
		disbursementClient: clientService,
	}
}


func (c *DisbursementClient) Create(ctx context.Context, request *models.CreateDisbursementRequest) (*models.CreateDisbursementResponse, error) {
	result, err := c.disbursementClient.PostDisbursementV10BankDisbursement(
		&disbursement.PostDisbursementV10BankDisbursementParams{
			Request: request,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}

func (c *DisbursementClient) CreateEMoney(ctx context.Context, request *models.CreateEwalletTopupRequest) (*models.CreateEwalletTopupResponse, error) {
	result, err := c.disbursementClient.PostDisbursementV10EmoneyDisbursement(
		&disbursement.PostDisbursementV10EmoneyDisbursementParams{
			Request: request,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}

func (c *DisbursementClient) ValidateBankAccount(ctx context.Context, request *models.BankAccountValidationRequest) (*models.BankAccountValidationResponse, error) {
	result, err := c.disbursementClient.PostDisbursementV10BankAccountValidation(
		&disbursement.PostDisbursementV10BankAccountValidationParams{
			Request: request,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}


func (c *DisbursementClient) ValidateEwalletAccount(ctx context.Context, request *models.EmoneyAccountValidationRequest) (*models.EmoneyAccountValidationResponse, error) {
	result, err := c.disbursementClient.PostDisbursementV10EmoneyAccountValidation(
		&disbursement.PostDisbursementV10EmoneyAccountValidationParams{
			Request: request,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}

func (c *DisbursementClient) GetStatus(ctx context.Context, request *models.InquiryDisbursementStatusRequest) (*models.InquiryDisbursementStatusResponse, error) {
	result, err := c.disbursementClient.PostDisbursementV10BankDisbursementStatus(
		&disbursement.PostDisbursementV10BankDisbursementStatusParams{
			Request: request,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}

func (c *DisbursementClient) GetStatusEMoney(ctx context.Context, request *models.InquiryEwalletTopupStatusRequest) (*models.InquiryEwalletTopupStatusResponse, error) {
	result, err := c.disbursementClient.PostDisbursementV10EmoneyDisbursementStatus(
		&disbursement.PostDisbursementV10EmoneyDisbursementStatusParams{
			Request: request,
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}

