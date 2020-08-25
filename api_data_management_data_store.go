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
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// DataManagementDataStoreApiService DataManagementDataStoreApi service
type DataManagementDataStoreApiService service

// AddDataMessageUsingPOSTOpts Optional parameters for the method 'AddDataMessageUsingPOST'
type AddDataMessageUsingPOSTOpts struct {
    Data optional.Interface
}

/*
AddDataMessageUsingPOST Insert a new Data into the stream
In order for the data to be enriched by decoders and pipelines functionalities, &#39;metadata.encoding&#39; field must be set.&lt;br&gt;&lt;br&gt;Restricted to API keys with at least one of the following roles : DATA_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param streamId StreamId in which the data will be added. Should not contains following character or space : ' \\ \" ; { } ( )
 * @param xAPIKEY a valid API key
 * @param optional nil or *AddDataMessageUsingPOSTOpts - Optional Parameters:
 * @param "Data" (optional.Interface of DataWeb) -  DataMessage to add
@return string
*/
func (a *DataManagementDataStoreApiService) AddDataMessageUsingPOST(ctx _context.Context, streamId string, xAPIKEY string, localVarOptionals *AddDataMessageUsingPOSTOpts) (string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  string
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v0/data/streams/{streamId}"
	localVarPath = strings.Replace(localVarPath, "{"+"streamId"+"}", _neturl.PathEscape(parameterToString(streamId, "")) , -1)

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
	if localVarOptionals != nil && localVarOptionals.Data.IsSet() {
		localVarOptionalData, localVarOptionalDataok := localVarOptionals.Data.Value().(DataWeb)
		if !localVarOptionalDataok {
			return localVarReturnValue, nil, reportError("data should be DataWeb")
		}
		localVarPostBody = &localVarOptionalData
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

// RetrieveDataUsingGETOpts Optional parameters for the method 'RetrieveDataUsingGET'
type RetrieveDataUsingGETOpts struct {
    Limit optional.Int32
    TimeRange optional.Interface
    BookmarkId optional.String
}

/*
RetrieveDataUsingGET Retrieve data from the streamId
return an array of StoredDataMessage matching the request parameters.&lt;br&gt;&lt;br&gt;Restricted to API keys with at least one of the following roles : DATA_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param streamId StreamId from which the data will be retrieved
 * @param xAPIKEY a valid API key
 * @param optional nil or *RetrieveDataUsingGETOpts - Optional Parameters:
 * @param "Limit" (optional.Int32) -  max number of data to return, value is limited to 1000
 * @param "TimeRange" (optional.Interface of []string) -  filter data where timestamp is in timeRange (must be ISO 8601 formatted) : ?timeRange=[lowerbound],[upperbound]. Coma is mandatory, lowerbound and upperbound are optionals. Lowerbound is inclusiv, upperbound is exclusiv.
 * @param "BookmarkId" (optional.String) -  id of the last document retrieved that can be used to paginate : first result will be the one following this document id
@return []DataStoredWeb
*/
func (a *DataManagementDataStoreApiService) RetrieveDataUsingGET(ctx _context.Context, streamId string, xAPIKEY string, localVarOptionals *RetrieveDataUsingGETOpts) ([]DataStoredWeb, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []DataStoredWeb
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v0/data/streams/{streamId}"
	localVarPath = strings.Replace(localVarPath, "{"+"streamId"+"}", _neturl.PathEscape(parameterToString(streamId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TimeRange.IsSet() {
		t:=localVarOptionals.TimeRange.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("timeRange", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("timeRange", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.BookmarkId.IsSet() {
		localVarQueryParams.Add("bookmarkId", parameterToString(localVarOptionals.BookmarkId.Value(), ""))
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