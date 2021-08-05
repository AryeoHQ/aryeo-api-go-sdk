# InteractiveContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the content. UUID Version 4. | 
**DisplayType** | **string** | Is the content branded, unbranded, or both? | 
**ContentType** | **string** | The type of interactive content. | 
**Url** | **string** | URL for the content. | 
**ThumbnailUrl** | Pointer to **NullableString** | A URL for a thumbnail-sized preview of the content. | [optional] 

## Methods

### NewInteractiveContent

`func NewInteractiveContent(id string, displayType string, contentType string, url string, ) *InteractiveContent`

NewInteractiveContent instantiates a new InteractiveContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteractiveContentWithDefaults

`func NewInteractiveContentWithDefaults() *InteractiveContent`

NewInteractiveContentWithDefaults instantiates a new InteractiveContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InteractiveContent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InteractiveContent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InteractiveContent) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayType

`func (o *InteractiveContent) GetDisplayType() string`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *InteractiveContent) GetDisplayTypeOk() (*string, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *InteractiveContent) SetDisplayType(v string)`

SetDisplayType sets DisplayType field to given value.


### GetContentType

`func (o *InteractiveContent) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *InteractiveContent) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *InteractiveContent) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetUrl

`func (o *InteractiveContent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InteractiveContent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InteractiveContent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetThumbnailUrl

`func (o *InteractiveContent) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *InteractiveContent) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *InteractiveContent) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *InteractiveContent) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *InteractiveContent) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *InteractiveContent) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


