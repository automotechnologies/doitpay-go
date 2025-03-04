// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteVirtualAccountRequest delete virtual account request
//
// swagger:model DeleteVirtualAccountRequest
type DeleteVirtualAccountRequest struct {

	// customer no
	CustomerNo string `json:"customerNo,omitempty"`

	// partner service Id
	PartnerServiceID string `json:"partnerServiceId,omitempty"`

	// trx Id
	TrxID string `json:"trxId,omitempty"`

	// virtual account no
	VirtualAccountNo string `json:"virtualAccountNo,omitempty"`
}

// Validate validates this delete virtual account request
func (m *DeleteVirtualAccountRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete virtual account request based on context it is used
func (m *DeleteVirtualAccountRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteVirtualAccountRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteVirtualAccountRequest) UnmarshalBinary(b []byte) error {
	var res DeleteVirtualAccountRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
