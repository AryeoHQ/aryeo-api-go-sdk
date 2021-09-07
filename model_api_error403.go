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

// ApiError403 A forbidden error returned by the API.
type ApiError403 struct {
	// What was the state of the request?
	Status string `json:"status"`
	// The error message.
	Message string `json:"message"`
	// A numeric code corresponding to the error.
	Code NullableInt32 `json:"code,omitempty"`
}

// NewApiError403 instantiates a new ApiError403 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiError403(status string, message string) *ApiError403 {
	this := ApiError403{}
	this.Status = status
	this.Message = message
	return &this
}

// NewApiError403WithDefaults instantiates a new ApiError403 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiError403WithDefaults() *ApiError403 {
	this := ApiError403{}
	return &this
}

// GetStatus returns the Status field value
func (o *ApiError403) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApiError403) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApiError403) SetStatus(v string) {
	o.Status = v
}

// GetMessage returns the Message field value
func (o *ApiError403) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ApiError403) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ApiError403) SetMessage(v string) {
	o.Message = v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiError403) GetCode() int32 {
	if o == nil || o.Code.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiError403) GetCodeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *ApiError403) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableInt32 and assigns it to the Code field.
func (o *ApiError403) SetCode(v int32) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *ApiError403) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *ApiError403) UnsetCode() {
	o.Code.Unset()
}

func (o ApiError403) MarshalJSON() ([]byte, error) {
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

type NullableApiError403 struct {
	value *ApiError403
	isSet bool
}

func (v NullableApiError403) Get() *ApiError403 {
	return v.value
}

func (v *NullableApiError403) Set(val *ApiError403) {
	v.value = val
	v.isSet = true
}

func (v NullableApiError403) IsSet() bool {
	return v.isSet
}

func (v *NullableApiError403) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiError403(val *ApiError403) *NullableApiError403 {
	return &NullableApiError403{value: val, isSet: true}
}

func (v NullableApiError403) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiError403) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


