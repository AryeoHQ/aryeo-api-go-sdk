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

// ProductItem A subtype or part of a product that a group is selling.
type ProductItem struct {
	// ID of the product item.
	Id NullableString `json:"id,omitempty"`
}

// NewProductItem instantiates a new ProductItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductItem() *ProductItem {
	this := ProductItem{}
	return &this
}

// NewProductItemWithDefaults instantiates a new ProductItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductItemWithDefaults() *ProductItem {
	this := ProductItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductItem) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductItem) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ProductItem) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ProductItem) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ProductItem) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ProductItem) UnsetId() {
	o.Id.Unset()
}

func (o ProductItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableProductItem struct {
	value *ProductItem
	isSet bool
}

func (v NullableProductItem) Get() *ProductItem {
	return v.value
}

func (v *NullableProductItem) Set(val *ProductItem) {
	v.value = val
	v.isSet = true
}

func (v NullableProductItem) IsSet() bool {
	return v.isSet
}

func (v *NullableProductItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductItem(val *ProductItem) *NullableProductItem {
	return &NullableProductItem{value: val, isSet: true}
}

func (v NullableProductItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


