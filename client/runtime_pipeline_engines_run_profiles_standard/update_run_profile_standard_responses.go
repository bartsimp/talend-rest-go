// Code generated by go-swagger; DO NOT EDIT.

package runtime_pipeline_engines_run_profiles_standard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bartsimp/talend-rest-go/models"
)

// UpdateRunProfileStandardReader is a Reader for the UpdateRunProfileStandard structure.
type UpdateRunProfileStandardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRunProfileStandardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateRunProfileStandardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRunProfileStandardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateRunProfileStandardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRunProfileStandardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRunProfileStandardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateRunProfileStandardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRunProfileStandardNoContent creates a UpdateRunProfileStandardNoContent with default headers values
func NewUpdateRunProfileStandardNoContent() *UpdateRunProfileStandardNoContent {
	return &UpdateRunProfileStandardNoContent{}
}

/*
UpdateRunProfileStandardNoContent describes a response with status code 204, with default header values.

No Content
*/
type UpdateRunProfileStandardNoContent struct {
}

// IsSuccess returns true when this update run profile standard no content response has a 2xx status code
func (o *UpdateRunProfileStandardNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update run profile standard no content response has a 3xx status code
func (o *UpdateRunProfileStandardNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update run profile standard no content response has a 4xx status code
func (o *UpdateRunProfileStandardNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this update run profile standard no content response has a 5xx status code
func (o *UpdateRunProfileStandardNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this update run profile standard no content response a status code equal to that given
func (o *UpdateRunProfileStandardNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UpdateRunProfileStandardNoContent) Error() string {
	return fmt.Sprintf("[PUT /runtimes/pipeline-engines/{engineId}/run-profiles/standard/{runProfileId}][%d] updateRunProfileStandardNoContent ", 204)
}

func (o *UpdateRunProfileStandardNoContent) String() string {
	return fmt.Sprintf("[PUT /runtimes/pipeline-engines/{engineId}/run-profiles/standard/{runProfileId}][%d] updateRunProfileStandardNoContent ", 204)
}

func (o *UpdateRunProfileStandardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateRunProfileStandardBadRequest creates a UpdateRunProfileStandardBadRequest with default headers values
func NewUpdateRunProfileStandardBadRequest() *UpdateRunProfileStandardBadRequest {
	return &UpdateRunProfileStandardBadRequest{}
}

/*
UpdateRunProfileStandardBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateRunProfileStandardBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update run profile standard bad request response has a 2xx status code
func (o *UpdateRunProfileStandardBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update run profile standard bad request response has a 3xx status code
func (o *UpdateRunProfileStandardBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update run profile standard bad request response has a 4xx status code
func (o *UpdateRunProfileStandardBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update run profile standard bad request response has a 5xx status code
func (o *UpdateRunProfileStandardBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update run profile standard bad request response a status code equal to that given
func (o *UpdateRunProfileStandardBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateRunProfileStandardBadRequest) Error() string {
	return fmt.Sprintf("[PUT /runtimes/pipeline-engines/{engineId}/run-profiles/standard/{runProfileId}][%d] updateRunProfileStandardBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRunProfileStandardBadRequest) String() string {
	return fmt.Sprintf("[PUT /runtimes/pipeline-engines/{engineId}/run-profiles/standard/{runProfileId}][%d] updateRunProfileStandardBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRunProfileStandardBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateRunProfileStandardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunProfileStandardUnauthorized creates a UpdateRunProfileStandardUnauthorized with default headers values
func NewUpdateRunProfileStandardUnauthorized() *UpdateRunProfileStandardUnauthorized {
	return &UpdateRunProfileStandardUnauthorized{}
}

/*
UpdateRunProfileStandardUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateRunProfileStandardUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update run profile standard unauthorized response has a 2xx status code
func (o *UpdateRunProfileStandardUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update run profile standard unauthorized response has a 3xx status code
func (o *UpdateRunProfileStandardUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update run profile standard unauthorized response has a 4xx status code
func (o *UpdateRunProfileStandardUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update run profile standard unauthorized response has a 5xx status code
func (o *UpdateRunProfileStandardUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update run profile standard unauthorized response a status code equal to that given
func (o *UpdateRunProfileStandardUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateRunProfileStandardUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /runtimes/pipeline-engines/{engineId}/run-profiles/standard/{runProfileId}][%d] updateRunProfileStandardUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateRunProfileStandardUnauthorized) String() string {
	return fmt.Sprintf("[PUT /runtimes/pipeline-engines/{engineId}/run-profiles/standard/{runProfileId}][%d] updateRunProfileStandardUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateRunProfileStandardUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateRunProfileStandardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunProfileStandardForbidden creates a UpdateRunProfileStandardForbidden with default headers values
func NewUpdateRunProfileStandardForbidden() *UpdateRunProfileStandardForbidden {
	return &UpdateRunProfileStandardForbidden{}
}

/*
UpdateRunProfileStandardForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateRunProfileStandardForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update run profile standard forbidden response has a 2xx status code
func (o *UpdateRunProfileStandardForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update run profile standard forbidden response has a 3xx status code
func (o *UpdateRunProfileStandardForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update run profile standard forbidden response has a 4xx status code
func (o *UpdateRunProfileStandardForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update run profile standard forbidden response has a 5xx status code
func (o *UpdateRunProfileStandardForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update run profile standard forbidden response a status code equal to that given
func (o *UpdateRunProfileStandardForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateRunProfileStandardForbidden) Error() string {
	return fmt.Sprintf("[PUT /runtimes/pipeline-engines/{engineId}/run-profiles/standard/{runProfileId}][%d] updateRunProfileStandardForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRunProfileStandardForbidden) String() string {
	return fmt.Sprintf("[PUT /runtimes/pipeline-engines/{engineId}/run-profiles/standard/{runProfileId}][%d] updateRunProfileStandardForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRunProfileStandardForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateRunProfileStandardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunProfileStandardNotFound creates a UpdateRunProfileStandardNotFound with default headers values
func NewUpdateRunProfileStandardNotFound() *UpdateRunProfileStandardNotFound {
	return &UpdateRunProfileStandardNotFound{}
}

/*
UpdateRunProfileStandardNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateRunProfileStandardNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update run profile standard not found response has a 2xx status code
func (o *UpdateRunProfileStandardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update run profile standard not found response has a 3xx status code
func (o *UpdateRunProfileStandardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update run profile standard not found response has a 4xx status code
func (o *UpdateRunProfileStandardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update run profile standard not found response has a 5xx status code
func (o *UpdateRunProfileStandardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update run profile standard not found response a status code equal to that given
func (o *UpdateRunProfileStandardNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateRunProfileStandardNotFound) Error() string {
	return fmt.Sprintf("[PUT /runtimes/pipeline-engines/{engineId}/run-profiles/standard/{runProfileId}][%d] updateRunProfileStandardNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRunProfileStandardNotFound) String() string {
	return fmt.Sprintf("[PUT /runtimes/pipeline-engines/{engineId}/run-profiles/standard/{runProfileId}][%d] updateRunProfileStandardNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRunProfileStandardNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateRunProfileStandardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRunProfileStandardInternalServerError creates a UpdateRunProfileStandardInternalServerError with default headers values
func NewUpdateRunProfileStandardInternalServerError() *UpdateRunProfileStandardInternalServerError {
	return &UpdateRunProfileStandardInternalServerError{}
}

/*
UpdateRunProfileStandardInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateRunProfileStandardInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this update run profile standard internal server error response has a 2xx status code
func (o *UpdateRunProfileStandardInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update run profile standard internal server error response has a 3xx status code
func (o *UpdateRunProfileStandardInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update run profile standard internal server error response has a 4xx status code
func (o *UpdateRunProfileStandardInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update run profile standard internal server error response has a 5xx status code
func (o *UpdateRunProfileStandardInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update run profile standard internal server error response a status code equal to that given
func (o *UpdateRunProfileStandardInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateRunProfileStandardInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /runtimes/pipeline-engines/{engineId}/run-profiles/standard/{runProfileId}][%d] updateRunProfileStandardInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRunProfileStandardInternalServerError) String() string {
	return fmt.Sprintf("[PUT /runtimes/pipeline-engines/{engineId}/run-profiles/standard/{runProfileId}][%d] updateRunProfileStandardInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRunProfileStandardInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateRunProfileStandardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}