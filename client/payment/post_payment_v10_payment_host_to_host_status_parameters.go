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

// NewPostPaymentV10PaymentHostToHostStatusParams creates a new PostPaymentV10PaymentHostToHostStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostPaymentV10PaymentHostToHostStatusParams() *PostPaymentV10PaymentHostToHostStatusParams {
	return &PostPaymentV10PaymentHostToHostStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostPaymentV10PaymentHostToHostStatusParamsWithTimeout creates a new PostPaymentV10PaymentHostToHostStatusParams object
// with the ability to set a timeout on a request.
func NewPostPaymentV10PaymentHostToHostStatusParamsWithTimeout(timeout time.Duration) *PostPaymentV10PaymentHostToHostStatusParams {
	return &PostPaymentV10PaymentHostToHostStatusParams{
		timeout: timeout,
	}
}

// NewPostPaymentV10PaymentHostToHostStatusParamsWithContext creates a new PostPaymentV10PaymentHostToHostStatusParams object
// with the ability to set a context for a request.
func NewPostPaymentV10PaymentHostToHostStatusParamsWithContext(ctx context.Context) *PostPaymentV10PaymentHostToHostStatusParams {
	return &PostPaymentV10PaymentHostToHostStatusParams{
		Context: ctx,
	}
}

// NewPostPaymentV10PaymentHostToHostStatusParamsWithHTTPClient creates a new PostPaymentV10PaymentHostToHostStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostPaymentV10PaymentHostToHostStatusParamsWithHTTPClient(client *http.Client) *PostPaymentV10PaymentHostToHostStatusParams {
	return &PostPaymentV10PaymentHostToHostStatusParams{
		HTTPClient: client,
	}
}

/*
PostPaymentV10PaymentHostToHostStatusParams contains all the parameters to send to the API endpoint

	for the post payment v10 payment host to host status operation.

	Typically these are written to a http.Request.
*/
type PostPaymentV10PaymentHostToHostStatusParams struct {

	/* Request.

	   Status check request
	*/
	Request *models.CheckPaymentStatusEwalletRequest

	/* XCHANNELID.

	   Channel ID
	*/
	XCHANNELID string

	/* XCLIENTKEY.

	   Client Key
	*/
	XCLIENTKEY string

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

// WithDefaults hydrates default values in the post payment v10 payment host to host status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostPaymentV10PaymentHostToHostStatusParams) WithDefaults() *PostPaymentV10PaymentHostToHostStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post payment v10 payment host to host status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostPaymentV10PaymentHostToHostStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) WithTimeout(timeout time.Duration) *PostPaymentV10PaymentHostToHostStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) WithContext(ctx context.Context) *PostPaymentV10PaymentHostToHostStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) WithHTTPClient(client *http.Client) *PostPaymentV10PaymentHostToHostStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) WithRequest(request *models.CheckPaymentStatusEwalletRequest) *PostPaymentV10PaymentHostToHostStatusParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) SetRequest(request *models.CheckPaymentStatusEwalletRequest) {
	o.Request = request
}

// WithXCHANNELID adds the xCHANNELID to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) WithXCHANNELID(xCHANNELID string) *PostPaymentV10PaymentHostToHostStatusParams {
	o.SetXCHANNELID(xCHANNELID)
	return o
}

// SetXCHANNELID adds the xCHANNELId to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) SetXCHANNELID(xCHANNELID string) {
	o.XCHANNELID = xCHANNELID
}

// WithXCLIENTKEY adds the xCLIENTKEY to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) WithXCLIENTKEY(xCLIENTKEY string) *PostPaymentV10PaymentHostToHostStatusParams {
	o.SetXCLIENTKEY(xCLIENTKEY)
	return o
}

// SetXCLIENTKEY adds the xCLIENTKEY to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) SetXCLIENTKEY(xCLIENTKEY string) {
	o.XCLIENTKEY = xCLIENTKEY
}

// WithXPARTNERID adds the xPARTNERID to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) WithXPARTNERID(xPARTNERID string) *PostPaymentV10PaymentHostToHostStatusParams {
	o.SetXPARTNERID(xPARTNERID)
	return o
}

// SetXPARTNERID adds the xPARTNERId to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) SetXPARTNERID(xPARTNERID string) {
	o.XPARTNERID = xPARTNERID
}

// WithXREQUESTID adds the xREQUESTID to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) WithXREQUESTID(xREQUESTID string) *PostPaymentV10PaymentHostToHostStatusParams {
	o.SetXREQUESTID(xREQUESTID)
	return o
}

// SetXREQUESTID adds the xREQUESTId to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) SetXREQUESTID(xREQUESTID string) {
	o.XREQUESTID = xREQUESTID
}

// WithXSIGNATURE adds the xSIGNATURE to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) WithXSIGNATURE(xSIGNATURE string) *PostPaymentV10PaymentHostToHostStatusParams {
	o.SetXSIGNATURE(xSIGNATURE)
	return o
}

// SetXSIGNATURE adds the xSIGNATURE to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) SetXSIGNATURE(xSIGNATURE string) {
	o.XSIGNATURE = xSIGNATURE
}

// WithXTIMESTAMP adds the xTIMESTAMP to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) WithXTIMESTAMP(xTIMESTAMP string) *PostPaymentV10PaymentHostToHostStatusParams {
	o.SetXTIMESTAMP(xTIMESTAMP)
	return o
}

// SetXTIMESTAMP adds the xTIMESTAMP to the post payment v10 payment host to host status params
func (o *PostPaymentV10PaymentHostToHostStatusParams) SetXTIMESTAMP(xTIMESTAMP string) {
	o.XTIMESTAMP = xTIMESTAMP
}

// WriteToRequest writes these params to a swagger request
func (o *PostPaymentV10PaymentHostToHostStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	// header param X-CHANNEL-ID
	if err := r.SetHeaderParam("X-CHANNEL-ID", o.XCHANNELID); err != nil {
		return err
	}

	// header param X-CLIENT-KEY
	if err := r.SetHeaderParam("X-CLIENT-KEY", o.XCLIENTKEY); err != nil {
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
