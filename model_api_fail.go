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

// ApiFail A generic failure returned by the API.
type ApiFail struct {
	// What was the state of the request?
	Status string `json:"status"`
}

// NewApiFail instantiates a new ApiFail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiFail(status string) *ApiFail {
	this := ApiFail{}
	this.Status = status
	return &this
}

// NewApiFailWithDefaults instantiates a new ApiFail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiFailWithDefaults() *ApiFail {
	this := ApiFail{}
	return &this
}

// GetStatus returns the Status field value
func (o *ApiFail) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApiFail) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApiFail) SetStatus(v string) {
	o.Status = v
}

func (o ApiFail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableApiFail struct {
	value *ApiFail
	isSet bool
}

func (v NullableApiFail) Get() *ApiFail {
	return v.value
}

func (v *NullableApiFail) Set(val *ApiFail) {
	v.value = val
	v.isSet = true
}

func (v NullableApiFail) IsSet() bool {
	return v.isSet
}

func (v *NullableApiFail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiFail(val *ApiFail) *NullableApiFail {
	return &NullableApiFail{value: val, isSet: true}
}

func (v NullableApiFail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiFail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


