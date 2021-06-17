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

// FloorPlan A scale diagram of the arrangement of a building.
type FloorPlan struct {
	// ID of the floor plan.
	Id int32 `json:"id"`
	// A URL for a thumbnail-sized version of the floor plan.
	ThumbnailUrl string `json:"thumbnail_url"`
	// A URL for a large version of the floor plan.
	LargeUrl string `json:"large_url"`
	// A URL for the original, full-resolution version of the floor plan. Useful for downloading.
	OriginalUrl string `json:"original_url"`
	// The title (or caption) of the floor plan.
	Title NullableString `json:"title,omitempty"`
	// Index order position of the floor plan.
	Index NullableInt32 `json:"index,omitempty"`
}

// NewFloorPlan instantiates a new FloorPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFloorPlan(id int32, thumbnailUrl string, largeUrl string, originalUrl string) *FloorPlan {
	this := FloorPlan{}
	this.Id = id
	this.ThumbnailUrl = thumbnailUrl
	this.LargeUrl = largeUrl
	this.OriginalUrl = originalUrl
	return &this
}

// NewFloorPlanWithDefaults instantiates a new FloorPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFloorPlanWithDefaults() *FloorPlan {
	this := FloorPlan{}
	return &this
}

// GetId returns the Id field value
func (o *FloorPlan) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FloorPlan) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FloorPlan) SetId(v int32) {
	o.Id = v
}

// GetThumbnailUrl returns the ThumbnailUrl field value
func (o *FloorPlan) GetThumbnailUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value
// and a boolean to check if the value has been set.
func (o *FloorPlan) GetThumbnailUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ThumbnailUrl, true
}

// SetThumbnailUrl sets field value
func (o *FloorPlan) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = v
}

// GetLargeUrl returns the LargeUrl field value
func (o *FloorPlan) GetLargeUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LargeUrl
}

// GetLargeUrlOk returns a tuple with the LargeUrl field value
// and a boolean to check if the value has been set.
func (o *FloorPlan) GetLargeUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LargeUrl, true
}

// SetLargeUrl sets field value
func (o *FloorPlan) SetLargeUrl(v string) {
	o.LargeUrl = v
}

// GetOriginalUrl returns the OriginalUrl field value
func (o *FloorPlan) GetOriginalUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalUrl
}

// GetOriginalUrlOk returns a tuple with the OriginalUrl field value
// and a boolean to check if the value has been set.
func (o *FloorPlan) GetOriginalUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginalUrl, true
}

// SetOriginalUrl sets field value
func (o *FloorPlan) SetOriginalUrl(v string) {
	o.OriginalUrl = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FloorPlan) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FloorPlan) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *FloorPlan) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *FloorPlan) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *FloorPlan) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *FloorPlan) UnsetTitle() {
	o.Title.Unset()
}

// GetIndex returns the Index field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FloorPlan) GetIndex() int32 {
	if o == nil || o.Index.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Index.Get()
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FloorPlan) GetIndexOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Index.Get(), o.Index.IsSet()
}

// HasIndex returns a boolean if a field has been set.
func (o *FloorPlan) HasIndex() bool {
	if o != nil && o.Index.IsSet() {
		return true
	}

	return false
}

// SetIndex gets a reference to the given NullableInt32 and assigns it to the Index field.
func (o *FloorPlan) SetIndex(v int32) {
	o.Index.Set(&v)
}
// SetIndexNil sets the value for Index to be an explicit nil
func (o *FloorPlan) SetIndexNil() {
	o.Index.Set(nil)
}

// UnsetIndex ensures that no value is present for Index, not even an explicit nil
func (o *FloorPlan) UnsetIndex() {
	o.Index.Unset()
}

func (o FloorPlan) MarshalJSON() ([]byte, error) {
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
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Index.IsSet() {
		toSerialize["index"] = o.Index.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableFloorPlan struct {
	value *FloorPlan
	isSet bool
}

func (v NullableFloorPlan) Get() *FloorPlan {
	return v.value
}

func (v *NullableFloorPlan) Set(val *FloorPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableFloorPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableFloorPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFloorPlan(val *FloorPlan) *NullableFloorPlan {
	return &NullableFloorPlan{value: val, isSet: true}
}

func (v NullableFloorPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFloorPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


