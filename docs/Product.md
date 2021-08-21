# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the product. UUID Version 4. | 
**Title** | **string** | The title of the product. | 
**Description** | Pointer to **string** | The description of the product. | [optional] 
**Type** | **string** | The type of product. | 
**Variants** | Pointer to [**[]ProductVariant**](ProductVariant.md) |  | [optional] 
**Categories** | Pointer to [**[]ProductCategory**](ProductCategory.md) |  | [optional] 

## Methods

### NewProduct

`func NewProduct(id string, title string, type_ string, ) *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Product) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Product) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Product) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *Product) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Product) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Product) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Product) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Product) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Product) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Product) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *Product) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Product) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Product) SetType(v string)`

SetType sets Type field to given value.


### GetVariants

`func (o *Product) GetVariants() []ProductVariant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *Product) GetVariantsOk() (*[]ProductVariant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *Product) SetVariants(v []ProductVariant)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *Product) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetCategories

`func (o *Product) GetCategories() []ProductCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Product) GetCategoriesOk() (*[]ProductCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Product) SetCategories(v []ProductCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Product) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


