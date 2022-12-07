// Code generated by go-swagger; DO NOT EDIT.

package workspaces_permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bartsimp/talend-rest-go/models"
)

// GetWorkspacePermissionReader is a Reader for the GetWorkspacePermission structure.
type GetWorkspacePermissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkspacePermissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkspacePermissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkspacePermissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkspacePermissionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkspacePermissionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkspacePermissionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkspacePermissionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkspacePermissionOK creates a GetWorkspacePermissionOK with default headers values
func NewGetWorkspacePermissionOK() *GetWorkspacePermissionOK {
	return &GetWorkspacePermissionOK{}
}

/*
GetWorkspacePermissionOK describes a response with status code 200, with default header values.

OK
*/
type GetWorkspacePermissionOK struct {
	Payload *models.Permission
}

// IsSuccess returns true when this get workspace permission o k response has a 2xx status code
func (o *GetWorkspacePermissionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workspace permission o k response has a 3xx status code
func (o *GetWorkspacePermissionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workspace permission o k response has a 4xx status code
func (o *GetWorkspacePermissionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workspace permission o k response has a 5xx status code
func (o *GetWorkspacePermissionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workspace permission o k response a status code equal to that given
func (o *GetWorkspacePermissionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkspacePermissionOK) Error() string {
	return fmt.Sprintf("[GET /workspaces/{workspaceId}/users/{userId}/permissions][%d] getWorkspacePermissionOK  %+v", 200, o.Payload)
}

func (o *GetWorkspacePermissionOK) String() string {
	return fmt.Sprintf("[GET /workspaces/{workspaceId}/users/{userId}/permissions][%d] getWorkspacePermissionOK  %+v", 200, o.Payload)
}

func (o *GetWorkspacePermissionOK) GetPayload() *models.Permission {
	return o.Payload
}

func (o *GetWorkspacePermissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Permission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkspacePermissionBadRequest creates a GetWorkspacePermissionBadRequest with default headers values
func NewGetWorkspacePermissionBadRequest() *GetWorkspacePermissionBadRequest {
	return &GetWorkspacePermissionBadRequest{}
}

/*
GetWorkspacePermissionBadRequest describes a response with status code 400, with default header values.

Parameters not valid
*/
type GetWorkspacePermissionBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get workspace permission bad request response has a 2xx status code
func (o *GetWorkspacePermissionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workspace permission bad request response has a 3xx status code
func (o *GetWorkspacePermissionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workspace permission bad request response has a 4xx status code
func (o *GetWorkspacePermissionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workspace permission bad request response has a 5xx status code
func (o *GetWorkspacePermissionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workspace permission bad request response a status code equal to that given
func (o *GetWorkspacePermissionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWorkspacePermissionBadRequest) Error() string {
	return fmt.Sprintf("[GET /workspaces/{workspaceId}/users/{userId}/permissions][%d] getWorkspacePermissionBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkspacePermissionBadRequest) String() string {
	return fmt.Sprintf("[GET /workspaces/{workspaceId}/users/{userId}/permissions][%d] getWorkspacePermissionBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkspacePermissionBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWorkspacePermissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkspacePermissionUnauthorized creates a GetWorkspacePermissionUnauthorized with default headers values
func NewGetWorkspacePermissionUnauthorized() *GetWorkspacePermissionUnauthorized {
	return &GetWorkspacePermissionUnauthorized{}
}

/*
GetWorkspacePermissionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetWorkspacePermissionUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get workspace permission unauthorized response has a 2xx status code
func (o *GetWorkspacePermissionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workspace permission unauthorized response has a 3xx status code
func (o *GetWorkspacePermissionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workspace permission unauthorized response has a 4xx status code
func (o *GetWorkspacePermissionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workspace permission unauthorized response has a 5xx status code
func (o *GetWorkspacePermissionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workspace permission unauthorized response a status code equal to that given
func (o *GetWorkspacePermissionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWorkspacePermissionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /workspaces/{workspaceId}/users/{userId}/permissions][%d] getWorkspacePermissionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkspacePermissionUnauthorized) String() string {
	return fmt.Sprintf("[GET /workspaces/{workspaceId}/users/{userId}/permissions][%d] getWorkspacePermissionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkspacePermissionUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWorkspacePermissionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkspacePermissionForbidden creates a GetWorkspacePermissionForbidden with default headers values
func NewGetWorkspacePermissionForbidden() *GetWorkspacePermissionForbidden {
	return &GetWorkspacePermissionForbidden{}
}

/*
GetWorkspacePermissionForbidden describes a response with status code 403, with default header values.

Not authorized to change permission
*/
type GetWorkspacePermissionForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get workspace permission forbidden response has a 2xx status code
func (o *GetWorkspacePermissionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workspace permission forbidden response has a 3xx status code
func (o *GetWorkspacePermissionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workspace permission forbidden response has a 4xx status code
func (o *GetWorkspacePermissionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workspace permission forbidden response has a 5xx status code
func (o *GetWorkspacePermissionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workspace permission forbidden response a status code equal to that given
func (o *GetWorkspacePermissionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWorkspacePermissionForbidden) Error() string {
	return fmt.Sprintf("[GET /workspaces/{workspaceId}/users/{userId}/permissions][%d] getWorkspacePermissionForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkspacePermissionForbidden) String() string {
	return fmt.Sprintf("[GET /workspaces/{workspaceId}/users/{userId}/permissions][%d] getWorkspacePermissionForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkspacePermissionForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWorkspacePermissionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkspacePermissionNotFound creates a GetWorkspacePermissionNotFound with default headers values
func NewGetWorkspacePermissionNotFound() *GetWorkspacePermissionNotFound {
	return &GetWorkspacePermissionNotFound{}
}

/*
GetWorkspacePermissionNotFound describes a response with status code 404, with default header values.

Permission not found
*/
type GetWorkspacePermissionNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get workspace permission not found response has a 2xx status code
func (o *GetWorkspacePermissionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workspace permission not found response has a 3xx status code
func (o *GetWorkspacePermissionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workspace permission not found response has a 4xx status code
func (o *GetWorkspacePermissionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workspace permission not found response has a 5xx status code
func (o *GetWorkspacePermissionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workspace permission not found response a status code equal to that given
func (o *GetWorkspacePermissionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWorkspacePermissionNotFound) Error() string {
	return fmt.Sprintf("[GET /workspaces/{workspaceId}/users/{userId}/permissions][%d] getWorkspacePermissionNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkspacePermissionNotFound) String() string {
	return fmt.Sprintf("[GET /workspaces/{workspaceId}/users/{userId}/permissions][%d] getWorkspacePermissionNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkspacePermissionNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWorkspacePermissionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkspacePermissionInternalServerError creates a GetWorkspacePermissionInternalServerError with default headers values
func NewGetWorkspacePermissionInternalServerError() *GetWorkspacePermissionInternalServerError {
	return &GetWorkspacePermissionInternalServerError{}
}

/*
GetWorkspacePermissionInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetWorkspacePermissionInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get workspace permission internal server error response has a 2xx status code
func (o *GetWorkspacePermissionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workspace permission internal server error response has a 3xx status code
func (o *GetWorkspacePermissionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workspace permission internal server error response has a 4xx status code
func (o *GetWorkspacePermissionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workspace permission internal server error response has a 5xx status code
func (o *GetWorkspacePermissionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workspace permission internal server error response a status code equal to that given
func (o *GetWorkspacePermissionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWorkspacePermissionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /workspaces/{workspaceId}/users/{userId}/permissions][%d] getWorkspacePermissionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkspacePermissionInternalServerError) String() string {
	return fmt.Sprintf("[GET /workspaces/{workspaceId}/users/{userId}/permissions][%d] getWorkspacePermissionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkspacePermissionInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetWorkspacePermissionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}