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
	"github.com/go-openapi/swag"
)

// NewEmsFilterRuleGetParams creates a new EmsFilterRuleGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmsFilterRuleGetParams() *EmsFilterRuleGetParams {
	return &EmsFilterRuleGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmsFilterRuleGetParamsWithTimeout creates a new EmsFilterRuleGetParams object
// with the ability to set a timeout on a request.
func NewEmsFilterRuleGetParamsWithTimeout(timeout time.Duration) *EmsFilterRuleGetParams {
	return &EmsFilterRuleGetParams{
		timeout: timeout,
	}
}

// NewEmsFilterRuleGetParamsWithContext creates a new EmsFilterRuleGetParams object
// with the ability to set a context for a request.
func NewEmsFilterRuleGetParamsWithContext(ctx context.Context) *EmsFilterRuleGetParams {
	return &EmsFilterRuleGetParams{
		Context: ctx,
	}
}

// NewEmsFilterRuleGetParamsWithHTTPClient creates a new EmsFilterRuleGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmsFilterRuleGetParamsWithHTTPClient(client *http.Client) *EmsFilterRuleGetParams {
	return &EmsFilterRuleGetParams{
		HTTPClient: client,
	}
}

/*
EmsFilterRuleGetParams contains all the parameters to send to the API endpoint

	for the ems filter rule get operation.

	Typically these are written to a http.Request.
*/
type EmsFilterRuleGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Index.

	   Filter index
	*/
	IndexPathParameter string

	/* Name.

	   Filter name
	*/
	NamePathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ems filter rule get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsFilterRuleGetParams) WithDefaults() *EmsFilterRuleGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ems filter rule get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsFilterRuleGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ems filter rule get params
func (o *EmsFilterRuleGetParams) WithTimeout(timeout time.Duration) *EmsFilterRuleGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ems filter rule get params
func (o *EmsFilterRuleGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ems filter rule get params
func (o *EmsFilterRuleGetParams) WithContext(ctx context.Context) *EmsFilterRuleGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ems filter rule get params
func (o *EmsFilterRuleGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ems filter rule get params
func (o *EmsFilterRuleGetParams) WithHTTPClient(client *http.Client) *EmsFilterRuleGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ems filter rule get params
func (o *EmsFilterRuleGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the ems filter rule get params
func (o *EmsFilterRuleGetParams) WithFieldsQueryParameter(fields []string) *EmsFilterRuleGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the ems filter rule get params
func (o *EmsFilterRuleGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIndexPathParameter adds the index to the ems filter rule get params
func (o *EmsFilterRuleGetParams) WithIndexPathParameter(index string) *EmsFilterRuleGetParams {
	o.SetIndexPathParameter(index)
	return o
}

// SetIndexPathParameter adds the index to the ems filter rule get params
func (o *EmsFilterRuleGetParams) SetIndexPathParameter(index string) {
	o.IndexPathParameter = index
}

// WithNamePathParameter adds the name to the ems filter rule get params
func (o *EmsFilterRuleGetParams) WithNamePathParameter(name string) *EmsFilterRuleGetParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the ems filter rule get params
func (o *EmsFilterRuleGetParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WriteToRequest writes these params to a swagger request
func (o *EmsFilterRuleGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param index
	if err := r.SetPathParam("index", o.IndexPathParameter); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.NamePathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamEmsFilterRuleGet binds the parameter fields
func (o *EmsFilterRuleGetParams) bindParamFields(formats strfmt.Registry) []string {
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
