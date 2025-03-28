// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateMerchantRequest create merchant request
//
// swagger:model CreateMerchantRequest
type CreateMerchantRequest struct {

	// Business name of the merchant
	// required: true
	// Required: true
	BusinessName *string `json:"business_name"`

	// Name of the merchant
	// required: true
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this create merchant request
func (m *CreateMerchantRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBusinessName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateMerchantRequest) validateBusinessName(formats strfmt.Registry) error {

	if err := validate.Required("business_name", "body", m.BusinessName); err != nil {
		return err
	}

	return nil
}

func (m *CreateMerchantRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create merchant request based on context it is used
func (m *CreateMerchantRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateMerchantRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateMerchantRequest) UnmarshalBinary(b []byte) error {
	var res CreateMerchantRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
