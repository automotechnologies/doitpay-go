// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateEwalletRequest create ewallet request
//
// swagger:model CreateEwalletRequest
type CreateEwalletRequest struct {

	// additional info
	AdditionalInfo *CreateEwalletRequestAdditionalInfo `json:"additionalInfo,omitempty"`

	// amount
	Amount *AmountDetail `json:"amount,omitempty"`

	// partner reference no
	PartnerReferenceNo string `json:"partnerReferenceNo,omitempty"`

	// point of initiation
	PointOfInitiation string `json:"pointOfInitiation,omitempty"`

	// url param
	URLParam []*URLParam `json:"urlParam"`

	// valid up to
	ValidUpTo string `json:"validUpTo,omitempty"`
}

// Validate validates this create ewallet request
func (m *CreateEwalletRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURLParam(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateEwalletRequest) validateAdditionalInfo(formats strfmt.Registry) error {
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

func (m *CreateEwalletRequest) validateAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	if m.Amount != nil {
		if err := m.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("amount")
			}
			return err
		}
	}

	return nil
}

func (m *CreateEwalletRequest) validateURLParam(formats strfmt.Registry) error {
	if swag.IsZero(m.URLParam) { // not required
		return nil
	}

	for i := 0; i < len(m.URLParam); i++ {
		if swag.IsZero(m.URLParam[i]) { // not required
			continue
		}

		if m.URLParam[i] != nil {
			if err := m.URLParam[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("urlParam" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("urlParam" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create ewallet request based on the context it is used
func (m *CreateEwalletRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURLParam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateEwalletRequest) contextValidateAdditionalInfo(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CreateEwalletRequest) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.Amount != nil {

		if swag.IsZero(m.Amount) { // not required
			return nil
		}

		if err := m.Amount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("amount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("amount")
			}
			return err
		}
	}

	return nil
}

func (m *CreateEwalletRequest) contextValidateURLParam(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.URLParam); i++ {

		if m.URLParam[i] != nil {

			if swag.IsZero(m.URLParam[i]) { // not required
				return nil
			}

			if err := m.URLParam[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("urlParam" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("urlParam" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateEwalletRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateEwalletRequest) UnmarshalBinary(b []byte) error {
	var res CreateEwalletRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
