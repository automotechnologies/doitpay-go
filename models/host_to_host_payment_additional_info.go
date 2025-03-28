// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HostToHostPaymentAdditionalInfo host to host payment additional info
//
// swagger:model HostToHostPaymentAdditionalInfo
type HostToHostPaymentAdditionalInfo struct {

	// channel
	Channel string `json:"channel,omitempty"`

	// channel code
	ChannelCode string `json:"channelCode,omitempty"`

	// contract Id
	ContractID string `json:"contractId,omitempty"`

	// customer email
	CustomerEmail string `json:"customerEmail,omitempty"`

	// customer name
	CustomerName string `json:"customerName,omitempty"`

	// customer phone
	CustomerPhone string `json:"customerPhone,omitempty"`

	// expired time
	ExpiredTime string `json:"expiredTime,omitempty"`

	// latest transaction status
	LatestTransactionStatus string `json:"latestTransactionStatus,omitempty"`

	// original reference no
	OriginalReferenceNo string `json:"originalReferenceNo,omitempty"`
}

// Validate validates this host to host payment additional info
func (m *HostToHostPaymentAdditionalInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this host to host payment additional info based on context it is used
func (m *HostToHostPaymentAdditionalInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HostToHostPaymentAdditionalInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostToHostPaymentAdditionalInfo) UnmarshalBinary(b []byte) error {
	var res HostToHostPaymentAdditionalInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
