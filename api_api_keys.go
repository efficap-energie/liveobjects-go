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
	"strings"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// ApiKeysApiService ApiKeysApi service
type ApiKeysApiService service

type ApiCreateApiKeyUsingPOSTRequest struct {
	ctx _context.Context
	ApiService *ApiKeysApiService
	xAPIKEY *string
	apiKeyCreationRequest *ApiKeyCreationReqWeb
}

func (r ApiCreateApiKeyUsingPOSTRequest) XAPIKEY(xAPIKEY string) ApiCreateApiKeyUsingPOSTRequest {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiCreateApiKeyUsingPOSTRequest) ApiKeyCreationRequest(apiKeyCreationRequest ApiKeyCreationReqWeb) ApiCreateApiKeyUsingPOSTRequest {
	r.apiKeyCreationRequest = &apiKeyCreationRequest
	return r
}

func (r ApiCreateApiKeyUsingPOSTRequest) Execute() (ApiKey, *_nethttp.Response, error) {
	return r.ApiService.CreateApiKeyUsingPOSTExecute(r)
}

/*
 * CreateApiKeyUsingPOST Create an API key
 * Usage of this API will be reported in your access log under 'api_key' category.<br><br>Restricted to API keys with at least one of the following roles : API_KEY_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCreateApiKeyUsingPOSTRequest
 */
func (a *ApiKeysApiService) CreateApiKeyUsingPOST(ctx _context.Context) ApiCreateApiKeyUsingPOSTRequest {
	return ApiCreateApiKeyUsingPOSTRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ApiKey
 */
func (a *ApiKeysApiService) CreateApiKeyUsingPOSTExecute(r ApiCreateApiKeyUsingPOSTRequest) (ApiKey, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ApiKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiKeysApiService.CreateApiKeyUsingPOST")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/apiKeys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return localVarReturnValue, nil, reportError("xAPIKEY is required and must be specified")
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
	localVarHeaderParams["X-API-KEY"] = parameterToString(*r.xAPIKEY, "")
	// body params
	localVarPostBody = r.apiKeyCreationRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponseWeb
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

type ApiDeleteApiKeyUsingDELETERequest struct {
	ctx _context.Context
	ApiService *ApiKeysApiService
	apiKeyId string
	xAPIKEY *string
	tenantId *string
}

func (r ApiDeleteApiKeyUsingDELETERequest) XAPIKEY(xAPIKEY string) ApiDeleteApiKeyUsingDELETERequest {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiDeleteApiKeyUsingDELETERequest) TenantId(tenantId string) ApiDeleteApiKeyUsingDELETERequest {
	r.tenantId = &tenantId
	return r
}

func (r ApiDeleteApiKeyUsingDELETERequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteApiKeyUsingDELETEExecute(r)
}

/*
 * DeleteApiKeyUsingDELETE Delete an API key
 * Usage of this API will be reported in your access log under 'api_key' category.<br><br>Restricted to API keys with at least one of the following roles : API_KEY_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param apiKeyId identifier of your API key. Expected identifier (max 24 characters)
 * @return ApiDeleteApiKeyUsingDELETERequest
 */
func (a *ApiKeysApiService) DeleteApiKeyUsingDELETE(ctx _context.Context, apiKeyId string) ApiDeleteApiKeyUsingDELETERequest {
	return ApiDeleteApiKeyUsingDELETERequest{
		ApiService: a,
		ctx: ctx,
		apiKeyId: apiKeyId,
	}
}

/*
 * Execute executes the request
 */
func (a *ApiKeysApiService) DeleteApiKeyUsingDELETEExecute(r ApiDeleteApiKeyUsingDELETERequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiKeysApiService.DeleteApiKeyUsingDELETE")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/apiKeys/{apiKeyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"apiKeyId"+"}", _neturl.PathEscape(parameterToString(r.apiKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return nil, reportError("xAPIKEY is required and must be specified")
	}

	if r.tenantId != nil {
		localVarQueryParams.Add("tenantId", parameterToString(*r.tenantId, ""))
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponseWeb
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

type ApiGetApiKeyFromAuthenticationUsingGET3Request struct {
	ctx _context.Context
	ApiService *ApiKeysApiService
	xAPIKEY *string
}

func (r ApiGetApiKeyFromAuthenticationUsingGET3Request) XAPIKEY(xAPIKEY string) ApiGetApiKeyFromAuthenticationUsingGET3Request {
	r.xAPIKEY = &xAPIKEY
	return r
}

func (r ApiGetApiKeyFromAuthenticationUsingGET3Request) Execute() (ApiKey, *_nethttp.Response, error) {
	return r.ApiService.GetApiKeyFromAuthenticationUsingGET3Execute(r)
}

/*
 * GetApiKeyFromAuthenticationUsingGET3 getApiKeyFromAuthentication
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetApiKeyFromAuthenticationUsingGET3Request
 */
func (a *ApiKeysApiService) GetApiKeyFromAuthenticationUsingGET3(ctx _context.Context) ApiGetApiKeyFromAuthenticationUsingGET3Request {
	return ApiGetApiKeyFromAuthenticationUsingGET3Request{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ApiKey
 */
func (a *ApiKeysApiService) GetApiKeyFromAuthenticationUsingGET3Execute(r ApiGetApiKeyFromAuthenticationUsingGET3Request) (ApiKey, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ApiKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiKeysApiService.GetApiKeyFromAuthenticationUsingGET3")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/apiKeys/current_key"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return localVarReturnValue, nil, reportError("xAPIKEY is required and must be specified")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponseWeb
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponseWeb
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

type ApiGetApiKeyUsingGET3Request struct {
	ctx _context.Context
	ApiService *ApiKeysApiService
	apiKeyId string
	xAPIKEY *string
	tenantId *string
}

func (r ApiGetApiKeyUsingGET3Request) XAPIKEY(xAPIKEY string) ApiGetApiKeyUsingGET3Request {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiGetApiKeyUsingGET3Request) TenantId(tenantId string) ApiGetApiKeyUsingGET3Request {
	r.tenantId = &tenantId
	return r
}

func (r ApiGetApiKeyUsingGET3Request) Execute() (ApiKey, *_nethttp.Response, error) {
	return r.ApiService.GetApiKeyUsingGET3Execute(r)
}

/*
 * GetApiKeyUsingGET3 Get an API key
 * Restricted to API keys with at least one of the following roles : API_KEY_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param apiKeyId the id of your API key. Expected identifier (max 24 characters)
 * @return ApiGetApiKeyUsingGET3Request
 */
func (a *ApiKeysApiService) GetApiKeyUsingGET3(ctx _context.Context, apiKeyId string) ApiGetApiKeyUsingGET3Request {
	return ApiGetApiKeyUsingGET3Request{
		ApiService: a,
		ctx: ctx,
		apiKeyId: apiKeyId,
	}
}

/*
 * Execute executes the request
 * @return ApiKey
 */
func (a *ApiKeysApiService) GetApiKeyUsingGET3Execute(r ApiGetApiKeyUsingGET3Request) (ApiKey, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ApiKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiKeysApiService.GetApiKeyUsingGET3")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/apiKeys/{apiKeyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"apiKeyId"+"}", _neturl.PathEscape(parameterToString(r.apiKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return localVarReturnValue, nil, reportError("xAPIKEY is required and must be specified")
	}

	if r.tenantId != nil {
		localVarQueryParams.Add("tenantId", parameterToString(*r.tenantId, ""))
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponseWeb
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

type ApiGetApiKeysUsingGET3Request struct {
	ctx _context.Context
	ApiService *ApiKeysApiService
	xAPIKEY *string
	size *int32
	page *int32
	tenantId *string
	parentId *string
	showSessionKeys *bool
	roles *[]string
	showMasterKey *bool
}

func (r ApiGetApiKeysUsingGET3Request) XAPIKEY(xAPIKEY string) ApiGetApiKeysUsingGET3Request {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiGetApiKeysUsingGET3Request) Size(size int32) ApiGetApiKeysUsingGET3Request {
	r.size = &size
	return r
}
func (r ApiGetApiKeysUsingGET3Request) Page(page int32) ApiGetApiKeysUsingGET3Request {
	r.page = &page
	return r
}
func (r ApiGetApiKeysUsingGET3Request) TenantId(tenantId string) ApiGetApiKeysUsingGET3Request {
	r.tenantId = &tenantId
	return r
}
func (r ApiGetApiKeysUsingGET3Request) ParentId(parentId string) ApiGetApiKeysUsingGET3Request {
	r.parentId = &parentId
	return r
}
func (r ApiGetApiKeysUsingGET3Request) ShowSessionKeys(showSessionKeys bool) ApiGetApiKeysUsingGET3Request {
	r.showSessionKeys = &showSessionKeys
	return r
}
func (r ApiGetApiKeysUsingGET3Request) Roles(roles []string) ApiGetApiKeysUsingGET3Request {
	r.roles = &roles
	return r
}
func (r ApiGetApiKeysUsingGET3Request) ShowMasterKey(showMasterKey bool) ApiGetApiKeysUsingGET3Request {
	r.showMasterKey = &showMasterKey
	return r
}

func (r ApiGetApiKeysUsingGET3Request) Execute() (PageableApiKey, *_nethttp.Response, error) {
	return r.ApiService.GetApiKeysUsingGET3Execute(r)
}

/*
 * GetApiKeysUsingGET3 List API keys
 * Restricted to API keys with at least one of the following roles : API_KEY_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetApiKeysUsingGET3Request
 */
func (a *ApiKeysApiService) GetApiKeysUsingGET3(ctx _context.Context) ApiGetApiKeysUsingGET3Request {
	return ApiGetApiKeysUsingGET3Request{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PageableApiKey
 */
func (a *ApiKeysApiService) GetApiKeysUsingGET3Execute(r ApiGetApiKeysUsingGET3Request) (PageableApiKey, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PageableApiKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiKeysApiService.GetApiKeysUsingGET3")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/apiKeys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return localVarReturnValue, nil, reportError("xAPIKEY is required and must be specified")
	}

	if r.size != nil {
		localVarQueryParams.Add("size", parameterToString(*r.size, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.tenantId != nil {
		localVarQueryParams.Add("tenantId", parameterToString(*r.tenantId, ""))
	}
	if r.parentId != nil {
		localVarQueryParams.Add("parentId", parameterToString(*r.parentId, ""))
	}
	if r.showSessionKeys != nil {
		localVarQueryParams.Add("showSessionKeys", parameterToString(*r.showSessionKeys, ""))
	}
	if r.roles != nil {
		t := *r.roles
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("roles", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("roles", parameterToString(t, "multi"))
		}
	}
	if r.showMasterKey != nil {
		localVarQueryParams.Add("showMasterKey", parameterToString(*r.showMasterKey, ""))
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponseWeb
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

type ApiSetApiKeyDebugModeUsingPUT3Request struct {
	ctx _context.Context
	ApiService *ApiKeysApiService
	apiKeyId string
	xAPIKEY *string
	setApiKeyDebugModeRequest *ApiKeySetDebugModeReqWeb
}

func (r ApiSetApiKeyDebugModeUsingPUT3Request) XAPIKEY(xAPIKEY string) ApiSetApiKeyDebugModeUsingPUT3Request {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiSetApiKeyDebugModeUsingPUT3Request) SetApiKeyDebugModeRequest(setApiKeyDebugModeRequest ApiKeySetDebugModeReqWeb) ApiSetApiKeyDebugModeUsingPUT3Request {
	r.setApiKeyDebugModeRequest = &setApiKeyDebugModeRequest
	return r
}

func (r ApiSetApiKeyDebugModeUsingPUT3Request) Execute() (ApiKey, *_nethttp.Response, error) {
	return r.ApiService.SetApiKeyDebugModeUsingPUT3Execute(r)
}

/*
 * SetApiKeyDebugModeUsingPUT3 Activate/Deactivate the debug mode on an API key
 * Usage of this API will be reported in your access log under 'api_key' category.<br><br>Restricted to API keys with at least one of the following roles : API_KEY_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param apiKeyId identifier of your API key. Expected identifier (max 24 characters)
 * @return ApiSetApiKeyDebugModeUsingPUT3Request
 */
func (a *ApiKeysApiService) SetApiKeyDebugModeUsingPUT3(ctx _context.Context, apiKeyId string) ApiSetApiKeyDebugModeUsingPUT3Request {
	return ApiSetApiKeyDebugModeUsingPUT3Request{
		ApiService: a,
		ctx: ctx,
		apiKeyId: apiKeyId,
	}
}

/*
 * Execute executes the request
 * @return ApiKey
 */
func (a *ApiKeysApiService) SetApiKeyDebugModeUsingPUT3Execute(r ApiSetApiKeyDebugModeUsingPUT3Request) (ApiKey, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ApiKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiKeysApiService.SetApiKeyDebugModeUsingPUT3")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/apiKeys/{apiKeyId}/debugMode"
	localVarPath = strings.Replace(localVarPath, "{"+"apiKeyId"+"}", _neturl.PathEscape(parameterToString(r.apiKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return localVarReturnValue, nil, reportError("xAPIKEY is required and must be specified")
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
	localVarHeaderParams["X-API-KEY"] = parameterToString(*r.xAPIKEY, "")
	// body params
	localVarPostBody = r.setApiKeyDebugModeRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponseWeb
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponseWeb
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponseWeb
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorResponseWeb
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

type ApiUpdateApiKeyUsingPOST1Request struct {
	ctx _context.Context
	ApiService *ApiKeysApiService
	apiKeyId string
	xAPIKEY *string
	apiKeyUpdateRequest *ApiKeyUpdateReqWeb
}

func (r ApiUpdateApiKeyUsingPOST1Request) XAPIKEY(xAPIKEY string) ApiUpdateApiKeyUsingPOST1Request {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiUpdateApiKeyUsingPOST1Request) ApiKeyUpdateRequest(apiKeyUpdateRequest ApiKeyUpdateReqWeb) ApiUpdateApiKeyUsingPOST1Request {
	r.apiKeyUpdateRequest = &apiKeyUpdateRequest
	return r
}

func (r ApiUpdateApiKeyUsingPOST1Request) Execute() (ApiKey, *_nethttp.Response, error) {
	return r.ApiService.UpdateApiKeyUsingPOST1Execute(r)
}

/*
 * UpdateApiKeyUsingPOST1 Update an API key
 * Update a set of properties of the selected API key<br><br>Usage of this API will be reported in your access log under 'api_key' category.<br><br>Restricted to API keys with at least one of the following roles : API_KEY_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param apiKeyId the id of your API key. Expected identifier (max 24 characters)
 * @return ApiUpdateApiKeyUsingPOST1Request
 */
func (a *ApiKeysApiService) UpdateApiKeyUsingPOST1(ctx _context.Context, apiKeyId string) ApiUpdateApiKeyUsingPOST1Request {
	return ApiUpdateApiKeyUsingPOST1Request{
		ApiService: a,
		ctx: ctx,
		apiKeyId: apiKeyId,
	}
}

/*
 * Execute executes the request
 * @return ApiKey
 */
func (a *ApiKeysApiService) UpdateApiKeyUsingPOST1Execute(r ApiUpdateApiKeyUsingPOST1Request) (ApiKey, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ApiKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiKeysApiService.UpdateApiKeyUsingPOST1")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/apiKeys/{apiKeyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"apiKeyId"+"}", _neturl.PathEscape(parameterToString(r.apiKeyId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return localVarReturnValue, nil, reportError("xAPIKEY is required and must be specified")
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
	localVarHeaderParams["X-API-KEY"] = parameterToString(*r.xAPIKEY, "")
	// body params
	localVarPostBody = r.apiKeyUpdateRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponseWeb
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
