# AzureEventHubsAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | A raw string or a mustache (https://mustache.github.io/mustache.5.html) template describing the webhook body. If empty, the raw event will be used. | [optional] 
**EventHubName** | **string** | The name of the targeted Event Hub. | 
**EventHubsNamespace** | **string** | the Event Hubs namespace where is located the targeted Event Hub. | 
**JsonPath** | **string** | The json path to extract from the considered message (or event), it will be taken as the root datacontext object when combined with a mustache template in content | [optional] 
**RetryOnFailure** | **bool** | Indicate if a retry policy should be set up in case of a Event Hub delivery failure | [optional] 
**SharedAccessKey** | **string** | The shared access key (also called shared access policy). | 
**SharedAccessKeyName** | **string** | The name of the shared access key (also called shared access policy) from the targeted Event Hub. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


