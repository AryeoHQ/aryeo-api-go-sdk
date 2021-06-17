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

// PartialListingCollection A collection of partial listings.
type PartialListingCollection struct {
	Data *[]PartialListing `json:"data,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
	Links *PaginationLinks `json:"links,omitempty"`
}

// NewPartialListingCollection instantiates a new PartialListingCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialListingCollection() *PartialListingCollection {
	this := PartialListingCollection{}
	return &this
}

// NewPartialListingCollectionWithDefaults instantiates a new PartialListingCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialListingCollectionWithDefaults() *PartialListingCollection {
	this := PartialListingCollection{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PartialListingCollection) GetData() []PartialListing {
	if o == nil || o.Data == nil {
		var ret []PartialListing
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialListingCollection) GetDataOk() (*[]PartialListing, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PartialListingCollection) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []PartialListing and assigns it to the Data field.
func (o *PartialListingCollection) SetData(v []PartialListing) {
	o.Data = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PartialListingCollection) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialListingCollection) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PartialListingCollection) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *PartialListingCollection) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PartialListingCollection) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialListingCollection) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PartialListingCollection) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *PartialListingCollection) SetLinks(v PaginationLinks) {
	o.Links = &v
}

func (o PartialListingCollection) MarshalJSON() ([]byte, error) {
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

type NullablePartialListingCollection struct {
	value *PartialListingCollection
	isSet bool
}

func (v NullablePartialListingCollection) Get() *PartialListingCollection {
	return v.value
}

func (v *NullablePartialListingCollection) Set(val *PartialListingCollection) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialListingCollection) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialListingCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialListingCollection(val *PartialListingCollection) *NullablePartialListingCollection {
	return &NullablePartialListingCollection{value: val, isSet: true}
}

func (v NullablePartialListingCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialListingCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


