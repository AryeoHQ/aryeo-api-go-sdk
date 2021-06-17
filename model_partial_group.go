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

// PartialGroup A collection of users that can interact with the Aryeo platform. Permissions and properties are determined based on the type which can be creator, agent, or brokerage.
type PartialGroup struct {
	// ID of the group.
	Id string `json:"id"`
	// The type of group.
	GroupType string `json:"group_type"`
	// The name of the group.
	Name string `json:"name"`
	// Group logo.
	Logo NullableString `json:"logo,omitempty"`
	// Email.
	Email NullableString `json:"email,omitempty"`
	// Phone number.
	Phone NullableString `json:"phone,omitempty"`
	// Does this group represent a brokerage or an agent who belongs to a brokerage?
	IsBrokerageOrBrokerageAgent bool `json:"is_brokerage_or_brokerage_agent"`
}

// NewPartialGroup instantiates a new PartialGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialGroup(id string, groupType string, name string, isBrokerageOrBrokerageAgent bool) *PartialGroup {
	this := PartialGroup{}
	this.Id = id
	this.GroupType = groupType
	this.Name = name
	this.IsBrokerageOrBrokerageAgent = isBrokerageOrBrokerageAgent
	return &this
}

// NewPartialGroupWithDefaults instantiates a new PartialGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialGroupWithDefaults() *PartialGroup {
	this := PartialGroup{}
	return &this
}

// GetId returns the Id field value
func (o *PartialGroup) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PartialGroup) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PartialGroup) SetId(v string) {
	o.Id = v
}

// GetGroupType returns the GroupType field value
func (o *PartialGroup) GetGroupType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupType
}

// GetGroupTypeOk returns a tuple with the GroupType field value
// and a boolean to check if the value has been set.
func (o *PartialGroup) GetGroupTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GroupType, true
}

// SetGroupType sets field value
func (o *PartialGroup) SetGroupType(v string) {
	o.GroupType = v
}

// GetName returns the Name field value
func (o *PartialGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PartialGroup) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PartialGroup) SetName(v string) {
	o.Name = v
}

// GetLogo returns the Logo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialGroup) GetLogo() string {
	if o == nil || o.Logo.Get() == nil {
		var ret string
		return ret
	}
	return *o.Logo.Get()
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialGroup) GetLogoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Logo.Get(), o.Logo.IsSet()
}

// HasLogo returns a boolean if a field has been set.
func (o *PartialGroup) HasLogo() bool {
	if o != nil && o.Logo.IsSet() {
		return true
	}

	return false
}

// SetLogo gets a reference to the given NullableString and assigns it to the Logo field.
func (o *PartialGroup) SetLogo(v string) {
	o.Logo.Set(&v)
}
// SetLogoNil sets the value for Logo to be an explicit nil
func (o *PartialGroup) SetLogoNil() {
	o.Logo.Set(nil)
}

// UnsetLogo ensures that no value is present for Logo, not even an explicit nil
func (o *PartialGroup) UnsetLogo() {
	o.Logo.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialGroup) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialGroup) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *PartialGroup) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *PartialGroup) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *PartialGroup) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *PartialGroup) UnsetEmail() {
	o.Email.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialGroup) GetPhone() string {
	if o == nil || o.Phone.Get() == nil {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialGroup) GetPhoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *PartialGroup) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *PartialGroup) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *PartialGroup) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *PartialGroup) UnsetPhone() {
	o.Phone.Unset()
}

// GetIsBrokerageOrBrokerageAgent returns the IsBrokerageOrBrokerageAgent field value
func (o *PartialGroup) GetIsBrokerageOrBrokerageAgent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBrokerageOrBrokerageAgent
}

// GetIsBrokerageOrBrokerageAgentOk returns a tuple with the IsBrokerageOrBrokerageAgent field value
// and a boolean to check if the value has been set.
func (o *PartialGroup) GetIsBrokerageOrBrokerageAgentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsBrokerageOrBrokerageAgent, true
}

// SetIsBrokerageOrBrokerageAgent sets field value
func (o *PartialGroup) SetIsBrokerageOrBrokerageAgent(v bool) {
	o.IsBrokerageOrBrokerageAgent = v
}

func (o PartialGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["group_type"] = o.GroupType
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Logo.IsSet() {
		toSerialize["logo"] = o.Logo.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	if true {
		toSerialize["is_brokerage_or_brokerage_agent"] = o.IsBrokerageOrBrokerageAgent
	}
	return json.Marshal(toSerialize)
}

type NullablePartialGroup struct {
	value *PartialGroup
	isSet bool
}

func (v NullablePartialGroup) Get() *PartialGroup {
	return v.value
}

func (v *NullablePartialGroup) Set(val *PartialGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialGroup(val *PartialGroup) *NullablePartialGroup {
	return &NullablePartialGroup{value: val, isSet: true}
}

func (v NullablePartialGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


