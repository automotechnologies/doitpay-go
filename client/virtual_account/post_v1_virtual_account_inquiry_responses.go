// Code generated by go-swagger; DO NOT EDIT.

package virtual_account

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

// PostV1VirtualAccountInquiryReader is a Reader for the PostV1VirtualAccountInquiry structure.
type PostV1VirtualAccountInquiryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1VirtualAccountInquiryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostV1VirtualAccountInquiryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1VirtualAccountInquiryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostV1VirtualAccountInquiryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostV1VirtualAccountInquiryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/virtual-account/inquiry] PostV1VirtualAccountInquiry", response, response.Code())
	}
}

// NewPostV1VirtualAccountInquiryOK creates a PostV1VirtualAccountInquiryOK with default headers values
func NewPostV1VirtualAccountInquiryOK() *PostV1VirtualAccountInquiryOK {
	return &PostV1VirtualAccountInquiryOK{}
}

/*
PostV1VirtualAccountInquiryOK describes a response with status code 200, with default header values.

OK
*/
type PostV1VirtualAccountInquiryOK struct {
	Payload *models.Apiv1InquiryVirtualAccountResponse
}

// IsSuccess returns true when this post v1 virtual account inquiry o k response has a 2xx status code
func (o *PostV1VirtualAccountInquiryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 virtual account inquiry o k response has a 3xx status code
func (o *PostV1VirtualAccountInquiryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 virtual account inquiry o k response has a 4xx status code
func (o *PostV1VirtualAccountInquiryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 virtual account inquiry o k response has a 5xx status code
func (o *PostV1VirtualAccountInquiryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 virtual account inquiry o k response a status code equal to that given
func (o *PostV1VirtualAccountInquiryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post v1 virtual account inquiry o k response
func (o *PostV1VirtualAccountInquiryOK) Code() int {
	return 200
}

func (o *PostV1VirtualAccountInquiryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/virtual-account/inquiry][%d] postV1VirtualAccountInquiryOK %s", 200, payload)
}

func (o *PostV1VirtualAccountInquiryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/virtual-account/inquiry][%d] postV1VirtualAccountInquiryOK %s", 200, payload)
}

func (o *PostV1VirtualAccountInquiryOK) GetPayload() *models.Apiv1InquiryVirtualAccountResponse {
	return o.Payload
}

func (o *PostV1VirtualAccountInquiryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Apiv1InquiryVirtualAccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1VirtualAccountInquiryBadRequest creates a PostV1VirtualAccountInquiryBadRequest with default headers values
func NewPostV1VirtualAccountInquiryBadRequest() *PostV1VirtualAccountInquiryBadRequest {
	return &PostV1VirtualAccountInquiryBadRequest{}
}

/*
PostV1VirtualAccountInquiryBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostV1VirtualAccountInquiryBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this post v1 virtual account inquiry bad request response has a 2xx status code
func (o *PostV1VirtualAccountInquiryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 virtual account inquiry bad request response has a 3xx status code
func (o *PostV1VirtualAccountInquiryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 virtual account inquiry bad request response has a 4xx status code
func (o *PostV1VirtualAccountInquiryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 virtual account inquiry bad request response has a 5xx status code
func (o *PostV1VirtualAccountInquiryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 virtual account inquiry bad request response a status code equal to that given
func (o *PostV1VirtualAccountInquiryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post v1 virtual account inquiry bad request response
func (o *PostV1VirtualAccountInquiryBadRequest) Code() int {
	return 400
}

func (o *PostV1VirtualAccountInquiryBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/virtual-account/inquiry][%d] postV1VirtualAccountInquiryBadRequest %s", 400, payload)
}

func (o *PostV1VirtualAccountInquiryBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/virtual-account/inquiry][%d] postV1VirtualAccountInquiryBadRequest %s", 400, payload)
}

func (o *PostV1VirtualAccountInquiryBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *PostV1VirtualAccountInquiryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1VirtualAccountInquiryUnauthorized creates a PostV1VirtualAccountInquiryUnauthorized with default headers values
func NewPostV1VirtualAccountInquiryUnauthorized() *PostV1VirtualAccountInquiryUnauthorized {
	return &PostV1VirtualAccountInquiryUnauthorized{}
}

/*
PostV1VirtualAccountInquiryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostV1VirtualAccountInquiryUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this post v1 virtual account inquiry unauthorized response has a 2xx status code
func (o *PostV1VirtualAccountInquiryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 virtual account inquiry unauthorized response has a 3xx status code
func (o *PostV1VirtualAccountInquiryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 virtual account inquiry unauthorized response has a 4xx status code
func (o *PostV1VirtualAccountInquiryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 virtual account inquiry unauthorized response has a 5xx status code
func (o *PostV1VirtualAccountInquiryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 virtual account inquiry unauthorized response a status code equal to that given
func (o *PostV1VirtualAccountInquiryUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post v1 virtual account inquiry unauthorized response
func (o *PostV1VirtualAccountInquiryUnauthorized) Code() int {
	return 401
}

func (o *PostV1VirtualAccountInquiryUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/virtual-account/inquiry][%d] postV1VirtualAccountInquiryUnauthorized %s", 401, payload)
}

func (o *PostV1VirtualAccountInquiryUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/virtual-account/inquiry][%d] postV1VirtualAccountInquiryUnauthorized %s", 401, payload)
}

func (o *PostV1VirtualAccountInquiryUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *PostV1VirtualAccountInquiryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1VirtualAccountInquiryInternalServerError creates a PostV1VirtualAccountInquiryInternalServerError with default headers values
func NewPostV1VirtualAccountInquiryInternalServerError() *PostV1VirtualAccountInquiryInternalServerError {
	return &PostV1VirtualAccountInquiryInternalServerError{}
}

/*
PostV1VirtualAccountInquiryInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostV1VirtualAccountInquiryInternalServerError struct {
	Payload interface{}
}

// IsSuccess returns true when this post v1 virtual account inquiry internal server error response has a 2xx status code
func (o *PostV1VirtualAccountInquiryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 virtual account inquiry internal server error response has a 3xx status code
func (o *PostV1VirtualAccountInquiryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 virtual account inquiry internal server error response has a 4xx status code
func (o *PostV1VirtualAccountInquiryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 virtual account inquiry internal server error response has a 5xx status code
func (o *PostV1VirtualAccountInquiryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post v1 virtual account inquiry internal server error response a status code equal to that given
func (o *PostV1VirtualAccountInquiryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post v1 virtual account inquiry internal server error response
func (o *PostV1VirtualAccountInquiryInternalServerError) Code() int {
	return 500
}

func (o *PostV1VirtualAccountInquiryInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/virtual-account/inquiry][%d] postV1VirtualAccountInquiryInternalServerError %s", 500, payload)
}

func (o *PostV1VirtualAccountInquiryInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/virtual-account/inquiry][%d] postV1VirtualAccountInquiryInternalServerError %s", 500, payload)
}

func (o *PostV1VirtualAccountInquiryInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *PostV1VirtualAccountInquiryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
