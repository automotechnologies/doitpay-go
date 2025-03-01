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

// Apiv1UpdateVirtualAccountResponse apiv1 update virtual account response
//
// swagger:model apiv1.UpdateVirtualAccountResponse
type Apiv1UpdateVirtualAccountResponse struct {

	// response code
	ResponseCode string `json:"responseCode,omitempty"`

	// response message
	ResponseMessage string `json:"responseMessage,omitempty"`

	// virtual account data
	VirtualAccountData *Apiv1GenericVirtualAccountData `json:"virtualAccountData,omitempty"`
}

// Validate validates this apiv1 update virtual account response
func (m *Apiv1UpdateVirtualAccountResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVirtualAccountData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Apiv1UpdateVirtualAccountResponse) validateVirtualAccountData(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualAccountData) { // not required
		return nil
	}

	if m.VirtualAccountData != nil {
		if err := m.VirtualAccountData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualAccountData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtualAccountData")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this apiv1 update virtual account response based on the context it is used
func (m *Apiv1UpdateVirtualAccountResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVirtualAccountData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Apiv1UpdateVirtualAccountResponse) contextValidateVirtualAccountData(ctx context.Context, formats strfmt.Registry) error {

	if m.VirtualAccountData != nil {

		if swag.IsZero(m.VirtualAccountData) { // not required
			return nil
		}

		if err := m.VirtualAccountData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualAccountData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtualAccountData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Apiv1UpdateVirtualAccountResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Apiv1UpdateVirtualAccountResponse) UnmarshalBinary(b []byte) error {
	var res Apiv1UpdateVirtualAccountResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
