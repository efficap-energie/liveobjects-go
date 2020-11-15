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
)

// Linger please
var (
	_ _context.Context
)

// DeprecatedTriggersAndActionsTestApiService DeprecatedTriggersAndActionsTestApi service
type DeprecatedTriggersAndActionsTestApiService service

type ApiTestHttpPushUsingPOSTRequest struct {
	ctx _context.Context
	ApiService *DeprecatedTriggersAndActionsTestApiService
	xAPIKEY *string
	req *HttpPushWebhookTest
}

func (r ApiTestHttpPushUsingPOSTRequest) XAPIKEY(xAPIKEY string) ApiTestHttpPushUsingPOSTRequest {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiTestHttpPushUsingPOSTRequest) Req(req HttpPushWebhookTest) ApiTestHttpPushUsingPOSTRequest {
	r.req = &req
	return r
}

func (r ApiTestHttpPushUsingPOSTRequest) Execute() (HttpPushTestResult, *_nethttp.Response, error) {
	return r.ApiService.TestHttpPushUsingPOSTExecute(r)
}

/*
 * TestHttpPushUsingPOST Post an http request for testing a webhook
 * Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiTestHttpPushUsingPOSTRequest
 */
func (a *DeprecatedTriggersAndActionsTestApiService) TestHttpPushUsingPOST(ctx _context.Context) ApiTestHttpPushUsingPOSTRequest {
	return ApiTestHttpPushUsingPOSTRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return HttpPushTestResult
 */
func (a *DeprecatedTriggersAndActionsTestApiService) TestHttpPushUsingPOSTExecute(r ApiTestHttpPushUsingPOSTRequest) (HttpPushTestResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  HttpPushTestResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeprecatedTriggersAndActionsTestApiService.TestHttpPushUsingPOST")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/event2action/test/http-push"

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
	localVarPostBody = r.req
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
