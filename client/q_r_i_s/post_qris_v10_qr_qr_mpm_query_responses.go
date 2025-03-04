// Code generated by go-swagger; DO NOT EDIT.

package q_r_i_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/automotechnologies/doitpay-go/v2/models"
)

// PostQrisV10QrQrMpmQueryReader is a Reader for the PostQrisV10QrQrMpmQuery structure.
type PostQrisV10QrQrMpmQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostQrisV10QrQrMpmQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostQrisV10QrQrMpmQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostQrisV10QrQrMpmQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostQrisV10QrQrMpmQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostQrisV10QrQrMpmQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostQrisV10QrQrMpmQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /qris/v1.0/qr/qr-mpm-query] PostQrisV10QrQrMpmQuery", response, response.Code())
	}
}

// NewPostQrisV10QrQrMpmQueryOK creates a PostQrisV10QrQrMpmQueryOK with default headers values
func NewPostQrisV10QrQrMpmQueryOK() *PostQrisV10QrQrMpmQueryOK {
	return &PostQrisV10QrQrMpmQueryOK{}
}

/*
PostQrisV10QrQrMpmQueryOK describes a response with status code 200, with default header values.

OK
*/
type PostQrisV10QrQrMpmQueryOK struct {
	Payload *models.QrisQueryPaymentResponse
}

// IsSuccess returns true when this post qris v10 qr qr mpm query o k response has a 2xx status code
func (o *PostQrisV10QrQrMpmQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post qris v10 qr qr mpm query o k response has a 3xx status code
func (o *PostQrisV10QrQrMpmQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post qris v10 qr qr mpm query o k response has a 4xx status code
func (o *PostQrisV10QrQrMpmQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post qris v10 qr qr mpm query o k response has a 5xx status code
func (o *PostQrisV10QrQrMpmQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post qris v10 qr qr mpm query o k response a status code equal to that given
func (o *PostQrisV10QrQrMpmQueryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post qris v10 qr qr mpm query o k response
func (o *PostQrisV10QrQrMpmQueryOK) Code() int {
	return 200
}

func (o *PostQrisV10QrQrMpmQueryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /qris/v1.0/qr/qr-mpm-query][%d] postQrisV10QrQrMpmQueryOK %s", 200, payload)
}

func (o *PostQrisV10QrQrMpmQueryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /qris/v1.0/qr/qr-mpm-query][%d] postQrisV10QrQrMpmQueryOK %s", 200, payload)
}

func (o *PostQrisV10QrQrMpmQueryOK) GetPayload() *models.QrisQueryPaymentResponse {
	return o.Payload
}

func (o *PostQrisV10QrQrMpmQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QrisQueryPaymentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQrisV10QrQrMpmQueryBadRequest creates a PostQrisV10QrQrMpmQueryBadRequest with default headers values
func NewPostQrisV10QrQrMpmQueryBadRequest() *PostQrisV10QrQrMpmQueryBadRequest {
	return &PostQrisV10QrQrMpmQueryBadRequest{}
}

/*
PostQrisV10QrQrMpmQueryBadRequest describes a response with status code 400, with default header values.

Invalid request parameters
*/
type PostQrisV10QrQrMpmQueryBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this post qris v10 qr qr mpm query bad request response has a 2xx status code
func (o *PostQrisV10QrQrMpmQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post qris v10 qr qr mpm query bad request response has a 3xx status code
func (o *PostQrisV10QrQrMpmQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post qris v10 qr qr mpm query bad request response has a 4xx status code
func (o *PostQrisV10QrQrMpmQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post qris v10 qr qr mpm query bad request response has a 5xx status code
func (o *PostQrisV10QrQrMpmQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post qris v10 qr qr mpm query bad request response a status code equal to that given
func (o *PostQrisV10QrQrMpmQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post qris v10 qr qr mpm query bad request response
func (o *PostQrisV10QrQrMpmQueryBadRequest) Code() int {
	return 400
}

func (o *PostQrisV10QrQrMpmQueryBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /qris/v1.0/qr/qr-mpm-query][%d] postQrisV10QrQrMpmQueryBadRequest %s", 400, payload)
}

func (o *PostQrisV10QrQrMpmQueryBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /qris/v1.0/qr/qr-mpm-query][%d] postQrisV10QrQrMpmQueryBadRequest %s", 400, payload)
}

func (o *PostQrisV10QrQrMpmQueryBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *PostQrisV10QrQrMpmQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQrisV10QrQrMpmQueryUnauthorized creates a PostQrisV10QrQrMpmQueryUnauthorized with default headers values
func NewPostQrisV10QrQrMpmQueryUnauthorized() *PostQrisV10QrQrMpmQueryUnauthorized {
	return &PostQrisV10QrQrMpmQueryUnauthorized{}
}

/*
PostQrisV10QrQrMpmQueryUnauthorized describes a response with status code 401, with default header values.

Authentication failed
*/
type PostQrisV10QrQrMpmQueryUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this post qris v10 qr qr mpm query unauthorized response has a 2xx status code
func (o *PostQrisV10QrQrMpmQueryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post qris v10 qr qr mpm query unauthorized response has a 3xx status code
func (o *PostQrisV10QrQrMpmQueryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post qris v10 qr qr mpm query unauthorized response has a 4xx status code
func (o *PostQrisV10QrQrMpmQueryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post qris v10 qr qr mpm query unauthorized response has a 5xx status code
func (o *PostQrisV10QrQrMpmQueryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post qris v10 qr qr mpm query unauthorized response a status code equal to that given
func (o *PostQrisV10QrQrMpmQueryUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post qris v10 qr qr mpm query unauthorized response
func (o *PostQrisV10QrQrMpmQueryUnauthorized) Code() int {
	return 401
}

func (o *PostQrisV10QrQrMpmQueryUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /qris/v1.0/qr/qr-mpm-query][%d] postQrisV10QrQrMpmQueryUnauthorized %s", 401, payload)
}

func (o *PostQrisV10QrQrMpmQueryUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /qris/v1.0/qr/qr-mpm-query][%d] postQrisV10QrQrMpmQueryUnauthorized %s", 401, payload)
}

func (o *PostQrisV10QrQrMpmQueryUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *PostQrisV10QrQrMpmQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQrisV10QrQrMpmQueryNotFound creates a PostQrisV10QrQrMpmQueryNotFound with default headers values
func NewPostQrisV10QrQrMpmQueryNotFound() *PostQrisV10QrQrMpmQueryNotFound {
	return &PostQrisV10QrQrMpmQueryNotFound{}
}

/*
PostQrisV10QrQrMpmQueryNotFound describes a response with status code 404, with default header values.

Payment not found
*/
type PostQrisV10QrQrMpmQueryNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this post qris v10 qr qr mpm query not found response has a 2xx status code
func (o *PostQrisV10QrQrMpmQueryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post qris v10 qr qr mpm query not found response has a 3xx status code
func (o *PostQrisV10QrQrMpmQueryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post qris v10 qr qr mpm query not found response has a 4xx status code
func (o *PostQrisV10QrQrMpmQueryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post qris v10 qr qr mpm query not found response has a 5xx status code
func (o *PostQrisV10QrQrMpmQueryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post qris v10 qr qr mpm query not found response a status code equal to that given
func (o *PostQrisV10QrQrMpmQueryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post qris v10 qr qr mpm query not found response
func (o *PostQrisV10QrQrMpmQueryNotFound) Code() int {
	return 404
}

func (o *PostQrisV10QrQrMpmQueryNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /qris/v1.0/qr/qr-mpm-query][%d] postQrisV10QrQrMpmQueryNotFound %s", 404, payload)
}

func (o *PostQrisV10QrQrMpmQueryNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /qris/v1.0/qr/qr-mpm-query][%d] postQrisV10QrQrMpmQueryNotFound %s", 404, payload)
}

func (o *PostQrisV10QrQrMpmQueryNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PostQrisV10QrQrMpmQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQrisV10QrQrMpmQueryInternalServerError creates a PostQrisV10QrQrMpmQueryInternalServerError with default headers values
func NewPostQrisV10QrQrMpmQueryInternalServerError() *PostQrisV10QrQrMpmQueryInternalServerError {
	return &PostQrisV10QrQrMpmQueryInternalServerError{}
}

/*
PostQrisV10QrQrMpmQueryInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostQrisV10QrQrMpmQueryInternalServerError struct {
	Payload interface{}
}

// IsSuccess returns true when this post qris v10 qr qr mpm query internal server error response has a 2xx status code
func (o *PostQrisV10QrQrMpmQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post qris v10 qr qr mpm query internal server error response has a 3xx status code
func (o *PostQrisV10QrQrMpmQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post qris v10 qr qr mpm query internal server error response has a 4xx status code
func (o *PostQrisV10QrQrMpmQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post qris v10 qr qr mpm query internal server error response has a 5xx status code
func (o *PostQrisV10QrQrMpmQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post qris v10 qr qr mpm query internal server error response a status code equal to that given
func (o *PostQrisV10QrQrMpmQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post qris v10 qr qr mpm query internal server error response
func (o *PostQrisV10QrQrMpmQueryInternalServerError) Code() int {
	return 500
}

func (o *PostQrisV10QrQrMpmQueryInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /qris/v1.0/qr/qr-mpm-query][%d] postQrisV10QrQrMpmQueryInternalServerError %s", 500, payload)
}

func (o *PostQrisV10QrQrMpmQueryInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /qris/v1.0/qr/qr-mpm-query][%d] postQrisV10QrQrMpmQueryInternalServerError %s", 500, payload)
}

func (o *PostQrisV10QrQrMpmQueryInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *PostQrisV10QrQrMpmQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
