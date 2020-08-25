# DeviceCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | device unique identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**Name** | **string** | human-readable device name. Expected string (max 255 characters) | [optional] 
**Description** | **string** | human-readable device description. Expected string (max 500 characters) | [optional] 
**DefaultDataStreamId** | **string** | default data streamId. Expected not empty string. Following character are forbidden &lt;code&gt;\&quot;&#39;\\\&quot;\\\\;{}() \&quot;&lt;/code&gt; (max 255 characters) | [optional] 
**Tags** | **[]string** | set of tags associated with the new device. Max number of tags depends of your offer settings. Tag value max length is 32. | [optional] 
**Properties** | **map[string]string** | map of key/value string pairs detailing properties of the device. Max number of properties depends of your offer settings. A property name must not include following characters &lt;code&gt;$.&lt;/code&gt; and max length is 128. Invalid property names are : &#39;class&#39;, &#39;_class&#39;. Property value max length is 256. | [optional] 
**Interfaces** | [**[]NewDeviceInterface**](NewDeviceInterface.md) | list of device network interfaces | [optional] 
**Group** | [**DeviceGroup**](DeviceGroup.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


