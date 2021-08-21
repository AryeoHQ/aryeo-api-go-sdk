# AppointmentReschedulePutPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartAt** | **NullableTime** | The new date and time (ISO 8601 format) when the appointment is set to start. | 
**EndAt** | **NullableTime** | The new date and time (ISO 8601 format) when the appointment is set to end. | 
**NotifyCustomer** | Pointer to **NullableBool** | Send a notification to the appointment&#39;s order&#39;s customer that the appointment was rescheduled. | [optional] 

## Methods

### NewAppointmentReschedulePutPayload

`func NewAppointmentReschedulePutPayload(startAt NullableTime, endAt NullableTime, ) *AppointmentReschedulePutPayload`

NewAppointmentReschedulePutPayload instantiates a new AppointmentReschedulePutPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppointmentReschedulePutPayloadWithDefaults

`func NewAppointmentReschedulePutPayloadWithDefaults() *AppointmentReschedulePutPayload`

NewAppointmentReschedulePutPayloadWithDefaults instantiates a new AppointmentReschedulePutPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartAt

`func (o *AppointmentReschedulePutPayload) GetStartAt() time.Time`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *AppointmentReschedulePutPayload) GetStartAtOk() (*time.Time, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *AppointmentReschedulePutPayload) SetStartAt(v time.Time)`

SetStartAt sets StartAt field to given value.


### SetStartAtNil

`func (o *AppointmentReschedulePutPayload) SetStartAtNil(b bool)`

 SetStartAtNil sets the value for StartAt to be an explicit nil

### UnsetStartAt
`func (o *AppointmentReschedulePutPayload) UnsetStartAt()`

UnsetStartAt ensures that no value is present for StartAt, not even an explicit nil
### GetEndAt

`func (o *AppointmentReschedulePutPayload) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *AppointmentReschedulePutPayload) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *AppointmentReschedulePutPayload) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.


### SetEndAtNil

`func (o *AppointmentReschedulePutPayload) SetEndAtNil(b bool)`

 SetEndAtNil sets the value for EndAt to be an explicit nil

### UnsetEndAt
`func (o *AppointmentReschedulePutPayload) UnsetEndAt()`

UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
### GetNotifyCustomer

`func (o *AppointmentReschedulePutPayload) GetNotifyCustomer() bool`

GetNotifyCustomer returns the NotifyCustomer field if non-nil, zero value otherwise.

### GetNotifyCustomerOk

`func (o *AppointmentReschedulePutPayload) GetNotifyCustomerOk() (*bool, bool)`

GetNotifyCustomerOk returns a tuple with the NotifyCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCustomer

`func (o *AppointmentReschedulePutPayload) SetNotifyCustomer(v bool)`

SetNotifyCustomer sets NotifyCustomer field to given value.

### HasNotifyCustomer

`func (o *AppointmentReschedulePutPayload) HasNotifyCustomer() bool`

HasNotifyCustomer returns a boolean if a field has been set.

### SetNotifyCustomerNil

`func (o *AppointmentReschedulePutPayload) SetNotifyCustomerNil(b bool)`

 SetNotifyCustomerNil sets the value for NotifyCustomer to be an explicit nil

### UnsetNotifyCustomer
`func (o *AppointmentReschedulePutPayload) UnsetNotifyCustomer()`

UnsetNotifyCustomer ensures that no value is present for NotifyCustomer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


