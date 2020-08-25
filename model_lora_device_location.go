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
// LoraDeviceLocation struct for LoraDeviceLocation
type LoraDeviceLocation struct {
	// Last global geolocation tolerance (meter)
	Accuracy float64 `json:"accuracy,omitempty"`
	// Last geolocation altitude (meter)
	Alt float64 `json:"alt,omitempty"`
	// Date/time of the last location
	LastUpdateTs string `json:"lastUpdateTs,omitempty"`
	// Last geolocation latitude (GPS coordinate system)
	Lat float64 `json:"lat,omitempty"`
	// Last geolocation longitude (GPS coordinate system)
	Lon float64 `json:"lon,omitempty"`
	// Computing geolocation method
	Provider string `json:"provider,omitempty"`
}
