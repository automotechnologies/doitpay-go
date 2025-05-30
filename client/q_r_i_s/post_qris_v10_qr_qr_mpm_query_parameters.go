// Code generated by go-swagger; DO NOT EDIT.

package q_r_i_s

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

// NewPostQrisV10QrQrMpmQueryParams creates a new PostQrisV10QrQrMpmQueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostQrisV10QrQrMpmQueryParams() *PostQrisV10QrQrMpmQueryParams {
	return &PostQrisV10QrQrMpmQueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostQrisV10QrQrMpmQueryParamsWithTimeout creates a new PostQrisV10QrQrMpmQueryParams object
// with the ability to set a timeout on a request.
func NewPostQrisV10QrQrMpmQueryParamsWithTimeout(timeout time.Duration) *PostQrisV10QrQrMpmQueryParams {
	return &PostQrisV10QrQrMpmQueryParams{
		timeout: timeout,
	}
}

// NewPostQrisV10QrQrMpmQueryParamsWithContext creates a new PostQrisV10QrQrMpmQueryParams object
// with the ability to set a context for a request.
func NewPostQrisV10QrQrMpmQueryParamsWithContext(ctx context.Context) *PostQrisV10QrQrMpmQueryParams {
	return &PostQrisV10QrQrMpmQueryParams{
		Context: ctx,
	}
}

// NewPostQrisV10QrQrMpmQueryParamsWithHTTPClient creates a new PostQrisV10QrQrMpmQueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostQrisV10QrQrMpmQueryParamsWithHTTPClient(client *http.Client) *PostQrisV10QrQrMpmQueryParams {
	return &PostQrisV10QrQrMpmQueryParams{
		HTTPClient: client,
	}
}

/*
PostQrisV10QrQrMpmQueryParams contains all the parameters to send to the API endpoint

	for the post qris v10 qr qr mpm query operation.

	Typically these are written to a http.Request.
*/
type PostQrisV10QrQrMpmQueryParams struct {

	/* CHANNELID.

	   Channel ID
	*/
	CHANNELID string

	/* Request.

	   Query request
	*/
	Request *models.QrisQueryPaymentRequest

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post qris v10 qr qr mpm query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostQrisV10QrQrMpmQueryParams) WithDefaults() *PostQrisV10QrQrMpmQueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post qris v10 qr qr mpm query params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostQrisV10QrQrMpmQueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post qris v10 qr qr mpm query params
func (o *PostQrisV10QrQrMpmQueryParams) WithTimeout(timeout time.Duration) *PostQrisV10QrQrMpmQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post qris v10 qr qr mpm query params
func (o *PostQrisV10QrQrMpmQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post qris v10 qr qr mpm query params
func (o *PostQrisV10QrQrMpmQueryParams) WithContext(ctx context.Context) *PostQrisV10QrQrMpmQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post qris v10 qr qr mpm query params
func (o *PostQrisV10QrQrMpmQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post qris v10 qr qr mpm query params
func (o *PostQrisV10QrQrMpmQueryParams) WithHTTPClient(client *http.Client) *PostQrisV10QrQrMpmQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post qris v10 qr qr mpm query params
func (o *PostQrisV10QrQrMpmQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCHANNELID adds the cHANNELID to the post qris v10 qr qr mpm query params
func (o *PostQrisV10QrQrMpmQueryParams) WithCHANNELID(cHANNELID string) *PostQrisV10QrQrMpmQueryParams {
	o.SetCHANNELID(cHANNELID)
	return o
}

// SetCHANNELID adds the cHANNELId to the post qris v10 qr qr mpm query params
func (o *PostQrisV10QrQrMpmQueryParams) SetCHANNELID(cHANNELID string) {
	o.CHANNELID = cHANNELID
}

// WithRequest adds the request to the post qris v10 qr qr mpm query params
func (o *PostQrisV10QrQrMpmQueryParams) WithRequest(request *models.QrisQueryPaymentRequest) *PostQrisV10QrQrMpmQueryParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post qris v10 qr qr mpm query params
func (o *PostQrisV10QrQrMpmQueryParams) SetRequest(request *models.QrisQueryPaymentRequest) {
	o.Request = request
}

// WithXEXTERNALID adds the xEXTERNALID to the post qris v10 qr qr mpm query params
func (o *PostQrisV10QrQrMpmQueryParams) WithXEXTERNALID(xEXTERNALID string) *PostQrisV10QrQrMpmQueryParams {
	o.SetXEXTERNALID(xEXTERNALID)
	return o
}

// SetXEXTERNALID adds the xEXTERNALId to the post qris v10 qr qr mpm query params
func (o *PostQrisV10QrQrMpmQueryParams) SetXEXTERNALID(xEXTERNALID string) {
	o.XEXTERNALID = xEXTERNALID
}

// WithXSIGNATURE adds the xSIGNATURE to the post qris v10 qr qr mpm query params
func (o *PostQrisV10QrQrMpmQueryParams) WithXSIGNATURE(xSIGNATURE string) *PostQrisV10QrQrMpmQueryParams {
	o.SetXSIGNATURE(xSIGNATURE)
	return o
}

// SetXSIGNATURE adds the xSIGNATURE to the post qris v10 qr qr mpm query params
func (o *PostQrisV10QrQrMpmQueryParams) SetXSIGNATURE(xSIGNATURE string) {
	o.XSIGNATURE = xSIGNATURE
}

// WithXTIMESTAMP adds the xTIMESTAMP to the post qris v10 qr qr mpm query params
func (o *PostQrisV10QrQrMpmQueryParams) WithXTIMESTAMP(xTIMESTAMP string) *PostQrisV10QrQrMpmQueryParams {
	o.SetXTIMESTAMP(xTIMESTAMP)
	return o
}

// SetXTIMESTAMP adds the xTIMESTAMP to the post qris v10 qr qr mpm query params
func (o *PostQrisV10QrQrMpmQueryParams) SetXTIMESTAMP(xTIMESTAMP string) {
	o.XTIMESTAMP = xTIMESTAMP
}

// WriteToRequest writes these params to a swagger request
func (o *PostQrisV10QrQrMpmQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param CHANNEL-ID
	if err := r.SetHeaderParam("CHANNEL-ID", o.CHANNELID); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
