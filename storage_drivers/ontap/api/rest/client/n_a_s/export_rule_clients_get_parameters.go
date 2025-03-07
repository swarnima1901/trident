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
	"github.com/go-openapi/swag"
)

// NewExportRuleClientsGetParams creates a new ExportRuleClientsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExportRuleClientsGetParams() *ExportRuleClientsGetParams {
	return &ExportRuleClientsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExportRuleClientsGetParamsWithTimeout creates a new ExportRuleClientsGetParams object
// with the ability to set a timeout on a request.
func NewExportRuleClientsGetParamsWithTimeout(timeout time.Duration) *ExportRuleClientsGetParams {
	return &ExportRuleClientsGetParams{
		timeout: timeout,
	}
}

// NewExportRuleClientsGetParamsWithContext creates a new ExportRuleClientsGetParams object
// with the ability to set a context for a request.
func NewExportRuleClientsGetParamsWithContext(ctx context.Context) *ExportRuleClientsGetParams {
	return &ExportRuleClientsGetParams{
		Context: ctx,
	}
}

// NewExportRuleClientsGetParamsWithHTTPClient creates a new ExportRuleClientsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewExportRuleClientsGetParamsWithHTTPClient(client *http.Client) *ExportRuleClientsGetParams {
	return &ExportRuleClientsGetParams{
		HTTPClient: client,
	}
}

/*
ExportRuleClientsGetParams contains all the parameters to send to the API endpoint

	for the export rule clients get operation.

	Typically these are written to a http.Request.
*/
type ExportRuleClientsGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Index.

	   Export Rule Index
	*/
	IndexPathParameter int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* PolicyID.

	   Export Policy ID
	*/
	PolicyIDPathParameter int64

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

// WithDefaults hydrates default values in the export rule clients get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportRuleClientsGetParams) WithDefaults() *ExportRuleClientsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the export rule clients get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportRuleClientsGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := ExportRuleClientsGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the export rule clients get params
func (o *ExportRuleClientsGetParams) WithTimeout(timeout time.Duration) *ExportRuleClientsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export rule clients get params
func (o *ExportRuleClientsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export rule clients get params
func (o *ExportRuleClientsGetParams) WithContext(ctx context.Context) *ExportRuleClientsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export rule clients get params
func (o *ExportRuleClientsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export rule clients get params
func (o *ExportRuleClientsGetParams) WithHTTPClient(client *http.Client) *ExportRuleClientsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export rule clients get params
func (o *ExportRuleClientsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the export rule clients get params
func (o *ExportRuleClientsGetParams) WithFieldsQueryParameter(fields []string) *ExportRuleClientsGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the export rule clients get params
func (o *ExportRuleClientsGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIndexPathParameter adds the index to the export rule clients get params
func (o *ExportRuleClientsGetParams) WithIndexPathParameter(index int64) *ExportRuleClientsGetParams {
	o.SetIndexPathParameter(index)
	return o
}

// SetIndexPathParameter adds the index to the export rule clients get params
func (o *ExportRuleClientsGetParams) SetIndexPathParameter(index int64) {
	o.IndexPathParameter = index
}

// WithMaxRecordsQueryParameter adds the maxRecords to the export rule clients get params
func (o *ExportRuleClientsGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *ExportRuleClientsGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the export rule clients get params
func (o *ExportRuleClientsGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the export rule clients get params
func (o *ExportRuleClientsGetParams) WithOrderByQueryParameter(orderBy []string) *ExportRuleClientsGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the export rule clients get params
func (o *ExportRuleClientsGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithPolicyIDPathParameter adds the policyID to the export rule clients get params
func (o *ExportRuleClientsGetParams) WithPolicyIDPathParameter(policyID int64) *ExportRuleClientsGetParams {
	o.SetPolicyIDPathParameter(policyID)
	return o
}

// SetPolicyIDPathParameter adds the policyId to the export rule clients get params
func (o *ExportRuleClientsGetParams) SetPolicyIDPathParameter(policyID int64) {
	o.PolicyIDPathParameter = policyID
}

// WithReturnRecordsQueryParameter adds the returnRecords to the export rule clients get params
func (o *ExportRuleClientsGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *ExportRuleClientsGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the export rule clients get params
func (o *ExportRuleClientsGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the export rule clients get params
func (o *ExportRuleClientsGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *ExportRuleClientsGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the export rule clients get params
func (o *ExportRuleClientsGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *ExportRuleClientsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("index", swag.FormatInt64(o.IndexPathParameter)); err != nil {
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

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	// path param policy.id
	if err := r.SetPathParam("policy.id", swag.FormatInt64(o.PolicyIDPathParameter)); err != nil {
		return err
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

// bindParamExportRuleClientsGet binds the parameter fields
func (o *ExportRuleClientsGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamExportRuleClientsGet binds the parameter order_by
func (o *ExportRuleClientsGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
