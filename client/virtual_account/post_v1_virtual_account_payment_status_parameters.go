// Code generated by go-swagger; DO NOT EDIT.

package virtual_account

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

	"github.com/automotechnologies/doitpay-go/v2/models"
)

// NewPostV1VirtualAccountPaymentStatusParams creates a new PostV1VirtualAccountPaymentStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1VirtualAccountPaymentStatusParams() *PostV1VirtualAccountPaymentStatusParams {
	return &PostV1VirtualAccountPaymentStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1VirtualAccountPaymentStatusParamsWithTimeout creates a new PostV1VirtualAccountPaymentStatusParams object
// with the ability to set a timeout on a request.
func NewPostV1VirtualAccountPaymentStatusParamsWithTimeout(timeout time.Duration) *PostV1VirtualAccountPaymentStatusParams {
	return &PostV1VirtualAccountPaymentStatusParams{
		timeout: timeout,
	}
}

// NewPostV1VirtualAccountPaymentStatusParamsWithContext creates a new PostV1VirtualAccountPaymentStatusParams object
// with the ability to set a context for a request.
func NewPostV1VirtualAccountPaymentStatusParamsWithContext(ctx context.Context) *PostV1VirtualAccountPaymentStatusParams {
	return &PostV1VirtualAccountPaymentStatusParams{
		Context: ctx,
	}
}

// NewPostV1VirtualAccountPaymentStatusParamsWithHTTPClient creates a new PostV1VirtualAccountPaymentStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1VirtualAccountPaymentStatusParamsWithHTTPClient(client *http.Client) *PostV1VirtualAccountPaymentStatusParams {
	return &PostV1VirtualAccountPaymentStatusParams{
		HTTPClient: client,
	}
}

/*
PostV1VirtualAccountPaymentStatusParams contains all the parameters to send to the API endpoint

	for the post v1 virtual account payment status operation.

	Typically these are written to a http.Request.
*/
type PostV1VirtualAccountPaymentStatusParams struct {

	/* XCLIENTKEY.

	   Client Key
	*/
	XCLIENTKEY string

	/* Body.

	   Virtual Account Request
	*/
	Body *models.CheckVirtualAccountPaymentStatusRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 virtual account payment status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1VirtualAccountPaymentStatusParams) WithDefaults() *PostV1VirtualAccountPaymentStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 virtual account payment status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1VirtualAccountPaymentStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 virtual account payment status params
func (o *PostV1VirtualAccountPaymentStatusParams) WithTimeout(timeout time.Duration) *PostV1VirtualAccountPaymentStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 virtual account payment status params
func (o *PostV1VirtualAccountPaymentStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 virtual account payment status params
func (o *PostV1VirtualAccountPaymentStatusParams) WithContext(ctx context.Context) *PostV1VirtualAccountPaymentStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 virtual account payment status params
func (o *PostV1VirtualAccountPaymentStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 virtual account payment status params
func (o *PostV1VirtualAccountPaymentStatusParams) WithHTTPClient(client *http.Client) *PostV1VirtualAccountPaymentStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 virtual account payment status params
func (o *PostV1VirtualAccountPaymentStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCLIENTKEY adds the xCLIENTKEY to the post v1 virtual account payment status params
func (o *PostV1VirtualAccountPaymentStatusParams) WithXCLIENTKEY(xCLIENTKEY string) *PostV1VirtualAccountPaymentStatusParams {
	o.SetXCLIENTKEY(xCLIENTKEY)
	return o
}

// SetXCLIENTKEY adds the xCLIENTKEY to the post v1 virtual account payment status params
func (o *PostV1VirtualAccountPaymentStatusParams) SetXCLIENTKEY(xCLIENTKEY string) {
	o.XCLIENTKEY = xCLIENTKEY
}

// WithBody adds the body to the post v1 virtual account payment status params
func (o *PostV1VirtualAccountPaymentStatusParams) WithBody(body *models.CheckVirtualAccountPaymentStatusRequest) *PostV1VirtualAccountPaymentStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post v1 virtual account payment status params
func (o *PostV1VirtualAccountPaymentStatusParams) SetBody(body *models.CheckVirtualAccountPaymentStatusRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1VirtualAccountPaymentStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-CLIENT-KEY
	if err := r.SetHeaderParam("X-CLIENT-KEY", o.XCLIENTKEY); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
