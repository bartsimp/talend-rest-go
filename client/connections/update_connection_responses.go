// Code generated by go-swagger; DO NOT EDIT.

package connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bartsimp/talend-rest-go/models"
)

// UpdateConnectionReader is a Reader for the UpdateConnection structure.
type UpdateConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateConnectionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateConnectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateConnectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateConnectionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateConnectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateConnectionOK creates a UpdateConnectionOK with default headers values
func NewUpdateConnectionOK() *UpdateConnectionOK {
	return &UpdateConnectionOK{}
}

/*
UpdateConnectionOK describes a response with status code 200, with default header values.

Updated
*/
type UpdateConnectionOK struct {
	Payload *models.ConnectionDetails
}

// IsSuccess returns true when this update connection o k response has a 2xx status code
func (o *UpdateConnectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update connection o k response has a 3xx status code
func (o *UpdateConnectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update connection o k response has a 4xx status code
func (o *UpdateConnectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update connection o k response has a 5xx status code
func (o *UpdateConnectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update connection o k response a status code equal to that given
func (o *UpdateConnectionOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateConnectionOK) Error() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionOK  %+v", 200, o.Payload)
}

func (o *UpdateConnectionOK) String() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionOK  %+v", 200, o.Payload)
}

func (o *UpdateConnectionOK) GetPayload() *models.ConnectionDetails {
	return o.Payload
}

func (o *UpdateConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConnectionDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConnectionBadRequest creates a UpdateConnectionBadRequest with default headers values
func NewUpdateConnectionBadRequest() *UpdateConnectionBadRequest {
	return &UpdateConnectionBadRequest{}
}

/*
UpdateConnectionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateConnectionBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update connection bad request response has a 2xx status code
func (o *UpdateConnectionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update connection bad request response has a 3xx status code
func (o *UpdateConnectionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update connection bad request response has a 4xx status code
func (o *UpdateConnectionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update connection bad request response has a 5xx status code
func (o *UpdateConnectionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update connection bad request response a status code equal to that given
func (o *UpdateConnectionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateConnectionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateConnectionBadRequest) String() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateConnectionBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConnectionUnauthorized creates a UpdateConnectionUnauthorized with default headers values
func NewUpdateConnectionUnauthorized() *UpdateConnectionUnauthorized {
	return &UpdateConnectionUnauthorized{}
}

/*
UpdateConnectionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateConnectionUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update connection unauthorized response has a 2xx status code
func (o *UpdateConnectionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update connection unauthorized response has a 3xx status code
func (o *UpdateConnectionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update connection unauthorized response has a 4xx status code
func (o *UpdateConnectionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update connection unauthorized response has a 5xx status code
func (o *UpdateConnectionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update connection unauthorized response a status code equal to that given
func (o *UpdateConnectionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateConnectionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateConnectionUnauthorized) String() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateConnectionUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateConnectionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConnectionForbidden creates a UpdateConnectionForbidden with default headers values
func NewUpdateConnectionForbidden() *UpdateConnectionForbidden {
	return &UpdateConnectionForbidden{}
}

/*
UpdateConnectionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateConnectionForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update connection forbidden response has a 2xx status code
func (o *UpdateConnectionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update connection forbidden response has a 3xx status code
func (o *UpdateConnectionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update connection forbidden response has a 4xx status code
func (o *UpdateConnectionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update connection forbidden response has a 5xx status code
func (o *UpdateConnectionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update connection forbidden response a status code equal to that given
func (o *UpdateConnectionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateConnectionForbidden) Error() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionForbidden  %+v", 403, o.Payload)
}

func (o *UpdateConnectionForbidden) String() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionForbidden  %+v", 403, o.Payload)
}

func (o *UpdateConnectionForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateConnectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConnectionNotFound creates a UpdateConnectionNotFound with default headers values
func NewUpdateConnectionNotFound() *UpdateConnectionNotFound {
	return &UpdateConnectionNotFound{}
}

/*
UpdateConnectionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateConnectionNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update connection not found response has a 2xx status code
func (o *UpdateConnectionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update connection not found response has a 3xx status code
func (o *UpdateConnectionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update connection not found response has a 4xx status code
func (o *UpdateConnectionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update connection not found response has a 5xx status code
func (o *UpdateConnectionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update connection not found response a status code equal to that given
func (o *UpdateConnectionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateConnectionNotFound) Error() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionNotFound  %+v", 404, o.Payload)
}

func (o *UpdateConnectionNotFound) String() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionNotFound  %+v", 404, o.Payload)
}

func (o *UpdateConnectionNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateConnectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConnectionConflict creates a UpdateConnectionConflict with default headers values
func NewUpdateConnectionConflict() *UpdateConnectionConflict {
	return &UpdateConnectionConflict{}
}

/*
UpdateConnectionConflict describes a response with status code 409, with default header values.

Connection name already exists
*/
type UpdateConnectionConflict struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update connection conflict response has a 2xx status code
func (o *UpdateConnectionConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update connection conflict response has a 3xx status code
func (o *UpdateConnectionConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update connection conflict response has a 4xx status code
func (o *UpdateConnectionConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update connection conflict response has a 5xx status code
func (o *UpdateConnectionConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update connection conflict response a status code equal to that given
func (o *UpdateConnectionConflict) IsCode(code int) bool {
	return code == 409
}

func (o *UpdateConnectionConflict) Error() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionConflict  %+v", 409, o.Payload)
}

func (o *UpdateConnectionConflict) String() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionConflict  %+v", 409, o.Payload)
}

func (o *UpdateConnectionConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateConnectionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConnectionInternalServerError creates a UpdateConnectionInternalServerError with default headers values
func NewUpdateConnectionInternalServerError() *UpdateConnectionInternalServerError {
	return &UpdateConnectionInternalServerError{}
}

/*
UpdateConnectionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateConnectionInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update connection internal server error response has a 2xx status code
func (o *UpdateConnectionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update connection internal server error response has a 3xx status code
func (o *UpdateConnectionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update connection internal server error response has a 4xx status code
func (o *UpdateConnectionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update connection internal server error response has a 5xx status code
func (o *UpdateConnectionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update connection internal server error response a status code equal to that given
func (o *UpdateConnectionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateConnectionInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateConnectionInternalServerError) String() string {
	return fmt.Sprintf("[PUT /connections/{id}][%d] updateConnectionInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateConnectionInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateConnectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
