/*
 * Aryeo
 *
 *
 * API version: 1.0.0
 * Contact: jarrod@aryeo.com
 */

package aryeo

import (
	"encoding/json"
)

// Video A recording of moving visual images. Provided either as a download URL (MP4) or a link (e.g. YouTube, Vimeo).
type Video struct {
	// ID of the video.
	Id int32 `json:"id"`
	// The title of the video given by the uploading user.
	Title NullableString `json:"title,omitempty"`
	// The display type determines if the video is branded or unbranded (can also be none or both). This affects whether the video is displayed on branded or unbranded marketing materials such as the property website.
	DisplayType string `json:"display_type"`
	// The original upload source of the video, used to determine how to handle the playback_url of the video and other display properties. 
	SourceType string `json:"source_type"`
	// Thumbnail URL for the video.
	ThumbnailUrl string `json:"thumbnail_url"`
	// A URL linking to the video.
	PlaybackUrl string `json:"playback_url"`
	// A URL for downloading the video.
	DownloadUrl NullableString `json:"download_url,omitempty"`
	// The video's runtime in seconds.
	Seconds NullableInt32 `json:"seconds,omitempty"`
	// Aryeo iFrame player URL
	ShareUrl NullableString `json:"share_url,omitempty"`
}

// NewVideo instantiates a new Video object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideo(id int32, displayType string, sourceType string, thumbnailUrl string, playbackUrl string) *Video {
	this := Video{}
	this.Id = id
	this.DisplayType = displayType
	this.SourceType = sourceType
	this.ThumbnailUrl = thumbnailUrl
	this.PlaybackUrl = playbackUrl
	return &this
}

// NewVideoWithDefaults instantiates a new Video object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoWithDefaults() *Video {
	this := Video{}
	return &this
}

// GetId returns the Id field value
func (o *Video) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Video) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Video) SetId(v int32) {
	o.Id = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Video) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Video) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Video) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Video) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Video) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Video) UnsetTitle() {
	o.Title.Unset()
}

// GetDisplayType returns the DisplayType field value
func (o *Video) GetDisplayType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayType
}

// GetDisplayTypeOk returns a tuple with the DisplayType field value
// and a boolean to check if the value has been set.
func (o *Video) GetDisplayTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayType, true
}

// SetDisplayType sets field value
func (o *Video) SetDisplayType(v string) {
	o.DisplayType = v
}

// GetSourceType returns the SourceType field value
func (o *Video) GetSourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *Video) GetSourceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *Video) SetSourceType(v string) {
	o.SourceType = v
}

// GetThumbnailUrl returns the ThumbnailUrl field value
func (o *Video) GetThumbnailUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value
// and a boolean to check if the value has been set.
func (o *Video) GetThumbnailUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ThumbnailUrl, true
}

// SetThumbnailUrl sets field value
func (o *Video) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = v
}

// GetPlaybackUrl returns the PlaybackUrl field value
func (o *Video) GetPlaybackUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlaybackUrl
}

// GetPlaybackUrlOk returns a tuple with the PlaybackUrl field value
// and a boolean to check if the value has been set.
func (o *Video) GetPlaybackUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PlaybackUrl, true
}

// SetPlaybackUrl sets field value
func (o *Video) SetPlaybackUrl(v string) {
	o.PlaybackUrl = v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Video) GetDownloadUrl() string {
	if o == nil || o.DownloadUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.DownloadUrl.Get()
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Video) GetDownloadUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DownloadUrl.Get(), o.DownloadUrl.IsSet()
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *Video) HasDownloadUrl() bool {
	if o != nil && o.DownloadUrl.IsSet() {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given NullableString and assigns it to the DownloadUrl field.
func (o *Video) SetDownloadUrl(v string) {
	o.DownloadUrl.Set(&v)
}
// SetDownloadUrlNil sets the value for DownloadUrl to be an explicit nil
func (o *Video) SetDownloadUrlNil() {
	o.DownloadUrl.Set(nil)
}

// UnsetDownloadUrl ensures that no value is present for DownloadUrl, not even an explicit nil
func (o *Video) UnsetDownloadUrl() {
	o.DownloadUrl.Unset()
}

// GetSeconds returns the Seconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Video) GetSeconds() int32 {
	if o == nil || o.Seconds.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Seconds.Get()
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Video) GetSecondsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Seconds.Get(), o.Seconds.IsSet()
}

// HasSeconds returns a boolean if a field has been set.
func (o *Video) HasSeconds() bool {
	if o != nil && o.Seconds.IsSet() {
		return true
	}

	return false
}

// SetSeconds gets a reference to the given NullableInt32 and assigns it to the Seconds field.
func (o *Video) SetSeconds(v int32) {
	o.Seconds.Set(&v)
}
// SetSecondsNil sets the value for Seconds to be an explicit nil
func (o *Video) SetSecondsNil() {
	o.Seconds.Set(nil)
}

// UnsetSeconds ensures that no value is present for Seconds, not even an explicit nil
func (o *Video) UnsetSeconds() {
	o.Seconds.Unset()
}

// GetShareUrl returns the ShareUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Video) GetShareUrl() string {
	if o == nil || o.ShareUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ShareUrl.Get()
}

// GetShareUrlOk returns a tuple with the ShareUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Video) GetShareUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ShareUrl.Get(), o.ShareUrl.IsSet()
}

// HasShareUrl returns a boolean if a field has been set.
func (o *Video) HasShareUrl() bool {
	if o != nil && o.ShareUrl.IsSet() {
		return true
	}

	return false
}

// SetShareUrl gets a reference to the given NullableString and assigns it to the ShareUrl field.
func (o *Video) SetShareUrl(v string) {
	o.ShareUrl.Set(&v)
}
// SetShareUrlNil sets the value for ShareUrl to be an explicit nil
func (o *Video) SetShareUrlNil() {
	o.ShareUrl.Set(nil)
}

// UnsetShareUrl ensures that no value is present for ShareUrl, not even an explicit nil
func (o *Video) UnsetShareUrl() {
	o.ShareUrl.Unset()
}

func (o Video) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if true {
		toSerialize["display_type"] = o.DisplayType
	}
	if true {
		toSerialize["source_type"] = o.SourceType
	}
	if true {
		toSerialize["thumbnail_url"] = o.ThumbnailUrl
	}
	if true {
		toSerialize["playback_url"] = o.PlaybackUrl
	}
	if o.DownloadUrl.IsSet() {
		toSerialize["download_url"] = o.DownloadUrl.Get()
	}
	if o.Seconds.IsSet() {
		toSerialize["seconds"] = o.Seconds.Get()
	}
	if o.ShareUrl.IsSet() {
		toSerialize["share_url"] = o.ShareUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableVideo struct {
	value *Video
	isSet bool
}

func (v NullableVideo) Get() *Video {
	return v.value
}

func (v *NullableVideo) Set(val *Video) {
	v.value = val
	v.isSet = true
}

func (v NullableVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideo(val *Video) *NullableVideo {
	return &NullableVideo{value: val, isSet: true}
}

func (v NullableVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


