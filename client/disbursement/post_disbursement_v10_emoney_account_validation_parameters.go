// Code generated by go-swagger; DO NOT EDIT.

package disbursement

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

// NewPostDisbursementV10EmoneyAccountValidationParams creates a new PostDisbursementV10EmoneyAccountValidationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDisbursementV10EmoneyAccountValidationParams() *PostDisbursementV10EmoneyAccountValidationParams {
	return &PostDisbursementV10EmoneyAccountValidationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDisbursementV10EmoneyAccountValidationParamsWithTimeout creates a new PostDisbursementV10EmoneyAccountValidationParams object
// with the ability to set a timeout on a request.
func NewPostDisbursementV10EmoneyAccountValidationParamsWithTimeout(timeout time.Duration) *PostDisbursementV10EmoneyAccountValidationParams {
	return &PostDisbursementV10EmoneyAccountValidationParams{
		timeout: timeout,
	}
}

// NewPostDisbursementV10EmoneyAccountValidationParamsWithContext creates a new PostDisbursementV10EmoneyAccountValidationParams object
// with the ability to set a context for a request.
func NewPostDisbursementV10EmoneyAccountValidationParamsWithContext(ctx context.Context) *PostDisbursementV10EmoneyAccountValidationParams {
	return &PostDisbursementV10EmoneyAccountValidationParams{
		Context: ctx,
	}
}

// NewPostDisbursementV10EmoneyAccountValidationParamsWithHTTPClient creates a new PostDisbursementV10EmoneyAccountValidationParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDisbursementV10EmoneyAccountValidationParamsWithHTTPClient(client *http.Client) *PostDisbursementV10EmoneyAccountValidationParams {
	return &PostDisbursementV10EmoneyAccountValidationParams{
		HTTPClient: client,
	}
}

/*
PostDisbursementV10EmoneyAccountValidationParams contains all the parameters to send to the API endpoint

	for the post disbursement v10 emoney account validation operation.

	Typically these are written to a http.Request.
*/
type PostDisbursementV10EmoneyAccountValidationParams struct {

	/* CHANNELID.

	   Channel ID
	*/
	CHANNELID string

	/* Request.

	   E-money account validation request
	*/
	Request *models.EmoneyAccountValidationRequest

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

// WithDefaults hydrates default values in the post disbursement v10 emoney account validation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDisbursementV10EmoneyAccountValidationParams) WithDefaults() *PostDisbursementV10EmoneyAccountValidationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post disbursement v10 emoney account validation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDisbursementV10EmoneyAccountValidationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post disbursement v10 emoney account validation params
func (o *PostDisbursementV10EmoneyAccountValidationParams) WithTimeout(timeout time.Duration) *PostDisbursementV10EmoneyAccountValidationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post disbursement v10 emoney account validation params
func (o *PostDisbursementV10EmoneyAccountValidationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post disbursement v10 emoney account validation params
func (o *PostDisbursementV10EmoneyAccountValidationParams) WithContext(ctx context.Context) *PostDisbursementV10EmoneyAccountValidationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post disbursement v10 emoney account validation params
func (o *PostDisbursementV10EmoneyAccountValidationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post disbursement v10 emoney account validation params
func (o *PostDisbursementV10EmoneyAccountValidationParams) WithHTTPClient(client *http.Client) *PostDisbursementV10EmoneyAccountValidationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post disbursement v10 emoney account validation params
func (o *PostDisbursementV10EmoneyAccountValidationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCHANNELID adds the cHANNELID to the post disbursement v10 emoney account validation params
func (o *PostDisbursementV10EmoneyAccountValidationParams) WithCHANNELID(cHANNELID string) *PostDisbursementV10EmoneyAccountValidationParams {
	o.SetCHANNELID(cHANNELID)
	return o
}

// SetCHANNELID adds the cHANNELId to the post disbursement v10 emoney account validation params
func (o *PostDisbursementV10EmoneyAccountValidationParams) SetCHANNELID(cHANNELID string) {
	o.CHANNELID = cHANNELID
}

// WithRequest adds the request to the post disbursement v10 emoney account validation params
func (o *PostDisbursementV10EmoneyAccountValidationParams) WithRequest(request *models.EmoneyAccountValidationRequest) *PostDisbursementV10EmoneyAccountValidationParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post disbursement v10 emoney account validation params
func (o *PostDisbursementV10EmoneyAccountValidationParams) SetRequest(request *models.EmoneyAccountValidationRequest) {
	o.Request = request
}

// WithXEXTERNALID adds the xEXTERNALID to the post disbursement v10 emoney account validation params
func (o *PostDisbursementV10EmoneyAccountValidationParams) WithXEXTERNALID(xEXTERNALID string) *PostDisbursementV10EmoneyAccountValidationParams {
	o.SetXEXTERNALID(xEXTERNALID)
	return o
}

// SetXEXTERNALID adds the xEXTERNALId to the post disbursement v10 emoney account validation params
func (o *PostDisbursementV10EmoneyAccountValidationParams) SetXEXTERNALID(xEXTERNALID string) {
	o.XEXTERNALID = xEXTERNALID
}

// WithXSIGNATURE adds the xSIGNATURE to the post disbursement v10 emoney account validation params
func (o *PostDisbursementV10EmoneyAccountValidationParams) WithXSIGNATURE(xSIGNATURE string) *PostDisbursementV10EmoneyAccountValidationParams {
	o.SetXSIGNATURE(xSIGNATURE)
	return o
}

// SetXSIGNATURE adds the xSIGNATURE to the post disbursement v10 emoney account validation params
func (o *PostDisbursementV10EmoneyAccountValidationParams) SetXSIGNATURE(xSIGNATURE string) {
	o.XSIGNATURE = xSIGNATURE
}

// WithXTIMESTAMP adds the xTIMESTAMP to the post disbursement v10 emoney account validation params
func (o *PostDisbursementV10EmoneyAccountValidationParams) WithXTIMESTAMP(xTIMESTAMP string) *PostDisbursementV10EmoneyAccountValidationParams {
	o.SetXTIMESTAMP(xTIMESTAMP)
	return o
}

// SetXTIMESTAMP adds the xTIMESTAMP to the post disbursement v10 emoney account validation params
func (o *PostDisbursementV10EmoneyAccountValidationParams) SetXTIMESTAMP(xTIMESTAMP string) {
	o.XTIMESTAMP = xTIMESTAMP
}

// WriteToRequest writes these params to a swagger request
func (o *PostDisbursementV10EmoneyAccountValidationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
