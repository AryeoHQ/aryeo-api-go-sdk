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

// ProductVariant A variant of a product.
type ProductVariant struct {
	// ID of the product variant. UUID Version 4.
	Id string `json:"id"`
	// The title of the product variant.
	Title string `json:"title"`
	// A positive integer in the smallest currency unit (that is, 100 cents for $1.00) representing the price of the product variant.
	Price int32 `json:"price"`
}

// NewProductVariant instantiates a new ProductVariant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductVariant(id string, title string, price int32) *ProductVariant {
	this := ProductVariant{}
	this.Id = id
	this.Title = title
	this.Price = price
	return &this
}

// NewProductVariantWithDefaults instantiates a new ProductVariant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductVariantWithDefaults() *ProductVariant {
	this := ProductVariant{}
	return &this
}

// GetId returns the Id field value
func (o *ProductVariant) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProductVariant) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProductVariant) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *ProductVariant) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ProductVariant) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ProductVariant) SetTitle(v string) {
	o.Title = v
}

// GetPrice returns the Price field value
func (o *ProductVariant) GetPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *ProductVariant) GetPriceOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *ProductVariant) SetPrice(v int32) {
	o.Price = v
}

func (o ProductVariant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["price"] = o.Price
	}
	return json.Marshal(toSerialize)
}

type NullableProductVariant struct {
	value *ProductVariant
	isSet bool
}

func (v NullableProductVariant) Get() *ProductVariant {
	return v.value
}

func (v *NullableProductVariant) Set(val *ProductVariant) {
	v.value = val
	v.isSet = true
}

func (v NullableProductVariant) IsSet() bool {
	return v.isSet
}

func (v *NullableProductVariant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductVariant(val *ProductVariant) *NullableProductVariant {
	return &NullableProductVariant{value: val, isSet: true}
}

func (v NullableProductVariant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductVariant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


