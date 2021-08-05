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

// User A record of a person on the Aryeo platform.
type User struct {
	// ID of the user. UUID Version 4.
	Id string `json:"id"`
	// Email address of the user.
	Email string `json:"email"`
	// First name of the user.
	FirstName NullableString `json:"first_name,omitempty"`
	// Last name of the user.
	LastName NullableString `json:"last_name,omitempty"`
	// A phone number represented in whichever standards specified by the user, typically ###-###-#### (separated by hyphens).
	Phone NullableString `json:"phone,omitempty"`
	// The avatar image URL of a user.
	AvatarUrl NullableString `json:"avatar_url,omitempty"`
	// Describes user's relationship (access level) to a specified group. Only returned if this resource is returned as a sub-resource of a group.
	Relationship NullableString `json:"relationship,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(id string, email string) *User {
	this := User{}
	this.Id = id
	this.Email = email
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value
func (o *User) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *User) SetId(v string) {
	o.Id = v
}

// GetEmail returns the Email field value
func (o *User) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *User) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetFirstName() string {
	if o == nil || o.FirstName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *User) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *User) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *User) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *User) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetLastName() string {
	if o == nil || o.LastName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *User) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *User) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *User) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *User) UnsetLastName() {
	o.LastName.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetPhone() string {
	if o == nil || o.Phone.Get() == nil {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetPhoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *User) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *User) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *User) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *User) UnsetPhone() {
	o.Phone.Unset()
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetAvatarUrl() string {
	if o == nil || o.AvatarUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.AvatarUrl.Get()
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetAvatarUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AvatarUrl.Get(), o.AvatarUrl.IsSet()
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *User) HasAvatarUrl() bool {
	if o != nil && o.AvatarUrl.IsSet() {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given NullableString and assigns it to the AvatarUrl field.
func (o *User) SetAvatarUrl(v string) {
	o.AvatarUrl.Set(&v)
}
// SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil
func (o *User) SetAvatarUrlNil() {
	o.AvatarUrl.Set(nil)
}

// UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
func (o *User) UnsetAvatarUrl() {
	o.AvatarUrl.Unset()
}

// GetRelationship returns the Relationship field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetRelationship() string {
	if o == nil || o.Relationship.Get() == nil {
		var ret string
		return ret
	}
	return *o.Relationship.Get()
}

// GetRelationshipOk returns a tuple with the Relationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetRelationshipOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Relationship.Get(), o.Relationship.IsSet()
}

// HasRelationship returns a boolean if a field has been set.
func (o *User) HasRelationship() bool {
	if o != nil && o.Relationship.IsSet() {
		return true
	}

	return false
}

// SetRelationship gets a reference to the given NullableString and assigns it to the Relationship field.
func (o *User) SetRelationship(v string) {
	o.Relationship.Set(&v)
}
// SetRelationshipNil sets the value for Relationship to be an explicit nil
func (o *User) SetRelationshipNil() {
	o.Relationship.Set(nil)
}

// UnsetRelationship ensures that no value is present for Relationship, not even an explicit nil
func (o *User) UnsetRelationship() {
	o.Relationship.Unset()
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	if o.AvatarUrl.IsSet() {
		toSerialize["avatar_url"] = o.AvatarUrl.Get()
	}
	if o.Relationship.IsSet() {
		toSerialize["relationship"] = o.Relationship.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


