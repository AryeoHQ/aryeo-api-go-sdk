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

// PropertyWebsite Website (in branded and unbranded versions) that displays information about a property.
type PropertyWebsite struct {
	// ID of the website.
	Id string `json:"id"`
	// URL for branded version of website.
	BrandedUrl string `json:"branded_url"`
	// URL for unbranded version of website.
	UnbrandedUrl string `json:"unbranded_url"`
}

// NewPropertyWebsite instantiates a new PropertyWebsite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyWebsite(id string, brandedUrl string, unbrandedUrl string) *PropertyWebsite {
	this := PropertyWebsite{}
	this.Id = id
	this.BrandedUrl = brandedUrl
	this.UnbrandedUrl = unbrandedUrl
	return &this
}

// NewPropertyWebsiteWithDefaults instantiates a new PropertyWebsite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyWebsiteWithDefaults() *PropertyWebsite {
	this := PropertyWebsite{}
	return &this
}

// GetId returns the Id field value
func (o *PropertyWebsite) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PropertyWebsite) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PropertyWebsite) SetId(v string) {
	o.Id = v
}

// GetBrandedUrl returns the BrandedUrl field value
func (o *PropertyWebsite) GetBrandedUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandedUrl
}

// GetBrandedUrlOk returns a tuple with the BrandedUrl field value
// and a boolean to check if the value has been set.
func (o *PropertyWebsite) GetBrandedUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BrandedUrl, true
}

// SetBrandedUrl sets field value
func (o *PropertyWebsite) SetBrandedUrl(v string) {
	o.BrandedUrl = v
}

// GetUnbrandedUrl returns the UnbrandedUrl field value
func (o *PropertyWebsite) GetUnbrandedUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnbrandedUrl
}

// GetUnbrandedUrlOk returns a tuple with the UnbrandedUrl field value
// and a boolean to check if the value has been set.
func (o *PropertyWebsite) GetUnbrandedUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UnbrandedUrl, true
}

// SetUnbrandedUrl sets field value
func (o *PropertyWebsite) SetUnbrandedUrl(v string) {
	o.UnbrandedUrl = v
}

func (o PropertyWebsite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["branded_url"] = o.BrandedUrl
	}
	if true {
		toSerialize["unbranded_url"] = o.UnbrandedUrl
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyWebsite struct {
	value *PropertyWebsite
	isSet bool
}

func (v NullablePropertyWebsite) Get() *PropertyWebsite {
	return v.value
}

func (v *NullablePropertyWebsite) Set(val *PropertyWebsite) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyWebsite) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyWebsite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyWebsite(val *PropertyWebsite) *NullablePropertyWebsite {
	return &NullablePropertyWebsite{value: val, isSet: true}
}

func (v NullablePropertyWebsite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyWebsite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


