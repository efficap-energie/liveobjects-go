# PayloadDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | decoder activation. Default is false. | 
**Encoding** | **string** | unique decoder name. Will be used to trigger the decoding function | 
**Id** | **string** | id of the decoder. Should be null when used for POST. Required for update. | [optional] 
**MathEvalEnabled** | **bool** | applying math evaluation on templated decoded result | [optional] 
**Metadata** | [**PayloadDescriptionMetadata**](PayloadDescriptionMetadata.md) |  | [optional] 
**Model** | **string** | if empty, the decoded data will use the value of &#39;model&#39; field from encoded data. If set, this value will be used to override &#39;model&#39; field in decoded data. | [optional] 
**Tags** | **[]string** | tags used for filtering | [optional] 
**Template** | **string** | decoding result optional template | [optional] 
**Type** | **string** | decoder type : csv, binary, javascript | [optional] 
**Updated** | **string** | date of the last update. ISO 8601 date time. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


