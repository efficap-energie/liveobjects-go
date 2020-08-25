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
// Device struct for Device
type Device struct {
	// Device URN
	Id string `json:"id"`
	// Human readable name
	Name string `json:"name,omitempty"`
	// Device description
	Description string `json:"description,omitempty"`
	// Device tags
	Tags []string `json:"tags,omitempty"`
	// Device properties (from device provisioning)
	Properties map[string]string `json:"properties,omitempty"`
	Group DeviceGroup `json:"group,omitempty"`
	// List of this device's interfaces (i.e. 'connectivity nodes')
	Interfaces []DeviceInterface `json:"interfaces,omitempty"`
	// Device configuration (last reported parameter values)
	Config map[string]DeviceParameterValue `json:"config,omitempty"`
	// Device firmware versions
	Firmwares map[string]string `json:"firmwares,omitempty"`
	// default data streamId
	DefaultDataStreamId string `json:"defaultDataStreamId,omitempty"`
	// Date/time when device was first registered
	Created string `json:"created"`
	// Date/time when device status has been lastly updated
	Updated string `json:"updated,omitempty"`
	// Activity state of the device according to the activity rules set for this device
	ActivityState string `json:"activityState,omitempty"`
}