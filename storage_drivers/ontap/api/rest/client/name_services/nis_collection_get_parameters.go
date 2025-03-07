// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// NewNisCollectionGetParams creates a new NisCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNisCollectionGetParams() *NisCollectionGetParams {
	return &NisCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNisCollectionGetParamsWithTimeout creates a new NisCollectionGetParams object
// with the ability to set a timeout on a request.
func NewNisCollectionGetParamsWithTimeout(timeout time.Duration) *NisCollectionGetParams {
	return &NisCollectionGetParams{
		timeout: timeout,
	}
}

// NewNisCollectionGetParamsWithContext creates a new NisCollectionGetParams object
// with the ability to set a context for a request.
func NewNisCollectionGetParamsWithContext(ctx context.Context) *NisCollectionGetParams {
	return &NisCollectionGetParams{
		Context: ctx,
	}
}

// NewNisCollectionGetParamsWithHTTPClient creates a new NisCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNisCollectionGetParamsWithHTTPClient(client *http.Client) *NisCollectionGetParams {
	return &NisCollectionGetParams{
		HTTPClient: client,
	}
}

/*
NisCollectionGetParams contains all the parameters to send to the API endpoint

	for the nis collection get operation.

	Typically these are written to a http.Request.
*/
type NisCollectionGetParams struct {

	/* BindingDetailsServer.

	   Filter by binding_details.server
	*/
	BindingDetailsServerQueryParameter *string

	/* BindingDetailsStatusCode.

	   Filter by binding_details.status.code
	*/
	BindingDetailsStatusCodeQueryParameter *string

	/* BindingDetailsStatusMessage.

	   Filter by binding_details.status.message
	*/
	BindingDetailsStatusMessageQueryParameter *string

	/* BoundServers.

	   Filter by bound_servers
	*/
	BoundServersQueryParameter *string

	/* Domain.

	   Filter by domain
	*/
	DomainQueryParameter *string

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

	/* Servers.

	   Filter by servers
	*/
	ServersQueryParameter *string

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nis collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NisCollectionGetParams) WithDefaults() *NisCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nis collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NisCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := NisCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the nis collection get params
func (o *NisCollectionGetParams) WithTimeout(timeout time.Duration) *NisCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nis collection get params
func (o *NisCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nis collection get params
func (o *NisCollectionGetParams) WithContext(ctx context.Context) *NisCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nis collection get params
func (o *NisCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nis collection get params
func (o *NisCollectionGetParams) WithHTTPClient(client *http.Client) *NisCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nis collection get params
func (o *NisCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBindingDetailsServerQueryParameter adds the bindingDetailsServer to the nis collection get params
func (o *NisCollectionGetParams) WithBindingDetailsServerQueryParameter(bindingDetailsServer *string) *NisCollectionGetParams {
	o.SetBindingDetailsServerQueryParameter(bindingDetailsServer)
	return o
}

// SetBindingDetailsServerQueryParameter adds the bindingDetailsServer to the nis collection get params
func (o *NisCollectionGetParams) SetBindingDetailsServerQueryParameter(bindingDetailsServer *string) {
	o.BindingDetailsServerQueryParameter = bindingDetailsServer
}

// WithBindingDetailsStatusCodeQueryParameter adds the bindingDetailsStatusCode to the nis collection get params
func (o *NisCollectionGetParams) WithBindingDetailsStatusCodeQueryParameter(bindingDetailsStatusCode *string) *NisCollectionGetParams {
	o.SetBindingDetailsStatusCodeQueryParameter(bindingDetailsStatusCode)
	return o
}

// SetBindingDetailsStatusCodeQueryParameter adds the bindingDetailsStatusCode to the nis collection get params
func (o *NisCollectionGetParams) SetBindingDetailsStatusCodeQueryParameter(bindingDetailsStatusCode *string) {
	o.BindingDetailsStatusCodeQueryParameter = bindingDetailsStatusCode
}

// WithBindingDetailsStatusMessageQueryParameter adds the bindingDetailsStatusMessage to the nis collection get params
func (o *NisCollectionGetParams) WithBindingDetailsStatusMessageQueryParameter(bindingDetailsStatusMessage *string) *NisCollectionGetParams {
	o.SetBindingDetailsStatusMessageQueryParameter(bindingDetailsStatusMessage)
	return o
}

// SetBindingDetailsStatusMessageQueryParameter adds the bindingDetailsStatusMessage to the nis collection get params
func (o *NisCollectionGetParams) SetBindingDetailsStatusMessageQueryParameter(bindingDetailsStatusMessage *string) {
	o.BindingDetailsStatusMessageQueryParameter = bindingDetailsStatusMessage
}

// WithBoundServersQueryParameter adds the boundServers to the nis collection get params
func (o *NisCollectionGetParams) WithBoundServersQueryParameter(boundServers *string) *NisCollectionGetParams {
	o.SetBoundServersQueryParameter(boundServers)
	return o
}

// SetBoundServersQueryParameter adds the boundServers to the nis collection get params
func (o *NisCollectionGetParams) SetBoundServersQueryParameter(boundServers *string) {
	o.BoundServersQueryParameter = boundServers
}

// WithDomainQueryParameter adds the domain to the nis collection get params
func (o *NisCollectionGetParams) WithDomainQueryParameter(domain *string) *NisCollectionGetParams {
	o.SetDomainQueryParameter(domain)
	return o
}

// SetDomainQueryParameter adds the domain to the nis collection get params
func (o *NisCollectionGetParams) SetDomainQueryParameter(domain *string) {
	o.DomainQueryParameter = domain
}

// WithFieldsQueryParameter adds the fields to the nis collection get params
func (o *NisCollectionGetParams) WithFieldsQueryParameter(fields []string) *NisCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the nis collection get params
func (o *NisCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the nis collection get params
func (o *NisCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *NisCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the nis collection get params
func (o *NisCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the nis collection get params
func (o *NisCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *NisCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the nis collection get params
func (o *NisCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the nis collection get params
func (o *NisCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *NisCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the nis collection get params
func (o *NisCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the nis collection get params
func (o *NisCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *NisCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the nis collection get params
func (o *NisCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithServersQueryParameter adds the servers to the nis collection get params
func (o *NisCollectionGetParams) WithServersQueryParameter(servers *string) *NisCollectionGetParams {
	o.SetServersQueryParameter(servers)
	return o
}

// SetServersQueryParameter adds the servers to the nis collection get params
func (o *NisCollectionGetParams) SetServersQueryParameter(servers *string) {
	o.ServersQueryParameter = servers
}

// WithSVMNameQueryParameter adds the svmName to the nis collection get params
func (o *NisCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *NisCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the nis collection get params
func (o *NisCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the nis collection get params
func (o *NisCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *NisCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the nis collection get params
func (o *NisCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NisCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BindingDetailsServerQueryParameter != nil {

		// query param binding_details.server
		var qrBindingDetailsServer string

		if o.BindingDetailsServerQueryParameter != nil {
			qrBindingDetailsServer = *o.BindingDetailsServerQueryParameter
		}
		qBindingDetailsServer := qrBindingDetailsServer
		if qBindingDetailsServer != "" {

			if err := r.SetQueryParam("binding_details.server", qBindingDetailsServer); err != nil {
				return err
			}
		}
	}

	if o.BindingDetailsStatusCodeQueryParameter != nil {

		// query param binding_details.status.code
		var qrBindingDetailsStatusCode string

		if o.BindingDetailsStatusCodeQueryParameter != nil {
			qrBindingDetailsStatusCode = *o.BindingDetailsStatusCodeQueryParameter
		}
		qBindingDetailsStatusCode := qrBindingDetailsStatusCode
		if qBindingDetailsStatusCode != "" {

			if err := r.SetQueryParam("binding_details.status.code", qBindingDetailsStatusCode); err != nil {
				return err
			}
		}
	}

	if o.BindingDetailsStatusMessageQueryParameter != nil {

		// query param binding_details.status.message
		var qrBindingDetailsStatusMessage string

		if o.BindingDetailsStatusMessageQueryParameter != nil {
			qrBindingDetailsStatusMessage = *o.BindingDetailsStatusMessageQueryParameter
		}
		qBindingDetailsStatusMessage := qrBindingDetailsStatusMessage
		if qBindingDetailsStatusMessage != "" {

			if err := r.SetQueryParam("binding_details.status.message", qBindingDetailsStatusMessage); err != nil {
				return err
			}
		}
	}

	if o.BoundServersQueryParameter != nil {

		// query param bound_servers
		var qrBoundServers string

		if o.BoundServersQueryParameter != nil {
			qrBoundServers = *o.BoundServersQueryParameter
		}
		qBoundServers := qrBoundServers
		if qBoundServers != "" {

			if err := r.SetQueryParam("bound_servers", qBoundServers); err != nil {
				return err
			}
		}
	}

	if o.DomainQueryParameter != nil {

		// query param domain
		var qrDomain string

		if o.DomainQueryParameter != nil {
			qrDomain = *o.DomainQueryParameter
		}
		qDomain := qrDomain
		if qDomain != "" {

			if err := r.SetQueryParam("domain", qDomain); err != nil {
				return err
			}
		}
	}

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

	if o.ServersQueryParameter != nil {

		// query param servers
		var qrServers string

		if o.ServersQueryParameter != nil {
			qrServers = *o.ServersQueryParameter
		}
		qServers := qrServers
		if qServers != "" {

			if err := r.SetQueryParam("servers", qServers); err != nil {
				return err
			}
		}
	}

	if o.SVMNameQueryParameter != nil {

		// query param svm.name
		var qrSvmName string

		if o.SVMNameQueryParameter != nil {
			qrSvmName = *o.SVMNameQueryParameter
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SVMUUIDQueryParameter != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SVMUUIDQueryParameter != nil {
			qrSvmUUID = *o.SVMUUIDQueryParameter
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNisCollectionGet binds the parameter fields
func (o *NisCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamNisCollectionGet binds the parameter order_by
func (o *NisCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
