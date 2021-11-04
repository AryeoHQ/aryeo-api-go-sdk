/*
 * Aryeo
 *
 *
 * API version: 2021-06-17
 * Contact: jarrod@aryeo.com
 */

package aryeo

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"time"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// AppointmentsApiService AppointmentsApi service
type AppointmentsApiService service

type ApiGetAppointmentsRequest struct {
	ctx _context.Context
	ApiService *AppointmentsApiService
	include *string
	filterTense *string
	filterStartAtGte *time.Time
	filterStartAtLte *time.Time
	filterUserIds *[]string
	sort *string
	perPage *string
	page *string
}

// Comma separated list of optional data to include in the response.
func (r ApiGetAppointmentsRequest) Include(include string) ApiGetAppointmentsRequest {
	r.include = &include
	return r
}
// Return appointments that are upcoming or in the past.
func (r ApiGetAppointmentsRequest) FilterTense(filterTense string) ApiGetAppointmentsRequest {
	r.filterTense = &filterTense
	return r
}
// Return appointments where the start_at field is greater than or equal to this value. Effectively, appointments that start after this date.
func (r ApiGetAppointmentsRequest) FilterStartAtGte(filterStartAtGte time.Time) ApiGetAppointmentsRequest {
	r.filterStartAtGte = &filterStartAtGte
	return r
}
// Return appointments where the start_at field is less than or equal to this value. Effectively, appointments that start before this date.
func (r ApiGetAppointmentsRequest) FilterStartAtLte(filterStartAtLte time.Time) ApiGetAppointmentsRequest {
	r.filterStartAtLte = &filterStartAtLte
	return r
}
// The IDs of users whose appointments will be retrieved. UUID Version 4.
func (r ApiGetAppointmentsRequest) FilterUserIds(filterUserIds []string) ApiGetAppointmentsRequest {
	r.filterUserIds = &filterUserIds
	return r
}
// Comma separated list of fields used for sorting. Placing a minus symbol in front of a field name sorts in descending order. Defaults to &#x60;-start_at&#x60;.
func (r ApiGetAppointmentsRequest) Sort(sort string) ApiGetAppointmentsRequest {
	r.sort = &sort
	return r
}
// The number of items per page. Defaults to 25.
func (r ApiGetAppointmentsRequest) PerPage(perPage string) ApiGetAppointmentsRequest {
	r.perPage = &perPage
	return r
}
// The requested page. Defaults to 1.
func (r ApiGetAppointmentsRequest) Page(page string) ApiGetAppointmentsRequest {
	r.page = &page
	return r
}

func (r ApiGetAppointmentsRequest) Execute() (AppointmentCollection, *_nethttp.Response, error) {
	return r.ApiService.GetAppointmentsExecute(r)
}

/*
GetAppointments List all appointments.

List all appointments. By default, returns a list of appointments that have been scheduled and have not been canceled

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAppointmentsRequest
*/
func (a *AppointmentsApiService) GetAppointments(ctx _context.Context) ApiGetAppointmentsRequest {
	return ApiGetAppointmentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AppointmentCollection
func (a *AppointmentsApiService) GetAppointmentsExecute(r ApiGetAppointmentsRequest) (AppointmentCollection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AppointmentCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppointmentsApiService.GetAppointments")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/appointments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}
	if r.filterTense != nil {
		localVarQueryParams.Add("filter[tense]", parameterToString(*r.filterTense, ""))
	}
	if r.filterStartAtGte != nil {
		localVarQueryParams.Add("filter[start_at_gte]", parameterToString(*r.filterStartAtGte, ""))
	}
	if r.filterStartAtLte != nil {
		localVarQueryParams.Add("filter[start_at_lte]", parameterToString(*r.filterStartAtLte, ""))
	}
	if r.filterUserIds != nil {
		t := *r.filterUserIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[user_ids]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[user_ids]", parameterToString(t, "multi"))
		}
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ApiFail422
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAvailableDatesRequest struct {
	ctx _context.Context
	ApiService *AppointmentsApiService
	filterUserIds *[]string
	filterAppointmentId *string
	filterStartAt *time.Time
	filterEndAt *time.Time
	filterTimeframe *[]string
	duration *int32
	interval *int32
	timezone *string
	page *int32
	perPage *int32
}

// The IDs of users whose availability will be retrieved. UUID Version 4.
func (r ApiGetAvailableDatesRequest) FilterUserIds(filterUserIds []string) ApiGetAvailableDatesRequest {
	r.filterUserIds = &filterUserIds
	return r
}
// Appointment ID used to fetch availability for an existing order
func (r ApiGetAvailableDatesRequest) FilterAppointmentId(filterAppointmentId string) ApiGetAvailableDatesRequest {
	r.filterAppointmentId = &filterAppointmentId
	return r
}
// Returns availability after start_at
func (r ApiGetAvailableDatesRequest) FilterStartAt(filterStartAt time.Time) ApiGetAvailableDatesRequest {
	r.filterStartAt = &filterStartAt
	return r
}
// Returns availability before end_at
func (r ApiGetAvailableDatesRequest) FilterEndAt(filterEndAt time.Time) ApiGetAvailableDatesRequest {
	r.filterEndAt = &filterEndAt
	return r
}
// Returns availability for a specific timeframe. Used instead of start_at &amp; end_at
func (r ApiGetAvailableDatesRequest) FilterTimeframe(filterTimeframe []string) ApiGetAvailableDatesRequest {
	r.filterTimeframe = &filterTimeframe
	return r
}
// Duration of the event to schedule. Required if appointment_id isn&#39;t specified
func (r ApiGetAvailableDatesRequest) Duration(duration int32) ApiGetAvailableDatesRequest {
	r.duration = &duration
	return r
}
// Interval of bookable timeslots starting at x minutes on the hour . Required if appointment_id isn&#39;t specified
func (r ApiGetAvailableDatesRequest) Interval(interval int32) ApiGetAvailableDatesRequest {
	r.interval = &interval
	return r
}
// Timezone of the client. Localizes the available days
func (r ApiGetAvailableDatesRequest) Timezone(timezone string) ApiGetAvailableDatesRequest {
	r.timezone = &timezone
	return r
}
// The requested page of results
func (r ApiGetAvailableDatesRequest) Page(page int32) ApiGetAvailableDatesRequest {
	r.page = &page
	return r
}
// The number of results per page. Only applies when using a date range
func (r ApiGetAvailableDatesRequest) PerPage(perPage int32) ApiGetAvailableDatesRequest {
	r.perPage = &perPage
	return r
}

func (r ApiGetAvailableDatesRequest) Execute() (CalendarDayCollection, *_nethttp.Response, error) {
	return r.ApiService.GetAvailableDatesExecute(r)
}

/*
GetAvailableDates Fetch available days for a user or group

Fetch available calendar days for scheduling or rescheduling an appointment. Availability can be retrieved using a specific start & end date range, or using a timeframe. When using a timeframe, the page parameter can be used to flip through weeks, months, etc.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAvailableDatesRequest
*/
func (a *AppointmentsApiService) GetAvailableDates(ctx _context.Context) ApiGetAvailableDatesRequest {
	return ApiGetAvailableDatesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CalendarDayCollection
func (a *AppointmentsApiService) GetAvailableDatesExecute(r ApiGetAvailableDatesRequest) (CalendarDayCollection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CalendarDayCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppointmentsApiService.GetAvailableDates")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/scheduling/available-dates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.filterUserIds != nil {
		t := *r.filterUserIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[user_ids]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[user_ids]", parameterToString(t, "multi"))
		}
	}
	if r.filterAppointmentId != nil {
		localVarQueryParams.Add("filter[appointment_id]", parameterToString(*r.filterAppointmentId, ""))
	}
	if r.filterStartAt != nil {
		localVarQueryParams.Add("filter[start_at]", parameterToString(*r.filterStartAt, ""))
	}
	if r.filterEndAt != nil {
		localVarQueryParams.Add("filter[end_at]", parameterToString(*r.filterEndAt, ""))
	}
	if r.filterTimeframe != nil {
		t := *r.filterTimeframe
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[timeframe]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[timeframe]", parameterToString(t, "multi"))
		}
	}
	if r.duration != nil {
		localVarQueryParams.Add("duration", parameterToString(*r.duration, ""))
	}
	if r.interval != nil {
		localVarQueryParams.Add("interval", parameterToString(*r.interval, ""))
	}
	if r.timezone != nil {
		localVarQueryParams.Add("timezone", parameterToString(*r.timezone, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ApiFail422
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAvailableTimeslotsRequest struct {
	ctx _context.Context
	ApiService *AppointmentsApiService
	filterUserIds *[]string
	filterAppointmentId *string
	filterStartAt *time.Time
	filterEndAt *time.Time
	filterTimeframe *[]string
	duration *int32
	interval *int32
	page *int32
	perPage *int32
}

// The IDs of users whose appointments will be retrieved. UUID Version 4.
func (r ApiGetAvailableTimeslotsRequest) FilterUserIds(filterUserIds []string) ApiGetAvailableTimeslotsRequest {
	r.filterUserIds = &filterUserIds
	return r
}
// Appointment ID used to fetch availability for an existing order
func (r ApiGetAvailableTimeslotsRequest) FilterAppointmentId(filterAppointmentId string) ApiGetAvailableTimeslotsRequest {
	r.filterAppointmentId = &filterAppointmentId
	return r
}
// Returns availability after start_at
func (r ApiGetAvailableTimeslotsRequest) FilterStartAt(filterStartAt time.Time) ApiGetAvailableTimeslotsRequest {
	r.filterStartAt = &filterStartAt
	return r
}
// Returns availability before end_at
func (r ApiGetAvailableTimeslotsRequest) FilterEndAt(filterEndAt time.Time) ApiGetAvailableTimeslotsRequest {
	r.filterEndAt = &filterEndAt
	return r
}
// Returns availability for a specific timeframe. Used instead of start_at &amp; end_at
func (r ApiGetAvailableTimeslotsRequest) FilterTimeframe(filterTimeframe []string) ApiGetAvailableTimeslotsRequest {
	r.filterTimeframe = &filterTimeframe
	return r
}
// Duration of the event to schedule. Required if appointment_id isn&#39;t specified
func (r ApiGetAvailableTimeslotsRequest) Duration(duration int32) ApiGetAvailableTimeslotsRequest {
	r.duration = &duration
	return r
}
// Interval of bookable timeslots starting at x minutes on the hour . Required if appointment_id isn&#39;t specified
func (r ApiGetAvailableTimeslotsRequest) Interval(interval int32) ApiGetAvailableTimeslotsRequest {
	r.interval = &interval
	return r
}
// The requested page of results
func (r ApiGetAvailableTimeslotsRequest) Page(page int32) ApiGetAvailableTimeslotsRequest {
	r.page = &page
	return r
}
// The number of results per page. Only applies when using a date range
func (r ApiGetAvailableTimeslotsRequest) PerPage(perPage int32) ApiGetAvailableTimeslotsRequest {
	r.perPage = &perPage
	return r
}

func (r ApiGetAvailableTimeslotsRequest) Execute() (TimeslotCollection, *_nethttp.Response, error) {
	return r.ApiService.GetAvailableTimeslotsExecute(r)
}

/*
GetAvailableTimeslots Fetch available timeslots for a user or group

Fetch available timeslots for scheduling or rescheduling an appointment. Timeslots can be retrieved using a specific start & end date range, or using a timeframe. When using a timeframe, the page parameter can be used to flip through days or weeks.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAvailableTimeslotsRequest
*/
func (a *AppointmentsApiService) GetAvailableTimeslots(ctx _context.Context) ApiGetAvailableTimeslotsRequest {
	return ApiGetAvailableTimeslotsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TimeslotCollection
func (a *AppointmentsApiService) GetAvailableTimeslotsExecute(r ApiGetAvailableTimeslotsRequest) (TimeslotCollection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TimeslotCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppointmentsApiService.GetAvailableTimeslots")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/scheduling/available-timeslots"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.filterUserIds != nil {
		t := *r.filterUserIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[user_ids]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[user_ids]", parameterToString(t, "multi"))
		}
	}
	if r.filterAppointmentId != nil {
		localVarQueryParams.Add("filter[appointment_id]", parameterToString(*r.filterAppointmentId, ""))
	}
	if r.filterStartAt != nil {
		localVarQueryParams.Add("filter[start_at]", parameterToString(*r.filterStartAt, ""))
	}
	if r.filterEndAt != nil {
		localVarQueryParams.Add("filter[end_at]", parameterToString(*r.filterEndAt, ""))
	}
	if r.filterTimeframe != nil {
		t := *r.filterTimeframe
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[timeframe]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[timeframe]", parameterToString(t, "multi"))
		}
	}
	if r.duration != nil {
		localVarQueryParams.Add("duration", parameterToString(*r.duration, ""))
	}
	if r.interval != nil {
		localVarQueryParams.Add("interval", parameterToString(*r.interval, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ApiFail422
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetUnconfirmedAppointmentsRequest struct {
	ctx _context.Context
	ApiService *AppointmentsApiService
	include *string
	filterUserIds *[]string
	sort *string
	perPage *string
	page *string
}

// Comma separated list of optional data to include in the response.
func (r ApiGetUnconfirmedAppointmentsRequest) Include(include string) ApiGetUnconfirmedAppointmentsRequest {
	r.include = &include
	return r
}
// The IDs of users whose appointments will be retrieved. UUID Version 4.
func (r ApiGetUnconfirmedAppointmentsRequest) FilterUserIds(filterUserIds []string) ApiGetUnconfirmedAppointmentsRequest {
	r.filterUserIds = &filterUserIds
	return r
}
// Comma separated list of fields used for sorting. Placing a minus symbol in front of a field name sorts in descending order. Defaults to &#x60;-start_at&#x60;.
func (r ApiGetUnconfirmedAppointmentsRequest) Sort(sort string) ApiGetUnconfirmedAppointmentsRequest {
	r.sort = &sort
	return r
}
// The number of items per page. Defaults to 25.
func (r ApiGetUnconfirmedAppointmentsRequest) PerPage(perPage string) ApiGetUnconfirmedAppointmentsRequest {
	r.perPage = &perPage
	return r
}
// The requested page. Defaults to 1.
func (r ApiGetUnconfirmedAppointmentsRequest) Page(page string) ApiGetUnconfirmedAppointmentsRequest {
	r.page = &page
	return r
}

func (r ApiGetUnconfirmedAppointmentsRequest) Execute() (UnconfirmedAppointmentCollection, *_nethttp.Response, error) {
	return r.ApiService.GetUnconfirmedAppointmentsExecute(r)
}

/*
GetUnconfirmedAppointments List all unconfirmed appointments.

List all unconfirmed appointments. These are appointments that have not been scheduled. 

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUnconfirmedAppointmentsRequest
*/
func (a *AppointmentsApiService) GetUnconfirmedAppointments(ctx _context.Context) ApiGetUnconfirmedAppointmentsRequest {
	return ApiGetUnconfirmedAppointmentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UnconfirmedAppointmentCollection
func (a *AppointmentsApiService) GetUnconfirmedAppointmentsExecute(r ApiGetUnconfirmedAppointmentsRequest) (UnconfirmedAppointmentCollection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UnconfirmedAppointmentCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppointmentsApiService.GetUnconfirmedAppointments")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/unconfirmed-appointments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}
	if r.filterUserIds != nil {
		t := *r.filterUserIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("filter[user_ids]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("filter[user_ids]", parameterToString(t, "multi"))
		}
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ApiFail422
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetUnconfirmedAppointmentsIdRequest struct {
	ctx _context.Context
	ApiService *AppointmentsApiService
	unconfirmedAppointmentId string
	include *string
}

// Comma separated list of optional data to include in the response.
func (r ApiGetUnconfirmedAppointmentsIdRequest) Include(include string) ApiGetUnconfirmedAppointmentsIdRequest {
	r.include = &include
	return r
}

func (r ApiGetUnconfirmedAppointmentsIdRequest) Execute() (UnconfirmedAppointmentResource, *_nethttp.Response, error) {
	return r.ApiService.GetUnconfirmedAppointmentsIdExecute(r)
}

/*
GetUnconfirmedAppointmentsId Retrieve an unconfirmed appointment.

Retrieves the details of an unconfirmed appointment with the given ID.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param unconfirmedAppointmentId The ID of an appointment.
 @return ApiGetUnconfirmedAppointmentsIdRequest
*/
func (a *AppointmentsApiService) GetUnconfirmedAppointmentsId(ctx _context.Context, unconfirmedAppointmentId string) ApiGetUnconfirmedAppointmentsIdRequest {
	return ApiGetUnconfirmedAppointmentsIdRequest{
		ApiService: a,
		ctx: ctx,
		unconfirmedAppointmentId: unconfirmedAppointmentId,
	}
}

// Execute executes the request
//  @return UnconfirmedAppointmentResource
func (a *AppointmentsApiService) GetUnconfirmedAppointmentsIdExecute(r ApiGetUnconfirmedAppointmentsIdRequest) (UnconfirmedAppointmentResource, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  UnconfirmedAppointmentResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppointmentsApiService.GetUnconfirmedAppointmentsId")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/unconfirmed-appointments/{unconfirmed_appointment_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"unconfirmed_appointment_id"+"}", _neturl.PathEscape(parameterToString(r.unconfirmedAppointmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(r.unconfirmedAppointmentId) < 0 {
		return localVarReturnValue, nil, reportError("unconfirmedAppointmentId must have at least 0 elements")
	}
	if strlen(r.unconfirmedAppointmentId) > 255 {
		return localVarReturnValue, nil, reportError("unconfirmedAppointmentId must have less than 255 elements")
	}

	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ApiFail422
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPutAppointmentsAppointmentIdCancelRequest struct {
	ctx _context.Context
	ApiService *AppointmentsApiService
	appointmentId string
	appointmentCancelPutPayload *AppointmentCancelPutPayload
}

func (r ApiPutAppointmentsAppointmentIdCancelRequest) AppointmentCancelPutPayload(appointmentCancelPutPayload AppointmentCancelPutPayload) ApiPutAppointmentsAppointmentIdCancelRequest {
	r.appointmentCancelPutPayload = &appointmentCancelPutPayload
	return r
}

func (r ApiPutAppointmentsAppointmentIdCancelRequest) Execute() (AppointmentResource, *_nethttp.Response, error) {
	return r.ApiService.PutAppointmentsAppointmentIdCancelExecute(r)
}

/*
PutAppointmentsAppointmentIdCancel Cancel an appointment.

Cancel an appointment. The appointments order's customer can be optionally notified of this change. 

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appointmentId The ID of an appointment.
 @return ApiPutAppointmentsAppointmentIdCancelRequest
*/
func (a *AppointmentsApiService) PutAppointmentsAppointmentIdCancel(ctx _context.Context, appointmentId string) ApiPutAppointmentsAppointmentIdCancelRequest {
	return ApiPutAppointmentsAppointmentIdCancelRequest{
		ApiService: a,
		ctx: ctx,
		appointmentId: appointmentId,
	}
}

// Execute executes the request
//  @return AppointmentResource
func (a *AppointmentsApiService) PutAppointmentsAppointmentIdCancelExecute(r ApiPutAppointmentsAppointmentIdCancelRequest) (AppointmentResource, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AppointmentResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppointmentsApiService.PutAppointmentsAppointmentIdCancel")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/appointments/{appointment_id}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"appointment_id"+"}", _neturl.PathEscape(parameterToString(r.appointmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(r.appointmentId) < 0 {
		return localVarReturnValue, nil, reportError("appointmentId must have at least 0 elements")
	}
	if strlen(r.appointmentId) > 255 {
		return localVarReturnValue, nil, reportError("appointmentId must have less than 255 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.appointmentCancelPutPayload
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiError409
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ApiFail422
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPutAppointmentsAppointmentIdRescheduleRequest struct {
	ctx _context.Context
	ApiService *AppointmentsApiService
	appointmentId string
	appointmentReschedulePutPayload *AppointmentReschedulePutPayload
}

func (r ApiPutAppointmentsAppointmentIdRescheduleRequest) AppointmentReschedulePutPayload(appointmentReschedulePutPayload AppointmentReschedulePutPayload) ApiPutAppointmentsAppointmentIdRescheduleRequest {
	r.appointmentReschedulePutPayload = &appointmentReschedulePutPayload
	return r
}

func (r ApiPutAppointmentsAppointmentIdRescheduleRequest) Execute() (AppointmentResource, *_nethttp.Response, error) {
	return r.ApiService.PutAppointmentsAppointmentIdRescheduleExecute(r)
}

/*
PutAppointmentsAppointmentIdReschedule Reschedule an appointment.

Reschedule an appointment. The appointments order's customer can be optionally notified of this change. 

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appointmentId The ID of an appointment.
 @return ApiPutAppointmentsAppointmentIdRescheduleRequest
*/
func (a *AppointmentsApiService) PutAppointmentsAppointmentIdReschedule(ctx _context.Context, appointmentId string) ApiPutAppointmentsAppointmentIdRescheduleRequest {
	return ApiPutAppointmentsAppointmentIdRescheduleRequest{
		ApiService: a,
		ctx: ctx,
		appointmentId: appointmentId,
	}
}

// Execute executes the request
//  @return AppointmentResource
func (a *AppointmentsApiService) PutAppointmentsAppointmentIdRescheduleExecute(r ApiPutAppointmentsAppointmentIdRescheduleRequest) (AppointmentResource, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AppointmentResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppointmentsApiService.PutAppointmentsAppointmentIdReschedule")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/appointments/{appointment_id}/reschedule"
	localVarPath = strings.Replace(localVarPath, "{"+"appointment_id"+"}", _neturl.PathEscape(parameterToString(r.appointmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(r.appointmentId) < 0 {
		return localVarReturnValue, nil, reportError("appointmentId must have at least 0 elements")
	}
	if strlen(r.appointmentId) > 255 {
		return localVarReturnValue, nil, reportError("appointmentId must have less than 255 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.appointmentReschedulePutPayload
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiError409
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ApiFail422
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
