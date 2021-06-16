# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the group. | 
**GroupType** | **string** | The type of group. | 
**Name** | **string** | The name of the group. | 
**Logo** | Pointer to **NullableString** | Group logo. | [optional] 
**Email** | Pointer to **NullableString** | Email. | [optional] 
**Phone** | Pointer to **NullableString** | Phone number. | [optional] 
**Website** | Pointer to **NullableString** | Website. | [optional] 
**IsBrokerageOrBrokerageAgent** | **bool** | Does this group represent a brokerage or an agent who belongs to a brokerage? | 
**SocialProfiles** | Pointer to [**SocialProfiles**](SocialProfiles.md) |  | [optional] 
**AgentProperties** | Pointer to [**GroupAgentProperties**](GroupAgentProperties.md) |  | [optional] 
**Users** | Pointer to [**[]User**](User.md) | users | [optional] 
**DefaultOrderForm** | Pointer to [**OrderForm**](OrderForm.md) |  | [optional] 
**OrderForms** | Pointer to [**[]OrderForm**](OrderForm.md) | An array of order forms. | [optional] 

## Methods

### NewGroup

`func NewGroup(id string, groupType string, name string, isBrokerageOrBrokerageAgent bool, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v string)`

SetId sets Id field to given value.


### GetGroupType

`func (o *Group) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *Group) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *Group) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.


### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.


### GetLogo

`func (o *Group) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *Group) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *Group) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *Group) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *Group) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *Group) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetEmail

`func (o *Group) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Group) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Group) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Group) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *Group) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Group) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *Group) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Group) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Group) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Group) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *Group) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *Group) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetWebsite

`func (o *Group) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Group) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Group) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *Group) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *Group) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *Group) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetIsBrokerageOrBrokerageAgent

`func (o *Group) GetIsBrokerageOrBrokerageAgent() bool`

GetIsBrokerageOrBrokerageAgent returns the IsBrokerageOrBrokerageAgent field if non-nil, zero value otherwise.

### GetIsBrokerageOrBrokerageAgentOk

`func (o *Group) GetIsBrokerageOrBrokerageAgentOk() (*bool, bool)`

GetIsBrokerageOrBrokerageAgentOk returns a tuple with the IsBrokerageOrBrokerageAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBrokerageOrBrokerageAgent

`func (o *Group) SetIsBrokerageOrBrokerageAgent(v bool)`

SetIsBrokerageOrBrokerageAgent sets IsBrokerageOrBrokerageAgent field to given value.


### GetSocialProfiles

`func (o *Group) GetSocialProfiles() SocialProfiles`

GetSocialProfiles returns the SocialProfiles field if non-nil, zero value otherwise.

### GetSocialProfilesOk

`func (o *Group) GetSocialProfilesOk() (*SocialProfiles, bool)`

GetSocialProfilesOk returns a tuple with the SocialProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProfiles

`func (o *Group) SetSocialProfiles(v SocialProfiles)`

SetSocialProfiles sets SocialProfiles field to given value.

### HasSocialProfiles

`func (o *Group) HasSocialProfiles() bool`

HasSocialProfiles returns a boolean if a field has been set.

### GetAgentProperties

`func (o *Group) GetAgentProperties() GroupAgentProperties`

GetAgentProperties returns the AgentProperties field if non-nil, zero value otherwise.

### GetAgentPropertiesOk

`func (o *Group) GetAgentPropertiesOk() (*GroupAgentProperties, bool)`

GetAgentPropertiesOk returns a tuple with the AgentProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentProperties

`func (o *Group) SetAgentProperties(v GroupAgentProperties)`

SetAgentProperties sets AgentProperties field to given value.

### HasAgentProperties

`func (o *Group) HasAgentProperties() bool`

HasAgentProperties returns a boolean if a field has been set.

### GetUsers

`func (o *Group) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Group) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Group) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Group) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetDefaultOrderForm

`func (o *Group) GetDefaultOrderForm() OrderForm`

GetDefaultOrderForm returns the DefaultOrderForm field if non-nil, zero value otherwise.

### GetDefaultOrderFormOk

`func (o *Group) GetDefaultOrderFormOk() (*OrderForm, bool)`

GetDefaultOrderFormOk returns a tuple with the DefaultOrderForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOrderForm

`func (o *Group) SetDefaultOrderForm(v OrderForm)`

SetDefaultOrderForm sets DefaultOrderForm field to given value.

### HasDefaultOrderForm

`func (o *Group) HasDefaultOrderForm() bool`

HasDefaultOrderForm returns a boolean if a field has been set.

### GetOrderForms

`func (o *Group) GetOrderForms() []OrderForm`

GetOrderForms returns the OrderForms field if non-nil, zero value otherwise.

### GetOrderFormsOk

`func (o *Group) GetOrderFormsOk() (*[]OrderForm, bool)`

GetOrderFormsOk returns a tuple with the OrderForms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderForms

`func (o *Group) SetOrderForms(v []OrderForm)`

SetOrderForms sets OrderForms field to given value.

### HasOrderForms

`func (o *Group) HasOrderForms() bool`

HasOrderForms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


