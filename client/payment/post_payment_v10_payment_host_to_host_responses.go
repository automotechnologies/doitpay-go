// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/automotechnologies/doitpay-go/models"
)

// PostPaymentV10PaymentHostToHostReader is a Reader for the PostPaymentV10PaymentHostToHost structure.
type PostPaymentV10PaymentHostToHostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentV10PaymentHostToHostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostPaymentV10PaymentHostToHostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPaymentV10PaymentHostToHostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostPaymentV10PaymentHostToHostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostPaymentV10PaymentHostToHostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPaymentV10PaymentHostToHostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /payment/v1.0/payment/host-to-host] PostPaymentV10PaymentHostToHost", response, response.Code())
	}
}

// NewPostPaymentV10PaymentHostToHostOK creates a PostPaymentV10PaymentHostToHostOK with default headers values
func NewPostPaymentV10PaymentHostToHostOK() *PostPaymentV10PaymentHostToHostOK {
	return &PostPaymentV10PaymentHostToHostOK{}
}

/*
PostPaymentV10PaymentHostToHostOK describes a response with status code 200, with default header values.

OK
*/
type PostPaymentV10PaymentHostToHostOK struct {
	Payload *models.CreateEwalletResponse
}

// IsSuccess returns true when this post payment v10 payment host to host o k response has a 2xx status code
func (o *PostPaymentV10PaymentHostToHostOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post payment v10 payment host to host o k response has a 3xx status code
func (o *PostPaymentV10PaymentHostToHostOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post payment v10 payment host to host o k response has a 4xx status code
func (o *PostPaymentV10PaymentHostToHostOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post payment v10 payment host to host o k response has a 5xx status code
func (o *PostPaymentV10PaymentHostToHostOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post payment v10 payment host to host o k response a status code equal to that given
func (o *PostPaymentV10PaymentHostToHostOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post payment v10 payment host to host o k response
func (o *PostPaymentV10PaymentHostToHostOK) Code() int {
	return 200
}

func (o *PostPaymentV10PaymentHostToHostOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host][%d] postPaymentV10PaymentHostToHostOK %s", 200, payload)
}

func (o *PostPaymentV10PaymentHostToHostOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host][%d] postPaymentV10PaymentHostToHostOK %s", 200, payload)
}

func (o *PostPaymentV10PaymentHostToHostOK) GetPayload() *models.CreateEwalletResponse {
	return o.Payload
}

func (o *PostPaymentV10PaymentHostToHostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateEwalletResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentV10PaymentHostToHostBadRequest creates a PostPaymentV10PaymentHostToHostBadRequest with default headers values
func NewPostPaymentV10PaymentHostToHostBadRequest() *PostPaymentV10PaymentHostToHostBadRequest {
	return &PostPaymentV10PaymentHostToHostBadRequest{}
}

/*
PostPaymentV10PaymentHostToHostBadRequest describes a response with status code 400, with default header values.

Invalid request parameters
*/
type PostPaymentV10PaymentHostToHostBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this post payment v10 payment host to host bad request response has a 2xx status code
func (o *PostPaymentV10PaymentHostToHostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post payment v10 payment host to host bad request response has a 3xx status code
func (o *PostPaymentV10PaymentHostToHostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post payment v10 payment host to host bad request response has a 4xx status code
func (o *PostPaymentV10PaymentHostToHostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post payment v10 payment host to host bad request response has a 5xx status code
func (o *PostPaymentV10PaymentHostToHostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post payment v10 payment host to host bad request response a status code equal to that given
func (o *PostPaymentV10PaymentHostToHostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post payment v10 payment host to host bad request response
func (o *PostPaymentV10PaymentHostToHostBadRequest) Code() int {
	return 400
}

func (o *PostPaymentV10PaymentHostToHostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host][%d] postPaymentV10PaymentHostToHostBadRequest %s", 400, payload)
}

func (o *PostPaymentV10PaymentHostToHostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host][%d] postPaymentV10PaymentHostToHostBadRequest %s", 400, payload)
}

func (o *PostPaymentV10PaymentHostToHostBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *PostPaymentV10PaymentHostToHostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentV10PaymentHostToHostUnauthorized creates a PostPaymentV10PaymentHostToHostUnauthorized with default headers values
func NewPostPaymentV10PaymentHostToHostUnauthorized() *PostPaymentV10PaymentHostToHostUnauthorized {
	return &PostPaymentV10PaymentHostToHostUnauthorized{}
}

/*
PostPaymentV10PaymentHostToHostUnauthorized describes a response with status code 401, with default header values.

Authentication failed
*/
type PostPaymentV10PaymentHostToHostUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this post payment v10 payment host to host unauthorized response has a 2xx status code
func (o *PostPaymentV10PaymentHostToHostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post payment v10 payment host to host unauthorized response has a 3xx status code
func (o *PostPaymentV10PaymentHostToHostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post payment v10 payment host to host unauthorized response has a 4xx status code
func (o *PostPaymentV10PaymentHostToHostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post payment v10 payment host to host unauthorized response has a 5xx status code
func (o *PostPaymentV10PaymentHostToHostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post payment v10 payment host to host unauthorized response a status code equal to that given
func (o *PostPaymentV10PaymentHostToHostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post payment v10 payment host to host unauthorized response
func (o *PostPaymentV10PaymentHostToHostUnauthorized) Code() int {
	return 401
}

func (o *PostPaymentV10PaymentHostToHostUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host][%d] postPaymentV10PaymentHostToHostUnauthorized %s", 401, payload)
}

func (o *PostPaymentV10PaymentHostToHostUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host][%d] postPaymentV10PaymentHostToHostUnauthorized %s", 401, payload)
}

func (o *PostPaymentV10PaymentHostToHostUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *PostPaymentV10PaymentHostToHostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentV10PaymentHostToHostConflict creates a PostPaymentV10PaymentHostToHostConflict with default headers values
func NewPostPaymentV10PaymentHostToHostConflict() *PostPaymentV10PaymentHostToHostConflict {
	return &PostPaymentV10PaymentHostToHostConflict{}
}

/*
PostPaymentV10PaymentHostToHostConflict describes a response with status code 409, with default header values.

Duplicate transaction
*/
type PostPaymentV10PaymentHostToHostConflict struct {
	Payload interface{}
}

// IsSuccess returns true when this post payment v10 payment host to host conflict response has a 2xx status code
func (o *PostPaymentV10PaymentHostToHostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post payment v10 payment host to host conflict response has a 3xx status code
func (o *PostPaymentV10PaymentHostToHostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post payment v10 payment host to host conflict response has a 4xx status code
func (o *PostPaymentV10PaymentHostToHostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post payment v10 payment host to host conflict response has a 5xx status code
func (o *PostPaymentV10PaymentHostToHostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post payment v10 payment host to host conflict response a status code equal to that given
func (o *PostPaymentV10PaymentHostToHostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the post payment v10 payment host to host conflict response
func (o *PostPaymentV10PaymentHostToHostConflict) Code() int {
	return 409
}

func (o *PostPaymentV10PaymentHostToHostConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host][%d] postPaymentV10PaymentHostToHostConflict %s", 409, payload)
}

func (o *PostPaymentV10PaymentHostToHostConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host][%d] postPaymentV10PaymentHostToHostConflict %s", 409, payload)
}

func (o *PostPaymentV10PaymentHostToHostConflict) GetPayload() interface{} {
	return o.Payload
}

func (o *PostPaymentV10PaymentHostToHostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentV10PaymentHostToHostInternalServerError creates a PostPaymentV10PaymentHostToHostInternalServerError with default headers values
func NewPostPaymentV10PaymentHostToHostInternalServerError() *PostPaymentV10PaymentHostToHostInternalServerError {
	return &PostPaymentV10PaymentHostToHostInternalServerError{}
}

/*
PostPaymentV10PaymentHostToHostInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostPaymentV10PaymentHostToHostInternalServerError struct {
	Payload interface{}
}

// IsSuccess returns true when this post payment v10 payment host to host internal server error response has a 2xx status code
func (o *PostPaymentV10PaymentHostToHostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post payment v10 payment host to host internal server error response has a 3xx status code
func (o *PostPaymentV10PaymentHostToHostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post payment v10 payment host to host internal server error response has a 4xx status code
func (o *PostPaymentV10PaymentHostToHostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post payment v10 payment host to host internal server error response has a 5xx status code
func (o *PostPaymentV10PaymentHostToHostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post payment v10 payment host to host internal server error response a status code equal to that given
func (o *PostPaymentV10PaymentHostToHostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post payment v10 payment host to host internal server error response
func (o *PostPaymentV10PaymentHostToHostInternalServerError) Code() int {
	return 500
}

func (o *PostPaymentV10PaymentHostToHostInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host][%d] postPaymentV10PaymentHostToHostInternalServerError %s", 500, payload)
}

func (o *PostPaymentV10PaymentHostToHostInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host][%d] postPaymentV10PaymentHostToHostInternalServerError %s", 500, payload)
}

func (o *PostPaymentV10PaymentHostToHostInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *PostPaymentV10PaymentHostToHostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
