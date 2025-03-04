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

// NewPostDisbursementV10BalanceParams creates a new PostDisbursementV10BalanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDisbursementV10BalanceParams() *PostDisbursementV10BalanceParams {
	return &PostDisbursementV10BalanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDisbursementV10BalanceParamsWithTimeout creates a new PostDisbursementV10BalanceParams object
// with the ability to set a timeout on a request.
func NewPostDisbursementV10BalanceParamsWithTimeout(timeout time.Duration) *PostDisbursementV10BalanceParams {
	return &PostDisbursementV10BalanceParams{
		timeout: timeout,
	}
}

// NewPostDisbursementV10BalanceParamsWithContext creates a new PostDisbursementV10BalanceParams object
// with the ability to set a context for a request.
func NewPostDisbursementV10BalanceParamsWithContext(ctx context.Context) *PostDisbursementV10BalanceParams {
	return &PostDisbursementV10BalanceParams{
		Context: ctx,
	}
}

// NewPostDisbursementV10BalanceParamsWithHTTPClient creates a new PostDisbursementV10BalanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDisbursementV10BalanceParamsWithHTTPClient(client *http.Client) *PostDisbursementV10BalanceParams {
	return &PostDisbursementV10BalanceParams{
		HTTPClient: client,
	}
}

/*
PostDisbursementV10BalanceParams contains all the parameters to send to the API endpoint

	for the post disbursement v10 balance operation.

	Typically these are written to a http.Request.
*/
type PostDisbursementV10BalanceParams struct {

	/* Request.

	   Balance inquiry request
	*/
	Request *models.InquiryBalanceRequest

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

// WithDefaults hydrates default values in the post disbursement v10 balance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDisbursementV10BalanceParams) WithDefaults() *PostDisbursementV10BalanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post disbursement v10 balance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDisbursementV10BalanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post disbursement v10 balance params
func (o *PostDisbursementV10BalanceParams) WithTimeout(timeout time.Duration) *PostDisbursementV10BalanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post disbursement v10 balance params
func (o *PostDisbursementV10BalanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post disbursement v10 balance params
func (o *PostDisbursementV10BalanceParams) WithContext(ctx context.Context) *PostDisbursementV10BalanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post disbursement v10 balance params
func (o *PostDisbursementV10BalanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post disbursement v10 balance params
func (o *PostDisbursementV10BalanceParams) WithHTTPClient(client *http.Client) *PostDisbursementV10BalanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post disbursement v10 balance params
func (o *PostDisbursementV10BalanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post disbursement v10 balance params
func (o *PostDisbursementV10BalanceParams) WithRequest(request *models.InquiryBalanceRequest) *PostDisbursementV10BalanceParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post disbursement v10 balance params
func (o *PostDisbursementV10BalanceParams) SetRequest(request *models.InquiryBalanceRequest) {
	o.Request = request
}

// WithXCLIENTKEY adds the xCLIENTKEY to the post disbursement v10 balance params
func (o *PostDisbursementV10BalanceParams) WithXCLIENTKEY(xCLIENTKEY string) *PostDisbursementV10BalanceParams {
	o.SetXCLIENTKEY(xCLIENTKEY)
	return o
}

// SetXCLIENTKEY adds the xCLIENTKEY to the post disbursement v10 balance params
func (o *PostDisbursementV10BalanceParams) SetXCLIENTKEY(xCLIENTKEY string) {
	o.XCLIENTKEY = xCLIENTKEY
}

// WithXSIGNATURE adds the xSIGNATURE to the post disbursement v10 balance params
func (o *PostDisbursementV10BalanceParams) WithXSIGNATURE(xSIGNATURE string) *PostDisbursementV10BalanceParams {
	o.SetXSIGNATURE(xSIGNATURE)
	return o
}

// SetXSIGNATURE adds the xSIGNATURE to the post disbursement v10 balance params
func (o *PostDisbursementV10BalanceParams) SetXSIGNATURE(xSIGNATURE string) {
	o.XSIGNATURE = xSIGNATURE
}

// WithXTIMESTAMP adds the xTIMESTAMP to the post disbursement v10 balance params
func (o *PostDisbursementV10BalanceParams) WithXTIMESTAMP(xTIMESTAMP string) *PostDisbursementV10BalanceParams {
	o.SetXTIMESTAMP(xTIMESTAMP)
	return o
}

// SetXTIMESTAMP adds the xTIMESTAMP to the post disbursement v10 balance params
func (o *PostDisbursementV10BalanceParams) SetXTIMESTAMP(xTIMESTAMP string) {
	o.XTIMESTAMP = xTIMESTAMP
}

// WriteToRequest writes these params to a swagger request
func (o *PostDisbursementV10BalanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
