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

// DataManagementDataStoreApiService DataManagementDataStoreApi service
type DataManagementDataStoreApiService service

type ApiAddDataMessageUsingPOSTRequest struct {
	ctx _context.Context
	ApiService *DataManagementDataStoreApiService
	streamId string
	xAPIKEY *string
	data *DataWeb
}

func (r ApiAddDataMessageUsingPOSTRequest) XAPIKEY(xAPIKEY string) ApiAddDataMessageUsingPOSTRequest {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiAddDataMessageUsingPOSTRequest) Data(data DataWeb) ApiAddDataMessageUsingPOSTRequest {
	r.data = &data
	return r
}

func (r ApiAddDataMessageUsingPOSTRequest) Execute() (string, *_nethttp.Response, error) {
	return r.ApiService.AddDataMessageUsingPOSTExecute(r)
}

/*
 * AddDataMessageUsingPOST Insert a new Data into the stream
 * In order for the data to be enriched by decoders and pipelines functionalities, 'metadata.encoding' field must be set.<br><br>Restricted to API keys with at least one of the following roles : DATA_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param streamId StreamId in which the data will be added. Should not contains following character or space : ' \\ \" ; { } ( )
 * @return ApiAddDataMessageUsingPOSTRequest
 */
func (a *DataManagementDataStoreApiService) AddDataMessageUsingPOST(ctx _context.Context, streamId string) ApiAddDataMessageUsingPOSTRequest {
	return ApiAddDataMessageUsingPOSTRequest{
		ApiService: a,
		ctx: ctx,
		streamId: streamId,
	}
}

/*
 * Execute executes the request
 * @return string
 */
func (a *DataManagementDataStoreApiService) AddDataMessageUsingPOSTExecute(r ApiAddDataMessageUsingPOSTRequest) (string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataManagementDataStoreApiService.AddDataMessageUsingPOST")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/data/streams/{streamId}"
	localVarPath = strings.Replace(localVarPath, "{"+"streamId"+"}", _neturl.PathEscape(parameterToString(r.streamId, "")), -1)

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
	localVarPostBody = r.data
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

type ApiRetrieveDataUsingGETRequest struct {
	ctx _context.Context
	ApiService *DataManagementDataStoreApiService
	streamId string
	xAPIKEY *string
	limit *int32
	timeRange *[]string
	bookmarkId *string
}

func (r ApiRetrieveDataUsingGETRequest) XAPIKEY(xAPIKEY string) ApiRetrieveDataUsingGETRequest {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiRetrieveDataUsingGETRequest) Limit(limit int32) ApiRetrieveDataUsingGETRequest {
	r.limit = &limit
	return r
}
func (r ApiRetrieveDataUsingGETRequest) TimeRange(timeRange []string) ApiRetrieveDataUsingGETRequest {
	r.timeRange = &timeRange
	return r
}
func (r ApiRetrieveDataUsingGETRequest) BookmarkId(bookmarkId string) ApiRetrieveDataUsingGETRequest {
	r.bookmarkId = &bookmarkId
	return r
}

func (r ApiRetrieveDataUsingGETRequest) Execute() ([]DataStoredWeb, *_nethttp.Response, error) {
	return r.ApiService.RetrieveDataUsingGETExecute(r)
}

/*
 * RetrieveDataUsingGET Retrieve data from the streamId
 * return an array of StoredDataMessage matching the request parameters.<br><br>Restricted to API keys with at least one of the following roles : DATA_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param streamId StreamId from which the data will be retrieved
 * @return ApiRetrieveDataUsingGETRequest
 */
func (a *DataManagementDataStoreApiService) RetrieveDataUsingGET(ctx _context.Context, streamId string) ApiRetrieveDataUsingGETRequest {
	return ApiRetrieveDataUsingGETRequest{
		ApiService: a,
		ctx: ctx,
		streamId: streamId,
	}
}

/*
 * Execute executes the request
 * @return []DataStoredWeb
 */
func (a *DataManagementDataStoreApiService) RetrieveDataUsingGETExecute(r ApiRetrieveDataUsingGETRequest) ([]DataStoredWeb, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []DataStoredWeb
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataManagementDataStoreApiService.RetrieveDataUsingGET")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/data/streams/{streamId}"
	localVarPath = strings.Replace(localVarPath, "{"+"streamId"+"}", _neturl.PathEscape(parameterToString(r.streamId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return localVarReturnValue, nil, reportError("xAPIKEY is required and must be specified")
	}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.timeRange != nil {
		t := *r.timeRange
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("timeRange", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("timeRange", parameterToString(t, "multi"))
		}
	}
	if r.bookmarkId != nil {
		localVarQueryParams.Add("bookmarkId", parameterToString(*r.bookmarkId, ""))
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
