# AppointmentCancelPutPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifyCustomer** | Pointer to **NullableBool** | Sends a notification to the appointment&#39;s order&#39;s customer that the appointment was canceled. | [optional] 

## Methods

### NewAppointmentCancelPutPayload

`func NewAppointmentCancelPutPayload() *AppointmentCancelPutPayload`

NewAppointmentCancelPutPayload instantiates a new AppointmentCancelPutPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppointmentCancelPutPayloadWithDefaults

`func NewAppointmentCancelPutPayloadWithDefaults() *AppointmentCancelPutPayload`

NewAppointmentCancelPutPayloadWithDefaults instantiates a new AppointmentCancelPutPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifyCustomer

`func (o *AppointmentCancelPutPayload) GetNotifyCustomer() bool`

GetNotifyCustomer returns the NotifyCustomer field if non-nil, zero value otherwise.

### GetNotifyCustomerOk

`func (o *AppointmentCancelPutPayload) GetNotifyCustomerOk() (*bool, bool)`

GetNotifyCustomerOk returns a tuple with the NotifyCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCustomer

`func (o *AppointmentCancelPutPayload) SetNotifyCustomer(v bool)`

SetNotifyCustomer sets NotifyCustomer field to given value.

### HasNotifyCustomer

`func (o *AppointmentCancelPutPayload) HasNotifyCustomer() bool`

HasNotifyCustomer returns a boolean if a field has been set.

### SetNotifyCustomerNil

`func (o *AppointmentCancelPutPayload) SetNotifyCustomerNil(b bool)`

 SetNotifyCustomerNil sets the value for NotifyCustomer to be an explicit nil

### UnsetNotifyCustomer
`func (o *AppointmentCancelPutPayload) UnsetNotifyCustomer()`

UnsetNotifyCustomer ensures that no value is present for NotifyCustomer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


