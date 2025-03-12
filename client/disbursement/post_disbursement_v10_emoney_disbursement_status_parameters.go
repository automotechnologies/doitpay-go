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

// NewPostDisbursementV10EmoneyDisbursementStatusParams creates a new PostDisbursementV10EmoneyDisbursementStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDisbursementV10EmoneyDisbursementStatusParams() *PostDisbursementV10EmoneyDisbursementStatusParams {
	return &PostDisbursementV10EmoneyDisbursementStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDisbursementV10EmoneyDisbursementStatusParamsWithTimeout creates a new PostDisbursementV10EmoneyDisbursementStatusParams object
// with the ability to set a timeout on a request.
func NewPostDisbursementV10EmoneyDisbursementStatusParamsWithTimeout(timeout time.Duration) *PostDisbursementV10EmoneyDisbursementStatusParams {
	return &PostDisbursementV10EmoneyDisbursementStatusParams{
		timeout: timeout,
	}
}

// NewPostDisbursementV10EmoneyDisbursementStatusParamsWithContext creates a new PostDisbursementV10EmoneyDisbursementStatusParams object
// with the ability to set a context for a request.
func NewPostDisbursementV10EmoneyDisbursementStatusParamsWithContext(ctx context.Context) *PostDisbursementV10EmoneyDisbursementStatusParams {
	return &PostDisbursementV10EmoneyDisbursementStatusParams{
		Context: ctx,
	}
}

// NewPostDisbursementV10EmoneyDisbursementStatusParamsWithHTTPClient creates a new PostDisbursementV10EmoneyDisbursementStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDisbursementV10EmoneyDisbursementStatusParamsWithHTTPClient(client *http.Client) *PostDisbursementV10EmoneyDisbursementStatusParams {
	return &PostDisbursementV10EmoneyDisbursementStatusParams{
		HTTPClient: client,
	}
}

/*
PostDisbursementV10EmoneyDisbursementStatusParams contains all the parameters to send to the API endpoint

	for the post disbursement v10 emoney disbursement status operation.

	Typically these are written to a http.Request.
*/
type PostDisbursementV10EmoneyDisbursementStatusParams struct {

	/* CHANNELID.

	   Channel ID
	*/
	CHANNELID string

	/* Request.

	   Status inquiry request
	*/
	Request *models.InquiryEwalletTopupStatusRequest

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

// WithDefaults hydrates default values in the post disbursement v10 emoney disbursement status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) WithDefaults() *PostDisbursementV10EmoneyDisbursementStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post disbursement v10 emoney disbursement status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post disbursement v10 emoney disbursement status params
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) WithTimeout(timeout time.Duration) *PostDisbursementV10EmoneyDisbursementStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post disbursement v10 emoney disbursement status params
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post disbursement v10 emoney disbursement status params
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) WithContext(ctx context.Context) *PostDisbursementV10EmoneyDisbursementStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post disbursement v10 emoney disbursement status params
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post disbursement v10 emoney disbursement status params
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) WithHTTPClient(client *http.Client) *PostDisbursementV10EmoneyDisbursementStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post disbursement v10 emoney disbursement status params
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCHANNELID adds the cHANNELID to the post disbursement v10 emoney disbursement status params
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) WithCHANNELID(cHANNELID string) *PostDisbursementV10EmoneyDisbursementStatusParams {
	o.SetCHANNELID(cHANNELID)
	return o
}

// SetCHANNELID adds the cHANNELId to the post disbursement v10 emoney disbursement status params
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) SetCHANNELID(cHANNELID string) {
	o.CHANNELID = cHANNELID
}

// WithRequest adds the request to the post disbursement v10 emoney disbursement status params
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) WithRequest(request *models.InquiryEwalletTopupStatusRequest) *PostDisbursementV10EmoneyDisbursementStatusParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post disbursement v10 emoney disbursement status params
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) SetRequest(request *models.InquiryEwalletTopupStatusRequest) {
	o.Request = request
}

// WithXEXTERNALID adds the xEXTERNALID to the post disbursement v10 emoney disbursement status params
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) WithXEXTERNALID(xEXTERNALID string) *PostDisbursementV10EmoneyDisbursementStatusParams {
	o.SetXEXTERNALID(xEXTERNALID)
	return o
}

// SetXEXTERNALID adds the xEXTERNALId to the post disbursement v10 emoney disbursement status params
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) SetXEXTERNALID(xEXTERNALID string) {
	o.XEXTERNALID = xEXTERNALID
}

// WithXSIGNATURE adds the xSIGNATURE to the post disbursement v10 emoney disbursement status params
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) WithXSIGNATURE(xSIGNATURE string) *PostDisbursementV10EmoneyDisbursementStatusParams {
	o.SetXSIGNATURE(xSIGNATURE)
	return o
}

// SetXSIGNATURE adds the xSIGNATURE to the post disbursement v10 emoney disbursement status params
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) SetXSIGNATURE(xSIGNATURE string) {
	o.XSIGNATURE = xSIGNATURE
}

// WithXTIMESTAMP adds the xTIMESTAMP to the post disbursement v10 emoney disbursement status params
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) WithXTIMESTAMP(xTIMESTAMP string) *PostDisbursementV10EmoneyDisbursementStatusParams {
	o.SetXTIMESTAMP(xTIMESTAMP)
	return o
}

// SetXTIMESTAMP adds the xTIMESTAMP to the post disbursement v10 emoney disbursement status params
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) SetXTIMESTAMP(xTIMESTAMP string) {
	o.XTIMESTAMP = xTIMESTAMP
}

// WriteToRequest writes these params to a swagger request
func (o *PostDisbursementV10EmoneyDisbursementStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
