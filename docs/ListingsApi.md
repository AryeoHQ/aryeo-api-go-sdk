# \ListingsApi

All URIs are relative to *https://api.aryeo.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetListings**](ListingsApi.md#GetListings) | **Get** /listings | List all listings.
[**GetListingsId**](ListingsApi.md#GetListingsId) | **Get** /listings/{listing_id} | Retrieve a listing.



## GetListings

> ListingCollection GetListings(ctx).Include(include).FilterSearch(filterSearch).FilterAddress(filterAddress).FilterListAgent(filterListAgent).FilterStatus(filterStatus).FilterActive(filterActive).FilterPriceGte(filterPriceGte).FilterPriceLte(filterPriceLte).FilterSquareFeetGte(filterSquareFeetGte).FilterSquareFeetLte(filterSquareFeetLte).FilterBedroomsGte(filterBedroomsGte).FilterBedroomsLte(filterBedroomsLte).FilterBathroomsGte(filterBathroomsGte).FilterBathroomsLte(filterBathroomsLte).Sort(sort).PerPage(perPage).Page(page).Execute()

List all listings.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    include := "images,videos,orders" // string | Comma separated list of optional data to include in the response. (optional)
    filterSearch := "123 Main St" // string | Return listings that have fields matching this term. (optional)
    filterAddress := "123 Main St" // string | Return listings that have an address matching this term. (optional)
    filterListAgent := "John Doe" // string | Return listings that have a listing agent or co-listing agent matching this term. (optional)
    filterStatus := "FOR_SALE" // string | Return listings that have a certain status. (optional)
    filterActive := true // bool | Set as true to return listings that have an active status (e.g. active statuses include `COMING_SOON`, `FOR_SALE`, `FOR_LEASE`, `PENDING_SALE`, `PENDING_LEASE`, `SOLD`, `LEASED`). (optional)
    filterPriceGte := float32(100000) // float32 | Return listings where the price field is greater than or equal to this value. (optional)
    filterPriceLte := float32(4000000) // float32 | Return listings where the price field is less than or equal to this value. (optional)
    filterSquareFeetGte := float32(1000) // float32 | Return listings where the square feet field is greater than or equal to this value. (optional)
    filterSquareFeetLte := float32(5000) // float32 | Return listings where the square feet field is less than or equal to this value. (optional)
    filterBedroomsGte := int32(2) // int32 | Return listings where the bedrooms field is greater than or equal to this value. (optional)
    filterBedroomsLte := int32(4) // int32 | Return listings where the bedrooms field is less than or equal to this value. (optional)
    filterBathroomsGte := float32(2.5) // float32 | Return listings where the bathrooms field is greater than or equal to this value. (optional)
    filterBathroomsLte := float32(5) // float32 | Return listings where the bathrooms field is less than or equal to this value. (optional)
    sort := "-created_at" // string | Comma separated list of fields used for sorting. Placing a minus symbol in front of a field name sorts in descending order. Defaults to `-created_at`. (optional)
    perPage := "25" // string | The number of items per page. Defaults to 25. (optional)
    page := "2" // string | The requested page. Defaults to 1. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ListingsApi.GetListings(context.Background()).Include(include).FilterSearch(filterSearch).FilterAddress(filterAddress).FilterListAgent(filterListAgent).FilterStatus(filterStatus).FilterActive(filterActive).FilterPriceGte(filterPriceGte).FilterPriceLte(filterPriceLte).FilterSquareFeetGte(filterSquareFeetGte).FilterSquareFeetLte(filterSquareFeetLte).FilterBedroomsGte(filterBedroomsGte).FilterBedroomsLte(filterBedroomsLte).FilterBathroomsGte(filterBathroomsGte).FilterBathroomsLte(filterBathroomsLte).Sort(sort).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.GetListings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListings`: ListingCollection
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.GetListings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetListingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Comma separated list of optional data to include in the response. | 
 **filterSearch** | **string** | Return listings that have fields matching this term. | 
 **filterAddress** | **string** | Return listings that have an address matching this term. | 
 **filterListAgent** | **string** | Return listings that have a listing agent or co-listing agent matching this term. | 
 **filterStatus** | **string** | Return listings that have a certain status. | 
 **filterActive** | **bool** | Set as true to return listings that have an active status (e.g. active statuses include &#x60;COMING_SOON&#x60;, &#x60;FOR_SALE&#x60;, &#x60;FOR_LEASE&#x60;, &#x60;PENDING_SALE&#x60;, &#x60;PENDING_LEASE&#x60;, &#x60;SOLD&#x60;, &#x60;LEASED&#x60;). | 
 **filterPriceGte** | **float32** | Return listings where the price field is greater than or equal to this value. | 
 **filterPriceLte** | **float32** | Return listings where the price field is less than or equal to this value. | 
 **filterSquareFeetGte** | **float32** | Return listings where the square feet field is greater than or equal to this value. | 
 **filterSquareFeetLte** | **float32** | Return listings where the square feet field is less than or equal to this value. | 
 **filterBedroomsGte** | **int32** | Return listings where the bedrooms field is greater than or equal to this value. | 
 **filterBedroomsLte** | **int32** | Return listings where the bedrooms field is less than or equal to this value. | 
 **filterBathroomsGte** | **float32** | Return listings where the bathrooms field is greater than or equal to this value. | 
 **filterBathroomsLte** | **float32** | Return listings where the bathrooms field is less than or equal to this value. | 
 **sort** | **string** | Comma separated list of fields used for sorting. Placing a minus symbol in front of a field name sorts in descending order. Defaults to &#x60;-created_at&#x60;. | 
 **perPage** | **string** | The number of items per page. Defaults to 25. | 
 **page** | **string** | The requested page. Defaults to 1. | 

### Return type

[**ListingCollection**](ListingCollection.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListingsId

> ListingResource GetListingsId(ctx, listingId).Include(include).Execute()

Retrieve a listing.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    listingId := TODO // string | The ID of a listing. UUID Version 4.
    include := "images,videos,orders" // string | Comma separated list of optional data to include in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ListingsApi.GetListingsId(context.Background(), listingId).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.GetListingsId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListingsId`: ListingResource
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.GetListingsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listingId** | [**string**](.md) | The ID of a listing. UUID Version 4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | Comma separated list of optional data to include in the response. | 

### Return type

[**ListingResource**](ListingResource.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

