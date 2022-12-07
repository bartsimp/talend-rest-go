// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bartsimp/talend-rest-go/models"
)

// GetTaskConfigurationReader is a Reader for the GetTaskConfiguration structure.
type GetTaskConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTaskConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTaskConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTaskConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTaskConfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTaskConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTaskConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTaskConfigurationOK creates a GetTaskConfigurationOK with default headers values
func NewGetTaskConfigurationOK() *GetTaskConfigurationOK {
	return &GetTaskConfigurationOK{}
}

/*
GetTaskConfigurationOK describes a response with status code 200, with default header values.

OK
*/
type GetTaskConfigurationOK struct {
	Payload *models.TaskRunConfig
}

// IsSuccess returns true when this get task configuration o k response has a 2xx status code
func (o *GetTaskConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get task configuration o k response has a 3xx status code
func (o *GetTaskConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task configuration o k response has a 4xx status code
func (o *GetTaskConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get task configuration o k response has a 5xx status code
func (o *GetTaskConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get task configuration o k response a status code equal to that given
func (o *GetTaskConfigurationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTaskConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}/run-config][%d] getTaskConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetTaskConfigurationOK) String() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}/run-config][%d] getTaskConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetTaskConfigurationOK) GetPayload() *models.TaskRunConfig {
	return o.Payload
}

func (o *GetTaskConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskRunConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskConfigurationBadRequest creates a GetTaskConfigurationBadRequest with default headers values
func NewGetTaskConfigurationBadRequest() *GetTaskConfigurationBadRequest {
	return &GetTaskConfigurationBadRequest{}
}

/*
GetTaskConfigurationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetTaskConfigurationBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get task configuration bad request response has a 2xx status code
func (o *GetTaskConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task configuration bad request response has a 3xx status code
func (o *GetTaskConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task configuration bad request response has a 4xx status code
func (o *GetTaskConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get task configuration bad request response has a 5xx status code
func (o *GetTaskConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get task configuration bad request response a status code equal to that given
func (o *GetTaskConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTaskConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}/run-config][%d] getTaskConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetTaskConfigurationBadRequest) String() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}/run-config][%d] getTaskConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetTaskConfigurationBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetTaskConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskConfigurationUnauthorized creates a GetTaskConfigurationUnauthorized with default headers values
func NewGetTaskConfigurationUnauthorized() *GetTaskConfigurationUnauthorized {
	return &GetTaskConfigurationUnauthorized{}
}

/*
GetTaskConfigurationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTaskConfigurationUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get task configuration unauthorized response has a 2xx status code
func (o *GetTaskConfigurationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task configuration unauthorized response has a 3xx status code
func (o *GetTaskConfigurationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task configuration unauthorized response has a 4xx status code
func (o *GetTaskConfigurationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get task configuration unauthorized response has a 5xx status code
func (o *GetTaskConfigurationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get task configuration unauthorized response a status code equal to that given
func (o *GetTaskConfigurationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTaskConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}/run-config][%d] getTaskConfigurationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTaskConfigurationUnauthorized) String() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}/run-config][%d] getTaskConfigurationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTaskConfigurationUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetTaskConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskConfigurationForbidden creates a GetTaskConfigurationForbidden with default headers values
func NewGetTaskConfigurationForbidden() *GetTaskConfigurationForbidden {
	return &GetTaskConfigurationForbidden{}
}

/*
GetTaskConfigurationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetTaskConfigurationForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get task configuration forbidden response has a 2xx status code
func (o *GetTaskConfigurationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task configuration forbidden response has a 3xx status code
func (o *GetTaskConfigurationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task configuration forbidden response has a 4xx status code
func (o *GetTaskConfigurationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get task configuration forbidden response has a 5xx status code
func (o *GetTaskConfigurationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get task configuration forbidden response a status code equal to that given
func (o *GetTaskConfigurationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTaskConfigurationForbidden) Error() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}/run-config][%d] getTaskConfigurationForbidden  %+v", 403, o.Payload)
}

func (o *GetTaskConfigurationForbidden) String() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}/run-config][%d] getTaskConfigurationForbidden  %+v", 403, o.Payload)
}

func (o *GetTaskConfigurationForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetTaskConfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskConfigurationNotFound creates a GetTaskConfigurationNotFound with default headers values
func NewGetTaskConfigurationNotFound() *GetTaskConfigurationNotFound {
	return &GetTaskConfigurationNotFound{}
}

/*
GetTaskConfigurationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetTaskConfigurationNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get task configuration not found response has a 2xx status code
func (o *GetTaskConfigurationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task configuration not found response has a 3xx status code
func (o *GetTaskConfigurationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task configuration not found response has a 4xx status code
func (o *GetTaskConfigurationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get task configuration not found response has a 5xx status code
func (o *GetTaskConfigurationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get task configuration not found response a status code equal to that given
func (o *GetTaskConfigurationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTaskConfigurationNotFound) Error() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}/run-config][%d] getTaskConfigurationNotFound  %+v", 404, o.Payload)
}

func (o *GetTaskConfigurationNotFound) String() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}/run-config][%d] getTaskConfigurationNotFound  %+v", 404, o.Payload)
}

func (o *GetTaskConfigurationNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetTaskConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskConfigurationInternalServerError creates a GetTaskConfigurationInternalServerError with default headers values
func NewGetTaskConfigurationInternalServerError() *GetTaskConfigurationInternalServerError {
	return &GetTaskConfigurationInternalServerError{}
}

/*
GetTaskConfigurationInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetTaskConfigurationInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get task configuration internal server error response has a 2xx status code
func (o *GetTaskConfigurationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task configuration internal server error response has a 3xx status code
func (o *GetTaskConfigurationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task configuration internal server error response has a 4xx status code
func (o *GetTaskConfigurationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get task configuration internal server error response has a 5xx status code
func (o *GetTaskConfigurationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get task configuration internal server error response a status code equal to that given
func (o *GetTaskConfigurationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTaskConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}/run-config][%d] getTaskConfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTaskConfigurationInternalServerError) String() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}/run-config][%d] getTaskConfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTaskConfigurationInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetTaskConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}