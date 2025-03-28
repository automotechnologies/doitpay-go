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

// PostPaymentV10PaymentHostToHostCancelReader is a Reader for the PostPaymentV10PaymentHostToHostCancel structure.
type PostPaymentV10PaymentHostToHostCancelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPaymentV10PaymentHostToHostCancelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostPaymentV10PaymentHostToHostCancelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostPaymentV10PaymentHostToHostCancelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostPaymentV10PaymentHostToHostCancelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostPaymentV10PaymentHostToHostCancelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostPaymentV10PaymentHostToHostCancelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /payment/v1.0/payment/host-to-host/cancel] PostPaymentV10PaymentHostToHostCancel", response, response.Code())
	}
}

// NewPostPaymentV10PaymentHostToHostCancelOK creates a PostPaymentV10PaymentHostToHostCancelOK with default headers values
func NewPostPaymentV10PaymentHostToHostCancelOK() *PostPaymentV10PaymentHostToHostCancelOK {
	return &PostPaymentV10PaymentHostToHostCancelOK{}
}

/*
PostPaymentV10PaymentHostToHostCancelOK describes a response with status code 200, with default header values.

OK
*/
type PostPaymentV10PaymentHostToHostCancelOK struct {
	Payload *models.CancelHostToHostPaymentResponse
}

// IsSuccess returns true when this post payment v10 payment host to host cancel o k response has a 2xx status code
func (o *PostPaymentV10PaymentHostToHostCancelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post payment v10 payment host to host cancel o k response has a 3xx status code
func (o *PostPaymentV10PaymentHostToHostCancelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post payment v10 payment host to host cancel o k response has a 4xx status code
func (o *PostPaymentV10PaymentHostToHostCancelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post payment v10 payment host to host cancel o k response has a 5xx status code
func (o *PostPaymentV10PaymentHostToHostCancelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post payment v10 payment host to host cancel o k response a status code equal to that given
func (o *PostPaymentV10PaymentHostToHostCancelOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post payment v10 payment host to host cancel o k response
func (o *PostPaymentV10PaymentHostToHostCancelOK) Code() int {
	return 200
}

func (o *PostPaymentV10PaymentHostToHostCancelOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host/cancel][%d] postPaymentV10PaymentHostToHostCancelOK %s", 200, payload)
}

func (o *PostPaymentV10PaymentHostToHostCancelOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host/cancel][%d] postPaymentV10PaymentHostToHostCancelOK %s", 200, payload)
}

func (o *PostPaymentV10PaymentHostToHostCancelOK) GetPayload() *models.CancelHostToHostPaymentResponse {
	return o.Payload
}

func (o *PostPaymentV10PaymentHostToHostCancelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CancelHostToHostPaymentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentV10PaymentHostToHostCancelBadRequest creates a PostPaymentV10PaymentHostToHostCancelBadRequest with default headers values
func NewPostPaymentV10PaymentHostToHostCancelBadRequest() *PostPaymentV10PaymentHostToHostCancelBadRequest {
	return &PostPaymentV10PaymentHostToHostCancelBadRequest{}
}

/*
PostPaymentV10PaymentHostToHostCancelBadRequest describes a response with status code 400, with default header values.

Invalid request parameters
*/
type PostPaymentV10PaymentHostToHostCancelBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this post payment v10 payment host to host cancel bad request response has a 2xx status code
func (o *PostPaymentV10PaymentHostToHostCancelBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post payment v10 payment host to host cancel bad request response has a 3xx status code
func (o *PostPaymentV10PaymentHostToHostCancelBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post payment v10 payment host to host cancel bad request response has a 4xx status code
func (o *PostPaymentV10PaymentHostToHostCancelBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post payment v10 payment host to host cancel bad request response has a 5xx status code
func (o *PostPaymentV10PaymentHostToHostCancelBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post payment v10 payment host to host cancel bad request response a status code equal to that given
func (o *PostPaymentV10PaymentHostToHostCancelBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post payment v10 payment host to host cancel bad request response
func (o *PostPaymentV10PaymentHostToHostCancelBadRequest) Code() int {
	return 400
}

func (o *PostPaymentV10PaymentHostToHostCancelBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host/cancel][%d] postPaymentV10PaymentHostToHostCancelBadRequest %s", 400, payload)
}

func (o *PostPaymentV10PaymentHostToHostCancelBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host/cancel][%d] postPaymentV10PaymentHostToHostCancelBadRequest %s", 400, payload)
}

func (o *PostPaymentV10PaymentHostToHostCancelBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *PostPaymentV10PaymentHostToHostCancelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentV10PaymentHostToHostCancelUnauthorized creates a PostPaymentV10PaymentHostToHostCancelUnauthorized with default headers values
func NewPostPaymentV10PaymentHostToHostCancelUnauthorized() *PostPaymentV10PaymentHostToHostCancelUnauthorized {
	return &PostPaymentV10PaymentHostToHostCancelUnauthorized{}
}

/*
PostPaymentV10PaymentHostToHostCancelUnauthorized describes a response with status code 401, with default header values.

Authentication failed
*/
type PostPaymentV10PaymentHostToHostCancelUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this post payment v10 payment host to host cancel unauthorized response has a 2xx status code
func (o *PostPaymentV10PaymentHostToHostCancelUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post payment v10 payment host to host cancel unauthorized response has a 3xx status code
func (o *PostPaymentV10PaymentHostToHostCancelUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post payment v10 payment host to host cancel unauthorized response has a 4xx status code
func (o *PostPaymentV10PaymentHostToHostCancelUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post payment v10 payment host to host cancel unauthorized response has a 5xx status code
func (o *PostPaymentV10PaymentHostToHostCancelUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post payment v10 payment host to host cancel unauthorized response a status code equal to that given
func (o *PostPaymentV10PaymentHostToHostCancelUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post payment v10 payment host to host cancel unauthorized response
func (o *PostPaymentV10PaymentHostToHostCancelUnauthorized) Code() int {
	return 401
}

func (o *PostPaymentV10PaymentHostToHostCancelUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host/cancel][%d] postPaymentV10PaymentHostToHostCancelUnauthorized %s", 401, payload)
}

func (o *PostPaymentV10PaymentHostToHostCancelUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host/cancel][%d] postPaymentV10PaymentHostToHostCancelUnauthorized %s", 401, payload)
}

func (o *PostPaymentV10PaymentHostToHostCancelUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *PostPaymentV10PaymentHostToHostCancelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentV10PaymentHostToHostCancelNotFound creates a PostPaymentV10PaymentHostToHostCancelNotFound with default headers values
func NewPostPaymentV10PaymentHostToHostCancelNotFound() *PostPaymentV10PaymentHostToHostCancelNotFound {
	return &PostPaymentV10PaymentHostToHostCancelNotFound{}
}

/*
PostPaymentV10PaymentHostToHostCancelNotFound describes a response with status code 404, with default header values.

Payment not found
*/
type PostPaymentV10PaymentHostToHostCancelNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this post payment v10 payment host to host cancel not found response has a 2xx status code
func (o *PostPaymentV10PaymentHostToHostCancelNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post payment v10 payment host to host cancel not found response has a 3xx status code
func (o *PostPaymentV10PaymentHostToHostCancelNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post payment v10 payment host to host cancel not found response has a 4xx status code
func (o *PostPaymentV10PaymentHostToHostCancelNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post payment v10 payment host to host cancel not found response has a 5xx status code
func (o *PostPaymentV10PaymentHostToHostCancelNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post payment v10 payment host to host cancel not found response a status code equal to that given
func (o *PostPaymentV10PaymentHostToHostCancelNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post payment v10 payment host to host cancel not found response
func (o *PostPaymentV10PaymentHostToHostCancelNotFound) Code() int {
	return 404
}

func (o *PostPaymentV10PaymentHostToHostCancelNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host/cancel][%d] postPaymentV10PaymentHostToHostCancelNotFound %s", 404, payload)
}

func (o *PostPaymentV10PaymentHostToHostCancelNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host/cancel][%d] postPaymentV10PaymentHostToHostCancelNotFound %s", 404, payload)
}

func (o *PostPaymentV10PaymentHostToHostCancelNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PostPaymentV10PaymentHostToHostCancelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentV10PaymentHostToHostCancelInternalServerError creates a PostPaymentV10PaymentHostToHostCancelInternalServerError with default headers values
func NewPostPaymentV10PaymentHostToHostCancelInternalServerError() *PostPaymentV10PaymentHostToHostCancelInternalServerError {
	return &PostPaymentV10PaymentHostToHostCancelInternalServerError{}
}

/*
PostPaymentV10PaymentHostToHostCancelInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostPaymentV10PaymentHostToHostCancelInternalServerError struct {
	Payload interface{}
}

// IsSuccess returns true when this post payment v10 payment host to host cancel internal server error response has a 2xx status code
func (o *PostPaymentV10PaymentHostToHostCancelInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post payment v10 payment host to host cancel internal server error response has a 3xx status code
func (o *PostPaymentV10PaymentHostToHostCancelInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post payment v10 payment host to host cancel internal server error response has a 4xx status code
func (o *PostPaymentV10PaymentHostToHostCancelInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post payment v10 payment host to host cancel internal server error response has a 5xx status code
func (o *PostPaymentV10PaymentHostToHostCancelInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post payment v10 payment host to host cancel internal server error response a status code equal to that given
func (o *PostPaymentV10PaymentHostToHostCancelInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post payment v10 payment host to host cancel internal server error response
func (o *PostPaymentV10PaymentHostToHostCancelInternalServerError) Code() int {
	return 500
}

func (o *PostPaymentV10PaymentHostToHostCancelInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host/cancel][%d] postPaymentV10PaymentHostToHostCancelInternalServerError %s", 500, payload)
}

func (o *PostPaymentV10PaymentHostToHostCancelInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /payment/v1.0/payment/host-to-host/cancel][%d] postPaymentV10PaymentHostToHostCancelInternalServerError %s", 500, payload)
}

func (o *PostPaymentV10PaymentHostToHostCancelInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *PostPaymentV10PaymentHostToHostCancelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
