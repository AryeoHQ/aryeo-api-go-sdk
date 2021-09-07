# ApiError403

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | What was the state of the request? | 
**Message** | **string** | The error message. | 
**Code** | Pointer to **NullableInt32** | A numeric code corresponding to the error. | [optional] 

## Methods

### NewApiError403

`func NewApiError403(status string, message string, ) *ApiError403`

NewApiError403 instantiates a new ApiError403 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiError403WithDefaults

`func NewApiError403WithDefaults() *ApiError403`

NewApiError403WithDefaults instantiates a new ApiError403 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ApiError403) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiError403) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiError403) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *ApiError403) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiError403) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiError403) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *ApiError403) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiError403) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiError403) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiError403) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ApiError403) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ApiError403) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


