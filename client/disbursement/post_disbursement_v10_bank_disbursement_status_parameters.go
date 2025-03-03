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

	"github.com/automotechnologies/doitpay-go/models"
)

// NewPostDisbursementV10BankDisbursementStatusParams creates a new PostDisbursementV10BankDisbursementStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDisbursementV10BankDisbursementStatusParams() *PostDisbursementV10BankDisbursementStatusParams {
	return &PostDisbursementV10BankDisbursementStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDisbursementV10BankDisbursementStatusParamsWithTimeout creates a new PostDisbursementV10BankDisbursementStatusParams object
// with the ability to set a timeout on a request.
func NewPostDisbursementV10BankDisbursementStatusParamsWithTimeout(timeout time.Duration) *PostDisbursementV10BankDisbursementStatusParams {
	return &PostDisbursementV10BankDisbursementStatusParams{
		timeout: timeout,
	}
}

// NewPostDisbursementV10BankDisbursementStatusParamsWithContext creates a new PostDisbursementV10BankDisbursementStatusParams object
// with the ability to set a context for a request.
func NewPostDisbursementV10BankDisbursementStatusParamsWithContext(ctx context.Context) *PostDisbursementV10BankDisbursementStatusParams {
	return &PostDisbursementV10BankDisbursementStatusParams{
		Context: ctx,
	}
}

// NewPostDisbursementV10BankDisbursementStatusParamsWithHTTPClient creates a new PostDisbursementV10BankDisbursementStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDisbursementV10BankDisbursementStatusParamsWithHTTPClient(client *http.Client) *PostDisbursementV10BankDisbursementStatusParams {
	return &PostDisbursementV10BankDisbursementStatusParams{
		HTTPClient: client,
	}
}

/*
PostDisbursementV10BankDisbursementStatusParams contains all the parameters to send to the API endpoint

	for the post disbursement v10 bank disbursement status operation.

	Typically these are written to a http.Request.
*/
type PostDisbursementV10BankDisbursementStatusParams struct {

	/* Request.

	   Status inquiry request
	*/
	Request *models.InquiryDisbursementStatusRequest

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

// WithDefaults hydrates default values in the post disbursement v10 bank disbursement status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDisbursementV10BankDisbursementStatusParams) WithDefaults() *PostDisbursementV10BankDisbursementStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post disbursement v10 bank disbursement status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDisbursementV10BankDisbursementStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post disbursement v10 bank disbursement status params
func (o *PostDisbursementV10BankDisbursementStatusParams) WithTimeout(timeout time.Duration) *PostDisbursementV10BankDisbursementStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post disbursement v10 bank disbursement status params
func (o *PostDisbursementV10BankDisbursementStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post disbursement v10 bank disbursement status params
func (o *PostDisbursementV10BankDisbursementStatusParams) WithContext(ctx context.Context) *PostDisbursementV10BankDisbursementStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post disbursement v10 bank disbursement status params
func (o *PostDisbursementV10BankDisbursementStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post disbursement v10 bank disbursement status params
func (o *PostDisbursementV10BankDisbursementStatusParams) WithHTTPClient(client *http.Client) *PostDisbursementV10BankDisbursementStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post disbursement v10 bank disbursement status params
func (o *PostDisbursementV10BankDisbursementStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post disbursement v10 bank disbursement status params
func (o *PostDisbursementV10BankDisbursementStatusParams) WithRequest(request *models.InquiryDisbursementStatusRequest) *PostDisbursementV10BankDisbursementStatusParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post disbursement v10 bank disbursement status params
func (o *PostDisbursementV10BankDisbursementStatusParams) SetRequest(request *models.InquiryDisbursementStatusRequest) {
	o.Request = request
}

// WithXCLIENTKEY adds the xCLIENTKEY to the post disbursement v10 bank disbursement status params
func (o *PostDisbursementV10BankDisbursementStatusParams) WithXCLIENTKEY(xCLIENTKEY string) *PostDisbursementV10BankDisbursementStatusParams {
	o.SetXCLIENTKEY(xCLIENTKEY)
	return o
}

// SetXCLIENTKEY adds the xCLIENTKEY to the post disbursement v10 bank disbursement status params
func (o *PostDisbursementV10BankDisbursementStatusParams) SetXCLIENTKEY(xCLIENTKEY string) {
	o.XCLIENTKEY = xCLIENTKEY
}

// WithXSIGNATURE adds the xSIGNATURE to the post disbursement v10 bank disbursement status params
func (o *PostDisbursementV10BankDisbursementStatusParams) WithXSIGNATURE(xSIGNATURE string) *PostDisbursementV10BankDisbursementStatusParams {
	o.SetXSIGNATURE(xSIGNATURE)
	return o
}

// SetXSIGNATURE adds the xSIGNATURE to the post disbursement v10 bank disbursement status params
func (o *PostDisbursementV10BankDisbursementStatusParams) SetXSIGNATURE(xSIGNATURE string) {
	o.XSIGNATURE = xSIGNATURE
}

// WithXTIMESTAMP adds the xTIMESTAMP to the post disbursement v10 bank disbursement status params
func (o *PostDisbursementV10BankDisbursementStatusParams) WithXTIMESTAMP(xTIMESTAMP string) *PostDisbursementV10BankDisbursementStatusParams {
	o.SetXTIMESTAMP(xTIMESTAMP)
	return o
}

// SetXTIMESTAMP adds the xTIMESTAMP to the post disbursement v10 bank disbursement status params
func (o *PostDisbursementV10BankDisbursementStatusParams) SetXTIMESTAMP(xTIMESTAMP string) {
	o.XTIMESTAMP = xTIMESTAMP
}

// WriteToRequest writes these params to a swagger request
func (o *PostDisbursementV10BankDisbursementStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
