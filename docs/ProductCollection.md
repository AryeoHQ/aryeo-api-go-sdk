# ProductCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | What was the state of the request? | 
**Data** | Pointer to [**[]Product**](Product.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 

## Methods

### NewProductCollection

`func NewProductCollection(status string, ) *ProductCollection`

NewProductCollection instantiates a new ProductCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductCollectionWithDefaults

`func NewProductCollectionWithDefaults() *ProductCollection`

NewProductCollectionWithDefaults instantiates a new ProductCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ProductCollection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductCollection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductCollection) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetData

`func (o *ProductCollection) GetData() []Product`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProductCollection) GetDataOk() (*[]Product, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProductCollection) SetData(v []Product)`

SetData sets Data field to given value.

### HasData

`func (o *ProductCollection) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ProductCollection) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ProductCollection) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMeta

`func (o *ProductCollection) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ProductCollection) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ProductCollection) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ProductCollection) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *ProductCollection) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProductCollection) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProductCollection) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProductCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


