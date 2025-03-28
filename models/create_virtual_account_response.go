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

// CreateVirtualAccountResponse create virtual account response
//
// swagger:model CreateVirtualAccountResponse
type CreateVirtualAccountResponse struct {

	// response code
	ResponseCode string `json:"responseCode,omitempty"`

	// response message
	ResponseMessage string `json:"responseMessage,omitempty"`

	// virtual account data
	VirtualAccountData *GenericVirtualAccountData `json:"virtualAccountData,omitempty"`
}

// Validate validates this create virtual account response
func (m *CreateVirtualAccountResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVirtualAccountData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateVirtualAccountResponse) validateVirtualAccountData(formats strfmt.Registry) error {
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

// ContextValidate validate this create virtual account response based on the context it is used
func (m *CreateVirtualAccountResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVirtualAccountData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateVirtualAccountResponse) contextValidateVirtualAccountData(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CreateVirtualAccountResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateVirtualAccountResponse) UnmarshalBinary(b []byte) error {
	var res CreateVirtualAccountResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
