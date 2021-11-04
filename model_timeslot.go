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

// Timeslot A bookable time range, including users that are available.
type Timeslot struct {
	// Start time of the available slot
	StartAt string `json:"start_at"`
	// End time of the available slot
	EndAt string `json:"end_at"`
	Users *[]User `json:"users,omitempty"`
}

// NewTimeslot instantiates a new Timeslot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeslot(startAt string, endAt string) *Timeslot {
	this := Timeslot{}
	this.StartAt = startAt
	this.EndAt = endAt
	return &this
}

// NewTimeslotWithDefaults instantiates a new Timeslot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeslotWithDefaults() *Timeslot {
	this := Timeslot{}
	return &this
}

// GetStartAt returns the StartAt field value
func (o *Timeslot) GetStartAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value
// and a boolean to check if the value has been set.
func (o *Timeslot) GetStartAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartAt, true
}

// SetStartAt sets field value
func (o *Timeslot) SetStartAt(v string) {
	o.StartAt = v
}

// GetEndAt returns the EndAt field value
func (o *Timeslot) GetEndAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndAt
}

// GetEndAtOk returns a tuple with the EndAt field value
// and a boolean to check if the value has been set.
func (o *Timeslot) GetEndAtOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndAt, true
}

// SetEndAt sets field value
func (o *Timeslot) SetEndAt(v string) {
	o.EndAt = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *Timeslot) GetUsers() []User {
	if o == nil || o.Users == nil {
		var ret []User
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeslot) GetUsersOk() (*[]User, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Timeslot) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *Timeslot) SetUsers(v []User) {
	o.Users = &v
}

func (o Timeslot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["start_at"] = o.StartAt
	}
	if true {
		toSerialize["end_at"] = o.EndAt
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableTimeslot struct {
	value *Timeslot
	isSet bool
}

func (v NullableTimeslot) Get() *Timeslot {
	return v.value
}

func (v *NullableTimeslot) Set(val *Timeslot) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeslot) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeslot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeslot(val *Timeslot) *NullableTimeslot {
	return &NullableTimeslot{value: val, isSet: true}
}

func (v NullableTimeslot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeslot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


