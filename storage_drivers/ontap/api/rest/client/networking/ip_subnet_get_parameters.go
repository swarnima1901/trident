// Code generated by go-swagger; DO NOT EDIT.

package networking

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
	"github.com/go-openapi/swag"
)

// NewIPSubnetGetParams creates a new IPSubnetGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIPSubnetGetParams() *IPSubnetGetParams {
	return &IPSubnetGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIPSubnetGetParamsWithTimeout creates a new IPSubnetGetParams object
// with the ability to set a timeout on a request.
func NewIPSubnetGetParamsWithTimeout(timeout time.Duration) *IPSubnetGetParams {
	return &IPSubnetGetParams{
		timeout: timeout,
	}
}

// NewIPSubnetGetParamsWithContext creates a new IPSubnetGetParams object
// with the ability to set a context for a request.
func NewIPSubnetGetParamsWithContext(ctx context.Context) *IPSubnetGetParams {
	return &IPSubnetGetParams{
		Context: ctx,
	}
}

// NewIPSubnetGetParamsWithHTTPClient creates a new IPSubnetGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewIPSubnetGetParamsWithHTTPClient(client *http.Client) *IPSubnetGetParams {
	return &IPSubnetGetParams{
		HTTPClient: client,
	}
}

/*
IPSubnetGetParams contains all the parameters to send to the API endpoint

	for the ip subnet get operation.

	Typically these are written to a http.Request.
*/
type IPSubnetGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* UUID.

	   IP subnet UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ip subnet get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IPSubnetGetParams) WithDefaults() *IPSubnetGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ip subnet get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IPSubnetGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ip subnet get params
func (o *IPSubnetGetParams) WithTimeout(timeout time.Duration) *IPSubnetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ip subnet get params
func (o *IPSubnetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ip subnet get params
func (o *IPSubnetGetParams) WithContext(ctx context.Context) *IPSubnetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ip subnet get params
func (o *IPSubnetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ip subnet get params
func (o *IPSubnetGetParams) WithHTTPClient(client *http.Client) *IPSubnetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ip subnet get params
func (o *IPSubnetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the ip subnet get params
func (o *IPSubnetGetParams) WithFieldsQueryParameter(fields []string) *IPSubnetGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the ip subnet get params
func (o *IPSubnetGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithUUIDPathParameter adds the uuid to the ip subnet get params
func (o *IPSubnetGetParams) WithUUIDPathParameter(uuid string) *IPSubnetGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the ip subnet get params
func (o *IPSubnetGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *IPSubnetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamIPSubnetGet binds the parameter fields
func (o *IPSubnetGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
