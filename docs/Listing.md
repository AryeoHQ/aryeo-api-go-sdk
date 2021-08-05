# Listing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the listing. UUID Version 4. | 
**Address** | [**Address**](Address.md) |  | 
**MlsNumber** | Pointer to **NullableString** | The identifier for a listing on its local MLS.  | [optional] 
**Type** | Pointer to **NullableString** | General type of the listing, primarily categorizing its use case. Examples include residential and commercial.  | [optional] 
**SubType** | Pointer to **NullableString** | Further specifies the listing type. Examples include family residence and condominium. | [optional] 
**Status** | Pointer to **NullableString** | Local, regional, or otherwise custom status for the listing used by the parties involved in the listing transaction. While variable, these statuses are typically mapped to the listing&#39;s standard status. | [optional] 
**StandardStatus** | Pointer to **NullableString** | The status of the listing as it reflects the state of the contract between the listing agent and seller or an agreement with a buyer, including Active, Active Under Contract, Canceled, Closed, Expired, Pending, and Withdrawn. | [optional] 
**Description** | Pointer to **NullableString** | Description of the selling points of the building and/or land for sale.  | [optional] 
**Lot** | Pointer to [**ListingLot**](ListingLot.md) |  | [optional] 
**Building** | Pointer to [**ListingBuilding**](ListingBuilding.md) |  | [optional] 
**Price** | Pointer to [**ListingPrice**](ListingPrice.md) |  | [optional] 
**ListAgent** | Pointer to [**Group**](Group.md) |  | [optional] 
**CoListAgent** | Pointer to [**Group**](Group.md) |  | [optional] 
**Images** | Pointer to [**[]Image**](Image.md) | images | [optional] 
**Videos** | Pointer to [**[]Video**](Video.md) | videos | [optional] 
**FloorPlans** | Pointer to [**[]FloorPlan**](FloorPlan.md) | floor_plans | [optional] 
**InteractiveContent** | Pointer to [**[]InteractiveContent**](InteractiveContent.md) | interactive_content | [optional] 
**PropertyWebsite** | Pointer to [**PropertyWebsite**](PropertyWebsite.md) |  | [optional] 
**Orders** | Pointer to [**[]Order**](Order.md) | orders | [optional] 
**DownloadsEnabled** | **bool** | Are downloads enabled for this listing? | 

## Methods

### NewListing

`func NewListing(id string, address Address, downloadsEnabled bool, ) *Listing`

NewListing instantiates a new Listing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingWithDefaults

`func NewListingWithDefaults() *Listing`

NewListingWithDefaults instantiates a new Listing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Listing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Listing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Listing) SetId(v string)`

SetId sets Id field to given value.


### GetAddress

`func (o *Listing) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Listing) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Listing) SetAddress(v Address)`

SetAddress sets Address field to given value.


### GetMlsNumber

`func (o *Listing) GetMlsNumber() string`

GetMlsNumber returns the MlsNumber field if non-nil, zero value otherwise.

### GetMlsNumberOk

`func (o *Listing) GetMlsNumberOk() (*string, bool)`

GetMlsNumberOk returns a tuple with the MlsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlsNumber

`func (o *Listing) SetMlsNumber(v string)`

SetMlsNumber sets MlsNumber field to given value.

### HasMlsNumber

`func (o *Listing) HasMlsNumber() bool`

HasMlsNumber returns a boolean if a field has been set.

### SetMlsNumberNil

`func (o *Listing) SetMlsNumberNil(b bool)`

 SetMlsNumberNil sets the value for MlsNumber to be an explicit nil

### UnsetMlsNumber
`func (o *Listing) UnsetMlsNumber()`

UnsetMlsNumber ensures that no value is present for MlsNumber, not even an explicit nil
### GetType

`func (o *Listing) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Listing) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Listing) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Listing) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Listing) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Listing) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetSubType

`func (o *Listing) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *Listing) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *Listing) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *Listing) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### SetSubTypeNil

`func (o *Listing) SetSubTypeNil(b bool)`

 SetSubTypeNil sets the value for SubType to be an explicit nil

### UnsetSubType
`func (o *Listing) UnsetSubType()`

UnsetSubType ensures that no value is present for SubType, not even an explicit nil
### GetStatus

`func (o *Listing) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Listing) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Listing) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Listing) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Listing) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Listing) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStandardStatus

`func (o *Listing) GetStandardStatus() string`

GetStandardStatus returns the StandardStatus field if non-nil, zero value otherwise.

### GetStandardStatusOk

`func (o *Listing) GetStandardStatusOk() (*string, bool)`

GetStandardStatusOk returns a tuple with the StandardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardStatus

`func (o *Listing) SetStandardStatus(v string)`

SetStandardStatus sets StandardStatus field to given value.

### HasStandardStatus

`func (o *Listing) HasStandardStatus() bool`

HasStandardStatus returns a boolean if a field has been set.

### SetStandardStatusNil

`func (o *Listing) SetStandardStatusNil(b bool)`

 SetStandardStatusNil sets the value for StandardStatus to be an explicit nil

### UnsetStandardStatus
`func (o *Listing) UnsetStandardStatus()`

UnsetStandardStatus ensures that no value is present for StandardStatus, not even an explicit nil
### GetDescription

`func (o *Listing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Listing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Listing) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Listing) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Listing) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Listing) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLot

`func (o *Listing) GetLot() ListingLot`

GetLot returns the Lot field if non-nil, zero value otherwise.

### GetLotOk

`func (o *Listing) GetLotOk() (*ListingLot, bool)`

GetLotOk returns a tuple with the Lot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLot

`func (o *Listing) SetLot(v ListingLot)`

SetLot sets Lot field to given value.

### HasLot

`func (o *Listing) HasLot() bool`

HasLot returns a boolean if a field has been set.

### GetBuilding

`func (o *Listing) GetBuilding() ListingBuilding`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *Listing) GetBuildingOk() (*ListingBuilding, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *Listing) SetBuilding(v ListingBuilding)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *Listing) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### GetPrice

`func (o *Listing) GetPrice() ListingPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Listing) GetPriceOk() (*ListingPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Listing) SetPrice(v ListingPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Listing) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetListAgent

`func (o *Listing) GetListAgent() Group`

GetListAgent returns the ListAgent field if non-nil, zero value otherwise.

### GetListAgentOk

`func (o *Listing) GetListAgentOk() (*Group, bool)`

GetListAgentOk returns a tuple with the ListAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListAgent

`func (o *Listing) SetListAgent(v Group)`

SetListAgent sets ListAgent field to given value.

### HasListAgent

`func (o *Listing) HasListAgent() bool`

HasListAgent returns a boolean if a field has been set.

### GetCoListAgent

`func (o *Listing) GetCoListAgent() Group`

GetCoListAgent returns the CoListAgent field if non-nil, zero value otherwise.

### GetCoListAgentOk

`func (o *Listing) GetCoListAgentOk() (*Group, bool)`

GetCoListAgentOk returns a tuple with the CoListAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoListAgent

`func (o *Listing) SetCoListAgent(v Group)`

SetCoListAgent sets CoListAgent field to given value.

### HasCoListAgent

`func (o *Listing) HasCoListAgent() bool`

HasCoListAgent returns a boolean if a field has been set.

### GetImages

`func (o *Listing) GetImages() []Image`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *Listing) GetImagesOk() (*[]Image, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *Listing) SetImages(v []Image)`

SetImages sets Images field to given value.

### HasImages

`func (o *Listing) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetVideos

`func (o *Listing) GetVideos() []Video`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *Listing) GetVideosOk() (*[]Video, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *Listing) SetVideos(v []Video)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *Listing) HasVideos() bool`

HasVideos returns a boolean if a field has been set.

### GetFloorPlans

`func (o *Listing) GetFloorPlans() []FloorPlan`

GetFloorPlans returns the FloorPlans field if non-nil, zero value otherwise.

### GetFloorPlansOk

`func (o *Listing) GetFloorPlansOk() (*[]FloorPlan, bool)`

GetFloorPlansOk returns a tuple with the FloorPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorPlans

`func (o *Listing) SetFloorPlans(v []FloorPlan)`

SetFloorPlans sets FloorPlans field to given value.

### HasFloorPlans

`func (o *Listing) HasFloorPlans() bool`

HasFloorPlans returns a boolean if a field has been set.

### GetInteractiveContent

`func (o *Listing) GetInteractiveContent() []InteractiveContent`

GetInteractiveContent returns the InteractiveContent field if non-nil, zero value otherwise.

### GetInteractiveContentOk

`func (o *Listing) GetInteractiveContentOk() (*[]InteractiveContent, bool)`

GetInteractiveContentOk returns a tuple with the InteractiveContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractiveContent

`func (o *Listing) SetInteractiveContent(v []InteractiveContent)`

SetInteractiveContent sets InteractiveContent field to given value.

### HasInteractiveContent

`func (o *Listing) HasInteractiveContent() bool`

HasInteractiveContent returns a boolean if a field has been set.

### GetPropertyWebsite

`func (o *Listing) GetPropertyWebsite() PropertyWebsite`

GetPropertyWebsite returns the PropertyWebsite field if non-nil, zero value otherwise.

### GetPropertyWebsiteOk

`func (o *Listing) GetPropertyWebsiteOk() (*PropertyWebsite, bool)`

GetPropertyWebsiteOk returns a tuple with the PropertyWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyWebsite

`func (o *Listing) SetPropertyWebsite(v PropertyWebsite)`

SetPropertyWebsite sets PropertyWebsite field to given value.

### HasPropertyWebsite

`func (o *Listing) HasPropertyWebsite() bool`

HasPropertyWebsite returns a boolean if a field has been set.

### GetOrders

`func (o *Listing) GetOrders() []Order`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *Listing) GetOrdersOk() (*[]Order, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *Listing) SetOrders(v []Order)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *Listing) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetDownloadsEnabled

`func (o *Listing) GetDownloadsEnabled() bool`

GetDownloadsEnabled returns the DownloadsEnabled field if non-nil, zero value otherwise.

### GetDownloadsEnabledOk

`func (o *Listing) GetDownloadsEnabledOk() (*bool, bool)`

GetDownloadsEnabledOk returns a tuple with the DownloadsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadsEnabled

`func (o *Listing) SetDownloadsEnabled(v bool)`

SetDownloadsEnabled sets DownloadsEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


