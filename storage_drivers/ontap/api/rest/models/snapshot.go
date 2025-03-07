// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Snapshot The Snapshot copy object represents a point in time Snapshot copy of a volume.
//
// swagger:model snapshot
type Snapshot struct {

	// links
	Links *SnapshotLinks `json:"_links,omitempty"`

	// A comment associated with the Snapshot copy. This is an optional attribute for POST or PATCH.
	Comment string `json:"comment,omitempty"`

	// Creation time of the Snapshot copy. It is the volume access time when the Snapshot copy was created.
	// Example: 2019-02-04T19:00:00Z
	// Read Only: true
	// Format: date-time
	CreateTime *strfmt.DateTime `json:"create_time,omitempty"`

	// The expiry time for the Snapshot copy. This is an optional attribute for POST or PATCH. Snapshot copies with an expiry time set are not allowed to be deleted until the retention time is reached.
	// Example: 2019-02-04T19:00:00Z
	// Format: date-time
	ExpiryTime *strfmt.DateTime `json:"expiry_time,omitempty"`

	// Snapshot copy. Valid in POST or PATCH.
	// Example: this_snapshot
	Name string `json:"name,omitempty"`

	// owners
	// Read Only: true
	Owners []string `json:"owners,omitempty"`

	// provenance volume
	ProvenanceVolume *SnapshotProvenanceVolume `json:"provenance_volume,omitempty"`

	// Space reclaimed when the Snapshot copy is deleted, in bytes.
	ReclaimableSpace int64 `json:"reclaimable_space,omitempty"`

	// Size of the active file system at the time the Snapshot copy is captured. The actual size of the Snapshot copy also includes those blocks trapped by other Snapshot copies. On a Snapshot copy deletion, the "size" amount of blocks is the maximum number of blocks available. On a Snapshot copy restore, the "afs-used size" value will match the Snapshot copy "size" value.
	// Example: 122880
	// Read Only: true
	Size int64 `json:"size,omitempty"`

	// SnapLock expiry time for the Snapshot copy, if the Snapshot copy is taken on a SnapLock volume. A Snapshot copy is not allowed to be deleted or renamed until the SnapLock ComplianceClock time goes beyond this retention time.
	// Example: 2019-02-04T19:00:00Z
	// Read Only: true
	// Format: date-time
	SnaplockExpiryTime *strfmt.DateTime `json:"snaplock_expiry_time,omitempty"`

	// Label for SnapMirror operations
	SnapmirrorLabel string `json:"snapmirror_label,omitempty"`

	// State of the Snapshot copy. There are cases where some Snapshot copies are not complete. In the "partial" state, the Snapshot copy is consistent but exists only on the subset of the constituents that existed prior to the FlexGroup's expansion. Partial Snapshot copies cannot be used for a Snapshot copy restore operation. A Snapshot copy is in an "invalid" state when it is present in some FlexGroup constituents but not in others. At all other times, a Snapshot copy is valid.
	// Read Only: true
	// Enum: [valid invalid partial]
	State string `json:"state,omitempty"`

	// svm
	Svm *SnapshotSvm `json:"svm,omitempty"`

	// The UUID of the Snapshot copy in the volume that uniquely identifies the Snapshot copy in that volume.
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID string `json:"uuid,omitempty"`

	// The 128 bit identifier that uniquely identifies a snapshot and its logical data layout.
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	VersionUUID string `json:"version_uuid,omitempty"`

	// volume
	Volume *SnapshotVolume `json:"volume,omitempty"`
}

// Validate validates this snapshot
func (m *Snapshot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiryTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwners(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvenanceVolume(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnaplockExpiryTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Snapshot) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *Snapshot) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) validateExpiryTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiryTime) { // not required
		return nil
	}

	if err := validate.FormatOf("expiry_time", "body", "date-time", m.ExpiryTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var snapshotOwnersItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unknown","snapmirror","snapmirror_dependent","sync_mirror","volume_clone","volume_clone_dependent","snap_restore","snap_restore_dependent","dump","cifs_share","volume_copy","ndmp","worm_volume","sis_clone","s2c_iron","lun_clone","backup_dependent","snaplock_dependent","file_clone_dependent","volume_move_dependent","svmdr_dependent","anti_ransomware_backup"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snapshotOwnersItemsEnum = append(snapshotOwnersItemsEnum, v)
	}
}

func (m *Snapshot) validateOwnersItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snapshotOwnersItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Snapshot) validateOwners(formats strfmt.Registry) error {
	if swag.IsZero(m.Owners) { // not required
		return nil
	}

	for i := 0; i < len(m.Owners); i++ {

		// value enum
		if err := m.validateOwnersItemsEnum("owners"+"."+strconv.Itoa(i), "body", m.Owners[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *Snapshot) validateProvenanceVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.ProvenanceVolume) { // not required
		return nil
	}

	if m.ProvenanceVolume != nil {
		if err := m.ProvenanceVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provenance_volume")
			}
			return err
		}
	}

	return nil
}

func (m *Snapshot) validateSnaplockExpiryTime(formats strfmt.Registry) error {
	if swag.IsZero(m.SnaplockExpiryTime) { // not required
		return nil
	}

	if err := validate.FormatOf("snaplock_expiry_time", "body", "date-time", m.SnaplockExpiryTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var snapshotTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["valid","invalid","partial"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snapshotTypeStatePropEnum = append(snapshotTypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// snapshot
	// Snapshot
	// state
	// State
	// valid
	// END DEBUGGING
	// SnapshotStateValid captures enum value "valid"
	SnapshotStateValid string = "valid"

	// BEGIN DEBUGGING
	// snapshot
	// Snapshot
	// state
	// State
	// invalid
	// END DEBUGGING
	// SnapshotStateInvalid captures enum value "invalid"
	SnapshotStateInvalid string = "invalid"

	// BEGIN DEBUGGING
	// snapshot
	// Snapshot
	// state
	// State
	// partial
	// END DEBUGGING
	// SnapshotStatePartial captures enum value "partial"
	SnapshotStatePartial string = "partial"
)

// prop value enum
func (m *Snapshot) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snapshotTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Snapshot) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

func (m *Snapshot) validateVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.Volume) { // not required
		return nil
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot based on the context it is used
func (m *Snapshot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwners(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProvenanceVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnaplockExpiryTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersionUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Snapshot) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *Snapshot) contextValidateCreateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "create_time", "body", m.CreateTime); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) contextValidateOwners(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "owners", "body", []string(m.Owners)); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) contextValidateProvenanceVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.ProvenanceVolume != nil {
		if err := m.ProvenanceVolume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("provenance_volume")
			}
			return err
		}
	}

	return nil
}

func (m *Snapshot) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "size", "body", int64(m.Size)); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) contextValidateSnaplockExpiryTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "snaplock_expiry_time", "body", m.SnaplockExpiryTime); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

func (m *Snapshot) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) contextValidateVersionUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "version_uuid", "body", string(m.VersionUUID)); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) contextValidateVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.Volume != nil {
		if err := m.Volume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Snapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Snapshot) UnmarshalBinary(b []byte) error {
	var res Snapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapshotLinks snapshot links
//
// swagger:model SnapshotLinks
type SnapshotLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snapshot links
func (m *SnapshotLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot links based on the context it is used
func (m *SnapshotLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotLinks) UnmarshalBinary(b []byte) error {
	var res SnapshotLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapshotProvenanceVolume snapshot provenance volume
//
// swagger:model SnapshotProvenanceVolume
type SnapshotProvenanceVolume struct {

	// UUID for the volume that is used to identify the source volume in a mirroring relationship. When the mirroring relationship is broken, a volume's Instance UUID and Provenance UUID are made identical. An unmirrored volume's Provenance UUID is the same as its Instance UUID. This field is valid for flexible volumes only.
	// Example: 4cd8a442-86d1-11e0-ae1c-125648563413
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this snapshot provenance volume
func (m *SnapshotProvenanceVolume) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this snapshot provenance volume based on the context it is used
func (m *SnapshotProvenanceVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotProvenanceVolume) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "provenance_volume"+"."+"uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotProvenanceVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotProvenanceVolume) UnmarshalBinary(b []byte) error {
	var res SnapshotProvenanceVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapshotSvm snapshot svm
//
// swagger:model SnapshotSvm
type SnapshotSvm struct {

	// links
	Links *SnapshotSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this snapshot svm
func (m *SnapshotSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot svm based on the context it is used
func (m *SnapshotSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotSvm) UnmarshalBinary(b []byte) error {
	var res SnapshotSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapshotSvmLinks snapshot svm links
//
// swagger:model SnapshotSvmLinks
type SnapshotSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snapshot svm links
func (m *SnapshotSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotSvmLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot svm links based on the context it is used
func (m *SnapshotSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotSvmLinks) UnmarshalBinary(b []byte) error {
	var res SnapshotSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapshotVolume snapshot volume
//
// swagger:model SnapshotVolume
type SnapshotVolume struct {

	// links
	Links *SnapshotVolumeLinks `json:"_links,omitempty"`

	// The name of the volume.
	// Example: volume1
	Name string `json:"name,omitempty"`

	// Unique identifier for the volume. This corresponds to the instance-uuid that is exposed in the CLI and ONTAPI. It does not change due to a volume move.
	// Example: 028baa66-41bd-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this snapshot volume
func (m *SnapshotVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotVolume) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot volume based on the context it is used
func (m *SnapshotVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotVolume) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotVolume) UnmarshalBinary(b []byte) error {
	var res SnapshotVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapshotVolumeLinks snapshot volume links
//
// swagger:model SnapshotVolumeLinks
type SnapshotVolumeLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snapshot volume links
func (m *SnapshotVolumeLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotVolumeLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot volume links based on the context it is used
func (m *SnapshotVolumeLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotVolumeLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotVolumeLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotVolumeLinks) UnmarshalBinary(b []byte) error {
	var res SnapshotVolumeLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
