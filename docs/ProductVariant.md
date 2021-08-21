# ProductVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the product variant. UUID Version 4. | 
**Title** | **string** | The title of the product variant. | 
**Price** | **int32** | A positive integer in the smallest currency unit (that is, 100 cents for $1.00) representing the price of the product variant. | 

## Methods

### NewProductVariant

`func NewProductVariant(id string, title string, price int32, ) *ProductVariant`

NewProductVariant instantiates a new ProductVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductVariantWithDefaults

`func NewProductVariantWithDefaults() *ProductVariant`

NewProductVariantWithDefaults instantiates a new ProductVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductVariant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductVariant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductVariant) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *ProductVariant) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProductVariant) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProductVariant) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPrice

`func (o *ProductVariant) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductVariant) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductVariant) SetPrice(v int32)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


