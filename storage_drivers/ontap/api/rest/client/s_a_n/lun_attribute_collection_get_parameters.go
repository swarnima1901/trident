// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewLunAttributeCollectionGetParams creates a new LunAttributeCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLunAttributeCollectionGetParams() *LunAttributeCollectionGetParams {
	return &LunAttributeCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLunAttributeCollectionGetParamsWithTimeout creates a new LunAttributeCollectionGetParams object
// with the ability to set a timeout on a request.
func NewLunAttributeCollectionGetParamsWithTimeout(timeout time.Duration) *LunAttributeCollectionGetParams {
	return &LunAttributeCollectionGetParams{
		timeout: timeout,
	}
}

// NewLunAttributeCollectionGetParamsWithContext creates a new LunAttributeCollectionGetParams object
// with the ability to set a context for a request.
func NewLunAttributeCollectionGetParamsWithContext(ctx context.Context) *LunAttributeCollectionGetParams {
	return &LunAttributeCollectionGetParams{
		Context: ctx,
	}
}

// NewLunAttributeCollectionGetParamsWithHTTPClient creates a new LunAttributeCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewLunAttributeCollectionGetParamsWithHTTPClient(client *http.Client) *LunAttributeCollectionGetParams {
	return &LunAttributeCollectionGetParams{
		HTTPClient: client,
	}
}

/*
LunAttributeCollectionGetParams contains all the parameters to send to the API endpoint

	for the lun attribute collection get operation.

	Typically these are written to a http.Request.
*/
type LunAttributeCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* LunUUID.

	   The unique identifier of the LUN.

	*/
	LunUUIDPathParameter string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

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

	/* Value.

	   Filter by value
	*/
	ValueQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lun attribute collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunAttributeCollectionGetParams) WithDefaults() *LunAttributeCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lun attribute collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunAttributeCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := LunAttributeCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) WithTimeout(timeout time.Duration) *LunAttributeCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) WithContext(ctx context.Context) *LunAttributeCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) WithHTTPClient(client *http.Client) *LunAttributeCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) WithFieldsQueryParameter(fields []string) *LunAttributeCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithLunUUIDPathParameter adds the lunUUID to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) WithLunUUIDPathParameter(lunUUID string) *LunAttributeCollectionGetParams {
	o.SetLunUUIDPathParameter(lunUUID)
	return o
}

// SetLunUUIDPathParameter adds the lunUuid to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) SetLunUUIDPathParameter(lunUUID string) {
	o.LunUUIDPathParameter = lunUUID
}

// WithMaxRecordsQueryParameter adds the maxRecords to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *LunAttributeCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNameQueryParameter adds the name to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) WithNameQueryParameter(name *string) *LunAttributeCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderByQueryParameter adds the orderBy to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *LunAttributeCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *LunAttributeCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *LunAttributeCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithValueQueryParameter adds the value to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) WithValueQueryParameter(value *string) *LunAttributeCollectionGetParams {
	o.SetValueQueryParameter(value)
	return o
}

// SetValueQueryParameter adds the value to the lun attribute collection get params
func (o *LunAttributeCollectionGetParams) SetValueQueryParameter(value *string) {
	o.ValueQueryParameter = value
}

// WriteToRequest writes these params to a swagger request
func (o *LunAttributeCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param lun.uuid
	if err := r.SetPathParam("lun.uuid", o.LunUUIDPathParameter); err != nil {
		return err
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

	if o.NameQueryParameter != nil {

		// query param name
		var qrName string

		if o.NameQueryParameter != nil {
			qrName = *o.NameQueryParameter
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
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

	if o.ValueQueryParameter != nil {

		// query param value
		var qrValue string

		if o.ValueQueryParameter != nil {
			qrValue = *o.ValueQueryParameter
		}
		qValue := qrValue
		if qValue != "" {

			if err := r.SetQueryParam("value", qValue); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamLunAttributeCollectionGet binds the parameter fields
func (o *LunAttributeCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamLunAttributeCollectionGet binds the parameter order_by
func (o *LunAttributeCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
