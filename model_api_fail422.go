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

// ApiFail422 A processing or validation failure returned by the API.
type ApiFail422 struct {
	// What was the state of the request?
	Status string `json:"status"`
}

// NewApiFail422 instantiates a new ApiFail422 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiFail422(status string) *ApiFail422 {
	this := ApiFail422{}
	this.Status = status
	return &this
}

// NewApiFail422WithDefaults instantiates a new ApiFail422 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiFail422WithDefaults() *ApiFail422 {
	this := ApiFail422{}
	return &this
}

// GetStatus returns the Status field value
func (o *ApiFail422) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApiFail422) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApiFail422) SetStatus(v string) {
	o.Status = v
}

func (o ApiFail422) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableApiFail422 struct {
	value *ApiFail422
	isSet bool
}

func (v NullableApiFail422) Get() *ApiFail422 {
	return v.value
}

func (v *NullableApiFail422) Set(val *ApiFail422) {
	v.value = val
	v.isSet = true
}

func (v NullableApiFail422) IsSet() bool {
	return v.isSet
}

func (v *NullableApiFail422) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiFail422(val *ApiFail422) *NullableApiFail422 {
	return &NullableApiFail422{value: val, isSet: true}
}

func (v NullableApiFail422) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiFail422) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


