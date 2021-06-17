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

// OrderPostPayload Payload for creating an order.
type OrderPostPayload struct {
	// The fulfillment status of the order.
	FulfillmentStatus NullableString `json:"fulfillment_status,omitempty"`
	// The payment status of the order.
	PaymentStatus NullableString `json:"payment_status,omitempty"`
	// product_items
	ProductItems *[]ProductItem `json:"product_items,omitempty"`
}

// NewOrderPostPayload instantiates a new OrderPostPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderPostPayload() *OrderPostPayload {
	this := OrderPostPayload{}
	return &this
}

// NewOrderPostPayloadWithDefaults instantiates a new OrderPostPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderPostPayloadWithDefaults() *OrderPostPayload {
	this := OrderPostPayload{}
	return &this
}

// GetFulfillmentStatus returns the FulfillmentStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderPostPayload) GetFulfillmentStatus() string {
	if o == nil || o.FulfillmentStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.FulfillmentStatus.Get()
}

// GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderPostPayload) GetFulfillmentStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FulfillmentStatus.Get(), o.FulfillmentStatus.IsSet()
}

// HasFulfillmentStatus returns a boolean if a field has been set.
func (o *OrderPostPayload) HasFulfillmentStatus() bool {
	if o != nil && o.FulfillmentStatus.IsSet() {
		return true
	}

	return false
}

// SetFulfillmentStatus gets a reference to the given NullableString and assigns it to the FulfillmentStatus field.
func (o *OrderPostPayload) SetFulfillmentStatus(v string) {
	o.FulfillmentStatus.Set(&v)
}
// SetFulfillmentStatusNil sets the value for FulfillmentStatus to be an explicit nil
func (o *OrderPostPayload) SetFulfillmentStatusNil() {
	o.FulfillmentStatus.Set(nil)
}

// UnsetFulfillmentStatus ensures that no value is present for FulfillmentStatus, not even an explicit nil
func (o *OrderPostPayload) UnsetFulfillmentStatus() {
	o.FulfillmentStatus.Unset()
}

// GetPaymentStatus returns the PaymentStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderPostPayload) GetPaymentStatus() string {
	if o == nil || o.PaymentStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.PaymentStatus.Get()
}

// GetPaymentStatusOk returns a tuple with the PaymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderPostPayload) GetPaymentStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentStatus.Get(), o.PaymentStatus.IsSet()
}

// HasPaymentStatus returns a boolean if a field has been set.
func (o *OrderPostPayload) HasPaymentStatus() bool {
	if o != nil && o.PaymentStatus.IsSet() {
		return true
	}

	return false
}

// SetPaymentStatus gets a reference to the given NullableString and assigns it to the PaymentStatus field.
func (o *OrderPostPayload) SetPaymentStatus(v string) {
	o.PaymentStatus.Set(&v)
}
// SetPaymentStatusNil sets the value for PaymentStatus to be an explicit nil
func (o *OrderPostPayload) SetPaymentStatusNil() {
	o.PaymentStatus.Set(nil)
}

// UnsetPaymentStatus ensures that no value is present for PaymentStatus, not even an explicit nil
func (o *OrderPostPayload) UnsetPaymentStatus() {
	o.PaymentStatus.Unset()
}

// GetProductItems returns the ProductItems field value if set, zero value otherwise.
func (o *OrderPostPayload) GetProductItems() []ProductItem {
	if o == nil || o.ProductItems == nil {
		var ret []ProductItem
		return ret
	}
	return *o.ProductItems
}

// GetProductItemsOk returns a tuple with the ProductItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPostPayload) GetProductItemsOk() (*[]ProductItem, bool) {
	if o == nil || o.ProductItems == nil {
		return nil, false
	}
	return o.ProductItems, true
}

// HasProductItems returns a boolean if a field has been set.
func (o *OrderPostPayload) HasProductItems() bool {
	if o != nil && o.ProductItems != nil {
		return true
	}

	return false
}

// SetProductItems gets a reference to the given []ProductItem and assigns it to the ProductItems field.
func (o *OrderPostPayload) SetProductItems(v []ProductItem) {
	o.ProductItems = &v
}

func (o OrderPostPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FulfillmentStatus.IsSet() {
		toSerialize["fulfillment_status"] = o.FulfillmentStatus.Get()
	}
	if o.PaymentStatus.IsSet() {
		toSerialize["payment_status"] = o.PaymentStatus.Get()
	}
	if o.ProductItems != nil {
		toSerialize["product_items"] = o.ProductItems
	}
	return json.Marshal(toSerialize)
}

type NullableOrderPostPayload struct {
	value *OrderPostPayload
	isSet bool
}

func (v NullableOrderPostPayload) Get() *OrderPostPayload {
	return v.value
}

func (v *NullableOrderPostPayload) Set(val *OrderPostPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderPostPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderPostPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderPostPayload(val *OrderPostPayload) *NullableOrderPostPayload {
	return &NullableOrderPostPayload{value: val, isSet: true}
}

func (v NullableOrderPostPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderPostPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


