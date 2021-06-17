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

// PropertyDetails Details about a real-estate listing.
type PropertyDetails struct {
	// The price of the property in dollars.
	Price NullableInt32 `json:"price,omitempty"`
	// A number assigned to a real-estate listing for the purposes of information sharing.
	MlsNumber NullableString `json:"mls_number,omitempty"`
	// Number of bedrooms.
	Bedrooms NullableInt32 `json:"bedrooms,omitempty"`
	// Number of bathrooms. Represented as a number in order to support half-baths.
	Bathrooms NullableFloat32 `json:"bathrooms,omitempty"`
	// Total number of square feet.
	SquareFeet NullableFloat32 `json:"square_feet,omitempty"`
	// Total acres.
	LotAcres NullableFloat32 `json:"lot_acres,omitempty"`
	// Number of parking spaces.
	ParkingSpots NullableInt32 `json:"parking_spots,omitempty"`
	// Year the property was built.
	YearBuilt NullableInt32 `json:"year_built,omitempty"`
	// Type of property.
	PropertyType NullableString `json:"property_type,omitempty"`
	// Property description.
	Description NullableString `json:"description,omitempty"`
}

// NewPropertyDetails instantiates a new PropertyDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyDetails() *PropertyDetails {
	this := PropertyDetails{}
	return &this
}

// NewPropertyDetailsWithDefaults instantiates a new PropertyDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyDetailsWithDefaults() *PropertyDetails {
	this := PropertyDetails{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PropertyDetails) GetPrice() int32 {
	if o == nil || o.Price.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PropertyDetails) GetPriceOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *PropertyDetails) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableInt32 and assigns it to the Price field.
func (o *PropertyDetails) SetPrice(v int32) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *PropertyDetails) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *PropertyDetails) UnsetPrice() {
	o.Price.Unset()
}

// GetMlsNumber returns the MlsNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PropertyDetails) GetMlsNumber() string {
	if o == nil || o.MlsNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.MlsNumber.Get()
}

// GetMlsNumberOk returns a tuple with the MlsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PropertyDetails) GetMlsNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MlsNumber.Get(), o.MlsNumber.IsSet()
}

// HasMlsNumber returns a boolean if a field has been set.
func (o *PropertyDetails) HasMlsNumber() bool {
	if o != nil && o.MlsNumber.IsSet() {
		return true
	}

	return false
}

// SetMlsNumber gets a reference to the given NullableString and assigns it to the MlsNumber field.
func (o *PropertyDetails) SetMlsNumber(v string) {
	o.MlsNumber.Set(&v)
}
// SetMlsNumberNil sets the value for MlsNumber to be an explicit nil
func (o *PropertyDetails) SetMlsNumberNil() {
	o.MlsNumber.Set(nil)
}

// UnsetMlsNumber ensures that no value is present for MlsNumber, not even an explicit nil
func (o *PropertyDetails) UnsetMlsNumber() {
	o.MlsNumber.Unset()
}

// GetBedrooms returns the Bedrooms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PropertyDetails) GetBedrooms() int32 {
	if o == nil || o.Bedrooms.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Bedrooms.Get()
}

// GetBedroomsOk returns a tuple with the Bedrooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PropertyDetails) GetBedroomsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bedrooms.Get(), o.Bedrooms.IsSet()
}

// HasBedrooms returns a boolean if a field has been set.
func (o *PropertyDetails) HasBedrooms() bool {
	if o != nil && o.Bedrooms.IsSet() {
		return true
	}

	return false
}

// SetBedrooms gets a reference to the given NullableInt32 and assigns it to the Bedrooms field.
func (o *PropertyDetails) SetBedrooms(v int32) {
	o.Bedrooms.Set(&v)
}
// SetBedroomsNil sets the value for Bedrooms to be an explicit nil
func (o *PropertyDetails) SetBedroomsNil() {
	o.Bedrooms.Set(nil)
}

// UnsetBedrooms ensures that no value is present for Bedrooms, not even an explicit nil
func (o *PropertyDetails) UnsetBedrooms() {
	o.Bedrooms.Unset()
}

// GetBathrooms returns the Bathrooms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PropertyDetails) GetBathrooms() float32 {
	if o == nil || o.Bathrooms.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Bathrooms.Get()
}

// GetBathroomsOk returns a tuple with the Bathrooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PropertyDetails) GetBathroomsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bathrooms.Get(), o.Bathrooms.IsSet()
}

// HasBathrooms returns a boolean if a field has been set.
func (o *PropertyDetails) HasBathrooms() bool {
	if o != nil && o.Bathrooms.IsSet() {
		return true
	}

	return false
}

// SetBathrooms gets a reference to the given NullableFloat32 and assigns it to the Bathrooms field.
func (o *PropertyDetails) SetBathrooms(v float32) {
	o.Bathrooms.Set(&v)
}
// SetBathroomsNil sets the value for Bathrooms to be an explicit nil
func (o *PropertyDetails) SetBathroomsNil() {
	o.Bathrooms.Set(nil)
}

// UnsetBathrooms ensures that no value is present for Bathrooms, not even an explicit nil
func (o *PropertyDetails) UnsetBathrooms() {
	o.Bathrooms.Unset()
}

// GetSquareFeet returns the SquareFeet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PropertyDetails) GetSquareFeet() float32 {
	if o == nil || o.SquareFeet.Get() == nil {
		var ret float32
		return ret
	}
	return *o.SquareFeet.Get()
}

// GetSquareFeetOk returns a tuple with the SquareFeet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PropertyDetails) GetSquareFeetOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SquareFeet.Get(), o.SquareFeet.IsSet()
}

// HasSquareFeet returns a boolean if a field has been set.
func (o *PropertyDetails) HasSquareFeet() bool {
	if o != nil && o.SquareFeet.IsSet() {
		return true
	}

	return false
}

// SetSquareFeet gets a reference to the given NullableFloat32 and assigns it to the SquareFeet field.
func (o *PropertyDetails) SetSquareFeet(v float32) {
	o.SquareFeet.Set(&v)
}
// SetSquareFeetNil sets the value for SquareFeet to be an explicit nil
func (o *PropertyDetails) SetSquareFeetNil() {
	o.SquareFeet.Set(nil)
}

// UnsetSquareFeet ensures that no value is present for SquareFeet, not even an explicit nil
func (o *PropertyDetails) UnsetSquareFeet() {
	o.SquareFeet.Unset()
}

// GetLotAcres returns the LotAcres field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PropertyDetails) GetLotAcres() float32 {
	if o == nil || o.LotAcres.Get() == nil {
		var ret float32
		return ret
	}
	return *o.LotAcres.Get()
}

// GetLotAcresOk returns a tuple with the LotAcres field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PropertyDetails) GetLotAcresOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LotAcres.Get(), o.LotAcres.IsSet()
}

// HasLotAcres returns a boolean if a field has been set.
func (o *PropertyDetails) HasLotAcres() bool {
	if o != nil && o.LotAcres.IsSet() {
		return true
	}

	return false
}

// SetLotAcres gets a reference to the given NullableFloat32 and assigns it to the LotAcres field.
func (o *PropertyDetails) SetLotAcres(v float32) {
	o.LotAcres.Set(&v)
}
// SetLotAcresNil sets the value for LotAcres to be an explicit nil
func (o *PropertyDetails) SetLotAcresNil() {
	o.LotAcres.Set(nil)
}

// UnsetLotAcres ensures that no value is present for LotAcres, not even an explicit nil
func (o *PropertyDetails) UnsetLotAcres() {
	o.LotAcres.Unset()
}

// GetParkingSpots returns the ParkingSpots field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PropertyDetails) GetParkingSpots() int32 {
	if o == nil || o.ParkingSpots.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ParkingSpots.Get()
}

// GetParkingSpotsOk returns a tuple with the ParkingSpots field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PropertyDetails) GetParkingSpotsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParkingSpots.Get(), o.ParkingSpots.IsSet()
}

// HasParkingSpots returns a boolean if a field has been set.
func (o *PropertyDetails) HasParkingSpots() bool {
	if o != nil && o.ParkingSpots.IsSet() {
		return true
	}

	return false
}

// SetParkingSpots gets a reference to the given NullableInt32 and assigns it to the ParkingSpots field.
func (o *PropertyDetails) SetParkingSpots(v int32) {
	o.ParkingSpots.Set(&v)
}
// SetParkingSpotsNil sets the value for ParkingSpots to be an explicit nil
func (o *PropertyDetails) SetParkingSpotsNil() {
	o.ParkingSpots.Set(nil)
}

// UnsetParkingSpots ensures that no value is present for ParkingSpots, not even an explicit nil
func (o *PropertyDetails) UnsetParkingSpots() {
	o.ParkingSpots.Unset()
}

// GetYearBuilt returns the YearBuilt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PropertyDetails) GetYearBuilt() int32 {
	if o == nil || o.YearBuilt.Get() == nil {
		var ret int32
		return ret
	}
	return *o.YearBuilt.Get()
}

// GetYearBuiltOk returns a tuple with the YearBuilt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PropertyDetails) GetYearBuiltOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.YearBuilt.Get(), o.YearBuilt.IsSet()
}

// HasYearBuilt returns a boolean if a field has been set.
func (o *PropertyDetails) HasYearBuilt() bool {
	if o != nil && o.YearBuilt.IsSet() {
		return true
	}

	return false
}

// SetYearBuilt gets a reference to the given NullableInt32 and assigns it to the YearBuilt field.
func (o *PropertyDetails) SetYearBuilt(v int32) {
	o.YearBuilt.Set(&v)
}
// SetYearBuiltNil sets the value for YearBuilt to be an explicit nil
func (o *PropertyDetails) SetYearBuiltNil() {
	o.YearBuilt.Set(nil)
}

// UnsetYearBuilt ensures that no value is present for YearBuilt, not even an explicit nil
func (o *PropertyDetails) UnsetYearBuilt() {
	o.YearBuilt.Unset()
}

// GetPropertyType returns the PropertyType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PropertyDetails) GetPropertyType() string {
	if o == nil || o.PropertyType.Get() == nil {
		var ret string
		return ret
	}
	return *o.PropertyType.Get()
}

// GetPropertyTypeOk returns a tuple with the PropertyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PropertyDetails) GetPropertyTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PropertyType.Get(), o.PropertyType.IsSet()
}

// HasPropertyType returns a boolean if a field has been set.
func (o *PropertyDetails) HasPropertyType() bool {
	if o != nil && o.PropertyType.IsSet() {
		return true
	}

	return false
}

// SetPropertyType gets a reference to the given NullableString and assigns it to the PropertyType field.
func (o *PropertyDetails) SetPropertyType(v string) {
	o.PropertyType.Set(&v)
}
// SetPropertyTypeNil sets the value for PropertyType to be an explicit nil
func (o *PropertyDetails) SetPropertyTypeNil() {
	o.PropertyType.Set(nil)
}

// UnsetPropertyType ensures that no value is present for PropertyType, not even an explicit nil
func (o *PropertyDetails) UnsetPropertyType() {
	o.PropertyType.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PropertyDetails) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PropertyDetails) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PropertyDetails) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PropertyDetails) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PropertyDetails) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PropertyDetails) UnsetDescription() {
	o.Description.Unset()
}

func (o PropertyDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.MlsNumber.IsSet() {
		toSerialize["mls_number"] = o.MlsNumber.Get()
	}
	if o.Bedrooms.IsSet() {
		toSerialize["bedrooms"] = o.Bedrooms.Get()
	}
	if o.Bathrooms.IsSet() {
		toSerialize["bathrooms"] = o.Bathrooms.Get()
	}
	if o.SquareFeet.IsSet() {
		toSerialize["square_feet"] = o.SquareFeet.Get()
	}
	if o.LotAcres.IsSet() {
		toSerialize["lot_acres"] = o.LotAcres.Get()
	}
	if o.ParkingSpots.IsSet() {
		toSerialize["parking_spots"] = o.ParkingSpots.Get()
	}
	if o.YearBuilt.IsSet() {
		toSerialize["year_built"] = o.YearBuilt.Get()
	}
	if o.PropertyType.IsSet() {
		toSerialize["property_type"] = o.PropertyType.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyDetails struct {
	value *PropertyDetails
	isSet bool
}

func (v NullablePropertyDetails) Get() *PropertyDetails {
	return v.value
}

func (v *NullablePropertyDetails) Set(val *PropertyDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyDetails(val *PropertyDetails) *NullablePropertyDetails {
	return &NullablePropertyDetails{value: val, isSet: true}
}

func (v NullablePropertyDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


