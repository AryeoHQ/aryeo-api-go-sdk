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

// GroupResource A group.
type GroupResource struct {
	// What was the state of the request?
	Status string `json:"status"`
	Data *Group `json:"data,omitempty"`
}

// NewGroupResource instantiates a new GroupResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupResource(status string) *GroupResource {
	this := GroupResource{}
	this.Status = status
	return &this
}

// NewGroupResourceWithDefaults instantiates a new GroupResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupResourceWithDefaults() *GroupResource {
	this := GroupResource{}
	return &this
}

// GetStatus returns the Status field value
func (o *GroupResource) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GroupResource) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GroupResource) SetStatus(v string) {
	o.Status = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GroupResource) GetData() Group {
	if o == nil || o.Data == nil {
		var ret Group
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupResource) GetDataOk() (*Group, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GroupResource) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given Group and assigns it to the Data field.
func (o *GroupResource) SetData(v Group) {
	o.Data = &v
}

func (o GroupResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGroupResource struct {
	value *GroupResource
	isSet bool
}

func (v NullableGroupResource) Get() *GroupResource {
	return v.value
}

func (v *NullableGroupResource) Set(val *GroupResource) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupResource) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupResource(val *GroupResource) *NullableGroupResource {
	return &NullableGroupResource{value: val, isSet: true}
}

func (v NullableGroupResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


