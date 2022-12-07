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

// DeleteScheduleMultipleTriggersReader is a Reader for the DeleteScheduleMultipleTriggers structure.
type DeleteScheduleMultipleTriggersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteScheduleMultipleTriggersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteScheduleMultipleTriggersNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteScheduleMultipleTriggersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteScheduleMultipleTriggersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteScheduleMultipleTriggersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteScheduleMultipleTriggersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteScheduleMultipleTriggersConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteScheduleMultipleTriggersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteScheduleMultipleTriggersNoContent creates a DeleteScheduleMultipleTriggersNoContent with default headers values
func NewDeleteScheduleMultipleTriggersNoContent() *DeleteScheduleMultipleTriggersNoContent {
	return &DeleteScheduleMultipleTriggersNoContent{}
}

/*
DeleteScheduleMultipleTriggersNoContent describes a response with status code 204, with default header values.

No content
*/
type DeleteScheduleMultipleTriggersNoContent struct {
}

// IsSuccess returns true when this delete schedule multiple triggers no content response has a 2xx status code
func (o *DeleteScheduleMultipleTriggersNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete schedule multiple triggers no content response has a 3xx status code
func (o *DeleteScheduleMultipleTriggersNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete schedule multiple triggers no content response has a 4xx status code
func (o *DeleteScheduleMultipleTriggersNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete schedule multiple triggers no content response has a 5xx status code
func (o *DeleteScheduleMultipleTriggersNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete schedule multiple triggers no content response a status code equal to that given
func (o *DeleteScheduleMultipleTriggersNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteScheduleMultipleTriggersNoContent) Error() string {
	return fmt.Sprintf("[DELETE /schedules/{scheduleId}][%d] deleteScheduleMultipleTriggersNoContent ", 204)
}

func (o *DeleteScheduleMultipleTriggersNoContent) String() string {
	return fmt.Sprintf("[DELETE /schedules/{scheduleId}][%d] deleteScheduleMultipleTriggersNoContent ", 204)
}

func (o *DeleteScheduleMultipleTriggersNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteScheduleMultipleTriggersBadRequest creates a DeleteScheduleMultipleTriggersBadRequest with default headers values
func NewDeleteScheduleMultipleTriggersBadRequest() *DeleteScheduleMultipleTriggersBadRequest {
	return &DeleteScheduleMultipleTriggersBadRequest{}
}

/*
DeleteScheduleMultipleTriggersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteScheduleMultipleTriggersBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete schedule multiple triggers bad request response has a 2xx status code
func (o *DeleteScheduleMultipleTriggersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete schedule multiple triggers bad request response has a 3xx status code
func (o *DeleteScheduleMultipleTriggersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete schedule multiple triggers bad request response has a 4xx status code
func (o *DeleteScheduleMultipleTriggersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete schedule multiple triggers bad request response has a 5xx status code
func (o *DeleteScheduleMultipleTriggersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete schedule multiple triggers bad request response a status code equal to that given
func (o *DeleteScheduleMultipleTriggersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteScheduleMultipleTriggersBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /schedules/{scheduleId}][%d] deleteScheduleMultipleTriggersBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteScheduleMultipleTriggersBadRequest) String() string {
	return fmt.Sprintf("[DELETE /schedules/{scheduleId}][%d] deleteScheduleMultipleTriggersBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteScheduleMultipleTriggersBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteScheduleMultipleTriggersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScheduleMultipleTriggersUnauthorized creates a DeleteScheduleMultipleTriggersUnauthorized with default headers values
func NewDeleteScheduleMultipleTriggersUnauthorized() *DeleteScheduleMultipleTriggersUnauthorized {
	return &DeleteScheduleMultipleTriggersUnauthorized{}
}

/*
DeleteScheduleMultipleTriggersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteScheduleMultipleTriggersUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete schedule multiple triggers unauthorized response has a 2xx status code
func (o *DeleteScheduleMultipleTriggersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete schedule multiple triggers unauthorized response has a 3xx status code
func (o *DeleteScheduleMultipleTriggersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete schedule multiple triggers unauthorized response has a 4xx status code
func (o *DeleteScheduleMultipleTriggersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete schedule multiple triggers unauthorized response has a 5xx status code
func (o *DeleteScheduleMultipleTriggersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete schedule multiple triggers unauthorized response a status code equal to that given
func (o *DeleteScheduleMultipleTriggersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteScheduleMultipleTriggersUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /schedules/{scheduleId}][%d] deleteScheduleMultipleTriggersUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteScheduleMultipleTriggersUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /schedules/{scheduleId}][%d] deleteScheduleMultipleTriggersUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteScheduleMultipleTriggersUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteScheduleMultipleTriggersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScheduleMultipleTriggersForbidden creates a DeleteScheduleMultipleTriggersForbidden with default headers values
func NewDeleteScheduleMultipleTriggersForbidden() *DeleteScheduleMultipleTriggersForbidden {
	return &DeleteScheduleMultipleTriggersForbidden{}
}

/*
DeleteScheduleMultipleTriggersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteScheduleMultipleTriggersForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete schedule multiple triggers forbidden response has a 2xx status code
func (o *DeleteScheduleMultipleTriggersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete schedule multiple triggers forbidden response has a 3xx status code
func (o *DeleteScheduleMultipleTriggersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete schedule multiple triggers forbidden response has a 4xx status code
func (o *DeleteScheduleMultipleTriggersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete schedule multiple triggers forbidden response has a 5xx status code
func (o *DeleteScheduleMultipleTriggersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete schedule multiple triggers forbidden response a status code equal to that given
func (o *DeleteScheduleMultipleTriggersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteScheduleMultipleTriggersForbidden) Error() string {
	return fmt.Sprintf("[DELETE /schedules/{scheduleId}][%d] deleteScheduleMultipleTriggersForbidden  %+v", 403, o.Payload)
}

func (o *DeleteScheduleMultipleTriggersForbidden) String() string {
	return fmt.Sprintf("[DELETE /schedules/{scheduleId}][%d] deleteScheduleMultipleTriggersForbidden  %+v", 403, o.Payload)
}

func (o *DeleteScheduleMultipleTriggersForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteScheduleMultipleTriggersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScheduleMultipleTriggersNotFound creates a DeleteScheduleMultipleTriggersNotFound with default headers values
func NewDeleteScheduleMultipleTriggersNotFound() *DeleteScheduleMultipleTriggersNotFound {
	return &DeleteScheduleMultipleTriggersNotFound{}
}

/*
DeleteScheduleMultipleTriggersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteScheduleMultipleTriggersNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete schedule multiple triggers not found response has a 2xx status code
func (o *DeleteScheduleMultipleTriggersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete schedule multiple triggers not found response has a 3xx status code
func (o *DeleteScheduleMultipleTriggersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete schedule multiple triggers not found response has a 4xx status code
func (o *DeleteScheduleMultipleTriggersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete schedule multiple triggers not found response has a 5xx status code
func (o *DeleteScheduleMultipleTriggersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete schedule multiple triggers not found response a status code equal to that given
func (o *DeleteScheduleMultipleTriggersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteScheduleMultipleTriggersNotFound) Error() string {
	return fmt.Sprintf("[DELETE /schedules/{scheduleId}][%d] deleteScheduleMultipleTriggersNotFound  %+v", 404, o.Payload)
}

func (o *DeleteScheduleMultipleTriggersNotFound) String() string {
	return fmt.Sprintf("[DELETE /schedules/{scheduleId}][%d] deleteScheduleMultipleTriggersNotFound  %+v", 404, o.Payload)
}

func (o *DeleteScheduleMultipleTriggersNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteScheduleMultipleTriggersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScheduleMultipleTriggersConflict creates a DeleteScheduleMultipleTriggersConflict with default headers values
func NewDeleteScheduleMultipleTriggersConflict() *DeleteScheduleMultipleTriggersConflict {
	return &DeleteScheduleMultipleTriggersConflict{}
}

/*
DeleteScheduleMultipleTriggersConflict describes a response with status code 409, with default header values.

The schedule is currently used by an executable
*/
type DeleteScheduleMultipleTriggersConflict struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete schedule multiple triggers conflict response has a 2xx status code
func (o *DeleteScheduleMultipleTriggersConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete schedule multiple triggers conflict response has a 3xx status code
func (o *DeleteScheduleMultipleTriggersConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete schedule multiple triggers conflict response has a 4xx status code
func (o *DeleteScheduleMultipleTriggersConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete schedule multiple triggers conflict response has a 5xx status code
func (o *DeleteScheduleMultipleTriggersConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete schedule multiple triggers conflict response a status code equal to that given
func (o *DeleteScheduleMultipleTriggersConflict) IsCode(code int) bool {
	return code == 409
}

func (o *DeleteScheduleMultipleTriggersConflict) Error() string {
	return fmt.Sprintf("[DELETE /schedules/{scheduleId}][%d] deleteScheduleMultipleTriggersConflict  %+v", 409, o.Payload)
}

func (o *DeleteScheduleMultipleTriggersConflict) String() string {
	return fmt.Sprintf("[DELETE /schedules/{scheduleId}][%d] deleteScheduleMultipleTriggersConflict  %+v", 409, o.Payload)
}

func (o *DeleteScheduleMultipleTriggersConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteScheduleMultipleTriggersConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScheduleMultipleTriggersInternalServerError creates a DeleteScheduleMultipleTriggersInternalServerError with default headers values
func NewDeleteScheduleMultipleTriggersInternalServerError() *DeleteScheduleMultipleTriggersInternalServerError {
	return &DeleteScheduleMultipleTriggersInternalServerError{}
}

/*
DeleteScheduleMultipleTriggersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteScheduleMultipleTriggersInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete schedule multiple triggers internal server error response has a 2xx status code
func (o *DeleteScheduleMultipleTriggersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete schedule multiple triggers internal server error response has a 3xx status code
func (o *DeleteScheduleMultipleTriggersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete schedule multiple triggers internal server error response has a 4xx status code
func (o *DeleteScheduleMultipleTriggersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete schedule multiple triggers internal server error response has a 5xx status code
func (o *DeleteScheduleMultipleTriggersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete schedule multiple triggers internal server error response a status code equal to that given
func (o *DeleteScheduleMultipleTriggersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteScheduleMultipleTriggersInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /schedules/{scheduleId}][%d] deleteScheduleMultipleTriggersInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteScheduleMultipleTriggersInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /schedules/{scheduleId}][%d] deleteScheduleMultipleTriggersInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteScheduleMultipleTriggersInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteScheduleMultipleTriggersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
