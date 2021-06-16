# OrderResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Order**](Order.md) |  | [optional] 

## Methods

### NewOrderResource

`func NewOrderResource() *OrderResource`

NewOrderResource instantiates a new OrderResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderResourceWithDefaults

`func NewOrderResourceWithDefaults() *OrderResource`

NewOrderResourceWithDefaults instantiates a new OrderResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OrderResource) GetData() Order`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrderResource) GetDataOk() (*Order, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrderResource) SetData(v Order)`

SetData sets Data field to given value.

### HasData

`func (o *OrderResource) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


