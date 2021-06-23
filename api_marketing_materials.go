/*
 * Aryeo
 *
 *
 * API version: 1.0.0
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

// MarketingMaterialsApiService MarketingMaterialsApi service
type MarketingMaterialsApiService service

type ApiPutMarketingMaterialsTemplatesUuidPublishRequest struct {
	ctx _context.Context
	ApiService *MarketingMaterialsApiService
	uuid string
	marketingMaterialTemplatePublishPayload *MarketingMaterialTemplatePublishPayload
}

func (r ApiPutMarketingMaterialsTemplatesUuidPublishRequest) MarketingMaterialTemplatePublishPayload(marketingMaterialTemplatePublishPayload MarketingMaterialTemplatePublishPayload) ApiPutMarketingMaterialsTemplatesUuidPublishRequest {
	r.marketingMaterialTemplatePublishPayload = &marketingMaterialTemplatePublishPayload
	return r
}

func (r ApiPutMarketingMaterialsTemplatesUuidPublishRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PutMarketingMaterialsTemplatesUuidPublishExecute(r)
}

/*
 * PutMarketingMaterialsTemplatesUuidPublish Publish a marketing material template.
 * Publish a marketing material template.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uuid UUID of the marketing material template record.
 * @return ApiPutMarketingMaterialsTemplatesUuidPublishRequest
 */
func (a *MarketingMaterialsApiService) PutMarketingMaterialsTemplatesUuidPublish(ctx _context.Context, uuid string) ApiPutMarketingMaterialsTemplatesUuidPublishRequest {
	return ApiPutMarketingMaterialsTemplatesUuidPublishRequest{
		ApiService: a,
		ctx: ctx,
		uuid: uuid,
	}
}

/*
 * Execute executes the request
 */
func (a *MarketingMaterialsApiService) PutMarketingMaterialsTemplatesUuidPublishExecute(r ApiPutMarketingMaterialsTemplatesUuidPublishRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketingMaterialsApiService.PutMarketingMaterialsTemplatesUuidPublish")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing-materials/templates/{uuid}/publish"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", _neturl.PathEscape(parameterToString(r.uuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(r.uuid) < 0 {
		return nil, reportError("uuid must have at least 0 elements")
	}
	if strlen(r.uuid) > 255 {
		return nil, reportError("uuid must have less than 255 elements")
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
	localVarPostBody = r.marketingMaterialTemplatePublishPayload
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPutMarketingMaterialsUuidPublishRequest struct {
	ctx _context.Context
	ApiService *MarketingMaterialsApiService
	uuid string
	marketingMaterialPublishPayload *MarketingMaterialPublishPayload
}

func (r ApiPutMarketingMaterialsUuidPublishRequest) MarketingMaterialPublishPayload(marketingMaterialPublishPayload MarketingMaterialPublishPayload) ApiPutMarketingMaterialsUuidPublishRequest {
	r.marketingMaterialPublishPayload = &marketingMaterialPublishPayload
	return r
}

func (r ApiPutMarketingMaterialsUuidPublishRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.PutMarketingMaterialsUuidPublishExecute(r)
}

/*
 * PutMarketingMaterialsUuidPublish Publish a marketing material.
 * Publish a marketing material.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param uuid UUID of the marketing material record.
 * @return ApiPutMarketingMaterialsUuidPublishRequest
 */
func (a *MarketingMaterialsApiService) PutMarketingMaterialsUuidPublish(ctx _context.Context, uuid string) ApiPutMarketingMaterialsUuidPublishRequest {
	return ApiPutMarketingMaterialsUuidPublishRequest{
		ApiService: a,
		ctx: ctx,
		uuid: uuid,
	}
}

/*
 * Execute executes the request
 */
func (a *MarketingMaterialsApiService) PutMarketingMaterialsUuidPublishExecute(r ApiPutMarketingMaterialsUuidPublishRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketingMaterialsApiService.PutMarketingMaterialsUuidPublish")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing-materials/{uuid}/publish"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", _neturl.PathEscape(parameterToString(r.uuid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if strlen(r.uuid) < 0 {
		return nil, reportError("uuid must have at least 0 elements")
	}
	if strlen(r.uuid) > 255 {
		return nil, reportError("uuid must have less than 255 elements")
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
	localVarPostBody = r.marketingMaterialPublishPayload
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
