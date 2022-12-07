// Code generated by go-swagger; DO NOT EDIT.

package runtime_engines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bartsimp/talend-rest-go/models"
)

// GetRemoteEnginesReader is a Reader for the GetRemoteEngines structure.
type GetRemoteEnginesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRemoteEnginesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRemoteEnginesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRemoteEnginesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRemoteEnginesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRemoteEnginesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRemoteEnginesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRemoteEnginesOK creates a GetRemoteEnginesOK with default headers values
func NewGetRemoteEnginesOK() *GetRemoteEnginesOK {
	return &GetRemoteEnginesOK{}
}

/*
GetRemoteEnginesOK describes a response with status code 200, with default header values.

OK
*/
type GetRemoteEnginesOK struct {
	Payload []*models.Engine
}

// IsSuccess returns true when this get remote engines o k response has a 2xx status code
func (o *GetRemoteEnginesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get remote engines o k response has a 3xx status code
func (o *GetRemoteEnginesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remote engines o k response has a 4xx status code
func (o *GetRemoteEnginesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get remote engines o k response has a 5xx status code
func (o *GetRemoteEnginesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get remote engines o k response a status code equal to that given
func (o *GetRemoteEnginesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRemoteEnginesOK) Error() string {
	return fmt.Sprintf("[GET /runtimes/remote-engines][%d] getRemoteEnginesOK  %+v", 200, o.Payload)
}

func (o *GetRemoteEnginesOK) String() string {
	return fmt.Sprintf("[GET /runtimes/remote-engines][%d] getRemoteEnginesOK  %+v", 200, o.Payload)
}

func (o *GetRemoteEnginesOK) GetPayload() []*models.Engine {
	return o.Payload
}

func (o *GetRemoteEnginesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRemoteEnginesUnauthorized creates a GetRemoteEnginesUnauthorized with default headers values
func NewGetRemoteEnginesUnauthorized() *GetRemoteEnginesUnauthorized {
	return &GetRemoteEnginesUnauthorized{}
}

/*
GetRemoteEnginesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRemoteEnginesUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get remote engines unauthorized response has a 2xx status code
func (o *GetRemoteEnginesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get remote engines unauthorized response has a 3xx status code
func (o *GetRemoteEnginesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remote engines unauthorized response has a 4xx status code
func (o *GetRemoteEnginesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get remote engines unauthorized response has a 5xx status code
func (o *GetRemoteEnginesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get remote engines unauthorized response a status code equal to that given
func (o *GetRemoteEnginesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRemoteEnginesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /runtimes/remote-engines][%d] getRemoteEnginesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRemoteEnginesUnauthorized) String() string {
	return fmt.Sprintf("[GET /runtimes/remote-engines][%d] getRemoteEnginesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRemoteEnginesUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRemoteEnginesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRemoteEnginesForbidden creates a GetRemoteEnginesForbidden with default headers values
func NewGetRemoteEnginesForbidden() *GetRemoteEnginesForbidden {
	return &GetRemoteEnginesForbidden{}
}

/*
GetRemoteEnginesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRemoteEnginesForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get remote engines forbidden response has a 2xx status code
func (o *GetRemoteEnginesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get remote engines forbidden response has a 3xx status code
func (o *GetRemoteEnginesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remote engines forbidden response has a 4xx status code
func (o *GetRemoteEnginesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get remote engines forbidden response has a 5xx status code
func (o *GetRemoteEnginesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get remote engines forbidden response a status code equal to that given
func (o *GetRemoteEnginesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRemoteEnginesForbidden) Error() string {
	return fmt.Sprintf("[GET /runtimes/remote-engines][%d] getRemoteEnginesForbidden  %+v", 403, o.Payload)
}

func (o *GetRemoteEnginesForbidden) String() string {
	return fmt.Sprintf("[GET /runtimes/remote-engines][%d] getRemoteEnginesForbidden  %+v", 403, o.Payload)
}

func (o *GetRemoteEnginesForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRemoteEnginesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRemoteEnginesNotFound creates a GetRemoteEnginesNotFound with default headers values
func NewGetRemoteEnginesNotFound() *GetRemoteEnginesNotFound {
	return &GetRemoteEnginesNotFound{}
}

/*
GetRemoteEnginesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRemoteEnginesNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get remote engines not found response has a 2xx status code
func (o *GetRemoteEnginesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get remote engines not found response has a 3xx status code
func (o *GetRemoteEnginesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remote engines not found response has a 4xx status code
func (o *GetRemoteEnginesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get remote engines not found response has a 5xx status code
func (o *GetRemoteEnginesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get remote engines not found response a status code equal to that given
func (o *GetRemoteEnginesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRemoteEnginesNotFound) Error() string {
	return fmt.Sprintf("[GET /runtimes/remote-engines][%d] getRemoteEnginesNotFound  %+v", 404, o.Payload)
}

func (o *GetRemoteEnginesNotFound) String() string {
	return fmt.Sprintf("[GET /runtimes/remote-engines][%d] getRemoteEnginesNotFound  %+v", 404, o.Payload)
}

func (o *GetRemoteEnginesNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRemoteEnginesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRemoteEnginesInternalServerError creates a GetRemoteEnginesInternalServerError with default headers values
func NewGetRemoteEnginesInternalServerError() *GetRemoteEnginesInternalServerError {
	return &GetRemoteEnginesInternalServerError{}
}

/*
GetRemoteEnginesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetRemoteEnginesInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this get remote engines internal server error response has a 2xx status code
func (o *GetRemoteEnginesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get remote engines internal server error response has a 3xx status code
func (o *GetRemoteEnginesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get remote engines internal server error response has a 4xx status code
func (o *GetRemoteEnginesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get remote engines internal server error response has a 5xx status code
func (o *GetRemoteEnginesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get remote engines internal server error response a status code equal to that given
func (o *GetRemoteEnginesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRemoteEnginesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /runtimes/remote-engines][%d] getRemoteEnginesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRemoteEnginesInternalServerError) String() string {
	return fmt.Sprintf("[GET /runtimes/remote-engines][%d] getRemoteEnginesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRemoteEnginesInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetRemoteEnginesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
