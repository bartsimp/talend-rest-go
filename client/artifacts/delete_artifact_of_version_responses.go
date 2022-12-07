// Code generated by go-swagger; DO NOT EDIT.

package artifacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bartsimp/talend-rest-go/models"
)

// DeleteArtifactOfVersionReader is a Reader for the DeleteArtifactOfVersion structure.
type DeleteArtifactOfVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteArtifactOfVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteArtifactOfVersionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteArtifactOfVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteArtifactOfVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteArtifactOfVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteArtifactOfVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteArtifactOfVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteArtifactOfVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteArtifactOfVersionNoContent creates a DeleteArtifactOfVersionNoContent with default headers values
func NewDeleteArtifactOfVersionNoContent() *DeleteArtifactOfVersionNoContent {
	return &DeleteArtifactOfVersionNoContent{}
}

/*
DeleteArtifactOfVersionNoContent describes a response with status code 204, with default header values.

No content
*/
type DeleteArtifactOfVersionNoContent struct {
}

// IsSuccess returns true when this delete artifact of version no content response has a 2xx status code
func (o *DeleteArtifactOfVersionNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete artifact of version no content response has a 3xx status code
func (o *DeleteArtifactOfVersionNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete artifact of version no content response has a 4xx status code
func (o *DeleteArtifactOfVersionNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete artifact of version no content response has a 5xx status code
func (o *DeleteArtifactOfVersionNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete artifact of version no content response a status code equal to that given
func (o *DeleteArtifactOfVersionNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteArtifactOfVersionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /artifacts/{id}/versions/{version}][%d] deleteArtifactOfVersionNoContent ", 204)
}

func (o *DeleteArtifactOfVersionNoContent) String() string {
	return fmt.Sprintf("[DELETE /artifacts/{id}/versions/{version}][%d] deleteArtifactOfVersionNoContent ", 204)
}

func (o *DeleteArtifactOfVersionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteArtifactOfVersionBadRequest creates a DeleteArtifactOfVersionBadRequest with default headers values
func NewDeleteArtifactOfVersionBadRequest() *DeleteArtifactOfVersionBadRequest {
	return &DeleteArtifactOfVersionBadRequest{}
}

/*
DeleteArtifactOfVersionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteArtifactOfVersionBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete artifact of version bad request response has a 2xx status code
func (o *DeleteArtifactOfVersionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete artifact of version bad request response has a 3xx status code
func (o *DeleteArtifactOfVersionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete artifact of version bad request response has a 4xx status code
func (o *DeleteArtifactOfVersionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete artifact of version bad request response has a 5xx status code
func (o *DeleteArtifactOfVersionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete artifact of version bad request response a status code equal to that given
func (o *DeleteArtifactOfVersionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteArtifactOfVersionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /artifacts/{id}/versions/{version}][%d] deleteArtifactOfVersionBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteArtifactOfVersionBadRequest) String() string {
	return fmt.Sprintf("[DELETE /artifacts/{id}/versions/{version}][%d] deleteArtifactOfVersionBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteArtifactOfVersionBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteArtifactOfVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArtifactOfVersionUnauthorized creates a DeleteArtifactOfVersionUnauthorized with default headers values
func NewDeleteArtifactOfVersionUnauthorized() *DeleteArtifactOfVersionUnauthorized {
	return &DeleteArtifactOfVersionUnauthorized{}
}

/*
DeleteArtifactOfVersionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteArtifactOfVersionUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete artifact of version unauthorized response has a 2xx status code
func (o *DeleteArtifactOfVersionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete artifact of version unauthorized response has a 3xx status code
func (o *DeleteArtifactOfVersionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete artifact of version unauthorized response has a 4xx status code
func (o *DeleteArtifactOfVersionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete artifact of version unauthorized response has a 5xx status code
func (o *DeleteArtifactOfVersionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete artifact of version unauthorized response a status code equal to that given
func (o *DeleteArtifactOfVersionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteArtifactOfVersionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /artifacts/{id}/versions/{version}][%d] deleteArtifactOfVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteArtifactOfVersionUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /artifacts/{id}/versions/{version}][%d] deleteArtifactOfVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteArtifactOfVersionUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteArtifactOfVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArtifactOfVersionForbidden creates a DeleteArtifactOfVersionForbidden with default headers values
func NewDeleteArtifactOfVersionForbidden() *DeleteArtifactOfVersionForbidden {
	return &DeleteArtifactOfVersionForbidden{}
}

/*
DeleteArtifactOfVersionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteArtifactOfVersionForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete artifact of version forbidden response has a 2xx status code
func (o *DeleteArtifactOfVersionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete artifact of version forbidden response has a 3xx status code
func (o *DeleteArtifactOfVersionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete artifact of version forbidden response has a 4xx status code
func (o *DeleteArtifactOfVersionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete artifact of version forbidden response has a 5xx status code
func (o *DeleteArtifactOfVersionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete artifact of version forbidden response a status code equal to that given
func (o *DeleteArtifactOfVersionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteArtifactOfVersionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /artifacts/{id}/versions/{version}][%d] deleteArtifactOfVersionForbidden  %+v", 403, o.Payload)
}

func (o *DeleteArtifactOfVersionForbidden) String() string {
	return fmt.Sprintf("[DELETE /artifacts/{id}/versions/{version}][%d] deleteArtifactOfVersionForbidden  %+v", 403, o.Payload)
}

func (o *DeleteArtifactOfVersionForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteArtifactOfVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArtifactOfVersionNotFound creates a DeleteArtifactOfVersionNotFound with default headers values
func NewDeleteArtifactOfVersionNotFound() *DeleteArtifactOfVersionNotFound {
	return &DeleteArtifactOfVersionNotFound{}
}

/*
DeleteArtifactOfVersionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteArtifactOfVersionNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete artifact of version not found response has a 2xx status code
func (o *DeleteArtifactOfVersionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete artifact of version not found response has a 3xx status code
func (o *DeleteArtifactOfVersionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete artifact of version not found response has a 4xx status code
func (o *DeleteArtifactOfVersionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete artifact of version not found response has a 5xx status code
func (o *DeleteArtifactOfVersionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete artifact of version not found response a status code equal to that given
func (o *DeleteArtifactOfVersionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteArtifactOfVersionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /artifacts/{id}/versions/{version}][%d] deleteArtifactOfVersionNotFound  %+v", 404, o.Payload)
}

func (o *DeleteArtifactOfVersionNotFound) String() string {
	return fmt.Sprintf("[DELETE /artifacts/{id}/versions/{version}][%d] deleteArtifactOfVersionNotFound  %+v", 404, o.Payload)
}

func (o *DeleteArtifactOfVersionNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteArtifactOfVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArtifactOfVersionConflict creates a DeleteArtifactOfVersionConflict with default headers values
func NewDeleteArtifactOfVersionConflict() *DeleteArtifactOfVersionConflict {
	return &DeleteArtifactOfVersionConflict{}
}

/*
DeleteArtifactOfVersionConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteArtifactOfVersionConflict struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete artifact of version conflict response has a 2xx status code
func (o *DeleteArtifactOfVersionConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete artifact of version conflict response has a 3xx status code
func (o *DeleteArtifactOfVersionConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete artifact of version conflict response has a 4xx status code
func (o *DeleteArtifactOfVersionConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete artifact of version conflict response has a 5xx status code
func (o *DeleteArtifactOfVersionConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete artifact of version conflict response a status code equal to that given
func (o *DeleteArtifactOfVersionConflict) IsCode(code int) bool {
	return code == 409
}

func (o *DeleteArtifactOfVersionConflict) Error() string {
	return fmt.Sprintf("[DELETE /artifacts/{id}/versions/{version}][%d] deleteArtifactOfVersionConflict  %+v", 409, o.Payload)
}

func (o *DeleteArtifactOfVersionConflict) String() string {
	return fmt.Sprintf("[DELETE /artifacts/{id}/versions/{version}][%d] deleteArtifactOfVersionConflict  %+v", 409, o.Payload)
}

func (o *DeleteArtifactOfVersionConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteArtifactOfVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArtifactOfVersionInternalServerError creates a DeleteArtifactOfVersionInternalServerError with default headers values
func NewDeleteArtifactOfVersionInternalServerError() *DeleteArtifactOfVersionInternalServerError {
	return &DeleteArtifactOfVersionInternalServerError{}
}

/*
DeleteArtifactOfVersionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteArtifactOfVersionInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete artifact of version internal server error response has a 2xx status code
func (o *DeleteArtifactOfVersionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete artifact of version internal server error response has a 3xx status code
func (o *DeleteArtifactOfVersionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete artifact of version internal server error response has a 4xx status code
func (o *DeleteArtifactOfVersionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete artifact of version internal server error response has a 5xx status code
func (o *DeleteArtifactOfVersionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete artifact of version internal server error response a status code equal to that given
func (o *DeleteArtifactOfVersionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteArtifactOfVersionInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /artifacts/{id}/versions/{version}][%d] deleteArtifactOfVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteArtifactOfVersionInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /artifacts/{id}/versions/{version}][%d] deleteArtifactOfVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteArtifactOfVersionInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteArtifactOfVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}