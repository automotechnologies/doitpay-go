// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateEwalletResponseAdditionalInfo create ewallet response additional info
//
// swagger:model CreateEwalletResponseAdditionalInfo
type CreateEwalletResponseAdditionalInfo struct {

	// channel code
	ChannelCode string `json:"channelCode,omitempty"`

	// expired time
	ExpiredTime string `json:"expiredTime,omitempty"`

	// original reference no
	OriginalReferenceNo string `json:"originalReferenceNo,omitempty"`
}

// Validate validates this create ewallet response additional info
func (m *CreateEwalletResponseAdditionalInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create ewallet response additional info based on context it is used
func (m *CreateEwalletResponseAdditionalInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateEwalletResponseAdditionalInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateEwalletResponseAdditionalInfo) UnmarshalBinary(b []byte) error {
	var res CreateEwalletResponseAdditionalInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
