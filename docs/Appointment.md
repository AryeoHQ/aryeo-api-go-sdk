# Appointment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the appointment. | 
**Status** | Pointer to **NullableString** | The status of the appointment. | [optional] 
**Title** | Pointer to **NullableString** | The title of the appointment. | [optional] 
**Description** | Pointer to **NullableString** | The multi-line description of the appointment. | [optional] 
**StartAt** | Pointer to **NullableTime** | The date and time (ISO 8601 format) when the appointment is set to start. | [optional] 
**EndAt** | Pointer to **NullableTime** | The date and time (ISO 8601 format) when the appointment is set to end. | [optional] 
**Duration** | Pointer to **NullableInt32** | The length of the appointment in minutes. | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Users** | Pointer to [**[]User**](User.md) | Users attached to the appointment. | [optional] 

## Methods

### NewAppointment

`func NewAppointment(id string, ) *Appointment`

NewAppointment instantiates a new Appointment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppointmentWithDefaults

`func NewAppointmentWithDefaults() *Appointment`

NewAppointmentWithDefaults instantiates a new Appointment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Appointment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Appointment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Appointment) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *Appointment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Appointment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Appointment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Appointment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Appointment) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Appointment) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTitle

`func (o *Appointment) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Appointment) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Appointment) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Appointment) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Appointment) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Appointment) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *Appointment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Appointment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Appointment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Appointment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Appointment) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Appointment) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStartAt

`func (o *Appointment) GetStartAt() time.Time`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *Appointment) GetStartAtOk() (*time.Time, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *Appointment) SetStartAt(v time.Time)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *Appointment) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### SetStartAtNil

`func (o *Appointment) SetStartAtNil(b bool)`

 SetStartAtNil sets the value for StartAt to be an explicit nil

### UnsetStartAt
`func (o *Appointment) UnsetStartAt()`

UnsetStartAt ensures that no value is present for StartAt, not even an explicit nil
### GetEndAt

`func (o *Appointment) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *Appointment) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *Appointment) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *Appointment) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### SetEndAtNil

`func (o *Appointment) SetEndAtNil(b bool)`

 SetEndAtNil sets the value for EndAt to be an explicit nil

### UnsetEndAt
`func (o *Appointment) UnsetEndAt()`

UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
### GetDuration

`func (o *Appointment) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Appointment) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Appointment) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Appointment) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *Appointment) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *Appointment) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetOrder

`func (o *Appointment) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Appointment) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Appointment) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Appointment) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetUsers

`func (o *Appointment) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Appointment) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Appointment) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Appointment) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *Appointment) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *Appointment) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


