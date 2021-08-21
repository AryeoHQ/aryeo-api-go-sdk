# UnconfirmedAppointmentCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | What was the state of the request? | 
**Data** | Pointer to [**[]UnconfirmedAppointment**](UnconfirmedAppointment.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 
**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] 

## Methods

### NewUnconfirmedAppointmentCollection

`func NewUnconfirmedAppointmentCollection(status string, ) *UnconfirmedAppointmentCollection`

NewUnconfirmedAppointmentCollection instantiates a new UnconfirmedAppointmentCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnconfirmedAppointmentCollectionWithDefaults

`func NewUnconfirmedAppointmentCollectionWithDefaults() *UnconfirmedAppointmentCollection`

NewUnconfirmedAppointmentCollectionWithDefaults instantiates a new UnconfirmedAppointmentCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UnconfirmedAppointmentCollection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnconfirmedAppointmentCollection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnconfirmedAppointmentCollection) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetData

`func (o *UnconfirmedAppointmentCollection) GetData() []UnconfirmedAppointment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UnconfirmedAppointmentCollection) GetDataOk() (*[]UnconfirmedAppointment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UnconfirmedAppointmentCollection) SetData(v []UnconfirmedAppointment)`

SetData sets Data field to given value.

### HasData

`func (o *UnconfirmedAppointmentCollection) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *UnconfirmedAppointmentCollection) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *UnconfirmedAppointmentCollection) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMeta

`func (o *UnconfirmedAppointmentCollection) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UnconfirmedAppointmentCollection) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UnconfirmedAppointmentCollection) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UnconfirmedAppointmentCollection) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *UnconfirmedAppointmentCollection) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UnconfirmedAppointmentCollection) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UnconfirmedAppointmentCollection) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UnconfirmedAppointmentCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


