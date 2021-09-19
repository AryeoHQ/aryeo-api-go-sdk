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
    uuid := "UUID"

    configuration := aryeo.NewConfiguration()
    api_client := aryeo.NewAPIClient(configuration)
    auth := context.WithValue(context.Background(), aryeo.ContextAccessToken, "API_KEY")
    resp, r, err := api_client.ListingsApi.GetListingsId(auth, uuid).Execute()
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
*AppointmentsApi* | [**GetAppointments**](docs/AppointmentsApi.md#getappointments) | **Get** /appointments | List all appointments.
*AppointmentsApi* | [**GetUnconfirmedAppointments**](docs/AppointmentsApi.md#getunconfirmedappointments) | **Get** /unconfirmed-appointments | List all unconfirmed appointments.
*AppointmentsApi* | [**GetUnconfirmedAppointmentsId**](docs/AppointmentsApi.md#getunconfirmedappointmentsid) | **Get** /unconfirmed-appointments/{unconfirmed_appointment_id} | Retrieve an unconfirmed appointment.
*AppointmentsApi* | [**PutAppointmentsAppointmentIdCancel**](docs/AppointmentsApi.md#putappointmentsappointmentidcancel) | **Put** /appointments/{appointment_id}/cancel | Cancel an appointment.
*AppointmentsApi* | [**PutAppointmentsAppointmentIdReschedule**](docs/AppointmentsApi.md#putappointmentsappointmentidreschedule) | **Put** /appointments/{appointment_id}/reschedule | Reschedule an appointment.
*ListingsApi* | [**GetListings**](docs/ListingsApi.md#getlistings) | **Get** /listings | List all listings.
*ListingsApi* | [**GetListingsId**](docs/ListingsApi.md#getlistingsid) | **Get** /listings/{listing_id} | Retrieve a listing.
*OrdersApi* | [**GetOrders**](docs/OrdersApi.md#getorders) | **Get** /orders | List all orders.
*OrdersApi* | [**GetOrdersId**](docs/OrdersApi.md#getordersid) | **Get** /orders/{order_id} | Retrieve an order.
*OrdersApi* | [**GetProducts**](docs/OrdersApi.md#getproducts) | **Get** /products | List all products.
*OrdersApi* | [**PostOrders**](docs/OrdersApi.md#postorders) | **Post** /orders | Create an order.
*VendorsApi* | [**GetVendors**](docs/VendorsApi.md#getvendors) | **Get** /vendors | List all vendors.
*VendorsApi* | [**GetVendorsId**](docs/VendorsApi.md#getvendorsid) | **Get** /vendors/{vendor_id} | Retrieve a vendor.


## Documentation For Models

 - [Address](docs/Address.md)
 - [ApiError403](docs/ApiError403.md)
 - [ApiError404](docs/ApiError404.md)
 - [ApiError409](docs/ApiError409.md)
 - [ApiError500](docs/ApiError500.md)
 - [ApiFail422](docs/ApiFail422.md)
 - [Appointment](docs/Appointment.md)
 - [AppointmentCancelPutPayload](docs/AppointmentCancelPutPayload.md)
 - [AppointmentCollection](docs/AppointmentCollection.md)
 - [AppointmentReschedulePutPayload](docs/AppointmentReschedulePutPayload.md)
 - [AppointmentResource](docs/AppointmentResource.md)
 - [FloorPlan](docs/FloorPlan.md)
 - [Group](docs/Group.md)
 - [GroupCollection](docs/GroupCollection.md)
 - [GroupResource](docs/GroupResource.md)
 - [Image](docs/Image.md)
 - [InteractiveContent](docs/InteractiveContent.md)
 - [Listing](docs/Listing.md)
 - [ListingBuilding](docs/ListingBuilding.md)
 - [ListingCollection](docs/ListingCollection.md)
 - [ListingLot](docs/ListingLot.md)
 - [ListingPrice](docs/ListingPrice.md)
 - [ListingResource](docs/ListingResource.md)
 - [Order](docs/Order.md)
 - [OrderCollection](docs/OrderCollection.md)
 - [OrderForm](docs/OrderForm.md)
 - [OrderItem](docs/OrderItem.md)
 - [OrderPostPayload](docs/OrderPostPayload.md)
 - [OrderResource](docs/OrderResource.md)
 - [PaginationLinks](docs/PaginationLinks.md)
 - [PaginationMeta](docs/PaginationMeta.md)
 - [Product](docs/Product.md)
 - [ProductCategory](docs/ProductCategory.md)
 - [ProductCollection](docs/ProductCollection.md)
 - [ProductVariant](docs/ProductVariant.md)
 - [PropertyWebsite](docs/PropertyWebsite.md)
 - [SocialProfiles](docs/SocialProfiles.md)
 - [UnconfirmedAppointment](docs/UnconfirmedAppointment.md)
 - [UnconfirmedAppointmentCollection](docs/UnconfirmedAppointmentCollection.md)
 - [UnconfirmedAppointmentResource](docs/UnconfirmedAppointmentResource.md)
 - [User](docs/User.md)
 - [Video](docs/Video.md)


## Documentation For Authorization



### Token

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
