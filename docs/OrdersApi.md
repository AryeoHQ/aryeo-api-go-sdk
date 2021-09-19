# \OrdersApi

All URIs are relative to *https://api.aryeo.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrders**](OrdersApi.md#GetOrders) | **Get** /orders | List all orders.
[**GetOrdersId**](OrdersApi.md#GetOrdersId) | **Get** /orders/{order_id} | Retrieve an order.
[**GetProducts**](OrdersApi.md#GetProducts) | **Get** /products | List all products.
[**PostOrders**](OrdersApi.md#PostOrders) | **Post** /orders | Create an order.



## GetOrders

> OrderCollection GetOrders(ctx).Sort(sort).PerPage(perPage).Page(page).Execute()

List all orders.



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
    sort := "-created_at" // string | Comma separated list of fields used for sorting. Placing a minus symbol in front of a field name sorts in descending order. Defaults to `-created_at`. (optional)
    perPage := "25" // string | The number of items per page. Defaults to 25. (optional)
    page := "2" // string | The requested page. Defaults to 1. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.GetOrders(context.Background()).Sort(sort).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrders`: OrderCollection
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Comma separated list of fields used for sorting. Placing a minus symbol in front of a field name sorts in descending order. Defaults to &#x60;-created_at&#x60;. | 
 **perPage** | **string** | The number of items per page. Defaults to 25. | 
 **page** | **string** | The requested page. Defaults to 1. | 

### Return type

[**OrderCollection**](OrderCollection.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrdersId

> OrderResource GetOrdersId(ctx, orderId).Include(include).Execute()

Retrieve an order.



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
    orderId := TODO // string | The ID of an order. UUID Version 4.
    include := "items,appointments,unconfirmed_appointments" // string | Comma separated list of optional data to include in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.GetOrdersId(context.Background(), orderId).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrdersId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrdersId`: OrderResource
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrdersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | [**string**](.md) | The ID of an order. UUID Version 4. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | Comma separated list of optional data to include in the response. | 

### Return type

[**OrderResource**](OrderResource.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProducts

> ProductCollection GetProducts(ctx).Sort(sort).PerPage(perPage).Page(page).FilterSearch(filterSearch).FilterCategoryIds(filterCategoryIds).FilterType(filterType).Execute()

List all products.



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
    sort := "-created_at" // string | Comma separated list of fields used for sorting. Placing a minus symbol in front of a field name sorts in descending order. Defaults to `title`. (optional)
    perPage := "25" // string | The number of items per page. Defaults to 25. (optional)
    page := "2" // string | The requested page. Defaults to 1. (optional)
    filterSearch := "Photography" // string | Return products that have fields matching this term. (optional)
    filterCategoryIds := []string{"00000000-0000-4000-8000-000000000000"} // []string | Return products in the given categories. (optional)
    filterType := "MAIN" // string | Return products matching the given type. Allowed values are: MAIN, ADDON. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.GetProducts(context.Background()).Sort(sort).PerPage(perPage).Page(page).FilterSearch(filterSearch).FilterCategoryIds(filterCategoryIds).FilterType(filterType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProducts`: ProductCollection
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Comma separated list of fields used for sorting. Placing a minus symbol in front of a field name sorts in descending order. Defaults to &#x60;title&#x60;. | 
 **perPage** | **string** | The number of items per page. Defaults to 25. | 
 **page** | **string** | The requested page. Defaults to 1. | 
 **filterSearch** | **string** | Return products that have fields matching this term. | 
 **filterCategoryIds** | **[]string** | Return products in the given categories. | 
 **filterType** | **string** | Return products matching the given type. Allowed values are: MAIN, ADDON. | 

### Return type

[**ProductCollection**](ProductCollection.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrders

> OrderResource PostOrders(ctx).OrderPostPayload(orderPostPayload).Execute()

Create an order.



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
    orderPostPayload := *openapiclient.NewOrderPostPayload() // OrderPostPayload | OrderPostPayload (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.PostOrders(context.Background()).OrderPostPayload(orderPostPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.PostOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostOrders`: OrderResource
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.PostOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderPostPayload** | [**OrderPostPayload**](OrderPostPayload.md) | OrderPostPayload | 

### Return type

[**OrderResource**](OrderResource.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

