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
	"time"
)

// UnconfirmedAppointment An unconfirmed appointment.
type UnconfirmedAppointment struct {
	// The ID of the appointment.
	Id string `json:"id"`
	// The title of the appointment.
	Title NullableString `json:"title,omitempty"`
	// The multi-line description of the appointment.
	Description NullableString `json:"description,omitempty"`
	Order *Order `json:"order,omitempty"`
	// Users attached to the appointment.
	Users []User `json:"users,omitempty"`
	// The type of preferred scheduling information provided by a customer to aid in scheduling this appointment.
	PreferenceType NullableString `json:"preference_type,omitempty"`
	// A preferred date and time (ISO 8601 format) for when the appointment could start. Only returned if unconfirmed appointment's preference type is TIME. 
	PreferredStartAt NullableTime `json:"preferred_start_at,omitempty"`
	// A preferred date (ISO 8601 format) for when the appointment could start. Only returned if unconfirmed appointment's preference type is TIME_OF_DAY. 
	PreferredStartAtDay NullableString `json:"preferred_start_at_day,omitempty"`
	// A preferred time of day for when the appointment could start. Only returned if unconfirmed appointment's preference type is TIME_OF_DAY. 
	PreferredStartAtTimeOfDay NullableString `json:"preferred_start_at_time_of_day,omitempty"`
	// The estimated length of the appointment in minutes, if available.
	Duration NullableInt32 `json:"duration,omitempty"`
}

// NewUnconfirmedAppointment instantiates a new UnconfirmedAppointment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnconfirmedAppointment(id string) *UnconfirmedAppointment {
	this := UnconfirmedAppointment{}
	this.Id = id
	return &this
}

// NewUnconfirmedAppointmentWithDefaults instantiates a new UnconfirmedAppointment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnconfirmedAppointmentWithDefaults() *UnconfirmedAppointment {
	this := UnconfirmedAppointment{}
	return &this
}

// GetId returns the Id field value
func (o *UnconfirmedAppointment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UnconfirmedAppointment) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UnconfirmedAppointment) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnconfirmedAppointment) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnconfirmedAppointment) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *UnconfirmedAppointment) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *UnconfirmedAppointment) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *UnconfirmedAppointment) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *UnconfirmedAppointment) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnconfirmedAppointment) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnconfirmedAppointment) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UnconfirmedAppointment) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UnconfirmedAppointment) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UnconfirmedAppointment) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UnconfirmedAppointment) UnsetDescription() {
	o.Description.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *UnconfirmedAppointment) GetOrder() Order {
	if o == nil || o.Order == nil {
		var ret Order
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnconfirmedAppointment) GetOrderOk() (*Order, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *UnconfirmedAppointment) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given Order and assigns it to the Order field.
func (o *UnconfirmedAppointment) SetOrder(v Order) {
	o.Order = &v
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnconfirmedAppointment) GetUsers() []User {
	if o == nil  {
		var ret []User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnconfirmedAppointment) GetUsersOk() (*[]User, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *UnconfirmedAppointment) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *UnconfirmedAppointment) SetUsers(v []User) {
	o.Users = v
}

// GetPreferenceType returns the PreferenceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnconfirmedAppointment) GetPreferenceType() string {
	if o == nil || o.PreferenceType.Get() == nil {
		var ret string
		return ret
	}
	return *o.PreferenceType.Get()
}

// GetPreferenceTypeOk returns a tuple with the PreferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnconfirmedAppointment) GetPreferenceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreferenceType.Get(), o.PreferenceType.IsSet()
}

// HasPreferenceType returns a boolean if a field has been set.
func (o *UnconfirmedAppointment) HasPreferenceType() bool {
	if o != nil && o.PreferenceType.IsSet() {
		return true
	}

	return false
}

// SetPreferenceType gets a reference to the given NullableString and assigns it to the PreferenceType field.
func (o *UnconfirmedAppointment) SetPreferenceType(v string) {
	o.PreferenceType.Set(&v)
}
// SetPreferenceTypeNil sets the value for PreferenceType to be an explicit nil
func (o *UnconfirmedAppointment) SetPreferenceTypeNil() {
	o.PreferenceType.Set(nil)
}

// UnsetPreferenceType ensures that no value is present for PreferenceType, not even an explicit nil
func (o *UnconfirmedAppointment) UnsetPreferenceType() {
	o.PreferenceType.Unset()
}

// GetPreferredStartAt returns the PreferredStartAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnconfirmedAppointment) GetPreferredStartAt() time.Time {
	if o == nil || o.PreferredStartAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.PreferredStartAt.Get()
}

// GetPreferredStartAtOk returns a tuple with the PreferredStartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnconfirmedAppointment) GetPreferredStartAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreferredStartAt.Get(), o.PreferredStartAt.IsSet()
}

// HasPreferredStartAt returns a boolean if a field has been set.
func (o *UnconfirmedAppointment) HasPreferredStartAt() bool {
	if o != nil && o.PreferredStartAt.IsSet() {
		return true
	}

	return false
}

// SetPreferredStartAt gets a reference to the given NullableTime and assigns it to the PreferredStartAt field.
func (o *UnconfirmedAppointment) SetPreferredStartAt(v time.Time) {
	o.PreferredStartAt.Set(&v)
}
// SetPreferredStartAtNil sets the value for PreferredStartAt to be an explicit nil
func (o *UnconfirmedAppointment) SetPreferredStartAtNil() {
	o.PreferredStartAt.Set(nil)
}

// UnsetPreferredStartAt ensures that no value is present for PreferredStartAt, not even an explicit nil
func (o *UnconfirmedAppointment) UnsetPreferredStartAt() {
	o.PreferredStartAt.Unset()
}

// GetPreferredStartAtDay returns the PreferredStartAtDay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnconfirmedAppointment) GetPreferredStartAtDay() string {
	if o == nil || o.PreferredStartAtDay.Get() == nil {
		var ret string
		return ret
	}
	return *o.PreferredStartAtDay.Get()
}

// GetPreferredStartAtDayOk returns a tuple with the PreferredStartAtDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnconfirmedAppointment) GetPreferredStartAtDayOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreferredStartAtDay.Get(), o.PreferredStartAtDay.IsSet()
}

// HasPreferredStartAtDay returns a boolean if a field has been set.
func (o *UnconfirmedAppointment) HasPreferredStartAtDay() bool {
	if o != nil && o.PreferredStartAtDay.IsSet() {
		return true
	}

	return false
}

// SetPreferredStartAtDay gets a reference to the given NullableString and assigns it to the PreferredStartAtDay field.
func (o *UnconfirmedAppointment) SetPreferredStartAtDay(v string) {
	o.PreferredStartAtDay.Set(&v)
}
// SetPreferredStartAtDayNil sets the value for PreferredStartAtDay to be an explicit nil
func (o *UnconfirmedAppointment) SetPreferredStartAtDayNil() {
	o.PreferredStartAtDay.Set(nil)
}

// UnsetPreferredStartAtDay ensures that no value is present for PreferredStartAtDay, not even an explicit nil
func (o *UnconfirmedAppointment) UnsetPreferredStartAtDay() {
	o.PreferredStartAtDay.Unset()
}

// GetPreferredStartAtTimeOfDay returns the PreferredStartAtTimeOfDay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnconfirmedAppointment) GetPreferredStartAtTimeOfDay() string {
	if o == nil || o.PreferredStartAtTimeOfDay.Get() == nil {
		var ret string
		return ret
	}
	return *o.PreferredStartAtTimeOfDay.Get()
}

// GetPreferredStartAtTimeOfDayOk returns a tuple with the PreferredStartAtTimeOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnconfirmedAppointment) GetPreferredStartAtTimeOfDayOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PreferredStartAtTimeOfDay.Get(), o.PreferredStartAtTimeOfDay.IsSet()
}

// HasPreferredStartAtTimeOfDay returns a boolean if a field has been set.
func (o *UnconfirmedAppointment) HasPreferredStartAtTimeOfDay() bool {
	if o != nil && o.PreferredStartAtTimeOfDay.IsSet() {
		return true
	}

	return false
}

// SetPreferredStartAtTimeOfDay gets a reference to the given NullableString and assigns it to the PreferredStartAtTimeOfDay field.
func (o *UnconfirmedAppointment) SetPreferredStartAtTimeOfDay(v string) {
	o.PreferredStartAtTimeOfDay.Set(&v)
}
// SetPreferredStartAtTimeOfDayNil sets the value for PreferredStartAtTimeOfDay to be an explicit nil
func (o *UnconfirmedAppointment) SetPreferredStartAtTimeOfDayNil() {
	o.PreferredStartAtTimeOfDay.Set(nil)
}

// UnsetPreferredStartAtTimeOfDay ensures that no value is present for PreferredStartAtTimeOfDay, not even an explicit nil
func (o *UnconfirmedAppointment) UnsetPreferredStartAtTimeOfDay() {
	o.PreferredStartAtTimeOfDay.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnconfirmedAppointment) GetDuration() int32 {
	if o == nil || o.Duration.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnconfirmedAppointment) GetDurationOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *UnconfirmedAppointment) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableInt32 and assigns it to the Duration field.
func (o *UnconfirmedAppointment) SetDuration(v int32) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *UnconfirmedAppointment) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *UnconfirmedAppointment) UnsetDuration() {
	o.Duration.Unset()
}

func (o UnconfirmedAppointment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.PreferenceType.IsSet() {
		toSerialize["preference_type"] = o.PreferenceType.Get()
	}
	if o.PreferredStartAt.IsSet() {
		toSerialize["preferred_start_at"] = o.PreferredStartAt.Get()
	}
	if o.PreferredStartAtDay.IsSet() {
		toSerialize["preferred_start_at_day"] = o.PreferredStartAtDay.Get()
	}
	if o.PreferredStartAtTimeOfDay.IsSet() {
		toSerialize["preferred_start_at_time_of_day"] = o.PreferredStartAtTimeOfDay.Get()
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUnconfirmedAppointment struct {
	value *UnconfirmedAppointment
	isSet bool
}

func (v NullableUnconfirmedAppointment) Get() *UnconfirmedAppointment {
	return v.value
}

func (v *NullableUnconfirmedAppointment) Set(val *UnconfirmedAppointment) {
	v.value = val
	v.isSet = true
}

func (v NullableUnconfirmedAppointment) IsSet() bool {
	return v.isSet
}

func (v *NullableUnconfirmedAppointment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnconfirmedAppointment(val *UnconfirmedAppointment) *NullableUnconfirmedAppointment {
	return &NullableUnconfirmedAppointment{value: val, isSet: true}
}

func (v NullableUnconfirmedAppointment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnconfirmedAppointment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


