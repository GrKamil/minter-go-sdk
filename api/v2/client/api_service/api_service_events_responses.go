// Code generated by go-swagger; DO NOT EDIT.

package api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/MinterTeam/minter-go-sdk/api/v2/models"
)

// APIServiceEventsReader is a Reader for the APIServiceEvents structure.
type APIServiceEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAPIServiceEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceEventsOK creates a APIServiceEventsOK with default headers values
func NewAPIServiceEventsOK() *APIServiceEventsOK {
	return &APIServiceEventsOK{}
}

/*APIServiceEventsOK handles this case with default header values.

A successful response.
*/
type APIServiceEventsOK struct {
	Payload *models.APIPbEventsResponse
}

func (o *APIServiceEventsOK) Error() string {
	return fmt.Sprintf("[GET /events/{height}][%d] apiServiceEventsOK  %+v", 200, o.Payload)
}

func (o *APIServiceEventsOK) GetPayload() *models.APIPbEventsResponse {
	return o.Payload
}

func (o *APIServiceEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceEventsDefault creates a APIServiceEventsDefault with default headers values
func NewAPIServiceEventsDefault(code int) *APIServiceEventsDefault {
	return &APIServiceEventsDefault{
		_statusCode: code,
	}
}

/*APIServiceEventsDefault handles this case with default header values.

An unexpected error response
*/
type APIServiceEventsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the Api service events default response
func (o *APIServiceEventsDefault) Code() int {
	return o._statusCode
}

func (o *APIServiceEventsDefault) Error() string {
	return fmt.Sprintf("[GET /events/{height}][%d] ApiService_Events default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceEventsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *APIServiceEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}