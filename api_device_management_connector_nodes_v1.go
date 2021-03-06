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

// DeviceManagementConnectorNodesV1ApiService DeviceManagementConnectorNodesV1Api service
type DeviceManagementConnectorNodesV1ApiService service

type ApiDeleteNodeUsingDELETERequest struct {
	ctx _context.Context
	ApiService *DeviceManagementConnectorNodesV1ApiService
	connector string
	nodeId string
	xAPIKEY *string
}

func (r ApiDeleteNodeUsingDELETERequest) XAPIKEY(xAPIKEY string) ApiDeleteNodeUsingDELETERequest {
	r.xAPIKEY = &xAPIKEY
	return r
}

func (r ApiDeleteNodeUsingDELETERequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteNodeUsingDELETEExecute(r)
}

/*
 * DeleteNodeUsingDELETE Delete a connector node
 * Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connector the connector id
 * @param nodeId the node id
 * @return ApiDeleteNodeUsingDELETERequest
 */
func (a *DeviceManagementConnectorNodesV1ApiService) DeleteNodeUsingDELETE(ctx _context.Context, connector string, nodeId string) ApiDeleteNodeUsingDELETERequest {
	return ApiDeleteNodeUsingDELETERequest{
		ApiService: a,
		ctx: ctx,
		connector: connector,
		nodeId: nodeId,
	}
}

/*
 * Execute executes the request
 */
func (a *DeviceManagementConnectorNodesV1ApiService) DeleteNodeUsingDELETEExecute(r ApiDeleteNodeUsingDELETERequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementConnectorNodesV1ApiService.DeleteNodeUsingDELETE")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/deviceMgt/connectors/{connector}/nodes/{nodeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"connector"+"}", _neturl.PathEscape(parameterToString(r.connector, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodeId"+"}", _neturl.PathEscape(parameterToString(r.nodeId, "")), -1)

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
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
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

type ApiGetNodeUsingGETRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementConnectorNodesV1ApiService
	connector string
	nodeId string
	xAPIKEY *string
}

func (r ApiGetNodeUsingGETRequest) XAPIKEY(xAPIKEY string) ApiGetNodeUsingGETRequest {
	r.xAPIKEY = &xAPIKEY
	return r
}

func (r ApiGetNodeUsingGETRequest) Execute() (ConnectorNode, *_nethttp.Response, error) {
	return r.ApiService.GetNodeUsingGETExecute(r)
}

/*
 * GetNodeUsingGET Get details of connector node
 * <p>The <strong>definition</strong> depends on connector.</p><p><span style="text-decoration: underline;">Lora definition</span>:</p><ul><li>"devEUI": interface EUI</li><li>"appEUI": appEUI of the interface</li><li>"appKey": appKey of the interface</li><li>"activationType": supported value: "OTAA"</li><li>"profile": profile of the interface</li><li>"encoding": (Optional) encoding used in message's payload (use /decoders APIs to define your own decoder)</li><li>"connectivityOptions": connectivity options used for the interface </li><ul><li>"tdoa": true/false </li><li>"ackUl": true/false</li></ul><li>"connectivityPlan": connectivity plan to use for the interface. Should be used preferably at connectivityOptions but at least one of two shall be defined.</li></ul><p><span style="text-decoration: underline;">SMS definition</span>:</p><ul><li>"msisdn": interface msisdn</li><li>"serverPhoneNumber": (Optional) server phone number</li><li>"encoding": (Optional) name of the decoder that will be used to decode received SMS</li></ul><p><span style="text-decoration: underline;">MQTT definition</span>:</p><ul><li>"clientId": interface client ID</li><li>"encoding": (Optional) encoding used in message's payload (use /decoders APIs to define your own decoder)</li></ul><p><span style="text-decoration: underline;">X-CONNECTOR definition</span>:</p><ul><li>"nodeId": id used by the external connector to identify the communicating device (e.g. PUB connector/v1/nodes/{nodeId}/status)</li><li>"encoding": (Optional) encoding used in message's payload (use /decoders APIs to define your own decoder)</li></ul><br /><p>The <strong>activity</strong> depends on connector.</p><p><span style="text-decoration: underline;">Lora activity</span>:</p><ul><li>"lastActivationTs": (Optional) last activation date of the interface</li><li>"lastDeactivationTs": (Optional) last deactivation date of the interface</li><li>"lastSignalLevel": (Optional) last signal level</li><li>"lastBatteryLevel": (Optional) last battery level</li><li>"lastDlFcnt": (Optional) last downlink frame counter</li><li>"lastUlFcnt": (Optional) last uplink frame counter</li></ul><p><span style="text-decoration: underline;">SMS activity</span>:</p><ul><li>"lastUplink": (Optional) last uplink date of the interface</li><ul><li>"timestamp": date of the activity</li><li>"serverPhoneNumber": server phone number used</li></ul><li>"lastDownlink": (Optional) last downlink date of the connector node</li><ul><li>"timestamp": date of the activity</li><li>"serverPhoneNumber": server phone number used</li></ul></ul><p><span style="text-decoration: underline;">MQTT activity</span>:</p><ul><li>"apiKeyId": api key ID used</li><li>"mqttVersion": mqtt version used</li><li>"mqttUsername": mqtt username used</li><li>"mqttTimeout": mqtt timeout</li><li>"remoteAddress"</li><li>"lastSessionStartTime"</li><li>"lastSessionEndTime"</li></ul><br><br>Restricted to API keys with at least one of the following roles : DEVICE_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connector the connector
 * @param nodeId the node id
 * @return ApiGetNodeUsingGETRequest
 */
func (a *DeviceManagementConnectorNodesV1ApiService) GetNodeUsingGET(ctx _context.Context, connector string, nodeId string) ApiGetNodeUsingGETRequest {
	return ApiGetNodeUsingGETRequest{
		ApiService: a,
		ctx: ctx,
		connector: connector,
		nodeId: nodeId,
	}
}

/*
 * Execute executes the request
 * @return ConnectorNode
 */
func (a *DeviceManagementConnectorNodesV1ApiService) GetNodeUsingGETExecute(r ApiGetNodeUsingGETRequest) (ConnectorNode, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConnectorNode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementConnectorNodesV1ApiService.GetNodeUsingGET")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/deviceMgt/connectors/{connector}/nodes/{nodeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"connector"+"}", _neturl.PathEscape(parameterToString(r.connector, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodeId"+"}", _neturl.PathEscape(parameterToString(r.nodeId, "")), -1)

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

type ApiListNodesUsingGETRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementConnectorNodesV1ApiService
	connector string
	xAPIKEY *string
	limit *int32
	offset *int32
	xTotalCount *bool
}

func (r ApiListNodesUsingGETRequest) XAPIKEY(xAPIKEY string) ApiListNodesUsingGETRequest {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiListNodesUsingGETRequest) Limit(limit int32) ApiListNodesUsingGETRequest {
	r.limit = &limit
	return r
}
func (r ApiListNodesUsingGETRequest) Offset(offset int32) ApiListNodesUsingGETRequest {
	r.offset = &offset
	return r
}
func (r ApiListNodesUsingGETRequest) XTotalCount(xTotalCount bool) ApiListNodesUsingGETRequest {
	r.xTotalCount = &xTotalCount
	return r
}

func (r ApiListNodesUsingGETRequest) Execute() ([]ConnectorNode, *_nethttp.Response, error) {
	return r.ApiService.ListNodesUsingGETExecute(r)
}

/*
 * ListNodesUsingGET List all connector nodes
 * <p>The <strong>definition</strong> depends on connector.</p><p><span style="text-decoration: underline;">Lora definition</span>:</p><ul><li>"devEUI": interface EUI</li><li>"appEUI": appEUI of the interface</li><li>"appKey": appKey of the interface</li><li>"activationType": supported value: "OTAA"</li><li>"profile": profile of the interface</li><li>"encoding": (Optional) encoding used in message's payload (use /decoders APIs to define your own decoder)</li><li>"connectivityOptions": connectivity options used for the interface </li><ul><li>"tdoa": true/false </li><li>"ackUl": true/false</li></ul><li>"connectivityPlan": connectivity plan to use for the interface. Should be used preferably at connectivityOptions but at least one of two shall be defined.</li></ul><p><span style="text-decoration: underline;">SMS definition</span>:</p><ul><li>"msisdn": interface msisdn</li><li>"serverPhoneNumber": (Optional) server phone number</li><li>"encoding": (Optional) name of the decoder that will be used to decode received SMS</li></ul><p><span style="text-decoration: underline;">MQTT definition</span>:</p><ul><li>"clientId": interface client ID</li><li>"encoding": (Optional) encoding used in message's payload (use /decoders APIs to define your own decoder)</li></ul><p><span style="text-decoration: underline;">X-CONNECTOR definition</span>:</p><ul><li>"nodeId": id used by the external connector to identify the communicating device (e.g. PUB connector/v1/nodes/{nodeId}/status)</li><li>"encoding": (Optional) encoding used in message's payload (use /decoders APIs to define your own decoder)</li></ul><br /><p>The <strong>activity</strong> depends on connector.</p><p><span style="text-decoration: underline;">Lora activity</span>:</p><ul><li>"lastActivationTs": (Optional) last activation date of the interface</li><li>"lastDeactivationTs": (Optional) last deactivation date of the interface</li><li>"lastSignalLevel": (Optional) last signal level</li><li>"lastBatteryLevel": (Optional) last battery level</li><li>"lastDlFcnt": (Optional) last downlink frame counter</li><li>"lastUlFcnt": (Optional) last uplink frame counter</li></ul><p><span style="text-decoration: underline;">SMS activity</span>:</p><ul><li>"lastUplink": (Optional) last uplink date of the interface</li><ul><li>"timestamp": date of the activity</li><li>"serverPhoneNumber": server phone number used</li></ul><li>"lastDownlink": (Optional) last downlink date of the connector node</li><ul><li>"timestamp": date of the activity</li><li>"serverPhoneNumber": server phone number used</li></ul></ul><p><span style="text-decoration: underline;">MQTT activity</span>:</p><ul><li>"apiKeyId": api key ID used</li><li>"mqttVersion": mqtt version used</li><li>"mqttUsername": mqtt username used</li><li>"mqttTimeout": mqtt timeout</li><li>"remoteAddress"</li><li>"lastSessionStartTime"</li><li>"lastSessionEndTime"</li></ul><br><br>Restricted to API keys with at least one of the following roles : DEVICE_R.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connector the connector
 * @return ApiListNodesUsingGETRequest
 */
func (a *DeviceManagementConnectorNodesV1ApiService) ListNodesUsingGET(ctx _context.Context, connector string) ApiListNodesUsingGETRequest {
	return ApiListNodesUsingGETRequest{
		ApiService: a,
		ctx: ctx,
		connector: connector,
	}
}

/*
 * Execute executes the request
 * @return []ConnectorNode
 */
func (a *DeviceManagementConnectorNodesV1ApiService) ListNodesUsingGETExecute(r ApiListNodesUsingGETRequest) ([]ConnectorNode, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []ConnectorNode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementConnectorNodesV1ApiService.ListNodesUsingGET")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/deviceMgt/connectors/{connector}/nodes"
	localVarPath = strings.Replace(localVarPath, "{"+"connector"+"}", _neturl.PathEscape(parameterToString(r.connector, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return localVarReturnValue, nil, reportError("xAPIKEY is required and must be specified")
	}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
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
	if r.xTotalCount != nil {
		localVarHeaderParams["X-Total-Count"] = parameterToString(*r.xTotalCount, "")
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

type ApiUpdateNodeUsingPATCHRequest struct {
	ctx _context.Context
	ApiService *DeviceManagementConnectorNodesV1ApiService
	connector string
	nodeId string
	xAPIKEY *string
	body *UpdateConnectorNodeRequest
}

func (r ApiUpdateNodeUsingPATCHRequest) XAPIKEY(xAPIKEY string) ApiUpdateNodeUsingPATCHRequest {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiUpdateNodeUsingPATCHRequest) Body(body UpdateConnectorNodeRequest) ApiUpdateNodeUsingPATCHRequest {
	r.body = &body
	return r
}

func (r ApiUpdateNodeUsingPATCHRequest) Execute() (ConnectorNode, *_nethttp.Response, error) {
	return r.ApiService.UpdateNodeUsingPATCHExecute(r)
}

/*
 * UpdateNodeUsingPATCH Update a connector node
 * <p>The <strong>definition</strong> depends on connector.</p><p><span style="text-decoration: underline;">Lora definition</span>:</p><ul><li>"devEUI": interface EUI</li><li>"appEUI": appEUI of the interface</li><li>"appKey": appKey of the interface</li><li>"activationType": supported value: "OTAA"</li><li>"profile": profile of the interface</li><li>"encoding": (Optional) encoding used in message's payload (use /decoders APIs to define your own decoder)</li><li>"connectivityOptions": connectivity options used for the interface </li><ul><li>"tdoa": true/false </li><li>"ackUl": true/false</li></ul><li>"connectivityPlan": connectivity plan to use for the interface. Should be used preferably at connectivityOptions but at least one of two shall be defined.</li></ul><p><span style="text-decoration: underline;">SMS definition</span>:</p><ul><li>"msisdn": interface msisdn</li><li>"serverPhoneNumber": (Optional) server phone number</li><li>"encoding": (Optional) name of the decoder that will be used to decode received SMS</li></ul><p><span style="text-decoration: underline;">MQTT definition</span>:</p><ul><li>"clientId": interface client ID</li><li>"encoding": (Optional) encoding used in message's payload (use /decoders APIs to define your own decoder)</li></ul><p><span style="text-decoration: underline;">X-CONNECTOR definition</span>:</p><ul><li>"nodeId": id used by the external connector to identify the communicating device (e.g. PUB connector/v1/nodes/{nodeId}/status)</li><li>"encoding": (Optional) encoding used in message's payload (use /decoders APIs to define your own decoder)</li></ul><br><br>Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connector the connector id
 * @param nodeId the node id
 * @return ApiUpdateNodeUsingPATCHRequest
 */
func (a *DeviceManagementConnectorNodesV1ApiService) UpdateNodeUsingPATCH(ctx _context.Context, connector string, nodeId string) ApiUpdateNodeUsingPATCHRequest {
	return ApiUpdateNodeUsingPATCHRequest{
		ApiService: a,
		ctx: ctx,
		connector: connector,
		nodeId: nodeId,
	}
}

/*
 * Execute executes the request
 * @return ConnectorNode
 */
func (a *DeviceManagementConnectorNodesV1ApiService) UpdateNodeUsingPATCHExecute(r ApiUpdateNodeUsingPATCHRequest) (ConnectorNode, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConnectorNode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceManagementConnectorNodesV1ApiService.UpdateNodeUsingPATCH")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/deviceMgt/connectors/{connector}/nodes/{nodeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"connector"+"}", _neturl.PathEscape(parameterToString(r.connector, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"nodeId"+"}", _neturl.PathEscape(parameterToString(r.nodeId, "")), -1)

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
	localVarPostBody = r.body
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
