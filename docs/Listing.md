# Listing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the listing. | 
**Address** | [**PartialAddress**](PartialAddress.md) |  | 
**DeliveryStatus** | **string** | Has this listing been delivered? | 
**ThumbnailUrl** | Pointer to **NullableString** | Thumbnail URL for the listing. | [optional] 
**Agent** | Pointer to [**Group**](Group.md) |  | [optional] 
**CoAgent** | Pointer to [**Group**](Group.md) |  | [optional] 
**Images** | Pointer to [**[]Image**](Image.md) | images | [optional] 
**Videos** | Pointer to [**[]Video**](Video.md) | videos | [optional] 
**FloorPlans** | Pointer to [**[]FloorPlan**](FloorPlan.md) | floor_plans | [optional] 
**PropertyWebsites** | Pointer to [**PropertyWebsites**](PropertyWebsites.md) |  | [optional] 
**InteractiveContent** | Pointer to [**[]InteractiveContent**](InteractiveContent.md) | interactive_content | [optional] 
**PropertyDetails** | Pointer to [**PropertyDetails**](PropertyDetails.md) |  | [optional] 
**DownloadsEnabled** | **bool** | Are downloads enabled for this listing? | 
**Orders** | Pointer to [**[]Order**](Order.md) | orders | [optional] 

## Methods

### NewListing

`func NewListing(id string, address PartialAddress, deliveryStatus string, downloadsEnabled bool, ) *Listing`

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

`func (o *Listing) GetAddress() PartialAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Listing) GetAddressOk() (*PartialAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Listing) SetAddress(v PartialAddress)`

SetAddress sets Address field to given value.


### GetDeliveryStatus

`func (o *Listing) GetDeliveryStatus() string`

GetDeliveryStatus returns the DeliveryStatus field if non-nil, zero value otherwise.

### GetDeliveryStatusOk

`func (o *Listing) GetDeliveryStatusOk() (*string, bool)`

GetDeliveryStatusOk returns a tuple with the DeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatus

`func (o *Listing) SetDeliveryStatus(v string)`

SetDeliveryStatus sets DeliveryStatus field to given value.


### GetThumbnailUrl

`func (o *Listing) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *Listing) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *Listing) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *Listing) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *Listing) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *Listing) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
### GetAgent

`func (o *Listing) GetAgent() Group`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *Listing) GetAgentOk() (*Group, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *Listing) SetAgent(v Group)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *Listing) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetCoAgent

`func (o *Listing) GetCoAgent() Group`

GetCoAgent returns the CoAgent field if non-nil, zero value otherwise.

### GetCoAgentOk

`func (o *Listing) GetCoAgentOk() (*Group, bool)`

GetCoAgentOk returns a tuple with the CoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoAgent

`func (o *Listing) SetCoAgent(v Group)`

SetCoAgent sets CoAgent field to given value.

### HasCoAgent

`func (o *Listing) HasCoAgent() bool`

HasCoAgent returns a boolean if a field has been set.

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

### GetPropertyWebsites

`func (o *Listing) GetPropertyWebsites() PropertyWebsites`

GetPropertyWebsites returns the PropertyWebsites field if non-nil, zero value otherwise.

### GetPropertyWebsitesOk

`func (o *Listing) GetPropertyWebsitesOk() (*PropertyWebsites, bool)`

GetPropertyWebsitesOk returns a tuple with the PropertyWebsites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyWebsites

`func (o *Listing) SetPropertyWebsites(v PropertyWebsites)`

SetPropertyWebsites sets PropertyWebsites field to given value.

### HasPropertyWebsites

`func (o *Listing) HasPropertyWebsites() bool`

HasPropertyWebsites returns a boolean if a field has been set.

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

### GetPropertyDetails

`func (o *Listing) GetPropertyDetails() PropertyDetails`

GetPropertyDetails returns the PropertyDetails field if non-nil, zero value otherwise.

### GetPropertyDetailsOk

`func (o *Listing) GetPropertyDetailsOk() (*PropertyDetails, bool)`

GetPropertyDetailsOk returns a tuple with the PropertyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyDetails

`func (o *Listing) SetPropertyDetails(v PropertyDetails)`

SetPropertyDetails sets PropertyDetails field to given value.

### HasPropertyDetails

`func (o *Listing) HasPropertyDetails() bool`

HasPropertyDetails returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


