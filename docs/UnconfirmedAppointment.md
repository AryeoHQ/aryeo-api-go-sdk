# UnconfirmedAppointment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the appointment. | 
**Title** | Pointer to **NullableString** | The title of the appointment. | [optional] 
**Description** | Pointer to **NullableString** | The multi-line description of the appointment. | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**Users** | Pointer to [**[]User**](User.md) | Users attached to the appointment. | [optional] 
**PreferenceType** | Pointer to **NullableString** | The type of preferred scheduling information provided by a customer to aid in scheduling this appointment. | [optional] 
**PreferredStartAt** | Pointer to **NullableTime** | A preferred date and time (ISO 8601 format) for when the appointment could start. Only returned if unconfirmed appointment&#39;s preference type is TIME.  | [optional] 
**PreferredStartAtDay** | Pointer to **NullableString** | A preferred date (ISO 8601 format) for when the appointment could start. Only returned if unconfirmed appointment&#39;s preference type is TIME_OF_DAY.  | [optional] 
**PreferredStartAtTimeOfDay** | Pointer to **NullableString** | A preferred time of day for when the appointment could start. Only returned if unconfirmed appointment&#39;s preference type is TIME_OF_DAY.  | [optional] 
**Duration** | Pointer to **NullableInt32** | The estimated length of the appointment in minutes, if available. | [optional] 

## Methods

### NewUnconfirmedAppointment

`func NewUnconfirmedAppointment(id string, ) *UnconfirmedAppointment`

NewUnconfirmedAppointment instantiates a new UnconfirmedAppointment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnconfirmedAppointmentWithDefaults

`func NewUnconfirmedAppointmentWithDefaults() *UnconfirmedAppointment`

NewUnconfirmedAppointmentWithDefaults instantiates a new UnconfirmedAppointment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnconfirmedAppointment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnconfirmedAppointment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnconfirmedAppointment) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *UnconfirmedAppointment) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UnconfirmedAppointment) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UnconfirmedAppointment) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UnconfirmedAppointment) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UnconfirmedAppointment) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UnconfirmedAppointment) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *UnconfirmedAppointment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnconfirmedAppointment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnconfirmedAppointment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnconfirmedAppointment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UnconfirmedAppointment) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UnconfirmedAppointment) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetOrder

`func (o *UnconfirmedAppointment) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *UnconfirmedAppointment) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *UnconfirmedAppointment) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *UnconfirmedAppointment) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetUsers

`func (o *UnconfirmedAppointment) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UnconfirmedAppointment) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UnconfirmedAppointment) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UnconfirmedAppointment) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *UnconfirmedAppointment) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *UnconfirmedAppointment) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetPreferenceType

`func (o *UnconfirmedAppointment) GetPreferenceType() string`

GetPreferenceType returns the PreferenceType field if non-nil, zero value otherwise.

### GetPreferenceTypeOk

`func (o *UnconfirmedAppointment) GetPreferenceTypeOk() (*string, bool)`

GetPreferenceTypeOk returns a tuple with the PreferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferenceType

`func (o *UnconfirmedAppointment) SetPreferenceType(v string)`

SetPreferenceType sets PreferenceType field to given value.

### HasPreferenceType

`func (o *UnconfirmedAppointment) HasPreferenceType() bool`

HasPreferenceType returns a boolean if a field has been set.

### SetPreferenceTypeNil

`func (o *UnconfirmedAppointment) SetPreferenceTypeNil(b bool)`

 SetPreferenceTypeNil sets the value for PreferenceType to be an explicit nil

### UnsetPreferenceType
`func (o *UnconfirmedAppointment) UnsetPreferenceType()`

UnsetPreferenceType ensures that no value is present for PreferenceType, not even an explicit nil
### GetPreferredStartAt

`func (o *UnconfirmedAppointment) GetPreferredStartAt() time.Time`

GetPreferredStartAt returns the PreferredStartAt field if non-nil, zero value otherwise.

### GetPreferredStartAtOk

`func (o *UnconfirmedAppointment) GetPreferredStartAtOk() (*time.Time, bool)`

GetPreferredStartAtOk returns a tuple with the PreferredStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredStartAt

`func (o *UnconfirmedAppointment) SetPreferredStartAt(v time.Time)`

SetPreferredStartAt sets PreferredStartAt field to given value.

### HasPreferredStartAt

`func (o *UnconfirmedAppointment) HasPreferredStartAt() bool`

HasPreferredStartAt returns a boolean if a field has been set.

### SetPreferredStartAtNil

`func (o *UnconfirmedAppointment) SetPreferredStartAtNil(b bool)`

 SetPreferredStartAtNil sets the value for PreferredStartAt to be an explicit nil

### UnsetPreferredStartAt
`func (o *UnconfirmedAppointment) UnsetPreferredStartAt()`

UnsetPreferredStartAt ensures that no value is present for PreferredStartAt, not even an explicit nil
### GetPreferredStartAtDay

`func (o *UnconfirmedAppointment) GetPreferredStartAtDay() string`

GetPreferredStartAtDay returns the PreferredStartAtDay field if non-nil, zero value otherwise.

### GetPreferredStartAtDayOk

`func (o *UnconfirmedAppointment) GetPreferredStartAtDayOk() (*string, bool)`

GetPreferredStartAtDayOk returns a tuple with the PreferredStartAtDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredStartAtDay

`func (o *UnconfirmedAppointment) SetPreferredStartAtDay(v string)`

SetPreferredStartAtDay sets PreferredStartAtDay field to given value.

### HasPreferredStartAtDay

`func (o *UnconfirmedAppointment) HasPreferredStartAtDay() bool`

HasPreferredStartAtDay returns a boolean if a field has been set.

### SetPreferredStartAtDayNil

`func (o *UnconfirmedAppointment) SetPreferredStartAtDayNil(b bool)`

 SetPreferredStartAtDayNil sets the value for PreferredStartAtDay to be an explicit nil

### UnsetPreferredStartAtDay
`func (o *UnconfirmedAppointment) UnsetPreferredStartAtDay()`

UnsetPreferredStartAtDay ensures that no value is present for PreferredStartAtDay, not even an explicit nil
### GetPreferredStartAtTimeOfDay

`func (o *UnconfirmedAppointment) GetPreferredStartAtTimeOfDay() string`

GetPreferredStartAtTimeOfDay returns the PreferredStartAtTimeOfDay field if non-nil, zero value otherwise.

### GetPreferredStartAtTimeOfDayOk

`func (o *UnconfirmedAppointment) GetPreferredStartAtTimeOfDayOk() (*string, bool)`

GetPreferredStartAtTimeOfDayOk returns a tuple with the PreferredStartAtTimeOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredStartAtTimeOfDay

`func (o *UnconfirmedAppointment) SetPreferredStartAtTimeOfDay(v string)`

SetPreferredStartAtTimeOfDay sets PreferredStartAtTimeOfDay field to given value.

### HasPreferredStartAtTimeOfDay

`func (o *UnconfirmedAppointment) HasPreferredStartAtTimeOfDay() bool`

HasPreferredStartAtTimeOfDay returns a boolean if a field has been set.

### SetPreferredStartAtTimeOfDayNil

`func (o *UnconfirmedAppointment) SetPreferredStartAtTimeOfDayNil(b bool)`

 SetPreferredStartAtTimeOfDayNil sets the value for PreferredStartAtTimeOfDay to be an explicit nil

### UnsetPreferredStartAtTimeOfDay
`func (o *UnconfirmedAppointment) UnsetPreferredStartAtTimeOfDay()`

UnsetPreferredStartAtTimeOfDay ensures that no value is present for PreferredStartAtTimeOfDay, not even an explicit nil
### GetDuration

`func (o *UnconfirmedAppointment) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UnconfirmedAppointment) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UnconfirmedAppointment) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *UnconfirmedAppointment) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *UnconfirmedAppointment) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *UnconfirmedAppointment) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


