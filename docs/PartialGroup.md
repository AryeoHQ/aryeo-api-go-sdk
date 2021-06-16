# PartialGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the group. | 
**GroupType** | **string** | The type of group. | 
**Name** | **string** | The name of the group. | 
**Logo** | Pointer to **NullableString** | Group logo. | [optional] 
**Email** | Pointer to **NullableString** | Email. | [optional] 
**Phone** | Pointer to **NullableString** | Phone number. | [optional] 
**IsBrokerageOrBrokerageAgent** | **bool** | Does this group represent a brokerage or an agent who belongs to a brokerage? | 

## Methods

### NewPartialGroup

`func NewPartialGroup(id string, groupType string, name string, isBrokerageOrBrokerageAgent bool, ) *PartialGroup`

NewPartialGroup instantiates a new PartialGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialGroupWithDefaults

`func NewPartialGroupWithDefaults() *PartialGroup`

NewPartialGroupWithDefaults instantiates a new PartialGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PartialGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartialGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartialGroup) SetId(v string)`

SetId sets Id field to given value.


### GetGroupType

`func (o *PartialGroup) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *PartialGroup) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *PartialGroup) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.


### GetName

`func (o *PartialGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartialGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartialGroup) SetName(v string)`

SetName sets Name field to given value.


### GetLogo

`func (o *PartialGroup) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *PartialGroup) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *PartialGroup) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *PartialGroup) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *PartialGroup) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *PartialGroup) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetEmail

`func (o *PartialGroup) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PartialGroup) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PartialGroup) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PartialGroup) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PartialGroup) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PartialGroup) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *PartialGroup) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PartialGroup) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PartialGroup) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PartialGroup) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *PartialGroup) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *PartialGroup) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetIsBrokerageOrBrokerageAgent

`func (o *PartialGroup) GetIsBrokerageOrBrokerageAgent() bool`

GetIsBrokerageOrBrokerageAgent returns the IsBrokerageOrBrokerageAgent field if non-nil, zero value otherwise.

### GetIsBrokerageOrBrokerageAgentOk

`func (o *PartialGroup) GetIsBrokerageOrBrokerageAgentOk() (*bool, bool)`

GetIsBrokerageOrBrokerageAgentOk returns a tuple with the IsBrokerageOrBrokerageAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBrokerageOrBrokerageAgent

`func (o *PartialGroup) SetIsBrokerageOrBrokerageAgent(v bool)`

SetIsBrokerageOrBrokerageAgent sets IsBrokerageOrBrokerageAgent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


