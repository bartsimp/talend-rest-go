// Code generated by go-swagger; DO NOT EDIT.

package promotions_executables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bartsimp/talend-rest-go/models"
)

// DeletePromotionReader is a Reader for the DeletePromotion structure.
type DeletePromotionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePromotionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePromotionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeletePromotionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeletePromotionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePromotionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePromotionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeletePromotionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeletePromotionNoContent creates a DeletePromotionNoContent with default headers values
func NewDeletePromotionNoContent() *DeletePromotionNoContent {
	return &DeletePromotionNoContent{}
}

/*
DeletePromotionNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeletePromotionNoContent struct {
}

// IsSuccess returns true when this delete promotion no content response has a 2xx status code
func (o *DeletePromotionNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete promotion no content response has a 3xx status code
func (o *DeletePromotionNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete promotion no content response has a 4xx status code
func (o *DeletePromotionNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete promotion no content response has a 5xx status code
func (o *DeletePromotionNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete promotion no content response a status code equal to that given
func (o *DeletePromotionNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeletePromotionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /executables/promotions/{promotionId}][%d] deletePromotionNoContent ", 204)
}

func (o *DeletePromotionNoContent) String() string {
	return fmt.Sprintf("[DELETE /executables/promotions/{promotionId}][%d] deletePromotionNoContent ", 204)
}

func (o *DeletePromotionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePromotionBadRequest creates a DeletePromotionBadRequest with default headers values
func NewDeletePromotionBadRequest() *DeletePromotionBadRequest {
	return &DeletePromotionBadRequest{}
}

/*
DeletePromotionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeletePromotionBadRequest struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete promotion bad request response has a 2xx status code
func (o *DeletePromotionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete promotion bad request response has a 3xx status code
func (o *DeletePromotionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete promotion bad request response has a 4xx status code
func (o *DeletePromotionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete promotion bad request response has a 5xx status code
func (o *DeletePromotionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete promotion bad request response a status code equal to that given
func (o *DeletePromotionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeletePromotionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /executables/promotions/{promotionId}][%d] deletePromotionBadRequest  %+v", 400, o.Payload)
}

func (o *DeletePromotionBadRequest) String() string {
	return fmt.Sprintf("[DELETE /executables/promotions/{promotionId}][%d] deletePromotionBadRequest  %+v", 400, o.Payload)
}

func (o *DeletePromotionBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeletePromotionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePromotionUnauthorized creates a DeletePromotionUnauthorized with default headers values
func NewDeletePromotionUnauthorized() *DeletePromotionUnauthorized {
	return &DeletePromotionUnauthorized{}
}

/*
DeletePromotionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeletePromotionUnauthorized struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete promotion unauthorized response has a 2xx status code
func (o *DeletePromotionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete promotion unauthorized response has a 3xx status code
func (o *DeletePromotionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete promotion unauthorized response has a 4xx status code
func (o *DeletePromotionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete promotion unauthorized response has a 5xx status code
func (o *DeletePromotionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete promotion unauthorized response a status code equal to that given
func (o *DeletePromotionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeletePromotionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /executables/promotions/{promotionId}][%d] deletePromotionUnauthorized  %+v", 401, o.Payload)
}

func (o *DeletePromotionUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /executables/promotions/{promotionId}][%d] deletePromotionUnauthorized  %+v", 401, o.Payload)
}

func (o *DeletePromotionUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeletePromotionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePromotionForbidden creates a DeletePromotionForbidden with default headers values
func NewDeletePromotionForbidden() *DeletePromotionForbidden {
	return &DeletePromotionForbidden{}
}

/*
DeletePromotionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeletePromotionForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete promotion forbidden response has a 2xx status code
func (o *DeletePromotionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete promotion forbidden response has a 3xx status code
func (o *DeletePromotionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete promotion forbidden response has a 4xx status code
func (o *DeletePromotionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete promotion forbidden response has a 5xx status code
func (o *DeletePromotionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete promotion forbidden response a status code equal to that given
func (o *DeletePromotionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeletePromotionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /executables/promotions/{promotionId}][%d] deletePromotionForbidden  %+v", 403, o.Payload)
}

func (o *DeletePromotionForbidden) String() string {
	return fmt.Sprintf("[DELETE /executables/promotions/{promotionId}][%d] deletePromotionForbidden  %+v", 403, o.Payload)
}

func (o *DeletePromotionForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeletePromotionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePromotionNotFound creates a DeletePromotionNotFound with default headers values
func NewDeletePromotionNotFound() *DeletePromotionNotFound {
	return &DeletePromotionNotFound{}
}

/*
DeletePromotionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeletePromotionNotFound struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete promotion not found response has a 2xx status code
func (o *DeletePromotionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete promotion not found response has a 3xx status code
func (o *DeletePromotionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete promotion not found response has a 4xx status code
func (o *DeletePromotionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete promotion not found response has a 5xx status code
func (o *DeletePromotionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete promotion not found response a status code equal to that given
func (o *DeletePromotionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeletePromotionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /executables/promotions/{promotionId}][%d] deletePromotionNotFound  %+v", 404, o.Payload)
}

func (o *DeletePromotionNotFound) String() string {
	return fmt.Sprintf("[DELETE /executables/promotions/{promotionId}][%d] deletePromotionNotFound  %+v", 404, o.Payload)
}

func (o *DeletePromotionNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeletePromotionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePromotionInternalServerError creates a DeletePromotionInternalServerError with default headers values
func NewDeletePromotionInternalServerError() *DeletePromotionInternalServerError {
	return &DeletePromotionInternalServerError{}
}

/*
DeletePromotionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeletePromotionInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this delete promotion internal server error response has a 2xx status code
func (o *DeletePromotionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete promotion internal server error response has a 3xx status code
func (o *DeletePromotionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete promotion internal server error response has a 4xx status code
func (o *DeletePromotionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete promotion internal server error response has a 5xx status code
func (o *DeletePromotionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete promotion internal server error response a status code equal to that given
func (o *DeletePromotionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeletePromotionInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /executables/promotions/{promotionId}][%d] deletePromotionInternalServerError  %+v", 500, o.Payload)
}

func (o *DeletePromotionInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /executables/promotions/{promotionId}][%d] deletePromotionInternalServerError  %+v", 500, o.Payload)
}

func (o *DeletePromotionInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeletePromotionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}