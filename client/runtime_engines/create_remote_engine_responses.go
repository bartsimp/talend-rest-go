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

// CreateRemoteEngineReader is a Reader for the CreateRemoteEngine structure.
type CreateRemoteEngineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRemoteEngineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRemoteEngineCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRemoteEngineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateRemoteEngineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRemoteEngineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRemoteEngineInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRemoteEngineCreated creates a CreateRemoteEngineCreated with default headers values
func NewCreateRemoteEngineCreated() *CreateRemoteEngineCreated {
	return &CreateRemoteEngineCreated{}
}

/*
CreateRemoteEngineCreated describes a response with status code 201, with default header values.

Created
*/
type CreateRemoteEngineCreated struct {
	Payload *models.Engine
}

// IsSuccess returns true when this create remote engine created response has a 2xx status code
func (o *CreateRemoteEngineCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create remote engine created response has a 3xx status code
func (o *CreateRemoteEngineCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create remote engine created response has a 4xx status code
func (o *CreateRemoteEngineCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create remote engine created response has a 5xx status code
func (o *CreateRemoteEngineCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create remote engine created response a status code equal to that given
func (o *CreateRemoteEngineCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateRemoteEngineCreated) Error() string {
	return fmt.Sprintf("[POST /runtimes/remote-engines][%d] createRemoteEngineCreated  %+v", 201, o.Payload)
}

func (o *CreateRemoteEngineCreated) String() string {
	return fmt.Sprintf("[POST /runtimes/remote-engines][%d] createRemoteEngineCreated  %+v", 201, o.Payload)
}

func (o *CreateRemoteEngineCreated) GetPayload() *models.Engine {
	return o.Payload
}

func (o *CreateRemoteEngineCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Engine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRemoteEngineBadRequest creates a CreateRemoteEngineBadRequest with default headers values
func NewCreateRemoteEngineBadRequest() *CreateRemoteEngineBadRequest {
	return &CreateRemoteEngineBadRequest{}
}

/*
CreateRemoteEngineBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateRemoteEngineBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create remote engine bad request response has a 2xx status code
func (o *CreateRemoteEngineBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create remote engine bad request response has a 3xx status code
func (o *CreateRemoteEngineBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create remote engine bad request response has a 4xx status code
func (o *CreateRemoteEngineBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create remote engine bad request response has a 5xx status code
func (o *CreateRemoteEngineBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create remote engine bad request response a status code equal to that given
func (o *CreateRemoteEngineBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateRemoteEngineBadRequest) Error() string {
	return fmt.Sprintf("[POST /runtimes/remote-engines][%d] createRemoteEngineBadRequest  %+v", 400, o.Payload)
}

func (o *CreateRemoteEngineBadRequest) String() string {
	return fmt.Sprintf("[POST /runtimes/remote-engines][%d] createRemoteEngineBadRequest  %+v", 400, o.Payload)
}

func (o *CreateRemoteEngineBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateRemoteEngineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRemoteEngineUnauthorized creates a CreateRemoteEngineUnauthorized with default headers values
func NewCreateRemoteEngineUnauthorized() *CreateRemoteEngineUnauthorized {
	return &CreateRemoteEngineUnauthorized{}
}

/*
CreateRemoteEngineUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateRemoteEngineUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create remote engine unauthorized response has a 2xx status code
func (o *CreateRemoteEngineUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create remote engine unauthorized response has a 3xx status code
func (o *CreateRemoteEngineUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create remote engine unauthorized response has a 4xx status code
func (o *CreateRemoteEngineUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create remote engine unauthorized response has a 5xx status code
func (o *CreateRemoteEngineUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create remote engine unauthorized response a status code equal to that given
func (o *CreateRemoteEngineUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateRemoteEngineUnauthorized) Error() string {
	return fmt.Sprintf("[POST /runtimes/remote-engines][%d] createRemoteEngineUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateRemoteEngineUnauthorized) String() string {
	return fmt.Sprintf("[POST /runtimes/remote-engines][%d] createRemoteEngineUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateRemoteEngineUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateRemoteEngineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRemoteEngineForbidden creates a CreateRemoteEngineForbidden with default headers values
func NewCreateRemoteEngineForbidden() *CreateRemoteEngineForbidden {
	return &CreateRemoteEngineForbidden{}
}

/*
CreateRemoteEngineForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateRemoteEngineForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create remote engine forbidden response has a 2xx status code
func (o *CreateRemoteEngineForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create remote engine forbidden response has a 3xx status code
func (o *CreateRemoteEngineForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create remote engine forbidden response has a 4xx status code
func (o *CreateRemoteEngineForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create remote engine forbidden response has a 5xx status code
func (o *CreateRemoteEngineForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create remote engine forbidden response a status code equal to that given
func (o *CreateRemoteEngineForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateRemoteEngineForbidden) Error() string {
	return fmt.Sprintf("[POST /runtimes/remote-engines][%d] createRemoteEngineForbidden  %+v", 403, o.Payload)
}

func (o *CreateRemoteEngineForbidden) String() string {
	return fmt.Sprintf("[POST /runtimes/remote-engines][%d] createRemoteEngineForbidden  %+v", 403, o.Payload)
}

func (o *CreateRemoteEngineForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateRemoteEngineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRemoteEngineInternalServerError creates a CreateRemoteEngineInternalServerError with default headers values
func NewCreateRemoteEngineInternalServerError() *CreateRemoteEngineInternalServerError {
	return &CreateRemoteEngineInternalServerError{}
}

/*
CreateRemoteEngineInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateRemoteEngineInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create remote engine internal server error response has a 2xx status code
func (o *CreateRemoteEngineInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create remote engine internal server error response has a 3xx status code
func (o *CreateRemoteEngineInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create remote engine internal server error response has a 4xx status code
func (o *CreateRemoteEngineInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create remote engine internal server error response has a 5xx status code
func (o *CreateRemoteEngineInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create remote engine internal server error response a status code equal to that given
func (o *CreateRemoteEngineInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateRemoteEngineInternalServerError) Error() string {
	return fmt.Sprintf("[POST /runtimes/remote-engines][%d] createRemoteEngineInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRemoteEngineInternalServerError) String() string {
	return fmt.Sprintf("[POST /runtimes/remote-engines][%d] createRemoteEngineInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRemoteEngineInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateRemoteEngineInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
