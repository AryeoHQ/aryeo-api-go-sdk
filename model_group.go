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

// Group A collection of users that can interact with the Aryeo platform. Permissions and properties are determined based on the group's type which can be creator, agent, or brokerage.
type Group struct {
	// ID of the group. UUID Version 4.
	Id string `json:"id"`
	// The type of the group. Can be CREATOR, AGENT, or BROKERAGE, and may dictate the attributes of the group returned.
	Type string `json:"type"`
	// The name of the group.
	Name string `json:"name"`
	// The email address of a group.
	Email NullableString `json:"email,omitempty"`
	// A phone number represented in whichever standards specified by the group, typically ###-###-#### (separated by hyphens).
	Phone NullableString `json:"phone,omitempty"`
	// The website URL of a group.
	WebsiteUrl NullableString `json:"website_url,omitempty"`
	// The logo URL of a group.
	LogoUrl NullableString `json:"logo_url,omitempty"`
	// The profile image URL of a real estate agent. Only returned if group's type is AGENT.
	AvatarUrl NullableString `json:"avatar_url,omitempty"`
	// The name of the brokerage or team of a real estate agent. Only returned if group's type is AGENT.
	OfficeName NullableString `json:"office_name,omitempty"`
	// The license number of a real estate agent. Only returned if group's type is AGENT.
	LicenseNumber NullableString `json:"license_number,omitempty"`
	SocialProfiles *SocialProfiles `json:"social_profiles,omitempty"`
	DefaultOrderForm *OrderForm `json:"default_order_form,omitempty"`
	// An array of order forms a vendor group provides for placing orders. Only returned if group's type is CREATOR. 
	OrderForms []OrderForm `json:"order_forms,omitempty"`
	Owner *User `json:"owner,omitempty"`
	// The Aryeo users associated with this group.
	Users []User `json:"users,omitempty"`
	// Does this group represent a brokerage or an agent who belongs to a brokerage?
	IsBrokerageOrBrokerageAgent bool `json:"is_brokerage_or_brokerage_agent"`
}

// NewGroup instantiates a new Group object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroup(id string, type_ string, name string, isBrokerageOrBrokerageAgent bool) *Group {
	this := Group{}
	this.Id = id
	this.Type = type_
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

// GetType returns the Type field value
func (o *Group) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Group) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Group) SetType(v string) {
	o.Type = v
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

// GetWebsiteUrl returns the WebsiteUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetWebsiteUrl() string {
	if o == nil || o.WebsiteUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebsiteUrl.Get()
}

// GetWebsiteUrlOk returns a tuple with the WebsiteUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetWebsiteUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WebsiteUrl.Get(), o.WebsiteUrl.IsSet()
}

// HasWebsiteUrl returns a boolean if a field has been set.
func (o *Group) HasWebsiteUrl() bool {
	if o != nil && o.WebsiteUrl.IsSet() {
		return true
	}

	return false
}

// SetWebsiteUrl gets a reference to the given NullableString and assigns it to the WebsiteUrl field.
func (o *Group) SetWebsiteUrl(v string) {
	o.WebsiteUrl.Set(&v)
}
// SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil
func (o *Group) SetWebsiteUrlNil() {
	o.WebsiteUrl.Set(nil)
}

// UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
func (o *Group) UnsetWebsiteUrl() {
	o.WebsiteUrl.Unset()
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetLogoUrl() string {
	if o == nil || o.LogoUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.LogoUrl.Get()
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetLogoUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LogoUrl.Get(), o.LogoUrl.IsSet()
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *Group) HasLogoUrl() bool {
	if o != nil && o.LogoUrl.IsSet() {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given NullableString and assigns it to the LogoUrl field.
func (o *Group) SetLogoUrl(v string) {
	o.LogoUrl.Set(&v)
}
// SetLogoUrlNil sets the value for LogoUrl to be an explicit nil
func (o *Group) SetLogoUrlNil() {
	o.LogoUrl.Set(nil)
}

// UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
func (o *Group) UnsetLogoUrl() {
	o.LogoUrl.Unset()
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetAvatarUrl() string {
	if o == nil || o.AvatarUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.AvatarUrl.Get()
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetAvatarUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AvatarUrl.Get(), o.AvatarUrl.IsSet()
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *Group) HasAvatarUrl() bool {
	if o != nil && o.AvatarUrl.IsSet() {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given NullableString and assigns it to the AvatarUrl field.
func (o *Group) SetAvatarUrl(v string) {
	o.AvatarUrl.Set(&v)
}
// SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil
func (o *Group) SetAvatarUrlNil() {
	o.AvatarUrl.Set(nil)
}

// UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
func (o *Group) UnsetAvatarUrl() {
	o.AvatarUrl.Unset()
}

// GetOfficeName returns the OfficeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetOfficeName() string {
	if o == nil || o.OfficeName.Get() == nil {
		var ret string
		return ret
	}
	return *o.OfficeName.Get()
}

// GetOfficeNameOk returns a tuple with the OfficeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetOfficeNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OfficeName.Get(), o.OfficeName.IsSet()
}

// HasOfficeName returns a boolean if a field has been set.
func (o *Group) HasOfficeName() bool {
	if o != nil && o.OfficeName.IsSet() {
		return true
	}

	return false
}

// SetOfficeName gets a reference to the given NullableString and assigns it to the OfficeName field.
func (o *Group) SetOfficeName(v string) {
	o.OfficeName.Set(&v)
}
// SetOfficeNameNil sets the value for OfficeName to be an explicit nil
func (o *Group) SetOfficeNameNil() {
	o.OfficeName.Set(nil)
}

// UnsetOfficeName ensures that no value is present for OfficeName, not even an explicit nil
func (o *Group) UnsetOfficeName() {
	o.OfficeName.Unset()
}

// GetLicenseNumber returns the LicenseNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetLicenseNumber() string {
	if o == nil || o.LicenseNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.LicenseNumber.Get()
}

// GetLicenseNumberOk returns a tuple with the LicenseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetLicenseNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LicenseNumber.Get(), o.LicenseNumber.IsSet()
}

// HasLicenseNumber returns a boolean if a field has been set.
func (o *Group) HasLicenseNumber() bool {
	if o != nil && o.LicenseNumber.IsSet() {
		return true
	}

	return false
}

// SetLicenseNumber gets a reference to the given NullableString and assigns it to the LicenseNumber field.
func (o *Group) SetLicenseNumber(v string) {
	o.LicenseNumber.Set(&v)
}
// SetLicenseNumberNil sets the value for LicenseNumber to be an explicit nil
func (o *Group) SetLicenseNumberNil() {
	o.LicenseNumber.Set(nil)
}

// UnsetLicenseNumber ensures that no value is present for LicenseNumber, not even an explicit nil
func (o *Group) UnsetLicenseNumber() {
	o.LicenseNumber.Unset()
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

// GetOrderForms returns the OrderForms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetOrderForms() []OrderForm {
	if o == nil  {
		var ret []OrderForm
		return ret
	}
	return o.OrderForms
}

// GetOrderFormsOk returns a tuple with the OrderForms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetOrderFormsOk() (*[]OrderForm, bool) {
	if o == nil || o.OrderForms == nil {
		return nil, false
	}
	return &o.OrderForms, true
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
	o.OrderForms = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Group) GetOwner() User {
	if o == nil || o.Owner == nil {
		var ret User
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetOwnerOk() (*User, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Group) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given User and assigns it to the Owner field.
func (o *Group) SetOwner(v User) {
	o.Owner = &v
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetUsers() []User {
	if o == nil  {
		var ret []User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetUsersOk() (*[]User, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
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
	o.Users = v
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

func (o Group) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	if o.WebsiteUrl.IsSet() {
		toSerialize["website_url"] = o.WebsiteUrl.Get()
	}
	if o.LogoUrl.IsSet() {
		toSerialize["logo_url"] = o.LogoUrl.Get()
	}
	if o.AvatarUrl.IsSet() {
		toSerialize["avatar_url"] = o.AvatarUrl.Get()
	}
	if o.OfficeName.IsSet() {
		toSerialize["office_name"] = o.OfficeName.Get()
	}
	if o.LicenseNumber.IsSet() {
		toSerialize["license_number"] = o.LicenseNumber.Get()
	}
	if o.SocialProfiles != nil {
		toSerialize["social_profiles"] = o.SocialProfiles
	}
	if o.DefaultOrderForm != nil {
		toSerialize["default_order_form"] = o.DefaultOrderForm
	}
	if o.OrderForms != nil {
		toSerialize["order_forms"] = o.OrderForms
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if true {
		toSerialize["is_brokerage_or_brokerage_agent"] = o.IsBrokerageOrBrokerageAgent
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


