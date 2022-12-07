// Code generated by go-swagger; DO NOT EDIT.

package runtime_engines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bartsimp/talend-rest-go/models"
)

// UnpairRemoteEngineReader is a Reader for the UnpairRemoteEngine structure.
type UnpairRemoteEngineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnpairRemoteEngineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnpairRemoteEngineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUnpairRemoteEngineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUnpairRemoteEngineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUnpairRemoteEngineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUnpairRemoteEngineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUnpairRemoteEngineInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUnpairRemoteEngineOK creates a UnpairRemoteEngineOK with default headers values
func NewUnpairRemoteEngineOK() *UnpairRemoteEngineOK {
	return &UnpairRemoteEngineOK{}
}

/*
UnpairRemoteEngineOK describes a response with status code 200, with default header values.

OK
*/
type UnpairRemoteEngineOK struct {
	Payload string
}

// IsSuccess returns true when this unpair remote engine o k response has a 2xx status code
func (o *UnpairRemoteEngineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unpair remote engine o k response has a 3xx status code
func (o *UnpairRemoteEngineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unpair remote engine o k response has a 4xx status code
func (o *UnpairRemoteEngineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unpair remote engine o k response has a 5xx status code
func (o *UnpairRemoteEngineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unpair remote engine o k response a status code equal to that given
func (o *UnpairRemoteEngineOK) IsCode(code int) bool {
	return code == 200
}

func (o *UnpairRemoteEngineOK) Error() string {
	return fmt.Sprintf("[DELETE /runtimes/remote-engines/{id}/pairing][%d] unpairRemoteEngineOK  %+v", 200, o.Payload)
}

func (o *UnpairRemoteEngineOK) String() string {
	return fmt.Sprintf("[DELETE /runtimes/remote-engines/{id}/pairing][%d] unpairRemoteEngineOK  %+v", 200, o.Payload)
}

func (o *UnpairRemoteEngineOK) GetPayload() string {
	return o.Payload
}

func (o *UnpairRemoteEngineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnpairRemoteEngineBadRequest creates a UnpairRemoteEngineBadRequest with default headers values
func NewUnpairRemoteEngineBadRequest() *UnpairRemoteEngineBadRequest {
	return &UnpairRemoteEngineBadRequest{}
}

/*
UnpairRemoteEngineBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UnpairRemoteEngineBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this unpair remote engine bad request response has a 2xx status code
func (o *UnpairRemoteEngineBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unpair remote engine bad request response has a 3xx status code
func (o *UnpairRemoteEngineBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unpair remote engine bad request response has a 4xx status code
func (o *UnpairRemoteEngineBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this unpair remote engine bad request response has a 5xx status code
func (o *UnpairRemoteEngineBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this unpair remote engine bad request response a status code equal to that given
func (o *UnpairRemoteEngineBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UnpairRemoteEngineBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /runtimes/remote-engines/{id}/pairing][%d] unpairRemoteEngineBadRequest  %+v", 400, o.Payload)
}

func (o *UnpairRemoteEngineBadRequest) String() string {
	return fmt.Sprintf("[DELETE /runtimes/remote-engines/{id}/pairing][%d] unpairRemoteEngineBadRequest  %+v", 400, o.Payload)
}

func (o *UnpairRemoteEngineBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnpairRemoteEngineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnpairRemoteEngineUnauthorized creates a UnpairRemoteEngineUnauthorized with default headers values
func NewUnpairRemoteEngineUnauthorized() *UnpairRemoteEngineUnauthorized {
	return &UnpairRemoteEngineUnauthorized{}
}

/*
UnpairRemoteEngineUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UnpairRemoteEngineUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this unpair remote engine unauthorized response has a 2xx status code
func (o *UnpairRemoteEngineUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unpair remote engine unauthorized response has a 3xx status code
func (o *UnpairRemoteEngineUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unpair remote engine unauthorized response has a 4xx status code
func (o *UnpairRemoteEngineUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this unpair remote engine unauthorized response has a 5xx status code
func (o *UnpairRemoteEngineUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this unpair remote engine unauthorized response a status code equal to that given
func (o *UnpairRemoteEngineUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UnpairRemoteEngineUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /runtimes/remote-engines/{id}/pairing][%d] unpairRemoteEngineUnauthorized  %+v", 401, o.Payload)
}

func (o *UnpairRemoteEngineUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /runtimes/remote-engines/{id}/pairing][%d] unpairRemoteEngineUnauthorized  %+v", 401, o.Payload)
}

func (o *UnpairRemoteEngineUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnpairRemoteEngineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnpairRemoteEngineForbidden creates a UnpairRemoteEngineForbidden with default headers values
func NewUnpairRemoteEngineForbidden() *UnpairRemoteEngineForbidden {
	return &UnpairRemoteEngineForbidden{}
}

/*
UnpairRemoteEngineForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UnpairRemoteEngineForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this unpair remote engine forbidden response has a 2xx status code
func (o *UnpairRemoteEngineForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unpair remote engine forbidden response has a 3xx status code
func (o *UnpairRemoteEngineForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unpair remote engine forbidden response has a 4xx status code
func (o *UnpairRemoteEngineForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this unpair remote engine forbidden response has a 5xx status code
func (o *UnpairRemoteEngineForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this unpair remote engine forbidden response a status code equal to that given
func (o *UnpairRemoteEngineForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UnpairRemoteEngineForbidden) Error() string {
	return fmt.Sprintf("[DELETE /runtimes/remote-engines/{id}/pairing][%d] unpairRemoteEngineForbidden  %+v", 403, o.Payload)
}

func (o *UnpairRemoteEngineForbidden) String() string {
	return fmt.Sprintf("[DELETE /runtimes/remote-engines/{id}/pairing][%d] unpairRemoteEngineForbidden  %+v", 403, o.Payload)
}

func (o *UnpairRemoteEngineForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnpairRemoteEngineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnpairRemoteEngineNotFound creates a UnpairRemoteEngineNotFound with default headers values
func NewUnpairRemoteEngineNotFound() *UnpairRemoteEngineNotFound {
	return &UnpairRemoteEngineNotFound{}
}

/*
UnpairRemoteEngineNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UnpairRemoteEngineNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this unpair remote engine not found response has a 2xx status code
func (o *UnpairRemoteEngineNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unpair remote engine not found response has a 3xx status code
func (o *UnpairRemoteEngineNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unpair remote engine not found response has a 4xx status code
func (o *UnpairRemoteEngineNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this unpair remote engine not found response has a 5xx status code
func (o *UnpairRemoteEngineNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this unpair remote engine not found response a status code equal to that given
func (o *UnpairRemoteEngineNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UnpairRemoteEngineNotFound) Error() string {
	return fmt.Sprintf("[DELETE /runtimes/remote-engines/{id}/pairing][%d] unpairRemoteEngineNotFound  %+v", 404, o.Payload)
}

func (o *UnpairRemoteEngineNotFound) String() string {
	return fmt.Sprintf("[DELETE /runtimes/remote-engines/{id}/pairing][%d] unpairRemoteEngineNotFound  %+v", 404, o.Payload)
}

func (o *UnpairRemoteEngineNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnpairRemoteEngineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnpairRemoteEngineInternalServerError creates a UnpairRemoteEngineInternalServerError with default headers values
func NewUnpairRemoteEngineInternalServerError() *UnpairRemoteEngineInternalServerError {
	return &UnpairRemoteEngineInternalServerError{}
}

/*
UnpairRemoteEngineInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UnpairRemoteEngineInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this unpair remote engine internal server error response has a 2xx status code
func (o *UnpairRemoteEngineInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unpair remote engine internal server error response has a 3xx status code
func (o *UnpairRemoteEngineInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unpair remote engine internal server error response has a 4xx status code
func (o *UnpairRemoteEngineInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this unpair remote engine internal server error response has a 5xx status code
func (o *UnpairRemoteEngineInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this unpair remote engine internal server error response a status code equal to that given
func (o *UnpairRemoteEngineInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UnpairRemoteEngineInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /runtimes/remote-engines/{id}/pairing][%d] unpairRemoteEngineInternalServerError  %+v", 500, o.Payload)
}

func (o *UnpairRemoteEngineInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /runtimes/remote-engines/{id}/pairing][%d] unpairRemoteEngineInternalServerError  %+v", 500, o.Payload)
}

func (o *UnpairRemoteEngineInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnpairRemoteEngineInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}