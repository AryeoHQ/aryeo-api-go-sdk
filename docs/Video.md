# Video

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the video. | 
**Title** | Pointer to **NullableString** | The title of the video given by the uploading user. | [optional] 
**Duration** | Pointer to **NullableInt32** | The video&#39;s runtime in seconds. | [optional] 
**DisplayType** | **string** | The display type determines if the video is branded or unbranded (can also be none or both). This affects whether the video is displayed on branded or unbranded marketing materials such as the property website. | 
**SourceType** | **string** | The original upload source of the video, used to determine how to handle the playback_url of the video and other display properties.  | 
**ThumbnailUrl** | **string** | A thumbnail image URL for the video. | 
**PlaybackUrl** | **string** | A URL linking to playback stream of the video. | 
**DownloadUrl** | Pointer to **NullableString** | A URL for downloading the video. | [optional] 
**ShareUrl** | Pointer to **NullableString** | A URL linking to a public viewing optimized webpage the video. | [optional] 

## Methods

### NewVideo

`func NewVideo(id string, displayType string, sourceType string, thumbnailUrl string, playbackUrl string, ) *Video`

NewVideo instantiates a new Video object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoWithDefaults

`func NewVideoWithDefaults() *Video`

NewVideoWithDefaults instantiates a new Video object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Video) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Video) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Video) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *Video) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Video) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Video) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Video) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Video) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Video) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDuration

`func (o *Video) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Video) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Video) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Video) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *Video) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *Video) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetDisplayType

`func (o *Video) GetDisplayType() string`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *Video) GetDisplayTypeOk() (*string, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *Video) SetDisplayType(v string)`

SetDisplayType sets DisplayType field to given value.


### GetSourceType

`func (o *Video) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *Video) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *Video) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetThumbnailUrl

`func (o *Video) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *Video) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *Video) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.


### GetPlaybackUrl

`func (o *Video) GetPlaybackUrl() string`

GetPlaybackUrl returns the PlaybackUrl field if non-nil, zero value otherwise.

### GetPlaybackUrlOk

`func (o *Video) GetPlaybackUrlOk() (*string, bool)`

GetPlaybackUrlOk returns a tuple with the PlaybackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackUrl

`func (o *Video) SetPlaybackUrl(v string)`

SetPlaybackUrl sets PlaybackUrl field to given value.


### GetDownloadUrl

`func (o *Video) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *Video) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *Video) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *Video) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### SetDownloadUrlNil

`func (o *Video) SetDownloadUrlNil(b bool)`

 SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil

### UnsetDownloadUrl
`func (o *Video) UnsetDownloadUrl()`

UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
### GetShareUrl

`func (o *Video) GetShareUrl() string`

GetShareUrl returns the ShareUrl field if non-nil, zero value otherwise.

### GetShareUrlOk

`func (o *Video) GetShareUrlOk() (*string, bool)`

GetShareUrlOk returns a tuple with the ShareUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareUrl

`func (o *Video) SetShareUrl(v string)`

SetShareUrl sets ShareUrl field to given value.

### HasShareUrl

`func (o *Video) HasShareUrl() bool`

HasShareUrl returns a boolean if a field has been set.

### SetShareUrlNil

`func (o *Video) SetShareUrlNil(b bool)`

 SetShareUrlNil sets the value for ShareUrl to be an explicit nil

### UnsetShareUrl
`func (o *Video) UnsetShareUrl()`

UnsetShareUrl ensures that no value is present for ShareUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


