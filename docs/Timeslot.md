# Timeslot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartAt** | **string** | Start time of the available slot | 
**EndAt** | **string** | End time of the available slot | 
**Users** | Pointer to [**[]User**](User.md) |  | [optional] 

## Methods

### NewTimeslot

`func NewTimeslot(startAt string, endAt string, ) *Timeslot`

NewTimeslot instantiates a new Timeslot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeslotWithDefaults

`func NewTimeslotWithDefaults() *Timeslot`

NewTimeslotWithDefaults instantiates a new Timeslot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartAt

`func (o *Timeslot) GetStartAt() string`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *Timeslot) GetStartAtOk() (*string, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *Timeslot) SetStartAt(v string)`

SetStartAt sets StartAt field to given value.


### GetEndAt

`func (o *Timeslot) GetEndAt() string`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *Timeslot) GetEndAtOk() (*string, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *Timeslot) SetEndAt(v string)`

SetEndAt sets EndAt field to given value.


### GetUsers

`func (o *Timeslot) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Timeslot) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Timeslot) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Timeslot) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


