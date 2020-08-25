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
// FifoBinding struct for FifoBinding
type FifoBinding struct {
	// Destination Fifo Name
	DestFifoName string `json:"destFifoName"`
	// routing Key filter
	RoutingKeyFilter string `json:"routingKeyFilter"`
}