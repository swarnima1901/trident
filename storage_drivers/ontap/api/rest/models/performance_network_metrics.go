// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PerformanceNetworkMetrics performance network metrics
//
// swagger:model performance_network_metrics
type PerformanceNetworkMetrics struct {

	// links
	Links *PerformanceNetworkMetricsLinks `json:"_links,omitempty"`

	// The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:
	//
	// Example: PT15S
	// Enum: [PT15S PT4M PT30M PT2H P1D PT5M]
	Duration string `json:"duration,omitempty"`

	// Errors associated with the sample. For example, if the aggregation of data over multiple nodes fails, then any partial errors might return "ok" on success or "error" on an internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "inconsistent_delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Enum: [ok error partial_no_data partial_no_uuid partial_no_response partial_other_error negative_delta backfilled_data inconsistent_delta_time inconsistent_old_data]
	Status string `json:"status,omitempty"`

	// throughput
	Throughput *PerformanceNetworkMetricsThroughput `json:"throughput,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25T11:20:13Z
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this performance network metrics
func (m *PerformanceNetworkMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThroughput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceNetworkMetrics) validateLinks(formats strfmt.Registry) error {
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

var performanceNetworkMetricsTypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceNetworkMetricsTypeDurationPropEnum = append(performanceNetworkMetricsTypeDurationPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// performance_network_metrics
	// PerformanceNetworkMetrics
	// duration
	// Duration
	// PT15S
	// END DEBUGGING
	// PerformanceNetworkMetricsDurationPT15S captures enum value "PT15S"
	PerformanceNetworkMetricsDurationPT15S string = "PT15S"

	// BEGIN DEBUGGING
	// performance_network_metrics
	// PerformanceNetworkMetrics
	// duration
	// Duration
	// PT4M
	// END DEBUGGING
	// PerformanceNetworkMetricsDurationPT4M captures enum value "PT4M"
	PerformanceNetworkMetricsDurationPT4M string = "PT4M"

	// BEGIN DEBUGGING
	// performance_network_metrics
	// PerformanceNetworkMetrics
	// duration
	// Duration
	// PT30M
	// END DEBUGGING
	// PerformanceNetworkMetricsDurationPT30M captures enum value "PT30M"
	PerformanceNetworkMetricsDurationPT30M string = "PT30M"

	// BEGIN DEBUGGING
	// performance_network_metrics
	// PerformanceNetworkMetrics
	// duration
	// Duration
	// PT2H
	// END DEBUGGING
	// PerformanceNetworkMetricsDurationPT2H captures enum value "PT2H"
	PerformanceNetworkMetricsDurationPT2H string = "PT2H"

	// BEGIN DEBUGGING
	// performance_network_metrics
	// PerformanceNetworkMetrics
	// duration
	// Duration
	// P1D
	// END DEBUGGING
	// PerformanceNetworkMetricsDurationP1D captures enum value "P1D"
	PerformanceNetworkMetricsDurationP1D string = "P1D"

	// BEGIN DEBUGGING
	// performance_network_metrics
	// PerformanceNetworkMetrics
	// duration
	// Duration
	// PT5M
	// END DEBUGGING
	// PerformanceNetworkMetricsDurationPT5M captures enum value "PT5M"
	PerformanceNetworkMetricsDurationPT5M string = "PT5M"
)

// prop value enum
func (m *PerformanceNetworkMetrics) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceNetworkMetricsTypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceNetworkMetrics) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

var performanceNetworkMetricsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_uuid","partial_no_response","partial_other_error","negative_delta","backfilled_data","inconsistent_delta_time","inconsistent_old_data"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceNetworkMetricsTypeStatusPropEnum = append(performanceNetworkMetricsTypeStatusPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// performance_network_metrics
	// PerformanceNetworkMetrics
	// status
	// Status
	// ok
	// END DEBUGGING
	// PerformanceNetworkMetricsStatusOk captures enum value "ok"
	PerformanceNetworkMetricsStatusOk string = "ok"

	// BEGIN DEBUGGING
	// performance_network_metrics
	// PerformanceNetworkMetrics
	// status
	// Status
	// error
	// END DEBUGGING
	// PerformanceNetworkMetricsStatusError captures enum value "error"
	PerformanceNetworkMetricsStatusError string = "error"

	// BEGIN DEBUGGING
	// performance_network_metrics
	// PerformanceNetworkMetrics
	// status
	// Status
	// partial_no_data
	// END DEBUGGING
	// PerformanceNetworkMetricsStatusPartialNoData captures enum value "partial_no_data"
	PerformanceNetworkMetricsStatusPartialNoData string = "partial_no_data"

	// BEGIN DEBUGGING
	// performance_network_metrics
	// PerformanceNetworkMetrics
	// status
	// Status
	// partial_no_uuid
	// END DEBUGGING
	// PerformanceNetworkMetricsStatusPartialNoUUID captures enum value "partial_no_uuid"
	PerformanceNetworkMetricsStatusPartialNoUUID string = "partial_no_uuid"

	// BEGIN DEBUGGING
	// performance_network_metrics
	// PerformanceNetworkMetrics
	// status
	// Status
	// partial_no_response
	// END DEBUGGING
	// PerformanceNetworkMetricsStatusPartialNoResponse captures enum value "partial_no_response"
	PerformanceNetworkMetricsStatusPartialNoResponse string = "partial_no_response"

	// BEGIN DEBUGGING
	// performance_network_metrics
	// PerformanceNetworkMetrics
	// status
	// Status
	// partial_other_error
	// END DEBUGGING
	// PerformanceNetworkMetricsStatusPartialOtherError captures enum value "partial_other_error"
	PerformanceNetworkMetricsStatusPartialOtherError string = "partial_other_error"

	// BEGIN DEBUGGING
	// performance_network_metrics
	// PerformanceNetworkMetrics
	// status
	// Status
	// negative_delta
	// END DEBUGGING
	// PerformanceNetworkMetricsStatusNegativeDelta captures enum value "negative_delta"
	PerformanceNetworkMetricsStatusNegativeDelta string = "negative_delta"

	// BEGIN DEBUGGING
	// performance_network_metrics
	// PerformanceNetworkMetrics
	// status
	// Status
	// backfilled_data
	// END DEBUGGING
	// PerformanceNetworkMetricsStatusBackfilledData captures enum value "backfilled_data"
	PerformanceNetworkMetricsStatusBackfilledData string = "backfilled_data"

	// BEGIN DEBUGGING
	// performance_network_metrics
	// PerformanceNetworkMetrics
	// status
	// Status
	// inconsistent_delta_time
	// END DEBUGGING
	// PerformanceNetworkMetricsStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PerformanceNetworkMetricsStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// BEGIN DEBUGGING
	// performance_network_metrics
	// PerformanceNetworkMetrics
	// status
	// Status
	// inconsistent_old_data
	// END DEBUGGING
	// PerformanceNetworkMetricsStatusInconsistentOldData captures enum value "inconsistent_old_data"
	PerformanceNetworkMetricsStatusInconsistentOldData string = "inconsistent_old_data"
)

// prop value enum
func (m *PerformanceNetworkMetrics) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceNetworkMetricsTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceNetworkMetrics) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceNetworkMetrics) validateThroughput(formats strfmt.Registry) error {
	if swag.IsZero(m.Throughput) { // not required
		return nil
	}

	if m.Throughput != nil {
		if err := m.Throughput.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceNetworkMetrics) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this performance network metrics based on the context it is used
func (m *PerformanceNetworkMetrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceNetworkMetrics) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceNetworkMetrics) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

	if m.Throughput != nil {
		if err := m.Throughput.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceNetworkMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceNetworkMetrics) UnmarshalBinary(b []byte) error {
	var res PerformanceNetworkMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceNetworkMetricsLinks performance network metrics links
//
// swagger:model PerformanceNetworkMetricsLinks
type PerformanceNetworkMetricsLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this performance network metrics links
func (m *PerformanceNetworkMetricsLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceNetworkMetricsLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this performance network metrics links based on the context it is used
func (m *PerformanceNetworkMetricsLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceNetworkMetricsLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PerformanceNetworkMetricsLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceNetworkMetricsLinks) UnmarshalBinary(b []byte) error {
	var res PerformanceNetworkMetricsLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceNetworkMetricsThroughput The rate of throughput bytes per second observed at the interface.
//
// swagger:model PerformanceNetworkMetricsThroughput
type PerformanceNetworkMetricsThroughput struct {

	// Performance metric for read I/O operations.
	// Example: 200
	Read int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total int64 `json:"total,omitempty"`

	// Peformance metric for write I/O operations.
	// Example: 100
	Write int64 `json:"write,omitempty"`
}

// Validate validates this performance network metrics throughput
func (m *PerformanceNetworkMetricsThroughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this performance network metrics throughput based on context it is used
func (m *PerformanceNetworkMetricsThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceNetworkMetricsThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceNetworkMetricsThroughput) UnmarshalBinary(b []byte) error {
	var res PerformanceNetworkMetricsThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
