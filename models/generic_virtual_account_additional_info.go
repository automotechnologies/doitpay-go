// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GenericVirtualAccountAdditionalInfo generic virtual account additional info
//
// swagger:model GenericVirtualAccountAdditionalInfo
type GenericVirtualAccountAdditionalInfo struct {

	// max amount
	MaxAmount string `json:"maxAmount,omitempty"`

	// min amount
	MinAmount string `json:"minAmount,omitempty"`
}

// Validate validates this generic virtual account additional info
func (m *GenericVirtualAccountAdditionalInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this generic virtual account additional info based on context it is used
func (m *GenericVirtualAccountAdditionalInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GenericVirtualAccountAdditionalInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericVirtualAccountAdditionalInfo) UnmarshalBinary(b []byte) error {
	var res GenericVirtualAccountAdditionalInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
