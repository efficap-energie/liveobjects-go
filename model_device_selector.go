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
// DeviceSelector struct for DeviceSelector
type DeviceSelector struct {
	// Filtering expression using RSQL notation. Supported device properties are 'groupPath', 'groupId', 'tags', 'properties'. Supported RSQL operators are '==','!=','=in=','=out=','=re=','=lt=','=le=','=gt=','=ge=','and','or'
	FilterQuery string `json:"filterQuery,omitempty"`
	// List of device IDs in the form urn:lo:nsid:${assetNamespace}:${assetId}
	IdList []string `json:"idList,omitempty"`
}
