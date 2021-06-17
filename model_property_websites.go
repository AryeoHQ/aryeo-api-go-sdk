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

// PropertyWebsites Websites that displays information about a property.
type PropertyWebsites struct {
	// URL for website.
	BrandedUrl string `json:"branded_url"`
	// URL for website.
	UnbrandedUrl string `json:"unbranded_url"`
	// ID for property website
	Id int32 `json:"id"`
}

// NewPropertyWebsites instantiates a new PropertyWebsites object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyWebsites(brandedUrl string, unbrandedUrl string, id int32) *PropertyWebsites {
	this := PropertyWebsites{}
	this.BrandedUrl = brandedUrl
	this.UnbrandedUrl = unbrandedUrl
	this.Id = id
	return &this
}

// NewPropertyWebsitesWithDefaults instantiates a new PropertyWebsites object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyWebsitesWithDefaults() *PropertyWebsites {
	this := PropertyWebsites{}
	return &this
}

// GetBrandedUrl returns the BrandedUrl field value
func (o *PropertyWebsites) GetBrandedUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandedUrl
}

// GetBrandedUrlOk returns a tuple with the BrandedUrl field value
// and a boolean to check if the value has been set.
func (o *PropertyWebsites) GetBrandedUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BrandedUrl, true
}

// SetBrandedUrl sets field value
func (o *PropertyWebsites) SetBrandedUrl(v string) {
	o.BrandedUrl = v
}

// GetUnbrandedUrl returns the UnbrandedUrl field value
func (o *PropertyWebsites) GetUnbrandedUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnbrandedUrl
}

// GetUnbrandedUrlOk returns a tuple with the UnbrandedUrl field value
// and a boolean to check if the value has been set.
func (o *PropertyWebsites) GetUnbrandedUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UnbrandedUrl, true
}

// SetUnbrandedUrl sets field value
func (o *PropertyWebsites) SetUnbrandedUrl(v string) {
	o.UnbrandedUrl = v
}

// GetId returns the Id field value
func (o *PropertyWebsites) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PropertyWebsites) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PropertyWebsites) SetId(v int32) {
	o.Id = v
}

func (o PropertyWebsites) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["branded_url"] = o.BrandedUrl
	}
	if true {
		toSerialize["unbranded_url"] = o.UnbrandedUrl
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyWebsites struct {
	value *PropertyWebsites
	isSet bool
}

func (v NullablePropertyWebsites) Get() *PropertyWebsites {
	return v.value
}

func (v *NullablePropertyWebsites) Set(val *PropertyWebsites) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyWebsites) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyWebsites) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyWebsites(val *PropertyWebsites) *NullablePropertyWebsites {
	return &NullablePropertyWebsites{value: val, isSet: true}
}

func (v NullablePropertyWebsites) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyWebsites) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


