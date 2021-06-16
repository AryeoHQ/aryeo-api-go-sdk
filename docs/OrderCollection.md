# OrderCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Order**](Order.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 

## Methods

### NewOrderCollection

`func NewOrderCollection() *OrderCollection`

NewOrderCollection instantiates a new OrderCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCollectionWithDefaults

`func NewOrderCollectionWithDefaults() *OrderCollection`

NewOrderCollectionWithDefaults instantiates a new OrderCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OrderCollection) GetData() []Order`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrderCollection) GetDataOk() (*[]Order, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrderCollection) SetData(v []Order)`

SetData sets Data field to given value.

### HasData

`func (o *OrderCollection) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *OrderCollection) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *OrderCollection) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *OrderCollection) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *OrderCollection) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *OrderCollection) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrderCollection) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrderCollection) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrderCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


