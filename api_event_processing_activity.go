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
)

// Linger please
var (
	_ _context.Context
)

// EventProcessingActivityApiService EventProcessingActivityApi service
type EventProcessingActivityApiService service

type ApiDeleteUsingDELETE14Request struct {
	ctx _context.Context
	ApiService *EventProcessingActivityApiService
	activityRuleId string
	xAPIKEY *string
}

func (r ApiDeleteUsingDELETE14Request) XAPIKEY(xAPIKEY string) ApiDeleteUsingDELETE14Request {
	r.xAPIKEY = &xAPIKEY
	return r
}

func (r ApiDeleteUsingDELETE14Request) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteUsingDELETE14Execute(r)
}

/*
 * DeleteUsingDELETE14 Delete an ActivityRule
 * Usage of this API will be reported in your access log under 'alarming' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param activityRuleId id of the ActivityRule to delete
 * @return ApiDeleteUsingDELETE14Request
 */
func (a *EventProcessingActivityApiService) DeleteUsingDELETE14(ctx _context.Context, activityRuleId string) ApiDeleteUsingDELETE14Request {
	return ApiDeleteUsingDELETE14Request{
		ApiService: a,
		ctx: ctx,
		activityRuleId: activityRuleId,
	}
}

/*
 * Execute executes the request
 */
func (a *EventProcessingActivityApiService) DeleteUsingDELETE14Execute(r ApiDeleteUsingDELETE14Request) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventProcessingActivityApiService.DeleteUsingDELETE14")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/eventprocessing/activity/rules/{activityRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"activityRuleId"+"}", _neturl.PathEscape(parameterToString(r.activityRuleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return nil, reportError("xAPIKEY is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetStatesUsingGETRequest struct {
	ctx _context.Context
	ApiService *EventProcessingActivityApiService
	xAPIKEY *string
	deviceId *string
	activityRuleId *string
	limit *int32
	bookmarkDeviceId *string
}

func (r ApiGetStatesUsingGETRequest) XAPIKEY(xAPIKEY string) ApiGetStatesUsingGETRequest {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiGetStatesUsingGETRequest) DeviceId(deviceId string) ApiGetStatesUsingGETRequest {
	r.deviceId = &deviceId
	return r
}
func (r ApiGetStatesUsingGETRequest) ActivityRuleId(activityRuleId string) ApiGetStatesUsingGETRequest {
	r.activityRuleId = &activityRuleId
	return r
}
func (r ApiGetStatesUsingGETRequest) Limit(limit int32) ApiGetStatesUsingGETRequest {
	r.limit = &limit
	return r
}
func (r ApiGetStatesUsingGETRequest) BookmarkDeviceId(bookmarkDeviceId string) ApiGetStatesUsingGETRequest {
	r.bookmarkDeviceId = &bookmarkDeviceId
	return r
}

func (r ApiGetStatesUsingGETRequest) Execute() ([]ActivityState, *_nethttp.Response, error) {
	return r.ApiService.GetStatesUsingGETExecute(r)
}

/*
 * GetStatesUsingGET Retrieve the list of all the ActivityStates linked to a specific device and/or rule
 * Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetStatesUsingGETRequest
 */
func (a *EventProcessingActivityApiService) GetStatesUsingGET(ctx _context.Context) ApiGetStatesUsingGETRequest {
	return ApiGetStatesUsingGETRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []ActivityState
 */
func (a *EventProcessingActivityApiService) GetStatesUsingGETExecute(r ApiGetStatesUsingGETRequest) ([]ActivityState, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []ActivityState
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventProcessingActivityApiService.GetStatesUsingGET")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/eventprocessing/activity/states"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return localVarReturnValue, nil, reportError("xAPIKEY is required and must be specified")
	}

	if r.deviceId != nil {
		localVarQueryParams.Add("deviceId", parameterToString(*r.deviceId, ""))
	}
	if r.activityRuleId != nil {
		localVarQueryParams.Add("activityRuleId", parameterToString(*r.activityRuleId, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.bookmarkDeviceId != nil {
		localVarQueryParams.Add("bookmarkDeviceId", parameterToString(*r.bookmarkDeviceId, ""))
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

type ApiGetUsingGET13Request struct {
	ctx _context.Context
	ApiService *EventProcessingActivityApiService
	activityRuleId string
	xAPIKEY *string
}

func (r ApiGetUsingGET13Request) XAPIKEY(xAPIKEY string) ApiGetUsingGET13Request {
	r.xAPIKEY = &xAPIKEY
	return r
}

func (r ApiGetUsingGET13Request) Execute() (ActivityRule, *_nethttp.Response, error) {
	return r.ApiService.GetUsingGET13Execute(r)
}

/*
 * GetUsingGET13 Retrieve an ActivityRule
 * Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param activityRuleId id of the ActivityRule to retrieve
 * @return ApiGetUsingGET13Request
 */
func (a *EventProcessingActivityApiService) GetUsingGET13(ctx _context.Context, activityRuleId string) ApiGetUsingGET13Request {
	return ApiGetUsingGET13Request{
		ApiService: a,
		ctx: ctx,
		activityRuleId: activityRuleId,
	}
}

/*
 * Execute executes the request
 * @return ActivityRule
 */
func (a *EventProcessingActivityApiService) GetUsingGET13Execute(r ApiGetUsingGET13Request) (ActivityRule, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ActivityRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventProcessingActivityApiService.GetUsingGET13")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/eventprocessing/activity/rules/{activityRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"activityRuleId"+"}", _neturl.PathEscape(parameterToString(r.activityRuleId, "")), -1)

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

type ApiListUsingGET16Request struct {
	ctx _context.Context
	ApiService *EventProcessingActivityApiService
	xAPIKEY *string
	name *string
}

func (r ApiListUsingGET16Request) XAPIKEY(xAPIKEY string) ApiListUsingGET16Request {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiListUsingGET16Request) Name(name string) ApiListUsingGET16Request {
	r.name = &name
	return r
}

func (r ApiListUsingGET16Request) Execute() ([]ActivityRule, *_nethttp.Response, error) {
	return r.ApiService.ListUsingGET16Execute(r)
}

/*
 * ListUsingGET16 Retrieve the list of all the ActivityRules or get an ActivityRule by its name
 * Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListUsingGET16Request
 */
func (a *EventProcessingActivityApiService) ListUsingGET16(ctx _context.Context) ApiListUsingGET16Request {
	return ApiListUsingGET16Request{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []ActivityRule
 */
func (a *EventProcessingActivityApiService) ListUsingGET16Execute(r ApiListUsingGET16Request) ([]ActivityRule, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []ActivityRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventProcessingActivityApiService.ListUsingGET16")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/eventprocessing/activity/rules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return localVarReturnValue, nil, reportError("xAPIKEY is required and must be specified")
	}

	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
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

type ApiMuteUsingPUTRequest struct {
	ctx _context.Context
	ApiService *EventProcessingActivityApiService
	xAPIKEY *string
	nextAlarmRequest *ActivityStateMuteRequest
}

func (r ApiMuteUsingPUTRequest) XAPIKEY(xAPIKEY string) ApiMuteUsingPUTRequest {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiMuteUsingPUTRequest) NextAlarmRequest(nextAlarmRequest ActivityStateMuteRequest) ApiMuteUsingPUTRequest {
	r.nextAlarmRequest = &nextAlarmRequest
	return r
}

func (r ApiMuteUsingPUTRequest) Execute() (int64, *_nethttp.Response, error) {
	return r.ApiService.MuteUsingPUTExecute(r)
}

/*
 * MuteUsingPUT Mute or reset nextAlarm of ActivityStates targeted by a specific deviceId/activityRuleId
 * Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiMuteUsingPUTRequest
 */
func (a *EventProcessingActivityApiService) MuteUsingPUT(ctx _context.Context) ApiMuteUsingPUTRequest {
	return ApiMuteUsingPUTRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return int64
 */
func (a *EventProcessingActivityApiService) MuteUsingPUTExecute(r ApiMuteUsingPUTRequest) (int64, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  int64
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventProcessingActivityApiService.MuteUsingPUT")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/eventprocessing/activity/states/mute"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return localVarReturnValue, nil, reportError("xAPIKEY is required and must be specified")
	}
	if r.nextAlarmRequest == nil {
		return localVarReturnValue, nil, reportError("nextAlarmRequest is required and must be specified")
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
	localVarPostBody = r.nextAlarmRequest
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

type ApiPostUsingPOST13Request struct {
	ctx _context.Context
	ApiService *EventProcessingActivityApiService
	xAPIKEY *string
	activityRule *ActivityRule
}

func (r ApiPostUsingPOST13Request) XAPIKEY(xAPIKEY string) ApiPostUsingPOST13Request {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiPostUsingPOST13Request) ActivityRule(activityRule ActivityRule) ApiPostUsingPOST13Request {
	r.activityRule = &activityRule
	return r
}

func (r ApiPostUsingPOST13Request) Execute() (ActivityRule, *_nethttp.Response, error) {
	return r.ApiService.PostUsingPOST13Execute(r)
}

/*
 * PostUsingPOST13 Create an ActivityRule
 * Total number of ActivityRules is limited. Contact the commercial team or see developer guide to get more information.<br><br>Usage of this API will be reported in your access log under 'alarming' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostUsingPOST13Request
 */
func (a *EventProcessingActivityApiService) PostUsingPOST13(ctx _context.Context) ApiPostUsingPOST13Request {
	return ApiPostUsingPOST13Request{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ActivityRule
 */
func (a *EventProcessingActivityApiService) PostUsingPOST13Execute(r ApiPostUsingPOST13Request) (ActivityRule, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ActivityRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventProcessingActivityApiService.PostUsingPOST13")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/eventprocessing/activity/rules"

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
	localVarPostBody = r.activityRule
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

type ApiUpdateUsingPUT7Request struct {
	ctx _context.Context
	ApiService *EventProcessingActivityApiService
	activityRuleId string
	xAPIKEY *string
	activityRule *ActivityRule
}

func (r ApiUpdateUsingPUT7Request) XAPIKEY(xAPIKEY string) ApiUpdateUsingPUT7Request {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiUpdateUsingPUT7Request) ActivityRule(activityRule ActivityRule) ApiUpdateUsingPUT7Request {
	r.activityRule = &activityRule
	return r
}

func (r ApiUpdateUsingPUT7Request) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UpdateUsingPUT7Execute(r)
}

/*
 * UpdateUsingPUT7 Update an ActivityRule
 * Usage of this API will be reported in your access log under 'alarming' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param activityRuleId id of the ActivityRule to update
 * @return ApiUpdateUsingPUT7Request
 */
func (a *EventProcessingActivityApiService) UpdateUsingPUT7(ctx _context.Context, activityRuleId string) ApiUpdateUsingPUT7Request {
	return ApiUpdateUsingPUT7Request{
		ApiService: a,
		ctx: ctx,
		activityRuleId: activityRuleId,
	}
}

/*
 * Execute executes the request
 */
func (a *EventProcessingActivityApiService) UpdateUsingPUT7Execute(r ApiUpdateUsingPUT7Request) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventProcessingActivityApiService.UpdateUsingPUT7")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/eventprocessing/activity/rules/{activityRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"activityRuleId"+"}", _neturl.PathEscape(parameterToString(r.activityRuleId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return nil, reportError("xAPIKEY is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-API-KEY"] = parameterToString(*r.xAPIKEY, "")
	// body params
	localVarPostBody = r.activityRule
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
