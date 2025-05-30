// Code generated by go-swagger; DO NOT EDIT.

package debit

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

// NewPostAPIV10DebitCancelParams creates a new PostAPIV10DebitCancelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAPIV10DebitCancelParams() *PostAPIV10DebitCancelParams {
	return &PostAPIV10DebitCancelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV10DebitCancelParamsWithTimeout creates a new PostAPIV10DebitCancelParams object
// with the ability to set a timeout on a request.
func NewPostAPIV10DebitCancelParamsWithTimeout(timeout time.Duration) *PostAPIV10DebitCancelParams {
	return &PostAPIV10DebitCancelParams{
		timeout: timeout,
	}
}

// NewPostAPIV10DebitCancelParamsWithContext creates a new PostAPIV10DebitCancelParams object
// with the ability to set a context for a request.
func NewPostAPIV10DebitCancelParamsWithContext(ctx context.Context) *PostAPIV10DebitCancelParams {
	return &PostAPIV10DebitCancelParams{
		Context: ctx,
	}
}

// NewPostAPIV10DebitCancelParamsWithHTTPClient creates a new PostAPIV10DebitCancelParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAPIV10DebitCancelParamsWithHTTPClient(client *http.Client) *PostAPIV10DebitCancelParams {
	return &PostAPIV10DebitCancelParams{
		HTTPClient: client,
	}
}

/*
PostAPIV10DebitCancelParams contains all the parameters to send to the API endpoint

	for the post API v10 debit cancel operation.

	Typically these are written to a http.Request.
*/
type PostAPIV10DebitCancelParams struct {

	/* CHANNELID.

	   Channel ID
	*/
	CHANNELID string

	/* Request.

	   Cancel payment request
	*/
	Request *models.CancelHostToHostPaymentRequest

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

// WithDefaults hydrates default values in the post API v10 debit cancel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV10DebitCancelParams) WithDefaults() *PostAPIV10DebitCancelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post API v10 debit cancel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAPIV10DebitCancelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post API v10 debit cancel params
func (o *PostAPIV10DebitCancelParams) WithTimeout(timeout time.Duration) *PostAPIV10DebitCancelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API v10 debit cancel params
func (o *PostAPIV10DebitCancelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API v10 debit cancel params
func (o *PostAPIV10DebitCancelParams) WithContext(ctx context.Context) *PostAPIV10DebitCancelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API v10 debit cancel params
func (o *PostAPIV10DebitCancelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API v10 debit cancel params
func (o *PostAPIV10DebitCancelParams) WithHTTPClient(client *http.Client) *PostAPIV10DebitCancelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API v10 debit cancel params
func (o *PostAPIV10DebitCancelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCHANNELID adds the cHANNELID to the post API v10 debit cancel params
func (o *PostAPIV10DebitCancelParams) WithCHANNELID(cHANNELID string) *PostAPIV10DebitCancelParams {
	o.SetCHANNELID(cHANNELID)
	return o
}

// SetCHANNELID adds the cHANNELId to the post API v10 debit cancel params
func (o *PostAPIV10DebitCancelParams) SetCHANNELID(cHANNELID string) {
	o.CHANNELID = cHANNELID
}

// WithRequest adds the request to the post API v10 debit cancel params
func (o *PostAPIV10DebitCancelParams) WithRequest(request *models.CancelHostToHostPaymentRequest) *PostAPIV10DebitCancelParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post API v10 debit cancel params
func (o *PostAPIV10DebitCancelParams) SetRequest(request *models.CancelHostToHostPaymentRequest) {
	o.Request = request
}

// WithXEXTERNALID adds the xEXTERNALID to the post API v10 debit cancel params
func (o *PostAPIV10DebitCancelParams) WithXEXTERNALID(xEXTERNALID string) *PostAPIV10DebitCancelParams {
	o.SetXEXTERNALID(xEXTERNALID)
	return o
}

// SetXEXTERNALID adds the xEXTERNALId to the post API v10 debit cancel params
func (o *PostAPIV10DebitCancelParams) SetXEXTERNALID(xEXTERNALID string) {
	o.XEXTERNALID = xEXTERNALID
}

// WithXSIGNATURE adds the xSIGNATURE to the post API v10 debit cancel params
func (o *PostAPIV10DebitCancelParams) WithXSIGNATURE(xSIGNATURE string) *PostAPIV10DebitCancelParams {
	o.SetXSIGNATURE(xSIGNATURE)
	return o
}

// SetXSIGNATURE adds the xSIGNATURE to the post API v10 debit cancel params
func (o *PostAPIV10DebitCancelParams) SetXSIGNATURE(xSIGNATURE string) {
	o.XSIGNATURE = xSIGNATURE
}

// WithXTIMESTAMP adds the xTIMESTAMP to the post API v10 debit cancel params
func (o *PostAPIV10DebitCancelParams) WithXTIMESTAMP(xTIMESTAMP string) *PostAPIV10DebitCancelParams {
	o.SetXTIMESTAMP(xTIMESTAMP)
	return o
}

// SetXTIMESTAMP adds the xTIMESTAMP to the post API v10 debit cancel params
func (o *PostAPIV10DebitCancelParams) SetXTIMESTAMP(xTIMESTAMP string) {
	o.XTIMESTAMP = xTIMESTAMP
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV10DebitCancelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
