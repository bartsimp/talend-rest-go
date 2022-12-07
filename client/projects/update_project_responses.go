// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bartsimp/talend-rest-go/models"
)

// UpdateProjectReader is a Reader for the UpdateProject structure.
type UpdateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateProjectConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProjectOK creates a UpdateProjectOK with default headers values
func NewUpdateProjectOK() *UpdateProjectOK {
	return &UpdateProjectOK{}
}

/*
UpdateProjectOK describes a response with status code 200, with default header values.

OK
*/
type UpdateProjectOK struct {
	Payload *models.ProjectView
}

// IsSuccess returns true when this update project o k response has a 2xx status code
func (o *UpdateProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update project o k response has a 3xx status code
func (o *UpdateProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project o k response has a 4xx status code
func (o *UpdateProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update project o k response has a 5xx status code
func (o *UpdateProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update project o k response a status code equal to that given
func (o *UpdateProjectOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateProjectOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{projectId}][%d] updateProjectOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectOK) String() string {
	return fmt.Sprintf("[PUT /projects/{projectId}][%d] updateProjectOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectOK) GetPayload() *models.ProjectView {
	return o.Payload
}

func (o *UpdateProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectView)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectBadRequest creates a UpdateProjectBadRequest with default headers values
func NewUpdateProjectBadRequest() *UpdateProjectBadRequest {
	return &UpdateProjectBadRequest{}
}

/*
UpdateProjectBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateProjectBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update project bad request response has a 2xx status code
func (o *UpdateProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project bad request response has a 3xx status code
func (o *UpdateProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project bad request response has a 4xx status code
func (o *UpdateProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project bad request response has a 5xx status code
func (o *UpdateProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update project bad request response a status code equal to that given
func (o *UpdateProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateProjectBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{projectId}][%d] updateProjectBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateProjectBadRequest) String() string {
	return fmt.Sprintf("[PUT /projects/{projectId}][%d] updateProjectBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateProjectBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectUnauthorized creates a UpdateProjectUnauthorized with default headers values
func NewUpdateProjectUnauthorized() *UpdateProjectUnauthorized {
	return &UpdateProjectUnauthorized{}
}

/*
UpdateProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateProjectUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update project unauthorized response has a 2xx status code
func (o *UpdateProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project unauthorized response has a 3xx status code
func (o *UpdateProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project unauthorized response has a 4xx status code
func (o *UpdateProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project unauthorized response has a 5xx status code
func (o *UpdateProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update project unauthorized response a status code equal to that given
func (o *UpdateProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateProjectUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projects/{projectId}][%d] updateProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateProjectUnauthorized) String() string {
	return fmt.Sprintf("[PUT /projects/{projectId}][%d] updateProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateProjectUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectForbidden creates a UpdateProjectForbidden with default headers values
func NewUpdateProjectForbidden() *UpdateProjectForbidden {
	return &UpdateProjectForbidden{}
}

/*
UpdateProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateProjectForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update project forbidden response has a 2xx status code
func (o *UpdateProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project forbidden response has a 3xx status code
func (o *UpdateProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project forbidden response has a 4xx status code
func (o *UpdateProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project forbidden response has a 5xx status code
func (o *UpdateProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update project forbidden response a status code equal to that given
func (o *UpdateProjectForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateProjectForbidden) Error() string {
	return fmt.Sprintf("[PUT /projects/{projectId}][%d] updateProjectForbidden  %+v", 403, o.Payload)
}

func (o *UpdateProjectForbidden) String() string {
	return fmt.Sprintf("[PUT /projects/{projectId}][%d] updateProjectForbidden  %+v", 403, o.Payload)
}

func (o *UpdateProjectForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectNotFound creates a UpdateProjectNotFound with default headers values
func NewUpdateProjectNotFound() *UpdateProjectNotFound {
	return &UpdateProjectNotFound{}
}

/*
UpdateProjectNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateProjectNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update project not found response has a 2xx status code
func (o *UpdateProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project not found response has a 3xx status code
func (o *UpdateProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project not found response has a 4xx status code
func (o *UpdateProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project not found response has a 5xx status code
func (o *UpdateProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update project not found response a status code equal to that given
func (o *UpdateProjectNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateProjectNotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{projectId}][%d] updateProjectNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProjectNotFound) String() string {
	return fmt.Sprintf("[PUT /projects/{projectId}][%d] updateProjectNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProjectNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectConflict creates a UpdateProjectConflict with default headers values
func NewUpdateProjectConflict() *UpdateProjectConflict {
	return &UpdateProjectConflict{}
}

/*
UpdateProjectConflict describes a response with status code 409, with default header values.

Project already exists
*/
type UpdateProjectConflict struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update project conflict response has a 2xx status code
func (o *UpdateProjectConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project conflict response has a 3xx status code
func (o *UpdateProjectConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project conflict response has a 4xx status code
func (o *UpdateProjectConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update project conflict response has a 5xx status code
func (o *UpdateProjectConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update project conflict response a status code equal to that given
func (o *UpdateProjectConflict) IsCode(code int) bool {
	return code == 409
}

func (o *UpdateProjectConflict) Error() string {
	return fmt.Sprintf("[PUT /projects/{projectId}][%d] updateProjectConflict  %+v", 409, o.Payload)
}

func (o *UpdateProjectConflict) String() string {
	return fmt.Sprintf("[PUT /projects/{projectId}][%d] updateProjectConflict  %+v", 409, o.Payload)
}

func (o *UpdateProjectConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateProjectConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectInternalServerError creates a UpdateProjectInternalServerError with default headers values
func NewUpdateProjectInternalServerError() *UpdateProjectInternalServerError {
	return &UpdateProjectInternalServerError{}
}

/*
UpdateProjectInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateProjectInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update project internal server error response has a 2xx status code
func (o *UpdateProjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update project internal server error response has a 3xx status code
func (o *UpdateProjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update project internal server error response has a 4xx status code
func (o *UpdateProjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update project internal server error response has a 5xx status code
func (o *UpdateProjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update project internal server error response a status code equal to that given
func (o *UpdateProjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateProjectInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{projectId}][%d] updateProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateProjectInternalServerError) String() string {
	return fmt.Sprintf("[PUT /projects/{projectId}][%d] updateProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateProjectInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
