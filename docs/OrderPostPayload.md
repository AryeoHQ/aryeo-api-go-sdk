# OrderPostPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FulfillmentStatus** | Pointer to **NullableString** | The fulfillment status of the order. | [optional] 
**PaymentStatus** | Pointer to **NullableString** | The payment status of the order. | [optional] 
**ProductItems** | Pointer to [**[]ProductItem**](ProductItem.md) | product_items | [optional] 

## Methods

### NewOrderPostPayload

`func NewOrderPostPayload() *OrderPostPayload`

NewOrderPostPayload instantiates a new OrderPostPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderPostPayloadWithDefaults

`func NewOrderPostPayloadWithDefaults() *OrderPostPayload`

NewOrderPostPayloadWithDefaults instantiates a new OrderPostPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFulfillmentStatus

`func (o *OrderPostPayload) GetFulfillmentStatus() string`

GetFulfillmentStatus returns the FulfillmentStatus field if non-nil, zero value otherwise.

### GetFulfillmentStatusOk

`func (o *OrderPostPayload) GetFulfillmentStatusOk() (*string, bool)`

GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentStatus

`func (o *OrderPostPayload) SetFulfillmentStatus(v string)`

SetFulfillmentStatus sets FulfillmentStatus field to given value.

### HasFulfillmentStatus

`func (o *OrderPostPayload) HasFulfillmentStatus() bool`

HasFulfillmentStatus returns a boolean if a field has been set.

### SetFulfillmentStatusNil

`func (o *OrderPostPayload) SetFulfillmentStatusNil(b bool)`

 SetFulfillmentStatusNil sets the value for FulfillmentStatus to be an explicit nil

### UnsetFulfillmentStatus
`func (o *OrderPostPayload) UnsetFulfillmentStatus()`

UnsetFulfillmentStatus ensures that no value is present for FulfillmentStatus, not even an explicit nil
### GetPaymentStatus

`func (o *OrderPostPayload) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *OrderPostPayload) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *OrderPostPayload) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *OrderPostPayload) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### SetPaymentStatusNil

`func (o *OrderPostPayload) SetPaymentStatusNil(b bool)`

 SetPaymentStatusNil sets the value for PaymentStatus to be an explicit nil

### UnsetPaymentStatus
`func (o *OrderPostPayload) UnsetPaymentStatus()`

UnsetPaymentStatus ensures that no value is present for PaymentStatus, not even an explicit nil
### GetProductItems

`func (o *OrderPostPayload) GetProductItems() []ProductItem`

GetProductItems returns the ProductItems field if non-nil, zero value otherwise.

### GetProductItemsOk

`func (o *OrderPostPayload) GetProductItemsOk() (*[]ProductItem, bool)`

GetProductItemsOk returns a tuple with the ProductItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductItems

`func (o *OrderPostPayload) SetProductItems(v []ProductItem)`

SetProductItems sets ProductItems field to given value.

### HasProductItems

`func (o *OrderPostPayload) HasProductItems() bool`

HasProductItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


