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

// CalendarDay A bookable time range with available users.
type CalendarDay struct {
	// Calendar day that has available timeslots.
	Date string `json:"date"`
	// Does the day have availability?
	IsAvailable bool `json:"is_available"`
}

// NewCalendarDay instantiates a new CalendarDay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalendarDay(date string, isAvailable bool) *CalendarDay {
	this := CalendarDay{}
	this.Date = date
	this.IsAvailable = isAvailable
	return &this
}

// NewCalendarDayWithDefaults instantiates a new CalendarDay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalendarDayWithDefaults() *CalendarDay {
	this := CalendarDay{}
	return &this
}

// GetDate returns the Date field value
func (o *CalendarDay) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *CalendarDay) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *CalendarDay) SetDate(v string) {
	o.Date = v
}

// GetIsAvailable returns the IsAvailable field value
func (o *CalendarDay) GetIsAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsAvailable
}

// GetIsAvailableOk returns a tuple with the IsAvailable field value
// and a boolean to check if the value has been set.
func (o *CalendarDay) GetIsAvailableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsAvailable, true
}

// SetIsAvailable sets field value
func (o *CalendarDay) SetIsAvailable(v bool) {
	o.IsAvailable = v
}

func (o CalendarDay) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["date"] = o.Date
	}
	if true {
		toSerialize["is_available"] = o.IsAvailable
	}
	return json.Marshal(toSerialize)
}

type NullableCalendarDay struct {
	value *CalendarDay
	isSet bool
}

func (v NullableCalendarDay) Get() *CalendarDay {
	return v.value
}

func (v *NullableCalendarDay) Set(val *CalendarDay) {
	v.value = val
	v.isSet = true
}

func (v NullableCalendarDay) IsSet() bool {
	return v.isSet
}

func (v *NullableCalendarDay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalendarDay(val *CalendarDay) *NullableCalendarDay {
	return &NullableCalendarDay{value: val, isSet: true}
}

func (v NullableCalendarDay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalendarDay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


