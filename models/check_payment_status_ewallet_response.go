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

// CheckPaymentStatusEwalletResponse check payment status ewallet response
//
// swagger:model CheckPaymentStatusEwalletResponse
type CheckPaymentStatusEwalletResponse struct {

	// additional info
	AdditionalInfo *HostToHostPaymentAdditionalInfo `json:"additionalInfo,omitempty"`

	// latest transaction status
	LatestTransactionStatus string `json:"latestTransactionStatus,omitempty"`

	// original partner reference no
	OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo,omitempty"`

	// response code
	ResponseCode string `json:"responseCode,omitempty"`

	// response message
	ResponseMessage string `json:"responseMessage,omitempty"`

	// service code
	ServiceCode string `json:"serviceCode,omitempty"`
}

// Validate validates this check payment status ewallet response
func (m *CheckPaymentStatusEwalletResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckPaymentStatusEwalletResponse) validateAdditionalInfo(formats strfmt.Registry) error {
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

// ContextValidate validate this check payment status ewallet response based on the context it is used
func (m *CheckPaymentStatusEwalletResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckPaymentStatusEwalletResponse) contextValidateAdditionalInfo(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *CheckPaymentStatusEwalletResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckPaymentStatusEwalletResponse) UnmarshalBinary(b []byte) error {
	var res CheckPaymentStatusEwalletResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
