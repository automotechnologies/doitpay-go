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

// Apiv1PaginatedMerchantResponse apiv1 paginated merchant response
//
// swagger:model apiv1.PaginatedMerchantResponse
type Apiv1PaginatedMerchantResponse struct {

	// List of merchant data
	Data []*Apiv1MerchantResponse `json:"data"`

	// pagination
	Pagination *Apiv1PaginatedMerchantResponsePagination `json:"pagination,omitempty"`
}

// Validate validates this apiv1 paginated merchant response
func (m *Apiv1PaginatedMerchantResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Apiv1PaginatedMerchantResponse) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Apiv1PaginatedMerchantResponse) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this apiv1 paginated merchant response based on the context it is used
func (m *Apiv1PaginatedMerchantResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Apiv1PaginatedMerchantResponse) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Data); i++ {

		if m.Data[i] != nil {

			if swag.IsZero(m.Data[i]) { // not required
				return nil
			}

			if err := m.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Apiv1PaginatedMerchantResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {

		if swag.IsZero(m.Pagination) { // not required
			return nil
		}

		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Apiv1PaginatedMerchantResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Apiv1PaginatedMerchantResponse) UnmarshalBinary(b []byte) error {
	var res Apiv1PaginatedMerchantResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Apiv1PaginatedMerchantResponsePagination Pagination information
//
// swagger:model Apiv1PaginatedMerchantResponsePagination
type Apiv1PaginatedMerchantResponsePagination struct {

	// Current page number
	CurrentPage int64 `json:"current_page,omitempty"`

	// Number of items per page
	Limit int64 `json:"limit,omitempty"`

	// Total number of items
	TotalItems int64 `json:"total_items,omitempty"`

	// Total number of pages
	TotalPages int64 `json:"total_pages,omitempty"`
}

// Validate validates this apiv1 paginated merchant response pagination
func (m *Apiv1PaginatedMerchantResponsePagination) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this apiv1 paginated merchant response pagination based on context it is used
func (m *Apiv1PaginatedMerchantResponsePagination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Apiv1PaginatedMerchantResponsePagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Apiv1PaginatedMerchantResponsePagination) UnmarshalBinary(b []byte) error {
	var res Apiv1PaginatedMerchantResponsePagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
