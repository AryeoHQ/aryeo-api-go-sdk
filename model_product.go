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

// Product A product available for purchase via an order.
type Product struct {
	// ID of the product. UUID Version 4.
	Id string `json:"id"`
	// The title of the product.
	Title string `json:"title"`
	// The description of the product.
	Description *string `json:"description,omitempty"`
	// The active status of a product.
	Active *bool `json:"active,omitempty"`
	// The type of product.
	Type string `json:"type"`
	Variants *[]ProductVariant `json:"variants,omitempty"`
	Categories *[]ProductCategory `json:"categories,omitempty"`
}

// NewProduct instantiates a new Product object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct(id string, title string, type_ string) *Product {
	this := Product{}
	this.Id = id
	this.Title = title
	this.Type = type_
	return &this
}

// NewProductWithDefaults instantiates a new Product object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductWithDefaults() *Product {
	this := Product{}
	return &this
}

// GetId returns the Id field value
func (o *Product) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Product) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Product) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *Product) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Product) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Product) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Product) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Product) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Product) SetDescription(v string) {
	o.Description = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Product) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Product) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Product) SetActive(v bool) {
	o.Active = &v
}

// GetType returns the Type field value
func (o *Product) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Product) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Product) SetType(v string) {
	o.Type = v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *Product) GetVariants() []ProductVariant {
	if o == nil || o.Variants == nil {
		var ret []ProductVariant
		return ret
	}
	return *o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetVariantsOk() (*[]ProductVariant, bool) {
	if o == nil || o.Variants == nil {
		return nil, false
	}
	return o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *Product) HasVariants() bool {
	if o != nil && o.Variants != nil {
		return true
	}

	return false
}

// SetVariants gets a reference to the given []ProductVariant and assigns it to the Variants field.
func (o *Product) SetVariants(v []ProductVariant) {
	o.Variants = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *Product) GetCategories() []ProductCategory {
	if o == nil || o.Categories == nil {
		var ret []ProductCategory
		return ret
	}
	return *o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product) GetCategoriesOk() (*[]ProductCategory, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *Product) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []ProductCategory and assigns it to the Categories field.
func (o *Product) SetCategories(v []ProductCategory) {
	o.Categories = &v
}

func (o Product) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Variants != nil {
		toSerialize["variants"] = o.Variants
	}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	return json.Marshal(toSerialize)
}

type NullableProduct struct {
	value *Product
	isSet bool
}

func (v NullableProduct) Get() *Product {
	return v.value
}

func (v *NullableProduct) Set(val *Product) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct(val *Product) *NullableProduct {
	return &NullableProduct{value: val, isSet: true}
}

func (v NullableProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


