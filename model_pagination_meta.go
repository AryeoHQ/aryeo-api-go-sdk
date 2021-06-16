/*
 * Aryeo
 *
 * # Introduction The Aryeo API gives access to the Aryeo platform. You can use your favorite HTTP/REST library for interfacing with the Aryeo API, or you can use one of our SDKs. Our SDKs are procedurally generated and a great starting point for experimental testing. If there is an additional language or framework that you want to see supported, then please reach and out and make a contribution!  <p align=\"center\"> <a href=\"https://github.com/AryeoHQ/aryeo-api-dart-sdk\"><img src=\"https://raw.githubusercontent.com/AryeoHQ/aryeo-api-docs/master/public/images/dart.svg\" alt=\"Dart\" width=\"44\" style=\"padding:10px;border: 1px solid #d3d3d3;border-radius: 5px;margin:4px;\"/></a> <a href=\"https://github.com/AryeoHQ/aryeo-api-go-sdk\"><img src=\"https://raw.githubusercontent.com/AryeoHQ/aryeo-api-docs/master/public/images/go.svg\" alt=\"Go\" width=\"44\" style=\"padding:10px;border: 1px solid #d3d3d3;border-radius: 5px;margin:4px;\"/></a> <a href=\"https://github.com/AryeoHQ/aryeo-api-js-sdk\"><img src=\"https://raw.githubusercontent.com/AryeoHQ/aryeo-api-docs/master/public/images/js.svg\" alt=\"Node JS\" width=\"44\" style=\"padding:10px;border: 1px solid #d3d3d3;border-radius: 5px;margin:4px;\"/></a> <a href=\"https://github.com/AryeoHQ/aryeo-api-php-sdk\"><img src=\"https://raw.githubusercontent.com/AryeoHQ/aryeo-api-docs/master/public/images/php.svg\" alt=\"PHP\" width=\"44\" style=\"padding:10px;border: 1px solid #d3d3d3;border-radius: 5px;margin:4px;\"/></a> <a href=\"https://github.com/AryeoHQ/aryeo-api-ruby-sdk\"><img src=\"https://raw.githubusercontent.com/AryeoHQ/aryeo-api-docs/master/public/images/ruby.svg\" alt=\"Ruby\" width=\"44\" style=\"padding:10px;border: 1px solid #d3d3d3;border-radius: 5px;margin:4px;\"/></a> <a href=\"https://github.com/AryeoHQ/aryeo-api-rust-sdk\"><img src=\"https://raw.githubusercontent.com/AryeoHQ/aryeo-api-docs/master/public/images/rust.svg\" alt=\"Rust\" width=\"44\" style=\"padding:10px;border: 1px solid #d3d3d3;border-radius: 5px;margin:4px;\"/></a> <a href=\"https://github.com/AryeoHQ/aryeo-api-swift-sdk\"><img src=\"https://raw.githubusercontent.com/AryeoHQ/aryeo-api-docs/master/public/images/swift.svg\" alt=\"Swift\" width=\"44\" style=\"padding:10px;border: 1px solid #d3d3d3;border-radius: 5px;margin:4px;\"/></a> </p>  **Note**: Some SDKs may require you to manually add the `Accept` header to all Aryeo API requests. If so, use the value `application/json`.  # Authentication To start using the Aryeo API, you will need to generate an API key from your group's developer settings. You can then authenticate to the Aryeo API by providing your key in the appropriate request header. Requests made without an API key will result in a `401 Unauthorized` error. 
 *
 * API version: 1.0.0
 * Contact: jarrod@aryeo.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aryeo

import (
	"encoding/json"
)

// PaginationMeta Metadata about a paginated response.
type PaginationMeta struct {
	// Total number of records.
	Total int32 `json:"total"`
	// Number of records per page.
	PerPage int32 `json:"per_page"`
	// The current page.
	CurrentPage int32 `json:"current_page"`
	// The last page of records.
	LastPage int32 `json:"last_page"`
	// The ID of the first record on this page. This is specified as either `integer` or `null` purely for contract testing purposes. The model which is autogenerated from this definition will be thrown out and written by-hand.
	From NullableInt32 `json:"from,omitempty"`
	// The ID of the last record on this page. This is specified as either `integer` or `null` purely for contract testing purposes. The model which is autogenerated from this definition will be thrown out and written by-hand.
	To NullableInt32 `json:"to,omitempty"`
	// The current paged path.
	Path string `json:"path"`
}

// NewPaginationMeta instantiates a new PaginationMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationMeta(total int32, perPage int32, currentPage int32, lastPage int32, path string) *PaginationMeta {
	this := PaginationMeta{}
	this.Total = total
	this.PerPage = perPage
	this.CurrentPage = currentPage
	this.LastPage = lastPage
	this.Path = path
	return &this
}

// NewPaginationMetaWithDefaults instantiates a new PaginationMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationMetaWithDefaults() *PaginationMeta {
	this := PaginationMeta{}
	return &this
}

// GetTotal returns the Total field value
func (o *PaginationMeta) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PaginationMeta) GetTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *PaginationMeta) SetTotal(v int32) {
	o.Total = v
}

// GetPerPage returns the PerPage field value
func (o *PaginationMeta) GetPerPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value
// and a boolean to check if the value has been set.
func (o *PaginationMeta) GetPerPageOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PerPage, true
}

// SetPerPage sets field value
func (o *PaginationMeta) SetPerPage(v int32) {
	o.PerPage = v
}

// GetCurrentPage returns the CurrentPage field value
func (o *PaginationMeta) GetCurrentPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value
// and a boolean to check if the value has been set.
func (o *PaginationMeta) GetCurrentPageOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CurrentPage, true
}

// SetCurrentPage sets field value
func (o *PaginationMeta) SetCurrentPage(v int32) {
	o.CurrentPage = v
}

// GetLastPage returns the LastPage field value
func (o *PaginationMeta) GetLastPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LastPage
}

// GetLastPageOk returns a tuple with the LastPage field value
// and a boolean to check if the value has been set.
func (o *PaginationMeta) GetLastPageOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastPage, true
}

// SetLastPage sets field value
func (o *PaginationMeta) SetLastPage(v int32) {
	o.LastPage = v
}

// GetFrom returns the From field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginationMeta) GetFrom() int32 {
	if o == nil || o.From.Get() == nil {
		var ret int32
		return ret
	}
	return *o.From.Get()
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationMeta) GetFromOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.From.Get(), o.From.IsSet()
}

// HasFrom returns a boolean if a field has been set.
func (o *PaginationMeta) HasFrom() bool {
	if o != nil && o.From.IsSet() {
		return true
	}

	return false
}

// SetFrom gets a reference to the given NullableInt32 and assigns it to the From field.
func (o *PaginationMeta) SetFrom(v int32) {
	o.From.Set(&v)
}
// SetFromNil sets the value for From to be an explicit nil
func (o *PaginationMeta) SetFromNil() {
	o.From.Set(nil)
}

// UnsetFrom ensures that no value is present for From, not even an explicit nil
func (o *PaginationMeta) UnsetFrom() {
	o.From.Unset()
}

// GetTo returns the To field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginationMeta) GetTo() int32 {
	if o == nil || o.To.Get() == nil {
		var ret int32
		return ret
	}
	return *o.To.Get()
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationMeta) GetToOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.To.Get(), o.To.IsSet()
}

// HasTo returns a boolean if a field has been set.
func (o *PaginationMeta) HasTo() bool {
	if o != nil && o.To.IsSet() {
		return true
	}

	return false
}

// SetTo gets a reference to the given NullableInt32 and assigns it to the To field.
func (o *PaginationMeta) SetTo(v int32) {
	o.To.Set(&v)
}
// SetToNil sets the value for To to be an explicit nil
func (o *PaginationMeta) SetToNil() {
	o.To.Set(nil)
}

// UnsetTo ensures that no value is present for To, not even an explicit nil
func (o *PaginationMeta) UnsetTo() {
	o.To.Unset()
}

// GetPath returns the Path field value
func (o *PaginationMeta) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *PaginationMeta) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *PaginationMeta) SetPath(v string) {
	o.Path = v
}

func (o PaginationMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["per_page"] = o.PerPage
	}
	if true {
		toSerialize["current_page"] = o.CurrentPage
	}
	if true {
		toSerialize["last_page"] = o.LastPage
	}
	if o.From.IsSet() {
		toSerialize["from"] = o.From.Get()
	}
	if o.To.IsSet() {
		toSerialize["to"] = o.To.Get()
	}
	if true {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullablePaginationMeta struct {
	value *PaginationMeta
	isSet bool
}

func (v NullablePaginationMeta) Get() *PaginationMeta {
	return v.value
}

func (v *NullablePaginationMeta) Set(val *PaginationMeta) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationMeta) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationMeta(val *PaginationMeta) *NullablePaginationMeta {
	return &NullablePaginationMeta{value: val, isSet: true}
}

func (v NullablePaginationMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

