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

// DeleteVaV10TransferVaDeleteVaReader is a Reader for the DeleteVaV10TransferVaDeleteVa structure.
type DeleteVaV10TransferVaDeleteVaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVaV10TransferVaDeleteVaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVaV10TransferVaDeleteVaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVaV10TransferVaDeleteVaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteVaV10TransferVaDeleteVaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteVaV10TransferVaDeleteVaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /va/v1.0/transfer-va/delete-va] DeleteVaV10TransferVaDeleteVa", response, response.Code())
	}
}

// NewDeleteVaV10TransferVaDeleteVaOK creates a DeleteVaV10TransferVaDeleteVaOK with default headers values
func NewDeleteVaV10TransferVaDeleteVaOK() *DeleteVaV10TransferVaDeleteVaOK {
	return &DeleteVaV10TransferVaDeleteVaOK{}
}

/*
DeleteVaV10TransferVaDeleteVaOK describes a response with status code 200, with default header values.

OK
*/
type DeleteVaV10TransferVaDeleteVaOK struct {
	Payload *models.DeleteVirtualAccountResponse
}

// IsSuccess returns true when this delete va v10 transfer va delete va o k response has a 2xx status code
func (o *DeleteVaV10TransferVaDeleteVaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete va v10 transfer va delete va o k response has a 3xx status code
func (o *DeleteVaV10TransferVaDeleteVaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete va v10 transfer va delete va o k response has a 4xx status code
func (o *DeleteVaV10TransferVaDeleteVaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete va v10 transfer va delete va o k response has a 5xx status code
func (o *DeleteVaV10TransferVaDeleteVaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete va v10 transfer va delete va o k response a status code equal to that given
func (o *DeleteVaV10TransferVaDeleteVaOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete va v10 transfer va delete va o k response
func (o *DeleteVaV10TransferVaDeleteVaOK) Code() int {
	return 200
}

func (o *DeleteVaV10TransferVaDeleteVaOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /va/v1.0/transfer-va/delete-va][%d] deleteVaV10TransferVaDeleteVaOK %s", 200, payload)
}

func (o *DeleteVaV10TransferVaDeleteVaOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /va/v1.0/transfer-va/delete-va][%d] deleteVaV10TransferVaDeleteVaOK %s", 200, payload)
}

func (o *DeleteVaV10TransferVaDeleteVaOK) GetPayload() *models.DeleteVirtualAccountResponse {
	return o.Payload
}

func (o *DeleteVaV10TransferVaDeleteVaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteVirtualAccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVaV10TransferVaDeleteVaBadRequest creates a DeleteVaV10TransferVaDeleteVaBadRequest with default headers values
func NewDeleteVaV10TransferVaDeleteVaBadRequest() *DeleteVaV10TransferVaDeleteVaBadRequest {
	return &DeleteVaV10TransferVaDeleteVaBadRequest{}
}

/*
DeleteVaV10TransferVaDeleteVaBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteVaV10TransferVaDeleteVaBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this delete va v10 transfer va delete va bad request response has a 2xx status code
func (o *DeleteVaV10TransferVaDeleteVaBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete va v10 transfer va delete va bad request response has a 3xx status code
func (o *DeleteVaV10TransferVaDeleteVaBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete va v10 transfer va delete va bad request response has a 4xx status code
func (o *DeleteVaV10TransferVaDeleteVaBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete va v10 transfer va delete va bad request response has a 5xx status code
func (o *DeleteVaV10TransferVaDeleteVaBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete va v10 transfer va delete va bad request response a status code equal to that given
func (o *DeleteVaV10TransferVaDeleteVaBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete va v10 transfer va delete va bad request response
func (o *DeleteVaV10TransferVaDeleteVaBadRequest) Code() int {
	return 400
}

func (o *DeleteVaV10TransferVaDeleteVaBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /va/v1.0/transfer-va/delete-va][%d] deleteVaV10TransferVaDeleteVaBadRequest %s", 400, payload)
}

func (o *DeleteVaV10TransferVaDeleteVaBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /va/v1.0/transfer-va/delete-va][%d] deleteVaV10TransferVaDeleteVaBadRequest %s", 400, payload)
}

func (o *DeleteVaV10TransferVaDeleteVaBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteVaV10TransferVaDeleteVaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVaV10TransferVaDeleteVaUnauthorized creates a DeleteVaV10TransferVaDeleteVaUnauthorized with default headers values
func NewDeleteVaV10TransferVaDeleteVaUnauthorized() *DeleteVaV10TransferVaDeleteVaUnauthorized {
	return &DeleteVaV10TransferVaDeleteVaUnauthorized{}
}

/*
DeleteVaV10TransferVaDeleteVaUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteVaV10TransferVaDeleteVaUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this delete va v10 transfer va delete va unauthorized response has a 2xx status code
func (o *DeleteVaV10TransferVaDeleteVaUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete va v10 transfer va delete va unauthorized response has a 3xx status code
func (o *DeleteVaV10TransferVaDeleteVaUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete va v10 transfer va delete va unauthorized response has a 4xx status code
func (o *DeleteVaV10TransferVaDeleteVaUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete va v10 transfer va delete va unauthorized response has a 5xx status code
func (o *DeleteVaV10TransferVaDeleteVaUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete va v10 transfer va delete va unauthorized response a status code equal to that given
func (o *DeleteVaV10TransferVaDeleteVaUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete va v10 transfer va delete va unauthorized response
func (o *DeleteVaV10TransferVaDeleteVaUnauthorized) Code() int {
	return 401
}

func (o *DeleteVaV10TransferVaDeleteVaUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /va/v1.0/transfer-va/delete-va][%d] deleteVaV10TransferVaDeleteVaUnauthorized %s", 401, payload)
}

func (o *DeleteVaV10TransferVaDeleteVaUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /va/v1.0/transfer-va/delete-va][%d] deleteVaV10TransferVaDeleteVaUnauthorized %s", 401, payload)
}

func (o *DeleteVaV10TransferVaDeleteVaUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteVaV10TransferVaDeleteVaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVaV10TransferVaDeleteVaInternalServerError creates a DeleteVaV10TransferVaDeleteVaInternalServerError with default headers values
func NewDeleteVaV10TransferVaDeleteVaInternalServerError() *DeleteVaV10TransferVaDeleteVaInternalServerError {
	return &DeleteVaV10TransferVaDeleteVaInternalServerError{}
}

/*
DeleteVaV10TransferVaDeleteVaInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteVaV10TransferVaDeleteVaInternalServerError struct {
	Payload interface{}
}

// IsSuccess returns true when this delete va v10 transfer va delete va internal server error response has a 2xx status code
func (o *DeleteVaV10TransferVaDeleteVaInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete va v10 transfer va delete va internal server error response has a 3xx status code
func (o *DeleteVaV10TransferVaDeleteVaInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete va v10 transfer va delete va internal server error response has a 4xx status code
func (o *DeleteVaV10TransferVaDeleteVaInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete va v10 transfer va delete va internal server error response has a 5xx status code
func (o *DeleteVaV10TransferVaDeleteVaInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete va v10 transfer va delete va internal server error response a status code equal to that given
func (o *DeleteVaV10TransferVaDeleteVaInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete va v10 transfer va delete va internal server error response
func (o *DeleteVaV10TransferVaDeleteVaInternalServerError) Code() int {
	return 500
}

func (o *DeleteVaV10TransferVaDeleteVaInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /va/v1.0/transfer-va/delete-va][%d] deleteVaV10TransferVaDeleteVaInternalServerError %s", 500, payload)
}

func (o *DeleteVaV10TransferVaDeleteVaInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /va/v1.0/transfer-va/delete-va][%d] deleteVaV10TransferVaDeleteVaInternalServerError %s", 500, payload)
}

func (o *DeleteVaV10TransferVaDeleteVaInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteVaV10TransferVaDeleteVaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
