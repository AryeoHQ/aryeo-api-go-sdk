# GroupResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | What was the state of the request? | 
**Data** | Pointer to [**Group**](Group.md) |  | [optional] 

## Methods

### NewGroupResource

`func NewGroupResource(status string, ) *GroupResource`

NewGroupResource instantiates a new GroupResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupResourceWithDefaults

`func NewGroupResourceWithDefaults() *GroupResource`

NewGroupResourceWithDefaults instantiates a new GroupResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GroupResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GroupResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GroupResource) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetData

`func (o *GroupResource) GetData() Group`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GroupResource) GetDataOk() (*Group, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GroupResource) SetData(v Group)`

SetData sets Data field to given value.

### HasData

`func (o *GroupResource) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


