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

// Apiv1AccountInfo apiv1 account info
//
// swagger:model apiv1.AccountInfo
type Apiv1AccountInfo struct {

	// available balance
	AvailableBalance *Apiv1BalanceDetail `json:"availableBalance,omitempty"`
}

// Validate validates this apiv1 account info
func (m *Apiv1AccountInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableBalance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Apiv1AccountInfo) validateAvailableBalance(formats strfmt.Registry) error {
	if swag.IsZero(m.AvailableBalance) { // not required
		return nil
	}

	if m.AvailableBalance != nil {
		if err := m.AvailableBalance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("availableBalance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("availableBalance")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this apiv1 account info based on the context it is used
func (m *Apiv1AccountInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAvailableBalance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Apiv1AccountInfo) contextValidateAvailableBalance(ctx context.Context, formats strfmt.Registry) error {

	if m.AvailableBalance != nil {

		if swag.IsZero(m.AvailableBalance) { // not required
			return nil
		}

		if err := m.AvailableBalance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("availableBalance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("availableBalance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Apiv1AccountInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Apiv1AccountInfo) UnmarshalBinary(b []byte) error {
	var res Apiv1AccountInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
