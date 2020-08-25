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
// DeferredResultVoid struct for DeferredResultVoid
type DeferredResultVoid struct {
	Result map[string]interface{} `json:"result,omitempty"`
	SetOrExpired bool `json:"setOrExpired,omitempty"`
}
