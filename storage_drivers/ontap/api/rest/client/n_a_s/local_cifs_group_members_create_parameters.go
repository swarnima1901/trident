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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewLocalCifsGroupMembersCreateParams creates a new LocalCifsGroupMembersCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocalCifsGroupMembersCreateParams() *LocalCifsGroupMembersCreateParams {
	return &LocalCifsGroupMembersCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocalCifsGroupMembersCreateParamsWithTimeout creates a new LocalCifsGroupMembersCreateParams object
// with the ability to set a timeout on a request.
func NewLocalCifsGroupMembersCreateParamsWithTimeout(timeout time.Duration) *LocalCifsGroupMembersCreateParams {
	return &LocalCifsGroupMembersCreateParams{
		timeout: timeout,
	}
}

// NewLocalCifsGroupMembersCreateParamsWithContext creates a new LocalCifsGroupMembersCreateParams object
// with the ability to set a context for a request.
func NewLocalCifsGroupMembersCreateParamsWithContext(ctx context.Context) *LocalCifsGroupMembersCreateParams {
	return &LocalCifsGroupMembersCreateParams{
		Context: ctx,
	}
}

// NewLocalCifsGroupMembersCreateParamsWithHTTPClient creates a new LocalCifsGroupMembersCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocalCifsGroupMembersCreateParamsWithHTTPClient(client *http.Client) *LocalCifsGroupMembersCreateParams {
	return &LocalCifsGroupMembersCreateParams{
		HTTPClient: client,
	}
}

/*
LocalCifsGroupMembersCreateParams contains all the parameters to send to the API endpoint

	for the local cifs group members create operation.

	Typically these are written to a http.Request.
*/
type LocalCifsGroupMembersCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.LocalCifsGroupMembers

	/* LocalCifsGroupSid.

	   Local group SID
	*/
	LocalCifsGroupSIDPathParameter string

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the local cifs group members create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocalCifsGroupMembersCreateParams) WithDefaults() *LocalCifsGroupMembersCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the local cifs group members create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocalCifsGroupMembersCreateParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)
	)

	val := LocalCifsGroupMembersCreateParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the local cifs group members create params
func (o *LocalCifsGroupMembersCreateParams) WithTimeout(timeout time.Duration) *LocalCifsGroupMembersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the local cifs group members create params
func (o *LocalCifsGroupMembersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the local cifs group members create params
func (o *LocalCifsGroupMembersCreateParams) WithContext(ctx context.Context) *LocalCifsGroupMembersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the local cifs group members create params
func (o *LocalCifsGroupMembersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the local cifs group members create params
func (o *LocalCifsGroupMembersCreateParams) WithHTTPClient(client *http.Client) *LocalCifsGroupMembersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the local cifs group members create params
func (o *LocalCifsGroupMembersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the local cifs group members create params
func (o *LocalCifsGroupMembersCreateParams) WithInfo(info *models.LocalCifsGroupMembers) *LocalCifsGroupMembersCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the local cifs group members create params
func (o *LocalCifsGroupMembersCreateParams) SetInfo(info *models.LocalCifsGroupMembers) {
	o.Info = info
}

// WithLocalCifsGroupSIDPathParameter adds the localCifsGroupSid to the local cifs group members create params
func (o *LocalCifsGroupMembersCreateParams) WithLocalCifsGroupSIDPathParameter(localCifsGroupSid string) *LocalCifsGroupMembersCreateParams {
	o.SetLocalCifsGroupSIDPathParameter(localCifsGroupSid)
	return o
}

// SetLocalCifsGroupSIDPathParameter adds the localCifsGroupSid to the local cifs group members create params
func (o *LocalCifsGroupMembersCreateParams) SetLocalCifsGroupSIDPathParameter(localCifsGroupSid string) {
	o.LocalCifsGroupSIDPathParameter = localCifsGroupSid
}

// WithReturnRecordsQueryParameter adds the returnRecords to the local cifs group members create params
func (o *LocalCifsGroupMembersCreateParams) WithReturnRecordsQueryParameter(returnRecords *bool) *LocalCifsGroupMembersCreateParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the local cifs group members create params
func (o *LocalCifsGroupMembersCreateParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithSVMUUIDPathParameter adds the svmUUID to the local cifs group members create params
func (o *LocalCifsGroupMembersCreateParams) WithSVMUUIDPathParameter(svmUUID string) *LocalCifsGroupMembersCreateParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the local cifs group members create params
func (o *LocalCifsGroupMembersCreateParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *LocalCifsGroupMembersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param local_cifs_group.sid
	if err := r.SetPathParam("local_cifs_group.sid", o.LocalCifsGroupSIDPathParameter); err != nil {
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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
