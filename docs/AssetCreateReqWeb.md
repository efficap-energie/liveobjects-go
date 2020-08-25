# AssetCreateReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | user-readable description of the asset. Expected string (max 500 characters) | [optional] 
**GroupId** | **string** | associate the asset to a group using the group id. Expected string (max 6 characters) | [optional] 
**GroupPath** | **string** | associate the asset to a group using the group path. Authorized: letter (lowercase and uppercase), accented characters, number, space, dash, underscore and simple quote. A valid path must respect the following regular expression &lt;code&gt;^[\\wÀ-ÖØ-öø-ÿ&#39; -]{1,255}&lt;/code&gt;.Expected string (max 255 characters) | [optional] 
**Id** | **string** | asset identifier value. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
**Name** | **string** | asset user-readable name (for display in web portal).Expected string (max 255 characters) | [optional] 
**Namespace** | **string** | asset identifier namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**Properties** | **map[string]string** | map of key/value string pairs detailed properties of the device. Max number of properties depends of your offer settings. A property name must not include following characters &lt;code&gt;$.&lt;/code&gt; and max length is 128. Invalid property names are : &#39;class&#39;, &#39;_class&#39;. Property value max length is 256. | [optional] 
**Tags** | **[]string** | series of tags associated with the asset. Max number of tags depends of your offer settings. Tag value max length is 32. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


