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

// GetTaskReader is a Reader for the GetTask structure.
type GetTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTaskUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTaskForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTaskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTaskOK creates a GetTaskOK with default headers values
func NewGetTaskOK() *GetTaskOK {
	return &GetTaskOK{}
}

/*
GetTaskOK describes a response with status code 200, with default header values.

OK
*/
type GetTaskOK struct {
	Payload *models.TaskV21
}

// IsSuccess returns true when this get task o k response has a 2xx status code
func (o *GetTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get task o k response has a 3xx status code
func (o *GetTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task o k response has a 4xx status code
func (o *GetTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get task o k response has a 5xx status code
func (o *GetTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get task o k response a status code equal to that given
func (o *GetTaskOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTaskOK) Error() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}][%d] getTaskOK  %+v", 200, o.Payload)
}

func (o *GetTaskOK) String() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}][%d] getTaskOK  %+v", 200, o.Payload)
}

func (o *GetTaskOK) GetPayload() *models.TaskV21 {
	return o.Payload
}

func (o *GetTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskV21)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskBadRequest creates a GetTaskBadRequest with default headers values
func NewGetTaskBadRequest() *GetTaskBadRequest {
	return &GetTaskBadRequest{}
}

/*
GetTaskBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetTaskBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get task bad request response has a 2xx status code
func (o *GetTaskBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task bad request response has a 3xx status code
func (o *GetTaskBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task bad request response has a 4xx status code
func (o *GetTaskBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get task bad request response has a 5xx status code
func (o *GetTaskBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get task bad request response a status code equal to that given
func (o *GetTaskBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTaskBadRequest) Error() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}][%d] getTaskBadRequest  %+v", 400, o.Payload)
}

func (o *GetTaskBadRequest) String() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}][%d] getTaskBadRequest  %+v", 400, o.Payload)
}

func (o *GetTaskBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskUnauthorized creates a GetTaskUnauthorized with default headers values
func NewGetTaskUnauthorized() *GetTaskUnauthorized {
	return &GetTaskUnauthorized{}
}

/*
GetTaskUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTaskUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get task unauthorized response has a 2xx status code
func (o *GetTaskUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task unauthorized response has a 3xx status code
func (o *GetTaskUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task unauthorized response has a 4xx status code
func (o *GetTaskUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get task unauthorized response has a 5xx status code
func (o *GetTaskUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get task unauthorized response a status code equal to that given
func (o *GetTaskUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTaskUnauthorized) Error() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}][%d] getTaskUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTaskUnauthorized) String() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}][%d] getTaskUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTaskUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetTaskUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskForbidden creates a GetTaskForbidden with default headers values
func NewGetTaskForbidden() *GetTaskForbidden {
	return &GetTaskForbidden{}
}

/*
GetTaskForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetTaskForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get task forbidden response has a 2xx status code
func (o *GetTaskForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task forbidden response has a 3xx status code
func (o *GetTaskForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task forbidden response has a 4xx status code
func (o *GetTaskForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get task forbidden response has a 5xx status code
func (o *GetTaskForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get task forbidden response a status code equal to that given
func (o *GetTaskForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTaskForbidden) Error() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}][%d] getTaskForbidden  %+v", 403, o.Payload)
}

func (o *GetTaskForbidden) String() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}][%d] getTaskForbidden  %+v", 403, o.Payload)
}

func (o *GetTaskForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetTaskForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskNotFound creates a GetTaskNotFound with default headers values
func NewGetTaskNotFound() *GetTaskNotFound {
	return &GetTaskNotFound{}
}

/*
GetTaskNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetTaskNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get task not found response has a 2xx status code
func (o *GetTaskNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task not found response has a 3xx status code
func (o *GetTaskNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task not found response has a 4xx status code
func (o *GetTaskNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get task not found response has a 5xx status code
func (o *GetTaskNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get task not found response a status code equal to that given
func (o *GetTaskNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTaskNotFound) Error() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}][%d] getTaskNotFound  %+v", 404, o.Payload)
}

func (o *GetTaskNotFound) String() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}][%d] getTaskNotFound  %+v", 404, o.Payload)
}

func (o *GetTaskNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskInternalServerError creates a GetTaskInternalServerError with default headers values
func NewGetTaskInternalServerError() *GetTaskInternalServerError {
	return &GetTaskInternalServerError{}
}

/*
GetTaskInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetTaskInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get task internal server error response has a 2xx status code
func (o *GetTaskInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get task internal server error response has a 3xx status code
func (o *GetTaskInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get task internal server error response has a 4xx status code
func (o *GetTaskInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get task internal server error response has a 5xx status code
func (o *GetTaskInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get task internal server error response a status code equal to that given
func (o *GetTaskInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTaskInternalServerError) Error() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}][%d] getTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTaskInternalServerError) String() string {
	return fmt.Sprintf("[GET /executables/tasks/{taskId}][%d] getTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTaskInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
