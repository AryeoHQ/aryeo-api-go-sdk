# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of the image. | 
**ThumbnailUrl** | **string** | A URL for a thumbnail-sized version of the image. | 
**LargeUrl** | **string** | A URL for a large version of the image. | 
**OriginalUrl** | **string** | A URL for the original, full-resolution version of the image. Useful for downloading. | 
**Index** | Pointer to **NullableInt32** | The order in which the image should be displayed amongst other related images. | [optional] 
**Caption** | Pointer to **NullableString** | A brief explanation of the image. | [optional] 
**DisplayInGallery** | **bool** | Should the image be displayed in a gallery? | 

## Methods

### NewImage

`func NewImage(id int32, thumbnailUrl string, largeUrl string, originalUrl string, displayInGallery bool, ) *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Image) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Image) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Image) SetId(v int32)`

SetId sets Id field to given value.


### GetThumbnailUrl

`func (o *Image) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *Image) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *Image) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.


### GetLargeUrl

`func (o *Image) GetLargeUrl() string`

GetLargeUrl returns the LargeUrl field if non-nil, zero value otherwise.

### GetLargeUrlOk

`func (o *Image) GetLargeUrlOk() (*string, bool)`

GetLargeUrlOk returns a tuple with the LargeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeUrl

`func (o *Image) SetLargeUrl(v string)`

SetLargeUrl sets LargeUrl field to given value.


### GetOriginalUrl

`func (o *Image) GetOriginalUrl() string`

GetOriginalUrl returns the OriginalUrl field if non-nil, zero value otherwise.

### GetOriginalUrlOk

`func (o *Image) GetOriginalUrlOk() (*string, bool)`

GetOriginalUrlOk returns a tuple with the OriginalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalUrl

`func (o *Image) SetOriginalUrl(v string)`

SetOriginalUrl sets OriginalUrl field to given value.


### GetIndex

`func (o *Image) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Image) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Image) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Image) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndexNil

`func (o *Image) SetIndexNil(b bool)`

 SetIndexNil sets the value for Index to be an explicit nil

### UnsetIndex
`func (o *Image) UnsetIndex()`

UnsetIndex ensures that no value is present for Index, not even an explicit nil
### GetCaption

`func (o *Image) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *Image) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *Image) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *Image) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### SetCaptionNil

`func (o *Image) SetCaptionNil(b bool)`

 SetCaptionNil sets the value for Caption to be an explicit nil

### UnsetCaption
`func (o *Image) UnsetCaption()`

UnsetCaption ensures that no value is present for Caption, not even an explicit nil
### GetDisplayInGallery

`func (o *Image) GetDisplayInGallery() bool`

GetDisplayInGallery returns the DisplayInGallery field if non-nil, zero value otherwise.

### GetDisplayInGalleryOk

`func (o *Image) GetDisplayInGalleryOk() (*bool, bool)`

GetDisplayInGalleryOk returns a tuple with the DisplayInGallery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInGallery

`func (o *Image) SetDisplayInGallery(v bool)`

SetDisplayInGallery sets DisplayInGallery field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


