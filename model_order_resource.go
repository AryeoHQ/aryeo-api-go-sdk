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

// OrderResource An order.
type OrderResource struct {
	Data *Order `json:"data,omitempty"`
}

// NewOrderResource instantiates a new OrderResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResource() *OrderResource {
	this := OrderResource{}
	return &this
}

// NewOrderResourceWithDefaults instantiates a new OrderResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResourceWithDefaults() *OrderResource {
	this := OrderResource{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderResource) GetData() Order {
	if o == nil || o.Data == nil {
		var ret Order
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResource) GetDataOk() (*Order, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderResource) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given Order and assigns it to the Data field.
func (o *OrderResource) SetData(v Order) {
	o.Data = &v
}

func (o OrderResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableOrderResource struct {
	value *OrderResource
	isSet bool
}

func (v NullableOrderResource) Get() *OrderResource {
	return v.value
}

func (v *NullableOrderResource) Set(val *OrderResource) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResource) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResource(val *OrderResource) *NullableOrderResource {
	return &NullableOrderResource{value: val, isSet: true}
}

func (v NullableOrderResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


