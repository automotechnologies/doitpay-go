// Code generated by go-swagger; DO NOT EDIT.

package disbursement

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

// PostDisbursementV10EmoneyDisbursementStatusReader is a Reader for the PostDisbursementV10EmoneyDisbursementStatus structure.
type PostDisbursementV10EmoneyDisbursementStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDisbursementV10EmoneyDisbursementStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDisbursementV10EmoneyDisbursementStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDisbursementV10EmoneyDisbursementStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDisbursementV10EmoneyDisbursementStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDisbursementV10EmoneyDisbursementStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDisbursementV10EmoneyDisbursementStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /disbursement/v1.0/emoney-disbursement/status] PostDisbursementV10EmoneyDisbursementStatus", response, response.Code())
	}
}

// NewPostDisbursementV10EmoneyDisbursementStatusOK creates a PostDisbursementV10EmoneyDisbursementStatusOK with default headers values
func NewPostDisbursementV10EmoneyDisbursementStatusOK() *PostDisbursementV10EmoneyDisbursementStatusOK {
	return &PostDisbursementV10EmoneyDisbursementStatusOK{}
}

/*
PostDisbursementV10EmoneyDisbursementStatusOK describes a response with status code 200, with default header values.

OK
*/
type PostDisbursementV10EmoneyDisbursementStatusOK struct {
	Payload *models.InquiryEwalletTopupStatusResponse
}

// IsSuccess returns true when this post disbursement v10 emoney disbursement status o k response has a 2xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post disbursement v10 emoney disbursement status o k response has a 3xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post disbursement v10 emoney disbursement status o k response has a 4xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post disbursement v10 emoney disbursement status o k response has a 5xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post disbursement v10 emoney disbursement status o k response a status code equal to that given
func (o *PostDisbursementV10EmoneyDisbursementStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post disbursement v10 emoney disbursement status o k response
func (o *PostDisbursementV10EmoneyDisbursementStatusOK) Code() int {
	return 200
}

func (o *PostDisbursementV10EmoneyDisbursementStatusOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /disbursement/v1.0/emoney-disbursement/status][%d] postDisbursementV10EmoneyDisbursementStatusOK %s", 200, payload)
}

func (o *PostDisbursementV10EmoneyDisbursementStatusOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /disbursement/v1.0/emoney-disbursement/status][%d] postDisbursementV10EmoneyDisbursementStatusOK %s", 200, payload)
}

func (o *PostDisbursementV10EmoneyDisbursementStatusOK) GetPayload() *models.InquiryEwalletTopupStatusResponse {
	return o.Payload
}

func (o *PostDisbursementV10EmoneyDisbursementStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InquiryEwalletTopupStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDisbursementV10EmoneyDisbursementStatusBadRequest creates a PostDisbursementV10EmoneyDisbursementStatusBadRequest with default headers values
func NewPostDisbursementV10EmoneyDisbursementStatusBadRequest() *PostDisbursementV10EmoneyDisbursementStatusBadRequest {
	return &PostDisbursementV10EmoneyDisbursementStatusBadRequest{}
}

/*
PostDisbursementV10EmoneyDisbursementStatusBadRequest describes a response with status code 400, with default header values.

Invalid request parameters
*/
type PostDisbursementV10EmoneyDisbursementStatusBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this post disbursement v10 emoney disbursement status bad request response has a 2xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post disbursement v10 emoney disbursement status bad request response has a 3xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post disbursement v10 emoney disbursement status bad request response has a 4xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post disbursement v10 emoney disbursement status bad request response has a 5xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post disbursement v10 emoney disbursement status bad request response a status code equal to that given
func (o *PostDisbursementV10EmoneyDisbursementStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post disbursement v10 emoney disbursement status bad request response
func (o *PostDisbursementV10EmoneyDisbursementStatusBadRequest) Code() int {
	return 400
}

func (o *PostDisbursementV10EmoneyDisbursementStatusBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /disbursement/v1.0/emoney-disbursement/status][%d] postDisbursementV10EmoneyDisbursementStatusBadRequest %s", 400, payload)
}

func (o *PostDisbursementV10EmoneyDisbursementStatusBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /disbursement/v1.0/emoney-disbursement/status][%d] postDisbursementV10EmoneyDisbursementStatusBadRequest %s", 400, payload)
}

func (o *PostDisbursementV10EmoneyDisbursementStatusBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *PostDisbursementV10EmoneyDisbursementStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDisbursementV10EmoneyDisbursementStatusUnauthorized creates a PostDisbursementV10EmoneyDisbursementStatusUnauthorized with default headers values
func NewPostDisbursementV10EmoneyDisbursementStatusUnauthorized() *PostDisbursementV10EmoneyDisbursementStatusUnauthorized {
	return &PostDisbursementV10EmoneyDisbursementStatusUnauthorized{}
}

/*
PostDisbursementV10EmoneyDisbursementStatusUnauthorized describes a response with status code 401, with default header values.

Authentication failed
*/
type PostDisbursementV10EmoneyDisbursementStatusUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this post disbursement v10 emoney disbursement status unauthorized response has a 2xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post disbursement v10 emoney disbursement status unauthorized response has a 3xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post disbursement v10 emoney disbursement status unauthorized response has a 4xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post disbursement v10 emoney disbursement status unauthorized response has a 5xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post disbursement v10 emoney disbursement status unauthorized response a status code equal to that given
func (o *PostDisbursementV10EmoneyDisbursementStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post disbursement v10 emoney disbursement status unauthorized response
func (o *PostDisbursementV10EmoneyDisbursementStatusUnauthorized) Code() int {
	return 401
}

func (o *PostDisbursementV10EmoneyDisbursementStatusUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /disbursement/v1.0/emoney-disbursement/status][%d] postDisbursementV10EmoneyDisbursementStatusUnauthorized %s", 401, payload)
}

func (o *PostDisbursementV10EmoneyDisbursementStatusUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /disbursement/v1.0/emoney-disbursement/status][%d] postDisbursementV10EmoneyDisbursementStatusUnauthorized %s", 401, payload)
}

func (o *PostDisbursementV10EmoneyDisbursementStatusUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *PostDisbursementV10EmoneyDisbursementStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDisbursementV10EmoneyDisbursementStatusNotFound creates a PostDisbursementV10EmoneyDisbursementStatusNotFound with default headers values
func NewPostDisbursementV10EmoneyDisbursementStatusNotFound() *PostDisbursementV10EmoneyDisbursementStatusNotFound {
	return &PostDisbursementV10EmoneyDisbursementStatusNotFound{}
}

/*
PostDisbursementV10EmoneyDisbursementStatusNotFound describes a response with status code 404, with default header values.

Disbursement not found
*/
type PostDisbursementV10EmoneyDisbursementStatusNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this post disbursement v10 emoney disbursement status not found response has a 2xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post disbursement v10 emoney disbursement status not found response has a 3xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post disbursement v10 emoney disbursement status not found response has a 4xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post disbursement v10 emoney disbursement status not found response has a 5xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post disbursement v10 emoney disbursement status not found response a status code equal to that given
func (o *PostDisbursementV10EmoneyDisbursementStatusNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post disbursement v10 emoney disbursement status not found response
func (o *PostDisbursementV10EmoneyDisbursementStatusNotFound) Code() int {
	return 404
}

func (o *PostDisbursementV10EmoneyDisbursementStatusNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /disbursement/v1.0/emoney-disbursement/status][%d] postDisbursementV10EmoneyDisbursementStatusNotFound %s", 404, payload)
}

func (o *PostDisbursementV10EmoneyDisbursementStatusNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /disbursement/v1.0/emoney-disbursement/status][%d] postDisbursementV10EmoneyDisbursementStatusNotFound %s", 404, payload)
}

func (o *PostDisbursementV10EmoneyDisbursementStatusNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PostDisbursementV10EmoneyDisbursementStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDisbursementV10EmoneyDisbursementStatusInternalServerError creates a PostDisbursementV10EmoneyDisbursementStatusInternalServerError with default headers values
func NewPostDisbursementV10EmoneyDisbursementStatusInternalServerError() *PostDisbursementV10EmoneyDisbursementStatusInternalServerError {
	return &PostDisbursementV10EmoneyDisbursementStatusInternalServerError{}
}

/*
PostDisbursementV10EmoneyDisbursementStatusInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostDisbursementV10EmoneyDisbursementStatusInternalServerError struct {
	Payload interface{}
}

// IsSuccess returns true when this post disbursement v10 emoney disbursement status internal server error response has a 2xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post disbursement v10 emoney disbursement status internal server error response has a 3xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post disbursement v10 emoney disbursement status internal server error response has a 4xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post disbursement v10 emoney disbursement status internal server error response has a 5xx status code
func (o *PostDisbursementV10EmoneyDisbursementStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post disbursement v10 emoney disbursement status internal server error response a status code equal to that given
func (o *PostDisbursementV10EmoneyDisbursementStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post disbursement v10 emoney disbursement status internal server error response
func (o *PostDisbursementV10EmoneyDisbursementStatusInternalServerError) Code() int {
	return 500
}

func (o *PostDisbursementV10EmoneyDisbursementStatusInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /disbursement/v1.0/emoney-disbursement/status][%d] postDisbursementV10EmoneyDisbursementStatusInternalServerError %s", 500, payload)
}

func (o *PostDisbursementV10EmoneyDisbursementStatusInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /disbursement/v1.0/emoney-disbursement/status][%d] postDisbursementV10EmoneyDisbursementStatusInternalServerError %s", 500, payload)
}

func (o *PostDisbursementV10EmoneyDisbursementStatusInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *PostDisbursementV10EmoneyDisbursementStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
