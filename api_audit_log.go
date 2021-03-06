/*
 * Live Objects REST API Guide v2.13.3
 *
 * API description for Live Objects service
 *
 * API version: 2.13.3
 * Contact: liveobjects.support@orange.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package liveobjects

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// AuditLogApiService AuditLogApi service
type AuditLogApiService service

type ApiSearchUsingGETRequest struct {
	ctx _context.Context
	ApiService *AuditLogApiService
	xAPIKEY *string
	from *string
	to *string
	offset *int32
	limit *int32
	sort *string
	filters *string
	any *[]string
}

func (r ApiSearchUsingGETRequest) XAPIKEY(xAPIKEY string) ApiSearchUsingGETRequest {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiSearchUsingGETRequest) From(from string) ApiSearchUsingGETRequest {
	r.from = &from
	return r
}
func (r ApiSearchUsingGETRequest) To(to string) ApiSearchUsingGETRequest {
	r.to = &to
	return r
}
func (r ApiSearchUsingGETRequest) Offset(offset int32) ApiSearchUsingGETRequest {
	r.offset = &offset
	return r
}
func (r ApiSearchUsingGETRequest) Limit(limit int32) ApiSearchUsingGETRequest {
	r.limit = &limit
	return r
}
func (r ApiSearchUsingGETRequest) Sort(sort string) ApiSearchUsingGETRequest {
	r.sort = &sort
	return r
}
func (r ApiSearchUsingGETRequest) Filters(filters string) ApiSearchUsingGETRequest {
	r.filters = &filters
	return r
}
func (r ApiSearchUsingGETRequest) Any(any []string) ApiSearchUsingGETRequest {
	r.any = &any
	return r
}

func (r ApiSearchUsingGETRequest) Execute() ([]AuditLogMessage, *_nethttp.Response, error) {
	return r.ApiService.SearchUsingGETExecute(r)
}

/*
 * SearchUsingGET Retrieve messages available in your AuditLog
 * Restricted to API keys with at least one of the following roles : LOGS_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSearchUsingGETRequest
 */
func (a *AuditLogApiService) SearchUsingGET(ctx _context.Context) ApiSearchUsingGETRequest {
	return ApiSearchUsingGETRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []AuditLogMessage
 */
func (a *AuditLogApiService) SearchUsingGETExecute(r ApiSearchUsingGETRequest) ([]AuditLogMessage, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []AuditLogMessage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditLogApiService.SearchUsingGET")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/auditlog/messages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return localVarReturnValue, nil, reportError("xAPIKEY is required and must be specified")
	}

	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.filters != nil {
		localVarQueryParams.Add("filters", parameterToString(*r.filters, ""))
	}
	if r.any != nil {
		t := *r.any
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("@any", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("@any", parameterToString(t, "multi"))
		}
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
	localVarHeaderParams["X-API-KEY"] = parameterToString(*r.xAPIKEY, "")
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
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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
