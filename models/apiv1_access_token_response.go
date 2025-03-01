// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Apiv1AccessTokenResponse apiv1 access token response
//
// swagger:model apiv1.AccessTokenResponse
type Apiv1AccessTokenResponse struct {

	// access token
	AccessToken string `json:"accessToken,omitempty"`

	// additional data
	AdditionalData string `json:"additionalData,omitempty"`

	// expires in
	ExpiresIn string `json:"expiresIn,omitempty"`

	// response code
	ResponseCode string `json:"responseCode,omitempty"`

	// response message
	ResponseMessage string `json:"responseMessage,omitempty"`

	// response status
	ResponseStatus string `json:"responseStatus,omitempty"`

	// token type
	TokenType string `json:"tokenType,omitempty"`
}

// Validate validates this apiv1 access token response
func (m *Apiv1AccessTokenResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this apiv1 access token response based on context it is used
func (m *Apiv1AccessTokenResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Apiv1AccessTokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Apiv1AccessTokenResponse) UnmarshalBinary(b []byte) error {
	var res Apiv1AccessTokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
