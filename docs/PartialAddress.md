# PartialAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of address. | 
**FullAddress** | Pointer to **NullableString** | The full address string containing address_1 and address_2. | [optional] 
**FormattedAddress1** | **string** | A formatted address string containing the street. | 
**FormattedAddress2** | **string** | A formatted address string containing the city, state, and zipcode. | 
**Latitude** | **float64** | Latitude of the address. | 
**Longitude** | **float64** | Longitude of the address. | 
**PlaceId** | Pointer to **NullableString** | ID of place. | [optional] 
**AddressLine1** | Pointer to **NullableString** | Address line 1 | [optional] 
**AddressLine2** | Pointer to **NullableString** | Address line 2 | [optional] 
**City** | Pointer to **NullableString** | City | [optional] 
**State** | Pointer to **NullableString** | State | [optional] 
**PostalCode** | Pointer to **NullableString** | Postal Code | [optional] 

## Methods

### NewPartialAddress

`func NewPartialAddress(id int32, formattedAddress1 string, formattedAddress2 string, latitude float64, longitude float64, ) *PartialAddress`

NewPartialAddress instantiates a new PartialAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialAddressWithDefaults

`func NewPartialAddressWithDefaults() *PartialAddress`

NewPartialAddressWithDefaults instantiates a new PartialAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PartialAddress) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartialAddress) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartialAddress) SetId(v int32)`

SetId sets Id field to given value.


### GetFullAddress

`func (o *PartialAddress) GetFullAddress() string`

GetFullAddress returns the FullAddress field if non-nil, zero value otherwise.

### GetFullAddressOk

`func (o *PartialAddress) GetFullAddressOk() (*string, bool)`

GetFullAddressOk returns a tuple with the FullAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullAddress

`func (o *PartialAddress) SetFullAddress(v string)`

SetFullAddress sets FullAddress field to given value.

### HasFullAddress

`func (o *PartialAddress) HasFullAddress() bool`

HasFullAddress returns a boolean if a field has been set.

### SetFullAddressNil

`func (o *PartialAddress) SetFullAddressNil(b bool)`

 SetFullAddressNil sets the value for FullAddress to be an explicit nil

### UnsetFullAddress
`func (o *PartialAddress) UnsetFullAddress()`

UnsetFullAddress ensures that no value is present for FullAddress, not even an explicit nil
### GetFormattedAddress1

`func (o *PartialAddress) GetFormattedAddress1() string`

GetFormattedAddress1 returns the FormattedAddress1 field if non-nil, zero value otherwise.

### GetFormattedAddress1Ok

`func (o *PartialAddress) GetFormattedAddress1Ok() (*string, bool)`

GetFormattedAddress1Ok returns a tuple with the FormattedAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAddress1

`func (o *PartialAddress) SetFormattedAddress1(v string)`

SetFormattedAddress1 sets FormattedAddress1 field to given value.


### GetFormattedAddress2

`func (o *PartialAddress) GetFormattedAddress2() string`

GetFormattedAddress2 returns the FormattedAddress2 field if non-nil, zero value otherwise.

### GetFormattedAddress2Ok

`func (o *PartialAddress) GetFormattedAddress2Ok() (*string, bool)`

GetFormattedAddress2Ok returns a tuple with the FormattedAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAddress2

`func (o *PartialAddress) SetFormattedAddress2(v string)`

SetFormattedAddress2 sets FormattedAddress2 field to given value.


### GetLatitude

`func (o *PartialAddress) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *PartialAddress) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *PartialAddress) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *PartialAddress) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *PartialAddress) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *PartialAddress) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetPlaceId

`func (o *PartialAddress) GetPlaceId() string`

GetPlaceId returns the PlaceId field if non-nil, zero value otherwise.

### GetPlaceIdOk

`func (o *PartialAddress) GetPlaceIdOk() (*string, bool)`

GetPlaceIdOk returns a tuple with the PlaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceId

`func (o *PartialAddress) SetPlaceId(v string)`

SetPlaceId sets PlaceId field to given value.

### HasPlaceId

`func (o *PartialAddress) HasPlaceId() bool`

HasPlaceId returns a boolean if a field has been set.

### SetPlaceIdNil

`func (o *PartialAddress) SetPlaceIdNil(b bool)`

 SetPlaceIdNil sets the value for PlaceId to be an explicit nil

### UnsetPlaceId
`func (o *PartialAddress) UnsetPlaceId()`

UnsetPlaceId ensures that no value is present for PlaceId, not even an explicit nil
### GetAddressLine1

`func (o *PartialAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *PartialAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *PartialAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *PartialAddress) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### SetAddressLine1Nil

`func (o *PartialAddress) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *PartialAddress) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *PartialAddress) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *PartialAddress) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *PartialAddress) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *PartialAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *PartialAddress) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *PartialAddress) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetCity

`func (o *PartialAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PartialAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PartialAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *PartialAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *PartialAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *PartialAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *PartialAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PartialAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PartialAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PartialAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *PartialAddress) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *PartialAddress) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetPostalCode

`func (o *PartialAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PartialAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PartialAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *PartialAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *PartialAddress) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *PartialAddress) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


