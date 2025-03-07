// Code generated by go-swagger; DO NOT EDIT.

package object_store

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

// NewS3UserDeleteParams creates a new S3UserDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3UserDeleteParams() *S3UserDeleteParams {
	return &S3UserDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3UserDeleteParamsWithTimeout creates a new S3UserDeleteParams object
// with the ability to set a timeout on a request.
func NewS3UserDeleteParamsWithTimeout(timeout time.Duration) *S3UserDeleteParams {
	return &S3UserDeleteParams{
		timeout: timeout,
	}
}

// NewS3UserDeleteParamsWithContext creates a new S3UserDeleteParams object
// with the ability to set a context for a request.
func NewS3UserDeleteParamsWithContext(ctx context.Context) *S3UserDeleteParams {
	return &S3UserDeleteParams{
		Context: ctx,
	}
}

// NewS3UserDeleteParamsWithHTTPClient creates a new S3UserDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3UserDeleteParamsWithHTTPClient(client *http.Client) *S3UserDeleteParams {
	return &S3UserDeleteParams{
		HTTPClient: client,
	}
}

/*
S3UserDeleteParams contains all the parameters to send to the API endpoint

	for the s3 user delete operation.

	Typically these are written to a http.Request.
*/
type S3UserDeleteParams struct {

	/* Name.

	   User name
	*/
	NamePathParameter string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 user delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3UserDeleteParams) WithDefaults() *S3UserDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 user delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3UserDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the s3 user delete params
func (o *S3UserDeleteParams) WithTimeout(timeout time.Duration) *S3UserDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 user delete params
func (o *S3UserDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 user delete params
func (o *S3UserDeleteParams) WithContext(ctx context.Context) *S3UserDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 user delete params
func (o *S3UserDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 user delete params
func (o *S3UserDeleteParams) WithHTTPClient(client *http.Client) *S3UserDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 user delete params
func (o *S3UserDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamePathParameter adds the name to the s3 user delete params
func (o *S3UserDeleteParams) WithNamePathParameter(name string) *S3UserDeleteParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the s3 user delete params
func (o *S3UserDeleteParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WithSVMUUIDPathParameter adds the svmUUID to the s3 user delete params
func (o *S3UserDeleteParams) WithSVMUUIDPathParameter(svmUUID string) *S3UserDeleteParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the s3 user delete params
func (o *S3UserDeleteParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *S3UserDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.NamePathParameter); err != nil {
		return err
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
