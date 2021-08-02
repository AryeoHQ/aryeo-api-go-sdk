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

// ListingLot Parcel data of the lot of a listing. Includes nearly everything about the land that makes up the listing.
type ListingLot struct {
	// Total area of the lot of a listing in acres. 
	SizeAcres NullableFloat32 `json:"size_acres,omitempty"`
	// Number of parking spaces.
	OpenParkingSpaces NullableFloat32 `json:"open_parking_spaces,omitempty"`
}

// NewListingLot instantiates a new ListingLot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListingLot() *ListingLot {
	this := ListingLot{}
	return &this
}

// NewListingLotWithDefaults instantiates a new ListingLot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListingLotWithDefaults() *ListingLot {
	this := ListingLot{}
	return &this
}

// GetSizeAcres returns the SizeAcres field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingLot) GetSizeAcres() float32 {
	if o == nil || o.SizeAcres.Get() == nil {
		var ret float32
		return ret
	}
	return *o.SizeAcres.Get()
}

// GetSizeAcresOk returns a tuple with the SizeAcres field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingLot) GetSizeAcresOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SizeAcres.Get(), o.SizeAcres.IsSet()
}

// HasSizeAcres returns a boolean if a field has been set.
func (o *ListingLot) HasSizeAcres() bool {
	if o != nil && o.SizeAcres.IsSet() {
		return true
	}

	return false
}

// SetSizeAcres gets a reference to the given NullableFloat32 and assigns it to the SizeAcres field.
func (o *ListingLot) SetSizeAcres(v float32) {
	o.SizeAcres.Set(&v)
}
// SetSizeAcresNil sets the value for SizeAcres to be an explicit nil
func (o *ListingLot) SetSizeAcresNil() {
	o.SizeAcres.Set(nil)
}

// UnsetSizeAcres ensures that no value is present for SizeAcres, not even an explicit nil
func (o *ListingLot) UnsetSizeAcres() {
	o.SizeAcres.Unset()
}

// GetOpenParkingSpaces returns the OpenParkingSpaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingLot) GetOpenParkingSpaces() float32 {
	if o == nil || o.OpenParkingSpaces.Get() == nil {
		var ret float32
		return ret
	}
	return *o.OpenParkingSpaces.Get()
}

// GetOpenParkingSpacesOk returns a tuple with the OpenParkingSpaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingLot) GetOpenParkingSpacesOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OpenParkingSpaces.Get(), o.OpenParkingSpaces.IsSet()
}

// HasOpenParkingSpaces returns a boolean if a field has been set.
func (o *ListingLot) HasOpenParkingSpaces() bool {
	if o != nil && o.OpenParkingSpaces.IsSet() {
		return true
	}

	return false
}

// SetOpenParkingSpaces gets a reference to the given NullableFloat32 and assigns it to the OpenParkingSpaces field.
func (o *ListingLot) SetOpenParkingSpaces(v float32) {
	o.OpenParkingSpaces.Set(&v)
}
// SetOpenParkingSpacesNil sets the value for OpenParkingSpaces to be an explicit nil
func (o *ListingLot) SetOpenParkingSpacesNil() {
	o.OpenParkingSpaces.Set(nil)
}

// UnsetOpenParkingSpaces ensures that no value is present for OpenParkingSpaces, not even an explicit nil
func (o *ListingLot) UnsetOpenParkingSpaces() {
	o.OpenParkingSpaces.Unset()
}

func (o ListingLot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SizeAcres.IsSet() {
		toSerialize["size_acres"] = o.SizeAcres.Get()
	}
	if o.OpenParkingSpaces.IsSet() {
		toSerialize["open_parking_spaces"] = o.OpenParkingSpaces.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableListingLot struct {
	value *ListingLot
	isSet bool
}

func (v NullableListingLot) Get() *ListingLot {
	return v.value
}

func (v *NullableListingLot) Set(val *ListingLot) {
	v.value = val
	v.isSet = true
}

func (v NullableListingLot) IsSet() bool {
	return v.isSet
}

func (v *NullableListingLot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListingLot(val *ListingLot) *NullableListingLot {
	return &NullableListingLot{value: val, isSet: true}
}

func (v NullableListingLot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListingLot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


