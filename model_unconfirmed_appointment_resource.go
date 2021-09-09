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

// UnconfirmedAppointmentResource An appointment.
type UnconfirmedAppointmentResource struct {
	// What was the state of the request?
	Status string `json:"status"`
	Data *UnconfirmedAppointment `json:"data,omitempty"`
}

// NewUnconfirmedAppointmentResource instantiates a new UnconfirmedAppointmentResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnconfirmedAppointmentResource(status string) *UnconfirmedAppointmentResource {
	this := UnconfirmedAppointmentResource{}
	this.Status = status
	return &this
}

// NewUnconfirmedAppointmentResourceWithDefaults instantiates a new UnconfirmedAppointmentResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnconfirmedAppointmentResourceWithDefaults() *UnconfirmedAppointmentResource {
	this := UnconfirmedAppointmentResource{}
	return &this
}

// GetStatus returns the Status field value
func (o *UnconfirmedAppointmentResource) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UnconfirmedAppointmentResource) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UnconfirmedAppointmentResource) SetStatus(v string) {
	o.Status = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UnconfirmedAppointmentResource) GetData() UnconfirmedAppointment {
	if o == nil || o.Data == nil {
		var ret UnconfirmedAppointment
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnconfirmedAppointmentResource) GetDataOk() (*UnconfirmedAppointment, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UnconfirmedAppointmentResource) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given UnconfirmedAppointment and assigns it to the Data field.
func (o *UnconfirmedAppointmentResource) SetData(v UnconfirmedAppointment) {
	o.Data = &v
}

func (o UnconfirmedAppointmentResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableUnconfirmedAppointmentResource struct {
	value *UnconfirmedAppointmentResource
	isSet bool
}

func (v NullableUnconfirmedAppointmentResource) Get() *UnconfirmedAppointmentResource {
	return v.value
}

func (v *NullableUnconfirmedAppointmentResource) Set(val *UnconfirmedAppointmentResource) {
	v.value = val
	v.isSet = true
}

func (v NullableUnconfirmedAppointmentResource) IsSet() bool {
	return v.isSet
}

func (v *NullableUnconfirmedAppointmentResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnconfirmedAppointmentResource(val *UnconfirmedAppointmentResource) *NullableUnconfirmedAppointmentResource {
	return &NullableUnconfirmedAppointmentResource{value: val, isSet: true}
}

func (v NullableUnconfirmedAppointmentResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnconfirmedAppointmentResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


