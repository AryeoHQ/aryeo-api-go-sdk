# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the address. | 
**Latitude** | **float32** | The geographic latitude of some reference point of the location, specified in degrees and decimal parts. Positive numbers must not include the plus symbol. | 
**Longitude** | **float32** | The geographic longitude of some reference point of the location, specified in degrees and decimal parts. Positive numbers must not include the plus symbol. | 
**StreetNumber** | Pointer to **NullableString** | The street number portion of a location&#39;s address. In some areas, the street number may contain non-numeric characters. This field can also contain extensions and modifiers to the street number, such as \&quot;1/2\&quot; or \&quot;-B\&quot;. | [optional] 
**StreetName** | Pointer to **NullableString** | The street name portion of a location&#39;s address. | [optional] 
**UnitNumber** | Pointer to **NullableString** | The number or portion of a larger building or complex. Examples are: \&quot;APT G\&quot;, \&quot;55\&quot;, etc. | [optional] 
**PostalCode** | Pointer to **NullableString** | The postal code portion of a location&#39;s address. | [optional] 
**City** | Pointer to **NullableString** | The city of a location&#39;s address. | [optional] 
**CityRegion** | Pointer to **NullableString** | A sub-section or area of a defined city. Examples would be SoHo in New York, NY, Ironbound in Newark, NJ or Inside the Beltway. | [optional] 
**CountyOrParish** | Pointer to **NullableString** | The County, Parish or other regional authority of the location. | [optional] 
**StateOrProvince** | Pointer to **NullableString** | The ISO 3166-2 subdivision code for the state or province of the location. For example, “MA” for Massachusetts, United States. | [optional] 
**StateOrProvinceRegion** | Pointer to **NullableString** | A sub-section or area of a defined state or province. Examples would be the Keys in FL or Hudson Valley in NY. | [optional] 
**Country** | Pointer to **NullableString** | The ISO 3166-1 country code for this for the country of the location. | [optional] 
**CountryRegion** | Pointer to **NullableString** | A sub-section or area of a defined country. Examples would be Napa Valley in the US, or the Amalfi Coast in Italy. | [optional] 
**UnparsedAddress** | Pointer to **NullableString** | Unparsed text representation of the address.  | [optional] 
**UnparsedAddressPartOne** | Pointer to **NullableString** | Unparsed text representation of the first part of the address, typically including the street number, street name, and unit number if applicable.   | [optional] 
**UnparsedAddressPartTwo** | Pointer to **NullableString** | Unparsed text representation of the second part of the address, typically including the city, state or province, and postal code.   | [optional] 

## Methods

### NewAddress

`func NewAddress(id string, latitude float32, longitude float32, ) *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Address) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Address) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Address) SetId(v string)`

SetId sets Id field to given value.


### GetLatitude

`func (o *Address) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *Address) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *Address) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *Address) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *Address) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *Address) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### GetStreetNumber

`func (o *Address) GetStreetNumber() string`

GetStreetNumber returns the StreetNumber field if non-nil, zero value otherwise.

### GetStreetNumberOk

`func (o *Address) GetStreetNumberOk() (*string, bool)`

GetStreetNumberOk returns a tuple with the StreetNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetNumber

`func (o *Address) SetStreetNumber(v string)`

SetStreetNumber sets StreetNumber field to given value.

### HasStreetNumber

`func (o *Address) HasStreetNumber() bool`

HasStreetNumber returns a boolean if a field has been set.

### SetStreetNumberNil

`func (o *Address) SetStreetNumberNil(b bool)`

 SetStreetNumberNil sets the value for StreetNumber to be an explicit nil

### UnsetStreetNumber
`func (o *Address) UnsetStreetNumber()`

UnsetStreetNumber ensures that no value is present for StreetNumber, not even an explicit nil
### GetStreetName

`func (o *Address) GetStreetName() string`

GetStreetName returns the StreetName field if non-nil, zero value otherwise.

### GetStreetNameOk

`func (o *Address) GetStreetNameOk() (*string, bool)`

GetStreetNameOk returns a tuple with the StreetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName

`func (o *Address) SetStreetName(v string)`

SetStreetName sets StreetName field to given value.

### HasStreetName

`func (o *Address) HasStreetName() bool`

HasStreetName returns a boolean if a field has been set.

### SetStreetNameNil

`func (o *Address) SetStreetNameNil(b bool)`

 SetStreetNameNil sets the value for StreetName to be an explicit nil

### UnsetStreetName
`func (o *Address) UnsetStreetName()`

UnsetStreetName ensures that no value is present for StreetName, not even an explicit nil
### GetUnitNumber

`func (o *Address) GetUnitNumber() string`

GetUnitNumber returns the UnitNumber field if non-nil, zero value otherwise.

### GetUnitNumberOk

`func (o *Address) GetUnitNumberOk() (*string, bool)`

GetUnitNumberOk returns a tuple with the UnitNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitNumber

`func (o *Address) SetUnitNumber(v string)`

SetUnitNumber sets UnitNumber field to given value.

### HasUnitNumber

`func (o *Address) HasUnitNumber() bool`

HasUnitNumber returns a boolean if a field has been set.

### SetUnitNumberNil

`func (o *Address) SetUnitNumberNil(b bool)`

 SetUnitNumberNil sets the value for UnitNumber to be an explicit nil

### UnsetUnitNumber
`func (o *Address) UnsetUnitNumber()`

UnsetUnitNumber ensures that no value is present for UnitNumber, not even an explicit nil
### GetPostalCode

`func (o *Address) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Address) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Address) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *Address) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *Address) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *Address) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCity

`func (o *Address) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Address) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Address) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Address) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *Address) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *Address) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCityRegion

`func (o *Address) GetCityRegion() string`

GetCityRegion returns the CityRegion field if non-nil, zero value otherwise.

### GetCityRegionOk

`func (o *Address) GetCityRegionOk() (*string, bool)`

GetCityRegionOk returns a tuple with the CityRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityRegion

`func (o *Address) SetCityRegion(v string)`

SetCityRegion sets CityRegion field to given value.

### HasCityRegion

`func (o *Address) HasCityRegion() bool`

HasCityRegion returns a boolean if a field has been set.

### SetCityRegionNil

`func (o *Address) SetCityRegionNil(b bool)`

 SetCityRegionNil sets the value for CityRegion to be an explicit nil

### UnsetCityRegion
`func (o *Address) UnsetCityRegion()`

UnsetCityRegion ensures that no value is present for CityRegion, not even an explicit nil
### GetCountyOrParish

`func (o *Address) GetCountyOrParish() string`

GetCountyOrParish returns the CountyOrParish field if non-nil, zero value otherwise.

### GetCountyOrParishOk

`func (o *Address) GetCountyOrParishOk() (*string, bool)`

GetCountyOrParishOk returns a tuple with the CountyOrParish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountyOrParish

`func (o *Address) SetCountyOrParish(v string)`

SetCountyOrParish sets CountyOrParish field to given value.

### HasCountyOrParish

`func (o *Address) HasCountyOrParish() bool`

HasCountyOrParish returns a boolean if a field has been set.

### SetCountyOrParishNil

`func (o *Address) SetCountyOrParishNil(b bool)`

 SetCountyOrParishNil sets the value for CountyOrParish to be an explicit nil

### UnsetCountyOrParish
`func (o *Address) UnsetCountyOrParish()`

UnsetCountyOrParish ensures that no value is present for CountyOrParish, not even an explicit nil
### GetStateOrProvince

`func (o *Address) GetStateOrProvince() string`

GetStateOrProvince returns the StateOrProvince field if non-nil, zero value otherwise.

### GetStateOrProvinceOk

`func (o *Address) GetStateOrProvinceOk() (*string, bool)`

GetStateOrProvinceOk returns a tuple with the StateOrProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvince

`func (o *Address) SetStateOrProvince(v string)`

SetStateOrProvince sets StateOrProvince field to given value.

### HasStateOrProvince

`func (o *Address) HasStateOrProvince() bool`

HasStateOrProvince returns a boolean if a field has been set.

### SetStateOrProvinceNil

`func (o *Address) SetStateOrProvinceNil(b bool)`

 SetStateOrProvinceNil sets the value for StateOrProvince to be an explicit nil

### UnsetStateOrProvince
`func (o *Address) UnsetStateOrProvince()`

UnsetStateOrProvince ensures that no value is present for StateOrProvince, not even an explicit nil
### GetStateOrProvinceRegion

`func (o *Address) GetStateOrProvinceRegion() string`

GetStateOrProvinceRegion returns the StateOrProvinceRegion field if non-nil, zero value otherwise.

### GetStateOrProvinceRegionOk

`func (o *Address) GetStateOrProvinceRegionOk() (*string, bool)`

GetStateOrProvinceRegionOk returns a tuple with the StateOrProvinceRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvinceRegion

`func (o *Address) SetStateOrProvinceRegion(v string)`

SetStateOrProvinceRegion sets StateOrProvinceRegion field to given value.

### HasStateOrProvinceRegion

`func (o *Address) HasStateOrProvinceRegion() bool`

HasStateOrProvinceRegion returns a boolean if a field has been set.

### SetStateOrProvinceRegionNil

`func (o *Address) SetStateOrProvinceRegionNil(b bool)`

 SetStateOrProvinceRegionNil sets the value for StateOrProvinceRegion to be an explicit nil

### UnsetStateOrProvinceRegion
`func (o *Address) UnsetStateOrProvinceRegion()`

UnsetStateOrProvinceRegion ensures that no value is present for StateOrProvinceRegion, not even an explicit nil
### GetCountry

`func (o *Address) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Address) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Address) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Address) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *Address) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Address) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCountryRegion

`func (o *Address) GetCountryRegion() string`

GetCountryRegion returns the CountryRegion field if non-nil, zero value otherwise.

### GetCountryRegionOk

`func (o *Address) GetCountryRegionOk() (*string, bool)`

GetCountryRegionOk returns a tuple with the CountryRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryRegion

`func (o *Address) SetCountryRegion(v string)`

SetCountryRegion sets CountryRegion field to given value.

### HasCountryRegion

`func (o *Address) HasCountryRegion() bool`

HasCountryRegion returns a boolean if a field has been set.

### SetCountryRegionNil

`func (o *Address) SetCountryRegionNil(b bool)`

 SetCountryRegionNil sets the value for CountryRegion to be an explicit nil

### UnsetCountryRegion
`func (o *Address) UnsetCountryRegion()`

UnsetCountryRegion ensures that no value is present for CountryRegion, not even an explicit nil
### GetUnparsedAddress

`func (o *Address) GetUnparsedAddress() string`

GetUnparsedAddress returns the UnparsedAddress field if non-nil, zero value otherwise.

### GetUnparsedAddressOk

`func (o *Address) GetUnparsedAddressOk() (*string, bool)`

GetUnparsedAddressOk returns a tuple with the UnparsedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnparsedAddress

`func (o *Address) SetUnparsedAddress(v string)`

SetUnparsedAddress sets UnparsedAddress field to given value.

### HasUnparsedAddress

`func (o *Address) HasUnparsedAddress() bool`

HasUnparsedAddress returns a boolean if a field has been set.

### SetUnparsedAddressNil

`func (o *Address) SetUnparsedAddressNil(b bool)`

 SetUnparsedAddressNil sets the value for UnparsedAddress to be an explicit nil

### UnsetUnparsedAddress
`func (o *Address) UnsetUnparsedAddress()`

UnsetUnparsedAddress ensures that no value is present for UnparsedAddress, not even an explicit nil
### GetUnparsedAddressPartOne

`func (o *Address) GetUnparsedAddressPartOne() string`

GetUnparsedAddressPartOne returns the UnparsedAddressPartOne field if non-nil, zero value otherwise.

### GetUnparsedAddressPartOneOk

`func (o *Address) GetUnparsedAddressPartOneOk() (*string, bool)`

GetUnparsedAddressPartOneOk returns a tuple with the UnparsedAddressPartOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnparsedAddressPartOne

`func (o *Address) SetUnparsedAddressPartOne(v string)`

SetUnparsedAddressPartOne sets UnparsedAddressPartOne field to given value.

### HasUnparsedAddressPartOne

`func (o *Address) HasUnparsedAddressPartOne() bool`

HasUnparsedAddressPartOne returns a boolean if a field has been set.

### SetUnparsedAddressPartOneNil

`func (o *Address) SetUnparsedAddressPartOneNil(b bool)`

 SetUnparsedAddressPartOneNil sets the value for UnparsedAddressPartOne to be an explicit nil

### UnsetUnparsedAddressPartOne
`func (o *Address) UnsetUnparsedAddressPartOne()`

UnsetUnparsedAddressPartOne ensures that no value is present for UnparsedAddressPartOne, not even an explicit nil
### GetUnparsedAddressPartTwo

`func (o *Address) GetUnparsedAddressPartTwo() string`

GetUnparsedAddressPartTwo returns the UnparsedAddressPartTwo field if non-nil, zero value otherwise.

### GetUnparsedAddressPartTwoOk

`func (o *Address) GetUnparsedAddressPartTwoOk() (*string, bool)`

GetUnparsedAddressPartTwoOk returns a tuple with the UnparsedAddressPartTwo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnparsedAddressPartTwo

`func (o *Address) SetUnparsedAddressPartTwo(v string)`

SetUnparsedAddressPartTwo sets UnparsedAddressPartTwo field to given value.

### HasUnparsedAddressPartTwo

`func (o *Address) HasUnparsedAddressPartTwo() bool`

HasUnparsedAddressPartTwo returns a boolean if a field has been set.

### SetUnparsedAddressPartTwoNil

`func (o *Address) SetUnparsedAddressPartTwoNil(b bool)`

 SetUnparsedAddressPartTwoNil sets the value for UnparsedAddressPartTwo to be an explicit nil

### UnsetUnparsedAddressPartTwo
`func (o *Address) UnsetUnparsedAddressPartTwo()`

UnsetUnparsedAddressPartTwo ensures that no value is present for UnparsedAddressPartTwo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


