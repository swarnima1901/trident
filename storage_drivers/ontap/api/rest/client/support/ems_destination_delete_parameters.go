// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewEmsDestinationDeleteParams creates a new EmsDestinationDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmsDestinationDeleteParams() *EmsDestinationDeleteParams {
	return &EmsDestinationDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmsDestinationDeleteParamsWithTimeout creates a new EmsDestinationDeleteParams object
// with the ability to set a timeout on a request.
func NewEmsDestinationDeleteParamsWithTimeout(timeout time.Duration) *EmsDestinationDeleteParams {
	return &EmsDestinationDeleteParams{
		timeout: timeout,
	}
}

// NewEmsDestinationDeleteParamsWithContext creates a new EmsDestinationDeleteParams object
// with the ability to set a context for a request.
func NewEmsDestinationDeleteParamsWithContext(ctx context.Context) *EmsDestinationDeleteParams {
	return &EmsDestinationDeleteParams{
		Context: ctx,
	}
}

// NewEmsDestinationDeleteParamsWithHTTPClient creates a new EmsDestinationDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmsDestinationDeleteParamsWithHTTPClient(client *http.Client) *EmsDestinationDeleteParams {
	return &EmsDestinationDeleteParams{
		HTTPClient: client,
	}
}

/*
EmsDestinationDeleteParams contains all the parameters to send to the API endpoint

	for the ems destination delete operation.

	Typically these are written to a http.Request.
*/
type EmsDestinationDeleteParams struct {

	/* Name.

	   Destination name
	*/
	NamePathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ems destination delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsDestinationDeleteParams) WithDefaults() *EmsDestinationDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ems destination delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsDestinationDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ems destination delete params
func (o *EmsDestinationDeleteParams) WithTimeout(timeout time.Duration) *EmsDestinationDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ems destination delete params
func (o *EmsDestinationDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ems destination delete params
func (o *EmsDestinationDeleteParams) WithContext(ctx context.Context) *EmsDestinationDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ems destination delete params
func (o *EmsDestinationDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ems destination delete params
func (o *EmsDestinationDeleteParams) WithHTTPClient(client *http.Client) *EmsDestinationDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ems destination delete params
func (o *EmsDestinationDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamePathParameter adds the name to the ems destination delete params
func (o *EmsDestinationDeleteParams) WithNamePathParameter(name string) *EmsDestinationDeleteParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the ems destination delete params
func (o *EmsDestinationDeleteParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WriteToRequest writes these params to a swagger request
func (o *EmsDestinationDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.NamePathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
