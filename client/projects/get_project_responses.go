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

// GetProjectReader is a Reader for the GetProject structure.
type GetProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectOK creates a GetProjectOK with default headers values
func NewGetProjectOK() *GetProjectOK {
	return &GetProjectOK{}
}

/*
GetProjectOK describes a response with status code 200, with default header values.

OK
*/
type GetProjectOK struct {
	Payload *models.ProjectView
}

// IsSuccess returns true when this get project o k response has a 2xx status code
func (o *GetProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project o k response has a 3xx status code
func (o *GetProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project o k response has a 4xx status code
func (o *GetProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project o k response has a 5xx status code
func (o *GetProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get project o k response a status code equal to that given
func (o *GetProjectOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetProjectOK) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectOK  %+v", 200, o.Payload)
}

func (o *GetProjectOK) String() string {
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectOK  %+v", 200, o.Payload)
}

func (o *GetProjectOK) GetPayload() *models.ProjectView {
	return o.Payload
}

func (o *GetProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectView)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectBadRequest creates a GetProjectBadRequest with default headers values
func NewGetProjectBadRequest() *GetProjectBadRequest {
	return &GetProjectBadRequest{}
}

/*
GetProjectBadRequest describes a response with status code 400, with default header values.

Project Id is not valid id.
*/
type GetProjectBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get project bad request response has a 2xx status code
func (o *GetProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project bad request response has a 3xx status code
func (o *GetProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project bad request response has a 4xx status code
func (o *GetProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project bad request response has a 5xx status code
func (o *GetProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get project bad request response a status code equal to that given
func (o *GetProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetProjectBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectBadRequest  %+v", 400, o.Payload)
}

func (o *GetProjectBadRequest) String() string {
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectBadRequest  %+v", 400, o.Payload)
}

func (o *GetProjectBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectUnauthorized creates a GetProjectUnauthorized with default headers values
func NewGetProjectUnauthorized() *GetProjectUnauthorized {
	return &GetProjectUnauthorized{}
}

/*
GetProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetProjectUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get project unauthorized response has a 2xx status code
func (o *GetProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project unauthorized response has a 3xx status code
func (o *GetProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project unauthorized response has a 4xx status code
func (o *GetProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project unauthorized response has a 5xx status code
func (o *GetProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get project unauthorized response a status code equal to that given
func (o *GetProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *GetProjectUnauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *GetProjectUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectForbidden creates a GetProjectForbidden with default headers values
func NewGetProjectForbidden() *GetProjectForbidden {
	return &GetProjectForbidden{}
}

/*
GetProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetProjectForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get project forbidden response has a 2xx status code
func (o *GetProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project forbidden response has a 3xx status code
func (o *GetProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project forbidden response has a 4xx status code
func (o *GetProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project forbidden response has a 5xx status code
func (o *GetProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get project forbidden response a status code equal to that given
func (o *GetProjectForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectForbidden  %+v", 403, o.Payload)
}

func (o *GetProjectForbidden) String() string {
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectForbidden  %+v", 403, o.Payload)
}

func (o *GetProjectForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectNotFound creates a GetProjectNotFound with default headers values
func NewGetProjectNotFound() *GetProjectNotFound {
	return &GetProjectNotFound{}
}

/*
GetProjectNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetProjectNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get project not found response has a 2xx status code
func (o *GetProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project not found response has a 3xx status code
func (o *GetProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project not found response has a 4xx status code
func (o *GetProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project not found response has a 5xx status code
func (o *GetProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get project not found response a status code equal to that given
func (o *GetProjectNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectNotFound  %+v", 404, o.Payload)
}

func (o *GetProjectNotFound) String() string {
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectNotFound  %+v", 404, o.Payload)
}

func (o *GetProjectNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectInternalServerError creates a GetProjectInternalServerError with default headers values
func NewGetProjectInternalServerError() *GetProjectInternalServerError {
	return &GetProjectInternalServerError{}
}

/*
GetProjectInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetProjectInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get project internal server error response has a 2xx status code
func (o *GetProjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project internal server error response has a 3xx status code
func (o *GetProjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project internal server error response has a 4xx status code
func (o *GetProjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project internal server error response has a 5xx status code
func (o *GetProjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get project internal server error response a status code equal to that given
func (o *GetProjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetProjectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProjectInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{projectId}][%d] getProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProjectInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
