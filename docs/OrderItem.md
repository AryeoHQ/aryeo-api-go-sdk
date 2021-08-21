# OrderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the item. UUID Version 4. | 
**Title** | Pointer to **string** | The title of the item. | [optional] 
**Description** | Pointer to **string** | The description of the item. | [optional] 
**Amount** | Pointer to **int32** | A positive integer in the smallest currency unit (that is, 100 cents for $1.00) representing the cost of a single instance of this item. This is multiplied by the quantity to calculate what was or will be charged. | [optional] 
**Quantity** | Pointer to **int32** | A positive integer representing the number of instances of this item that was or will be charged. | [optional] 

## Methods

### NewOrderItem

`func NewOrderItem(id string, ) *OrderItem`

NewOrderItem instantiates a new OrderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderItemWithDefaults

`func NewOrderItemWithDefaults() *OrderItem`

NewOrderItemWithDefaults instantiates a new OrderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderItem) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *OrderItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OrderItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OrderItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OrderItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *OrderItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrderItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrderItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrderItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmount

`func (o *OrderItem) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrderItem) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrderItem) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *OrderItem) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


