# CalendarDay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Calendar day that has available timeslots. | 
**IsAvailable** | **bool** | Does the day have availability? | 

## Methods

### NewCalendarDay

`func NewCalendarDay(date string, isAvailable bool, ) *CalendarDay`

NewCalendarDay instantiates a new CalendarDay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarDayWithDefaults

`func NewCalendarDayWithDefaults() *CalendarDay`

NewCalendarDayWithDefaults instantiates a new CalendarDay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CalendarDay) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CalendarDay) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CalendarDay) SetDate(v string)`

SetDate sets Date field to given value.


### GetIsAvailable

`func (o *CalendarDay) GetIsAvailable() bool`

GetIsAvailable returns the IsAvailable field if non-nil, zero value otherwise.

### GetIsAvailableOk

`func (o *CalendarDay) GetIsAvailableOk() (*bool, bool)`

GetIsAvailableOk returns a tuple with the IsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailable

`func (o *CalendarDay) SetIsAvailable(v bool)`

SetIsAvailable sets IsAvailable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


