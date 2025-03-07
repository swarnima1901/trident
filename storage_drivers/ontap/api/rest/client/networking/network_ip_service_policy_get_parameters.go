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

// NewNetworkIPServicePolicyGetParams creates a new NetworkIPServicePolicyGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkIPServicePolicyGetParams() *NetworkIPServicePolicyGetParams {
	return &NetworkIPServicePolicyGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkIPServicePolicyGetParamsWithTimeout creates a new NetworkIPServicePolicyGetParams object
// with the ability to set a timeout on a request.
func NewNetworkIPServicePolicyGetParamsWithTimeout(timeout time.Duration) *NetworkIPServicePolicyGetParams {
	return &NetworkIPServicePolicyGetParams{
		timeout: timeout,
	}
}

// NewNetworkIPServicePolicyGetParamsWithContext creates a new NetworkIPServicePolicyGetParams object
// with the ability to set a context for a request.
func NewNetworkIPServicePolicyGetParamsWithContext(ctx context.Context) *NetworkIPServicePolicyGetParams {
	return &NetworkIPServicePolicyGetParams{
		Context: ctx,
	}
}

// NewNetworkIPServicePolicyGetParamsWithHTTPClient creates a new NetworkIPServicePolicyGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkIPServicePolicyGetParamsWithHTTPClient(client *http.Client) *NetworkIPServicePolicyGetParams {
	return &NetworkIPServicePolicyGetParams{
		HTTPClient: client,
	}
}

/*
NetworkIPServicePolicyGetParams contains all the parameters to send to the API endpoint

	for the network ip service policy get operation.

	Typically these are written to a http.Request.
*/
type NetworkIPServicePolicyGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeoutQueryParameter *int64

	/* UUID.

	   Service policy UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network ip service policy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPServicePolicyGetParams) WithDefaults() *NetworkIPServicePolicyGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network ip service policy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPServicePolicyGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := NetworkIPServicePolicyGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the network ip service policy get params
func (o *NetworkIPServicePolicyGetParams) WithTimeout(timeout time.Duration) *NetworkIPServicePolicyGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network ip service policy get params
func (o *NetworkIPServicePolicyGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network ip service policy get params
func (o *NetworkIPServicePolicyGetParams) WithContext(ctx context.Context) *NetworkIPServicePolicyGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network ip service policy get params
func (o *NetworkIPServicePolicyGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network ip service policy get params
func (o *NetworkIPServicePolicyGetParams) WithHTTPClient(client *http.Client) *NetworkIPServicePolicyGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network ip service policy get params
func (o *NetworkIPServicePolicyGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the network ip service policy get params
func (o *NetworkIPServicePolicyGetParams) WithFieldsQueryParameter(fields []string) *NetworkIPServicePolicyGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the network ip service policy get params
func (o *NetworkIPServicePolicyGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the network ip service policy get params
func (o *NetworkIPServicePolicyGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *NetworkIPServicePolicyGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the network ip service policy get params
func (o *NetworkIPServicePolicyGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the network ip service policy get params
func (o *NetworkIPServicePolicyGetParams) WithOrderByQueryParameter(orderBy []string) *NetworkIPServicePolicyGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the network ip service policy get params
func (o *NetworkIPServicePolicyGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the network ip service policy get params
func (o *NetworkIPServicePolicyGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *NetworkIPServicePolicyGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the network ip service policy get params
func (o *NetworkIPServicePolicyGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the network ip service policy get params
func (o *NetworkIPServicePolicyGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *NetworkIPServicePolicyGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the network ip service policy get params
func (o *NetworkIPServicePolicyGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithUUIDPathParameter adds the uuid to the network ip service policy get params
func (o *NetworkIPServicePolicyGetParams) WithUUIDPathParameter(uuid string) *NetworkIPServicePolicyGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the network ip service policy get params
func (o *NetworkIPServicePolicyGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkIPServicePolicyGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.MaxRecordsQueryParameter != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecordsQueryParameter != nil {
			qrMaxRecords = *o.MaxRecordsQueryParameter
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
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

// bindParamNetworkIPServicePolicyGet binds the parameter fields
func (o *NetworkIPServicePolicyGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamNetworkIPServicePolicyGet binds the parameter order_by
func (o *NetworkIPServicePolicyGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderByQueryParameter

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}
