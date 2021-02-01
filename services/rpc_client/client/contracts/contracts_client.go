// Code generated by go-swagger; DO NOT EDIT.

package contracts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new contracts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for contracts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetContractBalance get contract balance API
*/
func (a *Client) GetContractBalance(params *GetContractBalanceParams) (*GetContractBalanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetContractBalanceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getContractBalance",
		Method:             "GET",
		PathPattern:        "/chains/main/blocks/head/context/contracts/{contract}/balance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetContractBalanceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetContractBalanceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getContractBalance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetContractManagerKey get contract manager key API
*/
func (a *Client) GetContractManagerKey(params *GetContractManagerKeyParams) (*GetContractManagerKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetContractManagerKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getContractManagerKey",
		Method:             "GET",
		PathPattern:        "/chains/main/blocks/head/context/contracts/{contract}/manager_key",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetContractManagerKeyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetContractManagerKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getContractManagerKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetContractScript get contract script API
*/
func (a *Client) GetContractScript(params *GetContractScriptParams) (*GetContractScriptOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetContractScriptParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getContractScript",
		Method:             "GET",
		PathPattern:        "/chains/main/blocks/head/context/contracts/{contract}/script",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetContractScriptReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetContractScriptOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getContractScript: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetContractStorage get contract storage API
*/
func (a *Client) GetContractStorage(params *GetContractStorageParams) (*GetContractStorageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetContractStorageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getContractStorage",
		Method:             "GET",
		PathPattern:        "/chains/main/blocks/head/context/contracts/{contract}/storage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetContractStorageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetContractStorageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getContractStorage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
