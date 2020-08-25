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
// LoraDevice struct for LoraDevice
type LoraDevice struct {
	// LoRa device activation type
	ActivationType string `json:"activationType"`
	// Application identifier (AppEUI)
	AppEUI string `json:"appEUI"`
	ConnectivityOptions LoraConnectivityOptions `json:"connectivityOptions"`
	// Device connectivity plan.
	ConnectivityPlan string `json:"connectivityPlan,omitempty"`
	// Date/time of the device creation
	CreationTs string `json:"creationTs,omitempty"`
	// LoRa End-device identifier (DevEUI)
	DevEUI string `json:"devEUI"`
	// Device status
	DeviceStatus string `json:"deviceStatus"`
	// LoRa encoder
	Encoding string `json:"encoding"`
	// Date/time of the last activation
	LastActivationTs string `json:"lastActivationTs,omitempty"`
	// Last battery level returned by the device (0: external power source, 1..254: 1=min / 254 = max, 255 : NA)
	LastBatteryLevel int32 `json:"lastBatteryLevel,omitempty"`
	// Date/time of the last communication
	LastCommunicationTs string `json:"lastCommunicationTs,omitempty"`
	// Date/time of the last deactivation
	LastDeactivationTs string `json:"lastDeactivationTs,omitempty"`
	// Frame counter of the last downlink
	LastDlFcnt int32 `json:"lastDlFcnt,omitempty"`
	// Signal level of the last uplink (between 0 to 5)
	LastSignalLevel int32 `json:"lastSignalLevel,omitempty"`
	// Frame counter of the last uplink
	LastUlFcnt int32 `json:"lastUlFcnt,omitempty"`
	Location LoraDeviceLocation `json:"location,omitempty"`
	// LoRa device name
	Name string `json:"name"`
	// LoRa device profile
	Profile string `json:"profile"`
	// List of tags, used to tag uplink from this device
	Tags []string `json:"tags,omitempty"`
	// Date/time of the last device modification
	UpdateTs string `json:"updateTs,omitempty"`
}
