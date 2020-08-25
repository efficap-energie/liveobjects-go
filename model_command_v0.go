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
// CommandV0 struct for CommandV0
type CommandV0 struct {
	AssetId string `json:"assetId,omitempty"`
	AssetIdNamespace string `json:"assetIdNamespace,omitempty"`
	CreationTs int64 `json:"creationTs,omitempty"`
	Id string `json:"id,omitempty"`
	NotifyTo string `json:"notifyTo,omitempty"`
	ReqData map[string]string `json:"reqData,omitempty"`
	ReqEvent string `json:"reqEvent,omitempty"`
	ReqPayload []string `json:"reqPayload,omitempty"`
	ResData map[string]string `json:"resData,omitempty"`
	ResPayload []string `json:"resPayload,omitempty"`
	Status string `json:"status,omitempty"`
	UpdateTs int64 `json:"updateTs,omitempty"`
}