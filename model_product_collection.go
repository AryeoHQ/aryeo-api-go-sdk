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

// ProductCollection A collection of products.
type ProductCollection struct {
	// What was the state of the request?
	Status string `json:"status"`
	// 
	Data []Product `json:"data,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
	Links *PaginationLinks `json:"links,omitempty"`
}

// NewProductCollection instantiates a new ProductCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductCollection(status string) *ProductCollection {
	this := ProductCollection{}
	this.Status = status
	return &this
}

// NewProductCollectionWithDefaults instantiates a new ProductCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductCollectionWithDefaults() *ProductCollection {
	this := ProductCollection{}
	return &this
}

// GetStatus returns the Status field value
func (o *ProductCollection) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ProductCollection) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ProductCollection) SetStatus(v string) {
	o.Status = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductCollection) GetData() []Product {
	if o == nil  {
		var ret []Product
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductCollection) GetDataOk() (*[]Product, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ProductCollection) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Product and assigns it to the Data field.
func (o *ProductCollection) SetData(v []Product) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ProductCollection) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCollection) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ProductCollection) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *ProductCollection) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProductCollection) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductCollection) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProductCollection) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *ProductCollection) SetLinks(v PaginationLinks) {
	o.Links = &v
}

func (o ProductCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
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

type NullableProductCollection struct {
	value *ProductCollection
	isSet bool
}

func (v NullableProductCollection) Get() *ProductCollection {
	return v.value
}

func (v *NullableProductCollection) Set(val *ProductCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableProductCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableProductCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductCollection(val *ProductCollection) *NullableProductCollection {
	return &NullableProductCollection{value: val, isSet: true}
}

func (v NullableProductCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


