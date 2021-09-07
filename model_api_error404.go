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

// ApiError404 A not found error returned by the API.
type ApiError404 struct {
	// What was the state of the request?
	Status string `json:"status"`
	// The error message.
	Message string `json:"message"`
	// A numeric code corresponding to the error.
	Code NullableInt32 `json:"code,omitempty"`
}

// NewApiError404 instantiates a new ApiError404 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiError404(status string, message string) *ApiError404 {
	this := ApiError404{}
	this.Status = status
	this.Message = message
	return &this
}

// NewApiError404WithDefaults instantiates a new ApiError404 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiError404WithDefaults() *ApiError404 {
	this := ApiError404{}
	return &this
}

// GetStatus returns the Status field value
func (o *ApiError404) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApiError404) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApiError404) SetStatus(v string) {
	o.Status = v
}

// GetMessage returns the Message field value
func (o *ApiError404) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ApiError404) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ApiError404) SetMessage(v string) {
	o.Message = v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiError404) GetCode() int32 {
	if o == nil || o.Code.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiError404) GetCodeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *ApiError404) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableInt32 and assigns it to the Code field.
func (o *ApiError404) SetCode(v int32) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *ApiError404) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *ApiError404) UnsetCode() {
	o.Code.Unset()
}

func (o ApiError404) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableApiError404 struct {
	value *ApiError404
	isSet bool
}

func (v NullableApiError404) Get() *ApiError404 {
	return v.value
}

func (v *NullableApiError404) Set(val *ApiError404) {
	v.value = val
	v.isSet = true
}

func (v NullableApiError404) IsSet() bool {
	return v.isSet
}

func (v *NullableApiError404) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiError404(val *ApiError404) *NullableApiError404 {
	return &NullableApiError404{value: val, isSet: true}
}

func (v NullableApiError404) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiError404) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


