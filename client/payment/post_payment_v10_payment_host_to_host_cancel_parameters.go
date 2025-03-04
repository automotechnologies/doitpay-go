// Code generated by go-swagger; DO NOT EDIT.

package payment

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

// NewPostPaymentV10PaymentHostToHostCancelParams creates a new PostPaymentV10PaymentHostToHostCancelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostPaymentV10PaymentHostToHostCancelParams() *PostPaymentV10PaymentHostToHostCancelParams {
	return &PostPaymentV10PaymentHostToHostCancelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostPaymentV10PaymentHostToHostCancelParamsWithTimeout creates a new PostPaymentV10PaymentHostToHostCancelParams object
// with the ability to set a timeout on a request.
func NewPostPaymentV10PaymentHostToHostCancelParamsWithTimeout(timeout time.Duration) *PostPaymentV10PaymentHostToHostCancelParams {
	return &PostPaymentV10PaymentHostToHostCancelParams{
		timeout: timeout,
	}
}

// NewPostPaymentV10PaymentHostToHostCancelParamsWithContext creates a new PostPaymentV10PaymentHostToHostCancelParams object
// with the ability to set a context for a request.
func NewPostPaymentV10PaymentHostToHostCancelParamsWithContext(ctx context.Context) *PostPaymentV10PaymentHostToHostCancelParams {
	return &PostPaymentV10PaymentHostToHostCancelParams{
		Context: ctx,
	}
}

// NewPostPaymentV10PaymentHostToHostCancelParamsWithHTTPClient creates a new PostPaymentV10PaymentHostToHostCancelParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostPaymentV10PaymentHostToHostCancelParamsWithHTTPClient(client *http.Client) *PostPaymentV10PaymentHostToHostCancelParams {
	return &PostPaymentV10PaymentHostToHostCancelParams{
		HTTPClient: client,
	}
}

/*
PostPaymentV10PaymentHostToHostCancelParams contains all the parameters to send to the API endpoint

	for the post payment v10 payment host to host cancel operation.

	Typically these are written to a http.Request.
*/
type PostPaymentV10PaymentHostToHostCancelParams struct {

	/* Request.

	   Cancel payment request
	*/
	Request *models.CancelHostToHostPaymentRequest

	/* XCLIENTKEY.

	   Client Key
	*/
	XCLIENTKEY string

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

// WithDefaults hydrates default values in the post payment v10 payment host to host cancel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostPaymentV10PaymentHostToHostCancelParams) WithDefaults() *PostPaymentV10PaymentHostToHostCancelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post payment v10 payment host to host cancel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostPaymentV10PaymentHostToHostCancelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post payment v10 payment host to host cancel params
func (o *PostPaymentV10PaymentHostToHostCancelParams) WithTimeout(timeout time.Duration) *PostPaymentV10PaymentHostToHostCancelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post payment v10 payment host to host cancel params
func (o *PostPaymentV10PaymentHostToHostCancelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post payment v10 payment host to host cancel params
func (o *PostPaymentV10PaymentHostToHostCancelParams) WithContext(ctx context.Context) *PostPaymentV10PaymentHostToHostCancelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post payment v10 payment host to host cancel params
func (o *PostPaymentV10PaymentHostToHostCancelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post payment v10 payment host to host cancel params
func (o *PostPaymentV10PaymentHostToHostCancelParams) WithHTTPClient(client *http.Client) *PostPaymentV10PaymentHostToHostCancelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post payment v10 payment host to host cancel params
func (o *PostPaymentV10PaymentHostToHostCancelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post payment v10 payment host to host cancel params
func (o *PostPaymentV10PaymentHostToHostCancelParams) WithRequest(request *models.CancelHostToHostPaymentRequest) *PostPaymentV10PaymentHostToHostCancelParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post payment v10 payment host to host cancel params
func (o *PostPaymentV10PaymentHostToHostCancelParams) SetRequest(request *models.CancelHostToHostPaymentRequest) {
	o.Request = request
}

// WithXCLIENTKEY adds the xCLIENTKEY to the post payment v10 payment host to host cancel params
func (o *PostPaymentV10PaymentHostToHostCancelParams) WithXCLIENTKEY(xCLIENTKEY string) *PostPaymentV10PaymentHostToHostCancelParams {
	o.SetXCLIENTKEY(xCLIENTKEY)
	return o
}

// SetXCLIENTKEY adds the xCLIENTKEY to the post payment v10 payment host to host cancel params
func (o *PostPaymentV10PaymentHostToHostCancelParams) SetXCLIENTKEY(xCLIENTKEY string) {
	o.XCLIENTKEY = xCLIENTKEY
}

// WithXSIGNATURE adds the xSIGNATURE to the post payment v10 payment host to host cancel params
func (o *PostPaymentV10PaymentHostToHostCancelParams) WithXSIGNATURE(xSIGNATURE string) *PostPaymentV10PaymentHostToHostCancelParams {
	o.SetXSIGNATURE(xSIGNATURE)
	return o
}

// SetXSIGNATURE adds the xSIGNATURE to the post payment v10 payment host to host cancel params
func (o *PostPaymentV10PaymentHostToHostCancelParams) SetXSIGNATURE(xSIGNATURE string) {
	o.XSIGNATURE = xSIGNATURE
}

// WithXTIMESTAMP adds the xTIMESTAMP to the post payment v10 payment host to host cancel params
func (o *PostPaymentV10PaymentHostToHostCancelParams) WithXTIMESTAMP(xTIMESTAMP string) *PostPaymentV10PaymentHostToHostCancelParams {
	o.SetXTIMESTAMP(xTIMESTAMP)
	return o
}

// SetXTIMESTAMP adds the xTIMESTAMP to the post payment v10 payment host to host cancel params
func (o *PostPaymentV10PaymentHostToHostCancelParams) SetXTIMESTAMP(xTIMESTAMP string) {
	o.XTIMESTAMP = xTIMESTAMP
}

// WriteToRequest writes these params to a swagger request
func (o *PostPaymentV10PaymentHostToHostCancelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	// header param X-CLIENT-KEY
	if err := r.SetHeaderParam("X-CLIENT-KEY", o.XCLIENTKEY); err != nil {
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
