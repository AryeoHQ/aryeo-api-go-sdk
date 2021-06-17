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

// OrderForm A mechanism for placing new orders on the Aryeo platform. 
type OrderForm struct {
	// UUID of the order form.
	Id string `json:"id"`
	// The name of the order form.
	Name NullableString `json:"name,omitempty"`
	// A URL of a publicly-accessible webpage for this order form.
	Url string `json:"url"`
}

// NewOrderForm instantiates a new OrderForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderForm(id string, url string) *OrderForm {
	this := OrderForm{}
	this.Id = id
	this.Url = url
	return &this
}

// NewOrderFormWithDefaults instantiates a new OrderForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderFormWithDefaults() *OrderForm {
	this := OrderForm{}
	return &this
}

// GetId returns the Id field value
func (o *OrderForm) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrderForm) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrderForm) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderForm) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderForm) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *OrderForm) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *OrderForm) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *OrderForm) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *OrderForm) UnsetName() {
	o.Name.Unset()
}

// GetUrl returns the Url field value
func (o *OrderForm) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *OrderForm) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *OrderForm) SetUrl(v string) {
	o.Url = v
}

func (o OrderForm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableOrderForm struct {
	value *OrderForm
	isSet bool
}

func (v NullableOrderForm) Get() *OrderForm {
	return v.value
}

func (v *NullableOrderForm) Set(val *OrderForm) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderForm) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderForm(val *OrderForm) *NullableOrderForm {
	return &NullableOrderForm{value: val, isSet: true}
}

func (v NullableOrderForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


