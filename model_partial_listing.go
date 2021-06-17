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

// PartialListing A real-estate property.
type PartialListing struct {
	// ID of the listing.
	Id string `json:"id"`
	Address PartialAddress `json:"address"`
	// Has the listing been delivered?
	DeliveryStatus string `json:"delivery_status"`
	// Thumbnail URL for the listing.
	ThumbnailUrl NullableString `json:"thumbnail_url,omitempty"`
	// The price of the property in dollars.
	Price NullableInt32 `json:"price,omitempty"`
	// URL for branded property website.
	BrandedUrl NullableString `json:"branded_url,omitempty"`
	// Total number of square feet.
	SquareFeet NullableFloat32 `json:"square_feet,omitempty"`
	// Number of bedrooms.
	Bedrooms NullableInt32 `json:"bedrooms,omitempty"`
	// Number of bathrooms.
	Bathrooms NullableFloat32 `json:"bathrooms,omitempty"`
	// Are downloads enabled for this listing?
	DownloadsEnabled bool `json:"downloads_enabled"`
	// Sales status for the listing.
	Status NullableString `json:"status,omitempty"`
	PropertyDetails *PropertyDetails `json:"property_details,omitempty"`
	Agent *PartialGroup `json:"agent,omitempty"`
	CoAgent *PartialGroup `json:"co_agent,omitempty"`
}

// NewPartialListing instantiates a new PartialListing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialListing(id string, address PartialAddress, deliveryStatus string, downloadsEnabled bool) *PartialListing {
	this := PartialListing{}
	this.Id = id
	this.Address = address
	this.DeliveryStatus = deliveryStatus
	this.DownloadsEnabled = downloadsEnabled
	return &this
}

// NewPartialListingWithDefaults instantiates a new PartialListing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialListingWithDefaults() *PartialListing {
	this := PartialListing{}
	return &this
}

// GetId returns the Id field value
func (o *PartialListing) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PartialListing) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PartialListing) SetId(v string) {
	o.Id = v
}

// GetAddress returns the Address field value
func (o *PartialListing) GetAddress() PartialAddress {
	if o == nil {
		var ret PartialAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *PartialListing) GetAddressOk() (*PartialAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *PartialListing) SetAddress(v PartialAddress) {
	o.Address = v
}

// GetDeliveryStatus returns the DeliveryStatus field value
func (o *PartialListing) GetDeliveryStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliveryStatus
}

// GetDeliveryStatusOk returns a tuple with the DeliveryStatus field value
// and a boolean to check if the value has been set.
func (o *PartialListing) GetDeliveryStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DeliveryStatus, true
}

// SetDeliveryStatus sets field value
func (o *PartialListing) SetDeliveryStatus(v string) {
	o.DeliveryStatus = v
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialListing) GetThumbnailUrl() string {
	if o == nil || o.ThumbnailUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl.Get()
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialListing) GetThumbnailUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ThumbnailUrl.Get(), o.ThumbnailUrl.IsSet()
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *PartialListing) HasThumbnailUrl() bool {
	if o != nil && o.ThumbnailUrl.IsSet() {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given NullableString and assigns it to the ThumbnailUrl field.
func (o *PartialListing) SetThumbnailUrl(v string) {
	o.ThumbnailUrl.Set(&v)
}
// SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil
func (o *PartialListing) SetThumbnailUrlNil() {
	o.ThumbnailUrl.Set(nil)
}

// UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
func (o *PartialListing) UnsetThumbnailUrl() {
	o.ThumbnailUrl.Unset()
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialListing) GetPrice() int32 {
	if o == nil || o.Price.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialListing) GetPriceOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *PartialListing) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableInt32 and assigns it to the Price field.
func (o *PartialListing) SetPrice(v int32) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *PartialListing) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *PartialListing) UnsetPrice() {
	o.Price.Unset()
}

// GetBrandedUrl returns the BrandedUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialListing) GetBrandedUrl() string {
	if o == nil || o.BrandedUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.BrandedUrl.Get()
}

// GetBrandedUrlOk returns a tuple with the BrandedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialListing) GetBrandedUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BrandedUrl.Get(), o.BrandedUrl.IsSet()
}

// HasBrandedUrl returns a boolean if a field has been set.
func (o *PartialListing) HasBrandedUrl() bool {
	if o != nil && o.BrandedUrl.IsSet() {
		return true
	}

	return false
}

// SetBrandedUrl gets a reference to the given NullableString and assigns it to the BrandedUrl field.
func (o *PartialListing) SetBrandedUrl(v string) {
	o.BrandedUrl.Set(&v)
}
// SetBrandedUrlNil sets the value for BrandedUrl to be an explicit nil
func (o *PartialListing) SetBrandedUrlNil() {
	o.BrandedUrl.Set(nil)
}

// UnsetBrandedUrl ensures that no value is present for BrandedUrl, not even an explicit nil
func (o *PartialListing) UnsetBrandedUrl() {
	o.BrandedUrl.Unset()
}

// GetSquareFeet returns the SquareFeet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialListing) GetSquareFeet() float32 {
	if o == nil || o.SquareFeet.Get() == nil {
		var ret float32
		return ret
	}
	return *o.SquareFeet.Get()
}

// GetSquareFeetOk returns a tuple with the SquareFeet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialListing) GetSquareFeetOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SquareFeet.Get(), o.SquareFeet.IsSet()
}

// HasSquareFeet returns a boolean if a field has been set.
func (o *PartialListing) HasSquareFeet() bool {
	if o != nil && o.SquareFeet.IsSet() {
		return true
	}

	return false
}

// SetSquareFeet gets a reference to the given NullableFloat32 and assigns it to the SquareFeet field.
func (o *PartialListing) SetSquareFeet(v float32) {
	o.SquareFeet.Set(&v)
}
// SetSquareFeetNil sets the value for SquareFeet to be an explicit nil
func (o *PartialListing) SetSquareFeetNil() {
	o.SquareFeet.Set(nil)
}

// UnsetSquareFeet ensures that no value is present for SquareFeet, not even an explicit nil
func (o *PartialListing) UnsetSquareFeet() {
	o.SquareFeet.Unset()
}

// GetBedrooms returns the Bedrooms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialListing) GetBedrooms() int32 {
	if o == nil || o.Bedrooms.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Bedrooms.Get()
}

// GetBedroomsOk returns a tuple with the Bedrooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialListing) GetBedroomsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bedrooms.Get(), o.Bedrooms.IsSet()
}

// HasBedrooms returns a boolean if a field has been set.
func (o *PartialListing) HasBedrooms() bool {
	if o != nil && o.Bedrooms.IsSet() {
		return true
	}

	return false
}

// SetBedrooms gets a reference to the given NullableInt32 and assigns it to the Bedrooms field.
func (o *PartialListing) SetBedrooms(v int32) {
	o.Bedrooms.Set(&v)
}
// SetBedroomsNil sets the value for Bedrooms to be an explicit nil
func (o *PartialListing) SetBedroomsNil() {
	o.Bedrooms.Set(nil)
}

// UnsetBedrooms ensures that no value is present for Bedrooms, not even an explicit nil
func (o *PartialListing) UnsetBedrooms() {
	o.Bedrooms.Unset()
}

// GetBathrooms returns the Bathrooms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialListing) GetBathrooms() float32 {
	if o == nil || o.Bathrooms.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Bathrooms.Get()
}

// GetBathroomsOk returns a tuple with the Bathrooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialListing) GetBathroomsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bathrooms.Get(), o.Bathrooms.IsSet()
}

// HasBathrooms returns a boolean if a field has been set.
func (o *PartialListing) HasBathrooms() bool {
	if o != nil && o.Bathrooms.IsSet() {
		return true
	}

	return false
}

// SetBathrooms gets a reference to the given NullableFloat32 and assigns it to the Bathrooms field.
func (o *PartialListing) SetBathrooms(v float32) {
	o.Bathrooms.Set(&v)
}
// SetBathroomsNil sets the value for Bathrooms to be an explicit nil
func (o *PartialListing) SetBathroomsNil() {
	o.Bathrooms.Set(nil)
}

// UnsetBathrooms ensures that no value is present for Bathrooms, not even an explicit nil
func (o *PartialListing) UnsetBathrooms() {
	o.Bathrooms.Unset()
}

// GetDownloadsEnabled returns the DownloadsEnabled field value
func (o *PartialListing) GetDownloadsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DownloadsEnabled
}

// GetDownloadsEnabledOk returns a tuple with the DownloadsEnabled field value
// and a boolean to check if the value has been set.
func (o *PartialListing) GetDownloadsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DownloadsEnabled, true
}

// SetDownloadsEnabled sets field value
func (o *PartialListing) SetDownloadsEnabled(v bool) {
	o.DownloadsEnabled = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PartialListing) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PartialListing) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *PartialListing) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *PartialListing) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *PartialListing) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *PartialListing) UnsetStatus() {
	o.Status.Unset()
}

// GetPropertyDetails returns the PropertyDetails field value if set, zero value otherwise.
func (o *PartialListing) GetPropertyDetails() PropertyDetails {
	if o == nil || o.PropertyDetails == nil {
		var ret PropertyDetails
		return ret
	}
	return *o.PropertyDetails
}

// GetPropertyDetailsOk returns a tuple with the PropertyDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialListing) GetPropertyDetailsOk() (*PropertyDetails, bool) {
	if o == nil || o.PropertyDetails == nil {
		return nil, false
	}
	return o.PropertyDetails, true
}

// HasPropertyDetails returns a boolean if a field has been set.
func (o *PartialListing) HasPropertyDetails() bool {
	if o != nil && o.PropertyDetails != nil {
		return true
	}

	return false
}

// SetPropertyDetails gets a reference to the given PropertyDetails and assigns it to the PropertyDetails field.
func (o *PartialListing) SetPropertyDetails(v PropertyDetails) {
	o.PropertyDetails = &v
}

// GetAgent returns the Agent field value if set, zero value otherwise.
func (o *PartialListing) GetAgent() PartialGroup {
	if o == nil || o.Agent == nil {
		var ret PartialGroup
		return ret
	}
	return *o.Agent
}

// GetAgentOk returns a tuple with the Agent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialListing) GetAgentOk() (*PartialGroup, bool) {
	if o == nil || o.Agent == nil {
		return nil, false
	}
	return o.Agent, true
}

// HasAgent returns a boolean if a field has been set.
func (o *PartialListing) HasAgent() bool {
	if o != nil && o.Agent != nil {
		return true
	}

	return false
}

// SetAgent gets a reference to the given PartialGroup and assigns it to the Agent field.
func (o *PartialListing) SetAgent(v PartialGroup) {
	o.Agent = &v
}

// GetCoAgent returns the CoAgent field value if set, zero value otherwise.
func (o *PartialListing) GetCoAgent() PartialGroup {
	if o == nil || o.CoAgent == nil {
		var ret PartialGroup
		return ret
	}
	return *o.CoAgent
}

// GetCoAgentOk returns a tuple with the CoAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialListing) GetCoAgentOk() (*PartialGroup, bool) {
	if o == nil || o.CoAgent == nil {
		return nil, false
	}
	return o.CoAgent, true
}

// HasCoAgent returns a boolean if a field has been set.
func (o *PartialListing) HasCoAgent() bool {
	if o != nil && o.CoAgent != nil {
		return true
	}

	return false
}

// SetCoAgent gets a reference to the given PartialGroup and assigns it to the CoAgent field.
func (o *PartialListing) SetCoAgent(v PartialGroup) {
	o.CoAgent = &v
}

func (o PartialListing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["delivery_status"] = o.DeliveryStatus
	}
	if o.ThumbnailUrl.IsSet() {
		toSerialize["thumbnail_url"] = o.ThumbnailUrl.Get()
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.BrandedUrl.IsSet() {
		toSerialize["branded_url"] = o.BrandedUrl.Get()
	}
	if o.SquareFeet.IsSet() {
		toSerialize["square_feet"] = o.SquareFeet.Get()
	}
	if o.Bedrooms.IsSet() {
		toSerialize["bedrooms"] = o.Bedrooms.Get()
	}
	if o.Bathrooms.IsSet() {
		toSerialize["bathrooms"] = o.Bathrooms.Get()
	}
	if true {
		toSerialize["downloads_enabled"] = o.DownloadsEnabled
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.PropertyDetails != nil {
		toSerialize["property_details"] = o.PropertyDetails
	}
	if o.Agent != nil {
		toSerialize["agent"] = o.Agent
	}
	if o.CoAgent != nil {
		toSerialize["co_agent"] = o.CoAgent
	}
	return json.Marshal(toSerialize)
}

type NullablePartialListing struct {
	value *PartialListing
	isSet bool
}

func (v NullablePartialListing) Get() *PartialListing {
	return v.value
}

func (v *NullablePartialListing) Set(val *PartialListing) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialListing) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialListing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialListing(val *PartialListing) *NullablePartialListing {
	return &NullablePartialListing{value: val, isSet: true}
}

func (v NullablePartialListing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialListing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


