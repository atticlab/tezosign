// Code generated by go-swagger; DO NOT EDIT.

package contracts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetContractStorageReader is a Reader for the GetContractStorage structure.
type GetContractStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContractStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContractStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetContractStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetContractStorageOK creates a GetContractStorageOK with default headers values
func NewGetContractStorageOK() *GetContractStorageOK {
	return &GetContractStorageOK{}
}

/*GetContractStorageOK handles this case with default header values.

Endpoint for contract storage
*/
type GetContractStorageOK struct {
	Payload interface{}
}

func (o *GetContractStorageOK) Error() string {
	return fmt.Sprintf("[GET /chains/main/blocks/head/context/contracts/{contract}/storage][%d] getContractStorageOK  %+v", 200, o.Payload)
}

func (o *GetContractStorageOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetContractStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContractStorageInternalServerError creates a GetContractStorageInternalServerError with default headers values
func NewGetContractStorageInternalServerError() *GetContractStorageInternalServerError {
	return &GetContractStorageInternalServerError{}
}

/*GetContractStorageInternalServerError handles this case with default header values.

Internal error
*/
type GetContractStorageInternalServerError struct {
}

func (o *GetContractStorageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /chains/main/blocks/head/context/contracts/{contract}/storage][%d] getContractStorageInternalServerError ", 500)
}

func (o *GetContractStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
