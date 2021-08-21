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

// OrderItem A product associated with an order.
type OrderItem struct {
	// ID of the item. UUID Version 4.
	Id string `json:"id"`
	// The title of the item.
	Title *string `json:"title,omitempty"`
	// The description of the item.
	Description *string `json:"description,omitempty"`
	// A positive integer in the smallest currency unit (that is, 100 cents for $1.00) representing the cost of a single instance of this item. This is multiplied by the quantity to calculate what was or will be charged.
	Amount *int32 `json:"amount,omitempty"`
	// A positive integer representing the number of instances of this item that was or will be charged.
	Quantity *int32 `json:"quantity,omitempty"`
}

// NewOrderItem instantiates a new OrderItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderItem(id string) *OrderItem {
	this := OrderItem{}
	this.Id = id
	return &this
}

// NewOrderItemWithDefaults instantiates a new OrderItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderItemWithDefaults() *OrderItem {
	this := OrderItem{}
	return &this
}

// GetId returns the Id field value
func (o *OrderItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrderItem) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrderItem) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *OrderItem) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *OrderItem) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *OrderItem) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrderItem) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrderItem) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrderItem) SetDescription(v string) {
	o.Description = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *OrderItem) GetAmount() int32 {
	if o == nil || o.Amount == nil {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetAmountOk() (*int32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *OrderItem) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *OrderItem) SetAmount(v int32) {
	o.Amount = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *OrderItem) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderItem) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *OrderItem) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *OrderItem) SetQuantity(v int32) {
	o.Quantity = &v
}

func (o OrderItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	return json.Marshal(toSerialize)
}

type NullableOrderItem struct {
	value *OrderItem
	isSet bool
}

func (v NullableOrderItem) Get() *OrderItem {
	return v.value
}

func (v *NullableOrderItem) Set(val *OrderItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderItem(val *OrderItem) *NullableOrderItem {
	return &NullableOrderItem{value: val, isSet: true}
}

func (v NullableOrderItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


