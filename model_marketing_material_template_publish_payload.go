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

// MarketingMaterialTemplatePublishPayload Payload for publishing a marketing material template record.
type MarketingMaterialTemplatePublishPayload struct {
	// String representation of a polotno JSON object.
	PolotnoJson *string `json:"polotno_json,omitempty"`
}

// NewMarketingMaterialTemplatePublishPayload instantiates a new MarketingMaterialTemplatePublishPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketingMaterialTemplatePublishPayload() *MarketingMaterialTemplatePublishPayload {
	this := MarketingMaterialTemplatePublishPayload{}
	return &this
}

// NewMarketingMaterialTemplatePublishPayloadWithDefaults instantiates a new MarketingMaterialTemplatePublishPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketingMaterialTemplatePublishPayloadWithDefaults() *MarketingMaterialTemplatePublishPayload {
	this := MarketingMaterialTemplatePublishPayload{}
	return &this
}

// GetPolotnoJson returns the PolotnoJson field value if set, zero value otherwise.
func (o *MarketingMaterialTemplatePublishPayload) GetPolotnoJson() string {
	if o == nil || o.PolotnoJson == nil {
		var ret string
		return ret
	}
	return *o.PolotnoJson
}

// GetPolotnoJsonOk returns a tuple with the PolotnoJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingMaterialTemplatePublishPayload) GetPolotnoJsonOk() (*string, bool) {
	if o == nil || o.PolotnoJson == nil {
		return nil, false
	}
	return o.PolotnoJson, true
}

// HasPolotnoJson returns a boolean if a field has been set.
func (o *MarketingMaterialTemplatePublishPayload) HasPolotnoJson() bool {
	if o != nil && o.PolotnoJson != nil {
		return true
	}

	return false
}

// SetPolotnoJson gets a reference to the given string and assigns it to the PolotnoJson field.
func (o *MarketingMaterialTemplatePublishPayload) SetPolotnoJson(v string) {
	o.PolotnoJson = &v
}

func (o MarketingMaterialTemplatePublishPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PolotnoJson != nil {
		toSerialize["polotno_json"] = o.PolotnoJson
	}
	return json.Marshal(toSerialize)
}

type NullableMarketingMaterialTemplatePublishPayload struct {
	value *MarketingMaterialTemplatePublishPayload
	isSet bool
}

func (v NullableMarketingMaterialTemplatePublishPayload) Get() *MarketingMaterialTemplatePublishPayload {
	return v.value
}

func (v *NullableMarketingMaterialTemplatePublishPayload) Set(val *MarketingMaterialTemplatePublishPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketingMaterialTemplatePublishPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketingMaterialTemplatePublishPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketingMaterialTemplatePublishPayload(val *MarketingMaterialTemplatePublishPayload) *NullableMarketingMaterialTemplatePublishPayload {
	return &NullableMarketingMaterialTemplatePublishPayload{value: val, isSet: true}
}

func (v NullableMarketingMaterialTemplatePublishPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketingMaterialTemplatePublishPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


