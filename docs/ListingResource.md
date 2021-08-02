# ListingResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | What was the state of the request? | 
**Data** | Pointer to [**Listing**](Listing.md) |  | [optional] 

## Methods

### NewListingResource

`func NewListingResource(status string, ) *ListingResource`

NewListingResource instantiates a new ListingResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingResourceWithDefaults

`func NewListingResourceWithDefaults() *ListingResource`

NewListingResourceWithDefaults instantiates a new ListingResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListingResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListingResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListingResource) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetData

`func (o *ListingResource) GetData() Listing`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListingResource) GetDataOk() (*Listing, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListingResource) SetData(v Listing)`

SetData sets Data field to given value.

### HasData

`func (o *ListingResource) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


