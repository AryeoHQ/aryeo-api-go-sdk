# \ListingsApi

All URIs are relative to *https://api.aryeo.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetListings**](ListingsApi.md#GetListings) | **Get** /listings | Get the listings available to a group.
[**GetListingsId**](ListingsApi.md#GetListingsId) | **Get** /listings/{id} | Get information about a listing.



## GetListings

> PartialListingCollection GetListings(ctx).Query(query).PerPage(perPage).Page(page).Status(status).Price(price).Bathrooms(bathrooms).Bedrooms(bedrooms).Execute()

Get the listings available to a group.



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
    query := "thoroughbred trail" // string | A search query. (optional)
    perPage := "25" // string | The number of items per page. Defaults to 25. (optional)
    page := "2" // string | The requested page. Defaults to 1. (optional)
    status := "coming_soon" // string | The status you want to scope down to, example sold,  coming_soon,  for_sale, sold (optional)
    price := int32(100000) // int32 | The price value and greater will be returned. (optional)
    bathrooms := float32(3.5) // float32 | Number of bathrooms minimum. (optional)
    bedrooms := int32(4) // int32 | Number of bedrooms minimum. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ListingsApi.GetListings(context.Background()).Query(query).PerPage(perPage).Page(page).Status(status).Price(price).Bathrooms(bathrooms).Bedrooms(bedrooms).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.GetListings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListings`: PartialListingCollection
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.GetListings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetListingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | A search query. | 
 **perPage** | **string** | The number of items per page. Defaults to 25. | 
 **page** | **string** | The requested page. Defaults to 1. | 
 **status** | **string** | The status you want to scope down to, example sold,  coming_soon,  for_sale, sold | 
 **price** | **int32** | The price value and greater will be returned. | 
 **bathrooms** | **float32** | Number of bathrooms minimum. | 
 **bedrooms** | **int32** | Number of bedrooms minimum. | 

### Return type

[**PartialListingCollection**](PartialListingCollection.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListingsId

> ListingResource GetListingsId(ctx, id).Execute()

Get information about a listing.



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
    id := TODO // string | The UUID of a listing.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ListingsApi.GetListingsId(context.Background(), id).Execute()
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
**id** | [**string**](.md) | The UUID of a listing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListingResource**](ListingResource.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

