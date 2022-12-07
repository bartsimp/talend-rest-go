// Code generated by go-swagger; DO NOT EDIT.

package runtime_pipeline_engines_run_profiles_data_integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bartsimp/talend-rest-go/models"
)

// CreateDataIntegrationRunProfileReader is a Reader for the CreateDataIntegrationRunProfile structure.
type CreateDataIntegrationRunProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDataIntegrationRunProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDataIntegrationRunProfileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDataIntegrationRunProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateDataIntegrationRunProfileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateDataIntegrationRunProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateDataIntegrationRunProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateDataIntegrationRunProfileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDataIntegrationRunProfileCreated creates a CreateDataIntegrationRunProfileCreated with default headers values
func NewCreateDataIntegrationRunProfileCreated() *CreateDataIntegrationRunProfileCreated {
	return &CreateDataIntegrationRunProfileCreated{}
}

/*
CreateDataIntegrationRunProfileCreated describes a response with status code 201, with default header values.

Created
*/
type CreateDataIntegrationRunProfileCreated struct {
}

// IsSuccess returns true when this create data integration run profile created response has a 2xx status code
func (o *CreateDataIntegrationRunProfileCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create data integration run profile created response has a 3xx status code
func (o *CreateDataIntegrationRunProfileCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create data integration run profile created response has a 4xx status code
func (o *CreateDataIntegrationRunProfileCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create data integration run profile created response has a 5xx status code
func (o *CreateDataIntegrationRunProfileCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create data integration run profile created response a status code equal to that given
func (o *CreateDataIntegrationRunProfileCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateDataIntegrationRunProfileCreated) Error() string {
	return fmt.Sprintf("[POST /runtimes/pipeline-engines/{engineId}/run-profiles/data-integration][%d] createDataIntegrationRunProfileCreated ", 201)
}

func (o *CreateDataIntegrationRunProfileCreated) String() string {
	return fmt.Sprintf("[POST /runtimes/pipeline-engines/{engineId}/run-profiles/data-integration][%d] createDataIntegrationRunProfileCreated ", 201)
}

func (o *CreateDataIntegrationRunProfileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDataIntegrationRunProfileBadRequest creates a CreateDataIntegrationRunProfileBadRequest with default headers values
func NewCreateDataIntegrationRunProfileBadRequest() *CreateDataIntegrationRunProfileBadRequest {
	return &CreateDataIntegrationRunProfileBadRequest{}
}

/*
CreateDataIntegrationRunProfileBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateDataIntegrationRunProfileBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create data integration run profile bad request response has a 2xx status code
func (o *CreateDataIntegrationRunProfileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create data integration run profile bad request response has a 3xx status code
func (o *CreateDataIntegrationRunProfileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create data integration run profile bad request response has a 4xx status code
func (o *CreateDataIntegrationRunProfileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create data integration run profile bad request response has a 5xx status code
func (o *CreateDataIntegrationRunProfileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create data integration run profile bad request response a status code equal to that given
func (o *CreateDataIntegrationRunProfileBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateDataIntegrationRunProfileBadRequest) Error() string {
	return fmt.Sprintf("[POST /runtimes/pipeline-engines/{engineId}/run-profiles/data-integration][%d] createDataIntegrationRunProfileBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDataIntegrationRunProfileBadRequest) String() string {
	return fmt.Sprintf("[POST /runtimes/pipeline-engines/{engineId}/run-profiles/data-integration][%d] createDataIntegrationRunProfileBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDataIntegrationRunProfileBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateDataIntegrationRunProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDataIntegrationRunProfileUnauthorized creates a CreateDataIntegrationRunProfileUnauthorized with default headers values
func NewCreateDataIntegrationRunProfileUnauthorized() *CreateDataIntegrationRunProfileUnauthorized {
	return &CreateDataIntegrationRunProfileUnauthorized{}
}

/*
CreateDataIntegrationRunProfileUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateDataIntegrationRunProfileUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create data integration run profile unauthorized response has a 2xx status code
func (o *CreateDataIntegrationRunProfileUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create data integration run profile unauthorized response has a 3xx status code
func (o *CreateDataIntegrationRunProfileUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create data integration run profile unauthorized response has a 4xx status code
func (o *CreateDataIntegrationRunProfileUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create data integration run profile unauthorized response has a 5xx status code
func (o *CreateDataIntegrationRunProfileUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create data integration run profile unauthorized response a status code equal to that given
func (o *CreateDataIntegrationRunProfileUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateDataIntegrationRunProfileUnauthorized) Error() string {
	return fmt.Sprintf("[POST /runtimes/pipeline-engines/{engineId}/run-profiles/data-integration][%d] createDataIntegrationRunProfileUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateDataIntegrationRunProfileUnauthorized) String() string {
	return fmt.Sprintf("[POST /runtimes/pipeline-engines/{engineId}/run-profiles/data-integration][%d] createDataIntegrationRunProfileUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateDataIntegrationRunProfileUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateDataIntegrationRunProfileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDataIntegrationRunProfileForbidden creates a CreateDataIntegrationRunProfileForbidden with default headers values
func NewCreateDataIntegrationRunProfileForbidden() *CreateDataIntegrationRunProfileForbidden {
	return &CreateDataIntegrationRunProfileForbidden{}
}

/*
CreateDataIntegrationRunProfileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateDataIntegrationRunProfileForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create data integration run profile forbidden response has a 2xx status code
func (o *CreateDataIntegrationRunProfileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create data integration run profile forbidden response has a 3xx status code
func (o *CreateDataIntegrationRunProfileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create data integration run profile forbidden response has a 4xx status code
func (o *CreateDataIntegrationRunProfileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create data integration run profile forbidden response has a 5xx status code
func (o *CreateDataIntegrationRunProfileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create data integration run profile forbidden response a status code equal to that given
func (o *CreateDataIntegrationRunProfileForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateDataIntegrationRunProfileForbidden) Error() string {
	return fmt.Sprintf("[POST /runtimes/pipeline-engines/{engineId}/run-profiles/data-integration][%d] createDataIntegrationRunProfileForbidden  %+v", 403, o.Payload)
}

func (o *CreateDataIntegrationRunProfileForbidden) String() string {
	return fmt.Sprintf("[POST /runtimes/pipeline-engines/{engineId}/run-profiles/data-integration][%d] createDataIntegrationRunProfileForbidden  %+v", 403, o.Payload)
}

func (o *CreateDataIntegrationRunProfileForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateDataIntegrationRunProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDataIntegrationRunProfileNotFound creates a CreateDataIntegrationRunProfileNotFound with default headers values
func NewCreateDataIntegrationRunProfileNotFound() *CreateDataIntegrationRunProfileNotFound {
	return &CreateDataIntegrationRunProfileNotFound{}
}

/*
CreateDataIntegrationRunProfileNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateDataIntegrationRunProfileNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create data integration run profile not found response has a 2xx status code
func (o *CreateDataIntegrationRunProfileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create data integration run profile not found response has a 3xx status code
func (o *CreateDataIntegrationRunProfileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create data integration run profile not found response has a 4xx status code
func (o *CreateDataIntegrationRunProfileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create data integration run profile not found response has a 5xx status code
func (o *CreateDataIntegrationRunProfileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create data integration run profile not found response a status code equal to that given
func (o *CreateDataIntegrationRunProfileNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateDataIntegrationRunProfileNotFound) Error() string {
	return fmt.Sprintf("[POST /runtimes/pipeline-engines/{engineId}/run-profiles/data-integration][%d] createDataIntegrationRunProfileNotFound  %+v", 404, o.Payload)
}

func (o *CreateDataIntegrationRunProfileNotFound) String() string {
	return fmt.Sprintf("[POST /runtimes/pipeline-engines/{engineId}/run-profiles/data-integration][%d] createDataIntegrationRunProfileNotFound  %+v", 404, o.Payload)
}

func (o *CreateDataIntegrationRunProfileNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateDataIntegrationRunProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDataIntegrationRunProfileInternalServerError creates a CreateDataIntegrationRunProfileInternalServerError with default headers values
func NewCreateDataIntegrationRunProfileInternalServerError() *CreateDataIntegrationRunProfileInternalServerError {
	return &CreateDataIntegrationRunProfileInternalServerError{}
}

/*
CreateDataIntegrationRunProfileInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateDataIntegrationRunProfileInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this create data integration run profile internal server error response has a 2xx status code
func (o *CreateDataIntegrationRunProfileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create data integration run profile internal server error response has a 3xx status code
func (o *CreateDataIntegrationRunProfileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create data integration run profile internal server error response has a 4xx status code
func (o *CreateDataIntegrationRunProfileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create data integration run profile internal server error response has a 5xx status code
func (o *CreateDataIntegrationRunProfileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create data integration run profile internal server error response a status code equal to that given
func (o *CreateDataIntegrationRunProfileInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateDataIntegrationRunProfileInternalServerError) Error() string {
	return fmt.Sprintf("[POST /runtimes/pipeline-engines/{engineId}/run-profiles/data-integration][%d] createDataIntegrationRunProfileInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateDataIntegrationRunProfileInternalServerError) String() string {
	return fmt.Sprintf("[POST /runtimes/pipeline-engines/{engineId}/run-profiles/data-integration][%d] createDataIntegrationRunProfileInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateDataIntegrationRunProfileInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateDataIntegrationRunProfileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
