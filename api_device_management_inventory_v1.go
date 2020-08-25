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

// DeviceManagementInventoryV1ApiService DeviceManagementInventoryV1Api service
type DeviceManagementInventoryV1ApiService service

// CreateDeviceUsingPOSTOpts Optional parameters for the method 'CreateDeviceUsingPOST'
type CreateDeviceUsingPOSTOpts struct {
    Body optional.Interface
}

/*
CreateDeviceUsingPOST Create a device
Please refer to the &#39;Device Management &gt; Interfaces&#39; API notes for more information about &#39;interfaces.[x].definition&#39; content&lt;br&gt;&lt;br&gt;Usage of this API will be reported in your access log under &#39;device_inventory&#39; category.&lt;br&gt;&lt;br&gt;Restricted to API keys with at least one of the following roles : DEVICE_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAPIKEY a valid API key
 * @param optional nil or *CreateDeviceUsingPOSTOpts - Optional Parameters:
 * @param "Body" (optional.Interface of DeviceCreateRequest) -  The device to register
@return Device
*/
func (a *DeviceManagementInventoryV1ApiService) CreateDeviceUsingPOST(ctx _context.Context, xAPIKEY string, localVarOptionals *CreateDeviceUsingPOSTOpts) (Device, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Device
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/deviceMgt/devices"
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
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		localVarOptionalBody, localVarOptionalBodyok := localVarOptionals.Body.Value().(DeviceCreateRequest)
		if !localVarOptionalBodyok {
			return localVarReturnValue, nil, reportError("body should be DeviceCreateRequest")
		}
		localVarPostBody = &localVarOptionalBody
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
		if localVarHTTPResponse.StatusCode == 409 {
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

/*
DeleteDeviceUsingDELETE Delete a device
Usage of this API will be reported in your access log under &#39;device_inventory&#39; category.&lt;br&gt;&lt;br&gt;Restricted to API keys with at least one of the following roles : DEVICE_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param deviceId Target device identifier. A Live Objects URN 'urn:lo:nsid:{namespace}:{id}' must respect the following regular expression <code>^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$</code> (max 269 characters)
 * @param xAPIKEY a valid API key
*/
func (a *DeviceManagementInventoryV1ApiService) DeleteDeviceUsingDELETE(ctx _context.Context, deviceId string, xAPIKEY string) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/deviceMgt/devices/{deviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", _neturl.PathEscape(parameterToString(deviceId, "")) , -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponseWeb
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
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

// GetDeviceStreamsUsingGETOpts Optional parameters for the method 'GetDeviceStreamsUsingGET'
type GetDeviceStreamsUsingGETOpts struct {
    Limit optional.Int32
}

/*
GetDeviceStreamsUsingGET Get a device's streamIds
Restricted to API keys with at least one of the following roles : DATA_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param deviceId Targeted device identifier. A Live Objects URN 'urn:lo:nsid:{namespace}:{id}' must respect the following regular expression <code>^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$</code> (max 269 characters)
 * @param xAPIKEY a valid API key
 * @param optional nil or *GetDeviceStreamsUsingGETOpts - Optional Parameters:
 * @param "Limit" (optional.Int32) -  maximum number of return items (optional, max 100 items)
@return []DeviceStreamsResponseWeb
*/
func (a *DeviceManagementInventoryV1ApiService) GetDeviceStreamsUsingGET(ctx _context.Context, deviceId string, xAPIKEY string, localVarOptionals *GetDeviceStreamsUsingGETOpts) ([]DeviceStreamsResponseWeb, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []DeviceStreamsResponseWeb
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/deviceMgt/devices/{deviceId}/data/streams"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", _neturl.PathEscape(parameterToString(deviceId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
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

/*
GetDeviceUsingGET2 Get a device
Restricted to API keys with at least one of the following roles : DEVICE_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param deviceId device identifier. A Live Objects URN 'urn:lo:nsid:{namespace}:{id}' must respect the following regular expression <code>^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$</code> (max 269 characters)
 * @param xAPIKEY a valid API key
@return Device
*/
func (a *DeviceManagementInventoryV1ApiService) GetDeviceUsingGET2(ctx _context.Context, deviceId string, xAPIKEY string) (Device, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Device
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/deviceMgt/devices/{deviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", _neturl.PathEscape(parameterToString(deviceId, "")) , -1)

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

// ListDevicesUsingGET2Opts Optional parameters for the method 'ListDevicesUsingGET2'
type ListDevicesUsingGET2Opts struct {
    Limit optional.Int32
    Offset optional.Int32
    Sort optional.Interface
    Id optional.String
    GroupPath optional.String
    GroupId optional.String
    Name optional.String
    Tags optional.Interface
    Connectors optional.Interface
    InterfacesNodeId optional.String
    InterfacesStatus optional.Interface
    InterfacesEnabled optional.Bool
    PropertyFilterName optional.String
    ActivityStates optional.Interface
    FilterQuery optional.String
    Fields optional.Interface
    XTotalCount optional.Bool
}

/*
ListDevicesUsingGET2 List registered devices
Restricted to API keys with at least one of the following roles : DEVICE_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xAPIKEY a valid API key
 * @param optional nil or *ListDevicesUsingGET2Opts - Optional Parameters:
 * @param "Limit" (optional.Int32) -  the maximum number of items per page (optional, highest value is 1000)
 * @param "Offset" (optional.Int32) -  number of items to skip (optional)
 * @param "Sort" (optional.Interface of []string) -  sorting list by attributes. Supported columns: id, name, created, updated, group, interfaces.status, interfaces.enabled, interfaces.lastContact). Example: [\"urn\",\"-created\"]. 
 * @param "Id" (optional.String) -  filter list by device identifier (regexp). Expected string (max 269 characters)
 * @param "GroupPath" (optional.String) -  filter list by device groupPath.  (with optional use of wildcard '/_*' at the end of search term)Expected string (max 255 characters)
 * @param "GroupId" (optional.String) -  filter list by device groupId. Expected string (max 6 characters)
 * @param "Name" (optional.String) -  filter list by device name.  (with optional use of wildcard '*' at the beginning or end of search term)Expected string (max 255 characters)
 * @param "Tags" (optional.Interface of []string) -  filter list by device tags. Max number of tags depends of your offer settings. Tag value max length is 32.
 * @param "Connectors" (optional.Interface of []string) -  list devices with interfaces of the specified connector(s). precede the connector with \"-\" to exclude it, use \"none\" to list devices with no interfaces. Example: \"mqtt, -lora, none\".
 * @param "InterfacesNodeId" (optional.String) -  filter list by nodeId. Expected string (max 269 characters)
 * @param "InterfacesStatus" (optional.Interface of []string) -  filter list by interface status. Supported values: REGISTERED, INITIALIZING, INITIALIZED, ONLINE, OFFLINE, ACTIVATED, REACTIVATED, DEACTIVATED, CONNECTIVITY_ERROR, DELETED.
 * @param "InterfacesEnabled" (optional.Bool) -  filter list by interface enabled state.
 * @param "PropertyFilterName" (optional.String) -  Multiple filters, Example : devices?property.temperature=25&property.humidity=58...[cannot be tested in swagger-ui]. Max number of properties depends of your offer settings. A property name must not include following characters <code>$.</code> and max length is 128. Invalid property names are : 'class', '_class'. Property value max length is 256.
 * @param "ActivityStates" (optional.Interface of []string) -  filter list by activity state. Supported values: ACTIVE, SILENT, UNKNOWN, NOT_MONITORED
 * @param "FilterQuery" (optional.String) -  device filter expression using RSQL notation.  Supported device properties are 'groupPath', 'groupId', 'tags', 'connector', 'properties'. Supported RSQL operators are '==','!=','=in=','=out=','=re=','=lt=','=le=','=gt=','=ge=','and','or'. Expected string (max 1000 characters)
 * @param "Fields" (optional.Interface of []string) -  list of fields to return.  Amongst: 'id', 'name', 'description', 'tags', 'properties', 'group', 'interfaces', 'config', 'firmwares', 'activityState', 'defaultDataStreamId', 'created', 'updated'), default: id, name, tags & group
 * @param "XTotalCount" (optional.Bool) -  true if a total count must be returned in response
@return []Device
*/
func (a *DeviceManagementInventoryV1ApiService) ListDevicesUsingGET2(ctx _context.Context, xAPIKEY string, localVarOptionals *ListDevicesUsingGET2Opts) ([]Device, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Device
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/deviceMgt/devices"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		t:=localVarOptionals.Sort.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("sort", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("sort", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.Id.IsSet() {
		localVarQueryParams.Add("id", parameterToString(localVarOptionals.Id.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GroupPath.IsSet() {
		localVarQueryParams.Add("groupPath", parameterToString(localVarOptionals.GroupPath.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GroupId.IsSet() {
		localVarQueryParams.Add("groupId", parameterToString(localVarOptionals.GroupId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Name.IsSet() {
		localVarQueryParams.Add("name", parameterToString(localVarOptionals.Name.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Tags.IsSet() {
		t:=localVarOptionals.Tags.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("tags", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("tags", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.Connectors.IsSet() {
		t:=localVarOptionals.Connectors.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("connectors", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("connectors", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.InterfacesNodeId.IsSet() {
		localVarQueryParams.Add("interfaces.nodeId", parameterToString(localVarOptionals.InterfacesNodeId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.InterfacesStatus.IsSet() {
		t:=localVarOptionals.InterfacesStatus.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("interfaces.status", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("interfaces.status", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.InterfacesEnabled.IsSet() {
		localVarQueryParams.Add("interfaces.enabled", parameterToString(localVarOptionals.InterfacesEnabled.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PropertyFilterName.IsSet() {
		localVarQueryParams.Add("property.{filterName}", parameterToString(localVarOptionals.PropertyFilterName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ActivityStates.IsSet() {
		t:=localVarOptionals.ActivityStates.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("activityStates", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("activityStates", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.FilterQuery.IsSet() {
		localVarQueryParams.Add("filterQuery", parameterToString(localVarOptionals.FilterQuery.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		t:=localVarOptionals.Fields.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("fields", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("fields", parameterToString(t, "multi"))
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
	if localVarOptionals != nil && localVarOptionals.XTotalCount.IsSet() {
		localVarHeaderParams["X-Total-Count"] = parameterToString(localVarOptionals.XTotalCount.Value(), "")
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
UpdateDeviceUsingPATCH Update a device
Usage of this API will be reported in your access log under &#39;device_inventory&#39; category.&lt;br&gt;&lt;br&gt;Restricted to API keys with at least one of the following roles : DEVICE_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param deviceId Targeted device identifier. A Live Objects URN 'urn:lo:nsid:{namespace}:{id}' must respect the following regular expression <code>^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$</code> (max 269 characters)
 * @param xAPIKEY a valid API key
 * @param deviceUpdate deviceUpdate
@return ErrorResponseWeb
*/
func (a *DeviceManagementInventoryV1ApiService) UpdateDeviceUsingPATCH(ctx _context.Context, deviceId string, xAPIKEY string, deviceUpdate DeviceUpdate) (ErrorResponseWeb, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ErrorResponseWeb
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/deviceMgt/devices/{deviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"deviceId"+"}", _neturl.PathEscape(parameterToString(deviceId, "")) , -1)

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
	localVarPostBody = &deviceUpdate
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