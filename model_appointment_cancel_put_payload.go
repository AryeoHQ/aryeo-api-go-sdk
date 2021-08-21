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
)

// AppointmentCancelPutPayload Payload for canceling an appointment record.
type AppointmentCancelPutPayload struct {
	// Sends a notification to the appointment's order's customer that the appointment was canceled.
	NotifyCustomer NullableBool `json:"notify_customer,omitempty"`
}

// NewAppointmentCancelPutPayload instantiates a new AppointmentCancelPutPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppointmentCancelPutPayload() *AppointmentCancelPutPayload {
	this := AppointmentCancelPutPayload{}
	return &this
}

// NewAppointmentCancelPutPayloadWithDefaults instantiates a new AppointmentCancelPutPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppointmentCancelPutPayloadWithDefaults() *AppointmentCancelPutPayload {
	this := AppointmentCancelPutPayload{}
	return &this
}

// GetNotifyCustomer returns the NotifyCustomer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppointmentCancelPutPayload) GetNotifyCustomer() bool {
	if o == nil || o.NotifyCustomer.Get() == nil {
		var ret bool
		return ret
	}
	return *o.NotifyCustomer.Get()
}

// GetNotifyCustomerOk returns a tuple with the NotifyCustomer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppointmentCancelPutPayload) GetNotifyCustomerOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NotifyCustomer.Get(), o.NotifyCustomer.IsSet()
}

// HasNotifyCustomer returns a boolean if a field has been set.
func (o *AppointmentCancelPutPayload) HasNotifyCustomer() bool {
	if o != nil && o.NotifyCustomer.IsSet() {
		return true
	}

	return false
}

// SetNotifyCustomer gets a reference to the given NullableBool and assigns it to the NotifyCustomer field.
func (o *AppointmentCancelPutPayload) SetNotifyCustomer(v bool) {
	o.NotifyCustomer.Set(&v)
}
// SetNotifyCustomerNil sets the value for NotifyCustomer to be an explicit nil
func (o *AppointmentCancelPutPayload) SetNotifyCustomerNil() {
	o.NotifyCustomer.Set(nil)
}

// UnsetNotifyCustomer ensures that no value is present for NotifyCustomer, not even an explicit nil
func (o *AppointmentCancelPutPayload) UnsetNotifyCustomer() {
	o.NotifyCustomer.Unset()
}

func (o AppointmentCancelPutPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NotifyCustomer.IsSet() {
		toSerialize["notify_customer"] = o.NotifyCustomer.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAppointmentCancelPutPayload struct {
	value *AppointmentCancelPutPayload
	isSet bool
}

func (v NullableAppointmentCancelPutPayload) Get() *AppointmentCancelPutPayload {
	return v.value
}

func (v *NullableAppointmentCancelPutPayload) Set(val *AppointmentCancelPutPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableAppointmentCancelPutPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableAppointmentCancelPutPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppointmentCancelPutPayload(val *AppointmentCancelPutPayload) *NullableAppointmentCancelPutPayload {
	return &NullableAppointmentCancelPutPayload{value: val, isSet: true}
}

func (v NullableAppointmentCancelPutPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppointmentCancelPutPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


