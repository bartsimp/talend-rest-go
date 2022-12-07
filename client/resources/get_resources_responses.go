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

// GetResourcesReader is a Reader for the GetResources structure.
type GetResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetResourcesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetResourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetResourcesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetResourcesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResourcesOK creates a GetResourcesOK with default headers values
func NewGetResourcesOK() *GetResourcesOK {
	return &GetResourcesOK{}
}

/*
GetResourcesOK describes a response with status code 200, with default header values.

OK
*/
type GetResourcesOK struct {
	Payload *models.Page
}

// IsSuccess returns true when this get resources o k response has a 2xx status code
func (o *GetResourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get resources o k response has a 3xx status code
func (o *GetResourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resources o k response has a 4xx status code
func (o *GetResourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resources o k response has a 5xx status code
func (o *GetResourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get resources o k response a status code equal to that given
func (o *GetResourcesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetResourcesOK) Error() string {
	return fmt.Sprintf("[GET /resources][%d] getResourcesOK  %+v", 200, o.Payload)
}

func (o *GetResourcesOK) String() string {
	return fmt.Sprintf("[GET /resources][%d] getResourcesOK  %+v", 200, o.Payload)
}

func (o *GetResourcesOK) GetPayload() *models.Page {
	return o.Payload
}

func (o *GetResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Page)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourcesBadRequest creates a GetResourcesBadRequest with default headers values
func NewGetResourcesBadRequest() *GetResourcesBadRequest {
	return &GetResourcesBadRequest{}
}

/*
GetResourcesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetResourcesBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get resources bad request response has a 2xx status code
func (o *GetResourcesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resources bad request response has a 3xx status code
func (o *GetResourcesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resources bad request response has a 4xx status code
func (o *GetResourcesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resources bad request response has a 5xx status code
func (o *GetResourcesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get resources bad request response a status code equal to that given
func (o *GetResourcesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetResourcesBadRequest) Error() string {
	return fmt.Sprintf("[GET /resources][%d] getResourcesBadRequest  %+v", 400, o.Payload)
}

func (o *GetResourcesBadRequest) String() string {
	return fmt.Sprintf("[GET /resources][%d] getResourcesBadRequest  %+v", 400, o.Payload)
}

func (o *GetResourcesBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetResourcesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourcesUnauthorized creates a GetResourcesUnauthorized with default headers values
func NewGetResourcesUnauthorized() *GetResourcesUnauthorized {
	return &GetResourcesUnauthorized{}
}

/*
GetResourcesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetResourcesUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get resources unauthorized response has a 2xx status code
func (o *GetResourcesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resources unauthorized response has a 3xx status code
func (o *GetResourcesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resources unauthorized response has a 4xx status code
func (o *GetResourcesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resources unauthorized response has a 5xx status code
func (o *GetResourcesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get resources unauthorized response a status code equal to that given
func (o *GetResourcesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetResourcesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /resources][%d] getResourcesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetResourcesUnauthorized) String() string {
	return fmt.Sprintf("[GET /resources][%d] getResourcesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetResourcesUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetResourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourcesForbidden creates a GetResourcesForbidden with default headers values
func NewGetResourcesForbidden() *GetResourcesForbidden {
	return &GetResourcesForbidden{}
}

/*
GetResourcesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetResourcesForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get resources forbidden response has a 2xx status code
func (o *GetResourcesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resources forbidden response has a 3xx status code
func (o *GetResourcesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resources forbidden response has a 4xx status code
func (o *GetResourcesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resources forbidden response has a 5xx status code
func (o *GetResourcesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get resources forbidden response a status code equal to that given
func (o *GetResourcesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetResourcesForbidden) Error() string {
	return fmt.Sprintf("[GET /resources][%d] getResourcesForbidden  %+v", 403, o.Payload)
}

func (o *GetResourcesForbidden) String() string {
	return fmt.Sprintf("[GET /resources][%d] getResourcesForbidden  %+v", 403, o.Payload)
}

func (o *GetResourcesForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetResourcesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourcesNotFound creates a GetResourcesNotFound with default headers values
func NewGetResourcesNotFound() *GetResourcesNotFound {
	return &GetResourcesNotFound{}
}

/*
GetResourcesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetResourcesNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get resources not found response has a 2xx status code
func (o *GetResourcesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resources not found response has a 3xx status code
func (o *GetResourcesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resources not found response has a 4xx status code
func (o *GetResourcesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resources not found response has a 5xx status code
func (o *GetResourcesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get resources not found response a status code equal to that given
func (o *GetResourcesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetResourcesNotFound) Error() string {
	return fmt.Sprintf("[GET /resources][%d] getResourcesNotFound  %+v", 404, o.Payload)
}

func (o *GetResourcesNotFound) String() string {
	return fmt.Sprintf("[GET /resources][%d] getResourcesNotFound  %+v", 404, o.Payload)
}

func (o *GetResourcesNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetResourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourcesInternalServerError creates a GetResourcesInternalServerError with default headers values
func NewGetResourcesInternalServerError() *GetResourcesInternalServerError {
	return &GetResourcesInternalServerError{}
}

/*
GetResourcesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetResourcesInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get resources internal server error response has a 2xx status code
func (o *GetResourcesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resources internal server error response has a 3xx status code
func (o *GetResourcesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resources internal server error response has a 4xx status code
func (o *GetResourcesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resources internal server error response has a 5xx status code
func (o *GetResourcesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get resources internal server error response a status code equal to that given
func (o *GetResourcesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetResourcesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /resources][%d] getResourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetResourcesInternalServerError) String() string {
	return fmt.Sprintf("[GET /resources][%d] getResourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetResourcesInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetResourcesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
