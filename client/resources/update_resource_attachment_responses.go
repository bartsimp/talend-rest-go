// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bartsimp/talend-rest-go/models"
)

// UpdateResourceAttachmentReader is a Reader for the UpdateResourceAttachment structure.
type UpdateResourceAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateResourceAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateResourceAttachmentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateResourceAttachmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateResourceAttachmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateResourceAttachmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateResourceAttachmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewUpdateResourceAttachmentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateResourceAttachmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateResourceAttachmentNoContent creates a UpdateResourceAttachmentNoContent with default headers values
func NewUpdateResourceAttachmentNoContent() *UpdateResourceAttachmentNoContent {
	return &UpdateResourceAttachmentNoContent{}
}

/*
UpdateResourceAttachmentNoContent describes a response with status code 204, with default header values.

no content
*/
type UpdateResourceAttachmentNoContent struct {
}

// IsSuccess returns true when this update resource attachment no content response has a 2xx status code
func (o *UpdateResourceAttachmentNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update resource attachment no content response has a 3xx status code
func (o *UpdateResourceAttachmentNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update resource attachment no content response has a 4xx status code
func (o *UpdateResourceAttachmentNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update resource attachment no content response has a 5xx status code
func (o *UpdateResourceAttachmentNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update resource attachment no content response a status code equal to that given
func (o *UpdateResourceAttachmentNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateResourceAttachmentNoContent) Error() string {
	return fmt.Sprintf("[PUT /resources/{id}/upload][%d] updateResourceAttachmentNoContent ", 204)
}

func (o *UpdateResourceAttachmentNoContent) String() string {
	return fmt.Sprintf("[PUT /resources/{id}/upload][%d] updateResourceAttachmentNoContent ", 204)
}

func (o *UpdateResourceAttachmentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateResourceAttachmentBadRequest creates a UpdateResourceAttachmentBadRequest with default headers values
func NewUpdateResourceAttachmentBadRequest() *UpdateResourceAttachmentBadRequest {
	return &UpdateResourceAttachmentBadRequest{}
}

/*
UpdateResourceAttachmentBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateResourceAttachmentBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update resource attachment bad request response has a 2xx status code
func (o *UpdateResourceAttachmentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update resource attachment bad request response has a 3xx status code
func (o *UpdateResourceAttachmentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update resource attachment bad request response has a 4xx status code
func (o *UpdateResourceAttachmentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update resource attachment bad request response has a 5xx status code
func (o *UpdateResourceAttachmentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update resource attachment bad request response a status code equal to that given
func (o *UpdateResourceAttachmentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateResourceAttachmentBadRequest) Error() string {
	return fmt.Sprintf("[PUT /resources/{id}/upload][%d] updateResourceAttachmentBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateResourceAttachmentBadRequest) String() string {
	return fmt.Sprintf("[PUT /resources/{id}/upload][%d] updateResourceAttachmentBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateResourceAttachmentBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateResourceAttachmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateResourceAttachmentUnauthorized creates a UpdateResourceAttachmentUnauthorized with default headers values
func NewUpdateResourceAttachmentUnauthorized() *UpdateResourceAttachmentUnauthorized {
	return &UpdateResourceAttachmentUnauthorized{}
}

/*
UpdateResourceAttachmentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateResourceAttachmentUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update resource attachment unauthorized response has a 2xx status code
func (o *UpdateResourceAttachmentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update resource attachment unauthorized response has a 3xx status code
func (o *UpdateResourceAttachmentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update resource attachment unauthorized response has a 4xx status code
func (o *UpdateResourceAttachmentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update resource attachment unauthorized response has a 5xx status code
func (o *UpdateResourceAttachmentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update resource attachment unauthorized response a status code equal to that given
func (o *UpdateResourceAttachmentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateResourceAttachmentUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /resources/{id}/upload][%d] updateResourceAttachmentUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateResourceAttachmentUnauthorized) String() string {
	return fmt.Sprintf("[PUT /resources/{id}/upload][%d] updateResourceAttachmentUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateResourceAttachmentUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateResourceAttachmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateResourceAttachmentForbidden creates a UpdateResourceAttachmentForbidden with default headers values
func NewUpdateResourceAttachmentForbidden() *UpdateResourceAttachmentForbidden {
	return &UpdateResourceAttachmentForbidden{}
}

/*
UpdateResourceAttachmentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateResourceAttachmentForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update resource attachment forbidden response has a 2xx status code
func (o *UpdateResourceAttachmentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update resource attachment forbidden response has a 3xx status code
func (o *UpdateResourceAttachmentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update resource attachment forbidden response has a 4xx status code
func (o *UpdateResourceAttachmentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update resource attachment forbidden response has a 5xx status code
func (o *UpdateResourceAttachmentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update resource attachment forbidden response a status code equal to that given
func (o *UpdateResourceAttachmentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateResourceAttachmentForbidden) Error() string {
	return fmt.Sprintf("[PUT /resources/{id}/upload][%d] updateResourceAttachmentForbidden  %+v", 403, o.Payload)
}

func (o *UpdateResourceAttachmentForbidden) String() string {
	return fmt.Sprintf("[PUT /resources/{id}/upload][%d] updateResourceAttachmentForbidden  %+v", 403, o.Payload)
}

func (o *UpdateResourceAttachmentForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateResourceAttachmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateResourceAttachmentNotFound creates a UpdateResourceAttachmentNotFound with default headers values
func NewUpdateResourceAttachmentNotFound() *UpdateResourceAttachmentNotFound {
	return &UpdateResourceAttachmentNotFound{}
}

/*
UpdateResourceAttachmentNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateResourceAttachmentNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update resource attachment not found response has a 2xx status code
func (o *UpdateResourceAttachmentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update resource attachment not found response has a 3xx status code
func (o *UpdateResourceAttachmentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update resource attachment not found response has a 4xx status code
func (o *UpdateResourceAttachmentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update resource attachment not found response has a 5xx status code
func (o *UpdateResourceAttachmentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update resource attachment not found response a status code equal to that given
func (o *UpdateResourceAttachmentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateResourceAttachmentNotFound) Error() string {
	return fmt.Sprintf("[PUT /resources/{id}/upload][%d] updateResourceAttachmentNotFound  %+v", 404, o.Payload)
}

func (o *UpdateResourceAttachmentNotFound) String() string {
	return fmt.Sprintf("[PUT /resources/{id}/upload][%d] updateResourceAttachmentNotFound  %+v", 404, o.Payload)
}

func (o *UpdateResourceAttachmentNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateResourceAttachmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateResourceAttachmentRequestEntityTooLarge creates a UpdateResourceAttachmentRequestEntityTooLarge with default headers values
func NewUpdateResourceAttachmentRequestEntityTooLarge() *UpdateResourceAttachmentRequestEntityTooLarge {
	return &UpdateResourceAttachmentRequestEntityTooLarge{}
}

/*
UpdateResourceAttachmentRequestEntityTooLarge describes a response with status code 413, with default header values.

Upload limit exceeded
*/
type UpdateResourceAttachmentRequestEntityTooLarge struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update resource attachment request entity too large response has a 2xx status code
func (o *UpdateResourceAttachmentRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update resource attachment request entity too large response has a 3xx status code
func (o *UpdateResourceAttachmentRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update resource attachment request entity too large response has a 4xx status code
func (o *UpdateResourceAttachmentRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this update resource attachment request entity too large response has a 5xx status code
func (o *UpdateResourceAttachmentRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this update resource attachment request entity too large response a status code equal to that given
func (o *UpdateResourceAttachmentRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *UpdateResourceAttachmentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /resources/{id}/upload][%d] updateResourceAttachmentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *UpdateResourceAttachmentRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /resources/{id}/upload][%d] updateResourceAttachmentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *UpdateResourceAttachmentRequestEntityTooLarge) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateResourceAttachmentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateResourceAttachmentInternalServerError creates a UpdateResourceAttachmentInternalServerError with default headers values
func NewUpdateResourceAttachmentInternalServerError() *UpdateResourceAttachmentInternalServerError {
	return &UpdateResourceAttachmentInternalServerError{}
}

/*
UpdateResourceAttachmentInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateResourceAttachmentInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update resource attachment internal server error response has a 2xx status code
func (o *UpdateResourceAttachmentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update resource attachment internal server error response has a 3xx status code
func (o *UpdateResourceAttachmentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update resource attachment internal server error response has a 4xx status code
func (o *UpdateResourceAttachmentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update resource attachment internal server error response has a 5xx status code
func (o *UpdateResourceAttachmentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update resource attachment internal server error response a status code equal to that given
func (o *UpdateResourceAttachmentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateResourceAttachmentInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /resources/{id}/upload][%d] updateResourceAttachmentInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateResourceAttachmentInternalServerError) String() string {
	return fmt.Sprintf("[PUT /resources/{id}/upload][%d] updateResourceAttachmentInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateResourceAttachmentInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateResourceAttachmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}