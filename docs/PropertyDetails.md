# PropertyDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to **NullableInt32** | The price of the property in dollars. | [optional] 
**MlsNumber** | Pointer to **NullableString** | A number assigned to a real-estate listing for the purposes of information sharing. | [optional] 
**Bedrooms** | Pointer to **NullableInt32** | Number of bedrooms. | [optional] 
**Bathrooms** | Pointer to **NullableFloat32** | Number of bathrooms. Represented as a number in order to support half-baths. | [optional] 
**SquareFeet** | Pointer to **NullableFloat32** | Total number of square feet. | [optional] 
**LotAcres** | Pointer to **NullableFloat32** | Total acres. | [optional] 
**ParkingSpots** | Pointer to **NullableInt32** | Number of parking spaces. | [optional] 
**YearBuilt** | Pointer to **NullableInt32** | Year the property was built. | [optional] 
**PropertyType** | Pointer to **NullableString** | Type of property. | [optional] 
**Description** | Pointer to **NullableString** | Property description. | [optional] 

## Methods

### NewPropertyDetails

`func NewPropertyDetails() *PropertyDetails`

NewPropertyDetails instantiates a new PropertyDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyDetailsWithDefaults

`func NewPropertyDetailsWithDefaults() *PropertyDetails`

NewPropertyDetailsWithDefaults instantiates a new PropertyDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *PropertyDetails) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PropertyDetails) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PropertyDetails) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PropertyDetails) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *PropertyDetails) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *PropertyDetails) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetMlsNumber

`func (o *PropertyDetails) GetMlsNumber() string`

GetMlsNumber returns the MlsNumber field if non-nil, zero value otherwise.

### GetMlsNumberOk

`func (o *PropertyDetails) GetMlsNumberOk() (*string, bool)`

GetMlsNumberOk returns a tuple with the MlsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlsNumber

`func (o *PropertyDetails) SetMlsNumber(v string)`

SetMlsNumber sets MlsNumber field to given value.

### HasMlsNumber

`func (o *PropertyDetails) HasMlsNumber() bool`

HasMlsNumber returns a boolean if a field has been set.

### SetMlsNumberNil

`func (o *PropertyDetails) SetMlsNumberNil(b bool)`

 SetMlsNumberNil sets the value for MlsNumber to be an explicit nil

### UnsetMlsNumber
`func (o *PropertyDetails) UnsetMlsNumber()`

UnsetMlsNumber ensures that no value is present for MlsNumber, not even an explicit nil
### GetBedrooms

`func (o *PropertyDetails) GetBedrooms() int32`

GetBedrooms returns the Bedrooms field if non-nil, zero value otherwise.

### GetBedroomsOk

`func (o *PropertyDetails) GetBedroomsOk() (*int32, bool)`

GetBedroomsOk returns a tuple with the Bedrooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBedrooms

`func (o *PropertyDetails) SetBedrooms(v int32)`

SetBedrooms sets Bedrooms field to given value.

### HasBedrooms

`func (o *PropertyDetails) HasBedrooms() bool`

HasBedrooms returns a boolean if a field has been set.

### SetBedroomsNil

`func (o *PropertyDetails) SetBedroomsNil(b bool)`

 SetBedroomsNil sets the value for Bedrooms to be an explicit nil

### UnsetBedrooms
`func (o *PropertyDetails) UnsetBedrooms()`

UnsetBedrooms ensures that no value is present for Bedrooms, not even an explicit nil
### GetBathrooms

`func (o *PropertyDetails) GetBathrooms() float32`

GetBathrooms returns the Bathrooms field if non-nil, zero value otherwise.

### GetBathroomsOk

`func (o *PropertyDetails) GetBathroomsOk() (*float32, bool)`

GetBathroomsOk returns a tuple with the Bathrooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBathrooms

`func (o *PropertyDetails) SetBathrooms(v float32)`

SetBathrooms sets Bathrooms field to given value.

### HasBathrooms

`func (o *PropertyDetails) HasBathrooms() bool`

HasBathrooms returns a boolean if a field has been set.

### SetBathroomsNil

`func (o *PropertyDetails) SetBathroomsNil(b bool)`

 SetBathroomsNil sets the value for Bathrooms to be an explicit nil

### UnsetBathrooms
`func (o *PropertyDetails) UnsetBathrooms()`

UnsetBathrooms ensures that no value is present for Bathrooms, not even an explicit nil
### GetSquareFeet

`func (o *PropertyDetails) GetSquareFeet() float32`

GetSquareFeet returns the SquareFeet field if non-nil, zero value otherwise.

### GetSquareFeetOk

`func (o *PropertyDetails) GetSquareFeetOk() (*float32, bool)`

GetSquareFeetOk returns a tuple with the SquareFeet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquareFeet

`func (o *PropertyDetails) SetSquareFeet(v float32)`

SetSquareFeet sets SquareFeet field to given value.

### HasSquareFeet

`func (o *PropertyDetails) HasSquareFeet() bool`

HasSquareFeet returns a boolean if a field has been set.

### SetSquareFeetNil

`func (o *PropertyDetails) SetSquareFeetNil(b bool)`

 SetSquareFeetNil sets the value for SquareFeet to be an explicit nil

### UnsetSquareFeet
`func (o *PropertyDetails) UnsetSquareFeet()`

UnsetSquareFeet ensures that no value is present for SquareFeet, not even an explicit nil
### GetLotAcres

`func (o *PropertyDetails) GetLotAcres() float32`

GetLotAcres returns the LotAcres field if non-nil, zero value otherwise.

### GetLotAcresOk

`func (o *PropertyDetails) GetLotAcresOk() (*float32, bool)`

GetLotAcresOk returns a tuple with the LotAcres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotAcres

`func (o *PropertyDetails) SetLotAcres(v float32)`

SetLotAcres sets LotAcres field to given value.

### HasLotAcres

`func (o *PropertyDetails) HasLotAcres() bool`

HasLotAcres returns a boolean if a field has been set.

### SetLotAcresNil

`func (o *PropertyDetails) SetLotAcresNil(b bool)`

 SetLotAcresNil sets the value for LotAcres to be an explicit nil

### UnsetLotAcres
`func (o *PropertyDetails) UnsetLotAcres()`

UnsetLotAcres ensures that no value is present for LotAcres, not even an explicit nil
### GetParkingSpots

`func (o *PropertyDetails) GetParkingSpots() int32`

GetParkingSpots returns the ParkingSpots field if non-nil, zero value otherwise.

### GetParkingSpotsOk

`func (o *PropertyDetails) GetParkingSpotsOk() (*int32, bool)`

GetParkingSpotsOk returns a tuple with the ParkingSpots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParkingSpots

`func (o *PropertyDetails) SetParkingSpots(v int32)`

SetParkingSpots sets ParkingSpots field to given value.

### HasParkingSpots

`func (o *PropertyDetails) HasParkingSpots() bool`

HasParkingSpots returns a boolean if a field has been set.

### SetParkingSpotsNil

`func (o *PropertyDetails) SetParkingSpotsNil(b bool)`

 SetParkingSpotsNil sets the value for ParkingSpots to be an explicit nil

### UnsetParkingSpots
`func (o *PropertyDetails) UnsetParkingSpots()`

UnsetParkingSpots ensures that no value is present for ParkingSpots, not even an explicit nil
### GetYearBuilt

`func (o *PropertyDetails) GetYearBuilt() int32`

GetYearBuilt returns the YearBuilt field if non-nil, zero value otherwise.

### GetYearBuiltOk

`func (o *PropertyDetails) GetYearBuiltOk() (*int32, bool)`

GetYearBuiltOk returns a tuple with the YearBuilt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearBuilt

`func (o *PropertyDetails) SetYearBuilt(v int32)`

SetYearBuilt sets YearBuilt field to given value.

### HasYearBuilt

`func (o *PropertyDetails) HasYearBuilt() bool`

HasYearBuilt returns a boolean if a field has been set.

### SetYearBuiltNil

`func (o *PropertyDetails) SetYearBuiltNil(b bool)`

 SetYearBuiltNil sets the value for YearBuilt to be an explicit nil

### UnsetYearBuilt
`func (o *PropertyDetails) UnsetYearBuilt()`

UnsetYearBuilt ensures that no value is present for YearBuilt, not even an explicit nil
### GetPropertyType

`func (o *PropertyDetails) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *PropertyDetails) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *PropertyDetails) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.

### HasPropertyType

`func (o *PropertyDetails) HasPropertyType() bool`

HasPropertyType returns a boolean if a field has been set.

### SetPropertyTypeNil

`func (o *PropertyDetails) SetPropertyTypeNil(b bool)`

 SetPropertyTypeNil sets the value for PropertyType to be an explicit nil

### UnsetPropertyType
`func (o *PropertyDetails) UnsetPropertyType()`

UnsetPropertyType ensures that no value is present for PropertyType, not even an explicit nil
### GetDescription

`func (o *PropertyDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PropertyDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PropertyDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PropertyDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PropertyDetails) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PropertyDetails) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


