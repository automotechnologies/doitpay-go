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

// NewPostVaV10TransferVaCreateVaParams creates a new PostVaV10TransferVaCreateVaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostVaV10TransferVaCreateVaParams() *PostVaV10TransferVaCreateVaParams {
	return &PostVaV10TransferVaCreateVaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostVaV10TransferVaCreateVaParamsWithTimeout creates a new PostVaV10TransferVaCreateVaParams object
// with the ability to set a timeout on a request.
func NewPostVaV10TransferVaCreateVaParamsWithTimeout(timeout time.Duration) *PostVaV10TransferVaCreateVaParams {
	return &PostVaV10TransferVaCreateVaParams{
		timeout: timeout,
	}
}

// NewPostVaV10TransferVaCreateVaParamsWithContext creates a new PostVaV10TransferVaCreateVaParams object
// with the ability to set a context for a request.
func NewPostVaV10TransferVaCreateVaParamsWithContext(ctx context.Context) *PostVaV10TransferVaCreateVaParams {
	return &PostVaV10TransferVaCreateVaParams{
		Context: ctx,
	}
}

// NewPostVaV10TransferVaCreateVaParamsWithHTTPClient creates a new PostVaV10TransferVaCreateVaParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostVaV10TransferVaCreateVaParamsWithHTTPClient(client *http.Client) *PostVaV10TransferVaCreateVaParams {
	return &PostVaV10TransferVaCreateVaParams{
		HTTPClient: client,
	}
}

/*
PostVaV10TransferVaCreateVaParams contains all the parameters to send to the API endpoint

	for the post va v10 transfer va create va operation.

	Typically these are written to a http.Request.
*/
type PostVaV10TransferVaCreateVaParams struct {

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

// WithDefaults hydrates default values in the post va v10 transfer va create va params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostVaV10TransferVaCreateVaParams) WithDefaults() *PostVaV10TransferVaCreateVaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post va v10 transfer va create va params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostVaV10TransferVaCreateVaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post va v10 transfer va create va params
func (o *PostVaV10TransferVaCreateVaParams) WithTimeout(timeout time.Duration) *PostVaV10TransferVaCreateVaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post va v10 transfer va create va params
func (o *PostVaV10TransferVaCreateVaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post va v10 transfer va create va params
func (o *PostVaV10TransferVaCreateVaParams) WithContext(ctx context.Context) *PostVaV10TransferVaCreateVaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post va v10 transfer va create va params
func (o *PostVaV10TransferVaCreateVaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post va v10 transfer va create va params
func (o *PostVaV10TransferVaCreateVaParams) WithHTTPClient(client *http.Client) *PostVaV10TransferVaCreateVaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post va v10 transfer va create va params
func (o *PostVaV10TransferVaCreateVaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCHANNELID adds the cHANNELID to the post va v10 transfer va create va params
func (o *PostVaV10TransferVaCreateVaParams) WithCHANNELID(cHANNELID string) *PostVaV10TransferVaCreateVaParams {
	o.SetCHANNELID(cHANNELID)
	return o
}

// SetCHANNELID adds the cHANNELId to the post va v10 transfer va create va params
func (o *PostVaV10TransferVaCreateVaParams) SetCHANNELID(cHANNELID string) {
	o.CHANNELID = cHANNELID
}

// WithXEXTERNALID adds the xEXTERNALID to the post va v10 transfer va create va params
func (o *PostVaV10TransferVaCreateVaParams) WithXEXTERNALID(xEXTERNALID string) *PostVaV10TransferVaCreateVaParams {
	o.SetXEXTERNALID(xEXTERNALID)
	return o
}

// SetXEXTERNALID adds the xEXTERNALId to the post va v10 transfer va create va params
func (o *PostVaV10TransferVaCreateVaParams) SetXEXTERNALID(xEXTERNALID string) {
	o.XEXTERNALID = xEXTERNALID
}

// WithXSIGNATURE adds the xSIGNATURE to the post va v10 transfer va create va params
func (o *PostVaV10TransferVaCreateVaParams) WithXSIGNATURE(xSIGNATURE string) *PostVaV10TransferVaCreateVaParams {
	o.SetXSIGNATURE(xSIGNATURE)
	return o
}

// SetXSIGNATURE adds the xSIGNATURE to the post va v10 transfer va create va params
func (o *PostVaV10TransferVaCreateVaParams) SetXSIGNATURE(xSIGNATURE string) {
	o.XSIGNATURE = xSIGNATURE
}

// WithXTIMESTAMP adds the xTIMESTAMP to the post va v10 transfer va create va params
func (o *PostVaV10TransferVaCreateVaParams) WithXTIMESTAMP(xTIMESTAMP string) *PostVaV10TransferVaCreateVaParams {
	o.SetXTIMESTAMP(xTIMESTAMP)
	return o
}

// SetXTIMESTAMP adds the xTIMESTAMP to the post va v10 transfer va create va params
func (o *PostVaV10TransferVaCreateVaParams) SetXTIMESTAMP(xTIMESTAMP string) {
	o.XTIMESTAMP = xTIMESTAMP
}

// WithBody adds the body to the post va v10 transfer va create va params
func (o *PostVaV10TransferVaCreateVaParams) WithBody(body *models.CreateVirtualAccountRequest) *PostVaV10TransferVaCreateVaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post va v10 transfer va create va params
func (o *PostVaV10TransferVaCreateVaParams) SetBody(body *models.CreateVirtualAccountRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostVaV10TransferVaCreateVaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
