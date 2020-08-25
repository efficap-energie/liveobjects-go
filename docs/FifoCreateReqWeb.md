# FifoCreateReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxLengthBytes** | **int64** | Maximum number of bytes that can be stored (0 means no storage, else min value is 524288 and max value depends of your account settings) | 
**MessageTtl** | **int32** | Time in seconds before messages are dropped if they are not consumed | [optional] 
**Name** | **string** | Topic name of the FIFO. Fifo name must respect the following regular expression &lt;code&gt;[A-Za-z\\d-_~]+&lt;/code&gt; (max 255 characters) | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


