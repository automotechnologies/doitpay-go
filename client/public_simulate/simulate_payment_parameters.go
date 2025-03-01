// Code generated by go-swagger; DO NOT EDIT.

package public_simulate

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

// NewSimulatePaymentParams creates a new SimulatePaymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSimulatePaymentParams() *SimulatePaymentParams {
	return &SimulatePaymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSimulatePaymentParamsWithTimeout creates a new SimulatePaymentParams object
// with the ability to set a timeout on a request.
func NewSimulatePaymentParamsWithTimeout(timeout time.Duration) *SimulatePaymentParams {
	return &SimulatePaymentParams{
		timeout: timeout,
	}
}

// NewSimulatePaymentParamsWithContext creates a new SimulatePaymentParams object
// with the ability to set a context for a request.
func NewSimulatePaymentParamsWithContext(ctx context.Context) *SimulatePaymentParams {
	return &SimulatePaymentParams{
		Context: ctx,
	}
}

// NewSimulatePaymentParamsWithHTTPClient creates a new SimulatePaymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewSimulatePaymentParamsWithHTTPClient(client *http.Client) *SimulatePaymentParams {
	return &SimulatePaymentParams{
		HTTPClient: client,
	}
}

/*
SimulatePaymentParams contains all the parameters to send to the API endpoint

	for the simulate payment operation.

	Typically these are written to a http.Request.
*/
type SimulatePaymentParams struct {

	/* Request.

	   Request payload to simulate payment
	*/
	Request *models.Apiv1QrisSimulateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the simulate payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SimulatePaymentParams) WithDefaults() *SimulatePaymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the simulate payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SimulatePaymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the simulate payment params
func (o *SimulatePaymentParams) WithTimeout(timeout time.Duration) *SimulatePaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the simulate payment params
func (o *SimulatePaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the simulate payment params
func (o *SimulatePaymentParams) WithContext(ctx context.Context) *SimulatePaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the simulate payment params
func (o *SimulatePaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the simulate payment params
func (o *SimulatePaymentParams) WithHTTPClient(client *http.Client) *SimulatePaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the simulate payment params
func (o *SimulatePaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the simulate payment params
func (o *SimulatePaymentParams) WithRequest(request *models.Apiv1QrisSimulateRequest) *SimulatePaymentParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the simulate payment params
func (o *SimulatePaymentParams) SetRequest(request *models.Apiv1QrisSimulateRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *SimulatePaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
