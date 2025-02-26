# TriggerPipelineParameters

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | The branch where the pipeline ran. The HEAD commit on this branch was used for the pipeline. Note that &#x60;branch&#x60; and &#x60;tag&#x60; are mutually exclusive. To trigger a pipeline for a PR by number use &#x60;pull/&lt;number&gt;/head&#x60; for the PR ref or &#x60;pull/&lt;number&gt;/merge&#x60; for the merge ref (GitHub only). | [optional] [default to null]
**Tag** | **string** | The tag used by the pipeline. The commit that this tag points to was used for the pipeline. Note that &#x60;branch&#x60; and &#x60;tag&#x60; are mutually exclusive. | [optional] [default to null]
**Parameters** | [**map[string]Object**](.md) | An object containing pipeline parameters and their values. Pipeline parameters have the following size limits: 100 max entries, 128 maximum key length, 512 maximum value length. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

