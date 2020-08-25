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
// TenantDayMetrics struct for TenantDayMetrics
type TenantDayMetrics struct {
	Tenant TenantExternalView `json:"tenant,omitempty"`
	// Statistics per day
	Days []DayMetrics `json:"days,omitempty"`
}