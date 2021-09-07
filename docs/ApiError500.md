# ApiError500

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | What was the state of the request? | 
**Message** | **string** | The error message. | 
**Code** | Pointer to **NullableInt32** | A numeric code corresponding to the error. | [optional] 

## Methods

### NewApiError500

`func NewApiError500(status string, message string, ) *ApiError500`

NewApiError500 instantiates a new ApiError500 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiError500WithDefaults

`func NewApiError500WithDefaults() *ApiError500`

NewApiError500WithDefaults instantiates a new ApiError500 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ApiError500) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiError500) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiError500) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *ApiError500) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiError500) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiError500) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *ApiError500) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiError500) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiError500) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiError500) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ApiError500) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ApiError500) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


