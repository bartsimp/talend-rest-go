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

// DeleteTaskReader is a Reader for the DeleteTask structure.
type DeleteTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteTaskNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteTaskUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTaskForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteTaskConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTaskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTaskNoContent creates a DeleteTaskNoContent with default headers values
func NewDeleteTaskNoContent() *DeleteTaskNoContent {
	return &DeleteTaskNoContent{}
}

/*
DeleteTaskNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteTaskNoContent struct {
}

// IsSuccess returns true when this delete task no content response has a 2xx status code
func (o *DeleteTaskNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete task no content response has a 3xx status code
func (o *DeleteTaskNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete task no content response has a 4xx status code
func (o *DeleteTaskNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete task no content response has a 5xx status code
func (o *DeleteTaskNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete task no content response a status code equal to that given
func (o *DeleteTaskNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteTaskNoContent) Error() string {
	return fmt.Sprintf("[DELETE /executables/tasks/{taskId}][%d] deleteTaskNoContent ", 204)
}

func (o *DeleteTaskNoContent) String() string {
	return fmt.Sprintf("[DELETE /executables/tasks/{taskId}][%d] deleteTaskNoContent ", 204)
}

func (o *DeleteTaskNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTaskBadRequest creates a DeleteTaskBadRequest with default headers values
func NewDeleteTaskBadRequest() *DeleteTaskBadRequest {
	return &DeleteTaskBadRequest{}
}

/*
DeleteTaskBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteTaskBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete task bad request response has a 2xx status code
func (o *DeleteTaskBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete task bad request response has a 3xx status code
func (o *DeleteTaskBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete task bad request response has a 4xx status code
func (o *DeleteTaskBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete task bad request response has a 5xx status code
func (o *DeleteTaskBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete task bad request response a status code equal to that given
func (o *DeleteTaskBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteTaskBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /executables/tasks/{taskId}][%d] deleteTaskBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTaskBadRequest) String() string {
	return fmt.Sprintf("[DELETE /executables/tasks/{taskId}][%d] deleteTaskBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTaskBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTaskUnauthorized creates a DeleteTaskUnauthorized with default headers values
func NewDeleteTaskUnauthorized() *DeleteTaskUnauthorized {
	return &DeleteTaskUnauthorized{}
}

/*
DeleteTaskUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteTaskUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete task unauthorized response has a 2xx status code
func (o *DeleteTaskUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete task unauthorized response has a 3xx status code
func (o *DeleteTaskUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete task unauthorized response has a 4xx status code
func (o *DeleteTaskUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete task unauthorized response has a 5xx status code
func (o *DeleteTaskUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete task unauthorized response a status code equal to that given
func (o *DeleteTaskUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteTaskUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /executables/tasks/{taskId}][%d] deleteTaskUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTaskUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /executables/tasks/{taskId}][%d] deleteTaskUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTaskUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteTaskUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTaskForbidden creates a DeleteTaskForbidden with default headers values
func NewDeleteTaskForbidden() *DeleteTaskForbidden {
	return &DeleteTaskForbidden{}
}

/*
DeleteTaskForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteTaskForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete task forbidden response has a 2xx status code
func (o *DeleteTaskForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete task forbidden response has a 3xx status code
func (o *DeleteTaskForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete task forbidden response has a 4xx status code
func (o *DeleteTaskForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete task forbidden response has a 5xx status code
func (o *DeleteTaskForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete task forbidden response a status code equal to that given
func (o *DeleteTaskForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteTaskForbidden) Error() string {
	return fmt.Sprintf("[DELETE /executables/tasks/{taskId}][%d] deleteTaskForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTaskForbidden) String() string {
	return fmt.Sprintf("[DELETE /executables/tasks/{taskId}][%d] deleteTaskForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTaskForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteTaskForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTaskNotFound creates a DeleteTaskNotFound with default headers values
func NewDeleteTaskNotFound() *DeleteTaskNotFound {
	return &DeleteTaskNotFound{}
}

/*
DeleteTaskNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteTaskNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete task not found response has a 2xx status code
func (o *DeleteTaskNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete task not found response has a 3xx status code
func (o *DeleteTaskNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete task not found response has a 4xx status code
func (o *DeleteTaskNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete task not found response has a 5xx status code
func (o *DeleteTaskNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete task not found response a status code equal to that given
func (o *DeleteTaskNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteTaskNotFound) Error() string {
	return fmt.Sprintf("[DELETE /executables/tasks/{taskId}][%d] deleteTaskNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTaskNotFound) String() string {
	return fmt.Sprintf("[DELETE /executables/tasks/{taskId}][%d] deleteTaskNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTaskNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTaskConflict creates a DeleteTaskConflict with default headers values
func NewDeleteTaskConflict() *DeleteTaskConflict {
	return &DeleteTaskConflict{}
}

/*
DeleteTaskConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteTaskConflict struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete task conflict response has a 2xx status code
func (o *DeleteTaskConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete task conflict response has a 3xx status code
func (o *DeleteTaskConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete task conflict response has a 4xx status code
func (o *DeleteTaskConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete task conflict response has a 5xx status code
func (o *DeleteTaskConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete task conflict response a status code equal to that given
func (o *DeleteTaskConflict) IsCode(code int) bool {
	return code == 409
}

func (o *DeleteTaskConflict) Error() string {
	return fmt.Sprintf("[DELETE /executables/tasks/{taskId}][%d] deleteTaskConflict  %+v", 409, o.Payload)
}

func (o *DeleteTaskConflict) String() string {
	return fmt.Sprintf("[DELETE /executables/tasks/{taskId}][%d] deleteTaskConflict  %+v", 409, o.Payload)
}

func (o *DeleteTaskConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteTaskConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTaskInternalServerError creates a DeleteTaskInternalServerError with default headers values
func NewDeleteTaskInternalServerError() *DeleteTaskInternalServerError {
	return &DeleteTaskInternalServerError{}
}

/*
DeleteTaskInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteTaskInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete task internal server error response has a 2xx status code
func (o *DeleteTaskInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete task internal server error response has a 3xx status code
func (o *DeleteTaskInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete task internal server error response has a 4xx status code
func (o *DeleteTaskInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete task internal server error response has a 5xx status code
func (o *DeleteTaskInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete task internal server error response a status code equal to that given
func (o *DeleteTaskInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteTaskInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /executables/tasks/{taskId}][%d] deleteTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTaskInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /executables/tasks/{taskId}][%d] deleteTaskInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTaskInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
