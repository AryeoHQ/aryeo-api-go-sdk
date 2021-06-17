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

// Order A payment request for some content or service.
type Order struct {
	// ID of the order.
	Id string `json:"id"`
	// A vanity id to be displayed for the order. For example, \"Order #1000\".
	DisplayId int32 `json:"display_id"`
	// The total price of the order given in cents.
	TotalPriceCents int32 `json:"total_price_cents"`
	Currency Currency `json:"currency"`
	// The payment status of the order.
	PaymentStatus string `json:"payment_status"`
	// A URL for to pay for the order.
	PaymentUrl NullableString `json:"payment_url,omitempty"`
	Listing *PartialListing `json:"listing,omitempty"`
	// The fulfillment status of the order.
	FulfillmentStatus string `json:"fulfillment_status"`
}

// NewOrder instantiates a new Order object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrder(id string, displayId int32, totalPriceCents int32, currency Currency, paymentStatus string, fulfillmentStatus string) *Order {
	this := Order{}
	this.Id = id
	this.DisplayId = displayId
	this.TotalPriceCents = totalPriceCents
	this.Currency = currency
	this.PaymentStatus = paymentStatus
	this.FulfillmentStatus = fulfillmentStatus
	return &this
}

// NewOrderWithDefaults instantiates a new Order object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderWithDefaults() *Order {
	this := Order{}
	return &this
}

// GetId returns the Id field value
func (o *Order) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Order) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Order) SetId(v string) {
	o.Id = v
}

// GetDisplayId returns the DisplayId field value
func (o *Order) GetDisplayId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value
// and a boolean to check if the value has been set.
func (o *Order) GetDisplayIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayId, true
}

// SetDisplayId sets field value
func (o *Order) SetDisplayId(v int32) {
	o.DisplayId = v
}

// GetTotalPriceCents returns the TotalPriceCents field value
func (o *Order) GetTotalPriceCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalPriceCents
}

// GetTotalPriceCentsOk returns a tuple with the TotalPriceCents field value
// and a boolean to check if the value has been set.
func (o *Order) GetTotalPriceCentsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalPriceCents, true
}

// SetTotalPriceCents sets field value
func (o *Order) SetTotalPriceCents(v int32) {
	o.TotalPriceCents = v
}

// GetCurrency returns the Currency field value
func (o *Order) GetCurrency() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Order) GetCurrencyOk() (*Currency, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Order) SetCurrency(v Currency) {
	o.Currency = v
}

// GetPaymentStatus returns the PaymentStatus field value
func (o *Order) GetPaymentStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentStatus
}

// GetPaymentStatusOk returns a tuple with the PaymentStatus field value
// and a boolean to check if the value has been set.
func (o *Order) GetPaymentStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentStatus, true
}

// SetPaymentStatus sets field value
func (o *Order) SetPaymentStatus(v string) {
	o.PaymentStatus = v
}

// GetPaymentUrl returns the PaymentUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Order) GetPaymentUrl() string {
	if o == nil || o.PaymentUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.PaymentUrl.Get()
}

// GetPaymentUrlOk returns a tuple with the PaymentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Order) GetPaymentUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentUrl.Get(), o.PaymentUrl.IsSet()
}

// HasPaymentUrl returns a boolean if a field has been set.
func (o *Order) HasPaymentUrl() bool {
	if o != nil && o.PaymentUrl.IsSet() {
		return true
	}

	return false
}

// SetPaymentUrl gets a reference to the given NullableString and assigns it to the PaymentUrl field.
func (o *Order) SetPaymentUrl(v string) {
	o.PaymentUrl.Set(&v)
}
// SetPaymentUrlNil sets the value for PaymentUrl to be an explicit nil
func (o *Order) SetPaymentUrlNil() {
	o.PaymentUrl.Set(nil)
}

// UnsetPaymentUrl ensures that no value is present for PaymentUrl, not even an explicit nil
func (o *Order) UnsetPaymentUrl() {
	o.PaymentUrl.Unset()
}

// GetListing returns the Listing field value if set, zero value otherwise.
func (o *Order) GetListing() PartialListing {
	if o == nil || o.Listing == nil {
		var ret PartialListing
		return ret
	}
	return *o.Listing
}

// GetListingOk returns a tuple with the Listing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetListingOk() (*PartialListing, bool) {
	if o == nil || o.Listing == nil {
		return nil, false
	}
	return o.Listing, true
}

// HasListing returns a boolean if a field has been set.
func (o *Order) HasListing() bool {
	if o != nil && o.Listing != nil {
		return true
	}

	return false
}

// SetListing gets a reference to the given PartialListing and assigns it to the Listing field.
func (o *Order) SetListing(v PartialListing) {
	o.Listing = &v
}

// GetFulfillmentStatus returns the FulfillmentStatus field value
func (o *Order) GetFulfillmentStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FulfillmentStatus
}

// GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field value
// and a boolean to check if the value has been set.
func (o *Order) GetFulfillmentStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FulfillmentStatus, true
}

// SetFulfillmentStatus sets field value
func (o *Order) SetFulfillmentStatus(v string) {
	o.FulfillmentStatus = v
}

func (o Order) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["display_id"] = o.DisplayId
	}
	if true {
		toSerialize["total_price_cents"] = o.TotalPriceCents
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["payment_status"] = o.PaymentStatus
	}
	if o.PaymentUrl.IsSet() {
		toSerialize["payment_url"] = o.PaymentUrl.Get()
	}
	if o.Listing != nil {
		toSerialize["listing"] = o.Listing
	}
	if true {
		toSerialize["fulfillment_status"] = o.FulfillmentStatus
	}
	return json.Marshal(toSerialize)
}

type NullableOrder struct {
	value *Order
	isSet bool
}

func (v NullableOrder) Get() *Order {
	return v.value
}

func (v *NullableOrder) Set(val *Order) {
	v.value = val
	v.isSet = true
}

func (v NullableOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrder(val *Order) *NullableOrder {
	return &NullableOrder{value: val, isSet: true}
}

func (v NullableOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


