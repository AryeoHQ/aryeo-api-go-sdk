# UnconfirmedAppointmentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | What was the state of the request? | 
**Data** | Pointer to [**UnconfirmedAppointment**](UnconfirmedAppointment.md) |  | [optional] 

## Methods

### NewUnconfirmedAppointmentResource

`func NewUnconfirmedAppointmentResource(status string, ) *UnconfirmedAppointmentResource`

NewUnconfirmedAppointmentResource instantiates a new UnconfirmedAppointmentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnconfirmedAppointmentResourceWithDefaults

`func NewUnconfirmedAppointmentResourceWithDefaults() *UnconfirmedAppointmentResource`

NewUnconfirmedAppointmentResourceWithDefaults instantiates a new UnconfirmedAppointmentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UnconfirmedAppointmentResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnconfirmedAppointmentResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnconfirmedAppointmentResource) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetData

`func (o *UnconfirmedAppointmentResource) GetData() UnconfirmedAppointment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UnconfirmedAppointmentResource) GetDataOk() (*UnconfirmedAppointment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UnconfirmedAppointmentResource) SetData(v UnconfirmedAppointment)`

SetData sets Data field to given value.

### HasData

`func (o *UnconfirmedAppointmentResource) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


