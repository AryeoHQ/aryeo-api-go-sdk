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

// Address A street address and additional metadata about a location.
type Address struct {
	// ID of the address. UUID Version 4.
	Id string `json:"id"`
	// The geographic latitude of some reference point of the location, specified in degrees and decimal parts. Positive numbers must not include the plus symbol.
	Latitude float32 `json:"latitude"`
	// The geographic longitude of some reference point of the location, specified in degrees and decimal parts. Positive numbers must not include the plus symbol.
	Longitude float32 `json:"longitude"`
	// The street number portion of a location's address. In some areas, the street number may contain non-numeric characters. This field can also contain extensions and modifiers to the street number, such as \"1/2\" or \"-B\".
	StreetNumber NullableString `json:"street_number,omitempty"`
	// The street name portion of a location's address.
	StreetName NullableString `json:"street_name,omitempty"`
	// The number or portion of a larger building or complex. Examples are: \"APT G\", \"55\", etc.
	UnitNumber NullableString `json:"unit_number,omitempty"`
	// The postal code portion of a location's address.
	PostalCode NullableString `json:"postal_code,omitempty"`
	// The city of a location's address.
	City NullableString `json:"city,omitempty"`
	// A sub-section or area of a defined city. Examples would be SoHo in New York, NY, Ironbound in Newark, NJ or Inside the Beltway.
	CityRegion NullableString `json:"city_region,omitempty"`
	// The County, Parish or other regional authority of the location.
	CountyOrParish NullableString `json:"county_or_parish,omitempty"`
	// The ISO 3166-2 subdivision code for the state or province of the location. For example, “MA” for Massachusetts, United States.
	StateOrProvince NullableString `json:"state_or_province,omitempty"`
	// A sub-section or area of a defined state or province. Examples would be the Keys in FL or Hudson Valley in NY.
	StateOrProvinceRegion NullableString `json:"state_or_province_region,omitempty"`
	// The ISO 3166-1 country code for this for the country of the location.
	Country NullableString `json:"country,omitempty"`
	// A sub-section or area of a defined country. Examples would be Napa Valley in the US, or the Amalfi Coast in Italy.
	CountryRegion NullableString `json:"country_region,omitempty"`
	// Unparsed text representation of the address. 
	UnparsedAddress NullableString `json:"unparsed_address,omitempty"`
	// Unparsed text representation of the first part of the address, typically including the street number, street name, and unit number if applicable.  
	UnparsedAddressPartOne NullableString `json:"unparsed_address_part_one,omitempty"`
	// Unparsed text representation of the second part of the address, typically including the city, state or province, and postal code.  
	UnparsedAddressPartTwo NullableString `json:"unparsed_address_part_two,omitempty"`
}

// NewAddress instantiates a new Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress(id string, latitude float32, longitude float32) *Address {
	this := Address{}
	this.Id = id
	this.Latitude = latitude
	this.Longitude = longitude
	return &this
}

// NewAddressWithDefaults instantiates a new Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithDefaults() *Address {
	this := Address{}
	return &this
}

// GetId returns the Id field value
func (o *Address) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Address) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Address) SetId(v string) {
	o.Id = v
}

// GetLatitude returns the Latitude field value
func (o *Address) GetLatitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *Address) GetLatitudeOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *Address) SetLatitude(v float32) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *Address) GetLongitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *Address) GetLongitudeOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *Address) SetLongitude(v float32) {
	o.Longitude = v
}

// GetStreetNumber returns the StreetNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetStreetNumber() string {
	if o == nil || o.StreetNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.StreetNumber.Get()
}

// GetStreetNumberOk returns a tuple with the StreetNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetStreetNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StreetNumber.Get(), o.StreetNumber.IsSet()
}

// HasStreetNumber returns a boolean if a field has been set.
func (o *Address) HasStreetNumber() bool {
	if o != nil && o.StreetNumber.IsSet() {
		return true
	}

	return false
}

// SetStreetNumber gets a reference to the given NullableString and assigns it to the StreetNumber field.
func (o *Address) SetStreetNumber(v string) {
	o.StreetNumber.Set(&v)
}
// SetStreetNumberNil sets the value for StreetNumber to be an explicit nil
func (o *Address) SetStreetNumberNil() {
	o.StreetNumber.Set(nil)
}

// UnsetStreetNumber ensures that no value is present for StreetNumber, not even an explicit nil
func (o *Address) UnsetStreetNumber() {
	o.StreetNumber.Unset()
}

// GetStreetName returns the StreetName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetStreetName() string {
	if o == nil || o.StreetName.Get() == nil {
		var ret string
		return ret
	}
	return *o.StreetName.Get()
}

// GetStreetNameOk returns a tuple with the StreetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetStreetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StreetName.Get(), o.StreetName.IsSet()
}

// HasStreetName returns a boolean if a field has been set.
func (o *Address) HasStreetName() bool {
	if o != nil && o.StreetName.IsSet() {
		return true
	}

	return false
}

// SetStreetName gets a reference to the given NullableString and assigns it to the StreetName field.
func (o *Address) SetStreetName(v string) {
	o.StreetName.Set(&v)
}
// SetStreetNameNil sets the value for StreetName to be an explicit nil
func (o *Address) SetStreetNameNil() {
	o.StreetName.Set(nil)
}

// UnsetStreetName ensures that no value is present for StreetName, not even an explicit nil
func (o *Address) UnsetStreetName() {
	o.StreetName.Unset()
}

// GetUnitNumber returns the UnitNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetUnitNumber() string {
	if o == nil || o.UnitNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.UnitNumber.Get()
}

// GetUnitNumberOk returns a tuple with the UnitNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetUnitNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnitNumber.Get(), o.UnitNumber.IsSet()
}

// HasUnitNumber returns a boolean if a field has been set.
func (o *Address) HasUnitNumber() bool {
	if o != nil && o.UnitNumber.IsSet() {
		return true
	}

	return false
}

// SetUnitNumber gets a reference to the given NullableString and assigns it to the UnitNumber field.
func (o *Address) SetUnitNumber(v string) {
	o.UnitNumber.Set(&v)
}
// SetUnitNumberNil sets the value for UnitNumber to be an explicit nil
func (o *Address) SetUnitNumberNil() {
	o.UnitNumber.Set(nil)
}

// UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil
func (o *Address) UnsetUnitNumber() {
	o.UnitNumber.Unset()
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// HasPostalCode returns a boolean if a field has been set.
func (o *Address) HasPostalCode() bool {
	if o != nil && o.PostalCode.IsSet() {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given NullableString and assigns it to the PostalCode field.
func (o *Address) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}
// SetPostalCodeNil sets the value for PostalCode to be an explicit nil
func (o *Address) SetPostalCodeNil() {
	o.PostalCode.Set(nil)
}

// UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
func (o *Address) UnsetPostalCode() {
	o.PostalCode.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *Address) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *Address) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *Address) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *Address) UnsetCity() {
	o.City.Unset()
}

// GetCityRegion returns the CityRegion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetCityRegion() string {
	if o == nil || o.CityRegion.Get() == nil {
		var ret string
		return ret
	}
	return *o.CityRegion.Get()
}

// GetCityRegionOk returns a tuple with the CityRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetCityRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CityRegion.Get(), o.CityRegion.IsSet()
}

// HasCityRegion returns a boolean if a field has been set.
func (o *Address) HasCityRegion() bool {
	if o != nil && o.CityRegion.IsSet() {
		return true
	}

	return false
}

// SetCityRegion gets a reference to the given NullableString and assigns it to the CityRegion field.
func (o *Address) SetCityRegion(v string) {
	o.CityRegion.Set(&v)
}
// SetCityRegionNil sets the value for CityRegion to be an explicit nil
func (o *Address) SetCityRegionNil() {
	o.CityRegion.Set(nil)
}

// UnsetCityRegion ensures that no value is present for CityRegion, not even an explicit nil
func (o *Address) UnsetCityRegion() {
	o.CityRegion.Unset()
}

// GetCountyOrParish returns the CountyOrParish field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetCountyOrParish() string {
	if o == nil || o.CountyOrParish.Get() == nil {
		var ret string
		return ret
	}
	return *o.CountyOrParish.Get()
}

// GetCountyOrParishOk returns a tuple with the CountyOrParish field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetCountyOrParishOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CountyOrParish.Get(), o.CountyOrParish.IsSet()
}

// HasCountyOrParish returns a boolean if a field has been set.
func (o *Address) HasCountyOrParish() bool {
	if o != nil && o.CountyOrParish.IsSet() {
		return true
	}

	return false
}

// SetCountyOrParish gets a reference to the given NullableString and assigns it to the CountyOrParish field.
func (o *Address) SetCountyOrParish(v string) {
	o.CountyOrParish.Set(&v)
}
// SetCountyOrParishNil sets the value for CountyOrParish to be an explicit nil
func (o *Address) SetCountyOrParishNil() {
	o.CountyOrParish.Set(nil)
}

// UnsetCountyOrParish ensures that no value is present for CountyOrParish, not even an explicit nil
func (o *Address) UnsetCountyOrParish() {
	o.CountyOrParish.Unset()
}

// GetStateOrProvince returns the StateOrProvince field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetStateOrProvince() string {
	if o == nil || o.StateOrProvince.Get() == nil {
		var ret string
		return ret
	}
	return *o.StateOrProvince.Get()
}

// GetStateOrProvinceOk returns a tuple with the StateOrProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetStateOrProvinceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StateOrProvince.Get(), o.StateOrProvince.IsSet()
}

// HasStateOrProvince returns a boolean if a field has been set.
func (o *Address) HasStateOrProvince() bool {
	if o != nil && o.StateOrProvince.IsSet() {
		return true
	}

	return false
}

// SetStateOrProvince gets a reference to the given NullableString and assigns it to the StateOrProvince field.
func (o *Address) SetStateOrProvince(v string) {
	o.StateOrProvince.Set(&v)
}
// SetStateOrProvinceNil sets the value for StateOrProvince to be an explicit nil
func (o *Address) SetStateOrProvinceNil() {
	o.StateOrProvince.Set(nil)
}

// UnsetStateOrProvince ensures that no value is present for StateOrProvince, not even an explicit nil
func (o *Address) UnsetStateOrProvince() {
	o.StateOrProvince.Unset()
}

// GetStateOrProvinceRegion returns the StateOrProvinceRegion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetStateOrProvinceRegion() string {
	if o == nil || o.StateOrProvinceRegion.Get() == nil {
		var ret string
		return ret
	}
	return *o.StateOrProvinceRegion.Get()
}

// GetStateOrProvinceRegionOk returns a tuple with the StateOrProvinceRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetStateOrProvinceRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StateOrProvinceRegion.Get(), o.StateOrProvinceRegion.IsSet()
}

// HasStateOrProvinceRegion returns a boolean if a field has been set.
func (o *Address) HasStateOrProvinceRegion() bool {
	if o != nil && o.StateOrProvinceRegion.IsSet() {
		return true
	}

	return false
}

// SetStateOrProvinceRegion gets a reference to the given NullableString and assigns it to the StateOrProvinceRegion field.
func (o *Address) SetStateOrProvinceRegion(v string) {
	o.StateOrProvinceRegion.Set(&v)
}
// SetStateOrProvinceRegionNil sets the value for StateOrProvinceRegion to be an explicit nil
func (o *Address) SetStateOrProvinceRegionNil() {
	o.StateOrProvinceRegion.Set(nil)
}

// UnsetStateOrProvinceRegion ensures that no value is present for StateOrProvinceRegion, not even an explicit nil
func (o *Address) UnsetStateOrProvinceRegion() {
	o.StateOrProvinceRegion.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *Address) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *Address) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *Address) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *Address) UnsetCountry() {
	o.Country.Unset()
}

// GetCountryRegion returns the CountryRegion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetCountryRegion() string {
	if o == nil || o.CountryRegion.Get() == nil {
		var ret string
		return ret
	}
	return *o.CountryRegion.Get()
}

// GetCountryRegionOk returns a tuple with the CountryRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetCountryRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CountryRegion.Get(), o.CountryRegion.IsSet()
}

// HasCountryRegion returns a boolean if a field has been set.
func (o *Address) HasCountryRegion() bool {
	if o != nil && o.CountryRegion.IsSet() {
		return true
	}

	return false
}

// SetCountryRegion gets a reference to the given NullableString and assigns it to the CountryRegion field.
func (o *Address) SetCountryRegion(v string) {
	o.CountryRegion.Set(&v)
}
// SetCountryRegionNil sets the value for CountryRegion to be an explicit nil
func (o *Address) SetCountryRegionNil() {
	o.CountryRegion.Set(nil)
}

// UnsetCountryRegion ensures that no value is present for CountryRegion, not even an explicit nil
func (o *Address) UnsetCountryRegion() {
	o.CountryRegion.Unset()
}

// GetUnparsedAddress returns the UnparsedAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetUnparsedAddress() string {
	if o == nil || o.UnparsedAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.UnparsedAddress.Get()
}

// GetUnparsedAddressOk returns a tuple with the UnparsedAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetUnparsedAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnparsedAddress.Get(), o.UnparsedAddress.IsSet()
}

// HasUnparsedAddress returns a boolean if a field has been set.
func (o *Address) HasUnparsedAddress() bool {
	if o != nil && o.UnparsedAddress.IsSet() {
		return true
	}

	return false
}

// SetUnparsedAddress gets a reference to the given NullableString and assigns it to the UnparsedAddress field.
func (o *Address) SetUnparsedAddress(v string) {
	o.UnparsedAddress.Set(&v)
}
// SetUnparsedAddressNil sets the value for UnparsedAddress to be an explicit nil
func (o *Address) SetUnparsedAddressNil() {
	o.UnparsedAddress.Set(nil)
}

// UnsetUnparsedAddress ensures that no value is present for UnparsedAddress, not even an explicit nil
func (o *Address) UnsetUnparsedAddress() {
	o.UnparsedAddress.Unset()
}

// GetUnparsedAddressPartOne returns the UnparsedAddressPartOne field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetUnparsedAddressPartOne() string {
	if o == nil || o.UnparsedAddressPartOne.Get() == nil {
		var ret string
		return ret
	}
	return *o.UnparsedAddressPartOne.Get()
}

// GetUnparsedAddressPartOneOk returns a tuple with the UnparsedAddressPartOne field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetUnparsedAddressPartOneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnparsedAddressPartOne.Get(), o.UnparsedAddressPartOne.IsSet()
}

// HasUnparsedAddressPartOne returns a boolean if a field has been set.
func (o *Address) HasUnparsedAddressPartOne() bool {
	if o != nil && o.UnparsedAddressPartOne.IsSet() {
		return true
	}

	return false
}

// SetUnparsedAddressPartOne gets a reference to the given NullableString and assigns it to the UnparsedAddressPartOne field.
func (o *Address) SetUnparsedAddressPartOne(v string) {
	o.UnparsedAddressPartOne.Set(&v)
}
// SetUnparsedAddressPartOneNil sets the value for UnparsedAddressPartOne to be an explicit nil
func (o *Address) SetUnparsedAddressPartOneNil() {
	o.UnparsedAddressPartOne.Set(nil)
}

// UnsetUnparsedAddressPartOne ensures that no value is present for UnparsedAddressPartOne, not even an explicit nil
func (o *Address) UnsetUnparsedAddressPartOne() {
	o.UnparsedAddressPartOne.Unset()
}

// GetUnparsedAddressPartTwo returns the UnparsedAddressPartTwo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetUnparsedAddressPartTwo() string {
	if o == nil || o.UnparsedAddressPartTwo.Get() == nil {
		var ret string
		return ret
	}
	return *o.UnparsedAddressPartTwo.Get()
}

// GetUnparsedAddressPartTwoOk returns a tuple with the UnparsedAddressPartTwo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetUnparsedAddressPartTwoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnparsedAddressPartTwo.Get(), o.UnparsedAddressPartTwo.IsSet()
}

// HasUnparsedAddressPartTwo returns a boolean if a field has been set.
func (o *Address) HasUnparsedAddressPartTwo() bool {
	if o != nil && o.UnparsedAddressPartTwo.IsSet() {
		return true
	}

	return false
}

// SetUnparsedAddressPartTwo gets a reference to the given NullableString and assigns it to the UnparsedAddressPartTwo field.
func (o *Address) SetUnparsedAddressPartTwo(v string) {
	o.UnparsedAddressPartTwo.Set(&v)
}
// SetUnparsedAddressPartTwoNil sets the value for UnparsedAddressPartTwo to be an explicit nil
func (o *Address) SetUnparsedAddressPartTwoNil() {
	o.UnparsedAddressPartTwo.Set(nil)
}

// UnsetUnparsedAddressPartTwo ensures that no value is present for UnparsedAddressPartTwo, not even an explicit nil
func (o *Address) UnsetUnparsedAddressPartTwo() {
	o.UnparsedAddressPartTwo.Unset()
}

func (o Address) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["latitude"] = o.Latitude
	}
	if true {
		toSerialize["longitude"] = o.Longitude
	}
	if o.StreetNumber.IsSet() {
		toSerialize["street_number"] = o.StreetNumber.Get()
	}
	if o.StreetName.IsSet() {
		toSerialize["street_name"] = o.StreetName.Get()
	}
	if o.UnitNumber.IsSet() {
		toSerialize["unit_number"] = o.UnitNumber.Get()
	}
	if o.PostalCode.IsSet() {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.CityRegion.IsSet() {
		toSerialize["city_region"] = o.CityRegion.Get()
	}
	if o.CountyOrParish.IsSet() {
		toSerialize["county_or_parish"] = o.CountyOrParish.Get()
	}
	if o.StateOrProvince.IsSet() {
		toSerialize["state_or_province"] = o.StateOrProvince.Get()
	}
	if o.StateOrProvinceRegion.IsSet() {
		toSerialize["state_or_province_region"] = o.StateOrProvinceRegion.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.CountryRegion.IsSet() {
		toSerialize["country_region"] = o.CountryRegion.Get()
	}
	if o.UnparsedAddress.IsSet() {
		toSerialize["unparsed_address"] = o.UnparsedAddress.Get()
	}
	if o.UnparsedAddressPartOne.IsSet() {
		toSerialize["unparsed_address_part_one"] = o.UnparsedAddressPartOne.Get()
	}
	if o.UnparsedAddressPartTwo.IsSet() {
		toSerialize["unparsed_address_part_two"] = o.UnparsedAddressPartTwo.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAddress struct {
	value *Address
	isSet bool
}

func (v NullableAddress) Get() *Address {
	return v.value
}

func (v *NullableAddress) Set(val *Address) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress(val *Address) *NullableAddress {
	return &NullableAddress{value: val, isSet: true}
}

func (v NullableAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


