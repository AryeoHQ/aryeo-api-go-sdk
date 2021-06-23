# \MarketingMaterialsApi

All URIs are relative to *https://api.aryeo.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutMarketingMaterialsTemplatesUuidPublish**](MarketingMaterialsApi.md#PutMarketingMaterialsTemplatesUuidPublish) | **Put** /marketing-materials/templates/{uuid}/publish | Publish a marketing material template.
[**PutMarketingMaterialsUuidPublish**](MarketingMaterialsApi.md#PutMarketingMaterialsUuidPublish) | **Put** /marketing-materials/{uuid}/publish | Publish a marketing material.



## PutMarketingMaterialsTemplatesUuidPublish

> PutMarketingMaterialsTemplatesUuidPublish(ctx, uuid).MarketingMaterialTemplatePublishPayload(marketingMaterialTemplatePublishPayload).Execute()

Publish a marketing material template.



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
    uuid := TODO // string | UUID of the marketing material template record.
    marketingMaterialTemplatePublishPayload := *openapiclient.NewMarketingMaterialTemplatePublishPayload() // MarketingMaterialTemplatePublishPayload |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketingMaterialsApi.PutMarketingMaterialsTemplatesUuidPublish(context.Background(), uuid).MarketingMaterialTemplatePublishPayload(marketingMaterialTemplatePublishPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingMaterialsApi.PutMarketingMaterialsTemplatesUuidPublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | UUID of the marketing material template record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMarketingMaterialsTemplatesUuidPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketingMaterialTemplatePublishPayload** | [**MarketingMaterialTemplatePublishPayload**](MarketingMaterialTemplatePublishPayload.md) |  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMarketingMaterialsUuidPublish

> PutMarketingMaterialsUuidPublish(ctx, uuid).MarketingMaterialPublishPayload(marketingMaterialPublishPayload).Execute()

Publish a marketing material.



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
    uuid := TODO // string | UUID of the marketing material record.
    marketingMaterialPublishPayload := *openapiclient.NewMarketingMaterialPublishPayload() // MarketingMaterialPublishPayload |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MarketingMaterialsApi.PutMarketingMaterialsUuidPublish(context.Background(), uuid).MarketingMaterialPublishPayload(marketingMaterialPublishPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingMaterialsApi.PutMarketingMaterialsUuidPublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | UUID of the marketing material record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMarketingMaterialsUuidPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketingMaterialPublishPayload** | [**MarketingMaterialPublishPayload**](MarketingMaterialPublishPayload.md) |  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

