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

// SocialProfiles External profile URLs for an agent or brokerage group.
type SocialProfiles struct {
	// URL for Facebook.
	FacebookProfileUrl NullableString `json:"facebook_profile_url,omitempty"`
	// URL for Instagram.
	InstagramProfileUrl NullableString `json:"instagram_profile_url,omitempty"`
	// URL for Twitter.
	TwitterProfileUrl NullableString `json:"twitter_profile_url,omitempty"`
	// URL for LinkedIn.
	LinkedinProfileUrl NullableString `json:"linkedin_profile_url,omitempty"`
	// URL for Zillow.
	ZillowProfileUrl NullableString `json:"zillow_profile_url,omitempty"`
}

// NewSocialProfiles instantiates a new SocialProfiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSocialProfiles() *SocialProfiles {
	this := SocialProfiles{}
	return &this
}

// NewSocialProfilesWithDefaults instantiates a new SocialProfiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSocialProfilesWithDefaults() *SocialProfiles {
	this := SocialProfiles{}
	return &this
}

// GetFacebookProfileUrl returns the FacebookProfileUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SocialProfiles) GetFacebookProfileUrl() string {
	if o == nil || o.FacebookProfileUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.FacebookProfileUrl.Get()
}

// GetFacebookProfileUrlOk returns a tuple with the FacebookProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SocialProfiles) GetFacebookProfileUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FacebookProfileUrl.Get(), o.FacebookProfileUrl.IsSet()
}

// HasFacebookProfileUrl returns a boolean if a field has been set.
func (o *SocialProfiles) HasFacebookProfileUrl() bool {
	if o != nil && o.FacebookProfileUrl.IsSet() {
		return true
	}

	return false
}

// SetFacebookProfileUrl gets a reference to the given NullableString and assigns it to the FacebookProfileUrl field.
func (o *SocialProfiles) SetFacebookProfileUrl(v string) {
	o.FacebookProfileUrl.Set(&v)
}
// SetFacebookProfileUrlNil sets the value for FacebookProfileUrl to be an explicit nil
func (o *SocialProfiles) SetFacebookProfileUrlNil() {
	o.FacebookProfileUrl.Set(nil)
}

// UnsetFacebookProfileUrl ensures that no value is present for FacebookProfileUrl, not even an explicit nil
func (o *SocialProfiles) UnsetFacebookProfileUrl() {
	o.FacebookProfileUrl.Unset()
}

// GetInstagramProfileUrl returns the InstagramProfileUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SocialProfiles) GetInstagramProfileUrl() string {
	if o == nil || o.InstagramProfileUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.InstagramProfileUrl.Get()
}

// GetInstagramProfileUrlOk returns a tuple with the InstagramProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SocialProfiles) GetInstagramProfileUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InstagramProfileUrl.Get(), o.InstagramProfileUrl.IsSet()
}

// HasInstagramProfileUrl returns a boolean if a field has been set.
func (o *SocialProfiles) HasInstagramProfileUrl() bool {
	if o != nil && o.InstagramProfileUrl.IsSet() {
		return true
	}

	return false
}

// SetInstagramProfileUrl gets a reference to the given NullableString and assigns it to the InstagramProfileUrl field.
func (o *SocialProfiles) SetInstagramProfileUrl(v string) {
	o.InstagramProfileUrl.Set(&v)
}
// SetInstagramProfileUrlNil sets the value for InstagramProfileUrl to be an explicit nil
func (o *SocialProfiles) SetInstagramProfileUrlNil() {
	o.InstagramProfileUrl.Set(nil)
}

// UnsetInstagramProfileUrl ensures that no value is present for InstagramProfileUrl, not even an explicit nil
func (o *SocialProfiles) UnsetInstagramProfileUrl() {
	o.InstagramProfileUrl.Unset()
}

// GetTwitterProfileUrl returns the TwitterProfileUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SocialProfiles) GetTwitterProfileUrl() string {
	if o == nil || o.TwitterProfileUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.TwitterProfileUrl.Get()
}

// GetTwitterProfileUrlOk returns a tuple with the TwitterProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SocialProfiles) GetTwitterProfileUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TwitterProfileUrl.Get(), o.TwitterProfileUrl.IsSet()
}

// HasTwitterProfileUrl returns a boolean if a field has been set.
func (o *SocialProfiles) HasTwitterProfileUrl() bool {
	if o != nil && o.TwitterProfileUrl.IsSet() {
		return true
	}

	return false
}

// SetTwitterProfileUrl gets a reference to the given NullableString and assigns it to the TwitterProfileUrl field.
func (o *SocialProfiles) SetTwitterProfileUrl(v string) {
	o.TwitterProfileUrl.Set(&v)
}
// SetTwitterProfileUrlNil sets the value for TwitterProfileUrl to be an explicit nil
func (o *SocialProfiles) SetTwitterProfileUrlNil() {
	o.TwitterProfileUrl.Set(nil)
}

// UnsetTwitterProfileUrl ensures that no value is present for TwitterProfileUrl, not even an explicit nil
func (o *SocialProfiles) UnsetTwitterProfileUrl() {
	o.TwitterProfileUrl.Unset()
}

// GetLinkedinProfileUrl returns the LinkedinProfileUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SocialProfiles) GetLinkedinProfileUrl() string {
	if o == nil || o.LinkedinProfileUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.LinkedinProfileUrl.Get()
}

// GetLinkedinProfileUrlOk returns a tuple with the LinkedinProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SocialProfiles) GetLinkedinProfileUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LinkedinProfileUrl.Get(), o.LinkedinProfileUrl.IsSet()
}

// HasLinkedinProfileUrl returns a boolean if a field has been set.
func (o *SocialProfiles) HasLinkedinProfileUrl() bool {
	if o != nil && o.LinkedinProfileUrl.IsSet() {
		return true
	}

	return false
}

// SetLinkedinProfileUrl gets a reference to the given NullableString and assigns it to the LinkedinProfileUrl field.
func (o *SocialProfiles) SetLinkedinProfileUrl(v string) {
	o.LinkedinProfileUrl.Set(&v)
}
// SetLinkedinProfileUrlNil sets the value for LinkedinProfileUrl to be an explicit nil
func (o *SocialProfiles) SetLinkedinProfileUrlNil() {
	o.LinkedinProfileUrl.Set(nil)
}

// UnsetLinkedinProfileUrl ensures that no value is present for LinkedinProfileUrl, not even an explicit nil
func (o *SocialProfiles) UnsetLinkedinProfileUrl() {
	o.LinkedinProfileUrl.Unset()
}

// GetZillowProfileUrl returns the ZillowProfileUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SocialProfiles) GetZillowProfileUrl() string {
	if o == nil || o.ZillowProfileUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ZillowProfileUrl.Get()
}

// GetZillowProfileUrlOk returns a tuple with the ZillowProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SocialProfiles) GetZillowProfileUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ZillowProfileUrl.Get(), o.ZillowProfileUrl.IsSet()
}

// HasZillowProfileUrl returns a boolean if a field has been set.
func (o *SocialProfiles) HasZillowProfileUrl() bool {
	if o != nil && o.ZillowProfileUrl.IsSet() {
		return true
	}

	return false
}

// SetZillowProfileUrl gets a reference to the given NullableString and assigns it to the ZillowProfileUrl field.
func (o *SocialProfiles) SetZillowProfileUrl(v string) {
	o.ZillowProfileUrl.Set(&v)
}
// SetZillowProfileUrlNil sets the value for ZillowProfileUrl to be an explicit nil
func (o *SocialProfiles) SetZillowProfileUrlNil() {
	o.ZillowProfileUrl.Set(nil)
}

// UnsetZillowProfileUrl ensures that no value is present for ZillowProfileUrl, not even an explicit nil
func (o *SocialProfiles) UnsetZillowProfileUrl() {
	o.ZillowProfileUrl.Unset()
}

func (o SocialProfiles) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FacebookProfileUrl.IsSet() {
		toSerialize["facebook_profile_url"] = o.FacebookProfileUrl.Get()
	}
	if o.InstagramProfileUrl.IsSet() {
		toSerialize["instagram_profile_url"] = o.InstagramProfileUrl.Get()
	}
	if o.TwitterProfileUrl.IsSet() {
		toSerialize["twitter_profile_url"] = o.TwitterProfileUrl.Get()
	}
	if o.LinkedinProfileUrl.IsSet() {
		toSerialize["linkedin_profile_url"] = o.LinkedinProfileUrl.Get()
	}
	if o.ZillowProfileUrl.IsSet() {
		toSerialize["zillow_profile_url"] = o.ZillowProfileUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSocialProfiles struct {
	value *SocialProfiles
	isSet bool
}

func (v NullableSocialProfiles) Get() *SocialProfiles {
	return v.value
}

func (v *NullableSocialProfiles) Set(val *SocialProfiles) {
	v.value = val
	v.isSet = true
}

func (v NullableSocialProfiles) IsSet() bool {
	return v.isSet
}

func (v *NullableSocialProfiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSocialProfiles(val *SocialProfiles) *NullableSocialProfiles {
	return &NullableSocialProfiles{value: val, isSet: true}
}

func (v NullableSocialProfiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSocialProfiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


