# \VendorsApi

All URIs are relative to *https://api.aryeo.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVendors**](VendorsApi.md#GetVendors) | **Get** /vendors | Get vendors available to a group.
[**GetVendorsSearch**](VendorsApi.md#GetVendorsSearch) | **Get** /vendors/search | Get vendors that can be added to the group&#39;s vendor list.



## GetVendors

> GroupCollection GetVendors(ctx).Execute()

Get vendors available to a group.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VendorsApi.GetVendors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VendorsApi.GetVendors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVendors`: GroupCollection
    fmt.Fprintf(os.Stdout, "Response from `VendorsApi.GetVendors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVendorsRequest struct via the builder pattern


### Return type

[**GroupCollection**](GroupCollection.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVendorsSearch

> GroupCollection GetVendorsSearch(ctx).Query(query).PerPage(perPage).Page(page).Execute()

Get vendors that can be added to the group's vendor list.



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
    query := "Demo Photography Company" // string | A search query. (optional)
    perPage := "25" // string | The number of items per page. Defaults to 25. (optional)
    page := "2" // string | The requested page. Defaults to 1. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VendorsApi.GetVendorsSearch(context.Background()).Query(query).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VendorsApi.GetVendorsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVendorsSearch`: GroupCollection
    fmt.Fprintf(os.Stdout, "Response from `VendorsApi.GetVendorsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVendorsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | A search query. | 
 **perPage** | **string** | The number of items per page. Defaults to 25. | 
 **page** | **string** | The requested page. Defaults to 1. | 

### Return type

[**GroupCollection**](GroupCollection.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

