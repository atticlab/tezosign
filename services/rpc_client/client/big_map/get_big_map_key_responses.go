// Code generated by go-swagger; DO NOT EDIT.

package big_map

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetBigMapKeyReader is a Reader for the GetBigMapKey structure.
type GetBigMapKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBigMapKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBigMapKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetBigMapKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetBigMapKeyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBigMapKeyOK creates a GetBigMapKeyOK with default headers values
func NewGetBigMapKeyOK() *GetBigMapKeyOK {
	return &GetBigMapKeyOK{}
}

/*GetBigMapKeyOK handles this case with default header values.

Endpoint for big map keys
*/
type GetBigMapKeyOK struct {
	Payload interface{}
}

func (o *GetBigMapKeyOK) Error() string {
	return fmt.Sprintf("[GET /chains/main/blocks/head/context/big_maps/{big_map_id}/{key_hash}][%d] getBigMapKeyOK  %+v", 200, o.Payload)
}

func (o *GetBigMapKeyOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetBigMapKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBigMapKeyNotFound creates a GetBigMapKeyNotFound with default headers values
func NewGetBigMapKeyNotFound() *GetBigMapKeyNotFound {
	return &GetBigMapKeyNotFound{}
}

/*GetBigMapKeyNotFound handles this case with default header values.

Key not found
*/
type GetBigMapKeyNotFound struct {
	Payload interface{}
}

func (o *GetBigMapKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /chains/main/blocks/head/context/big_maps/{big_map_id}/{key_hash}][%d] getBigMapKeyNotFound  %+v", 404, o.Payload)
}

func (o *GetBigMapKeyNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetBigMapKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBigMapKeyInternalServerError creates a GetBigMapKeyInternalServerError with default headers values
func NewGetBigMapKeyInternalServerError() *GetBigMapKeyInternalServerError {
	return &GetBigMapKeyInternalServerError{}
}

/*GetBigMapKeyInternalServerError handles this case with default header values.

Internal error
*/
type GetBigMapKeyInternalServerError struct {
}

func (o *GetBigMapKeyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /chains/main/blocks/head/context/big_maps/{big_map_id}/{key_hash}][%d] getBigMapKeyInternalServerError ", 500)
}

func (o *GetBigMapKeyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
