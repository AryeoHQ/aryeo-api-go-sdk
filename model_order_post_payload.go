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

// OrderPostPayload Payload for creating an order.
type OrderPostPayload struct {
	// The fulfillment status of the order. Defaults to \"UNFULFILLED\".
	FulfillmentStatus NullableString `json:"fulfillment_status,omitempty"`
	// Internal notes that will be attached to the order. Viewable only by the team.
	InternalNotes NullableString `json:"internal_notes,omitempty"`
	// The payment status of the order. Defaults to \"UNPAID\". 
	PaymentStatus NullableString `json:"payment_status,omitempty"`
	// ID of the address to associate with the order. UUID Version 4.
	AddressId *string `json:"address_id,omitempty"`
	// ID of the customer to associate with the order. UUID Version 4.
	CustomerId *string `json:"customer_id,omitempty"`
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

// GetInternalNotes returns the InternalNotes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderPostPayload) GetInternalNotes() string {
	if o == nil || o.InternalNotes.Get() == nil {
		var ret string
		return ret
	}
	return *o.InternalNotes.Get()
}

// GetInternalNotesOk returns a tuple with the InternalNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderPostPayload) GetInternalNotesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InternalNotes.Get(), o.InternalNotes.IsSet()
}

// HasInternalNotes returns a boolean if a field has been set.
func (o *OrderPostPayload) HasInternalNotes() bool {
	if o != nil && o.InternalNotes.IsSet() {
		return true
	}

	return false
}

// SetInternalNotes gets a reference to the given NullableString and assigns it to the InternalNotes field.
func (o *OrderPostPayload) SetInternalNotes(v string) {
	o.InternalNotes.Set(&v)
}
// SetInternalNotesNil sets the value for InternalNotes to be an explicit nil
func (o *OrderPostPayload) SetInternalNotesNil() {
	o.InternalNotes.Set(nil)
}

// UnsetInternalNotes ensures that no value is present for InternalNotes, not even an explicit nil
func (o *OrderPostPayload) UnsetInternalNotes() {
	o.InternalNotes.Unset()
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

// GetAddressId returns the AddressId field value if set, zero value otherwise.
func (o *OrderPostPayload) GetAddressId() string {
	if o == nil || o.AddressId == nil {
		var ret string
		return ret
	}
	return *o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPostPayload) GetAddressIdOk() (*string, bool) {
	if o == nil || o.AddressId == nil {
		return nil, false
	}
	return o.AddressId, true
}

// HasAddressId returns a boolean if a field has been set.
func (o *OrderPostPayload) HasAddressId() bool {
	if o != nil && o.AddressId != nil {
		return true
	}

	return false
}

// SetAddressId gets a reference to the given string and assigns it to the AddressId field.
func (o *OrderPostPayload) SetAddressId(v string) {
	o.AddressId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *OrderPostPayload) GetCustomerId() string {
	if o == nil || o.CustomerId == nil {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderPostPayload) GetCustomerIdOk() (*string, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *OrderPostPayload) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *OrderPostPayload) SetCustomerId(v string) {
	o.CustomerId = &v
}

func (o OrderPostPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FulfillmentStatus.IsSet() {
		toSerialize["fulfillment_status"] = o.FulfillmentStatus.Get()
	}
	if o.InternalNotes.IsSet() {
		toSerialize["internal_notes"] = o.InternalNotes.Get()
	}
	if o.PaymentStatus.IsSet() {
		toSerialize["payment_status"] = o.PaymentStatus.Get()
	}
	if o.AddressId != nil {
		toSerialize["address_id"] = o.AddressId
	}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
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


