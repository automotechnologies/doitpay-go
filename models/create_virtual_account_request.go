// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateVirtualAccountRequest create virtual account request
//
// swagger:model CreateVirtualAccountRequest
type CreateVirtualAccountRequest struct {

	// Object with more details
	AdditionalInfo struct {
		GenericVirtualAccountAdditionalInfo
	} `json:"additionalInfo,omitempty"`

	// Nomor untuk VA
	CustomerNo string `json:"customerNo,omitempty"`

	// ISO-8601
	ExpiredDate string `json:"expiredDate,omitempty"`

	// Prefix dari Duitku
	PartnerServiceID string `json:"partnerServiceId,omitempty"`

	// Object with more details
	TotalAmount struct {
		AmountDetail
	} `json:"totalAmount,omitempty"`

	// ID transaksi
	TrxID string `json:"trxId,omitempty"`

	// Nama akan ditampilkan di sisi bank
	VirtualAccountName string `json:"virtualAccountName,omitempty"`

	// PartnerServiceId + CustomerNo
	VirtualAccountNo string `json:"virtualAccountNo,omitempty"`

	// Close Amount : C, Open Amount : O
	VirtualAccountTrxType string `json:"virtualAccountTrxType,omitempty"`
}

// Validate validates this create virtual account request
func (m *CreateVirtualAccountRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateVirtualAccountRequest) validateAdditionalInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalInfo) { // not required
		return nil
	}

	return nil
}

func (m *CreateVirtualAccountRequest) validateTotalAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.TotalAmount) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this create virtual account request based on the context it is used
func (m *CreateVirtualAccountRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateVirtualAccountRequest) contextValidateAdditionalInfo(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *CreateVirtualAccountRequest) contextValidateTotalAmount(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *CreateVirtualAccountRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateVirtualAccountRequest) UnmarshalBinary(b []byte) error {
	var res CreateVirtualAccountRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
