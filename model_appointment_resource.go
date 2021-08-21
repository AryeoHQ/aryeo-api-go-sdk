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

// AppointmentResource An appointment.
type AppointmentResource struct {
	// What was the state of the request?
	Status string `json:"status"`
	Data *Appointment `json:"data,omitempty"`
}

// NewAppointmentResource instantiates a new AppointmentResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppointmentResource(status string) *AppointmentResource {
	this := AppointmentResource{}
	this.Status = status
	return &this
}

// NewAppointmentResourceWithDefaults instantiates a new AppointmentResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppointmentResourceWithDefaults() *AppointmentResource {
	this := AppointmentResource{}
	return &this
}

// GetStatus returns the Status field value
func (o *AppointmentResource) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AppointmentResource) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AppointmentResource) SetStatus(v string) {
	o.Status = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppointmentResource) GetData() Appointment {
	if o == nil || o.Data == nil {
		var ret Appointment
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentResource) GetDataOk() (*Appointment, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppointmentResource) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given Appointment and assigns it to the Data field.
func (o *AppointmentResource) SetData(v Appointment) {
	o.Data = &v
}

func (o AppointmentResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAppointmentResource struct {
	value *AppointmentResource
	isSet bool
}

func (v NullableAppointmentResource) Get() *AppointmentResource {
	return v.value
}

func (v *NullableAppointmentResource) Set(val *AppointmentResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAppointmentResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAppointmentResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppointmentResource(val *AppointmentResource) *NullableAppointmentResource {
	return &NullableAppointmentResource{value: val, isSet: true}
}

func (v NullableAppointmentResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppointmentResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


