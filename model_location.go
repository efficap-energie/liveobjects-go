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
// Location geographic location
type Location struct {
	// Accuracy value
	Accuracy float64 `json:"accuracy,omitempty"`
	// Altitude value
	Alt float64 `json:"alt,omitempty"`
	// Latitude value
	Lat float64 `json:"lat"`
	// Longitude value
	Lon float64 `json:"lon"`
	// Provider of the location
	Provider string `json:"provider,omitempty"`
}