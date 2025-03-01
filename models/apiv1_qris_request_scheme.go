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

// Apiv1QrisRequestScheme apiv1 qris request scheme
//
// swagger:model apiv1.QrisRequestScheme
type Apiv1QrisRequestScheme struct {

	// additional info
	AdditionalInfo map[string]string `json:"additionalInfo,omitempty"`

	// amount
	Amount *Apiv1QrisAmount `json:"amount,omitempty"`

	// Merchant identifier that is unique per each merchant (MPAN)
	MerchantID string `json:"merchantId,omitempty"`

	// partner reference no
	PartnerReferenceNo string `json:"partnerReferenceNo,omitempty"`

	// StoreId this should be defined merchant to PG (NMID)
	StoreID string `json:"storeId,omitempty"`

	// SubMerchantId Sub merchant id (MID)
	SubMerchantID string `json:"subMerchantId,omitempty"`

	// validity period
	ValidityPeriod string `json:"validityPeriod,omitempty"`
}

// Validate validates this apiv1 qris request scheme
func (m *Apiv1QrisRequestScheme) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Apiv1QrisRequestScheme) validateAmount(formats strfmt.Registry) error {
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

// ContextValidate validate this apiv1 qris request scheme based on the context it is used
func (m *Apiv1QrisRequestScheme) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Apiv1QrisRequestScheme) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Apiv1QrisRequestScheme) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Apiv1QrisRequestScheme) UnmarshalBinary(b []byte) error {
	var res Apiv1QrisRequestScheme
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
