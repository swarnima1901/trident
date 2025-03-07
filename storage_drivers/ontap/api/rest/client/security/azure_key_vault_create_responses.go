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

// AzureKeyVaultCreateReader is a Reader for the AzureKeyVaultCreate structure.
type AzureKeyVaultCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AzureKeyVaultCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAzureKeyVaultCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAzureKeyVaultCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAzureKeyVaultCreateCreated creates a AzureKeyVaultCreateCreated with default headers values
func NewAzureKeyVaultCreateCreated() *AzureKeyVaultCreateCreated {
	return &AzureKeyVaultCreateCreated{}
}

/*
AzureKeyVaultCreateCreated describes a response with status code 201, with default header values.

Created
*/
type AzureKeyVaultCreateCreated struct {
	Payload *models.AzureKeyVaultResponse
}

// IsSuccess returns true when this azure key vault create created response has a 2xx status code
func (o *AzureKeyVaultCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this azure key vault create created response has a 3xx status code
func (o *AzureKeyVaultCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this azure key vault create created response has a 4xx status code
func (o *AzureKeyVaultCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this azure key vault create created response has a 5xx status code
func (o *AzureKeyVaultCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this azure key vault create created response a status code equal to that given
func (o *AzureKeyVaultCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *AzureKeyVaultCreateCreated) Error() string {
	return fmt.Sprintf("[POST /security/azure-key-vaults][%d] azureKeyVaultCreateCreated  %+v", 201, o.Payload)
}

func (o *AzureKeyVaultCreateCreated) String() string {
	return fmt.Sprintf("[POST /security/azure-key-vaults][%d] azureKeyVaultCreateCreated  %+v", 201, o.Payload)
}

func (o *AzureKeyVaultCreateCreated) GetPayload() *models.AzureKeyVaultResponse {
	return o.Payload
}

func (o *AzureKeyVaultCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureKeyVaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAzureKeyVaultCreateDefault creates a AzureKeyVaultCreateDefault with default headers values
func NewAzureKeyVaultCreateDefault(code int) *AzureKeyVaultCreateDefault {
	return &AzureKeyVaultCreateDefault{
		_statusCode: code,
	}
}

/*
	AzureKeyVaultCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 3735553 | Failed to create self-signed certificate. |
| 3735664 | The specified key size is not supported in FIPS mode. |
| 3735665 | The specified hash function is not supported in FIPS mode. |
| 3735700 | The specified key size is not supported. |
| 52559972 | The certificates start date is later than the current date. |
| 65537500 | A key manager has already been configured for this SVM. |
| 65537504 | Internal error. Failed to store configuration in internal database. |
| 65537505 | One or more volume encryption keys of the given SVM are stored on a key manager configured for the admin SVM. |
| 65537506 | AKV is not supported in MetroCluster configurations. |
| 65537512 | AKV cannot be configured for the given SVM as not all nodes in the cluster can enable the Azure Key Vault feature. |
| 65537514 | Failed to check if the Azure Key Vault feature is enabled. |
| 65537518 | Failed to find an interface with Cluster role. |
| 65537523 | Invalid key ID format. Example key ID format\":" "https://mykeyvault.vault.azure.net/keys/key1". |
| 65537526 | Failed to enable Azure Key Vault feature. |
| 65537567 | No authentication method provided. |
| 65537573 | Invalid client certificate. |
*/
type AzureKeyVaultCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the azure key vault create default response
func (o *AzureKeyVaultCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this azure key vault create default response has a 2xx status code
func (o *AzureKeyVaultCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this azure key vault create default response has a 3xx status code
func (o *AzureKeyVaultCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this azure key vault create default response has a 4xx status code
func (o *AzureKeyVaultCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this azure key vault create default response has a 5xx status code
func (o *AzureKeyVaultCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this azure key vault create default response a status code equal to that given
func (o *AzureKeyVaultCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *AzureKeyVaultCreateDefault) Error() string {
	return fmt.Sprintf("[POST /security/azure-key-vaults][%d] azure_key_vault_create default  %+v", o._statusCode, o.Payload)
}

func (o *AzureKeyVaultCreateDefault) String() string {
	return fmt.Sprintf("[POST /security/azure-key-vaults][%d] azure_key_vault_create default  %+v", o._statusCode, o.Payload)
}

func (o *AzureKeyVaultCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AzureKeyVaultCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
