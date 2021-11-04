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

// TimeslotCollection A collection of bookable timeslots.
type TimeslotCollection struct {
	// What was the state of the request?
	Status string `json:"status"`
	// 
	Data []Timeslot `json:"data,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
	Links *PaginationLinks `json:"links,omitempty"`
}

// NewTimeslotCollection instantiates a new TimeslotCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeslotCollection(status string) *TimeslotCollection {
	this := TimeslotCollection{}
	this.Status = status
	return &this
}

// NewTimeslotCollectionWithDefaults instantiates a new TimeslotCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeslotCollectionWithDefaults() *TimeslotCollection {
	this := TimeslotCollection{}
	return &this
}

// GetStatus returns the Status field value
func (o *TimeslotCollection) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TimeslotCollection) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TimeslotCollection) SetStatus(v string) {
	o.Status = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimeslotCollection) GetData() []Timeslot {
	if o == nil  {
		var ret []Timeslot
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimeslotCollection) GetDataOk() (*[]Timeslot, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TimeslotCollection) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Timeslot and assigns it to the Data field.
func (o *TimeslotCollection) SetData(v []Timeslot) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TimeslotCollection) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeslotCollection) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TimeslotCollection) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *TimeslotCollection) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TimeslotCollection) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeslotCollection) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TimeslotCollection) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *TimeslotCollection) SetLinks(v PaginationLinks) {
	o.Links = &v
}

func (o TimeslotCollection) MarshalJSON() ([]byte, error) {
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

type NullableTimeslotCollection struct {
	value *TimeslotCollection
	isSet bool
}

func (v NullableTimeslotCollection) Get() *TimeslotCollection {
	return v.value
}

func (v *NullableTimeslotCollection) Set(val *TimeslotCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeslotCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeslotCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeslotCollection(val *TimeslotCollection) *NullableTimeslotCollection {
	return &NullableTimeslotCollection{value: val, isSet: true}
}

func (v NullableTimeslotCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeslotCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


