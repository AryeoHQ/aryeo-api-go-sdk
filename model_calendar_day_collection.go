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

// CalendarDayCollection A collection of calendar days that have available timeslots
type CalendarDayCollection struct {
	// What was the state of the request?
	Status string `json:"status"`
	// 
	Data []CalendarDay `json:"data,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
	Links *PaginationLinks `json:"links,omitempty"`
}

// NewCalendarDayCollection instantiates a new CalendarDayCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalendarDayCollection(status string) *CalendarDayCollection {
	this := CalendarDayCollection{}
	this.Status = status
	return &this
}

// NewCalendarDayCollectionWithDefaults instantiates a new CalendarDayCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalendarDayCollectionWithDefaults() *CalendarDayCollection {
	this := CalendarDayCollection{}
	return &this
}

// GetStatus returns the Status field value
func (o *CalendarDayCollection) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CalendarDayCollection) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CalendarDayCollection) SetStatus(v string) {
	o.Status = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CalendarDayCollection) GetData() []CalendarDay {
	if o == nil  {
		var ret []CalendarDay
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CalendarDayCollection) GetDataOk() (*[]CalendarDay, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CalendarDayCollection) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []CalendarDay and assigns it to the Data field.
func (o *CalendarDayCollection) SetData(v []CalendarDay) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CalendarDayCollection) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarDayCollection) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CalendarDayCollection) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *CalendarDayCollection) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CalendarDayCollection) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarDayCollection) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CalendarDayCollection) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *CalendarDayCollection) SetLinks(v PaginationLinks) {
	o.Links = &v
}

func (o CalendarDayCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableCalendarDayCollection struct {
	value *CalendarDayCollection
	isSet bool
}

func (v NullableCalendarDayCollection) Get() *CalendarDayCollection {
	return v.value
}

func (v *NullableCalendarDayCollection) Set(val *CalendarDayCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableCalendarDayCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableCalendarDayCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalendarDayCollection(val *CalendarDayCollection) *NullableCalendarDayCollection {
	return &NullableCalendarDayCollection{value: val, isSet: true}
}

func (v NullableCalendarDayCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalendarDayCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


