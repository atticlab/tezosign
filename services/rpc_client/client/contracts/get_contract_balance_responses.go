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

// GetContractBalanceReader is a Reader for the GetContractBalance structure.
type GetContractBalanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContractBalanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContractBalanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetContractBalanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetContractBalanceOK creates a GetContractBalanceOK with default headers values
func NewGetContractBalanceOK() *GetContractBalanceOK {
	return &GetContractBalanceOK{}
}

/*GetContractBalanceOK handles this case with default header values.

Endpoint for contract balance
*/
type GetContractBalanceOK struct {
	Payload string
}

func (o *GetContractBalanceOK) Error() string {
	return fmt.Sprintf("[GET /chains/main/blocks/head/context/contracts/{contract}/balance][%d] getContractBalanceOK  %+v", 200, o.Payload)
}

func (o *GetContractBalanceOK) GetPayload() string {
	return o.Payload
}

func (o *GetContractBalanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContractBalanceInternalServerError creates a GetContractBalanceInternalServerError with default headers values
func NewGetContractBalanceInternalServerError() *GetContractBalanceInternalServerError {
	return &GetContractBalanceInternalServerError{}
}

/*GetContractBalanceInternalServerError handles this case with default header values.

Internal error
*/
type GetContractBalanceInternalServerError struct {
}

func (o *GetContractBalanceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /chains/main/blocks/head/context/contracts/{contract}/balance][%d] getContractBalanceInternalServerError ", 500)
}

func (o *GetContractBalanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
