# GroupCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Group**](Group.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 

## Methods

### NewGroupCollection

`func NewGroupCollection() *GroupCollection`

NewGroupCollection instantiates a new GroupCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupCollectionWithDefaults

`func NewGroupCollectionWithDefaults() *GroupCollection`

NewGroupCollectionWithDefaults instantiates a new GroupCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GroupCollection) GetData() []Group`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GroupCollection) GetDataOk() (*[]Group, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GroupCollection) SetData(v []Group)`

SetData sets Data field to given value.

### HasData

`func (o *GroupCollection) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GroupCollection) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GroupCollection) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GroupCollection) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GroupCollection) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *GroupCollection) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GroupCollection) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GroupCollection) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GroupCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


