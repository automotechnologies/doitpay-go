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

// NewPostQrisV10QrQrMpmGenerateParams creates a new PostQrisV10QrQrMpmGenerateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostQrisV10QrQrMpmGenerateParams() *PostQrisV10QrQrMpmGenerateParams {
	return &PostQrisV10QrQrMpmGenerateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostQrisV10QrQrMpmGenerateParamsWithTimeout creates a new PostQrisV10QrQrMpmGenerateParams object
// with the ability to set a timeout on a request.
func NewPostQrisV10QrQrMpmGenerateParamsWithTimeout(timeout time.Duration) *PostQrisV10QrQrMpmGenerateParams {
	return &PostQrisV10QrQrMpmGenerateParams{
		timeout: timeout,
	}
}

// NewPostQrisV10QrQrMpmGenerateParamsWithContext creates a new PostQrisV10QrQrMpmGenerateParams object
// with the ability to set a context for a request.
func NewPostQrisV10QrQrMpmGenerateParamsWithContext(ctx context.Context) *PostQrisV10QrQrMpmGenerateParams {
	return &PostQrisV10QrQrMpmGenerateParams{
		Context: ctx,
	}
}

// NewPostQrisV10QrQrMpmGenerateParamsWithHTTPClient creates a new PostQrisV10QrQrMpmGenerateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostQrisV10QrQrMpmGenerateParamsWithHTTPClient(client *http.Client) *PostQrisV10QrQrMpmGenerateParams {
	return &PostQrisV10QrQrMpmGenerateParams{
		HTTPClient: client,
	}
}

/*
PostQrisV10QrQrMpmGenerateParams contains all the parameters to send to the API endpoint

	for the post qris v10 qr qr mpm generate operation.

	Typically these are written to a http.Request.
*/
type PostQrisV10QrQrMpmGenerateParams struct {

	/* CHANNELID.

	   Channel ID
	*/
	CHANNELID string

	/* Request.

	   QRIS generation request
	*/
	Request *models.QrisRequestScheme

	/* XEXTERNALID.

	   External ID
	*/
	XEXTERNALID string

	/* XPARTNERID.

	   Partner ID
	*/
	XPARTNERID string

	/* XREQUESTID.

	   Request ID
	*/
	XREQUESTID string

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

// WithDefaults hydrates default values in the post qris v10 qr qr mpm generate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostQrisV10QrQrMpmGenerateParams) WithDefaults() *PostQrisV10QrQrMpmGenerateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post qris v10 qr qr mpm generate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostQrisV10QrQrMpmGenerateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) WithTimeout(timeout time.Duration) *PostQrisV10QrQrMpmGenerateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) WithContext(ctx context.Context) *PostQrisV10QrQrMpmGenerateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) WithHTTPClient(client *http.Client) *PostQrisV10QrQrMpmGenerateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCHANNELID adds the cHANNELID to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) WithCHANNELID(cHANNELID string) *PostQrisV10QrQrMpmGenerateParams {
	o.SetCHANNELID(cHANNELID)
	return o
}

// SetCHANNELID adds the cHANNELId to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) SetCHANNELID(cHANNELID string) {
	o.CHANNELID = cHANNELID
}

// WithRequest adds the request to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) WithRequest(request *models.QrisRequestScheme) *PostQrisV10QrQrMpmGenerateParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) SetRequest(request *models.QrisRequestScheme) {
	o.Request = request
}

// WithXEXTERNALID adds the xEXTERNALID to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) WithXEXTERNALID(xEXTERNALID string) *PostQrisV10QrQrMpmGenerateParams {
	o.SetXEXTERNALID(xEXTERNALID)
	return o
}

// SetXEXTERNALID adds the xEXTERNALId to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) SetXEXTERNALID(xEXTERNALID string) {
	o.XEXTERNALID = xEXTERNALID
}

// WithXPARTNERID adds the xPARTNERID to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) WithXPARTNERID(xPARTNERID string) *PostQrisV10QrQrMpmGenerateParams {
	o.SetXPARTNERID(xPARTNERID)
	return o
}

// SetXPARTNERID adds the xPARTNERId to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) SetXPARTNERID(xPARTNERID string) {
	o.XPARTNERID = xPARTNERID
}

// WithXREQUESTID adds the xREQUESTID to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) WithXREQUESTID(xREQUESTID string) *PostQrisV10QrQrMpmGenerateParams {
	o.SetXREQUESTID(xREQUESTID)
	return o
}

// SetXREQUESTID adds the xREQUESTId to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) SetXREQUESTID(xREQUESTID string) {
	o.XREQUESTID = xREQUESTID
}

// WithXSIGNATURE adds the xSIGNATURE to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) WithXSIGNATURE(xSIGNATURE string) *PostQrisV10QrQrMpmGenerateParams {
	o.SetXSIGNATURE(xSIGNATURE)
	return o
}

// SetXSIGNATURE adds the xSIGNATURE to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) SetXSIGNATURE(xSIGNATURE string) {
	o.XSIGNATURE = xSIGNATURE
}

// WithXTIMESTAMP adds the xTIMESTAMP to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) WithXTIMESTAMP(xTIMESTAMP string) *PostQrisV10QrQrMpmGenerateParams {
	o.SetXTIMESTAMP(xTIMESTAMP)
	return o
}

// SetXTIMESTAMP adds the xTIMESTAMP to the post qris v10 qr qr mpm generate params
func (o *PostQrisV10QrQrMpmGenerateParams) SetXTIMESTAMP(xTIMESTAMP string) {
	o.XTIMESTAMP = xTIMESTAMP
}

// WriteToRequest writes these params to a swagger request
func (o *PostQrisV10QrQrMpmGenerateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// header param X-PARTNER-ID
	if err := r.SetHeaderParam("X-PARTNER-ID", o.XPARTNERID); err != nil {
		return err
	}

	// header param X-REQUEST-ID
	if err := r.SetHeaderParam("X-REQUEST-ID", o.XREQUESTID); err != nil {
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
