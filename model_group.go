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

// Group A collection of users that can interact with the Aryeo platform. Permissions and properties are determined based on the group's type which can be creator, agent, or brokerage.
type Group struct {
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
	// Website.
	Website NullableString `json:"website,omitempty"`
	// Does this group represent a brokerage or an agent who belongs to a brokerage?
	IsBrokerageOrBrokerageAgent bool `json:"is_brokerage_or_brokerage_agent"`
	SocialProfiles *SocialProfiles `json:"social_profiles,omitempty"`
	AgentProperties *GroupAgentProperties `json:"agent_properties,omitempty"`
	// users
	Users *[]User `json:"users,omitempty"`
	DefaultOrderForm *OrderForm `json:"default_order_form,omitempty"`
	// An array of order forms.
	OrderForms *[]OrderForm `json:"order_forms,omitempty"`
}

// NewGroup instantiates a new Group object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroup(id string, groupType string, name string, isBrokerageOrBrokerageAgent bool) *Group {
	this := Group{}
	this.Id = id
	this.GroupType = groupType
	this.Name = name
	this.IsBrokerageOrBrokerageAgent = isBrokerageOrBrokerageAgent
	return &this
}

// NewGroupWithDefaults instantiates a new Group object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWithDefaults() *Group {
	this := Group{}
	return &this
}

// GetId returns the Id field value
func (o *Group) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Group) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Group) SetId(v string) {
	o.Id = v
}

// GetGroupType returns the GroupType field value
func (o *Group) GetGroupType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupType
}

// GetGroupTypeOk returns a tuple with the GroupType field value
// and a boolean to check if the value has been set.
func (o *Group) GetGroupTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GroupType, true
}

// SetGroupType sets field value
func (o *Group) SetGroupType(v string) {
	o.GroupType = v
}

// GetName returns the Name field value
func (o *Group) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Group) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Group) SetName(v string) {
	o.Name = v
}

// GetLogo returns the Logo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetLogo() string {
	if o == nil || o.Logo.Get() == nil {
		var ret string
		return ret
	}
	return *o.Logo.Get()
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetLogoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Logo.Get(), o.Logo.IsSet()
}

// HasLogo returns a boolean if a field has been set.
func (o *Group) HasLogo() bool {
	if o != nil && o.Logo.IsSet() {
		return true
	}

	return false
}

// SetLogo gets a reference to the given NullableString and assigns it to the Logo field.
func (o *Group) SetLogo(v string) {
	o.Logo.Set(&v)
}
// SetLogoNil sets the value for Logo to be an explicit nil
func (o *Group) SetLogoNil() {
	o.Logo.Set(nil)
}

// UnsetLogo ensures that no value is present for Logo, not even an explicit nil
func (o *Group) UnsetLogo() {
	o.Logo.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *Group) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *Group) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *Group) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *Group) UnsetEmail() {
	o.Email.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetPhone() string {
	if o == nil || o.Phone.Get() == nil {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetPhoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *Group) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *Group) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *Group) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *Group) UnsetPhone() {
	o.Phone.Unset()
}

// GetWebsite returns the Website field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetWebsite() string {
	if o == nil || o.Website.Get() == nil {
		var ret string
		return ret
	}
	return *o.Website.Get()
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetWebsiteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Website.Get(), o.Website.IsSet()
}

// HasWebsite returns a boolean if a field has been set.
func (o *Group) HasWebsite() bool {
	if o != nil && o.Website.IsSet() {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given NullableString and assigns it to the Website field.
func (o *Group) SetWebsite(v string) {
	o.Website.Set(&v)
}
// SetWebsiteNil sets the value for Website to be an explicit nil
func (o *Group) SetWebsiteNil() {
	o.Website.Set(nil)
}

// UnsetWebsite ensures that no value is present for Website, not even an explicit nil
func (o *Group) UnsetWebsite() {
	o.Website.Unset()
}

// GetIsBrokerageOrBrokerageAgent returns the IsBrokerageOrBrokerageAgent field value
func (o *Group) GetIsBrokerageOrBrokerageAgent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBrokerageOrBrokerageAgent
}

// GetIsBrokerageOrBrokerageAgentOk returns a tuple with the IsBrokerageOrBrokerageAgent field value
// and a boolean to check if the value has been set.
func (o *Group) GetIsBrokerageOrBrokerageAgentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsBrokerageOrBrokerageAgent, true
}

// SetIsBrokerageOrBrokerageAgent sets field value
func (o *Group) SetIsBrokerageOrBrokerageAgent(v bool) {
	o.IsBrokerageOrBrokerageAgent = v
}

// GetSocialProfiles returns the SocialProfiles field value if set, zero value otherwise.
func (o *Group) GetSocialProfiles() SocialProfiles {
	if o == nil || o.SocialProfiles == nil {
		var ret SocialProfiles
		return ret
	}
	return *o.SocialProfiles
}

// GetSocialProfilesOk returns a tuple with the SocialProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetSocialProfilesOk() (*SocialProfiles, bool) {
	if o == nil || o.SocialProfiles == nil {
		return nil, false
	}
	return o.SocialProfiles, true
}

// HasSocialProfiles returns a boolean if a field has been set.
func (o *Group) HasSocialProfiles() bool {
	if o != nil && o.SocialProfiles != nil {
		return true
	}

	return false
}

// SetSocialProfiles gets a reference to the given SocialProfiles and assigns it to the SocialProfiles field.
func (o *Group) SetSocialProfiles(v SocialProfiles) {
	o.SocialProfiles = &v
}

// GetAgentProperties returns the AgentProperties field value if set, zero value otherwise.
func (o *Group) GetAgentProperties() GroupAgentProperties {
	if o == nil || o.AgentProperties == nil {
		var ret GroupAgentProperties
		return ret
	}
	return *o.AgentProperties
}

// GetAgentPropertiesOk returns a tuple with the AgentProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetAgentPropertiesOk() (*GroupAgentProperties, bool) {
	if o == nil || o.AgentProperties == nil {
		return nil, false
	}
	return o.AgentProperties, true
}

// HasAgentProperties returns a boolean if a field has been set.
func (o *Group) HasAgentProperties() bool {
	if o != nil && o.AgentProperties != nil {
		return true
	}

	return false
}

// SetAgentProperties gets a reference to the given GroupAgentProperties and assigns it to the AgentProperties field.
func (o *Group) SetAgentProperties(v GroupAgentProperties) {
	o.AgentProperties = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Group) GetUsers() []User {
	if o == nil || o.Users == nil {
		var ret []User
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetUsersOk() (*[]User, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Group) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *Group) SetUsers(v []User) {
	o.Users = &v
}

// GetDefaultOrderForm returns the DefaultOrderForm field value if set, zero value otherwise.
func (o *Group) GetDefaultOrderForm() OrderForm {
	if o == nil || o.DefaultOrderForm == nil {
		var ret OrderForm
		return ret
	}
	return *o.DefaultOrderForm
}

// GetDefaultOrderFormOk returns a tuple with the DefaultOrderForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetDefaultOrderFormOk() (*OrderForm, bool) {
	if o == nil || o.DefaultOrderForm == nil {
		return nil, false
	}
	return o.DefaultOrderForm, true
}

// HasDefaultOrderForm returns a boolean if a field has been set.
func (o *Group) HasDefaultOrderForm() bool {
	if o != nil && o.DefaultOrderForm != nil {
		return true
	}

	return false
}

// SetDefaultOrderForm gets a reference to the given OrderForm and assigns it to the DefaultOrderForm field.
func (o *Group) SetDefaultOrderForm(v OrderForm) {
	o.DefaultOrderForm = &v
}

// GetOrderForms returns the OrderForms field value if set, zero value otherwise.
func (o *Group) GetOrderForms() []OrderForm {
	if o == nil || o.OrderForms == nil {
		var ret []OrderForm
		return ret
	}
	return *o.OrderForms
}

// GetOrderFormsOk returns a tuple with the OrderForms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetOrderFormsOk() (*[]OrderForm, bool) {
	if o == nil || o.OrderForms == nil {
		return nil, false
	}
	return o.OrderForms, true
}

// HasOrderForms returns a boolean if a field has been set.
func (o *Group) HasOrderForms() bool {
	if o != nil && o.OrderForms != nil {
		return true
	}

	return false
}

// SetOrderForms gets a reference to the given []OrderForm and assigns it to the OrderForms field.
func (o *Group) SetOrderForms(v []OrderForm) {
	o.OrderForms = &v
}

func (o Group) MarshalJSON() ([]byte, error) {
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
	if o.Website.IsSet() {
		toSerialize["website"] = o.Website.Get()
	}
	if true {
		toSerialize["is_brokerage_or_brokerage_agent"] = o.IsBrokerageOrBrokerageAgent
	}
	if o.SocialProfiles != nil {
		toSerialize["social_profiles"] = o.SocialProfiles
	}
	if o.AgentProperties != nil {
		toSerialize["agent_properties"] = o.AgentProperties
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.DefaultOrderForm != nil {
		toSerialize["default_order_form"] = o.DefaultOrderForm
	}
	if o.OrderForms != nil {
		toSerialize["order_forms"] = o.OrderForms
	}
	return json.Marshal(toSerialize)
}

type NullableGroup struct {
	value *Group
	isSet bool
}

func (v NullableGroup) Get() *Group {
	return v.value
}

func (v *NullableGroup) Set(val *Group) {
	v.value = val
	v.isSet = true
}

func (v NullableGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroup(val *Group) *NullableGroup {
	return &NullableGroup{value: val, isSet: true}
}

func (v NullableGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


