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
// TenantMonthMetrics struct for TenantMonthMetrics
type TenantMonthMetrics struct {
	Tenant TenantExternalView `json:"tenant,omitempty"`
	// Statistics per month
	Months []MonthMetrics `json:"months,omitempty"`
}
