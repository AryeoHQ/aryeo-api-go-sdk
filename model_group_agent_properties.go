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

// GroupAgentProperties A subset of group properties relevant to agents.
type GroupAgentProperties struct {
	// Name of brokerage.
	BrokerageName NullableString `json:"brokerage_name,omitempty"`
	// Agent's license number.
	AgentLicenseNumber NullableString `json:"agent_license_number,omitempty"`
	// Agent's profile image URL.
	AgentAvatar NullableString `json:"agent_avatar,omitempty"`
}

// NewGroupAgentProperties instantiates a new GroupAgentProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupAgentProperties() *GroupAgentProperties {
	this := GroupAgentProperties{}
	return &this
}

// NewGroupAgentPropertiesWithDefaults instantiates a new GroupAgentProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupAgentPropertiesWithDefaults() *GroupAgentProperties {
	this := GroupAgentProperties{}
	return &this
}

// GetBrokerageName returns the BrokerageName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupAgentProperties) GetBrokerageName() string {
	if o == nil || o.BrokerageName.Get() == nil {
		var ret string
		return ret
	}
	return *o.BrokerageName.Get()
}

// GetBrokerageNameOk returns a tuple with the BrokerageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupAgentProperties) GetBrokerageNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BrokerageName.Get(), o.BrokerageName.IsSet()
}

// HasBrokerageName returns a boolean if a field has been set.
func (o *GroupAgentProperties) HasBrokerageName() bool {
	if o != nil && o.BrokerageName.IsSet() {
		return true
	}

	return false
}

// SetBrokerageName gets a reference to the given NullableString and assigns it to the BrokerageName field.
func (o *GroupAgentProperties) SetBrokerageName(v string) {
	o.BrokerageName.Set(&v)
}
// SetBrokerageNameNil sets the value for BrokerageName to be an explicit nil
func (o *GroupAgentProperties) SetBrokerageNameNil() {
	o.BrokerageName.Set(nil)
}

// UnsetBrokerageName ensures that no value is present for BrokerageName, not even an explicit nil
func (o *GroupAgentProperties) UnsetBrokerageName() {
	o.BrokerageName.Unset()
}

// GetAgentLicenseNumber returns the AgentLicenseNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupAgentProperties) GetAgentLicenseNumber() string {
	if o == nil || o.AgentLicenseNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.AgentLicenseNumber.Get()
}

// GetAgentLicenseNumberOk returns a tuple with the AgentLicenseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupAgentProperties) GetAgentLicenseNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AgentLicenseNumber.Get(), o.AgentLicenseNumber.IsSet()
}

// HasAgentLicenseNumber returns a boolean if a field has been set.
func (o *GroupAgentProperties) HasAgentLicenseNumber() bool {
	if o != nil && o.AgentLicenseNumber.IsSet() {
		return true
	}

	return false
}

// SetAgentLicenseNumber gets a reference to the given NullableString and assigns it to the AgentLicenseNumber field.
func (o *GroupAgentProperties) SetAgentLicenseNumber(v string) {
	o.AgentLicenseNumber.Set(&v)
}
// SetAgentLicenseNumberNil sets the value for AgentLicenseNumber to be an explicit nil
func (o *GroupAgentProperties) SetAgentLicenseNumberNil() {
	o.AgentLicenseNumber.Set(nil)
}

// UnsetAgentLicenseNumber ensures that no value is present for AgentLicenseNumber, not even an explicit nil
func (o *GroupAgentProperties) UnsetAgentLicenseNumber() {
	o.AgentLicenseNumber.Unset()
}

// GetAgentAvatar returns the AgentAvatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupAgentProperties) GetAgentAvatar() string {
	if o == nil || o.AgentAvatar.Get() == nil {
		var ret string
		return ret
	}
	return *o.AgentAvatar.Get()
}

// GetAgentAvatarOk returns a tuple with the AgentAvatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupAgentProperties) GetAgentAvatarOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AgentAvatar.Get(), o.AgentAvatar.IsSet()
}

// HasAgentAvatar returns a boolean if a field has been set.
func (o *GroupAgentProperties) HasAgentAvatar() bool {
	if o != nil && o.AgentAvatar.IsSet() {
		return true
	}

	return false
}

// SetAgentAvatar gets a reference to the given NullableString and assigns it to the AgentAvatar field.
func (o *GroupAgentProperties) SetAgentAvatar(v string) {
	o.AgentAvatar.Set(&v)
}
// SetAgentAvatarNil sets the value for AgentAvatar to be an explicit nil
func (o *GroupAgentProperties) SetAgentAvatarNil() {
	o.AgentAvatar.Set(nil)
}

// UnsetAgentAvatar ensures that no value is present for AgentAvatar, not even an explicit nil
func (o *GroupAgentProperties) UnsetAgentAvatar() {
	o.AgentAvatar.Unset()
}

func (o GroupAgentProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BrokerageName.IsSet() {
		toSerialize["brokerage_name"] = o.BrokerageName.Get()
	}
	if o.AgentLicenseNumber.IsSet() {
		toSerialize["agent_license_number"] = o.AgentLicenseNumber.Get()
	}
	if o.AgentAvatar.IsSet() {
		toSerialize["agent_avatar"] = o.AgentAvatar.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGroupAgentProperties struct {
	value *GroupAgentProperties
	isSet bool
}

func (v NullableGroupAgentProperties) Get() *GroupAgentProperties {
	return v.value
}

func (v *NullableGroupAgentProperties) Set(val *GroupAgentProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupAgentProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupAgentProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupAgentProperties(val *GroupAgentProperties) *NullableGroupAgentProperties {
	return &NullableGroupAgentProperties{value: val, isSet: true}
}

func (v NullableGroupAgentProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupAgentProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


