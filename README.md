# Aryeo SDK

## Introduction

This is an auto-generated client SDK for interfacing with the Aryeo API. We support a variety of languages and frameworks that are a great starting point for experimenting with the API. If there is an additional language or framework that you want to see supported, then please reach out and make a contribution!

<p align="center"> <a href="https://github.com/AryeoHQ/aryeo-api-dart-sdk"><img src="https://raw.githubusercontent.com/AryeoHQ/aryeo-api-docs/master/public/images/dart.svg" alt="Dart" width="44" style="padding:10px;border: 1px solid #d3d3d3;border-radius: 5px;margin:4px;"/></a> <a href="https://github.com/AryeoHQ/aryeo-api-go-sdk"><img src="https://raw.githubusercontent.com/AryeoHQ/aryeo-api-docs/master/public/images/go.svg" alt="Go" width="44" style="padding:10px;border: 1px solid #d3d3d3;border-radius: 5px;margin:4px;"/></a> <a href="https://github.com/AryeoHQ/aryeo-api-js-sdk"><img src="https://raw.githubusercontent.com/AryeoHQ/aryeo-api-docs/master/public/images/js.svg" alt="Node JS" width="44" style="padding:10px;border: 1px solid #d3d3d3;border-radius: 5px;margin:4px;"/></a> <a href="https://github.com/AryeoHQ/aryeo-api-php-sdk"><img src="https://raw.githubusercontent.com/AryeoHQ/aryeo-api-docs/master/public/images/php.svg" alt="PHP" width="44" style="padding:10px;border: 1px solid #d3d3d3;border-radius: 5px;margin:4px;"/></a> <a href="https://github.com/AryeoHQ/aryeo-api-ruby-sdk"><img src="https://raw.githubusercontent.com/AryeoHQ/aryeo-api-docs/master/public/images/ruby.svg" alt="Ruby" width="44" style="padding:10px;border: 1px solid #d3d3d3;border-radius: 5px;margin:4px;"/></a> <a href="https://github.com/AryeoHQ/aryeo-api-rust-sdk"><img src="https://raw.githubusercontent.com/AryeoHQ/aryeo-api-docs/master/public/images/rust.svg" alt="Rust" width="44" style="padding:10px;border: 1px solid #d3d3d3;border-radius: 5px;margin:4px;"/></a> <a href="https://github.com/AryeoHQ/aryeo-api-swift-sdk"><img src="https://raw.githubusercontent.com/AryeoHQ/aryeo-api-docs/master/public/images/swift.svg" alt="Swift" width="44" style="padding:10px;border: 1px solid #d3d3d3;border-radius: 5px;margin:4px;"/></a> </p>

## Authentication

To start using the Aryeo API, you will need to generate an API key from your group's developer settings. Then, make sure to provide your API key as a bearer token. Requests made without an API key will result in a `401 Unauthorized` error.

Example: `Authorization: Bearer {API_KEY}`

## Installation

Add the following block to `go.mod`:

```
require (
	github.com/AryeoHQ/aryeo-api-go-sdk
)
```

## Getting Started

```go
package main

import (
    "context"
    "fmt"
    "os"
    aryeo "github.com/AryeoHQ/aryeo-api-go-sdk"
)

func main() {
    id := "UUID"

    configuration := aryeo.NewConfiguration()
    api_client := aryeo.NewAPIClient(configuration)
    auth := context.WithValue(context.Background(), aryeo.ContextAccessToken, "API_KEY")
    resp, r, err := api_client.ListingsApi.GetListingsId(auth, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.GetListingsId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    
    fmt.Fprintf(os.Stdout, "%v\n", resp.Data.Address.FormattedAddress1)
}
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.aryeo.com/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ListingsApi* | [**GetListings**](docs/ListingsApi.md#getlistings) | **Get** /listings | Get the listings available to a group.
*ListingsApi* | [**GetListingsId**](docs/ListingsApi.md#getlistingsid) | **Get** /listings/{id} | Get information about a listing.
*MarketingMaterialsApi* | [**PutMarketingMaterialsUuidPublish**](docs/MarketingMaterialsApi.md#putmarketingmaterialsuuidpublish) | **Put** /marketing-materials/{uuid}/publish | Publish a marketing material.
*OrdersApi* | [**GetOrders**](docs/OrdersApi.md#getorders) | **Get** /orders | Get orders available to a group.
*OrdersApi* | [**PostOrders**](docs/OrdersApi.md#postorders) | **Post** /orders | Create an order.
*VendorsApi* | [**GetVendors**](docs/VendorsApi.md#getvendors) | **Get** /vendors | Get vendors available to a group.
*VendorsApi* | [**GetVendorsSearch**](docs/VendorsApi.md#getvendorssearch) | **Get** /vendors/search | Get vendors that can be added to the group&#39;s vendor list.


## Documentation For Models

 - [ApiError](docs/ApiError.md)
 - [Currency](docs/Currency.md)
 - [FloorPlan](docs/FloorPlan.md)
 - [Group](docs/Group.md)
 - [GroupAgentProperties](docs/GroupAgentProperties.md)
 - [GroupCollection](docs/GroupCollection.md)
 - [Image](docs/Image.md)
 - [InteractiveContent](docs/InteractiveContent.md)
 - [Listing](docs/Listing.md)
 - [ListingResource](docs/ListingResource.md)
 - [MarketingMaterialPublishPayload](docs/MarketingMaterialPublishPayload.md)
 - [Order](docs/Order.md)
 - [OrderCollection](docs/OrderCollection.md)
 - [OrderForm](docs/OrderForm.md)
 - [OrderPostPayload](docs/OrderPostPayload.md)
 - [OrderResource](docs/OrderResource.md)
 - [PaginationLinks](docs/PaginationLinks.md)
 - [PaginationMeta](docs/PaginationMeta.md)
 - [PartialAddress](docs/PartialAddress.md)
 - [PartialGroup](docs/PartialGroup.md)
 - [PartialListing](docs/PartialListing.md)
 - [PartialListingCollection](docs/PartialListingCollection.md)
 - [ProductItem](docs/ProductItem.md)
 - [PropertyDetails](docs/PropertyDetails.md)
 - [PropertyWebsites](docs/PropertyWebsites.md)
 - [SocialProfiles](docs/SocialProfiles.md)
 - [User](docs/User.md)
 - [Video](docs/Video.md)


## Documentation For Authorization



### JWT

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARERTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

jarrod@aryeo.com
