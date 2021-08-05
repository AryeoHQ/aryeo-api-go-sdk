/*
 * Aryeo
 *
 *
 * API version: 2021-06-17
 * Contact: jarrod@aryeo.com
 */

package aryeo

import (
	"encoding/json"
)

// InteractiveContent A 3D virtual tour.
type InteractiveContent struct {
	// ID of the content. UUID Version 4.
	Id string `json:"id"`
	// Is the content branded, unbranded, or both?
	DisplayType string `json:"display_type"`
	// The type of interactive content.
	ContentType string `json:"content_type"`
	// URL for the content.
	Url string `json:"url"`
	// A URL for a thumbnail-sized preview of the content.
	ThumbnailUrl NullableString `json:"thumbnail_url,omitempty"`
}

// NewInteractiveContent instantiates a new InteractiveContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInteractiveContent(id string, displayType string, contentType string, url string) *InteractiveContent {
	this := InteractiveContent{}
	this.Id = id
	this.DisplayType = displayType
	this.ContentType = contentType
	this.Url = url
	return &this
}

// NewInteractiveContentWithDefaults instantiates a new InteractiveContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInteractiveContentWithDefaults() *InteractiveContent {
	this := InteractiveContent{}
	return &this
}

// GetId returns the Id field value
func (o *InteractiveContent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InteractiveContent) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InteractiveContent) SetId(v string) {
	o.Id = v
}

// GetDisplayType returns the DisplayType field value
func (o *InteractiveContent) GetDisplayType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayType
}

// GetDisplayTypeOk returns a tuple with the DisplayType field value
// and a boolean to check if the value has been set.
func (o *InteractiveContent) GetDisplayTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayType, true
}

// SetDisplayType sets field value
func (o *InteractiveContent) SetDisplayType(v string) {
	o.DisplayType = v
}

// GetContentType returns the ContentType field value
func (o *InteractiveContent) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *InteractiveContent) GetContentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *InteractiveContent) SetContentType(v string) {
	o.ContentType = v
}

// GetUrl returns the Url field value
func (o *InteractiveContent) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *InteractiveContent) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *InteractiveContent) SetUrl(v string) {
	o.Url = v
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InteractiveContent) GetThumbnailUrl() string {
	if o == nil || o.ThumbnailUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl.Get()
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InteractiveContent) GetThumbnailUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ThumbnailUrl.Get(), o.ThumbnailUrl.IsSet()
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *InteractiveContent) HasThumbnailUrl() bool {
	if o != nil && o.ThumbnailUrl.IsSet() {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given NullableString and assigns it to the ThumbnailUrl field.
func (o *InteractiveContent) SetThumbnailUrl(v string) {
	o.ThumbnailUrl.Set(&v)
}
// SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil
func (o *InteractiveContent) SetThumbnailUrlNil() {
	o.ThumbnailUrl.Set(nil)
}

// UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
func (o *InteractiveContent) UnsetThumbnailUrl() {
	o.ThumbnailUrl.Unset()
}

func (o InteractiveContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["display_type"] = o.DisplayType
	}
	if true {
		toSerialize["content_type"] = o.ContentType
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.ThumbnailUrl.IsSet() {
		toSerialize["thumbnail_url"] = o.ThumbnailUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInteractiveContent struct {
	value *InteractiveContent
	isSet bool
}

func (v NullableInteractiveContent) Get() *InteractiveContent {
	return v.value
}

func (v *NullableInteractiveContent) Set(val *InteractiveContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInteractiveContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInteractiveContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInteractiveContent(val *InteractiveContent) *NullableInteractiveContent {
	return &NullableInteractiveContent{value: val, isSet: true}
}

func (v NullableInteractiveContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInteractiveContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


