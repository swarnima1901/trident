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

// NewNetworkIPBgpPeerGroupsGetParams creates a new NetworkIPBgpPeerGroupsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkIPBgpPeerGroupsGetParams() *NetworkIPBgpPeerGroupsGetParams {
	return &NetworkIPBgpPeerGroupsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkIPBgpPeerGroupsGetParamsWithTimeout creates a new NetworkIPBgpPeerGroupsGetParams object
// with the ability to set a timeout on a request.
func NewNetworkIPBgpPeerGroupsGetParamsWithTimeout(timeout time.Duration) *NetworkIPBgpPeerGroupsGetParams {
	return &NetworkIPBgpPeerGroupsGetParams{
		timeout: timeout,
	}
}

// NewNetworkIPBgpPeerGroupsGetParamsWithContext creates a new NetworkIPBgpPeerGroupsGetParams object
// with the ability to set a context for a request.
func NewNetworkIPBgpPeerGroupsGetParamsWithContext(ctx context.Context) *NetworkIPBgpPeerGroupsGetParams {
	return &NetworkIPBgpPeerGroupsGetParams{
		Context: ctx,
	}
}

// NewNetworkIPBgpPeerGroupsGetParamsWithHTTPClient creates a new NetworkIPBgpPeerGroupsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkIPBgpPeerGroupsGetParamsWithHTTPClient(client *http.Client) *NetworkIPBgpPeerGroupsGetParams {
	return &NetworkIPBgpPeerGroupsGetParams{
		HTTPClient: client,
	}
}

/*
NetworkIPBgpPeerGroupsGetParams contains all the parameters to send to the API endpoint

	for the network ip bgp peer groups get operation.

	Typically these are written to a http.Request.
*/
type NetworkIPBgpPeerGroupsGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* IpspaceName.

	   Filter by ipspace.name
	*/
	IpspaceNameQueryParameter *string

	/* IpspaceUUID.

	   Filter by ipspace.uuid
	*/
	IpspaceUUIDQueryParameter *string

	/* LocalInterfaceIPAddress.

	   Filter by local.interface.ip.address
	*/
	LocalInterfaceIPAddressQueryParameter *string

	/* LocalInterfaceName.

	   Filter by local.interface.name
	*/
	LocalInterfaceNameQueryParameter *string

	/* LocalInterfaceUUID.

	   Filter by local.interface.uuid
	*/
	LocalInterfaceUUIDQueryParameter *string

	/* LocalPortName.

	   Filter by local.port.name
	*/
	LocalPortNameQueryParameter *string

	/* LocalPortNodeName.

	   Filter by local.port.node.name
	*/
	LocalPortNodeNameQueryParameter *string

	/* LocalPortUUID.

	   Filter by local.port.uuid
	*/
	LocalPortUUIDQueryParameter *string

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

	/* PeerAddress.

	   Filter by peer.address
	*/
	PeerAddressQueryParameter *string

	/* PeerAsn.

	   Filter by peer.asn
	*/
	PeerAsnQueryParameter *int64

	/* PeerIsNextHop.

	   Filter by peer.is_next_hop
	*/
	PeerIsNextHopQueryParameter *bool

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

	/* State.

	   Filter by state
	*/
	StateQueryParameter *string

	/* UUID.

	   Filter by uuid
	*/
	UUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network ip bgp peer groups get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPBgpPeerGroupsGetParams) WithDefaults() *NetworkIPBgpPeerGroupsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network ip bgp peer groups get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPBgpPeerGroupsGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := NetworkIPBgpPeerGroupsGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithTimeout(timeout time.Duration) *NetworkIPBgpPeerGroupsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithContext(ctx context.Context) *NetworkIPBgpPeerGroupsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithHTTPClient(client *http.Client) *NetworkIPBgpPeerGroupsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithFieldsQueryParameter(fields []string) *NetworkIPBgpPeerGroupsGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIpspaceNameQueryParameter adds the ipspaceName to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithIpspaceNameQueryParameter(ipspaceName *string) *NetworkIPBgpPeerGroupsGetParams {
	o.SetIpspaceNameQueryParameter(ipspaceName)
	return o
}

// SetIpspaceNameQueryParameter adds the ipspaceName to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetIpspaceNameQueryParameter(ipspaceName *string) {
	o.IpspaceNameQueryParameter = ipspaceName
}

// WithIpspaceUUIDQueryParameter adds the ipspaceUUID to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithIpspaceUUIDQueryParameter(ipspaceUUID *string) *NetworkIPBgpPeerGroupsGetParams {
	o.SetIpspaceUUIDQueryParameter(ipspaceUUID)
	return o
}

// SetIpspaceUUIDQueryParameter adds the ipspaceUuid to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetIpspaceUUIDQueryParameter(ipspaceUUID *string) {
	o.IpspaceUUIDQueryParameter = ipspaceUUID
}

// WithLocalInterfaceIPAddressQueryParameter adds the localInterfaceIPAddress to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithLocalInterfaceIPAddressQueryParameter(localInterfaceIPAddress *string) *NetworkIPBgpPeerGroupsGetParams {
	o.SetLocalInterfaceIPAddressQueryParameter(localInterfaceIPAddress)
	return o
}

// SetLocalInterfaceIPAddressQueryParameter adds the localInterfaceIpAddress to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetLocalInterfaceIPAddressQueryParameter(localInterfaceIPAddress *string) {
	o.LocalInterfaceIPAddressQueryParameter = localInterfaceIPAddress
}

// WithLocalInterfaceNameQueryParameter adds the localInterfaceName to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithLocalInterfaceNameQueryParameter(localInterfaceName *string) *NetworkIPBgpPeerGroupsGetParams {
	o.SetLocalInterfaceNameQueryParameter(localInterfaceName)
	return o
}

// SetLocalInterfaceNameQueryParameter adds the localInterfaceName to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetLocalInterfaceNameQueryParameter(localInterfaceName *string) {
	o.LocalInterfaceNameQueryParameter = localInterfaceName
}

// WithLocalInterfaceUUIDQueryParameter adds the localInterfaceUUID to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithLocalInterfaceUUIDQueryParameter(localInterfaceUUID *string) *NetworkIPBgpPeerGroupsGetParams {
	o.SetLocalInterfaceUUIDQueryParameter(localInterfaceUUID)
	return o
}

// SetLocalInterfaceUUIDQueryParameter adds the localInterfaceUuid to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetLocalInterfaceUUIDQueryParameter(localInterfaceUUID *string) {
	o.LocalInterfaceUUIDQueryParameter = localInterfaceUUID
}

// WithLocalPortNameQueryParameter adds the localPortName to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithLocalPortNameQueryParameter(localPortName *string) *NetworkIPBgpPeerGroupsGetParams {
	o.SetLocalPortNameQueryParameter(localPortName)
	return o
}

// SetLocalPortNameQueryParameter adds the localPortName to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetLocalPortNameQueryParameter(localPortName *string) {
	o.LocalPortNameQueryParameter = localPortName
}

// WithLocalPortNodeNameQueryParameter adds the localPortNodeName to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithLocalPortNodeNameQueryParameter(localPortNodeName *string) *NetworkIPBgpPeerGroupsGetParams {
	o.SetLocalPortNodeNameQueryParameter(localPortNodeName)
	return o
}

// SetLocalPortNodeNameQueryParameter adds the localPortNodeName to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetLocalPortNodeNameQueryParameter(localPortNodeName *string) {
	o.LocalPortNodeNameQueryParameter = localPortNodeName
}

// WithLocalPortUUIDQueryParameter adds the localPortUUID to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithLocalPortUUIDQueryParameter(localPortUUID *string) *NetworkIPBgpPeerGroupsGetParams {
	o.SetLocalPortUUIDQueryParameter(localPortUUID)
	return o
}

// SetLocalPortUUIDQueryParameter adds the localPortUuid to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetLocalPortUUIDQueryParameter(localPortUUID *string) {
	o.LocalPortUUIDQueryParameter = localPortUUID
}

// WithMaxRecordsQueryParameter adds the maxRecords to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *NetworkIPBgpPeerGroupsGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNameQueryParameter adds the name to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithNameQueryParameter(name *string) *NetworkIPBgpPeerGroupsGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderByQueryParameter adds the orderBy to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithOrderByQueryParameter(orderBy []string) *NetworkIPBgpPeerGroupsGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithPeerAddressQueryParameter adds the peerAddress to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithPeerAddressQueryParameter(peerAddress *string) *NetworkIPBgpPeerGroupsGetParams {
	o.SetPeerAddressQueryParameter(peerAddress)
	return o
}

// SetPeerAddressQueryParameter adds the peerAddress to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetPeerAddressQueryParameter(peerAddress *string) {
	o.PeerAddressQueryParameter = peerAddress
}

// WithPeerAsnQueryParameter adds the peerAsn to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithPeerAsnQueryParameter(peerAsn *int64) *NetworkIPBgpPeerGroupsGetParams {
	o.SetPeerAsnQueryParameter(peerAsn)
	return o
}

// SetPeerAsnQueryParameter adds the peerAsn to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetPeerAsnQueryParameter(peerAsn *int64) {
	o.PeerAsnQueryParameter = peerAsn
}

// WithPeerIsNextHopQueryParameter adds the peerIsNextHop to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithPeerIsNextHopQueryParameter(peerIsNextHop *bool) *NetworkIPBgpPeerGroupsGetParams {
	o.SetPeerIsNextHopQueryParameter(peerIsNextHop)
	return o
}

// SetPeerIsNextHopQueryParameter adds the peerIsNextHop to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetPeerIsNextHopQueryParameter(peerIsNextHop *bool) {
	o.PeerIsNextHopQueryParameter = peerIsNextHop
}

// WithReturnRecordsQueryParameter adds the returnRecords to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *NetworkIPBgpPeerGroupsGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *NetworkIPBgpPeerGroupsGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithStateQueryParameter adds the state to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithStateQueryParameter(state *string) *NetworkIPBgpPeerGroupsGetParams {
	o.SetStateQueryParameter(state)
	return o
}

// SetStateQueryParameter adds the state to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetStateQueryParameter(state *string) {
	o.StateQueryParameter = state
}

// WithUUIDQueryParameter adds the uuid to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) WithUUIDQueryParameter(uuid *string) *NetworkIPBgpPeerGroupsGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the network ip bgp peer groups get params
func (o *NetworkIPBgpPeerGroupsGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkIPBgpPeerGroupsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IpspaceNameQueryParameter != nil {

		// query param ipspace.name
		var qrIpspaceName string

		if o.IpspaceNameQueryParameter != nil {
			qrIpspaceName = *o.IpspaceNameQueryParameter
		}
		qIpspaceName := qrIpspaceName
		if qIpspaceName != "" {

			if err := r.SetQueryParam("ipspace.name", qIpspaceName); err != nil {
				return err
			}
		}
	}

	if o.IpspaceUUIDQueryParameter != nil {

		// query param ipspace.uuid
		var qrIpspaceUUID string

		if o.IpspaceUUIDQueryParameter != nil {
			qrIpspaceUUID = *o.IpspaceUUIDQueryParameter
		}
		qIpspaceUUID := qrIpspaceUUID
		if qIpspaceUUID != "" {

			if err := r.SetQueryParam("ipspace.uuid", qIpspaceUUID); err != nil {
				return err
			}
		}
	}

	if o.LocalInterfaceIPAddressQueryParameter != nil {

		// query param local.interface.ip.address
		var qrLocalInterfaceIPAddress string

		if o.LocalInterfaceIPAddressQueryParameter != nil {
			qrLocalInterfaceIPAddress = *o.LocalInterfaceIPAddressQueryParameter
		}
		qLocalInterfaceIPAddress := qrLocalInterfaceIPAddress
		if qLocalInterfaceIPAddress != "" {

			if err := r.SetQueryParam("local.interface.ip.address", qLocalInterfaceIPAddress); err != nil {
				return err
			}
		}
	}

	if o.LocalInterfaceNameQueryParameter != nil {

		// query param local.interface.name
		var qrLocalInterfaceName string

		if o.LocalInterfaceNameQueryParameter != nil {
			qrLocalInterfaceName = *o.LocalInterfaceNameQueryParameter
		}
		qLocalInterfaceName := qrLocalInterfaceName
		if qLocalInterfaceName != "" {

			if err := r.SetQueryParam("local.interface.name", qLocalInterfaceName); err != nil {
				return err
			}
		}
	}

	if o.LocalInterfaceUUIDQueryParameter != nil {

		// query param local.interface.uuid
		var qrLocalInterfaceUUID string

		if o.LocalInterfaceUUIDQueryParameter != nil {
			qrLocalInterfaceUUID = *o.LocalInterfaceUUIDQueryParameter
		}
		qLocalInterfaceUUID := qrLocalInterfaceUUID
		if qLocalInterfaceUUID != "" {

			if err := r.SetQueryParam("local.interface.uuid", qLocalInterfaceUUID); err != nil {
				return err
			}
		}
	}

	if o.LocalPortNameQueryParameter != nil {

		// query param local.port.name
		var qrLocalPortName string

		if o.LocalPortNameQueryParameter != nil {
			qrLocalPortName = *o.LocalPortNameQueryParameter
		}
		qLocalPortName := qrLocalPortName
		if qLocalPortName != "" {

			if err := r.SetQueryParam("local.port.name", qLocalPortName); err != nil {
				return err
			}
		}
	}

	if o.LocalPortNodeNameQueryParameter != nil {

		// query param local.port.node.name
		var qrLocalPortNodeName string

		if o.LocalPortNodeNameQueryParameter != nil {
			qrLocalPortNodeName = *o.LocalPortNodeNameQueryParameter
		}
		qLocalPortNodeName := qrLocalPortNodeName
		if qLocalPortNodeName != "" {

			if err := r.SetQueryParam("local.port.node.name", qLocalPortNodeName); err != nil {
				return err
			}
		}
	}

	if o.LocalPortUUIDQueryParameter != nil {

		// query param local.port.uuid
		var qrLocalPortUUID string

		if o.LocalPortUUIDQueryParameter != nil {
			qrLocalPortUUID = *o.LocalPortUUIDQueryParameter
		}
		qLocalPortUUID := qrLocalPortUUID
		if qLocalPortUUID != "" {

			if err := r.SetQueryParam("local.port.uuid", qLocalPortUUID); err != nil {
				return err
			}
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

	if o.PeerAddressQueryParameter != nil {

		// query param peer.address
		var qrPeerAddress string

		if o.PeerAddressQueryParameter != nil {
			qrPeerAddress = *o.PeerAddressQueryParameter
		}
		qPeerAddress := qrPeerAddress
		if qPeerAddress != "" {

			if err := r.SetQueryParam("peer.address", qPeerAddress); err != nil {
				return err
			}
		}
	}

	if o.PeerAsnQueryParameter != nil {

		// query param peer.asn
		var qrPeerAsn int64

		if o.PeerAsnQueryParameter != nil {
			qrPeerAsn = *o.PeerAsnQueryParameter
		}
		qPeerAsn := swag.FormatInt64(qrPeerAsn)
		if qPeerAsn != "" {

			if err := r.SetQueryParam("peer.asn", qPeerAsn); err != nil {
				return err
			}
		}
	}

	if o.PeerIsNextHopQueryParameter != nil {

		// query param peer.is_next_hop
		var qrPeerIsNextHop bool

		if o.PeerIsNextHopQueryParameter != nil {
			qrPeerIsNextHop = *o.PeerIsNextHopQueryParameter
		}
		qPeerIsNextHop := swag.FormatBool(qrPeerIsNextHop)
		if qPeerIsNextHop != "" {

			if err := r.SetQueryParam("peer.is_next_hop", qPeerIsNextHop); err != nil {
				return err
			}
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

	if o.StateQueryParameter != nil {

		// query param state
		var qrState string

		if o.StateQueryParameter != nil {
			qrState = *o.StateQueryParameter
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	if o.UUIDQueryParameter != nil {

		// query param uuid
		var qrUUID string

		if o.UUIDQueryParameter != nil {
			qrUUID = *o.UUIDQueryParameter
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNetworkIPBgpPeerGroupsGet binds the parameter fields
func (o *NetworkIPBgpPeerGroupsGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamNetworkIPBgpPeerGroupsGet binds the parameter order_by
func (o *NetworkIPBgpPeerGroupsGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
