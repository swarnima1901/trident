// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewFpolicyConnectionGetParams creates a new FpolicyConnectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFpolicyConnectionGetParams() *FpolicyConnectionGetParams {
	return &FpolicyConnectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFpolicyConnectionGetParamsWithTimeout creates a new FpolicyConnectionGetParams object
// with the ability to set a timeout on a request.
func NewFpolicyConnectionGetParamsWithTimeout(timeout time.Duration) *FpolicyConnectionGetParams {
	return &FpolicyConnectionGetParams{
		timeout: timeout,
	}
}

// NewFpolicyConnectionGetParamsWithContext creates a new FpolicyConnectionGetParams object
// with the ability to set a context for a request.
func NewFpolicyConnectionGetParamsWithContext(ctx context.Context) *FpolicyConnectionGetParams {
	return &FpolicyConnectionGetParams{
		Context: ctx,
	}
}

// NewFpolicyConnectionGetParamsWithHTTPClient creates a new FpolicyConnectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFpolicyConnectionGetParamsWithHTTPClient(client *http.Client) *FpolicyConnectionGetParams {
	return &FpolicyConnectionGetParams{
		HTTPClient: client,
	}
}

/*
FpolicyConnectionGetParams contains all the parameters to send to the API endpoint

	for the fpolicy connection get operation.

	Typically these are written to a http.Request.
*/
type FpolicyConnectionGetParams struct {

	/* NodeUUID.

	   Node UUID.
	*/
	NodeUUIDPathParameter string

	/* PolicyName.

	   Policy name
	*/
	PolicyNamePathParameter string

	/* Server.

	   FPolicy Server IP address
	*/
	ServerPathParameter string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fpolicy connection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyConnectionGetParams) WithDefaults() *FpolicyConnectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fpolicy connection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyConnectionGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the fpolicy connection get params
func (o *FpolicyConnectionGetParams) WithTimeout(timeout time.Duration) *FpolicyConnectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fpolicy connection get params
func (o *FpolicyConnectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fpolicy connection get params
func (o *FpolicyConnectionGetParams) WithContext(ctx context.Context) *FpolicyConnectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fpolicy connection get params
func (o *FpolicyConnectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fpolicy connection get params
func (o *FpolicyConnectionGetParams) WithHTTPClient(client *http.Client) *FpolicyConnectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fpolicy connection get params
func (o *FpolicyConnectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNodeUUIDPathParameter adds the nodeUUID to the fpolicy connection get params
func (o *FpolicyConnectionGetParams) WithNodeUUIDPathParameter(nodeUUID string) *FpolicyConnectionGetParams {
	o.SetNodeUUIDPathParameter(nodeUUID)
	return o
}

// SetNodeUUIDPathParameter adds the nodeUuid to the fpolicy connection get params
func (o *FpolicyConnectionGetParams) SetNodeUUIDPathParameter(nodeUUID string) {
	o.NodeUUIDPathParameter = nodeUUID
}

// WithPolicyNamePathParameter adds the policyName to the fpolicy connection get params
func (o *FpolicyConnectionGetParams) WithPolicyNamePathParameter(policyName string) *FpolicyConnectionGetParams {
	o.SetPolicyNamePathParameter(policyName)
	return o
}

// SetPolicyNamePathParameter adds the policyName to the fpolicy connection get params
func (o *FpolicyConnectionGetParams) SetPolicyNamePathParameter(policyName string) {
	o.PolicyNamePathParameter = policyName
}

// WithServerPathParameter adds the server to the fpolicy connection get params
func (o *FpolicyConnectionGetParams) WithServerPathParameter(server string) *FpolicyConnectionGetParams {
	o.SetServerPathParameter(server)
	return o
}

// SetServerPathParameter adds the server to the fpolicy connection get params
func (o *FpolicyConnectionGetParams) SetServerPathParameter(server string) {
	o.ServerPathParameter = server
}

// WithSvmUUID adds the svmUUID to the fpolicy connection get params
func (o *FpolicyConnectionGetParams) WithSvmUUID(svmUUID string) *FpolicyConnectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the fpolicy connection get params
func (o *FpolicyConnectionGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *FpolicyConnectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param node.uuid
	if err := r.SetPathParam("node.uuid", o.NodeUUIDPathParameter); err != nil {
		return err
	}

	// path param policy.name
	if err := r.SetPathParam("policy.name", o.PolicyNamePathParameter); err != nil {
		return err
	}

	// path param server
	if err := r.SetPathParam("server", o.ServerPathParameter); err != nil {
		return err
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
