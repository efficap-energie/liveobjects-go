# DeviceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDataStreamId** | **string** | default data streamId. Expected not empty string. Following character are forbidden &lt;code&gt;\&quot;&#39;\\\&quot;\\\\;{}() \&quot;&lt;/code&gt; (max 255 characters) | [optional] 
**Description** | **string** | new device description. Expected string (max 500 characters) | [optional] 
**Group** | [**DeviceGroup**](DeviceGroup.md) |  | [optional] 
**Id** | **string** | new device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**Name** | **string** | new device human-readable name. Expected string (max 255 characters) | [optional] 
**Properties** | **map[string]string** | map of key/value string pairs detailing device properties to update. Max number of properties depends of your offer settings. A property name must not include following characters &lt;code&gt;$.&lt;/code&gt; and max length is 128. Invalid property names are : &#39;class&#39;, &#39;_class&#39;. Property value max length is 256. | [optional] 
**Tags** | **[]string** | new device set of tags. Max number of tags depends of your offer settings. Tag value max length is 32. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


