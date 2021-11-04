# \AppointmentsApi

All URIs are relative to *https://api.aryeo.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAppointments**](AppointmentsApi.md#GetAppointments) | **Get** /appointments | List all appointments.
[**GetAvailableDates**](AppointmentsApi.md#GetAvailableDates) | **Get** /scheduling/available-dates | Fetch available days for a user or group
[**GetAvailableTimeslots**](AppointmentsApi.md#GetAvailableTimeslots) | **Get** /scheduling/available-timeslots | Fetch available timeslots for a user or group
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
    filterUserIds := []string{"00000000-0000-4000-8000-000000000000"} // []string | The IDs of users whose appointments will be retrieved. UUID Version 4. (optional)
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
 **filterUserIds** | **[]string** | The IDs of users whose appointments will be retrieved. UUID Version 4. | 
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


## GetAvailableDates

> CalendarDayCollection GetAvailableDates(ctx).FilterUserIds(filterUserIds).FilterAppointmentId(filterAppointmentId).FilterStartAt(filterStartAt).FilterEndAt(filterEndAt).FilterTimeframe(filterTimeframe).Duration(duration).Interval(interval).Timezone(timezone).Page(page).PerPage(perPage).Execute()

Fetch available days for a user or group



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
    filterUserIds := []string{"00000000-0000-4000-8000-000000000000"} // []string | The IDs of users whose availability will be retrieved. UUID Version 4. (optional)
    filterAppointmentId := "00000000-0000-4000-8000-000000000000" // string | Appointment ID used to fetch availability for an existing order (optional)
    filterStartAt := time.Now() // time.Time | Returns availability after start_at (optional)
    filterEndAt := time.Now() // time.Time | Returns availability before end_at (optional)
    filterTimeframe := []string{"MONTH"} // []string | Returns availability for a specific timeframe. Used instead of start_at & end_at (optional)
    duration := int32(60) // int32 | Duration of the event to schedule. Required if appointment_id isn't specified (optional)
    interval := int32(15) // int32 | Interval of bookable timeslots starting at x minutes on the hour . Required if appointment_id isn't specified (optional)
    timezone := "2" // string | Timezone of the client. Localizes the available days (optional)
    page := int32(1) // int32 | The requested page of results (optional)
    perPage := int32(5) // int32 | The number of results per page. Only applies when using a date range (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppointmentsApi.GetAvailableDates(context.Background()).FilterUserIds(filterUserIds).FilterAppointmentId(filterAppointmentId).FilterStartAt(filterStartAt).FilterEndAt(filterEndAt).FilterTimeframe(filterTimeframe).Duration(duration).Interval(interval).Timezone(timezone).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppointmentsApi.GetAvailableDates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailableDates`: CalendarDayCollection
    fmt.Fprintf(os.Stdout, "Response from `AppointmentsApi.GetAvailableDates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableDatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterUserIds** | **[]string** | The IDs of users whose availability will be retrieved. UUID Version 4. | 
 **filterAppointmentId** | **string** | Appointment ID used to fetch availability for an existing order | 
 **filterStartAt** | **time.Time** | Returns availability after start_at | 
 **filterEndAt** | **time.Time** | Returns availability before end_at | 
 **filterTimeframe** | **[]string** | Returns availability for a specific timeframe. Used instead of start_at &amp; end_at | 
 **duration** | **int32** | Duration of the event to schedule. Required if appointment_id isn&#39;t specified | 
 **interval** | **int32** | Interval of bookable timeslots starting at x minutes on the hour . Required if appointment_id isn&#39;t specified | 
 **timezone** | **string** | Timezone of the client. Localizes the available days | 
 **page** | **int32** | The requested page of results | 
 **perPage** | **int32** | The number of results per page. Only applies when using a date range | 

### Return type

[**CalendarDayCollection**](CalendarDayCollection.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableTimeslots

> TimeslotCollection GetAvailableTimeslots(ctx).FilterUserIds(filterUserIds).FilterAppointmentId(filterAppointmentId).FilterStartAt(filterStartAt).FilterEndAt(filterEndAt).FilterTimeframe(filterTimeframe).Duration(duration).Interval(interval).Page(page).PerPage(perPage).Execute()

Fetch available timeslots for a user or group



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
    filterUserIds := []string{"00000000-0000-4000-8000-000000000000"} // []string | The IDs of users whose appointments will be retrieved. UUID Version 4. (optional)
    filterAppointmentId := "00000000-0000-4000-8000-000000000000" // string | Appointment ID used to fetch availability for an existing order (optional)
    filterStartAt := time.Now() // time.Time | Returns availability after start_at (optional)
    filterEndAt := time.Now() // time.Time | Returns availability before end_at (optional)
    filterTimeframe := []string{"MONTH"} // []string | Returns availability for a specific timeframe. Used instead of start_at & end_at (optional)
    duration := int32(60) // int32 | Duration of the event to schedule. Required if appointment_id isn't specified (optional)
    interval := int32(25) // int32 | Interval of bookable timeslots starting at x minutes on the hour . Required if appointment_id isn't specified (optional)
    page := int32(1) // int32 | The requested page of results (optional)
    perPage := int32(5) // int32 | The number of results per page. Only applies when using a date range (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppointmentsApi.GetAvailableTimeslots(context.Background()).FilterUserIds(filterUserIds).FilterAppointmentId(filterAppointmentId).FilterStartAt(filterStartAt).FilterEndAt(filterEndAt).FilterTimeframe(filterTimeframe).Duration(duration).Interval(interval).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppointmentsApi.GetAvailableTimeslots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailableTimeslots`: TimeslotCollection
    fmt.Fprintf(os.Stdout, "Response from `AppointmentsApi.GetAvailableTimeslots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableTimeslotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterUserIds** | **[]string** | The IDs of users whose appointments will be retrieved. UUID Version 4. | 
 **filterAppointmentId** | **string** | Appointment ID used to fetch availability for an existing order | 
 **filterStartAt** | **time.Time** | Returns availability after start_at | 
 **filterEndAt** | **time.Time** | Returns availability before end_at | 
 **filterTimeframe** | **[]string** | Returns availability for a specific timeframe. Used instead of start_at &amp; end_at | 
 **duration** | **int32** | Duration of the event to schedule. Required if appointment_id isn&#39;t specified | 
 **interval** | **int32** | Interval of bookable timeslots starting at x minutes on the hour . Required if appointment_id isn&#39;t specified | 
 **page** | **int32** | The requested page of results | 
 **perPage** | **int32** | The number of results per page. Only applies when using a date range | 

### Return type

[**TimeslotCollection**](TimeslotCollection.md)

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
    filterUserIds := []string{"00000000-0000-4000-8000-000000000000"} // []string | The IDs of users whose appointments will be retrieved. UUID Version 4. (optional)
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
 **filterUserIds** | **[]string** | The IDs of users whose appointments will be retrieved. UUID Version 4. | 
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

