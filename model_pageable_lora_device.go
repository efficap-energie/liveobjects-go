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
// PageableLoraDevice struct for PageableLoraDevice
type PageableLoraDevice struct {
	// number of the current page: starts at 0.
	Page int64 `json:"page"`
	// number of data per page (= maximum number of data in the associated list of data:the last page can have less data)
	Size int64 `json:"size"`
	// total count of data in the complete list.
	TotalCount int64 `json:"totalCount"`
	// list of data in this page.
	Data []LoraDevice `json:"data"`
}
