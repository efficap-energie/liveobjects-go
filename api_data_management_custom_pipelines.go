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

// DataManagementCustomPipelinesApiService DataManagementCustomPipelinesApi service
type DataManagementCustomPipelinesApiService service

type ApiDeleteUsingDELETE11Request struct {
	ctx _context.Context
	ApiService *DataManagementCustomPipelinesApiService
	pipelineId string
	xAPIKEY *string
}

func (r ApiDeleteUsingDELETE11Request) XAPIKEY(xAPIKEY string) ApiDeleteUsingDELETE11Request {
	r.xAPIKEY = &xAPIKEY
	return r
}

func (r ApiDeleteUsingDELETE11Request) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteUsingDELETE11Execute(r)
}

/*
 * DeleteUsingDELETE11 Delete a DataMessage pipeline
 * Usage of this API will be reported in your access log under 'data_pipeline' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pipelineId id of the pipeline to delete
 * @return ApiDeleteUsingDELETE11Request
 */
func (a *DataManagementCustomPipelinesApiService) DeleteUsingDELETE11(ctx _context.Context, pipelineId string) ApiDeleteUsingDELETE11Request {
	return ApiDeleteUsingDELETE11Request{
		ApiService: a,
		ctx: ctx,
		pipelineId: pipelineId,
	}
}

/*
 * Execute executes the request
 */
func (a *DataManagementCustomPipelinesApiService) DeleteUsingDELETE11Execute(r ApiDeleteUsingDELETE11Request) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataManagementCustomPipelinesApiService.DeleteUsingDELETE11")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/pipelines/{pipelineId}"
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineId"+"}", _neturl.PathEscape(parameterToString(r.pipelineId, "")), -1)

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

type ApiGetUsingGET12Request struct {
	ctx _context.Context
	ApiService *DataManagementCustomPipelinesApiService
	pipelineId string
	xAPIKEY *string
}

func (r ApiGetUsingGET12Request) XAPIKEY(xAPIKEY string) ApiGetUsingGET12Request {
	r.xAPIKEY = &xAPIKEY
	return r
}

func (r ApiGetUsingGET12Request) Execute() (Pipeline, *_nethttp.Response, error) {
	return r.ApiService.GetUsingGET12Execute(r)
}

/*
 * GetUsingGET12 Retrieve a DataMessage pipeline
 * Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pipelineId id of the pipeline to retrieve
 * @return ApiGetUsingGET12Request
 */
func (a *DataManagementCustomPipelinesApiService) GetUsingGET12(ctx _context.Context, pipelineId string) ApiGetUsingGET12Request {
	return ApiGetUsingGET12Request{
		ApiService: a,
		ctx: ctx,
		pipelineId: pipelineId,
	}
}

/*
 * Execute executes the request
 * @return Pipeline
 */
func (a *DataManagementCustomPipelinesApiService) GetUsingGET12Execute(r ApiGetUsingGET12Request) (Pipeline, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Pipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataManagementCustomPipelinesApiService.GetUsingGET12")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/pipelines/{pipelineId}"
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineId"+"}", _neturl.PathEscape(parameterToString(r.pipelineId, "")), -1)

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

type ApiListUsingGET13Request struct {
	ctx _context.Context
	ApiService *DataManagementCustomPipelinesApiService
	xAPIKEY *string
}

func (r ApiListUsingGET13Request) XAPIKEY(xAPIKEY string) ApiListUsingGET13Request {
	r.xAPIKEY = &xAPIKEY
	return r
}

func (r ApiListUsingGET13Request) Execute() ([]Pipeline, *_nethttp.Response, error) {
	return r.ApiService.ListUsingGET13Execute(r)
}

/*
 * ListUsingGET13 Retrieve the list of DataMessage pipelines, ordered by priorityLevel
 * Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListUsingGET13Request
 */
func (a *DataManagementCustomPipelinesApiService) ListUsingGET13(ctx _context.Context) ApiListUsingGET13Request {
	return ApiListUsingGET13Request{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []Pipeline
 */
func (a *DataManagementCustomPipelinesApiService) ListUsingGET13Execute(r ApiListUsingGET13Request) ([]Pipeline, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Pipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataManagementCustomPipelinesApiService.ListUsingGET13")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/pipelines"

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

type ApiPostUsingPOST12Request struct {
	ctx _context.Context
	ApiService *DataManagementCustomPipelinesApiService
	xAPIKEY *string
	pipelineDescription *Pipeline
}

func (r ApiPostUsingPOST12Request) XAPIKEY(xAPIKEY string) ApiPostUsingPOST12Request {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiPostUsingPOST12Request) PipelineDescription(pipelineDescription Pipeline) ApiPostUsingPOST12Request {
	r.pipelineDescription = &pipelineDescription
	return r
}

func (r ApiPostUsingPOST12Request) Execute() (Pipeline, *_nethttp.Response, error) {
	return r.ApiService.PostUsingPOST12Execute(r)
}

/*
 * PostUsingPOST12 Create a DataMessage pipeline
 * Usage of this API will be reported in your access log under 'data_pipeline' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPostUsingPOST12Request
 */
func (a *DataManagementCustomPipelinesApiService) PostUsingPOST12(ctx _context.Context) ApiPostUsingPOST12Request {
	return ApiPostUsingPOST12Request{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return Pipeline
 */
func (a *DataManagementCustomPipelinesApiService) PostUsingPOST12Execute(r ApiPostUsingPOST12Request) (Pipeline, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Pipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataManagementCustomPipelinesApiService.PostUsingPOST12")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/pipelines"

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
	localVarPostBody = r.pipelineDescription
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

type ApiUpdateUsingPUT6Request struct {
	ctx _context.Context
	ApiService *DataManagementCustomPipelinesApiService
	pipelineId string
	xAPIKEY *string
	pipelineDescription *Pipeline
}

func (r ApiUpdateUsingPUT6Request) XAPIKEY(xAPIKEY string) ApiUpdateUsingPUT6Request {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiUpdateUsingPUT6Request) PipelineDescription(pipelineDescription Pipeline) ApiUpdateUsingPUT6Request {
	r.pipelineDescription = &pipelineDescription
	return r
}

func (r ApiUpdateUsingPUT6Request) Execute() (Pipeline, *_nethttp.Response, error) {
	return r.ApiService.UpdateUsingPUT6Execute(r)
}

/*
 * UpdateUsingPUT6 Update a DataMessage pipeline
 * Usage of this API will be reported in your access log under 'data_pipeline' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param pipelineId id of the pipeline to update
 * @return ApiUpdateUsingPUT6Request
 */
func (a *DataManagementCustomPipelinesApiService) UpdateUsingPUT6(ctx _context.Context, pipelineId string) ApiUpdateUsingPUT6Request {
	return ApiUpdateUsingPUT6Request{
		ApiService: a,
		ctx: ctx,
		pipelineId: pipelineId,
	}
}

/*
 * Execute executes the request
 * @return Pipeline
 */
func (a *DataManagementCustomPipelinesApiService) UpdateUsingPUT6Execute(r ApiUpdateUsingPUT6Request) (Pipeline, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Pipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataManagementCustomPipelinesApiService.UpdateUsingPUT6")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/pipelines/{pipelineId}"
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineId"+"}", _neturl.PathEscape(parameterToString(r.pipelineId, "")), -1)

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
	localVarPostBody = r.pipelineDescription
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
