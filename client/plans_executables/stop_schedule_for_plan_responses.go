// Code generated by go-swagger; DO NOT EDIT.

package plans_executables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bartsimp/talend-rest-go/models"
)

// StopScheduleForPlanReader is a Reader for the StopScheduleForPlan structure.
type StopScheduleForPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopScheduleForPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewStopScheduleForPlanNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStopScheduleForPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStopScheduleForPlanUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStopScheduleForPlanForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopScheduleForPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStopScheduleForPlanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopScheduleForPlanNoContent creates a StopScheduleForPlanNoContent with default headers values
func NewStopScheduleForPlanNoContent() *StopScheduleForPlanNoContent {
	return &StopScheduleForPlanNoContent{}
}

/*
StopScheduleForPlanNoContent describes a response with status code 204, with default header values.

No Content
*/
type StopScheduleForPlanNoContent struct {
}

// IsSuccess returns true when this stop schedule for plan no content response has a 2xx status code
func (o *StopScheduleForPlanNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop schedule for plan no content response has a 3xx status code
func (o *StopScheduleForPlanNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop schedule for plan no content response has a 4xx status code
func (o *StopScheduleForPlanNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop schedule for plan no content response has a 5xx status code
func (o *StopScheduleForPlanNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this stop schedule for plan no content response a status code equal to that given
func (o *StopScheduleForPlanNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *StopScheduleForPlanNoContent) Error() string {
	return fmt.Sprintf("[DELETE /executables/plans/{planId}/run-config][%d] stopScheduleForPlanNoContent ", 204)
}

func (o *StopScheduleForPlanNoContent) String() string {
	return fmt.Sprintf("[DELETE /executables/plans/{planId}/run-config][%d] stopScheduleForPlanNoContent ", 204)
}

func (o *StopScheduleForPlanNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopScheduleForPlanBadRequest creates a StopScheduleForPlanBadRequest with default headers values
func NewStopScheduleForPlanBadRequest() *StopScheduleForPlanBadRequest {
	return &StopScheduleForPlanBadRequest{}
}

/*
StopScheduleForPlanBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StopScheduleForPlanBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this stop schedule for plan bad request response has a 2xx status code
func (o *StopScheduleForPlanBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop schedule for plan bad request response has a 3xx status code
func (o *StopScheduleForPlanBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop schedule for plan bad request response has a 4xx status code
func (o *StopScheduleForPlanBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop schedule for plan bad request response has a 5xx status code
func (o *StopScheduleForPlanBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stop schedule for plan bad request response a status code equal to that given
func (o *StopScheduleForPlanBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StopScheduleForPlanBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /executables/plans/{planId}/run-config][%d] stopScheduleForPlanBadRequest  %+v", 400, o.Payload)
}

func (o *StopScheduleForPlanBadRequest) String() string {
	return fmt.Sprintf("[DELETE /executables/plans/{planId}/run-config][%d] stopScheduleForPlanBadRequest  %+v", 400, o.Payload)
}

func (o *StopScheduleForPlanBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *StopScheduleForPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopScheduleForPlanUnauthorized creates a StopScheduleForPlanUnauthorized with default headers values
func NewStopScheduleForPlanUnauthorized() *StopScheduleForPlanUnauthorized {
	return &StopScheduleForPlanUnauthorized{}
}

/*
StopScheduleForPlanUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StopScheduleForPlanUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this stop schedule for plan unauthorized response has a 2xx status code
func (o *StopScheduleForPlanUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop schedule for plan unauthorized response has a 3xx status code
func (o *StopScheduleForPlanUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop schedule for plan unauthorized response has a 4xx status code
func (o *StopScheduleForPlanUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop schedule for plan unauthorized response has a 5xx status code
func (o *StopScheduleForPlanUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stop schedule for plan unauthorized response a status code equal to that given
func (o *StopScheduleForPlanUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StopScheduleForPlanUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /executables/plans/{planId}/run-config][%d] stopScheduleForPlanUnauthorized  %+v", 401, o.Payload)
}

func (o *StopScheduleForPlanUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /executables/plans/{planId}/run-config][%d] stopScheduleForPlanUnauthorized  %+v", 401, o.Payload)
}

func (o *StopScheduleForPlanUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *StopScheduleForPlanUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopScheduleForPlanForbidden creates a StopScheduleForPlanForbidden with default headers values
func NewStopScheduleForPlanForbidden() *StopScheduleForPlanForbidden {
	return &StopScheduleForPlanForbidden{}
}

/*
StopScheduleForPlanForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StopScheduleForPlanForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this stop schedule for plan forbidden response has a 2xx status code
func (o *StopScheduleForPlanForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop schedule for plan forbidden response has a 3xx status code
func (o *StopScheduleForPlanForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop schedule for plan forbidden response has a 4xx status code
func (o *StopScheduleForPlanForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop schedule for plan forbidden response has a 5xx status code
func (o *StopScheduleForPlanForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stop schedule for plan forbidden response a status code equal to that given
func (o *StopScheduleForPlanForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StopScheduleForPlanForbidden) Error() string {
	return fmt.Sprintf("[DELETE /executables/plans/{planId}/run-config][%d] stopScheduleForPlanForbidden  %+v", 403, o.Payload)
}

func (o *StopScheduleForPlanForbidden) String() string {
	return fmt.Sprintf("[DELETE /executables/plans/{planId}/run-config][%d] stopScheduleForPlanForbidden  %+v", 403, o.Payload)
}

func (o *StopScheduleForPlanForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *StopScheduleForPlanForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopScheduleForPlanNotFound creates a StopScheduleForPlanNotFound with default headers values
func NewStopScheduleForPlanNotFound() *StopScheduleForPlanNotFound {
	return &StopScheduleForPlanNotFound{}
}

/*
StopScheduleForPlanNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StopScheduleForPlanNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this stop schedule for plan not found response has a 2xx status code
func (o *StopScheduleForPlanNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop schedule for plan not found response has a 3xx status code
func (o *StopScheduleForPlanNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop schedule for plan not found response has a 4xx status code
func (o *StopScheduleForPlanNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop schedule for plan not found response has a 5xx status code
func (o *StopScheduleForPlanNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stop schedule for plan not found response a status code equal to that given
func (o *StopScheduleForPlanNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StopScheduleForPlanNotFound) Error() string {
	return fmt.Sprintf("[DELETE /executables/plans/{planId}/run-config][%d] stopScheduleForPlanNotFound  %+v", 404, o.Payload)
}

func (o *StopScheduleForPlanNotFound) String() string {
	return fmt.Sprintf("[DELETE /executables/plans/{planId}/run-config][%d] stopScheduleForPlanNotFound  %+v", 404, o.Payload)
}

func (o *StopScheduleForPlanNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *StopScheduleForPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopScheduleForPlanInternalServerError creates a StopScheduleForPlanInternalServerError with default headers values
func NewStopScheduleForPlanInternalServerError() *StopScheduleForPlanInternalServerError {
	return &StopScheduleForPlanInternalServerError{}
}

/*
StopScheduleForPlanInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type StopScheduleForPlanInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this stop schedule for plan internal server error response has a 2xx status code
func (o *StopScheduleForPlanInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop schedule for plan internal server error response has a 3xx status code
func (o *StopScheduleForPlanInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop schedule for plan internal server error response has a 4xx status code
func (o *StopScheduleForPlanInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop schedule for plan internal server error response has a 5xx status code
func (o *StopScheduleForPlanInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stop schedule for plan internal server error response a status code equal to that given
func (o *StopScheduleForPlanInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StopScheduleForPlanInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /executables/plans/{planId}/run-config][%d] stopScheduleForPlanInternalServerError  %+v", 500, o.Payload)
}

func (o *StopScheduleForPlanInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /executables/plans/{planId}/run-config][%d] stopScheduleForPlanInternalServerError  %+v", 500, o.Payload)
}

func (o *StopScheduleForPlanInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *StopScheduleForPlanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}