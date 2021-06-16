# PartialListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the listing. | 
**Address** | [**PartialAddress**](PartialAddress.md) |  | 
**DeliveryStatus** | **string** | Has the listing been delivered? | 
**ThumbnailUrl** | Pointer to **NullableString** | Thumbnail URL for the listing. | [optional] 
**Price** | Pointer to **NullableInt32** | The price of the property in dollars. | [optional] 
**BrandedUrl** | Pointer to **NullableString** | URL for branded property website. | [optional] 
**SquareFeet** | Pointer to **NullableFloat32** | Total number of square feet. | [optional] 
**Bedrooms** | Pointer to **NullableInt32** | Number of bedrooms. | [optional] 
**Bathrooms** | Pointer to **NullableFloat32** | Number of bathrooms. | [optional] 
**DownloadsEnabled** | **bool** | Are downloads enabled for this listing? | 
**Status** | Pointer to **NullableString** | Sales status for the listing. | [optional] 
**PropertyDetails** | Pointer to [**PropertyDetails**](PropertyDetails.md) |  | [optional] 
**Agent** | Pointer to [**PartialGroup**](PartialGroup.md) |  | [optional] 
**CoAgent** | Pointer to [**PartialGroup**](PartialGroup.md) |  | [optional] 

## Methods

### NewPartialListing

`func NewPartialListing(id string, address PartialAddress, deliveryStatus string, downloadsEnabled bool, ) *PartialListing`

NewPartialListing instantiates a new PartialListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialListingWithDefaults

`func NewPartialListingWithDefaults() *PartialListing`

NewPartialListingWithDefaults instantiates a new PartialListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PartialListing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartialListing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartialListing) SetId(v string)`

SetId sets Id field to given value.


### GetAddress

`func (o *PartialListing) GetAddress() PartialAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PartialListing) GetAddressOk() (*PartialAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PartialListing) SetAddress(v PartialAddress)`

SetAddress sets Address field to given value.


### GetDeliveryStatus

`func (o *PartialListing) GetDeliveryStatus() string`

GetDeliveryStatus returns the DeliveryStatus field if non-nil, zero value otherwise.

### GetDeliveryStatusOk

`func (o *PartialListing) GetDeliveryStatusOk() (*string, bool)`

GetDeliveryStatusOk returns a tuple with the DeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatus

`func (o *PartialListing) SetDeliveryStatus(v string)`

SetDeliveryStatus sets DeliveryStatus field to given value.


### GetThumbnailUrl

`func (o *PartialListing) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *PartialListing) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *PartialListing) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *PartialListing) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *PartialListing) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *PartialListing) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
### GetPrice

`func (o *PartialListing) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PartialListing) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PartialListing) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PartialListing) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *PartialListing) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *PartialListing) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetBrandedUrl

`func (o *PartialListing) GetBrandedUrl() string`

GetBrandedUrl returns the BrandedUrl field if non-nil, zero value otherwise.

### GetBrandedUrlOk

`func (o *PartialListing) GetBrandedUrlOk() (*string, bool)`

GetBrandedUrlOk returns a tuple with the BrandedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedUrl

`func (o *PartialListing) SetBrandedUrl(v string)`

SetBrandedUrl sets BrandedUrl field to given value.

### HasBrandedUrl

`func (o *PartialListing) HasBrandedUrl() bool`

HasBrandedUrl returns a boolean if a field has been set.

### SetBrandedUrlNil

`func (o *PartialListing) SetBrandedUrlNil(b bool)`

 SetBrandedUrlNil sets the value for BrandedUrl to be an explicit nil

### UnsetBrandedUrl
`func (o *PartialListing) UnsetBrandedUrl()`

UnsetBrandedUrl ensures that no value is present for BrandedUrl, not even an explicit nil
### GetSquareFeet

`func (o *PartialListing) GetSquareFeet() float32`

GetSquareFeet returns the SquareFeet field if non-nil, zero value otherwise.

### GetSquareFeetOk

`func (o *PartialListing) GetSquareFeetOk() (*float32, bool)`

GetSquareFeetOk returns a tuple with the SquareFeet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquareFeet

`func (o *PartialListing) SetSquareFeet(v float32)`

SetSquareFeet sets SquareFeet field to given value.

### HasSquareFeet

`func (o *PartialListing) HasSquareFeet() bool`

HasSquareFeet returns a boolean if a field has been set.

### SetSquareFeetNil

`func (o *PartialListing) SetSquareFeetNil(b bool)`

 SetSquareFeetNil sets the value for SquareFeet to be an explicit nil

### UnsetSquareFeet
`func (o *PartialListing) UnsetSquareFeet()`

UnsetSquareFeet ensures that no value is present for SquareFeet, not even an explicit nil
### GetBedrooms

`func (o *PartialListing) GetBedrooms() int32`

GetBedrooms returns the Bedrooms field if non-nil, zero value otherwise.

### GetBedroomsOk

`func (o *PartialListing) GetBedroomsOk() (*int32, bool)`

GetBedroomsOk returns a tuple with the Bedrooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBedrooms

`func (o *PartialListing) SetBedrooms(v int32)`

SetBedrooms sets Bedrooms field to given value.

### HasBedrooms

`func (o *PartialListing) HasBedrooms() bool`

HasBedrooms returns a boolean if a field has been set.

### SetBedroomsNil

`func (o *PartialListing) SetBedroomsNil(b bool)`

 SetBedroomsNil sets the value for Bedrooms to be an explicit nil

### UnsetBedrooms
`func (o *PartialListing) UnsetBedrooms()`

UnsetBedrooms ensures that no value is present for Bedrooms, not even an explicit nil
### GetBathrooms

`func (o *PartialListing) GetBathrooms() float32`

GetBathrooms returns the Bathrooms field if non-nil, zero value otherwise.

### GetBathroomsOk

`func (o *PartialListing) GetBathroomsOk() (*float32, bool)`

GetBathroomsOk returns a tuple with the Bathrooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBathrooms

`func (o *PartialListing) SetBathrooms(v float32)`

SetBathrooms sets Bathrooms field to given value.

### HasBathrooms

`func (o *PartialListing) HasBathrooms() bool`

HasBathrooms returns a boolean if a field has been set.

### SetBathroomsNil

`func (o *PartialListing) SetBathroomsNil(b bool)`

 SetBathroomsNil sets the value for Bathrooms to be an explicit nil

### UnsetBathrooms
`func (o *PartialListing) UnsetBathrooms()`

UnsetBathrooms ensures that no value is present for Bathrooms, not even an explicit nil
### GetDownloadsEnabled

`func (o *PartialListing) GetDownloadsEnabled() bool`

GetDownloadsEnabled returns the DownloadsEnabled field if non-nil, zero value otherwise.

### GetDownloadsEnabledOk

`func (o *PartialListing) GetDownloadsEnabledOk() (*bool, bool)`

GetDownloadsEnabledOk returns a tuple with the DownloadsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadsEnabled

`func (o *PartialListing) SetDownloadsEnabled(v bool)`

SetDownloadsEnabled sets DownloadsEnabled field to given value.


### GetStatus

`func (o *PartialListing) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PartialListing) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PartialListing) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PartialListing) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *PartialListing) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *PartialListing) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetPropertyDetails

`func (o *PartialListing) GetPropertyDetails() PropertyDetails`

GetPropertyDetails returns the PropertyDetails field if non-nil, zero value otherwise.

### GetPropertyDetailsOk

`func (o *PartialListing) GetPropertyDetailsOk() (*PropertyDetails, bool)`

GetPropertyDetailsOk returns a tuple with the PropertyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyDetails

`func (o *PartialListing) SetPropertyDetails(v PropertyDetails)`

SetPropertyDetails sets PropertyDetails field to given value.

### HasPropertyDetails

`func (o *PartialListing) HasPropertyDetails() bool`

HasPropertyDetails returns a boolean if a field has been set.

### GetAgent

`func (o *PartialListing) GetAgent() PartialGroup`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *PartialListing) GetAgentOk() (*PartialGroup, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *PartialListing) SetAgent(v PartialGroup)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *PartialListing) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetCoAgent

`func (o *PartialListing) GetCoAgent() PartialGroup`

GetCoAgent returns the CoAgent field if non-nil, zero value otherwise.

### GetCoAgentOk

`func (o *PartialListing) GetCoAgentOk() (*PartialGroup, bool)`

GetCoAgentOk returns a tuple with the CoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoAgent

`func (o *PartialListing) SetCoAgent(v PartialGroup)`

SetCoAgent sets CoAgent field to given value.

### HasCoAgent

`func (o *PartialListing) HasCoAgent() bool`

HasCoAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


