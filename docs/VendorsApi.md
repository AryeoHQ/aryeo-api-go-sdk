# \VendorsApi

All URIs are relative to *https://api.aryeo.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVendors**](VendorsApi.md#GetVendors) | **Get** /vendors | Get vendors available to a group.
[**GetVendorsId**](VendorsApi.md#GetVendorsId) | **Get** /vendors/{vendor_id} | Get vendors available to a group.



## GetVendors

> GroupCollection GetVendors(ctx).Include(include).Execute()

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
    include := "users" // string | Comma separated list of optional data to include in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VendorsApi.GetVendors(context.Background()).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VendorsApi.GetVendors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVendors`: GroupCollection
    fmt.Fprintf(os.Stdout, "Response from `VendorsApi.GetVendors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVendorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Comma separated list of optional data to include in the response. | 

### Return type

[**GroupCollection**](GroupCollection.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVendorsId

> GroupResource GetVendorsId(ctx, vendorId).Include(include).Execute()

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
    vendorId := TODO // string | The ID of the group that is associated as a vendor. UUID Version 4.
    include := "default_order_form" // string | Comma separated list of optional data to include in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VendorsApi.GetVendorsId(context.Background(), vendorId).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VendorsApi.GetVendorsId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVendorsId`: GroupResource
    fmt.Fprintf(os.Stdout, "Response from `VendorsApi.GetVendorsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vendorId** | [**string**](.md) | The ID of the group that is associated as a vendor. UUID Version 4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVendorsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | Comma separated list of optional data to include in the response. | 

### Return type

[**GroupResource**](GroupResource.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

