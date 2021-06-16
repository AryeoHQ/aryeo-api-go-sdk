# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the order. | 
**DisplayId** | **int32** | A vanity id to be displayed for the order. For example, \&quot;Order #1000\&quot;. | 
**TotalPriceCents** | **int32** | The total price of the order given in cents. | 
**Currency** | [**Currency**](Currency.md) |  | 
**PaymentStatus** | **string** | The payment status of the order. | 
**PaymentUrl** | Pointer to **NullableString** | A URL for to pay for the order. | [optional] 
**Listing** | Pointer to [**PartialListing**](PartialListing.md) |  | [optional] 
**FulfillmentStatus** | **string** | The fulfillment status of the order. | 

## Methods

### NewOrder

`func NewOrder(id string, displayId int32, totalPriceCents int32, currency Currency, paymentStatus string, fulfillmentStatus string, ) *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Order) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Order) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Order) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayId

`func (o *Order) GetDisplayId() int32`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *Order) GetDisplayIdOk() (*int32, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *Order) SetDisplayId(v int32)`

SetDisplayId sets DisplayId field to given value.


### GetTotalPriceCents

`func (o *Order) GetTotalPriceCents() int32`

GetTotalPriceCents returns the TotalPriceCents field if non-nil, zero value otherwise.

### GetTotalPriceCentsOk

`func (o *Order) GetTotalPriceCentsOk() (*int32, bool)`

GetTotalPriceCentsOk returns a tuple with the TotalPriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPriceCents

`func (o *Order) SetTotalPriceCents(v int32)`

SetTotalPriceCents sets TotalPriceCents field to given value.


### GetCurrency

`func (o *Order) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Order) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Order) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetPaymentStatus

`func (o *Order) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *Order) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *Order) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.


### GetPaymentUrl

`func (o *Order) GetPaymentUrl() string`

GetPaymentUrl returns the PaymentUrl field if non-nil, zero value otherwise.

### GetPaymentUrlOk

`func (o *Order) GetPaymentUrlOk() (*string, bool)`

GetPaymentUrlOk returns a tuple with the PaymentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentUrl

`func (o *Order) SetPaymentUrl(v string)`

SetPaymentUrl sets PaymentUrl field to given value.

### HasPaymentUrl

`func (o *Order) HasPaymentUrl() bool`

HasPaymentUrl returns a boolean if a field has been set.

### SetPaymentUrlNil

`func (o *Order) SetPaymentUrlNil(b bool)`

 SetPaymentUrlNil sets the value for PaymentUrl to be an explicit nil

### UnsetPaymentUrl
`func (o *Order) UnsetPaymentUrl()`

UnsetPaymentUrl ensures that no value is present for PaymentUrl, not even an explicit nil
### GetListing

`func (o *Order) GetListing() PartialListing`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *Order) GetListingOk() (*PartialListing, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *Order) SetListing(v PartialListing)`

SetListing sets Listing field to given value.

### HasListing

`func (o *Order) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetFulfillmentStatus

`func (o *Order) GetFulfillmentStatus() string`

GetFulfillmentStatus returns the FulfillmentStatus field if non-nil, zero value otherwise.

### GetFulfillmentStatusOk

`func (o *Order) GetFulfillmentStatusOk() (*string, bool)`

GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentStatus

`func (o *Order) SetFulfillmentStatus(v string)`

SetFulfillmentStatus sets FulfillmentStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


