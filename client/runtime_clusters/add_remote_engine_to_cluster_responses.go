// Code generated by go-swagger; DO NOT EDIT.

package runtime_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bartsimp/talend-rest-go/models"
)

// AddRemoteEngineToClusterReader is a Reader for the AddRemoteEngineToCluster structure.
type AddRemoteEngineToClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRemoteEngineToClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAddRemoteEngineToClusterNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddRemoteEngineToClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddRemoteEngineToClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddRemoteEngineToClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddRemoteEngineToClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddRemoteEngineToClusterConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddRemoteEngineToClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddRemoteEngineToClusterNoContent creates a AddRemoteEngineToClusterNoContent with default headers values
func NewAddRemoteEngineToClusterNoContent() *AddRemoteEngineToClusterNoContent {
	return &AddRemoteEngineToClusterNoContent{}
}

/*
AddRemoteEngineToClusterNoContent describes a response with status code 204, with default header values.

No Content
*/
type AddRemoteEngineToClusterNoContent struct {
}

// IsSuccess returns true when this add remote engine to cluster no content response has a 2xx status code
func (o *AddRemoteEngineToClusterNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add remote engine to cluster no content response has a 3xx status code
func (o *AddRemoteEngineToClusterNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add remote engine to cluster no content response has a 4xx status code
func (o *AddRemoteEngineToClusterNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this add remote engine to cluster no content response has a 5xx status code
func (o *AddRemoteEngineToClusterNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this add remote engine to cluster no content response a status code equal to that given
func (o *AddRemoteEngineToClusterNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *AddRemoteEngineToClusterNoContent) Error() string {
	return fmt.Sprintf("[PUT /runtimes/remote-engine-clusters/{clusterId}/engines/{engineId}][%d] addRemoteEngineToClusterNoContent ", 204)
}

func (o *AddRemoteEngineToClusterNoContent) String() string {
	return fmt.Sprintf("[PUT /runtimes/remote-engine-clusters/{clusterId}/engines/{engineId}][%d] addRemoteEngineToClusterNoContent ", 204)
}

func (o *AddRemoteEngineToClusterNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddRemoteEngineToClusterBadRequest creates a AddRemoteEngineToClusterBadRequest with default headers values
func NewAddRemoteEngineToClusterBadRequest() *AddRemoteEngineToClusterBadRequest {
	return &AddRemoteEngineToClusterBadRequest{}
}

/*
AddRemoteEngineToClusterBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AddRemoteEngineToClusterBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this add remote engine to cluster bad request response has a 2xx status code
func (o *AddRemoteEngineToClusterBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add remote engine to cluster bad request response has a 3xx status code
func (o *AddRemoteEngineToClusterBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add remote engine to cluster bad request response has a 4xx status code
func (o *AddRemoteEngineToClusterBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add remote engine to cluster bad request response has a 5xx status code
func (o *AddRemoteEngineToClusterBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add remote engine to cluster bad request response a status code equal to that given
func (o *AddRemoteEngineToClusterBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AddRemoteEngineToClusterBadRequest) Error() string {
	return fmt.Sprintf("[PUT /runtimes/remote-engine-clusters/{clusterId}/engines/{engineId}][%d] addRemoteEngineToClusterBadRequest  %+v", 400, o.Payload)
}

func (o *AddRemoteEngineToClusterBadRequest) String() string {
	return fmt.Sprintf("[PUT /runtimes/remote-engine-clusters/{clusterId}/engines/{engineId}][%d] addRemoteEngineToClusterBadRequest  %+v", 400, o.Payload)
}

func (o *AddRemoteEngineToClusterBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddRemoteEngineToClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddRemoteEngineToClusterUnauthorized creates a AddRemoteEngineToClusterUnauthorized with default headers values
func NewAddRemoteEngineToClusterUnauthorized() *AddRemoteEngineToClusterUnauthorized {
	return &AddRemoteEngineToClusterUnauthorized{}
}

/*
AddRemoteEngineToClusterUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AddRemoteEngineToClusterUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this add remote engine to cluster unauthorized response has a 2xx status code
func (o *AddRemoteEngineToClusterUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add remote engine to cluster unauthorized response has a 3xx status code
func (o *AddRemoteEngineToClusterUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add remote engine to cluster unauthorized response has a 4xx status code
func (o *AddRemoteEngineToClusterUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add remote engine to cluster unauthorized response has a 5xx status code
func (o *AddRemoteEngineToClusterUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add remote engine to cluster unauthorized response a status code equal to that given
func (o *AddRemoteEngineToClusterUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AddRemoteEngineToClusterUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /runtimes/remote-engine-clusters/{clusterId}/engines/{engineId}][%d] addRemoteEngineToClusterUnauthorized  %+v", 401, o.Payload)
}

func (o *AddRemoteEngineToClusterUnauthorized) String() string {
	return fmt.Sprintf("[PUT /runtimes/remote-engine-clusters/{clusterId}/engines/{engineId}][%d] addRemoteEngineToClusterUnauthorized  %+v", 401, o.Payload)
}

func (o *AddRemoteEngineToClusterUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddRemoteEngineToClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddRemoteEngineToClusterForbidden creates a AddRemoteEngineToClusterForbidden with default headers values
func NewAddRemoteEngineToClusterForbidden() *AddRemoteEngineToClusterForbidden {
	return &AddRemoteEngineToClusterForbidden{}
}

/*
AddRemoteEngineToClusterForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AddRemoteEngineToClusterForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this add remote engine to cluster forbidden response has a 2xx status code
func (o *AddRemoteEngineToClusterForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add remote engine to cluster forbidden response has a 3xx status code
func (o *AddRemoteEngineToClusterForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add remote engine to cluster forbidden response has a 4xx status code
func (o *AddRemoteEngineToClusterForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add remote engine to cluster forbidden response has a 5xx status code
func (o *AddRemoteEngineToClusterForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add remote engine to cluster forbidden response a status code equal to that given
func (o *AddRemoteEngineToClusterForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AddRemoteEngineToClusterForbidden) Error() string {
	return fmt.Sprintf("[PUT /runtimes/remote-engine-clusters/{clusterId}/engines/{engineId}][%d] addRemoteEngineToClusterForbidden  %+v", 403, o.Payload)
}

func (o *AddRemoteEngineToClusterForbidden) String() string {
	return fmt.Sprintf("[PUT /runtimes/remote-engine-clusters/{clusterId}/engines/{engineId}][%d] addRemoteEngineToClusterForbidden  %+v", 403, o.Payload)
}

func (o *AddRemoteEngineToClusterForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddRemoteEngineToClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddRemoteEngineToClusterNotFound creates a AddRemoteEngineToClusterNotFound with default headers values
func NewAddRemoteEngineToClusterNotFound() *AddRemoteEngineToClusterNotFound {
	return &AddRemoteEngineToClusterNotFound{}
}

/*
AddRemoteEngineToClusterNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AddRemoteEngineToClusterNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this add remote engine to cluster not found response has a 2xx status code
func (o *AddRemoteEngineToClusterNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add remote engine to cluster not found response has a 3xx status code
func (o *AddRemoteEngineToClusterNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add remote engine to cluster not found response has a 4xx status code
func (o *AddRemoteEngineToClusterNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add remote engine to cluster not found response has a 5xx status code
func (o *AddRemoteEngineToClusterNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add remote engine to cluster not found response a status code equal to that given
func (o *AddRemoteEngineToClusterNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AddRemoteEngineToClusterNotFound) Error() string {
	return fmt.Sprintf("[PUT /runtimes/remote-engine-clusters/{clusterId}/engines/{engineId}][%d] addRemoteEngineToClusterNotFound  %+v", 404, o.Payload)
}

func (o *AddRemoteEngineToClusterNotFound) String() string {
	return fmt.Sprintf("[PUT /runtimes/remote-engine-clusters/{clusterId}/engines/{engineId}][%d] addRemoteEngineToClusterNotFound  %+v", 404, o.Payload)
}

func (o *AddRemoteEngineToClusterNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddRemoteEngineToClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddRemoteEngineToClusterConflict creates a AddRemoteEngineToClusterConflict with default headers values
func NewAddRemoteEngineToClusterConflict() *AddRemoteEngineToClusterConflict {
	return &AddRemoteEngineToClusterConflict{}
}

/*
AddRemoteEngineToClusterConflict describes a response with status code 409, with default header values.

Conflict
*/
type AddRemoteEngineToClusterConflict struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this add remote engine to cluster conflict response has a 2xx status code
func (o *AddRemoteEngineToClusterConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add remote engine to cluster conflict response has a 3xx status code
func (o *AddRemoteEngineToClusterConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add remote engine to cluster conflict response has a 4xx status code
func (o *AddRemoteEngineToClusterConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this add remote engine to cluster conflict response has a 5xx status code
func (o *AddRemoteEngineToClusterConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this add remote engine to cluster conflict response a status code equal to that given
func (o *AddRemoteEngineToClusterConflict) IsCode(code int) bool {
	return code == 409
}

func (o *AddRemoteEngineToClusterConflict) Error() string {
	return fmt.Sprintf("[PUT /runtimes/remote-engine-clusters/{clusterId}/engines/{engineId}][%d] addRemoteEngineToClusterConflict  %+v", 409, o.Payload)
}

func (o *AddRemoteEngineToClusterConflict) String() string {
	return fmt.Sprintf("[PUT /runtimes/remote-engine-clusters/{clusterId}/engines/{engineId}][%d] addRemoteEngineToClusterConflict  %+v", 409, o.Payload)
}

func (o *AddRemoteEngineToClusterConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddRemoteEngineToClusterConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddRemoteEngineToClusterInternalServerError creates a AddRemoteEngineToClusterInternalServerError with default headers values
func NewAddRemoteEngineToClusterInternalServerError() *AddRemoteEngineToClusterInternalServerError {
	return &AddRemoteEngineToClusterInternalServerError{}
}

/*
AddRemoteEngineToClusterInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type AddRemoteEngineToClusterInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this add remote engine to cluster internal server error response has a 2xx status code
func (o *AddRemoteEngineToClusterInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add remote engine to cluster internal server error response has a 3xx status code
func (o *AddRemoteEngineToClusterInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add remote engine to cluster internal server error response has a 4xx status code
func (o *AddRemoteEngineToClusterInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add remote engine to cluster internal server error response has a 5xx status code
func (o *AddRemoteEngineToClusterInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add remote engine to cluster internal server error response a status code equal to that given
func (o *AddRemoteEngineToClusterInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AddRemoteEngineToClusterInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /runtimes/remote-engine-clusters/{clusterId}/engines/{engineId}][%d] addRemoteEngineToClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *AddRemoteEngineToClusterInternalServerError) String() string {
	return fmt.Sprintf("[PUT /runtimes/remote-engine-clusters/{clusterId}/engines/{engineId}][%d] addRemoteEngineToClusterInternalServerError  %+v", 500, o.Payload)
}

func (o *AddRemoteEngineToClusterInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddRemoteEngineToClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
