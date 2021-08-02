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

// Listing A real estate listing.
type Listing struct {
	// ID of the listing.
	Id string `json:"id"`
	Address Address `json:"address"`
	// The identifier for a listing on its local MLS. 
	MlsNumber NullableString `json:"mls_number,omitempty"`
	// General type of the listing, primarily categorizing its use case. Examples include residential and commercial. 
	Type NullableString `json:"type,omitempty"`
	// Further specifies the listing type. Examples include family residence and condominium.
	SubType NullableString `json:"sub_type,omitempty"`
	// Local, regional, or otherwise custom status for the listing used by the parties involved in the listing transaction. While variable, these statuses are typically mapped to the listing's standard status.
	Status NullableString `json:"status,omitempty"`
	// The status of the listing as it reflects the state of the contract between the listing agent and seller or an agreement with a buyer, including Active, Active Under Contract, Canceled, Closed, Expired, Pending, and Withdrawn.
	StandardStatus NullableString `json:"standard_status,omitempty"`
	// Description of the selling points of the building and/or land for sale. 
	Description NullableString `json:"description,omitempty"`
	Lot *ListingLot `json:"lot,omitempty"`
	Building *ListingBuilding `json:"building,omitempty"`
	Price *ListingPrice `json:"price,omitempty"`
	ListAgent *Group `json:"list_agent,omitempty"`
	CoListAgent *Group `json:"co_list_agent,omitempty"`
	// images
	Images *[]Image `json:"images,omitempty"`
	// videos
	Videos *[]Video `json:"videos,omitempty"`
	// floor_plans
	FloorPlans *[]FloorPlan `json:"floor_plans,omitempty"`
	// interactive_content
	InteractiveContent *[]InteractiveContent `json:"interactive_content,omitempty"`
	PropertyWebsite *PropertyWebsite `json:"property_website,omitempty"`
	// orders
	Orders *[]Order `json:"orders,omitempty"`
	// Are downloads enabled for this listing?
	DownloadsEnabled bool `json:"downloads_enabled"`
}

// NewListing instantiates a new Listing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListing(id string, address Address, downloadsEnabled bool) *Listing {
	this := Listing{}
	this.Id = id
	this.Address = address
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
func (o *Listing) GetAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *Listing) GetAddressOk() (*Address, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *Listing) SetAddress(v Address) {
	o.Address = v
}

// GetMlsNumber returns the MlsNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Listing) GetMlsNumber() string {
	if o == nil || o.MlsNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.MlsNumber.Get()
}

// GetMlsNumberOk returns a tuple with the MlsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Listing) GetMlsNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MlsNumber.Get(), o.MlsNumber.IsSet()
}

// HasMlsNumber returns a boolean if a field has been set.
func (o *Listing) HasMlsNumber() bool {
	if o != nil && o.MlsNumber.IsSet() {
		return true
	}

	return false
}

// SetMlsNumber gets a reference to the given NullableString and assigns it to the MlsNumber field.
func (o *Listing) SetMlsNumber(v string) {
	o.MlsNumber.Set(&v)
}
// SetMlsNumberNil sets the value for MlsNumber to be an explicit nil
func (o *Listing) SetMlsNumberNil() {
	o.MlsNumber.Set(nil)
}

// UnsetMlsNumber ensures that no value is present for MlsNumber, not even an explicit nil
func (o *Listing) UnsetMlsNumber() {
	o.MlsNumber.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Listing) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Listing) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Listing) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *Listing) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Listing) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Listing) UnsetType() {
	o.Type.Unset()
}

// GetSubType returns the SubType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Listing) GetSubType() string {
	if o == nil || o.SubType.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubType.Get()
}

// GetSubTypeOk returns a tuple with the SubType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Listing) GetSubTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubType.Get(), o.SubType.IsSet()
}

// HasSubType returns a boolean if a field has been set.
func (o *Listing) HasSubType() bool {
	if o != nil && o.SubType.IsSet() {
		return true
	}

	return false
}

// SetSubType gets a reference to the given NullableString and assigns it to the SubType field.
func (o *Listing) SetSubType(v string) {
	o.SubType.Set(&v)
}
// SetSubTypeNil sets the value for SubType to be an explicit nil
func (o *Listing) SetSubTypeNil() {
	o.SubType.Set(nil)
}

// UnsetSubType ensures that no value is present for SubType, not even an explicit nil
func (o *Listing) UnsetSubType() {
	o.SubType.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Listing) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Listing) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *Listing) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *Listing) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *Listing) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *Listing) UnsetStatus() {
	o.Status.Unset()
}

// GetStandardStatus returns the StandardStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Listing) GetStandardStatus() string {
	if o == nil || o.StandardStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.StandardStatus.Get()
}

// GetStandardStatusOk returns a tuple with the StandardStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Listing) GetStandardStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StandardStatus.Get(), o.StandardStatus.IsSet()
}

// HasStandardStatus returns a boolean if a field has been set.
func (o *Listing) HasStandardStatus() bool {
	if o != nil && o.StandardStatus.IsSet() {
		return true
	}

	return false
}

// SetStandardStatus gets a reference to the given NullableString and assigns it to the StandardStatus field.
func (o *Listing) SetStandardStatus(v string) {
	o.StandardStatus.Set(&v)
}
// SetStandardStatusNil sets the value for StandardStatus to be an explicit nil
func (o *Listing) SetStandardStatusNil() {
	o.StandardStatus.Set(nil)
}

// UnsetStandardStatus ensures that no value is present for StandardStatus, not even an explicit nil
func (o *Listing) UnsetStandardStatus() {
	o.StandardStatus.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Listing) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Listing) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Listing) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Listing) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Listing) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Listing) UnsetDescription() {
	o.Description.Unset()
}

// GetLot returns the Lot field value if set, zero value otherwise.
func (o *Listing) GetLot() ListingLot {
	if o == nil || o.Lot == nil {
		var ret ListingLot
		return ret
	}
	return *o.Lot
}

// GetLotOk returns a tuple with the Lot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listing) GetLotOk() (*ListingLot, bool) {
	if o == nil || o.Lot == nil {
		return nil, false
	}
	return o.Lot, true
}

// HasLot returns a boolean if a field has been set.
func (o *Listing) HasLot() bool {
	if o != nil && o.Lot != nil {
		return true
	}

	return false
}

// SetLot gets a reference to the given ListingLot and assigns it to the Lot field.
func (o *Listing) SetLot(v ListingLot) {
	o.Lot = &v
}

// GetBuilding returns the Building field value if set, zero value otherwise.
func (o *Listing) GetBuilding() ListingBuilding {
	if o == nil || o.Building == nil {
		var ret ListingBuilding
		return ret
	}
	return *o.Building
}

// GetBuildingOk returns a tuple with the Building field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listing) GetBuildingOk() (*ListingBuilding, bool) {
	if o == nil || o.Building == nil {
		return nil, false
	}
	return o.Building, true
}

// HasBuilding returns a boolean if a field has been set.
func (o *Listing) HasBuilding() bool {
	if o != nil && o.Building != nil {
		return true
	}

	return false
}

// SetBuilding gets a reference to the given ListingBuilding and assigns it to the Building field.
func (o *Listing) SetBuilding(v ListingBuilding) {
	o.Building = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *Listing) GetPrice() ListingPrice {
	if o == nil || o.Price == nil {
		var ret ListingPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listing) GetPriceOk() (*ListingPrice, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *Listing) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given ListingPrice and assigns it to the Price field.
func (o *Listing) SetPrice(v ListingPrice) {
	o.Price = &v
}

// GetListAgent returns the ListAgent field value if set, zero value otherwise.
func (o *Listing) GetListAgent() Group {
	if o == nil || o.ListAgent == nil {
		var ret Group
		return ret
	}
	return *o.ListAgent
}

// GetListAgentOk returns a tuple with the ListAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listing) GetListAgentOk() (*Group, bool) {
	if o == nil || o.ListAgent == nil {
		return nil, false
	}
	return o.ListAgent, true
}

// HasListAgent returns a boolean if a field has been set.
func (o *Listing) HasListAgent() bool {
	if o != nil && o.ListAgent != nil {
		return true
	}

	return false
}

// SetListAgent gets a reference to the given Group and assigns it to the ListAgent field.
func (o *Listing) SetListAgent(v Group) {
	o.ListAgent = &v
}

// GetCoListAgent returns the CoListAgent field value if set, zero value otherwise.
func (o *Listing) GetCoListAgent() Group {
	if o == nil || o.CoListAgent == nil {
		var ret Group
		return ret
	}
	return *o.CoListAgent
}

// GetCoListAgentOk returns a tuple with the CoListAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listing) GetCoListAgentOk() (*Group, bool) {
	if o == nil || o.CoListAgent == nil {
		return nil, false
	}
	return o.CoListAgent, true
}

// HasCoListAgent returns a boolean if a field has been set.
func (o *Listing) HasCoListAgent() bool {
	if o != nil && o.CoListAgent != nil {
		return true
	}

	return false
}

// SetCoListAgent gets a reference to the given Group and assigns it to the CoListAgent field.
func (o *Listing) SetCoListAgent(v Group) {
	o.CoListAgent = &v
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

// GetPropertyWebsite returns the PropertyWebsite field value if set, zero value otherwise.
func (o *Listing) GetPropertyWebsite() PropertyWebsite {
	if o == nil || o.PropertyWebsite == nil {
		var ret PropertyWebsite
		return ret
	}
	return *o.PropertyWebsite
}

// GetPropertyWebsiteOk returns a tuple with the PropertyWebsite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listing) GetPropertyWebsiteOk() (*PropertyWebsite, bool) {
	if o == nil || o.PropertyWebsite == nil {
		return nil, false
	}
	return o.PropertyWebsite, true
}

// HasPropertyWebsite returns a boolean if a field has been set.
func (o *Listing) HasPropertyWebsite() bool {
	if o != nil && o.PropertyWebsite != nil {
		return true
	}

	return false
}

// SetPropertyWebsite gets a reference to the given PropertyWebsite and assigns it to the PropertyWebsite field.
func (o *Listing) SetPropertyWebsite(v PropertyWebsite) {
	o.PropertyWebsite = &v
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

func (o Listing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if o.MlsNumber.IsSet() {
		toSerialize["mls_number"] = o.MlsNumber.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.SubType.IsSet() {
		toSerialize["sub_type"] = o.SubType.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.StandardStatus.IsSet() {
		toSerialize["standard_status"] = o.StandardStatus.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Lot != nil {
		toSerialize["lot"] = o.Lot
	}
	if o.Building != nil {
		toSerialize["building"] = o.Building
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.ListAgent != nil {
		toSerialize["list_agent"] = o.ListAgent
	}
	if o.CoListAgent != nil {
		toSerialize["co_list_agent"] = o.CoListAgent
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
	if o.InteractiveContent != nil {
		toSerialize["interactive_content"] = o.InteractiveContent
	}
	if o.PropertyWebsite != nil {
		toSerialize["property_website"] = o.PropertyWebsite
	}
	if o.Orders != nil {
		toSerialize["orders"] = o.Orders
	}
	if true {
		toSerialize["downloads_enabled"] = o.DownloadsEnabled
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


