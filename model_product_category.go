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

// ProductCategory A category for products.
type ProductCategory struct {
	// ID of the product category. UUID Version 4.
	Id string `json:"id"`
	// The title of the product category.
	Title string `json:"title"`
}

// NewProductCategory instantiates a new ProductCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductCategory(id string, title string) *ProductCategory {
	this := ProductCategory{}
	this.Id = id
	this.Title = title
	return &this
}

// NewProductCategoryWithDefaults instantiates a new ProductCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductCategoryWithDefaults() *ProductCategory {
	this := ProductCategory{}
	return &this
}

// GetId returns the Id field value
func (o *ProductCategory) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductCategory) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProductCategory) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *ProductCategory) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ProductCategory) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ProductCategory) SetTitle(v string) {
	o.Title = v
}

func (o ProductCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableProductCategory struct {
	value *ProductCategory
	isSet bool
}

func (v NullableProductCategory) Get() *ProductCategory {
	return v.value
}

func (v *NullableProductCategory) Set(val *ProductCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableProductCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableProductCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductCategory(val *ProductCategory) *NullableProductCategory {
	return &NullableProductCategory{value: val, isSet: true}
}

func (v NullableProductCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


