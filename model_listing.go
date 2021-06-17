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

// Listing A real-estate property.
type Listing struct {
	// ID of the listing.
	Id string `json:"id"`
	Address PartialAddress `json:"address"`
	// Has this listing been delivered?
	DeliveryStatus string `json:"delivery_status"`
	// Thumbnail URL for the listing.
	ThumbnailUrl NullableString `json:"thumbnail_url,omitempty"`
	Agent *Group `json:"agent,omitempty"`
	CoAgent *Group `json:"co_agent,omitempty"`
	// images
	Images *[]Image `json:"images,omitempty"`
	// videos
	Videos *[]Video `json:"videos,omitempty"`
	// floor_plans
	FloorPlans *[]FloorPlan `json:"floor_plans,omitempty"`
	PropertyWebsites *PropertyWebsites `json:"property_websites,omitempty"`
	// interactive_content
	InteractiveContent *[]InteractiveContent `json:"interactive_content,omitempty"`
	PropertyDetails *PropertyDetails `json:"property_details,omitempty"`
	// Are downloads enabled for this listing?
	DownloadsEnabled bool `json:"downloads_enabled"`
	// orders
	Orders *[]Order `json:"orders,omitempty"`
}

// NewListing instantiates a new Listing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListing(id string, address PartialAddress, deliveryStatus string, downloadsEnabled bool) *Listing {
	this := Listing{}
	this.Id = id
	this.Address = address
	this.DeliveryStatus = deliveryStatus
	this.DownloadsEnabled = downloadsEnabled
	return &this
}

// NewListingWithDefaults instantiates a new Listing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListingWithDefaults() *Listing {
	this := Listing{}
	return &this
}

// GetId returns the Id field value
func (o *Listing) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Listing) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Listing) SetId(v string) {
	o.Id = v
}

// GetAddress returns the Address field value
func (o *Listing) GetAddress() PartialAddress {
	if o == nil {
		var ret PartialAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *Listing) GetAddressOk() (*PartialAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *Listing) SetAddress(v PartialAddress) {
	o.Address = v
}

// GetDeliveryStatus returns the DeliveryStatus field value
func (o *Listing) GetDeliveryStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliveryStatus
}

// GetDeliveryStatusOk returns a tuple with the DeliveryStatus field value
// and a boolean to check if the value has been set.
func (o *Listing) GetDeliveryStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DeliveryStatus, true
}

// SetDeliveryStatus sets field value
func (o *Listing) SetDeliveryStatus(v string) {
	o.DeliveryStatus = v
}

// GetThumbnailUrl returns the ThumbnailUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Listing) GetThumbnailUrl() string {
	if o == nil || o.ThumbnailUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ThumbnailUrl.Get()
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Listing) GetThumbnailUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ThumbnailUrl.Get(), o.ThumbnailUrl.IsSet()
}

// HasThumbnailUrl returns a boolean if a field has been set.
func (o *Listing) HasThumbnailUrl() bool {
	if o != nil && o.ThumbnailUrl.IsSet() {
		return true
	}

	return false
}

// SetThumbnailUrl gets a reference to the given NullableString and assigns it to the ThumbnailUrl field.
func (o *Listing) SetThumbnailUrl(v string) {
	o.ThumbnailUrl.Set(&v)
}
// SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil
func (o *Listing) SetThumbnailUrlNil() {
	o.ThumbnailUrl.Set(nil)
}

// UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
func (o *Listing) UnsetThumbnailUrl() {
	o.ThumbnailUrl.Unset()
}

// GetAgent returns the Agent field value if set, zero value otherwise.
func (o *Listing) GetAgent() Group {
	if o == nil || o.Agent == nil {
		var ret Group
		return ret
	}
	return *o.Agent
}

// GetAgentOk returns a tuple with the Agent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listing) GetAgentOk() (*Group, bool) {
	if o == nil || o.Agent == nil {
		return nil, false
	}
	return o.Agent, true
}

// HasAgent returns a boolean if a field has been set.
func (o *Listing) HasAgent() bool {
	if o != nil && o.Agent != nil {
		return true
	}

	return false
}

// SetAgent gets a reference to the given Group and assigns it to the Agent field.
func (o *Listing) SetAgent(v Group) {
	o.Agent = &v
}

// GetCoAgent returns the CoAgent field value if set, zero value otherwise.
func (o *Listing) GetCoAgent() Group {
	if o == nil || o.CoAgent == nil {
		var ret Group
		return ret
	}
	return *o.CoAgent
}

// GetCoAgentOk returns a tuple with the CoAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listing) GetCoAgentOk() (*Group, bool) {
	if o == nil || o.CoAgent == nil {
		return nil, false
	}
	return o.CoAgent, true
}

// HasCoAgent returns a boolean if a field has been set.
func (o *Listing) HasCoAgent() bool {
	if o != nil && o.CoAgent != nil {
		return true
	}

	return false
}

// SetCoAgent gets a reference to the given Group and assigns it to the CoAgent field.
func (o *Listing) SetCoAgent(v Group) {
	o.CoAgent = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *Listing) GetImages() []Image {
	if o == nil || o.Images == nil {
		var ret []Image
		return ret
	}
	return *o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listing) GetImagesOk() (*[]Image, bool) {
	if o == nil || o.Images == nil {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *Listing) HasImages() bool {
	if o != nil && o.Images != nil {
		return true
	}

	return false
}

// SetImages gets a reference to the given []Image and assigns it to the Images field.
func (o *Listing) SetImages(v []Image) {
	o.Images = &v
}

// GetVideos returns the Videos field value if set, zero value otherwise.
func (o *Listing) GetVideos() []Video {
	if o == nil || o.Videos == nil {
		var ret []Video
		return ret
	}
	return *o.Videos
}

// GetVideosOk returns a tuple with the Videos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listing) GetVideosOk() (*[]Video, bool) {
	if o == nil || o.Videos == nil {
		return nil, false
	}
	return o.Videos, true
}

// HasVideos returns a boolean if a field has been set.
func (o *Listing) HasVideos() bool {
	if o != nil && o.Videos != nil {
		return true
	}

	return false
}

// SetVideos gets a reference to the given []Video and assigns it to the Videos field.
func (o *Listing) SetVideos(v []Video) {
	o.Videos = &v
}

// GetFloorPlans returns the FloorPlans field value if set, zero value otherwise.
func (o *Listing) GetFloorPlans() []FloorPlan {
	if o == nil || o.FloorPlans == nil {
		var ret []FloorPlan
		return ret
	}
	return *o.FloorPlans
}

// GetFloorPlansOk returns a tuple with the FloorPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listing) GetFloorPlansOk() (*[]FloorPlan, bool) {
	if o == nil || o.FloorPlans == nil {
		return nil, false
	}
	return o.FloorPlans, true
}

// HasFloorPlans returns a boolean if a field has been set.
func (o *Listing) HasFloorPlans() bool {
	if o != nil && o.FloorPlans != nil {
		return true
	}

	return false
}

// SetFloorPlans gets a reference to the given []FloorPlan and assigns it to the FloorPlans field.
func (o *Listing) SetFloorPlans(v []FloorPlan) {
	o.FloorPlans = &v
}

// GetPropertyWebsites returns the PropertyWebsites field value if set, zero value otherwise.
func (o *Listing) GetPropertyWebsites() PropertyWebsites {
	if o == nil || o.PropertyWebsites == nil {
		var ret PropertyWebsites
		return ret
	}
	return *o.PropertyWebsites
}

// GetPropertyWebsitesOk returns a tuple with the PropertyWebsites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listing) GetPropertyWebsitesOk() (*PropertyWebsites, bool) {
	if o == nil || o.PropertyWebsites == nil {
		return nil, false
	}
	return o.PropertyWebsites, true
}

// HasPropertyWebsites returns a boolean if a field has been set.
func (o *Listing) HasPropertyWebsites() bool {
	if o != nil && o.PropertyWebsites != nil {
		return true
	}

	return false
}

// SetPropertyWebsites gets a reference to the given PropertyWebsites and assigns it to the PropertyWebsites field.
func (o *Listing) SetPropertyWebsites(v PropertyWebsites) {
	o.PropertyWebsites = &v
}

// GetInteractiveContent returns the InteractiveContent field value if set, zero value otherwise.
func (o *Listing) GetInteractiveContent() []InteractiveContent {
	if o == nil || o.InteractiveContent == nil {
		var ret []InteractiveContent
		return ret
	}
	return *o.InteractiveContent
}

// GetInteractiveContentOk returns a tuple with the InteractiveContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listing) GetInteractiveContentOk() (*[]InteractiveContent, bool) {
	if o == nil || o.InteractiveContent == nil {
		return nil, false
	}
	return o.InteractiveContent, true
}

// HasInteractiveContent returns a boolean if a field has been set.
func (o *Listing) HasInteractiveContent() bool {
	if o != nil && o.InteractiveContent != nil {
		return true
	}

	return false
}

// SetInteractiveContent gets a reference to the given []InteractiveContent and assigns it to the InteractiveContent field.
func (o *Listing) SetInteractiveContent(v []InteractiveContent) {
	o.InteractiveContent = &v
}

// GetPropertyDetails returns the PropertyDetails field value if set, zero value otherwise.
func (o *Listing) GetPropertyDetails() PropertyDetails {
	if o == nil || o.PropertyDetails == nil {
		var ret PropertyDetails
		return ret
	}
	return *o.PropertyDetails
}

// GetPropertyDetailsOk returns a tuple with the PropertyDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listing) GetPropertyDetailsOk() (*PropertyDetails, bool) {
	if o == nil || o.PropertyDetails == nil {
		return nil, false
	}
	return o.PropertyDetails, true
}

// HasPropertyDetails returns a boolean if a field has been set.
func (o *Listing) HasPropertyDetails() bool {
	if o != nil && o.PropertyDetails != nil {
		return true
	}

	return false
}

// SetPropertyDetails gets a reference to the given PropertyDetails and assigns it to the PropertyDetails field.
func (o *Listing) SetPropertyDetails(v PropertyDetails) {
	o.PropertyDetails = &v
}

// GetDownloadsEnabled returns the DownloadsEnabled field value
func (o *Listing) GetDownloadsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DownloadsEnabled
}

// GetDownloadsEnabledOk returns a tuple with the DownloadsEnabled field value
// and a boolean to check if the value has been set.
func (o *Listing) GetDownloadsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DownloadsEnabled, true
}

// SetDownloadsEnabled sets field value
func (o *Listing) SetDownloadsEnabled(v bool) {
	o.DownloadsEnabled = v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *Listing) GetOrders() []Order {
	if o == nil || o.Orders == nil {
		var ret []Order
		return ret
	}
	return *o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listing) GetOrdersOk() (*[]Order, bool) {
	if o == nil || o.Orders == nil {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *Listing) HasOrders() bool {
	if o != nil && o.Orders != nil {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []Order and assigns it to the Orders field.
func (o *Listing) SetOrders(v []Order) {
	o.Orders = &v
}

func (o Listing) MarshalJSON() ([]byte, error) {
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
	if o.Agent != nil {
		toSerialize["agent"] = o.Agent
	}
	if o.CoAgent != nil {
		toSerialize["co_agent"] = o.CoAgent
	}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	if o.Videos != nil {
		toSerialize["videos"] = o.Videos
	}
	if o.FloorPlans != nil {
		toSerialize["floor_plans"] = o.FloorPlans
	}
	if o.PropertyWebsites != nil {
		toSerialize["property_websites"] = o.PropertyWebsites
	}
	if o.InteractiveContent != nil {
		toSerialize["interactive_content"] = o.InteractiveContent
	}
	if o.PropertyDetails != nil {
		toSerialize["property_details"] = o.PropertyDetails
	}
	if true {
		toSerialize["downloads_enabled"] = o.DownloadsEnabled
	}
	if o.Orders != nil {
		toSerialize["orders"] = o.Orders
	}
	return json.Marshal(toSerialize)
}

type NullableListing struct {
	value *Listing
	isSet bool
}

func (v NullableListing) Get() *Listing {
	return v.value
}

func (v *NullableListing) Set(val *Listing) {
	v.value = val
	v.isSet = true
}

func (v NullableListing) IsSet() bool {
	return v.isSet
}

func (v *NullableListing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListing(val *Listing) *NullableListing {
	return &NullableListing{value: val, isSet: true}
}

func (v NullableListing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


