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
)

// Linger please
var (
	_ _context.Context
)

// ListingsApiService ListingsApi service
type ListingsApiService service

type ApiGetListingsRequest struct {
	ctx _context.Context
	ApiService *ListingsApiService
	include *string
	filterSearch *string
	filterAddress *string
	filterListAgent *string
	filterStatus *string
	filterActive *bool
	filterPriceGte *float32
	filterPriceLte *float32
	filterSquareFeetGte *float32
	filterSquareFeetLte *float32
	filterBedroomsGte *int32
	filterBedroomsLte *int32
	filterBathroomsGte *float32
	filterBathroomsLte *float32
	sort *string
	perPage *string
	page *string
}

// Comma separated list of optional data to include in the response.
func (r ApiGetListingsRequest) Include(include string) ApiGetListingsRequest {
	r.include = &include
	return r
}
// Return listings that have fields matching this term.
func (r ApiGetListingsRequest) FilterSearch(filterSearch string) ApiGetListingsRequest {
	r.filterSearch = &filterSearch
	return r
}
// Return listings that have an address matching this term.
func (r ApiGetListingsRequest) FilterAddress(filterAddress string) ApiGetListingsRequest {
	r.filterAddress = &filterAddress
	return r
}
// Return listings that have a listing agent or co-listing agent matching this term.
func (r ApiGetListingsRequest) FilterListAgent(filterListAgent string) ApiGetListingsRequest {
	r.filterListAgent = &filterListAgent
	return r
}
// Return listings that have a certain status.
func (r ApiGetListingsRequest) FilterStatus(filterStatus string) ApiGetListingsRequest {
	r.filterStatus = &filterStatus
	return r
}
// Set as true to return listings that have an active status (e.g. active statuses include &#x60;COMING_SOON&#x60;, &#x60;FOR_SALE&#x60;, &#x60;FOR_LEASE&#x60;, &#x60;PENDING_SALE&#x60;, &#x60;PENDING_LEASE&#x60;, &#x60;SOLD&#x60;, &#x60;LEASED&#x60;). 
func (r ApiGetListingsRequest) FilterActive(filterActive bool) ApiGetListingsRequest {
	r.filterActive = &filterActive
	return r
}
// Return listings where the price field is greater than or equal to this value.
func (r ApiGetListingsRequest) FilterPriceGte(filterPriceGte float32) ApiGetListingsRequest {
	r.filterPriceGte = &filterPriceGte
	return r
}
// Return listings where the price field is less than or equal to this value.
func (r ApiGetListingsRequest) FilterPriceLte(filterPriceLte float32) ApiGetListingsRequest {
	r.filterPriceLte = &filterPriceLte
	return r
}
// Return listings where the square feet field is greater than or equal to this value.
func (r ApiGetListingsRequest) FilterSquareFeetGte(filterSquareFeetGte float32) ApiGetListingsRequest {
	r.filterSquareFeetGte = &filterSquareFeetGte
	return r
}
// Return listings where the square feet field is less than or equal to this value.
func (r ApiGetListingsRequest) FilterSquareFeetLte(filterSquareFeetLte float32) ApiGetListingsRequest {
	r.filterSquareFeetLte = &filterSquareFeetLte
	return r
}
// Return listings where the bedrooms field is greater than or equal to this value.
func (r ApiGetListingsRequest) FilterBedroomsGte(filterBedroomsGte int32) ApiGetListingsRequest {
	r.filterBedroomsGte = &filterBedroomsGte
	return r
}
// Return listings where the bedrooms field is less than or equal to this value.
func (r ApiGetListingsRequest) FilterBedroomsLte(filterBedroomsLte int32) ApiGetListingsRequest {
	r.filterBedroomsLte = &filterBedroomsLte
	return r
}
// Return listings where the bathrooms field is greater than or equal to this value.
func (r ApiGetListingsRequest) FilterBathroomsGte(filterBathroomsGte float32) ApiGetListingsRequest {
	r.filterBathroomsGte = &filterBathroomsGte
	return r
}
// Return listings where the bathrooms field is less than or equal to this value.
func (r ApiGetListingsRequest) FilterBathroomsLte(filterBathroomsLte float32) ApiGetListingsRequest {
	r.filterBathroomsLte = &filterBathroomsLte
	return r
}
// Comma separated list of fields used for sorting. Placing a minus symbol in front of a field name sorts in descending order. Defaults to &#x60;-created_at&#x60;.
func (r ApiGetListingsRequest) Sort(sort string) ApiGetListingsRequest {
	r.sort = &sort
	return r
}
// The number of items per page. Defaults to 25.
func (r ApiGetListingsRequest) PerPage(perPage string) ApiGetListingsRequest {
	r.perPage = &perPage
	return r
}
// The requested page. Defaults to 1.
func (r ApiGetListingsRequest) Page(page string) ApiGetListingsRequest {
	r.page = &page
	return r
}

func (r ApiGetListingsRequest) Execute() (ListingCollection, *_nethttp.Response, error) {
	return r.ApiService.GetListingsExecute(r)
}

/*
GetListings List all listings.

Lists all listings available to a group.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetListingsRequest
*/
func (a *ListingsApiService) GetListings(ctx _context.Context) ApiGetListingsRequest {
	return ApiGetListingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListingCollection
func (a *ListingsApiService) GetListingsExecute(r ApiGetListingsRequest) (ListingCollection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListingCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ListingsApiService.GetListings")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/listings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.include != nil {
		localVarQueryParams.Add("include", parameterToString(*r.include, ""))
	}
	if r.filterSearch != nil {
		localVarQueryParams.Add("filter[search]", parameterToString(*r.filterSearch, ""))
	}
	if r.filterAddress != nil {
		localVarQueryParams.Add("filter[address]", parameterToString(*r.filterAddress, ""))
	}
	if r.filterListAgent != nil {
		localVarQueryParams.Add("filter[list_agent]", parameterToString(*r.filterListAgent, ""))
	}
	if r.filterStatus != nil {
		localVarQueryParams.Add("filter[status]", parameterToString(*r.filterStatus, ""))
	}
	if r.filterActive != nil {
		localVarQueryParams.Add("filter[active]", parameterToString(*r.filterActive, ""))
	}
	if r.filterPriceGte != nil {
		localVarQueryParams.Add("filter[price_gte]", parameterToString(*r.filterPriceGte, ""))
	}
	if r.filterPriceLte != nil {
		localVarQueryParams.Add("filter[price_lte]", parameterToString(*r.filterPriceLte, ""))
	}
	if r.filterSquareFeetGte != nil {
		localVarQueryParams.Add("filter[square_feet_gte]", parameterToString(*r.filterSquareFeetGte, ""))
	}
	if r.filterSquareFeetLte != nil {
		localVarQueryParams.Add("filter[square_feet_lte]", parameterToString(*r.filterSquareFeetLte, ""))
	}
	if r.filterBedroomsGte != nil {
		localVarQueryParams.Add("filter[bedrooms_gte]", parameterToString(*r.filterBedroomsGte, ""))
	}
	if r.filterBedroomsLte != nil {
		localVarQueryParams.Add("filter[bedrooms_lte]", parameterToString(*r.filterBedroomsLte, ""))
	}
	if r.filterBathroomsGte != nil {
		localVarQueryParams.Add("filter[bathrooms_gte]", parameterToString(*r.filterBathroomsGte, ""))
	}
	if r.filterBathroomsLte != nil {
		localVarQueryParams.Add("filter[bathrooms_lte]", parameterToString(*r.filterBathroomsLte, ""))
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
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ApiFail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError
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

type ApiGetListingsIdRequest struct {
	ctx _context.Context
	ApiService *ListingsApiService
	listingId string
	include *string
}

// Comma separated list of optional data to include in the response.
func (r ApiGetListingsIdRequest) Include(include string) ApiGetListingsIdRequest {
	r.include = &include
	return r
}

func (r ApiGetListingsIdRequest) Execute() (ListingResource, *_nethttp.Response, error) {
	return r.ApiService.GetListingsIdExecute(r)
}

/*
GetListingsId Retrieve a listing.

Retrieves the details of a listing with the given ID.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param listingId The ID of a listing. UUID Version 4.
 @return ApiGetListingsIdRequest
*/
func (a *ListingsApiService) GetListingsId(ctx _context.Context, listingId string) ApiGetListingsIdRequest {
	return ApiGetListingsIdRequest{
		ApiService: a,
		ctx: ctx,
		listingId: listingId,
	}
}

// Execute executes the request
//  @return ListingResource
func (a *ListingsApiService) GetListingsIdExecute(r ApiGetListingsIdRequest) (ListingResource, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListingResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ListingsApiService.GetListingsId")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/listings/{listing_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"listing_id"+"}", _neturl.PathEscape(parameterToString(r.listingId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(r.listingId) < 36 {
		return localVarReturnValue, nil, reportError("listingId must have at least 36 elements")
	}
	if strlen(r.listingId) > 36 {
		return localVarReturnValue, nil, reportError("listingId must have less than 36 elements")
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
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ApiFail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError
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
