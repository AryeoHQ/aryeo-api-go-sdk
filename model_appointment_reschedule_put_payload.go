/*
 * Aryeo
 *
 *
 * API version: 2021-06-17
 * Contact: jarrod@aryeo.com
 */

package aryeo

import (
	"encoding/json"
	"time"
)

// AppointmentReschedulePutPayload Payload for rescheduling an appointment record.
type AppointmentReschedulePutPayload struct {
	// The new date and time (ISO 8601 format) when the appointment is set to start.
	StartAt NullableTime `json:"start_at"`
	// The new date and time (ISO 8601 format) when the appointment is set to end.
	EndAt NullableTime `json:"end_at"`
	// Send a notification to the appointment's order's customer that the appointment was rescheduled.
	NotifyCustomer NullableBool `json:"notify_customer,omitempty"`
}

// NewAppointmentReschedulePutPayload instantiates a new AppointmentReschedulePutPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppointmentReschedulePutPayload(startAt NullableTime, endAt NullableTime) *AppointmentReschedulePutPayload {
	this := AppointmentReschedulePutPayload{}
	this.StartAt = startAt
	this.EndAt = endAt
	return &this
}

// NewAppointmentReschedulePutPayloadWithDefaults instantiates a new AppointmentReschedulePutPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppointmentReschedulePutPayloadWithDefaults() *AppointmentReschedulePutPayload {
	this := AppointmentReschedulePutPayload{}
	return &this
}

// GetStartAt returns the StartAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *AppointmentReschedulePutPayload) GetStartAt() time.Time {
	if o == nil || o.StartAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.StartAt.Get()
}

// GetStartAtOk returns a tuple with the StartAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppointmentReschedulePutPayload) GetStartAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartAt.Get(), o.StartAt.IsSet()
}

// SetStartAt sets field value
func (o *AppointmentReschedulePutPayload) SetStartAt(v time.Time) {
	o.StartAt.Set(&v)
}

// GetEndAt returns the EndAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *AppointmentReschedulePutPayload) GetEndAt() time.Time {
	if o == nil || o.EndAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.EndAt.Get()
}

// GetEndAtOk returns a tuple with the EndAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppointmentReschedulePutPayload) GetEndAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndAt.Get(), o.EndAt.IsSet()
}

// SetEndAt sets field value
func (o *AppointmentReschedulePutPayload) SetEndAt(v time.Time) {
	o.EndAt.Set(&v)
}

// GetNotifyCustomer returns the NotifyCustomer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppointmentReschedulePutPayload) GetNotifyCustomer() bool {
	if o == nil || o.NotifyCustomer.Get() == nil {
		var ret bool
		return ret
	}
	return *o.NotifyCustomer.Get()
}

// GetNotifyCustomerOk returns a tuple with the NotifyCustomer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppointmentReschedulePutPayload) GetNotifyCustomerOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NotifyCustomer.Get(), o.NotifyCustomer.IsSet()
}

// HasNotifyCustomer returns a boolean if a field has been set.
func (o *AppointmentReschedulePutPayload) HasNotifyCustomer() bool {
	if o != nil && o.NotifyCustomer.IsSet() {
		return true
	}

	return false
}

// SetNotifyCustomer gets a reference to the given NullableBool and assigns it to the NotifyCustomer field.
func (o *AppointmentReschedulePutPayload) SetNotifyCustomer(v bool) {
	o.NotifyCustomer.Set(&v)
}
// SetNotifyCustomerNil sets the value for NotifyCustomer to be an explicit nil
func (o *AppointmentReschedulePutPayload) SetNotifyCustomerNil() {
	o.NotifyCustomer.Set(nil)
}

// UnsetNotifyCustomer ensures that no value is present for NotifyCustomer, not even an explicit nil
func (o *AppointmentReschedulePutPayload) UnsetNotifyCustomer() {
	o.NotifyCustomer.Unset()
}

func (o AppointmentReschedulePutPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["start_at"] = o.StartAt.Get()
	}
	if true {
		toSerialize["end_at"] = o.EndAt.Get()
	}
	if o.NotifyCustomer.IsSet() {
		toSerialize["notify_customer"] = o.NotifyCustomer.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAppointmentReschedulePutPayload struct {
	value *AppointmentReschedulePutPayload
	isSet bool
}

func (v NullableAppointmentReschedulePutPayload) Get() *AppointmentReschedulePutPayload {
	return v.value
}

func (v *NullableAppointmentReschedulePutPayload) Set(val *AppointmentReschedulePutPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableAppointmentReschedulePutPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableAppointmentReschedulePutPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppointmentReschedulePutPayload(val *AppointmentReschedulePutPayload) *NullableAppointmentReschedulePutPayload {
	return &NullableAppointmentReschedulePutPayload{value: val, isSet: true}
}

func (v NullableAppointmentReschedulePutPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppointmentReschedulePutPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


