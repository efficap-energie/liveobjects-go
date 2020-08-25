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
// AssetSetTargetResourceVersionReqWeb Set Target Resource Version data in Set Request
type AssetSetTargetResourceVersionReqWeb struct {
	// the asset's resource target version. Expected string (max 255 characters)
	TargetVersion string `json:"targetVersion"`
}
