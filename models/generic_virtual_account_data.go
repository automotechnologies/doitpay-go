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

// GenericVirtualAccountData generic virtual account data
//
// swagger:model GenericVirtualAccountData
type GenericVirtualAccountData struct {

	// additional info
	AdditionalInfo *GenericVirtualAccountAdditionalInfo `json:"additionalInfo,omitempty"`

	// customer no
	CustomerNo string `json:"customerNo,omitempty"`

	// expired date
	ExpiredDate string `json:"expiredDate,omitempty"`

	// partner service Id
	PartnerServiceID string `json:"partnerServiceId,omitempty"`

	// total amount
	TotalAmount *AmountDetail `json:"totalAmount,omitempty"`

	// trx Id
	TrxID string `json:"trxId,omitempty"`

	// virtual account name
	VirtualAccountName string `json:"virtualAccountName,omitempty"`

	// virtual account no
	VirtualAccountNo string `json:"virtualAccountNo,omitempty"`

	// virtual account trx type
	VirtualAccountTrxType string `json:"virtualAccountTrxType,omitempty"`
}

// Validate validates this generic virtual account data
func (m *GenericVirtualAccountData) Validate(formats strfmt.Registry) error {
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

func (m *GenericVirtualAccountData) validateAdditionalInfo(formats strfmt.Registry) error {
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

func (m *GenericVirtualAccountData) validateTotalAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.TotalAmount) { // not required
		return nil
	}

	if m.TotalAmount != nil {
		if err := m.TotalAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("totalAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("totalAmount")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this generic virtual account data based on the context it is used
func (m *GenericVirtualAccountData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GenericVirtualAccountData) contextValidateAdditionalInfo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GenericVirtualAccountData) contextValidateTotalAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.TotalAmount != nil {

		if swag.IsZero(m.TotalAmount) { // not required
			return nil
		}

		if err := m.TotalAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("totalAmount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("totalAmount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GenericVirtualAccountData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericVirtualAccountData) UnmarshalBinary(b []byte) error {
	var res GenericVirtualAccountData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
