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

// ListingBuilding Structural data of the primary building on a listing. Includes everything from square footage of certain spaces to construction dates. 
type ListingBuilding struct {
	// Total number of bedrooms.
	Bedrooms NullableInt32 `json:"bedrooms,omitempty"`
	// Sum of the number of bathrooms. Represented as a number in order to support half-baths.
	Bathrooms NullableFloat32 `json:"bathrooms,omitempty"`
	// Total number of square feet.
	SquareFeet NullableFloat32 `json:"square_feet,omitempty"`
	// Year the property was built.
	YearBuilt NullableInt32 `json:"year_built,omitempty"`
}

// NewListingBuilding instantiates a new ListingBuilding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListingBuilding() *ListingBuilding {
	this := ListingBuilding{}
	return &this
}

// NewListingBuildingWithDefaults instantiates a new ListingBuilding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListingBuildingWithDefaults() *ListingBuilding {
	this := ListingBuilding{}
	return &this
}

// GetBedrooms returns the Bedrooms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingBuilding) GetBedrooms() int32 {
	if o == nil || o.Bedrooms.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Bedrooms.Get()
}

// GetBedroomsOk returns a tuple with the Bedrooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingBuilding) GetBedroomsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bedrooms.Get(), o.Bedrooms.IsSet()
}

// HasBedrooms returns a boolean if a field has been set.
func (o *ListingBuilding) HasBedrooms() bool {
	if o != nil && o.Bedrooms.IsSet() {
		return true
	}

	return false
}

// SetBedrooms gets a reference to the given NullableInt32 and assigns it to the Bedrooms field.
func (o *ListingBuilding) SetBedrooms(v int32) {
	o.Bedrooms.Set(&v)
}
// SetBedroomsNil sets the value for Bedrooms to be an explicit nil
func (o *ListingBuilding) SetBedroomsNil() {
	o.Bedrooms.Set(nil)
}

// UnsetBedrooms ensures that no value is present for Bedrooms, not even an explicit nil
func (o *ListingBuilding) UnsetBedrooms() {
	o.Bedrooms.Unset()
}

// GetBathrooms returns the Bathrooms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingBuilding) GetBathrooms() float32 {
	if o == nil || o.Bathrooms.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Bathrooms.Get()
}

// GetBathroomsOk returns a tuple with the Bathrooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingBuilding) GetBathroomsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bathrooms.Get(), o.Bathrooms.IsSet()
}

// HasBathrooms returns a boolean if a field has been set.
func (o *ListingBuilding) HasBathrooms() bool {
	if o != nil && o.Bathrooms.IsSet() {
		return true
	}

	return false
}

// SetBathrooms gets a reference to the given NullableFloat32 and assigns it to the Bathrooms field.
func (o *ListingBuilding) SetBathrooms(v float32) {
	o.Bathrooms.Set(&v)
}
// SetBathroomsNil sets the value for Bathrooms to be an explicit nil
func (o *ListingBuilding) SetBathroomsNil() {
	o.Bathrooms.Set(nil)
}

// UnsetBathrooms ensures that no value is present for Bathrooms, not even an explicit nil
func (o *ListingBuilding) UnsetBathrooms() {
	o.Bathrooms.Unset()
}

// GetSquareFeet returns the SquareFeet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingBuilding) GetSquareFeet() float32 {
	if o == nil || o.SquareFeet.Get() == nil {
		var ret float32
		return ret
	}
	return *o.SquareFeet.Get()
}

// GetSquareFeetOk returns a tuple with the SquareFeet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingBuilding) GetSquareFeetOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SquareFeet.Get(), o.SquareFeet.IsSet()
}

// HasSquareFeet returns a boolean if a field has been set.
func (o *ListingBuilding) HasSquareFeet() bool {
	if o != nil && o.SquareFeet.IsSet() {
		return true
	}

	return false
}

// SetSquareFeet gets a reference to the given NullableFloat32 and assigns it to the SquareFeet field.
func (o *ListingBuilding) SetSquareFeet(v float32) {
	o.SquareFeet.Set(&v)
}
// SetSquareFeetNil sets the value for SquareFeet to be an explicit nil
func (o *ListingBuilding) SetSquareFeetNil() {
	o.SquareFeet.Set(nil)
}

// UnsetSquareFeet ensures that no value is present for SquareFeet, not even an explicit nil
func (o *ListingBuilding) UnsetSquareFeet() {
	o.SquareFeet.Unset()
}

// GetYearBuilt returns the YearBuilt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingBuilding) GetYearBuilt() int32 {
	if o == nil || o.YearBuilt.Get() == nil {
		var ret int32
		return ret
	}
	return *o.YearBuilt.Get()
}

// GetYearBuiltOk returns a tuple with the YearBuilt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingBuilding) GetYearBuiltOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.YearBuilt.Get(), o.YearBuilt.IsSet()
}

// HasYearBuilt returns a boolean if a field has been set.
func (o *ListingBuilding) HasYearBuilt() bool {
	if o != nil && o.YearBuilt.IsSet() {
		return true
	}

	return false
}

// SetYearBuilt gets a reference to the given NullableInt32 and assigns it to the YearBuilt field.
func (o *ListingBuilding) SetYearBuilt(v int32) {
	o.YearBuilt.Set(&v)
}
// SetYearBuiltNil sets the value for YearBuilt to be an explicit nil
func (o *ListingBuilding) SetYearBuiltNil() {
	o.YearBuilt.Set(nil)
}

// UnsetYearBuilt ensures that no value is present for YearBuilt, not even an explicit nil
func (o *ListingBuilding) UnsetYearBuilt() {
	o.YearBuilt.Unset()
}

func (o ListingBuilding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bedrooms.IsSet() {
		toSerialize["bedrooms"] = o.Bedrooms.Get()
	}
	if o.Bathrooms.IsSet() {
		toSerialize["bathrooms"] = o.Bathrooms.Get()
	}
	if o.SquareFeet.IsSet() {
		toSerialize["square_feet"] = o.SquareFeet.Get()
	}
	if o.YearBuilt.IsSet() {
		toSerialize["year_built"] = o.YearBuilt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableListingBuilding struct {
	value *ListingBuilding
	isSet bool
}

func (v NullableListingBuilding) Get() *ListingBuilding {
	return v.value
}

func (v *NullableListingBuilding) Set(val *ListingBuilding) {
	v.value = val
	v.isSet = true
}

func (v NullableListingBuilding) IsSet() bool {
	return v.isSet
}

func (v *NullableListingBuilding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListingBuilding(val *ListingBuilding) *NullableListingBuilding {
	return &NullableListingBuilding{value: val, isSet: true}
}

func (v NullableListingBuilding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListingBuilding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


