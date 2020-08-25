/*
 * Live Objects REST API Guide v2.12.2
 *
 * API description for Live Objects service
 *
 * API version: 2.12.2
 * Contact: liveobjects.support@orange.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package liveobjects

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	_bytes "bytes"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// EventProcessingFiringApiService EventProcessingFiringApi service
type EventProcessingFiringApiService service

/*
DeleteUsingDELETE16 Delete a FiringRule
Usage of this API will be reported in your access log under &#39;alarming&#39; category.&lt;br&gt;&lt;br&gt;Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param firingRuleId id of the FiringRule to delete
 * @param xAPIKEY a valid API key
*/
func (a *EventProcessingFiringApiService) DeleteUsingDELETE16(ctx _context.Context, firingRuleId string, xAPIKEY string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v0/eventprocessing/firing-rule/{firingRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"firingRuleId"+"}", _neturl.PathEscape(parameterToString(firingRuleId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarHeaderParams["X-API-KEY"] = parameterToString(xAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
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

/*
GetFiringGuardUsingGET Get a FiringGuard
Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param firingGuardId id of the FiringGuard
 * @param xAPIKEY a valid API key
@return FiringGuard
*/
func (a *EventProcessingFiringApiService) GetFiringGuardUsingGET(ctx _context.Context, firingGuardId string, xAPIKEY string) (FiringGuard, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FiringGuard
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v0/eventprocessing/firing-guard/{firingGuardId}"
	localVarPath = strings.Replace(localVarPath, "{"+"firingGuardId"+"}", _neturl.PathEscape(parameterToString(firingGuardId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarHeaderParams["X-API-KEY"] = parameterToString(xAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
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

// GetFiringGuardsUsingPOSTOpts Optional parameters for the method 'GetFiringGuardsUsingPOST'
type GetFiringGuardsUsingPOSTOpts struct {
    SearchParam optional.Interface
}

/*
GetFiringGuardsUsingPOST Get FiringGuards linked to a FiringRule, and where FiringGuards selection criteria match.
Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAPIKEY a valid API key
 * @param optional nil or *GetFiringGuardsUsingPOSTOpts - Optional Parameters:
 * @param "SearchParam" (optional.Interface of FiringGuardGetRequest) -  Search parameters to select FiringGuard
@return []FiringGuard
*/
func (a *EventProcessingFiringApiService) GetFiringGuardsUsingPOST(ctx _context.Context, xAPIKEY string, localVarOptionals *GetFiringGuardsUsingPOSTOpts) ([]FiringGuard, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []FiringGuard
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v0/eventprocessing/firing-guard/search"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarHeaderParams["X-API-KEY"] = parameterToString(xAPIKEY, "")
	// body params
	if localVarOptionals != nil && localVarOptionals.SearchParam.IsSet() {
		localVarOptionalSearchParam, localVarOptionalSearchParamok := localVarOptionals.SearchParam.Value().(FiringGuardGetRequest)
		if !localVarOptionalSearchParamok {
			return localVarReturnValue, nil, reportError("searchParam should be FiringGuardGetRequest")
		}
		localVarPostBody = &localVarOptionalSearchParam
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
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

/*
GetUsingGET15 Retrieve a FiringRule
Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param firingRuleId id of the FiringRule to retrieve
 * @param xAPIKEY a valid API key
@return FiringRule
*/
func (a *EventProcessingFiringApiService) GetUsingGET15(ctx _context.Context, firingRuleId string, xAPIKEY string) (FiringRule, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FiringRule
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v0/eventprocessing/firing-rule/{firingRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"firingRuleId"+"}", _neturl.PathEscape(parameterToString(firingRuleId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarHeaderParams["X-API-KEY"] = parameterToString(xAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
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

// ListUsingGET18Opts Optional parameters for the method 'ListUsingGET18'
type ListUsingGET18Opts struct {
    Name optional.String
}

/*
ListUsingGET18 Retrieve the list of all the FiringRules or get a FiringRule by its name
Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAPIKEY a valid API key
 * @param optional nil or *ListUsingGET18Opts - Optional Parameters:
 * @param "Name" (optional.String) -  name of the FiringRule to retrieve
@return []FiringRule
*/
func (a *EventProcessingFiringApiService) ListUsingGET18(ctx _context.Context, xAPIKEY string, localVarOptionals *ListUsingGET18Opts) ([]FiringRule, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []FiringRule
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v0/eventprocessing/firing-rule"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Name.IsSet() {
		localVarQueryParams.Add("name", parameterToString(localVarOptionals.Name.Value(), ""))
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
	localVarHeaderParams["X-API-KEY"] = parameterToString(xAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
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

// PostUsingPOST14Opts Optional parameters for the method 'PostUsingPOST14'
type PostUsingPOST14Opts struct {
    FiringRule optional.Interface
}

/*
PostUsingPOST14 Create a FiringRule
Total number of FiringRules is limited. Contact the commercial team or see developer guide to get more information.&lt;br&gt;&lt;br&gt;Usage of this API will be reported in your access log under &#39;alarming&#39; category.&lt;br&gt;&lt;br&gt;Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAPIKEY a valid API key
 * @param optional nil or *PostUsingPOST14Opts - Optional Parameters:
 * @param "FiringRule" (optional.Interface of FiringRule) -  FiringRule to add
@return FiringRule
*/
func (a *EventProcessingFiringApiService) PostUsingPOST14(ctx _context.Context, xAPIKEY string, localVarOptionals *PostUsingPOST14Opts) (FiringRule, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FiringRule
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v0/eventprocessing/firing-rule"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarHeaderParams["X-API-KEY"] = parameterToString(xAPIKEY, "")
	// body params
	if localVarOptionals != nil && localVarOptionals.FiringRule.IsSet() {
		localVarOptionalFiringRule, localVarOptionalFiringRuleok := localVarOptionals.FiringRule.Value().(FiringRule)
		if !localVarOptionalFiringRuleok {
			return localVarReturnValue, nil, reportError("firingRule should be FiringRule")
		}
		localVarPostBody = &localVarOptionalFiringRule
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
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

/*
RemoveFiringGuardUsingDELETE Remove a FiringGuard
Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param firingGuardId id of the FiringGuard to remove
 * @param xAPIKEY a valid API key
*/
func (a *EventProcessingFiringApiService) RemoveFiringGuardUsingDELETE(ctx _context.Context, firingGuardId string, xAPIKEY string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v0/eventprocessing/firing-guard/{firingGuardId}"
	localVarPath = strings.Replace(localVarPath, "{"+"firingGuardId"+"}", _neturl.PathEscape(parameterToString(firingGuardId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarHeaderParams["X-API-KEY"] = parameterToString(xAPIKEY, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
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

// RemoveFiringGuardsUsingDELETEOpts Optional parameters for the method 'RemoveFiringGuardsUsingDELETE'
type RemoveFiringGuardsUsingDELETEOpts struct {
    ResetParams optional.Interface
}

/*
RemoveFiringGuardsUsingDELETE Remove the FiringGuards linked to FiringRule, and where FiringGuards selection criteria match.
Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAPIKEY a valid API key
 * @param optional nil or *RemoveFiringGuardsUsingDELETEOpts - Optional Parameters:
 * @param "ResetParams" (optional.Interface of FiringGuardResetRequest) -  search parameters to select FiringGuard to remove
@return int32
*/
func (a *EventProcessingFiringApiService) RemoveFiringGuardsUsingDELETE(ctx _context.Context, xAPIKEY string, localVarOptionals *RemoveFiringGuardsUsingDELETEOpts) (int32, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  int32
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v0/eventprocessing/firing-guard"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarHeaderParams["X-API-KEY"] = parameterToString(xAPIKEY, "")
	// body params
	if localVarOptionals != nil && localVarOptionals.ResetParams.IsSet() {
		localVarOptionalResetParams, localVarOptionalResetParamsok := localVarOptionals.ResetParams.Value().(FiringGuardResetRequest)
		if !localVarOptionalResetParamsok {
			return localVarReturnValue, nil, reportError("resetParams should be FiringGuardResetRequest")
		}
		localVarPostBody = &localVarOptionalResetParams
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
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

// UpdateUsingPUT8Opts Optional parameters for the method 'UpdateUsingPUT8'
type UpdateUsingPUT8Opts struct {
    FiringRule optional.Interface
}

/*
UpdateUsingPUT8 Update a FiringRule
Usage of this API will be reported in your access log under &#39;alarming&#39; category.&lt;br&gt;&lt;br&gt;Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param firingRuleId id of the FiringRule to update
 * @param xAPIKEY a valid API key
 * @param optional nil or *UpdateUsingPUT8Opts - Optional Parameters:
 * @param "FiringRule" (optional.Interface of FiringRule) -  updated FiringRule
*/
func (a *EventProcessingFiringApiService) UpdateUsingPUT8(ctx _context.Context, firingRuleId string, xAPIKEY string, localVarOptionals *UpdateUsingPUT8Opts) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v0/eventprocessing/firing-rule/{firingRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"firingRuleId"+"}", _neturl.PathEscape(parameterToString(firingRuleId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	localVarHeaderParams["X-API-KEY"] = parameterToString(xAPIKEY, "")
	// body params
	if localVarOptionals != nil && localVarOptionals.FiringRule.IsSet() {
		localVarOptionalFiringRule, localVarOptionalFiringRuleok := localVarOptionals.FiringRule.Value().(FiringRule)
		if !localVarOptionalFiringRuleok {
			return nil, reportError("firingRule should be FiringRule")
		}
		localVarPostBody = &localVarOptionalFiringRule
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(_bytes.NewBuffer(localVarBody))
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
