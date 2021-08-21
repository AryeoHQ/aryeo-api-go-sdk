# AppointmentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | What was the state of the request? | 
**Data** | Pointer to [**Appointment**](Appointment.md) |  | [optional] 

## Methods

### NewAppointmentResource

`func NewAppointmentResource(status string, ) *AppointmentResource`

NewAppointmentResource instantiates a new AppointmentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppointmentResourceWithDefaults

`func NewAppointmentResourceWithDefaults() *AppointmentResource`

NewAppointmentResourceWithDefaults instantiates a new AppointmentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AppointmentResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppointmentResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppointmentResource) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetData

`func (o *AppointmentResource) GetData() Appointment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppointmentResource) GetDataOk() (*Appointment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppointmentResource) SetData(v Appointment)`

SetData sets Data field to given value.

### HasData

`func (o *AppointmentResource) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


