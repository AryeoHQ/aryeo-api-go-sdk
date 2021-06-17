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

// PartialAddress A structure containing a street address and additional metadata about a location.
type PartialAddress struct {
	// ID of address.
	Id int32 `json:"id"`
	// The full address string containing address_1 and address_2.
	FullAddress NullableString `json:"full_address,omitempty"`
	// A formatted address string containing the street.
	FormattedAddress1 string `json:"formatted_address_1"`
	// A formatted address string containing the city, state, and zipcode.
	FormattedAddress2 string `json:"formatted_address_2"`
	// Latitude of the address.
	Latitude float64 `json:"latitude"`
	// Longitude of the address.
	Longitude float64 `json:"longitude"`
	// ID of place.
	PlaceId NullableString `json:"place_id,omitempty"`
	// Address line 1
	AddressLine1 NullableString `json:"address_line_1,omitempty"`
	// Address line 2
	AddressLine2 NullableString `json:"address_line_2,omitempty"`
	// City
	City NullableString `json:"city,omitempty"`
	// State
	State NullableString `json:"state,omitempty"`
	// Postal Code
	PostalCode NullableString `json:"postal_code,omitempty"`
}

// NewPartialAddress instantiates a new PartialAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialAddress(id int32, formattedAddress1 string, formattedAddress2 string, latitude float64, longitude float64) *PartialAddress {
	this := PartialAddress{}
	this.Id = id
	this.FormattedAddress1 = formattedAddress1
	this.FormattedAddress2 = formattedAddress2
	this.Latitude = latitude
	this.Longitude = longitude
	return &this
}

// NewPartialAddressWithDefaults instantiates a new PartialAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialAddressWithDefaults() *PartialAddress {
	this := PartialAddress{}
	return &this
}

// GetId returns the Id field value
func (o *PartialAddress) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PartialAddress) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PartialAddress) SetId(v int32) {
	o.Id = v
}

// GetFullAddress returns the FullAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialAddress) GetFullAddress() string {
	if o == nil || o.FullAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.FullAddress.Get()
}

// GetFullAddressOk returns a tuple with the FullAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialAddress) GetFullAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FullAddress.Get(), o.FullAddress.IsSet()
}

// HasFullAddress returns a boolean if a field has been set.
func (o *PartialAddress) HasFullAddress() bool {
	if o != nil && o.FullAddress.IsSet() {
		return true
	}

	return false
}

// SetFullAddress gets a reference to the given NullableString and assigns it to the FullAddress field.
func (o *PartialAddress) SetFullAddress(v string) {
	o.FullAddress.Set(&v)
}
// SetFullAddressNil sets the value for FullAddress to be an explicit nil
func (o *PartialAddress) SetFullAddressNil() {
	o.FullAddress.Set(nil)
}

// UnsetFullAddress ensures that no value is present for FullAddress, not even an explicit nil
func (o *PartialAddress) UnsetFullAddress() {
	o.FullAddress.Unset()
}

// GetFormattedAddress1 returns the FormattedAddress1 field value
func (o *PartialAddress) GetFormattedAddress1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormattedAddress1
}

// GetFormattedAddress1Ok returns a tuple with the FormattedAddress1 field value
// and a boolean to check if the value has been set.
func (o *PartialAddress) GetFormattedAddress1Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FormattedAddress1, true
}

// SetFormattedAddress1 sets field value
func (o *PartialAddress) SetFormattedAddress1(v string) {
	o.FormattedAddress1 = v
}

// GetFormattedAddress2 returns the FormattedAddress2 field value
func (o *PartialAddress) GetFormattedAddress2() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormattedAddress2
}

// GetFormattedAddress2Ok returns a tuple with the FormattedAddress2 field value
// and a boolean to check if the value has been set.
func (o *PartialAddress) GetFormattedAddress2Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FormattedAddress2, true
}

// SetFormattedAddress2 sets field value
func (o *PartialAddress) SetFormattedAddress2(v string) {
	o.FormattedAddress2 = v
}

// GetLatitude returns the Latitude field value
func (o *PartialAddress) GetLatitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *PartialAddress) GetLatitudeOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *PartialAddress) SetLatitude(v float64) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *PartialAddress) GetLongitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *PartialAddress) GetLongitudeOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *PartialAddress) SetLongitude(v float64) {
	o.Longitude = v
}

// GetPlaceId returns the PlaceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialAddress) GetPlaceId() string {
	if o == nil || o.PlaceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PlaceId.Get()
}

// GetPlaceIdOk returns a tuple with the PlaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialAddress) GetPlaceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PlaceId.Get(), o.PlaceId.IsSet()
}

// HasPlaceId returns a boolean if a field has been set.
func (o *PartialAddress) HasPlaceId() bool {
	if o != nil && o.PlaceId.IsSet() {
		return true
	}

	return false
}

// SetPlaceId gets a reference to the given NullableString and assigns it to the PlaceId field.
func (o *PartialAddress) SetPlaceId(v string) {
	o.PlaceId.Set(&v)
}
// SetPlaceIdNil sets the value for PlaceId to be an explicit nil
func (o *PartialAddress) SetPlaceIdNil() {
	o.PlaceId.Set(nil)
}

// UnsetPlaceId ensures that no value is present for PlaceId, not even an explicit nil
func (o *PartialAddress) UnsetPlaceId() {
	o.PlaceId.Unset()
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialAddress) GetAddressLine1() string {
	if o == nil || o.AddressLine1.Get() == nil {
		var ret string
		return ret
	}
	return *o.AddressLine1.Get()
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialAddress) GetAddressLine1Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AddressLine1.Get(), o.AddressLine1.IsSet()
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *PartialAddress) HasAddressLine1() bool {
	if o != nil && o.AddressLine1.IsSet() {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given NullableString and assigns it to the AddressLine1 field.
func (o *PartialAddress) SetAddressLine1(v string) {
	o.AddressLine1.Set(&v)
}
// SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil
func (o *PartialAddress) SetAddressLine1Nil() {
	o.AddressLine1.Set(nil)
}

// UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
func (o *PartialAddress) UnsetAddressLine1() {
	o.AddressLine1.Unset()
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialAddress) GetAddressLine2() string {
	if o == nil || o.AddressLine2.Get() == nil {
		var ret string
		return ret
	}
	return *o.AddressLine2.Get()
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialAddress) GetAddressLine2Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AddressLine2.Get(), o.AddressLine2.IsSet()
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *PartialAddress) HasAddressLine2() bool {
	if o != nil && o.AddressLine2.IsSet() {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given NullableString and assigns it to the AddressLine2 field.
func (o *PartialAddress) SetAddressLine2(v string) {
	o.AddressLine2.Set(&v)
}
// SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil
func (o *PartialAddress) SetAddressLine2Nil() {
	o.AddressLine2.Set(nil)
}

// UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
func (o *PartialAddress) UnsetAddressLine2() {
	o.AddressLine2.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialAddress) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialAddress) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *PartialAddress) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *PartialAddress) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *PartialAddress) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *PartialAddress) UnsetCity() {
	o.City.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialAddress) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialAddress) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *PartialAddress) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *PartialAddress) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *PartialAddress) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *PartialAddress) UnsetState() {
	o.State.Unset()
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialAddress) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// HasPostalCode returns a boolean if a field has been set.
func (o *PartialAddress) HasPostalCode() bool {
	if o != nil && o.PostalCode.IsSet() {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given NullableString and assigns it to the PostalCode field.
func (o *PartialAddress) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}
// SetPostalCodeNil sets the value for PostalCode to be an explicit nil
func (o *PartialAddress) SetPostalCodeNil() {
	o.PostalCode.Set(nil)
}

// UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
func (o *PartialAddress) UnsetPostalCode() {
	o.PostalCode.Unset()
}

func (o PartialAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.FullAddress.IsSet() {
		toSerialize["full_address"] = o.FullAddress.Get()
	}
	if true {
		toSerialize["formatted_address_1"] = o.FormattedAddress1
	}
	if true {
		toSerialize["formatted_address_2"] = o.FormattedAddress2
	}
	if true {
		toSerialize["latitude"] = o.Latitude
	}
	if true {
		toSerialize["longitude"] = o.Longitude
	}
	if o.PlaceId.IsSet() {
		toSerialize["place_id"] = o.PlaceId.Get()
	}
	if o.AddressLine1.IsSet() {
		toSerialize["address_line_1"] = o.AddressLine1.Get()
	}
	if o.AddressLine2.IsSet() {
		toSerialize["address_line_2"] = o.AddressLine2.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.PostalCode.IsSet() {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePartialAddress struct {
	value *PartialAddress
	isSet bool
}

func (v NullablePartialAddress) Get() *PartialAddress {
	return v.value
}

func (v *NullablePartialAddress) Set(val *PartialAddress) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialAddress) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialAddress(val *PartialAddress) *NullablePartialAddress {
	return &NullablePartialAddress{value: val, isSet: true}
}

func (v NullablePartialAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


