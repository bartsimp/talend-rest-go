// Code generated by go-swagger; DO NOT EDIT.

package schedules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bartsimp/talend-rest-go/models"
)

// UpdateTriggerInScheduleMultipleReader is a Reader for the UpdateTriggerInScheduleMultiple structure.
type UpdateTriggerInScheduleMultipleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTriggerInScheduleMultipleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTriggerInScheduleMultipleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTriggerInScheduleMultipleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateTriggerInScheduleMultipleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateTriggerInScheduleMultipleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTriggerInScheduleMultipleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateTriggerInScheduleMultipleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 428:
		result := NewUpdateTriggerInScheduleMultiplePreconditionRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateTriggerInScheduleMultipleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateTriggerInScheduleMultipleOK creates a UpdateTriggerInScheduleMultipleOK with default headers values
func NewUpdateTriggerInScheduleMultipleOK() *UpdateTriggerInScheduleMultipleOK {
	return &UpdateTriggerInScheduleMultipleOK{}
}

/*
UpdateTriggerInScheduleMultipleOK describes a response with status code 200, with default header values.

OK
*/
type UpdateTriggerInScheduleMultipleOK struct {
}

// IsSuccess returns true when this update trigger in schedule multiple o k response has a 2xx status code
func (o *UpdateTriggerInScheduleMultipleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update trigger in schedule multiple o k response has a 3xx status code
func (o *UpdateTriggerInScheduleMultipleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trigger in schedule multiple o k response has a 4xx status code
func (o *UpdateTriggerInScheduleMultipleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update trigger in schedule multiple o k response has a 5xx status code
func (o *UpdateTriggerInScheduleMultipleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update trigger in schedule multiple o k response a status code equal to that given
func (o *UpdateTriggerInScheduleMultipleOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateTriggerInScheduleMultipleOK) Error() string {
	return fmt.Sprintf("[PUT /schedules/{scheduleId}/triggers][%d] updateTriggerInScheduleMultipleOK ", 200)
}

func (o *UpdateTriggerInScheduleMultipleOK) String() string {
	return fmt.Sprintf("[PUT /schedules/{scheduleId}/triggers][%d] updateTriggerInScheduleMultipleOK ", 200)
}

func (o *UpdateTriggerInScheduleMultipleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTriggerInScheduleMultipleBadRequest creates a UpdateTriggerInScheduleMultipleBadRequest with default headers values
func NewUpdateTriggerInScheduleMultipleBadRequest() *UpdateTriggerInScheduleMultipleBadRequest {
	return &UpdateTriggerInScheduleMultipleBadRequest{}
}

/*
UpdateTriggerInScheduleMultipleBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateTriggerInScheduleMultipleBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update trigger in schedule multiple bad request response has a 2xx status code
func (o *UpdateTriggerInScheduleMultipleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trigger in schedule multiple bad request response has a 3xx status code
func (o *UpdateTriggerInScheduleMultipleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trigger in schedule multiple bad request response has a 4xx status code
func (o *UpdateTriggerInScheduleMultipleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trigger in schedule multiple bad request response has a 5xx status code
func (o *UpdateTriggerInScheduleMultipleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update trigger in schedule multiple bad request response a status code equal to that given
func (o *UpdateTriggerInScheduleMultipleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateTriggerInScheduleMultipleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /schedules/{scheduleId}/triggers][%d] updateTriggerInScheduleMultipleBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTriggerInScheduleMultipleBadRequest) String() string {
	return fmt.Sprintf("[PUT /schedules/{scheduleId}/triggers][%d] updateTriggerInScheduleMultipleBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTriggerInScheduleMultipleBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateTriggerInScheduleMultipleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTriggerInScheduleMultipleUnauthorized creates a UpdateTriggerInScheduleMultipleUnauthorized with default headers values
func NewUpdateTriggerInScheduleMultipleUnauthorized() *UpdateTriggerInScheduleMultipleUnauthorized {
	return &UpdateTriggerInScheduleMultipleUnauthorized{}
}

/*
UpdateTriggerInScheduleMultipleUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateTriggerInScheduleMultipleUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update trigger in schedule multiple unauthorized response has a 2xx status code
func (o *UpdateTriggerInScheduleMultipleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trigger in schedule multiple unauthorized response has a 3xx status code
func (o *UpdateTriggerInScheduleMultipleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trigger in schedule multiple unauthorized response has a 4xx status code
func (o *UpdateTriggerInScheduleMultipleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trigger in schedule multiple unauthorized response has a 5xx status code
func (o *UpdateTriggerInScheduleMultipleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update trigger in schedule multiple unauthorized response a status code equal to that given
func (o *UpdateTriggerInScheduleMultipleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateTriggerInScheduleMultipleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /schedules/{scheduleId}/triggers][%d] updateTriggerInScheduleMultipleUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateTriggerInScheduleMultipleUnauthorized) String() string {
	return fmt.Sprintf("[PUT /schedules/{scheduleId}/triggers][%d] updateTriggerInScheduleMultipleUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateTriggerInScheduleMultipleUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateTriggerInScheduleMultipleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTriggerInScheduleMultipleForbidden creates a UpdateTriggerInScheduleMultipleForbidden with default headers values
func NewUpdateTriggerInScheduleMultipleForbidden() *UpdateTriggerInScheduleMultipleForbidden {
	return &UpdateTriggerInScheduleMultipleForbidden{}
}

/*
UpdateTriggerInScheduleMultipleForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateTriggerInScheduleMultipleForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update trigger in schedule multiple forbidden response has a 2xx status code
func (o *UpdateTriggerInScheduleMultipleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trigger in schedule multiple forbidden response has a 3xx status code
func (o *UpdateTriggerInScheduleMultipleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trigger in schedule multiple forbidden response has a 4xx status code
func (o *UpdateTriggerInScheduleMultipleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trigger in schedule multiple forbidden response has a 5xx status code
func (o *UpdateTriggerInScheduleMultipleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update trigger in schedule multiple forbidden response a status code equal to that given
func (o *UpdateTriggerInScheduleMultipleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateTriggerInScheduleMultipleForbidden) Error() string {
	return fmt.Sprintf("[PUT /schedules/{scheduleId}/triggers][%d] updateTriggerInScheduleMultipleForbidden  %+v", 403, o.Payload)
}

func (o *UpdateTriggerInScheduleMultipleForbidden) String() string {
	return fmt.Sprintf("[PUT /schedules/{scheduleId}/triggers][%d] updateTriggerInScheduleMultipleForbidden  %+v", 403, o.Payload)
}

func (o *UpdateTriggerInScheduleMultipleForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateTriggerInScheduleMultipleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTriggerInScheduleMultipleNotFound creates a UpdateTriggerInScheduleMultipleNotFound with default headers values
func NewUpdateTriggerInScheduleMultipleNotFound() *UpdateTriggerInScheduleMultipleNotFound {
	return &UpdateTriggerInScheduleMultipleNotFound{}
}

/*
UpdateTriggerInScheduleMultipleNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateTriggerInScheduleMultipleNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update trigger in schedule multiple not found response has a 2xx status code
func (o *UpdateTriggerInScheduleMultipleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trigger in schedule multiple not found response has a 3xx status code
func (o *UpdateTriggerInScheduleMultipleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trigger in schedule multiple not found response has a 4xx status code
func (o *UpdateTriggerInScheduleMultipleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trigger in schedule multiple not found response has a 5xx status code
func (o *UpdateTriggerInScheduleMultipleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update trigger in schedule multiple not found response a status code equal to that given
func (o *UpdateTriggerInScheduleMultipleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateTriggerInScheduleMultipleNotFound) Error() string {
	return fmt.Sprintf("[PUT /schedules/{scheduleId}/triggers][%d] updateTriggerInScheduleMultipleNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTriggerInScheduleMultipleNotFound) String() string {
	return fmt.Sprintf("[PUT /schedules/{scheduleId}/triggers][%d] updateTriggerInScheduleMultipleNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTriggerInScheduleMultipleNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateTriggerInScheduleMultipleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTriggerInScheduleMultipleConflict creates a UpdateTriggerInScheduleMultipleConflict with default headers values
func NewUpdateTriggerInScheduleMultipleConflict() *UpdateTriggerInScheduleMultipleConflict {
	return &UpdateTriggerInScheduleMultipleConflict{}
}

/*
UpdateTriggerInScheduleMultipleConflict describes a response with status code 409, with default header values.

Conflict
*/
type UpdateTriggerInScheduleMultipleConflict struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update trigger in schedule multiple conflict response has a 2xx status code
func (o *UpdateTriggerInScheduleMultipleConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trigger in schedule multiple conflict response has a 3xx status code
func (o *UpdateTriggerInScheduleMultipleConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trigger in schedule multiple conflict response has a 4xx status code
func (o *UpdateTriggerInScheduleMultipleConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trigger in schedule multiple conflict response has a 5xx status code
func (o *UpdateTriggerInScheduleMultipleConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update trigger in schedule multiple conflict response a status code equal to that given
func (o *UpdateTriggerInScheduleMultipleConflict) IsCode(code int) bool {
	return code == 409
}

func (o *UpdateTriggerInScheduleMultipleConflict) Error() string {
	return fmt.Sprintf("[PUT /schedules/{scheduleId}/triggers][%d] updateTriggerInScheduleMultipleConflict  %+v", 409, o.Payload)
}

func (o *UpdateTriggerInScheduleMultipleConflict) String() string {
	return fmt.Sprintf("[PUT /schedules/{scheduleId}/triggers][%d] updateTriggerInScheduleMultipleConflict  %+v", 409, o.Payload)
}

func (o *UpdateTriggerInScheduleMultipleConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateTriggerInScheduleMultipleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTriggerInScheduleMultiplePreconditionRequired creates a UpdateTriggerInScheduleMultiplePreconditionRequired with default headers values
func NewUpdateTriggerInScheduleMultiplePreconditionRequired() *UpdateTriggerInScheduleMultiplePreconditionRequired {
	return &UpdateTriggerInScheduleMultiplePreconditionRequired{}
}

/*
UpdateTriggerInScheduleMultiplePreconditionRequired describes a response with status code 428, with default header values.

Precondition Required
*/
type UpdateTriggerInScheduleMultiplePreconditionRequired struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update trigger in schedule multiple precondition required response has a 2xx status code
func (o *UpdateTriggerInScheduleMultiplePreconditionRequired) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trigger in schedule multiple precondition required response has a 3xx status code
func (o *UpdateTriggerInScheduleMultiplePreconditionRequired) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trigger in schedule multiple precondition required response has a 4xx status code
func (o *UpdateTriggerInScheduleMultiplePreconditionRequired) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trigger in schedule multiple precondition required response has a 5xx status code
func (o *UpdateTriggerInScheduleMultiplePreconditionRequired) IsServerError() bool {
	return false
}

// IsCode returns true when this update trigger in schedule multiple precondition required response a status code equal to that given
func (o *UpdateTriggerInScheduleMultiplePreconditionRequired) IsCode(code int) bool {
	return code == 428
}

func (o *UpdateTriggerInScheduleMultiplePreconditionRequired) Error() string {
	return fmt.Sprintf("[PUT /schedules/{scheduleId}/triggers][%d] updateTriggerInScheduleMultiplePreconditionRequired  %+v", 428, o.Payload)
}

func (o *UpdateTriggerInScheduleMultiplePreconditionRequired) String() string {
	return fmt.Sprintf("[PUT /schedules/{scheduleId}/triggers][%d] updateTriggerInScheduleMultiplePreconditionRequired  %+v", 428, o.Payload)
}

func (o *UpdateTriggerInScheduleMultiplePreconditionRequired) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateTriggerInScheduleMultiplePreconditionRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTriggerInScheduleMultipleInternalServerError creates a UpdateTriggerInScheduleMultipleInternalServerError with default headers values
func NewUpdateTriggerInScheduleMultipleInternalServerError() *UpdateTriggerInScheduleMultipleInternalServerError {
	return &UpdateTriggerInScheduleMultipleInternalServerError{}
}

/*
UpdateTriggerInScheduleMultipleInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateTriggerInScheduleMultipleInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update trigger in schedule multiple internal server error response has a 2xx status code
func (o *UpdateTriggerInScheduleMultipleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trigger in schedule multiple internal server error response has a 3xx status code
func (o *UpdateTriggerInScheduleMultipleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trigger in schedule multiple internal server error response has a 4xx status code
func (o *UpdateTriggerInScheduleMultipleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update trigger in schedule multiple internal server error response has a 5xx status code
func (o *UpdateTriggerInScheduleMultipleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update trigger in schedule multiple internal server error response a status code equal to that given
func (o *UpdateTriggerInScheduleMultipleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateTriggerInScheduleMultipleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /schedules/{scheduleId}/triggers][%d] updateTriggerInScheduleMultipleInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTriggerInScheduleMultipleInternalServerError) String() string {
	return fmt.Sprintf("[PUT /schedules/{scheduleId}/triggers][%d] updateTriggerInScheduleMultipleInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTriggerInScheduleMultipleInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateTriggerInScheduleMultipleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}