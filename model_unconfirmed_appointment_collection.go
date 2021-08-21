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

// UnconfirmedAppointmentCollection A collection of appointments.
type UnconfirmedAppointmentCollection struct {
	// What was the state of the request?
	Status string `json:"status"`
	// 
	Data []UnconfirmedAppointment `json:"data,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
	Links *PaginationLinks `json:"links,omitempty"`
}

// NewUnconfirmedAppointmentCollection instantiates a new UnconfirmedAppointmentCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnconfirmedAppointmentCollection(status string) *UnconfirmedAppointmentCollection {
	this := UnconfirmedAppointmentCollection{}
	this.Status = status
	return &this
}

// NewUnconfirmedAppointmentCollectionWithDefaults instantiates a new UnconfirmedAppointmentCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnconfirmedAppointmentCollectionWithDefaults() *UnconfirmedAppointmentCollection {
	this := UnconfirmedAppointmentCollection{}
	return &this
}

// GetStatus returns the Status field value
func (o *UnconfirmedAppointmentCollection) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UnconfirmedAppointmentCollection) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UnconfirmedAppointmentCollection) SetStatus(v string) {
	o.Status = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnconfirmedAppointmentCollection) GetData() []UnconfirmedAppointment {
	if o == nil  {
		var ret []UnconfirmedAppointment
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnconfirmedAppointmentCollection) GetDataOk() (*[]UnconfirmedAppointment, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UnconfirmedAppointmentCollection) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []UnconfirmedAppointment and assigns it to the Data field.
func (o *UnconfirmedAppointmentCollection) SetData(v []UnconfirmedAppointment) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *UnconfirmedAppointmentCollection) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnconfirmedAppointmentCollection) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *UnconfirmedAppointmentCollection) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *UnconfirmedAppointmentCollection) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UnconfirmedAppointmentCollection) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnconfirmedAppointmentCollection) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UnconfirmedAppointmentCollection) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *UnconfirmedAppointmentCollection) SetLinks(v PaginationLinks) {
	o.Links = &v
}

func (o UnconfirmedAppointmentCollection) MarshalJSON() ([]byte, error) {
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

type NullableUnconfirmedAppointmentCollection struct {
	value *UnconfirmedAppointmentCollection
	isSet bool
}

func (v NullableUnconfirmedAppointmentCollection) Get() *UnconfirmedAppointmentCollection {
	return v.value
}

func (v *NullableUnconfirmedAppointmentCollection) Set(val *UnconfirmedAppointmentCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableUnconfirmedAppointmentCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableUnconfirmedAppointmentCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnconfirmedAppointmentCollection(val *UnconfirmedAppointmentCollection) *NullableUnconfirmedAppointmentCollection {
	return &NullableUnconfirmedAppointmentCollection{value: val, isSet: true}
}

func (v NullableUnconfirmedAppointmentCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnconfirmedAppointmentCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


