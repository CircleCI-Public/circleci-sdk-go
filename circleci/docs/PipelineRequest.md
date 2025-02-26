# PipelineRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefinitionId** | **string** | The unique id for the pipeline definition. This can be found in the page Project Settings &gt; Pipelines. | [optional] [default to null]
**Config** | [***PipelineRequestConfig**](pipelineRequest_config.md) |  | [optional] [default to null]
**Checkout** | [***PipelineRequestCheckout**](pipelineRequest_checkout.md) |  | [optional] [default to null]
**Parameters** | [**ModelMap**](interface{}.md) | An object containing pipeline parameters and their values. Pipeline parameters have the following size limits: 100 max entries, 128 maximum key length, 512 maximum value length.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

