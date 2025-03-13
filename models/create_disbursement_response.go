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

// CreateDisbursementResponse create disbursement response
//
// swagger:model CreateDisbursementResponse
type CreateDisbursementResponse struct {

	// additional info
	AdditionalInfo *DisbursementAdditionalData `json:"additionalInfo,omitempty"`

	// amount
	Amount *AmountDetail `json:"amount,omitempty"`

	// beneficiary account no
	BeneficiaryAccountNo string `json:"beneficiaryAccountNo,omitempty"`

	// beneficiary bank code
	BeneficiaryBankCode string `json:"beneficiaryBankCode,omitempty"`

	// partner reference no
	PartnerReferenceNo string `json:"partnerReferenceNo,omitempty"`

	// reference no
	ReferenceNo string `json:"referenceNo,omitempty"`

	// response code
	ResponseCode string `json:"responseCode,omitempty"`

	// response message
	ResponseMessage string `json:"responseMessage,omitempty"`
}

// Validate validates this create disbursement response
func (m *CreateDisbursementResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateDisbursementResponse) validateAdditionalInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalInfo) { // not required
		return nil
	}

	if m.AdditionalInfo != nil {
		if err := m.AdditionalInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("additionalInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("additionalInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CreateDisbursementResponse) validateAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	if m.Amount != nil {
		if err := m.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("amount")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create disbursement response based on the context it is used
func (m *CreateDisbursementResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateDisbursementResponse) contextValidateAdditionalInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.AdditionalInfo != nil {

		if swag.IsZero(m.AdditionalInfo) { // not required
			return nil
		}

		if err := m.AdditionalInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("additionalInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("additionalInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CreateDisbursementResponse) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.Amount != nil {

		if swag.IsZero(m.Amount) { // not required
			return nil
		}

		if err := m.Amount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("amount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateDisbursementResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateDisbursementResponse) UnmarshalBinary(b []byte) error {
	var res CreateDisbursementResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
