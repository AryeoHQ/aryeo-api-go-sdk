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

// ListingPrice Valuation data relating to the price of a listing.
type ListingPrice struct {
	// The current price of the listing.
	ListPrice NullableInt32 `json:"list_price,omitempty"`
}

// NewListingPrice instantiates a new ListingPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListingPrice() *ListingPrice {
	this := ListingPrice{}
	return &this
}

// NewListingPriceWithDefaults instantiates a new ListingPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListingPriceWithDefaults() *ListingPrice {
	this := ListingPrice{}
	return &this
}

// GetListPrice returns the ListPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingPrice) GetListPrice() int32 {
	if o == nil || o.ListPrice.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ListPrice.Get()
}

// GetListPriceOk returns a tuple with the ListPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingPrice) GetListPriceOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ListPrice.Get(), o.ListPrice.IsSet()
}

// HasListPrice returns a boolean if a field has been set.
func (o *ListingPrice) HasListPrice() bool {
	if o != nil && o.ListPrice.IsSet() {
		return true
	}

	return false
}

// SetListPrice gets a reference to the given NullableInt32 and assigns it to the ListPrice field.
func (o *ListingPrice) SetListPrice(v int32) {
	o.ListPrice.Set(&v)
}
// SetListPriceNil sets the value for ListPrice to be an explicit nil
func (o *ListingPrice) SetListPriceNil() {
	o.ListPrice.Set(nil)
}

// UnsetListPrice ensures that no value is present for ListPrice, not even an explicit nil
func (o *ListingPrice) UnsetListPrice() {
	o.ListPrice.Unset()
}

func (o ListingPrice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ListPrice.IsSet() {
		toSerialize["list_price"] = o.ListPrice.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableListingPrice struct {
	value *ListingPrice
	isSet bool
}

func (v NullableListingPrice) Get() *ListingPrice {
	return v.value
}

func (v *NullableListingPrice) Set(val *ListingPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableListingPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableListingPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListingPrice(val *ListingPrice) *NullableListingPrice {
	return &NullableListingPrice{value: val, isSet: true}
}

func (v NullableListingPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListingPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


