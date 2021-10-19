# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** | String representing the object’s type. Objects of the same type share the same schema. | [optional] 
**Id** | **string** | ID of the group. UUID Version 4. | 
**Type** | **string** | The type of the group. Can be CREATOR, AGENT, or BROKERAGE, and may dictate the attributes of the group returned. | 
**Name** | **string** | The name of the group. | 
**Email** | Pointer to **NullableString** | The email address of a group. | [optional] 
**Phone** | Pointer to **NullableString** | A phone number represented in whichever standards specified by the group, typically ###-###-#### (separated by hyphens). | [optional] 
**WebsiteUrl** | Pointer to **NullableString** | The website URL of a group. | [optional] 
**LogoUrl** | Pointer to **NullableString** | The logo URL of a group. | [optional] 
**AvatarUrl** | Pointer to **NullableString** | The profile image URL of a real estate agent. Only returned if group&#39;s type is AGENT. | [optional] 
**OfficeName** | Pointer to **NullableString** | The name of the brokerage or team of a real estate agent. Only returned if group&#39;s type is AGENT. | [optional] 
**LicenseNumber** | Pointer to **NullableString** | The license number of a real estate agent. Only returned if group&#39;s type is AGENT. | [optional] 
**SocialProfiles** | Pointer to [**SocialProfiles**](SocialProfiles.md) |  | [optional] 
**DefaultOrderForm** | Pointer to [**OrderForm**](OrderForm.md) |  | [optional] 
**OrderForms** | Pointer to [**[]OrderForm**](OrderForm.md) | An array of order forms a vendor group provides for placing orders. Only returned if group&#39;s type is CREATOR.  | [optional] 
**Owner** | Pointer to [**User**](User.md) |  | [optional] 
**Users** | Pointer to [**[]User**](User.md) | The Aryeo users associated with this group. | [optional] 
**IsBrokerageOrBrokerageAgent** | **bool** | Does this group represent a brokerage or an agent who belongs to a brokerage? | 

## Methods

### NewGroup

`func NewGroup(id string, type_ string, name string, isBrokerageOrBrokerageAgent bool, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Group) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Group) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Group) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Group) HasObject() bool`

HasObject returns a boolean if a field has been set.

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


### GetType

`func (o *Group) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Group) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Group) SetType(v string)`

SetType sets Type field to given value.


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
### GetWebsiteUrl

`func (o *Group) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *Group) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *Group) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *Group) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *Group) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *Group) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetLogoUrl

`func (o *Group) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *Group) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *Group) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *Group) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *Group) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *Group) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetAvatarUrl

`func (o *Group) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *Group) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *Group) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *Group) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### SetAvatarUrlNil

`func (o *Group) SetAvatarUrlNil(b bool)`

 SetAvatarUrlNil sets the value for AvatarUrl to be an explicit nil

### UnsetAvatarUrl
`func (o *Group) UnsetAvatarUrl()`

UnsetAvatarUrl ensures that no value is present for AvatarUrl, not even an explicit nil
### GetOfficeName

`func (o *Group) GetOfficeName() string`

GetOfficeName returns the OfficeName field if non-nil, zero value otherwise.

### GetOfficeNameOk

`func (o *Group) GetOfficeNameOk() (*string, bool)`

GetOfficeNameOk returns a tuple with the OfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeName

`func (o *Group) SetOfficeName(v string)`

SetOfficeName sets OfficeName field to given value.

### HasOfficeName

`func (o *Group) HasOfficeName() bool`

HasOfficeName returns a boolean if a field has been set.

### SetOfficeNameNil

`func (o *Group) SetOfficeNameNil(b bool)`

 SetOfficeNameNil sets the value for OfficeName to be an explicit nil

### UnsetOfficeName
`func (o *Group) UnsetOfficeName()`

UnsetOfficeName ensures that no value is present for OfficeName, not even an explicit nil
### GetLicenseNumber

`func (o *Group) GetLicenseNumber() string`

GetLicenseNumber returns the LicenseNumber field if non-nil, zero value otherwise.

### GetLicenseNumberOk

`func (o *Group) GetLicenseNumberOk() (*string, bool)`

GetLicenseNumberOk returns a tuple with the LicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseNumber

`func (o *Group) SetLicenseNumber(v string)`

SetLicenseNumber sets LicenseNumber field to given value.

### HasLicenseNumber

`func (o *Group) HasLicenseNumber() bool`

HasLicenseNumber returns a boolean if a field has been set.

### SetLicenseNumberNil

`func (o *Group) SetLicenseNumberNil(b bool)`

 SetLicenseNumberNil sets the value for LicenseNumber to be an explicit nil

### UnsetLicenseNumber
`func (o *Group) UnsetLicenseNumber()`

UnsetLicenseNumber ensures that no value is present for LicenseNumber, not even an explicit nil
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

### SetOrderFormsNil

`func (o *Group) SetOrderFormsNil(b bool)`

 SetOrderFormsNil sets the value for OrderForms to be an explicit nil

### UnsetOrderForms
`func (o *Group) UnsetOrderForms()`

UnsetOrderForms ensures that no value is present for OrderForms, not even an explicit nil
### GetOwner

`func (o *Group) GetOwner() User`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Group) GetOwnerOk() (*User, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Group) SetOwner(v User)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Group) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

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

### SetUsersNil

`func (o *Group) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *Group) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


