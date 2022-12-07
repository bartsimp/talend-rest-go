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

// GetAllSchedulesMultipleTriggersReader is a Reader for the GetAllSchedulesMultipleTriggers structure.
type GetAllSchedulesMultipleTriggersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllSchedulesMultipleTriggersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllSchedulesMultipleTriggersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAllSchedulesMultipleTriggersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAllSchedulesMultipleTriggersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllSchedulesMultipleTriggersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllSchedulesMultipleTriggersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllSchedulesMultipleTriggersOK creates a GetAllSchedulesMultipleTriggersOK with default headers values
func NewGetAllSchedulesMultipleTriggersOK() *GetAllSchedulesMultipleTriggersOK {
	return &GetAllSchedulesMultipleTriggersOK{}
}

/*
GetAllSchedulesMultipleTriggersOK describes a response with status code 200, with default header values.

OK
*/
type GetAllSchedulesMultipleTriggersOK struct {
	Payload *models.Page
}

// IsSuccess returns true when this get all schedules multiple triggers o k response has a 2xx status code
func (o *GetAllSchedulesMultipleTriggersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all schedules multiple triggers o k response has a 3xx status code
func (o *GetAllSchedulesMultipleTriggersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all schedules multiple triggers o k response has a 4xx status code
func (o *GetAllSchedulesMultipleTriggersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all schedules multiple triggers o k response has a 5xx status code
func (o *GetAllSchedulesMultipleTriggersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all schedules multiple triggers o k response a status code equal to that given
func (o *GetAllSchedulesMultipleTriggersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAllSchedulesMultipleTriggersOK) Error() string {
	return fmt.Sprintf("[GET /schedules][%d] getAllSchedulesMultipleTriggersOK  %+v", 200, o.Payload)
}

func (o *GetAllSchedulesMultipleTriggersOK) String() string {
	return fmt.Sprintf("[GET /schedules][%d] getAllSchedulesMultipleTriggersOK  %+v", 200, o.Payload)
}

func (o *GetAllSchedulesMultipleTriggersOK) GetPayload() *models.Page {
	return o.Payload
}

func (o *GetAllSchedulesMultipleTriggersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Page)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllSchedulesMultipleTriggersBadRequest creates a GetAllSchedulesMultipleTriggersBadRequest with default headers values
func NewGetAllSchedulesMultipleTriggersBadRequest() *GetAllSchedulesMultipleTriggersBadRequest {
	return &GetAllSchedulesMultipleTriggersBadRequest{}
}

/*
GetAllSchedulesMultipleTriggersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAllSchedulesMultipleTriggersBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get all schedules multiple triggers bad request response has a 2xx status code
func (o *GetAllSchedulesMultipleTriggersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all schedules multiple triggers bad request response has a 3xx status code
func (o *GetAllSchedulesMultipleTriggersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all schedules multiple triggers bad request response has a 4xx status code
func (o *GetAllSchedulesMultipleTriggersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all schedules multiple triggers bad request response has a 5xx status code
func (o *GetAllSchedulesMultipleTriggersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get all schedules multiple triggers bad request response a status code equal to that given
func (o *GetAllSchedulesMultipleTriggersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAllSchedulesMultipleTriggersBadRequest) Error() string {
	return fmt.Sprintf("[GET /schedules][%d] getAllSchedulesMultipleTriggersBadRequest  %+v", 400, o.Payload)
}

func (o *GetAllSchedulesMultipleTriggersBadRequest) String() string {
	return fmt.Sprintf("[GET /schedules][%d] getAllSchedulesMultipleTriggersBadRequest  %+v", 400, o.Payload)
}

func (o *GetAllSchedulesMultipleTriggersBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAllSchedulesMultipleTriggersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllSchedulesMultipleTriggersUnauthorized creates a GetAllSchedulesMultipleTriggersUnauthorized with default headers values
func NewGetAllSchedulesMultipleTriggersUnauthorized() *GetAllSchedulesMultipleTriggersUnauthorized {
	return &GetAllSchedulesMultipleTriggersUnauthorized{}
}

/*
GetAllSchedulesMultipleTriggersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAllSchedulesMultipleTriggersUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get all schedules multiple triggers unauthorized response has a 2xx status code
func (o *GetAllSchedulesMultipleTriggersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all schedules multiple triggers unauthorized response has a 3xx status code
func (o *GetAllSchedulesMultipleTriggersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all schedules multiple triggers unauthorized response has a 4xx status code
func (o *GetAllSchedulesMultipleTriggersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all schedules multiple triggers unauthorized response has a 5xx status code
func (o *GetAllSchedulesMultipleTriggersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get all schedules multiple triggers unauthorized response a status code equal to that given
func (o *GetAllSchedulesMultipleTriggersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAllSchedulesMultipleTriggersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /schedules][%d] getAllSchedulesMultipleTriggersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAllSchedulesMultipleTriggersUnauthorized) String() string {
	return fmt.Sprintf("[GET /schedules][%d] getAllSchedulesMultipleTriggersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAllSchedulesMultipleTriggersUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAllSchedulesMultipleTriggersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllSchedulesMultipleTriggersForbidden creates a GetAllSchedulesMultipleTriggersForbidden with default headers values
func NewGetAllSchedulesMultipleTriggersForbidden() *GetAllSchedulesMultipleTriggersForbidden {
	return &GetAllSchedulesMultipleTriggersForbidden{}
}

/*
GetAllSchedulesMultipleTriggersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllSchedulesMultipleTriggersForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get all schedules multiple triggers forbidden response has a 2xx status code
func (o *GetAllSchedulesMultipleTriggersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all schedules multiple triggers forbidden response has a 3xx status code
func (o *GetAllSchedulesMultipleTriggersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all schedules multiple triggers forbidden response has a 4xx status code
func (o *GetAllSchedulesMultipleTriggersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all schedules multiple triggers forbidden response has a 5xx status code
func (o *GetAllSchedulesMultipleTriggersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get all schedules multiple triggers forbidden response a status code equal to that given
func (o *GetAllSchedulesMultipleTriggersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAllSchedulesMultipleTriggersForbidden) Error() string {
	return fmt.Sprintf("[GET /schedules][%d] getAllSchedulesMultipleTriggersForbidden  %+v", 403, o.Payload)
}

func (o *GetAllSchedulesMultipleTriggersForbidden) String() string {
	return fmt.Sprintf("[GET /schedules][%d] getAllSchedulesMultipleTriggersForbidden  %+v", 403, o.Payload)
}

func (o *GetAllSchedulesMultipleTriggersForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAllSchedulesMultipleTriggersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllSchedulesMultipleTriggersInternalServerError creates a GetAllSchedulesMultipleTriggersInternalServerError with default headers values
func NewGetAllSchedulesMultipleTriggersInternalServerError() *GetAllSchedulesMultipleTriggersInternalServerError {
	return &GetAllSchedulesMultipleTriggersInternalServerError{}
}

/*
GetAllSchedulesMultipleTriggersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAllSchedulesMultipleTriggersInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get all schedules multiple triggers internal server error response has a 2xx status code
func (o *GetAllSchedulesMultipleTriggersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all schedules multiple triggers internal server error response has a 3xx status code
func (o *GetAllSchedulesMultipleTriggersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all schedules multiple triggers internal server error response has a 4xx status code
func (o *GetAllSchedulesMultipleTriggersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all schedules multiple triggers internal server error response has a 5xx status code
func (o *GetAllSchedulesMultipleTriggersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get all schedules multiple triggers internal server error response a status code equal to that given
func (o *GetAllSchedulesMultipleTriggersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAllSchedulesMultipleTriggersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /schedules][%d] getAllSchedulesMultipleTriggersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllSchedulesMultipleTriggersInternalServerError) String() string {
	return fmt.Sprintf("[GET /schedules][%d] getAllSchedulesMultipleTriggersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllSchedulesMultipleTriggersInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAllSchedulesMultipleTriggersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
