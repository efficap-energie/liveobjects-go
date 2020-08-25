# HttpPushAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | A raw string or a mustache (https://mustache.github.io/mustache.5.html) template describing the webhook body. If empty, the raw event will be used. | [optional] 
**Headers** | [**map[string][]string**](array.md) | A map of header to be passed in the http POST headers | [optional] 
**JsonPath** | **string** | The json path to extract from the considered message (or event), it will be taken as the root datacontext object when combined with a mustache template in content | [optional] 
**RetryOnFailure** | **bool** | Indicate if a retry policy should be set up in case of a http push delivery failure | [optional] 
**WebhookUrl** | **string** | The url of the target service webhook (only the ports 80, 443, 8080, 8443 and 9243 are allowed). | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


