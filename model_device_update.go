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
// DeviceUpdate an update to a device definition
type DeviceUpdate struct {
	// default data streamId. Expected not empty string. Following character are forbidden <code>\"'\\\"\\\\;{}() \"</code> (max 255 characters)
	DefaultDataStreamId string `json:"defaultDataStreamId,omitempty"`
	// new device description. Expected string (max 500 characters)
	Description string `json:"description,omitempty"`
	Group DeviceGroup `json:"group,omitempty"`
	// new device identifier. A Live Objects URN 'urn:lo:nsid:{namespace}:{id}' must respect the following regular expression <code>^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$</code> (max 269 characters)
	Id string `json:"id"`
	// new device human-readable name. Expected string (max 255 characters)
	Name string `json:"name,omitempty"`
	// map of key/value string pairs detailing device properties to update. Max number of properties depends of your offer settings. A property name must not include following characters <code>$.</code> and max length is 128. Invalid property names are : 'class', '_class'. Property value max length is 256.
	Properties map[string]string `json:"properties,omitempty"`
	// new device set of tags. Max number of tags depends of your offer settings. Tag value max length is 32.
	Tags []string `json:"tags,omitempty"`
}