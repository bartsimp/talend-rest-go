// Code generated by go-swagger; DO NOT EDIT.

package runtime_pipeline_engines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bartsimp/talend-rest-go/models"
)

// GetPipelineEngineReader is a Reader for the GetPipelineEngine structure.
type GetPipelineEngineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPipelineEngineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPipelineEngineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPipelineEngineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPipelineEngineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPipelineEngineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPipelineEngineInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPipelineEngineOK creates a GetPipelineEngineOK with default headers values
func NewGetPipelineEngineOK() *GetPipelineEngineOK {
	return &GetPipelineEngineOK{}
}

/*
GetPipelineEngineOK describes a response with status code 200, with default header values.

OK
*/
type GetPipelineEngineOK struct {
	Payload *models.PipelineEngine
}

// IsSuccess returns true when this get pipeline engine o k response has a 2xx status code
func (o *GetPipelineEngineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get pipeline engine o k response has a 3xx status code
func (o *GetPipelineEngineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline engine o k response has a 4xx status code
func (o *GetPipelineEngineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get pipeline engine o k response has a 5xx status code
func (o *GetPipelineEngineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipeline engine o k response a status code equal to that given
func (o *GetPipelineEngineOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPipelineEngineOK) Error() string {
	return fmt.Sprintf("[GET /runtimes/pipeline-engines/{engineId}][%d] getPipelineEngineOK  %+v", 200, o.Payload)
}

func (o *GetPipelineEngineOK) String() string {
	return fmt.Sprintf("[GET /runtimes/pipeline-engines/{engineId}][%d] getPipelineEngineOK  %+v", 200, o.Payload)
}

func (o *GetPipelineEngineOK) GetPayload() *models.PipelineEngine {
	return o.Payload
}

func (o *GetPipelineEngineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PipelineEngine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPipelineEngineUnauthorized creates a GetPipelineEngineUnauthorized with default headers values
func NewGetPipelineEngineUnauthorized() *GetPipelineEngineUnauthorized {
	return &GetPipelineEngineUnauthorized{}
}

/*
GetPipelineEngineUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetPipelineEngineUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get pipeline engine unauthorized response has a 2xx status code
func (o *GetPipelineEngineUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pipeline engine unauthorized response has a 3xx status code
func (o *GetPipelineEngineUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline engine unauthorized response has a 4xx status code
func (o *GetPipelineEngineUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get pipeline engine unauthorized response has a 5xx status code
func (o *GetPipelineEngineUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipeline engine unauthorized response a status code equal to that given
func (o *GetPipelineEngineUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetPipelineEngineUnauthorized) Error() string {
	return fmt.Sprintf("[GET /runtimes/pipeline-engines/{engineId}][%d] getPipelineEngineUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPipelineEngineUnauthorized) String() string {
	return fmt.Sprintf("[GET /runtimes/pipeline-engines/{engineId}][%d] getPipelineEngineUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPipelineEngineUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetPipelineEngineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPipelineEngineForbidden creates a GetPipelineEngineForbidden with default headers values
func NewGetPipelineEngineForbidden() *GetPipelineEngineForbidden {
	return &GetPipelineEngineForbidden{}
}

/*
GetPipelineEngineForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPipelineEngineForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get pipeline engine forbidden response has a 2xx status code
func (o *GetPipelineEngineForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pipeline engine forbidden response has a 3xx status code
func (o *GetPipelineEngineForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline engine forbidden response has a 4xx status code
func (o *GetPipelineEngineForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get pipeline engine forbidden response has a 5xx status code
func (o *GetPipelineEngineForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipeline engine forbidden response a status code equal to that given
func (o *GetPipelineEngineForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetPipelineEngineForbidden) Error() string {
	return fmt.Sprintf("[GET /runtimes/pipeline-engines/{engineId}][%d] getPipelineEngineForbidden  %+v", 403, o.Payload)
}

func (o *GetPipelineEngineForbidden) String() string {
	return fmt.Sprintf("[GET /runtimes/pipeline-engines/{engineId}][%d] getPipelineEngineForbidden  %+v", 403, o.Payload)
}

func (o *GetPipelineEngineForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetPipelineEngineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPipelineEngineNotFound creates a GetPipelineEngineNotFound with default headers values
func NewGetPipelineEngineNotFound() *GetPipelineEngineNotFound {
	return &GetPipelineEngineNotFound{}
}

/*
GetPipelineEngineNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetPipelineEngineNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get pipeline engine not found response has a 2xx status code
func (o *GetPipelineEngineNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pipeline engine not found response has a 3xx status code
func (o *GetPipelineEngineNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline engine not found response has a 4xx status code
func (o *GetPipelineEngineNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get pipeline engine not found response has a 5xx status code
func (o *GetPipelineEngineNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipeline engine not found response a status code equal to that given
func (o *GetPipelineEngineNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPipelineEngineNotFound) Error() string {
	return fmt.Sprintf("[GET /runtimes/pipeline-engines/{engineId}][%d] getPipelineEngineNotFound  %+v", 404, o.Payload)
}

func (o *GetPipelineEngineNotFound) String() string {
	return fmt.Sprintf("[GET /runtimes/pipeline-engines/{engineId}][%d] getPipelineEngineNotFound  %+v", 404, o.Payload)
}

func (o *GetPipelineEngineNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetPipelineEngineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPipelineEngineInternalServerError creates a GetPipelineEngineInternalServerError with default headers values
func NewGetPipelineEngineInternalServerError() *GetPipelineEngineInternalServerError {
	return &GetPipelineEngineInternalServerError{}
}

/*
GetPipelineEngineInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetPipelineEngineInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get pipeline engine internal server error response has a 2xx status code
func (o *GetPipelineEngineInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get pipeline engine internal server error response has a 3xx status code
func (o *GetPipelineEngineInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline engine internal server error response has a 4xx status code
func (o *GetPipelineEngineInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get pipeline engine internal server error response has a 5xx status code
func (o *GetPipelineEngineInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get pipeline engine internal server error response a status code equal to that given
func (o *GetPipelineEngineInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetPipelineEngineInternalServerError) Error() string {
	return fmt.Sprintf("[GET /runtimes/pipeline-engines/{engineId}][%d] getPipelineEngineInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPipelineEngineInternalServerError) String() string {
	return fmt.Sprintf("[GET /runtimes/pipeline-engines/{engineId}][%d] getPipelineEngineInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPipelineEngineInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetPipelineEngineInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}