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

// Image A visual representation of something.
type Image struct {
	// ID of the image.
	Id int32 `json:"id"`
	// A URL for a thumbnail-sized version of the image.
	ThumbnailUrl string `json:"thumbnail_url"`
	// A URL for a large version of the image.
	LargeUrl string `json:"large_url"`
	// A URL for the original, full-resolution version of the image. Useful for downloading.
	OriginalUrl string `json:"original_url"`
	// The order in which the image should be displayed amongst other related images.
	Index NullableInt32 `json:"index,omitempty"`
	// A brief explanation of the image.
	Caption NullableString `json:"caption,omitempty"`
	// Should the image be displayed in a gallery?
	DisplayInGallery bool `json:"display_in_gallery"`
}

// NewImage instantiates a new Image object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImage(id int32, thumbnailUrl string, largeUrl string, originalUrl string, displayInGallery bool) *Image {
	this := Image{}
	this.Id = id
	this.ThumbnailUrl = thumbnailUrl
	this.LargeUrl = largeUrl
	this.OriginalUrl = originalUrl
	this.DisplayInGallery = displayInGallery
	return &this
}

// NewImageWithDefaults instantiates a new Image object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageWithDefaults() *Image {
	this := Image{}
	return &this
}

// GetId returns the Id field value
func (o *Image) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Image) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Image) SetId(v int32) {
	o.Id = v
}

// GetThumbnailUrl returns the ThumbnailUrl field value
func (o *Image) GetThumbnailUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value
// and a boolean to check if the value has been set.
func (o *Image) GetThumbnailUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ThumbnailUrl, true
}

// SetThumbnailUrl sets field value
func (o *Image) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = v
}

// GetLargeUrl returns the LargeUrl field value
func (o *Image) GetLargeUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LargeUrl
}

// GetLargeUrlOk returns a tuple with the LargeUrl field value
// and a boolean to check if the value has been set.
func (o *Image) GetLargeUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LargeUrl, true
}

// SetLargeUrl sets field value
func (o *Image) SetLargeUrl(v string) {
	o.LargeUrl = v
}

// GetOriginalUrl returns the OriginalUrl field value
func (o *Image) GetOriginalUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalUrl
}

// GetOriginalUrlOk returns a tuple with the OriginalUrl field value
// and a boolean to check if the value has been set.
func (o *Image) GetOriginalUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginalUrl, true
}

// SetOriginalUrl sets field value
func (o *Image) SetOriginalUrl(v string) {
	o.OriginalUrl = v
}

// GetIndex returns the Index field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Image) GetIndex() int32 {
	if o == nil || o.Index.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Index.Get()
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetIndexOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Index.Get(), o.Index.IsSet()
}

// HasIndex returns a boolean if a field has been set.
func (o *Image) HasIndex() bool {
	if o != nil && o.Index.IsSet() {
		return true
	}

	return false
}

// SetIndex gets a reference to the given NullableInt32 and assigns it to the Index field.
func (o *Image) SetIndex(v int32) {
	o.Index.Set(&v)
}
// SetIndexNil sets the value for Index to be an explicit nil
func (o *Image) SetIndexNil() {
	o.Index.Set(nil)
}

// UnsetIndex ensures that no value is present for Index, not even an explicit nil
func (o *Image) UnsetIndex() {
	o.Index.Unset()
}

// GetCaption returns the Caption field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Image) GetCaption() string {
	if o == nil || o.Caption.Get() == nil {
		var ret string
		return ret
	}
	return *o.Caption.Get()
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Image) GetCaptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Caption.Get(), o.Caption.IsSet()
}

// HasCaption returns a boolean if a field has been set.
func (o *Image) HasCaption() bool {
	if o != nil && o.Caption.IsSet() {
		return true
	}

	return false
}

// SetCaption gets a reference to the given NullableString and assigns it to the Caption field.
func (o *Image) SetCaption(v string) {
	o.Caption.Set(&v)
}
// SetCaptionNil sets the value for Caption to be an explicit nil
func (o *Image) SetCaptionNil() {
	o.Caption.Set(nil)
}

// UnsetCaption ensures that no value is present for Caption, not even an explicit nil
func (o *Image) UnsetCaption() {
	o.Caption.Unset()
}

// GetDisplayInGallery returns the DisplayInGallery field value
func (o *Image) GetDisplayInGallery() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisplayInGallery
}

// GetDisplayInGalleryOk returns a tuple with the DisplayInGallery field value
// and a boolean to check if the value has been set.
func (o *Image) GetDisplayInGalleryOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayInGallery, true
}

// SetDisplayInGallery sets field value
func (o *Image) SetDisplayInGallery(v bool) {
	o.DisplayInGallery = v
}

func (o Image) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["thumbnail_url"] = o.ThumbnailUrl
	}
	if true {
		toSerialize["large_url"] = o.LargeUrl
	}
	if true {
		toSerialize["original_url"] = o.OriginalUrl
	}
	if o.Index.IsSet() {
		toSerialize["index"] = o.Index.Get()
	}
	if o.Caption.IsSet() {
		toSerialize["caption"] = o.Caption.Get()
	}
	if true {
		toSerialize["display_in_gallery"] = o.DisplayInGallery
	}
	return json.Marshal(toSerialize)
}

type NullableImage struct {
	value *Image
	isSet bool
}

func (v NullableImage) Get() *Image {
	return v.value
}

func (v *NullableImage) Set(val *Image) {
	v.value = val
	v.isSet = true
}

func (v NullableImage) IsSet() bool {
	return v.isSet
}

func (v *NullableImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImage(val *Image) *NullableImage {
	return &NullableImage{value: val, isSet: true}
}

func (v NullableImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


