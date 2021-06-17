/*
 * Aryeo
 *
 *
 * API version: 1.0.0
 * Contact: jarrod@aryeo.com
 */

package aryeo

import (
	"encoding/json"
)

// GroupCollection A collection of groups.
type GroupCollection struct {
	Data *[]Group `json:"data,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
	Links *PaginationLinks `json:"links,omitempty"`
}

// NewGroupCollection instantiates a new GroupCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupCollection() *GroupCollection {
	this := GroupCollection{}
	return &this
}

// NewGroupCollectionWithDefaults instantiates a new GroupCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupCollectionWithDefaults() *GroupCollection {
	this := GroupCollection{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GroupCollection) GetData() []Group {
	if o == nil || o.Data == nil {
		var ret []Group
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupCollection) GetDataOk() (*[]Group, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GroupCollection) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Group and assigns it to the Data field.
func (o *GroupCollection) SetData(v []Group) {
	o.Data = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GroupCollection) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupCollection) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GroupCollection) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *GroupCollection) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GroupCollection) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupCollection) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GroupCollection) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *GroupCollection) SetLinks(v PaginationLinks) {
	o.Links = &v
}

func (o GroupCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableGroupCollection struct {
	value *GroupCollection
	isSet bool
}

func (v NullableGroupCollection) Get() *GroupCollection {
	return v.value
}

func (v *NullableGroupCollection) Set(val *GroupCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupCollection(val *GroupCollection) *NullableGroupCollection {
	return &NullableGroupCollection{value: val, isSet: true}
}

func (v NullableGroupCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


