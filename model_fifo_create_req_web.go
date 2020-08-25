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
// FifoCreateReqWeb  Fifo Topic data in Create Request
type FifoCreateReqWeb struct {
	// Maximum number of bytes that can be stored (0 means no storage, else min value is 524288 and max value depends of your account settings)
	MaxLengthBytes int64 `json:"maxLengthBytes"`
	// Time in seconds before messages are dropped if they are not consumed
	MessageTtl int32 `json:"messageTtl,omitempty"`
	// Topic name of the FIFO. Fifo name must respect the following regular expression <code>[A-Za-z\\d-_~]+</code> (max 255 characters)
	Name string `json:"name"`
}
