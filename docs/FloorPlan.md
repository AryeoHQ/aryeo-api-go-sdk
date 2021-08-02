# FloorPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the floor plan. | 
**OriginalUrl** | **string** | A URL for the original, full-resolution version of the floor plan. Useful for downloading. | 
**LargeUrl** | **string** | A URL for a large version of the floor plan. | 
**ThumbnailUrl** | **string** | A URL for a thumbnail-sized version of the floor plan. | 
**Title** | Pointer to **NullableString** | The title (or caption) of the floor plan. | [optional] 
**Index** | Pointer to **NullableInt32** | Index order position of the floor plan. | [optional] 

## Methods

### NewFloorPlan

`func NewFloorPlan(id string, originalUrl string, largeUrl string, thumbnailUrl string, ) *FloorPlan`

NewFloorPlan instantiates a new FloorPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFloorPlanWithDefaults

`func NewFloorPlanWithDefaults() *FloorPlan`

NewFloorPlanWithDefaults instantiates a new FloorPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FloorPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FloorPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FloorPlan) SetId(v string)`

SetId sets Id field to given value.


### GetOriginalUrl

`func (o *FloorPlan) GetOriginalUrl() string`

GetOriginalUrl returns the OriginalUrl field if non-nil, zero value otherwise.

### GetOriginalUrlOk

`func (o *FloorPlan) GetOriginalUrlOk() (*string, bool)`

GetOriginalUrlOk returns a tuple with the OriginalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalUrl

`func (o *FloorPlan) SetOriginalUrl(v string)`

SetOriginalUrl sets OriginalUrl field to given value.


### GetLargeUrl

`func (o *FloorPlan) GetLargeUrl() string`

GetLargeUrl returns the LargeUrl field if non-nil, zero value otherwise.

### GetLargeUrlOk

`func (o *FloorPlan) GetLargeUrlOk() (*string, bool)`

GetLargeUrlOk returns a tuple with the LargeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeUrl

`func (o *FloorPlan) SetLargeUrl(v string)`

SetLargeUrl sets LargeUrl field to given value.


### GetThumbnailUrl

`func (o *FloorPlan) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *FloorPlan) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *FloorPlan) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.


### GetTitle

`func (o *FloorPlan) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FloorPlan) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FloorPlan) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FloorPlan) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *FloorPlan) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *FloorPlan) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetIndex

`func (o *FloorPlan) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *FloorPlan) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *FloorPlan) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *FloorPlan) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndexNil

`func (o *FloorPlan) SetIndexNil(b bool)`

 SetIndexNil sets the value for Index to be an explicit nil

### UnsetIndex
`func (o *FloorPlan) UnsetIndex()`

UnsetIndex ensures that no value is present for Index, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


