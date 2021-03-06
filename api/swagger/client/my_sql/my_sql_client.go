// Code generated by go-swagger; DO NOT EDIT.

package my_sql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new my sql API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for my sql API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Add add API
*/
func (a *Client) Add(params *AddParams) (*AddOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Add",
		Method:             "POST",
		PathPattern:        "/v0/mysql",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddOK), nil

}

/*
List list API
*/
func (a *Client) List(params *ListParams) (*ListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "List",
		Method:             "GET",
		PathPattern:        "/v0/mysql",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListOK), nil

}

/*
Remove remove API
*/
func (a *Client) Remove(params *RemoveParams) (*RemoveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Remove",
		Method:             "DELETE",
		PathPattern:        "/v0/mysql/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemoveOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
