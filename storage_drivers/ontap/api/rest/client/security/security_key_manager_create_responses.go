// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SecurityKeyManagerCreateReader is a Reader for the SecurityKeyManagerCreate structure.
type SecurityKeyManagerCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityKeyManagerCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSecurityKeyManagerCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityKeyManagerCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityKeyManagerCreateCreated creates a SecurityKeyManagerCreateCreated with default headers values
func NewSecurityKeyManagerCreateCreated() *SecurityKeyManagerCreateCreated {
	return &SecurityKeyManagerCreateCreated{}
}

/*
SecurityKeyManagerCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SecurityKeyManagerCreateCreated struct {
	Payload *models.SecurityKeyManagerResponse
}

// IsSuccess returns true when this security key manager create created response has a 2xx status code
func (o *SecurityKeyManagerCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security key manager create created response has a 3xx status code
func (o *SecurityKeyManagerCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security key manager create created response has a 4xx status code
func (o *SecurityKeyManagerCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this security key manager create created response has a 5xx status code
func (o *SecurityKeyManagerCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this security key manager create created response a status code equal to that given
func (o *SecurityKeyManagerCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *SecurityKeyManagerCreateCreated) Error() string {
	return fmt.Sprintf("[POST /security/key-managers][%d] securityKeyManagerCreateCreated  %+v", 201, o.Payload)
}

func (o *SecurityKeyManagerCreateCreated) String() string {
	return fmt.Sprintf("[POST /security/key-managers][%d] securityKeyManagerCreateCreated  %+v", 201, o.Payload)
}

func (o *SecurityKeyManagerCreateCreated) GetPayload() *models.SecurityKeyManagerResponse {
	return o.Payload
}

func (o *SecurityKeyManagerCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityKeyManagerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityKeyManagerCreateDefault creates a SecurityKeyManagerCreateDefault with default headers values
func NewSecurityKeyManagerCreateDefault(code int) *SecurityKeyManagerCreateDefault {
	return &SecurityKeyManagerCreateDefault{
		_statusCode: code,
	}
}

/*
	SecurityKeyManagerCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 65536038 | A maximum of 4 active primary key servers are allowed. |
| 65536214 | Failed to generate cluster key encryption key. |
| 65536216 | Failed to add cluster key encryption key. |
| 65536310 | Failed to setup the Onboard Key Manager because the MetroCluster peer is unhealthy. |
| 65536341 | Failed to setup the Onboard Key Manager because the MetroCluster peer is unhealthy. |
| 65536508 | The platform does not support data at rest encryption. |
| 65536821 | The certificate is not installed. |
| 65536822 | Multitenant key management is not supported in the current cluster version. |
| 65536823 | The SVM has key manager already configured. |
| 65536824 | Multitenant key management is not supported in MetroCluster configurations. |
| 65536834 | Failed to get existing key-server details for the SVM. |
| 65536852 | Failed to query supported KMIP protocol versions. |
| 65536870 | Key management servers already configured. |
| 65536871 | Duplicate key management servers exist. |
| 65536876 | External key management requires client and server CA certificates installed and with one or more key servers provided. |
| 65536878 | External key management cannot be configured as one or more volume encryption keys of the SVM are stored in cluster key management server. |
| 65536895 | External key manager cannnot be configured since this cluster is part of a MetroCluster configuration and the partner site of this MetroCluster configuration has Onboard Key Manager configured. |
| 65536900 | The Onboard Key Manager cannot be configured because this cluster is part of a MetroCluster configuration and the partner site has the external key manager configured. |
| 65536903 | The Onboard Key Manager has failed to configure on some nodes in the cluster. Use the CLI to sync the Onboard Key Manager configuration on failed nodes. |
| 65536906 | The Onboard Key Manager has already been configured at the partner site. Use the CLI to sync the Onboard Key Manager with the same passphrase. |
| 65536907 | The Onboard Key Manager is already configured. Use the CLI to sync any nodes with the Onboard Key Manager configuration. |
| 65536916 | The Onboard Key Manager is only supported for an admin SVM. |
| 65536920 | The Onboard Key Manager passphrase length is incorrect. |
| 65537240 | The Onboard Key Manager passphrase must be provided when performing a POST/synchronize operation. |
| 65537241 | The Onboard Key Manager existing_passphrase must not be provided when performing a POST/synchronize operation. |
| 65537244 | Unable to sync/create Onboard Key Manager on the local cluster; Onboard Key Manager is already configured on the cluster. |
| 65537245 | Unable to sync/create Onboard Key Manager on the local cluster; Onboard Key Manager is not configured on the partner cluster. |
| 65537246 | Unable to sync/create Onboard Key Manager on local cluster. This cluster is not part of a MetroCluster configuration. |
| 66060338 | Failed to establish secure connection for a key management server due to incorrect server_ca certificates. |
| 66060339 | Failed to establish secure connection for a key management server due to incorrect client certificates. |
| 66060340 | Failed to establish secure connection for a key management server due to Cryptsoft error. |
*/
type SecurityKeyManagerCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the security key manager create default response
func (o *SecurityKeyManagerCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this security key manager create default response has a 2xx status code
func (o *SecurityKeyManagerCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security key manager create default response has a 3xx status code
func (o *SecurityKeyManagerCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security key manager create default response has a 4xx status code
func (o *SecurityKeyManagerCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security key manager create default response has a 5xx status code
func (o *SecurityKeyManagerCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security key manager create default response a status code equal to that given
func (o *SecurityKeyManagerCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SecurityKeyManagerCreateDefault) Error() string {
	return fmt.Sprintf("[POST /security/key-managers][%d] security_key_manager_create default  %+v", o._statusCode, o.Payload)
}

func (o *SecurityKeyManagerCreateDefault) String() string {
	return fmt.Sprintf("[POST /security/key-managers][%d] security_key_manager_create default  %+v", o._statusCode, o.Payload)
}

func (o *SecurityKeyManagerCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityKeyManagerCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
