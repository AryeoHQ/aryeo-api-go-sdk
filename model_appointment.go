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

// Appointment An appointment.
type Appointment struct {
	// The ID of the appointment.
	Id string `json:"id"`
	// The status of the appointment.
	Status NullableString `json:"status,omitempty"`
	// The title of the appointment.
	Title NullableString `json:"title,omitempty"`
	// The multi-line description of the appointment.
	Description NullableString `json:"description,omitempty"`
	// The date and time (ISO 8601 format) when the appointment is set to start.
	StartAt NullableTime `json:"start_at,omitempty"`
	// The date and time (ISO 8601 format) when the appointment is set to end.
	EndAt NullableTime `json:"end_at,omitempty"`
	// The length of the appointment in minutes.
	Duration NullableInt32 `json:"duration,omitempty"`
	Order *Order `json:"order,omitempty"`
	// Users attached to the appointment.
	Users []User `json:"users,omitempty"`
}

// NewAppointment instantiates a new Appointment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppointment(id string) *Appointment {
	this := Appointment{}
	this.Id = id
	return &this
}

// NewAppointmentWithDefaults instantiates a new Appointment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppointmentWithDefaults() *Appointment {
	this := Appointment{}
	return &this
}

// GetId returns the Id field value
func (o *Appointment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Appointment) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Appointment) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Appointment) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Appointment) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *Appointment) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *Appointment) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *Appointment) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *Appointment) UnsetStatus() {
	o.Status.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Appointment) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Appointment) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Appointment) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Appointment) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Appointment) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Appointment) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Appointment) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Appointment) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Appointment) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Appointment) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Appointment) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Appointment) UnsetDescription() {
	o.Description.Unset()
}

// GetStartAt returns the StartAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Appointment) GetStartAt() time.Time {
	if o == nil || o.StartAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartAt.Get()
}

// GetStartAtOk returns a tuple with the StartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Appointment) GetStartAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartAt.Get(), o.StartAt.IsSet()
}

// HasStartAt returns a boolean if a field has been set.
func (o *Appointment) HasStartAt() bool {
	if o != nil && o.StartAt.IsSet() {
		return true
	}

	return false
}

// SetStartAt gets a reference to the given NullableTime and assigns it to the StartAt field.
func (o *Appointment) SetStartAt(v time.Time) {
	o.StartAt.Set(&v)
}
// SetStartAtNil sets the value for StartAt to be an explicit nil
func (o *Appointment) SetStartAtNil() {
	o.StartAt.Set(nil)
}

// UnsetStartAt ensures that no value is present for StartAt, not even an explicit nil
func (o *Appointment) UnsetStartAt() {
	o.StartAt.Unset()
}

// GetEndAt returns the EndAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Appointment) GetEndAt() time.Time {
	if o == nil || o.EndAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndAt.Get()
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Appointment) GetEndAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndAt.Get(), o.EndAt.IsSet()
}

// HasEndAt returns a boolean if a field has been set.
func (o *Appointment) HasEndAt() bool {
	if o != nil && o.EndAt.IsSet() {
		return true
	}

	return false
}

// SetEndAt gets a reference to the given NullableTime and assigns it to the EndAt field.
func (o *Appointment) SetEndAt(v time.Time) {
	o.EndAt.Set(&v)
}
// SetEndAtNil sets the value for EndAt to be an explicit nil
func (o *Appointment) SetEndAtNil() {
	o.EndAt.Set(nil)
}

// UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
func (o *Appointment) UnsetEndAt() {
	o.EndAt.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Appointment) GetDuration() int32 {
	if o == nil || o.Duration.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Appointment) GetDurationOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *Appointment) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableInt32 and assigns it to the Duration field.
func (o *Appointment) SetDuration(v int32) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *Appointment) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *Appointment) UnsetDuration() {
	o.Duration.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *Appointment) GetOrder() Order {
	if o == nil || o.Order == nil {
		var ret Order
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appointment) GetOrderOk() (*Order, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *Appointment) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given Order and assigns it to the Order field.
func (o *Appointment) SetOrder(v Order) {
	o.Order = &v
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Appointment) GetUsers() []User {
	if o == nil  {
		var ret []User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Appointment) GetUsersOk() (*[]User, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *Appointment) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *Appointment) SetUsers(v []User) {
	o.Users = v
}

func (o Appointment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.StartAt.IsSet() {
		toSerialize["start_at"] = o.StartAt.Get()
	}
	if o.EndAt.IsSet() {
		toSerialize["end_at"] = o.EndAt.Get()
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableAppointment struct {
	value *Appointment
	isSet bool
}

func (v NullableAppointment) Get() *Appointment {
	return v.value
}

func (v *NullableAppointment) Set(val *Appointment) {
	v.value = val
	v.isSet = true
}

func (v NullableAppointment) IsSet() bool {
	return v.isSet
}

func (v *NullableAppointment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppointment(val *Appointment) *NullableAppointment {
	return &NullableAppointment{value: val, isSet: true}
}

func (v NullableAppointment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppointment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


