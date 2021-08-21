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

// AppointmentCollection A collection of appointments.
type AppointmentCollection struct {
	// What was the state of the request?
	Status string `json:"status"`
	// 
	Data []Appointment `json:"data,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
	Links *PaginationLinks `json:"links,omitempty"`
}

// NewAppointmentCollection instantiates a new AppointmentCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppointmentCollection(status string) *AppointmentCollection {
	this := AppointmentCollection{}
	this.Status = status
	return &this
}

// NewAppointmentCollectionWithDefaults instantiates a new AppointmentCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppointmentCollectionWithDefaults() *AppointmentCollection {
	this := AppointmentCollection{}
	return &this
}

// GetStatus returns the Status field value
func (o *AppointmentCollection) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AppointmentCollection) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AppointmentCollection) SetStatus(v string) {
	o.Status = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppointmentCollection) GetData() []Appointment {
	if o == nil  {
		var ret []Appointment
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppointmentCollection) GetDataOk() (*[]Appointment, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppointmentCollection) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Appointment and assigns it to the Data field.
func (o *AppointmentCollection) SetData(v []Appointment) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppointmentCollection) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentCollection) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppointmentCollection) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *AppointmentCollection) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppointmentCollection) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentCollection) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppointmentCollection) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *AppointmentCollection) SetLinks(v PaginationLinks) {
	o.Links = &v
}

func (o AppointmentCollection) MarshalJSON() ([]byte, error) {
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

type NullableAppointmentCollection struct {
	value *AppointmentCollection
	isSet bool
}

func (v NullableAppointmentCollection) Get() *AppointmentCollection {
	return v.value
}

func (v *NullableAppointmentCollection) Set(val *AppointmentCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableAppointmentCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableAppointmentCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppointmentCollection(val *AppointmentCollection) *NullableAppointmentCollection {
	return &NullableAppointmentCollection{value: val, isSet: true}
}

func (v NullableAppointmentCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppointmentCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


