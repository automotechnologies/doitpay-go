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

	"github.com/automotechnologies/doitpay-go/models"
)

// NewPostV1VirtualAccountParams creates a new PostV1VirtualAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1VirtualAccountParams() *PostV1VirtualAccountParams {
	return &PostV1VirtualAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1VirtualAccountParamsWithTimeout creates a new PostV1VirtualAccountParams object
// with the ability to set a timeout on a request.
func NewPostV1VirtualAccountParamsWithTimeout(timeout time.Duration) *PostV1VirtualAccountParams {
	return &PostV1VirtualAccountParams{
		timeout: timeout,
	}
}

// NewPostV1VirtualAccountParamsWithContext creates a new PostV1VirtualAccountParams object
// with the ability to set a context for a request.
func NewPostV1VirtualAccountParamsWithContext(ctx context.Context) *PostV1VirtualAccountParams {
	return &PostV1VirtualAccountParams{
		Context: ctx,
	}
}

// NewPostV1VirtualAccountParamsWithHTTPClient creates a new PostV1VirtualAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1VirtualAccountParamsWithHTTPClient(client *http.Client) *PostV1VirtualAccountParams {
	return &PostV1VirtualAccountParams{
		HTTPClient: client,
	}
}

/*
PostV1VirtualAccountParams contains all the parameters to send to the API endpoint

	for the post v1 virtual account operation.

	Typically these are written to a http.Request.
*/
type PostV1VirtualAccountParams struct {

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
	Body *models.CreateVirtualAccountRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 virtual account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1VirtualAccountParams) WithDefaults() *PostV1VirtualAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 virtual account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1VirtualAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 virtual account params
func (o *PostV1VirtualAccountParams) WithTimeout(timeout time.Duration) *PostV1VirtualAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 virtual account params
func (o *PostV1VirtualAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 virtual account params
func (o *PostV1VirtualAccountParams) WithContext(ctx context.Context) *PostV1VirtualAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 virtual account params
func (o *PostV1VirtualAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 virtual account params
func (o *PostV1VirtualAccountParams) WithHTTPClient(client *http.Client) *PostV1VirtualAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 virtual account params
func (o *PostV1VirtualAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCHANNELID adds the cHANNELID to the post v1 virtual account params
func (o *PostV1VirtualAccountParams) WithCHANNELID(cHANNELID string) *PostV1VirtualAccountParams {
	o.SetCHANNELID(cHANNELID)
	return o
}

// SetCHANNELID adds the cHANNELId to the post v1 virtual account params
func (o *PostV1VirtualAccountParams) SetCHANNELID(cHANNELID string) {
	o.CHANNELID = cHANNELID
}

// WithXEXTERNALID adds the xEXTERNALID to the post v1 virtual account params
func (o *PostV1VirtualAccountParams) WithXEXTERNALID(xEXTERNALID string) *PostV1VirtualAccountParams {
	o.SetXEXTERNALID(xEXTERNALID)
	return o
}

// SetXEXTERNALID adds the xEXTERNALId to the post v1 virtual account params
func (o *PostV1VirtualAccountParams) SetXEXTERNALID(xEXTERNALID string) {
	o.XEXTERNALID = xEXTERNALID
}

// WithXSIGNATURE adds the xSIGNATURE to the post v1 virtual account params
func (o *PostV1VirtualAccountParams) WithXSIGNATURE(xSIGNATURE string) *PostV1VirtualAccountParams {
	o.SetXSIGNATURE(xSIGNATURE)
	return o
}

// SetXSIGNATURE adds the xSIGNATURE to the post v1 virtual account params
func (o *PostV1VirtualAccountParams) SetXSIGNATURE(xSIGNATURE string) {
	o.XSIGNATURE = xSIGNATURE
}

// WithXTIMESTAMP adds the xTIMESTAMP to the post v1 virtual account params
func (o *PostV1VirtualAccountParams) WithXTIMESTAMP(xTIMESTAMP string) *PostV1VirtualAccountParams {
	o.SetXTIMESTAMP(xTIMESTAMP)
	return o
}

// SetXTIMESTAMP adds the xTIMESTAMP to the post v1 virtual account params
func (o *PostV1VirtualAccountParams) SetXTIMESTAMP(xTIMESTAMP string) {
	o.XTIMESTAMP = xTIMESTAMP
}

// WithBody adds the body to the post v1 virtual account params
func (o *PostV1VirtualAccountParams) WithBody(body *models.CreateVirtualAccountRequest) *PostV1VirtualAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post v1 virtual account params
func (o *PostV1VirtualAccountParams) SetBody(body *models.CreateVirtualAccountRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1VirtualAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
