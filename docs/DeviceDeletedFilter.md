# DeviceDeletedFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectors** | **[]string** | list of filtered connectors. | [optional] 
**GroupPaths** | [**[]GroupPath**](GroupPath.md) | list of filtered group paths. Null or empty to accept all group paths | [optional] 
**Tags** | [**[][]string**](array.md) | list of groups of tags that should be contained in message tags. There is a match if at least one group of tags is a match. A group of tags is a match if the tags of the message contains all elements of this group.&lt;br&gt;e.g. [[\&quot;ALERT\&quot;]] will match all messages containing tag &#39;ALERT&#39;.&lt;br&gt;e.g. [[\&quot;HIGH\&quot;, \&quot;ALERT\&quot;],[\&quot;PROD\&quot;]] will match messages with either tag &#39;PROD&#39; or both tags &#39;ALERT&#39; and &#39;HIGH&#39;. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


