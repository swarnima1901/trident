// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewCounterTableCollectionGetParams creates a new CounterTableCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCounterTableCollectionGetParams() *CounterTableCollectionGetParams {
	return &CounterTableCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCounterTableCollectionGetParamsWithTimeout creates a new CounterTableCollectionGetParams object
// with the ability to set a timeout on a request.
func NewCounterTableCollectionGetParamsWithTimeout(timeout time.Duration) *CounterTableCollectionGetParams {
	return &CounterTableCollectionGetParams{
		timeout: timeout,
	}
}

// NewCounterTableCollectionGetParamsWithContext creates a new CounterTableCollectionGetParams object
// with the ability to set a context for a request.
func NewCounterTableCollectionGetParamsWithContext(ctx context.Context) *CounterTableCollectionGetParams {
	return &CounterTableCollectionGetParams{
		Context: ctx,
	}
}

// NewCounterTableCollectionGetParamsWithHTTPClient creates a new CounterTableCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCounterTableCollectionGetParamsWithHTTPClient(client *http.Client) *CounterTableCollectionGetParams {
	return &CounterTableCollectionGetParams{
		HTTPClient: client,
	}
}

/*
CounterTableCollectionGetParams contains all the parameters to send to the API endpoint

	for the counter table collection get operation.

	Typically these are written to a http.Request.
*/
type CounterTableCollectionGetParams struct {

	/* CounterSchemasDenominatorName.

	   Filter by counter_schemas.denominator.name
	*/
	CounterSchemasDenominatorNameQueryParameter *string

	/* CounterSchemasDescription.

	   Filter by counter_schemas.description
	*/
	CounterSchemasDescriptionQueryParameter *string

	/* CounterSchemasName.

	   Filter by counter_schemas.name
	*/
	CounterSchemasNameQueryParameter *string

	/* CounterSchemasType.

	   Filter by counter_schemas.type
	*/
	CounterSchemasTypeQueryParameter *string

	/* CounterSchemasUnit.

	   Filter by counter_schemas.unit
	*/
	CounterSchemasUnitQueryParameter *string

	/* Description.

	   Filter by description
	*/
	DescriptionQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the counter table collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CounterTableCollectionGetParams) WithDefaults() *CounterTableCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the counter table collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CounterTableCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := CounterTableCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the counter table collection get params
func (o *CounterTableCollectionGetParams) WithTimeout(timeout time.Duration) *CounterTableCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the counter table collection get params
func (o *CounterTableCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the counter table collection get params
func (o *CounterTableCollectionGetParams) WithContext(ctx context.Context) *CounterTableCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the counter table collection get params
func (o *CounterTableCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the counter table collection get params
func (o *CounterTableCollectionGetParams) WithHTTPClient(client *http.Client) *CounterTableCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the counter table collection get params
func (o *CounterTableCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCounterSchemasDenominatorNameQueryParameter adds the counterSchemasDenominatorName to the counter table collection get params
func (o *CounterTableCollectionGetParams) WithCounterSchemasDenominatorNameQueryParameter(counterSchemasDenominatorName *string) *CounterTableCollectionGetParams {
	o.SetCounterSchemasDenominatorNameQueryParameter(counterSchemasDenominatorName)
	return o
}

// SetCounterSchemasDenominatorNameQueryParameter adds the counterSchemasDenominatorName to the counter table collection get params
func (o *CounterTableCollectionGetParams) SetCounterSchemasDenominatorNameQueryParameter(counterSchemasDenominatorName *string) {
	o.CounterSchemasDenominatorNameQueryParameter = counterSchemasDenominatorName
}

// WithCounterSchemasDescriptionQueryParameter adds the counterSchemasDescription to the counter table collection get params
func (o *CounterTableCollectionGetParams) WithCounterSchemasDescriptionQueryParameter(counterSchemasDescription *string) *CounterTableCollectionGetParams {
	o.SetCounterSchemasDescriptionQueryParameter(counterSchemasDescription)
	return o
}

// SetCounterSchemasDescriptionQueryParameter adds the counterSchemasDescription to the counter table collection get params
func (o *CounterTableCollectionGetParams) SetCounterSchemasDescriptionQueryParameter(counterSchemasDescription *string) {
	o.CounterSchemasDescriptionQueryParameter = counterSchemasDescription
}

// WithCounterSchemasNameQueryParameter adds the counterSchemasName to the counter table collection get params
func (o *CounterTableCollectionGetParams) WithCounterSchemasNameQueryParameter(counterSchemasName *string) *CounterTableCollectionGetParams {
	o.SetCounterSchemasNameQueryParameter(counterSchemasName)
	return o
}

// SetCounterSchemasNameQueryParameter adds the counterSchemasName to the counter table collection get params
func (o *CounterTableCollectionGetParams) SetCounterSchemasNameQueryParameter(counterSchemasName *string) {
	o.CounterSchemasNameQueryParameter = counterSchemasName
}

// WithCounterSchemasTypeQueryParameter adds the counterSchemasType to the counter table collection get params
func (o *CounterTableCollectionGetParams) WithCounterSchemasTypeQueryParameter(counterSchemasType *string) *CounterTableCollectionGetParams {
	o.SetCounterSchemasTypeQueryParameter(counterSchemasType)
	return o
}

// SetCounterSchemasTypeQueryParameter adds the counterSchemasType to the counter table collection get params
func (o *CounterTableCollectionGetParams) SetCounterSchemasTypeQueryParameter(counterSchemasType *string) {
	o.CounterSchemasTypeQueryParameter = counterSchemasType
}

// WithCounterSchemasUnitQueryParameter adds the counterSchemasUnit to the counter table collection get params
func (o *CounterTableCollectionGetParams) WithCounterSchemasUnitQueryParameter(counterSchemasUnit *string) *CounterTableCollectionGetParams {
	o.SetCounterSchemasUnitQueryParameter(counterSchemasUnit)
	return o
}

// SetCounterSchemasUnitQueryParameter adds the counterSchemasUnit to the counter table collection get params
func (o *CounterTableCollectionGetParams) SetCounterSchemasUnitQueryParameter(counterSchemasUnit *string) {
	o.CounterSchemasUnitQueryParameter = counterSchemasUnit
}

// WithDescriptionQueryParameter adds the description to the counter table collection get params
func (o *CounterTableCollectionGetParams) WithDescriptionQueryParameter(description *string) *CounterTableCollectionGetParams {
	o.SetDescriptionQueryParameter(description)
	return o
}

// SetDescriptionQueryParameter adds the description to the counter table collection get params
func (o *CounterTableCollectionGetParams) SetDescriptionQueryParameter(description *string) {
	o.DescriptionQueryParameter = description
}

// WithFieldsQueryParameter adds the fields to the counter table collection get params
func (o *CounterTableCollectionGetParams) WithFieldsQueryParameter(fields []string) *CounterTableCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the counter table collection get params
func (o *CounterTableCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the counter table collection get params
func (o *CounterTableCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *CounterTableCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the counter table collection get params
func (o *CounterTableCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNameQueryParameter adds the name to the counter table collection get params
func (o *CounterTableCollectionGetParams) WithNameQueryParameter(name *string) *CounterTableCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the counter table collection get params
func (o *CounterTableCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderByQueryParameter adds the orderBy to the counter table collection get params
func (o *CounterTableCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *CounterTableCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the counter table collection get params
func (o *CounterTableCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the counter table collection get params
func (o *CounterTableCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *CounterTableCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the counter table collection get params
func (o *CounterTableCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the counter table collection get params
func (o *CounterTableCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *CounterTableCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the counter table collection get params
func (o *CounterTableCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *CounterTableCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CounterSchemasDenominatorNameQueryParameter != nil {

		// query param counter_schemas.denominator.name
		var qrCounterSchemasDenominatorName string

		if o.CounterSchemasDenominatorNameQueryParameter != nil {
			qrCounterSchemasDenominatorName = *o.CounterSchemasDenominatorNameQueryParameter
		}
		qCounterSchemasDenominatorName := qrCounterSchemasDenominatorName
		if qCounterSchemasDenominatorName != "" {

			if err := r.SetQueryParam("counter_schemas.denominator.name", qCounterSchemasDenominatorName); err != nil {
				return err
			}
		}
	}

	if o.CounterSchemasDescriptionQueryParameter != nil {

		// query param counter_schemas.description
		var qrCounterSchemasDescription string

		if o.CounterSchemasDescriptionQueryParameter != nil {
			qrCounterSchemasDescription = *o.CounterSchemasDescriptionQueryParameter
		}
		qCounterSchemasDescription := qrCounterSchemasDescription
		if qCounterSchemasDescription != "" {

			if err := r.SetQueryParam("counter_schemas.description", qCounterSchemasDescription); err != nil {
				return err
			}
		}
	}

	if o.CounterSchemasNameQueryParameter != nil {

		// query param counter_schemas.name
		var qrCounterSchemasName string

		if o.CounterSchemasNameQueryParameter != nil {
			qrCounterSchemasName = *o.CounterSchemasNameQueryParameter
		}
		qCounterSchemasName := qrCounterSchemasName
		if qCounterSchemasName != "" {

			if err := r.SetQueryParam("counter_schemas.name", qCounterSchemasName); err != nil {
				return err
			}
		}
	}

	if o.CounterSchemasTypeQueryParameter != nil {

		// query param counter_schemas.type
		var qrCounterSchemasType string

		if o.CounterSchemasTypeQueryParameter != nil {
			qrCounterSchemasType = *o.CounterSchemasTypeQueryParameter
		}
		qCounterSchemasType := qrCounterSchemasType
		if qCounterSchemasType != "" {

			if err := r.SetQueryParam("counter_schemas.type", qCounterSchemasType); err != nil {
				return err
			}
		}
	}

	if o.CounterSchemasUnitQueryParameter != nil {

		// query param counter_schemas.unit
		var qrCounterSchemasUnit string

		if o.CounterSchemasUnitQueryParameter != nil {
			qrCounterSchemasUnit = *o.CounterSchemasUnitQueryParameter
		}
		qCounterSchemasUnit := qrCounterSchemasUnit
		if qCounterSchemasUnit != "" {

			if err := r.SetQueryParam("counter_schemas.unit", qCounterSchemasUnit); err != nil {
				return err
			}
		}
	}

	if o.DescriptionQueryParameter != nil {

		// query param description
		var qrDescription string

		if o.DescriptionQueryParameter != nil {
			qrDescription = *o.DescriptionQueryParameter
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCounterTableCollectionGet binds the parameter fields
func (o *CounterTableCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamCounterTableCollectionGet binds the parameter order_by
func (o *CounterTableCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
