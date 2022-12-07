// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bartsimp/talend-rest-go/models"
)

// CreateEnvironmentReader is a Reader for the CreateEnvironment structure.
type CreateEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateEnvironmentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateEnvironmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateEnvironmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateEnvironmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateEnvironmentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateEnvironmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateEnvironmentCreated creates a CreateEnvironmentCreated with default headers values
func NewCreateEnvironmentCreated() *CreateEnvironmentCreated {
	return &CreateEnvironmentCreated{}
}

/*
CreateEnvironmentCreated describes a response with status code 201, with default header values.

Environment successfully created
*/
type CreateEnvironmentCreated struct {
	Payload *models.EnvironmentInfo
}

// IsSuccess returns true when this create environment created response has a 2xx status code
func (o *CreateEnvironmentCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create environment created response has a 3xx status code
func (o *CreateEnvironmentCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create environment created response has a 4xx status code
func (o *CreateEnvironmentCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create environment created response has a 5xx status code
func (o *CreateEnvironmentCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create environment created response a status code equal to that given
func (o *CreateEnvironmentCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateEnvironmentCreated) Error() string {
	return fmt.Sprintf("[POST /environments][%d] createEnvironmentCreated  %+v", 201, o.Payload)
}

func (o *CreateEnvironmentCreated) String() string {
	return fmt.Sprintf("[POST /environments][%d] createEnvironmentCreated  %+v", 201, o.Payload)
}

func (o *CreateEnvironmentCreated) GetPayload() *models.EnvironmentInfo {
	return o.Payload
}

func (o *CreateEnvironmentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvironmentInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnvironmentBadRequest creates a CreateEnvironmentBadRequest with default headers values
func NewCreateEnvironmentBadRequest() *CreateEnvironmentBadRequest {
	return &CreateEnvironmentBadRequest{}
}

/*
CreateEnvironmentBadRequest describes a response with status code 400, with default header values.

Parameters not valid
*/
type CreateEnvironmentBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create environment bad request response has a 2xx status code
func (o *CreateEnvironmentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create environment bad request response has a 3xx status code
func (o *CreateEnvironmentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create environment bad request response has a 4xx status code
func (o *CreateEnvironmentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create environment bad request response has a 5xx status code
func (o *CreateEnvironmentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create environment bad request response a status code equal to that given
func (o *CreateEnvironmentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateEnvironmentBadRequest) Error() string {
	return fmt.Sprintf("[POST /environments][%d] createEnvironmentBadRequest  %+v", 400, o.Payload)
}

func (o *CreateEnvironmentBadRequest) String() string {
	return fmt.Sprintf("[POST /environments][%d] createEnvironmentBadRequest  %+v", 400, o.Payload)
}

func (o *CreateEnvironmentBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateEnvironmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnvironmentUnauthorized creates a CreateEnvironmentUnauthorized with default headers values
func NewCreateEnvironmentUnauthorized() *CreateEnvironmentUnauthorized {
	return &CreateEnvironmentUnauthorized{}
}

/*
CreateEnvironmentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateEnvironmentUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create environment unauthorized response has a 2xx status code
func (o *CreateEnvironmentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create environment unauthorized response has a 3xx status code
func (o *CreateEnvironmentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create environment unauthorized response has a 4xx status code
func (o *CreateEnvironmentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create environment unauthorized response has a 5xx status code
func (o *CreateEnvironmentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create environment unauthorized response a status code equal to that given
func (o *CreateEnvironmentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateEnvironmentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /environments][%d] createEnvironmentUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateEnvironmentUnauthorized) String() string {
	return fmt.Sprintf("[POST /environments][%d] createEnvironmentUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateEnvironmentUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateEnvironmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnvironmentForbidden creates a CreateEnvironmentForbidden with default headers values
func NewCreateEnvironmentForbidden() *CreateEnvironmentForbidden {
	return &CreateEnvironmentForbidden{}
}

/*
CreateEnvironmentForbidden describes a response with status code 403, with default header values.

Not authorized to create the environment
*/
type CreateEnvironmentForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create environment forbidden response has a 2xx status code
func (o *CreateEnvironmentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create environment forbidden response has a 3xx status code
func (o *CreateEnvironmentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create environment forbidden response has a 4xx status code
func (o *CreateEnvironmentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create environment forbidden response has a 5xx status code
func (o *CreateEnvironmentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create environment forbidden response a status code equal to that given
func (o *CreateEnvironmentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateEnvironmentForbidden) Error() string {
	return fmt.Sprintf("[POST /environments][%d] createEnvironmentForbidden  %+v", 403, o.Payload)
}

func (o *CreateEnvironmentForbidden) String() string {
	return fmt.Sprintf("[POST /environments][%d] createEnvironmentForbidden  %+v", 403, o.Payload)
}

func (o *CreateEnvironmentForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateEnvironmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnvironmentConflict creates a CreateEnvironmentConflict with default headers values
func NewCreateEnvironmentConflict() *CreateEnvironmentConflict {
	return &CreateEnvironmentConflict{}
}

/*
CreateEnvironmentConflict describes a response with status code 409, with default header values.

Environment name should be unique
*/
type CreateEnvironmentConflict struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create environment conflict response has a 2xx status code
func (o *CreateEnvironmentConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create environment conflict response has a 3xx status code
func (o *CreateEnvironmentConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create environment conflict response has a 4xx status code
func (o *CreateEnvironmentConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create environment conflict response has a 5xx status code
func (o *CreateEnvironmentConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create environment conflict response a status code equal to that given
func (o *CreateEnvironmentConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateEnvironmentConflict) Error() string {
	return fmt.Sprintf("[POST /environments][%d] createEnvironmentConflict  %+v", 409, o.Payload)
}

func (o *CreateEnvironmentConflict) String() string {
	return fmt.Sprintf("[POST /environments][%d] createEnvironmentConflict  %+v", 409, o.Payload)
}

func (o *CreateEnvironmentConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateEnvironmentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnvironmentInternalServerError creates a CreateEnvironmentInternalServerError with default headers values
func NewCreateEnvironmentInternalServerError() *CreateEnvironmentInternalServerError {
	return &CreateEnvironmentInternalServerError{}
}

/*
CreateEnvironmentInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateEnvironmentInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create environment internal server error response has a 2xx status code
func (o *CreateEnvironmentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create environment internal server error response has a 3xx status code
func (o *CreateEnvironmentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create environment internal server error response has a 4xx status code
func (o *CreateEnvironmentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create environment internal server error response has a 5xx status code
func (o *CreateEnvironmentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create environment internal server error response a status code equal to that given
func (o *CreateEnvironmentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateEnvironmentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /environments][%d] createEnvironmentInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateEnvironmentInternalServerError) String() string {
	return fmt.Sprintf("[POST /environments][%d] createEnvironmentInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateEnvironmentInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateEnvironmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
