# OrderPostPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FulfillmentStatus** | Pointer to **NullableString** | The fulfillment status of the order. Defaults to \&quot;UNFULFILLED\&quot;. | [optional] 
**InternalNotes** | Pointer to **NullableString** | Internal notes that will be attached to the order. Viewable only by the team. | [optional] 
**AddressId** | Pointer to **string** | ID of the address to associate with the order. UUID Version 4. | [optional] 
**CustomerId** | Pointer to **string** | ID of the customer to associate with the order. UUID Version 4. | [optional] 
**Notify** | Pointer to **NullableBool** | Indicates if the customer and creator notifications should be sent when creating the order. Requires an address and customer to be set in order for the notifications to be sent. | [optional] 
**LockDownloadForPayment** | Pointer to **NullableBool** | Indicates if the downloads for the attached listing should be locked while there is an outstanding balance on the order. | [optional] 
**AllowPaymentsBeforeFulfillment** | Pointer to **NullableBool** | Indicates if the order will allow payments from the customer before the order is marked as fulfilled. | [optional] 

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
### GetInternalNotes

`func (o *OrderPostPayload) GetInternalNotes() string`

GetInternalNotes returns the InternalNotes field if non-nil, zero value otherwise.

### GetInternalNotesOk

`func (o *OrderPostPayload) GetInternalNotesOk() (*string, bool)`

GetInternalNotesOk returns a tuple with the InternalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNotes

`func (o *OrderPostPayload) SetInternalNotes(v string)`

SetInternalNotes sets InternalNotes field to given value.

### HasInternalNotes

`func (o *OrderPostPayload) HasInternalNotes() bool`

HasInternalNotes returns a boolean if a field has been set.

### SetInternalNotesNil

`func (o *OrderPostPayload) SetInternalNotesNil(b bool)`

 SetInternalNotesNil sets the value for InternalNotes to be an explicit nil

### UnsetInternalNotes
`func (o *OrderPostPayload) UnsetInternalNotes()`

UnsetInternalNotes ensures that no value is present for InternalNotes, not even an explicit nil
### GetAddressId

`func (o *OrderPostPayload) GetAddressId() string`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *OrderPostPayload) GetAddressIdOk() (*string, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *OrderPostPayload) SetAddressId(v string)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *OrderPostPayload) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetCustomerId

`func (o *OrderPostPayload) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *OrderPostPayload) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *OrderPostPayload) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *OrderPostPayload) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetNotify

`func (o *OrderPostPayload) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *OrderPostPayload) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *OrderPostPayload) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *OrderPostPayload) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### SetNotifyNil

`func (o *OrderPostPayload) SetNotifyNil(b bool)`

 SetNotifyNil sets the value for Notify to be an explicit nil

### UnsetNotify
`func (o *OrderPostPayload) UnsetNotify()`

UnsetNotify ensures that no value is present for Notify, not even an explicit nil
### GetLockDownloadForPayment

`func (o *OrderPostPayload) GetLockDownloadForPayment() bool`

GetLockDownloadForPayment returns the LockDownloadForPayment field if non-nil, zero value otherwise.

### GetLockDownloadForPaymentOk

`func (o *OrderPostPayload) GetLockDownloadForPaymentOk() (*bool, bool)`

GetLockDownloadForPaymentOk returns a tuple with the LockDownloadForPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockDownloadForPayment

`func (o *OrderPostPayload) SetLockDownloadForPayment(v bool)`

SetLockDownloadForPayment sets LockDownloadForPayment field to given value.

### HasLockDownloadForPayment

`func (o *OrderPostPayload) HasLockDownloadForPayment() bool`

HasLockDownloadForPayment returns a boolean if a field has been set.

### SetLockDownloadForPaymentNil

`func (o *OrderPostPayload) SetLockDownloadForPaymentNil(b bool)`

 SetLockDownloadForPaymentNil sets the value for LockDownloadForPayment to be an explicit nil

### UnsetLockDownloadForPayment
`func (o *OrderPostPayload) UnsetLockDownloadForPayment()`

UnsetLockDownloadForPayment ensures that no value is present for LockDownloadForPayment, not even an explicit nil
### GetAllowPaymentsBeforeFulfillment

`func (o *OrderPostPayload) GetAllowPaymentsBeforeFulfillment() bool`

GetAllowPaymentsBeforeFulfillment returns the AllowPaymentsBeforeFulfillment field if non-nil, zero value otherwise.

### GetAllowPaymentsBeforeFulfillmentOk

`func (o *OrderPostPayload) GetAllowPaymentsBeforeFulfillmentOk() (*bool, bool)`

GetAllowPaymentsBeforeFulfillmentOk returns a tuple with the AllowPaymentsBeforeFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPaymentsBeforeFulfillment

`func (o *OrderPostPayload) SetAllowPaymentsBeforeFulfillment(v bool)`

SetAllowPaymentsBeforeFulfillment sets AllowPaymentsBeforeFulfillment field to given value.

### HasAllowPaymentsBeforeFulfillment

`func (o *OrderPostPayload) HasAllowPaymentsBeforeFulfillment() bool`

HasAllowPaymentsBeforeFulfillment returns a boolean if a field has been set.

### SetAllowPaymentsBeforeFulfillmentNil

`func (o *OrderPostPayload) SetAllowPaymentsBeforeFulfillmentNil(b bool)`

 SetAllowPaymentsBeforeFulfillmentNil sets the value for AllowPaymentsBeforeFulfillment to be an explicit nil

### UnsetAllowPaymentsBeforeFulfillment
`func (o *OrderPostPayload) UnsetAllowPaymentsBeforeFulfillment()`

UnsetAllowPaymentsBeforeFulfillment ensures that no value is present for AllowPaymentsBeforeFulfillment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


