# ResourceVersionAddReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checksum** | **string** | The Base64-encoded MD5 checksum of the version&#39;s raw file (non Base64-encoded)Expected string (max 255 characters) | 
**CompatibleVersions** | **[]string** | set of the versions from which a resource update to the version can be done. Max number of versions is 100. A version max length is 255.  | [optional] 
**Description** | **string** | the version&#39;s description. Expected string (max 255 characters) | [optional] 
**File** | **string** | the version&#39;s Base64-encoded binary. File max size is 10MB | 
**ResourceVersionId** | **string** | the resource version id. Expected string (max 255 characters) | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


