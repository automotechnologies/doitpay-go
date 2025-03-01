// Code generated by go-swagger; DO NOT EDIT.

package merchants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/automotechnologies/doitpay-go/models"
)

// NewPostV1MerchantParams creates a new PostV1MerchantParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1MerchantParams() *PostV1MerchantParams {
	return &PostV1MerchantParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1MerchantParamsWithTimeout creates a new PostV1MerchantParams object
// with the ability to set a timeout on a request.
func NewPostV1MerchantParamsWithTimeout(timeout time.Duration) *PostV1MerchantParams {
	return &PostV1MerchantParams{
		timeout: timeout,
	}
}

// NewPostV1MerchantParamsWithContext creates a new PostV1MerchantParams object
// with the ability to set a context for a request.
func NewPostV1MerchantParamsWithContext(ctx context.Context) *PostV1MerchantParams {
	return &PostV1MerchantParams{
		Context: ctx,
	}
}

// NewPostV1MerchantParamsWithHTTPClient creates a new PostV1MerchantParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1MerchantParamsWithHTTPClient(client *http.Client) *PostV1MerchantParams {
	return &PostV1MerchantParams{
		HTTPClient: client,
	}
}

/*
PostV1MerchantParams contains all the parameters to send to the API endpoint

	for the post v1 merchant operation.

	Typically these are written to a http.Request.
*/
type PostV1MerchantParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* Request.

	   Merchant creation request
	*/
	Request *models.Apiv1CreateMerchantRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 merchant params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1MerchantParams) WithDefaults() *PostV1MerchantParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 merchant params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1MerchantParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 merchant params
func (o *PostV1MerchantParams) WithTimeout(timeout time.Duration) *PostV1MerchantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 merchant params
func (o *PostV1MerchantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 merchant params
func (o *PostV1MerchantParams) WithContext(ctx context.Context) *PostV1MerchantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 merchant params
func (o *PostV1MerchantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 merchant params
func (o *PostV1MerchantParams) WithHTTPClient(client *http.Client) *PostV1MerchantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 merchant params
func (o *PostV1MerchantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post v1 merchant params
func (o *PostV1MerchantParams) WithAuthorization(authorization string) *PostV1MerchantParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post v1 merchant params
func (o *PostV1MerchantParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithRequest adds the request to the post v1 merchant params
func (o *PostV1MerchantParams) WithRequest(request *models.Apiv1CreateMerchantRequest) *PostV1MerchantParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post v1 merchant params
func (o *PostV1MerchantParams) SetRequest(request *models.Apiv1CreateMerchantRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1MerchantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
