/*
 * Aryeo
 *
 *
 * API version: 1.0.0
 * Contact: jarrod@aryeo.com
 */

package aryeo

import (
	"encoding/json"
)

// MarketingMaterialPublishPayload Payload for publishing a marketing material record.
type MarketingMaterialPublishPayload struct {
	// String representation of a polotno JSON object.
	PolotnoJson *string `json:"polotno_json,omitempty"`
}

// NewMarketingMaterialPublishPayload instantiates a new MarketingMaterialPublishPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketingMaterialPublishPayload() *MarketingMaterialPublishPayload {
	this := MarketingMaterialPublishPayload{}
	return &this
}

// NewMarketingMaterialPublishPayloadWithDefaults instantiates a new MarketingMaterialPublishPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketingMaterialPublishPayloadWithDefaults() *MarketingMaterialPublishPayload {
	this := MarketingMaterialPublishPayload{}
	return &this
}

// GetPolotnoJson returns the PolotnoJson field value if set, zero value otherwise.
func (o *MarketingMaterialPublishPayload) GetPolotnoJson() string {
	if o == nil || o.PolotnoJson == nil {
		var ret string
		return ret
	}
	return *o.PolotnoJson
}

// GetPolotnoJsonOk returns a tuple with the PolotnoJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingMaterialPublishPayload) GetPolotnoJsonOk() (*string, bool) {
	if o == nil || o.PolotnoJson == nil {
		return nil, false
	}
	return o.PolotnoJson, true
}

// HasPolotnoJson returns a boolean if a field has been set.
func (o *MarketingMaterialPublishPayload) HasPolotnoJson() bool {
	if o != nil && o.PolotnoJson != nil {
		return true
	}

	return false
}

// SetPolotnoJson gets a reference to the given string and assigns it to the PolotnoJson field.
func (o *MarketingMaterialPublishPayload) SetPolotnoJson(v string) {
	o.PolotnoJson = &v
}

func (o MarketingMaterialPublishPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PolotnoJson != nil {
		toSerialize["polotno_json"] = o.PolotnoJson
	}
	return json.Marshal(toSerialize)
}

type NullableMarketingMaterialPublishPayload struct {
	value *MarketingMaterialPublishPayload
	isSet bool
}

func (v NullableMarketingMaterialPublishPayload) Get() *MarketingMaterialPublishPayload {
	return v.value
}

func (v *NullableMarketingMaterialPublishPayload) Set(val *MarketingMaterialPublishPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketingMaterialPublishPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketingMaterialPublishPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketingMaterialPublishPayload(val *MarketingMaterialPublishPayload) *NullableMarketingMaterialPublishPayload {
	return &NullableMarketingMaterialPublishPayload{value: val, isSet: true}
}

func (v NullableMarketingMaterialPublishPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketingMaterialPublishPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


