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

// ListingResource A listing.
type ListingResource struct {
	// What was the state of the request?
	Status string `json:"status"`
	Data *Listing `json:"data,omitempty"`
}

// NewListingResource instantiates a new ListingResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListingResource(status string) *ListingResource {
	this := ListingResource{}
	this.Status = status
	return &this
}

// NewListingResourceWithDefaults instantiates a new ListingResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListingResourceWithDefaults() *ListingResource {
	this := ListingResource{}
	return &this
}

// GetStatus returns the Status field value
func (o *ListingResource) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ListingResource) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ListingResource) SetStatus(v string) {
	o.Status = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListingResource) GetData() Listing {
	if o == nil || o.Data == nil {
		var ret Listing
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingResource) GetDataOk() (*Listing, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListingResource) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given Listing and assigns it to the Data field.
func (o *ListingResource) SetData(v Listing) {
	o.Data = &v
}

func (o ListingResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListingResource struct {
	value *ListingResource
	isSet bool
}

func (v NullableListingResource) Get() *ListingResource {
	return v.value
}

func (v *NullableListingResource) Set(val *ListingResource) {
	v.value = val
	v.isSet = true
}

func (v NullableListingResource) IsSet() bool {
	return v.isSet
}

func (v *NullableListingResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListingResource(val *ListingResource) *NullableListingResource {
	return &NullableListingResource{value: val, isSet: true}
}

func (v NullableListingResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListingResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


