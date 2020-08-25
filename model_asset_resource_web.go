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
// AssetResourceWeb Resource data associated to an Asset
type AssetResourceWeb struct {
	// the asset id
	AssetId string `json:"assetId,omitempty"`
	// the asset namespace
	AssetIdNamespace string `json:"assetIdNamespace,omitempty"`
	// the connector metadata
	ConnectorMetadata map[string]string `json:"connectorMetadata,omitempty"`
	// the current resource version
	CurrentVersion string `json:"currentVersion,omitempty"`
	// the current resource version timestamp
	CurrentVersionChangeTs int64 `json:"currentVersionChangeTs,omitempty"`
	// the current resource version timestamp
	CurrentVersionTs int64 `json:"currentVersionTs,omitempty"`
	// the resource's identifier
	ResourceId string `json:"resourceId,omitempty"`
	// the target resource version timestamp
	TargetVersion string `json:"targetVersion,omitempty"`
	// the target resource version timestamp
	TargetVersionTs int64 `json:"targetVersionTs,omitempty"`
	// identifier of tenant account
	TenantId string `json:"tenantId,omitempty"`
}