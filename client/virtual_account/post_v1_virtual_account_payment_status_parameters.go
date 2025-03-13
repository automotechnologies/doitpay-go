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

	/* CHANNELID.

	   Channel ID
	*/
	CHANNELID string

	/* XEXTERNALID.

	   External ID
	*/
	XEXTERNALID string

	/* XSIGNATURE.

	   Request signature
	*/
	XSIGNATURE string

	/* XTIMESTAMP.

	   Request timestamp
	*/
	XTIMESTAMP string

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

// WithCHANNELID adds the cHANNELID to the post v1 virtual account payment status params
func (o *PostV1VirtualAccountPaymentStatusParams) WithCHANNELID(cHANNELID string) *PostV1VirtualAccountPaymentStatusParams {
	o.SetCHANNELID(cHANNELID)
	return o
}

// SetCHANNELID adds the cHANNELId to the post v1 virtual account payment status params
func (o *PostV1VirtualAccountPaymentStatusParams) SetCHANNELID(cHANNELID string) {
	o.CHANNELID = cHANNELID
}

// WithXEXTERNALID adds the xEXTERNALID to the post v1 virtual account payment status params
func (o *PostV1VirtualAccountPaymentStatusParams) WithXEXTERNALID(xEXTERNALID string) *PostV1VirtualAccountPaymentStatusParams {
	o.SetXEXTERNALID(xEXTERNALID)
	return o
}

// SetXEXTERNALID adds the xEXTERNALId to the post v1 virtual account payment status params
func (o *PostV1VirtualAccountPaymentStatusParams) SetXEXTERNALID(xEXTERNALID string) {
	o.XEXTERNALID = xEXTERNALID
}

// WithXSIGNATURE adds the xSIGNATURE to the post v1 virtual account payment status params
func (o *PostV1VirtualAccountPaymentStatusParams) WithXSIGNATURE(xSIGNATURE string) *PostV1VirtualAccountPaymentStatusParams {
	o.SetXSIGNATURE(xSIGNATURE)
	return o
}

// SetXSIGNATURE adds the xSIGNATURE to the post v1 virtual account payment status params
func (o *PostV1VirtualAccountPaymentStatusParams) SetXSIGNATURE(xSIGNATURE string) {
	o.XSIGNATURE = xSIGNATURE
}

// WithXTIMESTAMP adds the xTIMESTAMP to the post v1 virtual account payment status params
func (o *PostV1VirtualAccountPaymentStatusParams) WithXTIMESTAMP(xTIMESTAMP string) *PostV1VirtualAccountPaymentStatusParams {
	o.SetXTIMESTAMP(xTIMESTAMP)
	return o
}

// SetXTIMESTAMP adds the xTIMESTAMP to the post v1 virtual account payment status params
func (o *PostV1VirtualAccountPaymentStatusParams) SetXTIMESTAMP(xTIMESTAMP string) {
	o.XTIMESTAMP = xTIMESTAMP
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

	// header param CHANNEL-ID
	if err := r.SetHeaderParam("CHANNEL-ID", o.CHANNELID); err != nil {
		return err
	}

	// header param X-EXTERNAL-ID
	if err := r.SetHeaderParam("X-EXTERNAL-ID", o.XEXTERNALID); err != nil {
		return err
	}

	// header param X-SIGNATURE
	if err := r.SetHeaderParam("X-SIGNATURE", o.XSIGNATURE); err != nil {
		return err
	}

	// header param X-TIMESTAMP
	if err := r.SetHeaderParam("X-TIMESTAMP", o.XTIMESTAMP); err != nil {
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
