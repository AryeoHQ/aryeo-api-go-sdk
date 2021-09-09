# \AppointmentsApi

All URIs are relative to *https://api.aryeo.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAppointments**](AppointmentsApi.md#GetAppointments) | **Get** /appointments | List all appointments.
[**GetUnconfirmedAppointments**](AppointmentsApi.md#GetUnconfirmedAppointments) | **Get** /unconfirmed-appointments | List all unconfirmed appointments.
[**GetUnconfirmedAppointmentsId**](AppointmentsApi.md#GetUnconfirmedAppointmentsId) | **Get** /unconfirmed-appointments/{unconfirmed_appointment_id} | Retrieve an unconfirmed appointment.
[**PutAppointmentsAppointmentIdCancel**](AppointmentsApi.md#PutAppointmentsAppointmentIdCancel) | **Put** /appointments/{appointment_id}/cancel | Cancel an appointment.
[**PutAppointmentsAppointmentIdReschedule**](AppointmentsApi.md#PutAppointmentsAppointmentIdReschedule) | **Put** /appointments/{appointment_id}/reschedule | Reschedule an appointment.



## GetAppointments

> AppointmentCollection GetAppointments(ctx).Include(include).FilterTense(filterTense).FilterStartAtGte(filterStartAtGte).FilterStartAtLte(filterStartAtLte).FilterUserIds(filterUserIds).Sort(sort).PerPage(perPage).Page(page).Execute()

List all appointments.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    include := "order,users" // string | Comma separated list of optional data to include in the response. (optional)
    filterTense := "UPCOMING" // string | Return appointments that are upcoming or in the past. (optional)
    filterStartAtGte := time.Now() // time.Time | Return appointments where the start_at field is greater than or equal to this value. Effectively, appointments that start after this date. (optional)
    filterStartAtLte := time.Now() // time.Time | Return appointments where the start_at field is less than or equal to this value. Effectively, appointments that start before this date. (optional)
    filterUserIds := TODO // Array | The IDs of users whose appointments will be retrieved. UUID Version 4. (optional)
    sort := "-created_at" // string | Comma separated list of fields used for sorting. Placing a minus symbol in front of a field name sorts in descending order. Defaults to `-start_at`. (optional)
    perPage := "25" // string | The number of items per page. Defaults to 25. (optional)
    page := "2" // string | The requested page. Defaults to 1. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppointmentsApi.GetAppointments(context.Background()).Include(include).FilterTense(filterTense).FilterStartAtGte(filterStartAtGte).FilterStartAtLte(filterStartAtLte).FilterUserIds(filterUserIds).Sort(sort).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppointmentsApi.GetAppointments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppointments`: AppointmentCollection
    fmt.Fprintf(os.Stdout, "Response from `AppointmentsApi.GetAppointments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppointmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Comma separated list of optional data to include in the response. | 
 **filterTense** | **string** | Return appointments that are upcoming or in the past. | 
 **filterStartAtGte** | **time.Time** | Return appointments where the start_at field is greater than or equal to this value. Effectively, appointments that start after this date. | 
 **filterStartAtLte** | **time.Time** | Return appointments where the start_at field is less than or equal to this value. Effectively, appointments that start before this date. | 
 **filterUserIds** | [**Array**](Array.md) | The IDs of users whose appointments will be retrieved. UUID Version 4. | 
 **sort** | **string** | Comma separated list of fields used for sorting. Placing a minus symbol in front of a field name sorts in descending order. Defaults to &#x60;-start_at&#x60;. | 
 **perPage** | **string** | The number of items per page. Defaults to 25. | 
 **page** | **string** | The requested page. Defaults to 1. | 

### Return type

[**AppointmentCollection**](AppointmentCollection.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnconfirmedAppointments

> UnconfirmedAppointmentCollection GetUnconfirmedAppointments(ctx).Include(include).FilterUserIds(filterUserIds).Sort(sort).PerPage(perPage).Page(page).Execute()

List all unconfirmed appointments.



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
    include := "order,users" // string | Comma separated list of optional data to include in the response. (optional)
    filterUserIds := TODO // Array | The IDs of users whose appointments will be retrieved. UUID Version 4. (optional)
    sort := "-created_at" // string | Comma separated list of fields used for sorting. Placing a minus symbol in front of a field name sorts in descending order. Defaults to `-start_at`. (optional)
    perPage := "25" // string | The number of items per page. Defaults to 25. (optional)
    page := "2" // string | The requested page. Defaults to 1. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppointmentsApi.GetUnconfirmedAppointments(context.Background()).Include(include).FilterUserIds(filterUserIds).Sort(sort).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppointmentsApi.GetUnconfirmedAppointments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUnconfirmedAppointments`: UnconfirmedAppointmentCollection
    fmt.Fprintf(os.Stdout, "Response from `AppointmentsApi.GetUnconfirmedAppointments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUnconfirmedAppointmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **string** | Comma separated list of optional data to include in the response. | 
 **filterUserIds** | [**Array**](Array.md) | The IDs of users whose appointments will be retrieved. UUID Version 4. | 
 **sort** | **string** | Comma separated list of fields used for sorting. Placing a minus symbol in front of a field name sorts in descending order. Defaults to &#x60;-start_at&#x60;. | 
 **perPage** | **string** | The number of items per page. Defaults to 25. | 
 **page** | **string** | The requested page. Defaults to 1. | 

### Return type

[**UnconfirmedAppointmentCollection**](UnconfirmedAppointmentCollection.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnconfirmedAppointmentsId

> UnconfirmedAppointmentResource GetUnconfirmedAppointmentsId(ctx, unconfirmedAppointmentId).Include(include).Execute()

Retrieve an unconfirmed appointment.



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
    unconfirmedAppointmentId := TODO // string | The ID of an appointment.
    include := "order,users" // string | Comma separated list of optional data to include in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppointmentsApi.GetUnconfirmedAppointmentsId(context.Background(), unconfirmedAppointmentId).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppointmentsApi.GetUnconfirmedAppointmentsId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUnconfirmedAppointmentsId`: UnconfirmedAppointmentResource
    fmt.Fprintf(os.Stdout, "Response from `AppointmentsApi.GetUnconfirmedAppointmentsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**unconfirmedAppointmentId** | [**string**](.md) | The ID of an appointment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnconfirmedAppointmentsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **string** | Comma separated list of optional data to include in the response. | 

### Return type

[**UnconfirmedAppointmentResource**](UnconfirmedAppointmentResource.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAppointmentsAppointmentIdCancel

> AppointmentResource PutAppointmentsAppointmentIdCancel(ctx, appointmentId).AppointmentCancelPutPayload(appointmentCancelPutPayload).Execute()

Cancel an appointment.



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
    appointmentId := TODO // string | The ID of an appointment.
    appointmentCancelPutPayload := *openapiclient.NewAppointmentCancelPutPayload() // AppointmentCancelPutPayload |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppointmentsApi.PutAppointmentsAppointmentIdCancel(context.Background(), appointmentId).AppointmentCancelPutPayload(appointmentCancelPutPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppointmentsApi.PutAppointmentsAppointmentIdCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAppointmentsAppointmentIdCancel`: AppointmentResource
    fmt.Fprintf(os.Stdout, "Response from `AppointmentsApi.PutAppointmentsAppointmentIdCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appointmentId** | [**string**](.md) | The ID of an appointment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAppointmentsAppointmentIdCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appointmentCancelPutPayload** | [**AppointmentCancelPutPayload**](AppointmentCancelPutPayload.md) |  | 

### Return type

[**AppointmentResource**](AppointmentResource.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAppointmentsAppointmentIdReschedule

> AppointmentResource PutAppointmentsAppointmentIdReschedule(ctx, appointmentId).AppointmentReschedulePutPayload(appointmentReschedulePutPayload).Execute()

Reschedule an appointment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    appointmentId := TODO // string | The ID of an appointment.
    appointmentReschedulePutPayload := *openapiclient.NewAppointmentReschedulePutPayload(time.Now(), time.Now()) // AppointmentReschedulePutPayload |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppointmentsApi.PutAppointmentsAppointmentIdReschedule(context.Background(), appointmentId).AppointmentReschedulePutPayload(appointmentReschedulePutPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppointmentsApi.PutAppointmentsAppointmentIdReschedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAppointmentsAppointmentIdReschedule`: AppointmentResource
    fmt.Fprintf(os.Stdout, "Response from `AppointmentsApi.PutAppointmentsAppointmentIdReschedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appointmentId** | [**string**](.md) | The ID of an appointment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAppointmentsAppointmentIdRescheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appointmentReschedulePutPayload** | [**AppointmentReschedulePutPayload**](AppointmentReschedulePutPayload.md) |  | 

### Return type

[**AppointmentResource**](AppointmentResource.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

