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
	// ID of the address to associate with the order. UUID Version 4.
	AddressId *string `json:"address_id,omitempty"`
	// ID of the customer to associate with the order. UUID Version 4.
	CustomerId *string `json:"customer_id,omitempty"`
	// Indicates if the customer and creator notifications should be sent when creating the order. Requires an address and customer to be set in order for the notifications to be sent.
	Notify NullableBool `json:"notify,omitempty"`
	// Indicates if the downloads for the attached listing should be locked while there is an outstanding balance on the order.
	LockDownloadForPayment NullableBool `json:"lock_download_for_payment,omitempty"`
	// Indicates if the order will allow payments from the customer before the order is marked as fulfilled.
	AllowPaymentsBeforeFulfillment NullableBool `json:"allow_payments_before_fulfillment,omitempty"`
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

// GetNotify returns the Notify field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderPostPayload) GetNotify() bool {
	if o == nil || o.Notify.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Notify.Get()
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderPostPayload) GetNotifyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Notify.Get(), o.Notify.IsSet()
}

// HasNotify returns a boolean if a field has been set.
func (o *OrderPostPayload) HasNotify() bool {
	if o != nil && o.Notify.IsSet() {
		return true
	}

	return false
}

// SetNotify gets a reference to the given NullableBool and assigns it to the Notify field.
func (o *OrderPostPayload) SetNotify(v bool) {
	o.Notify.Set(&v)
}
// SetNotifyNil sets the value for Notify to be an explicit nil
func (o *OrderPostPayload) SetNotifyNil() {
	o.Notify.Set(nil)
}

// UnsetNotify ensures that no value is present for Notify, not even an explicit nil
func (o *OrderPostPayload) UnsetNotify() {
	o.Notify.Unset()
}

// GetLockDownloadForPayment returns the LockDownloadForPayment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderPostPayload) GetLockDownloadForPayment() bool {
	if o == nil || o.LockDownloadForPayment.Get() == nil {
		var ret bool
		return ret
	}
	return *o.LockDownloadForPayment.Get()
}

// GetLockDownloadForPaymentOk returns a tuple with the LockDownloadForPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderPostPayload) GetLockDownloadForPaymentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LockDownloadForPayment.Get(), o.LockDownloadForPayment.IsSet()
}

// HasLockDownloadForPayment returns a boolean if a field has been set.
func (o *OrderPostPayload) HasLockDownloadForPayment() bool {
	if o != nil && o.LockDownloadForPayment.IsSet() {
		return true
	}

	return false
}

// SetLockDownloadForPayment gets a reference to the given NullableBool and assigns it to the LockDownloadForPayment field.
func (o *OrderPostPayload) SetLockDownloadForPayment(v bool) {
	o.LockDownloadForPayment.Set(&v)
}
// SetLockDownloadForPaymentNil sets the value for LockDownloadForPayment to be an explicit nil
func (o *OrderPostPayload) SetLockDownloadForPaymentNil() {
	o.LockDownloadForPayment.Set(nil)
}

// UnsetLockDownloadForPayment ensures that no value is present for LockDownloadForPayment, not even an explicit nil
func (o *OrderPostPayload) UnsetLockDownloadForPayment() {
	o.LockDownloadForPayment.Unset()
}

// GetAllowPaymentsBeforeFulfillment returns the AllowPaymentsBeforeFulfillment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderPostPayload) GetAllowPaymentsBeforeFulfillment() bool {
	if o == nil || o.AllowPaymentsBeforeFulfillment.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AllowPaymentsBeforeFulfillment.Get()
}

// GetAllowPaymentsBeforeFulfillmentOk returns a tuple with the AllowPaymentsBeforeFulfillment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderPostPayload) GetAllowPaymentsBeforeFulfillmentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AllowPaymentsBeforeFulfillment.Get(), o.AllowPaymentsBeforeFulfillment.IsSet()
}

// HasAllowPaymentsBeforeFulfillment returns a boolean if a field has been set.
func (o *OrderPostPayload) HasAllowPaymentsBeforeFulfillment() bool {
	if o != nil && o.AllowPaymentsBeforeFulfillment.IsSet() {
		return true
	}

	return false
}

// SetAllowPaymentsBeforeFulfillment gets a reference to the given NullableBool and assigns it to the AllowPaymentsBeforeFulfillment field.
func (o *OrderPostPayload) SetAllowPaymentsBeforeFulfillment(v bool) {
	o.AllowPaymentsBeforeFulfillment.Set(&v)
}
// SetAllowPaymentsBeforeFulfillmentNil sets the value for AllowPaymentsBeforeFulfillment to be an explicit nil
func (o *OrderPostPayload) SetAllowPaymentsBeforeFulfillmentNil() {
	o.AllowPaymentsBeforeFulfillment.Set(nil)
}

// UnsetAllowPaymentsBeforeFulfillment ensures that no value is present for AllowPaymentsBeforeFulfillment, not even an explicit nil
func (o *OrderPostPayload) UnsetAllowPaymentsBeforeFulfillment() {
	o.AllowPaymentsBeforeFulfillment.Unset()
}

func (o OrderPostPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FulfillmentStatus.IsSet() {
		toSerialize["fulfillment_status"] = o.FulfillmentStatus.Get()
	}
	if o.InternalNotes.IsSet() {
		toSerialize["internal_notes"] = o.InternalNotes.Get()
	}
	if o.AddressId != nil {
		toSerialize["address_id"] = o.AddressId
	}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
	}
	if o.Notify.IsSet() {
		toSerialize["notify"] = o.Notify.Get()
	}
	if o.LockDownloadForPayment.IsSet() {
		toSerialize["lock_download_for_payment"] = o.LockDownloadForPayment.Get()
	}
	if o.AllowPaymentsBeforeFulfillment.IsSet() {
		toSerialize["allow_payments_before_fulfillment"] = o.AllowPaymentsBeforeFulfillment.Get()
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


