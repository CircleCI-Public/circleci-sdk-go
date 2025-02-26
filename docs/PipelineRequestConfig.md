# PipelineRequestConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | The branch that should be used to fetch the config file. Note that branch and tag are mutually exclusive. To trigger a pipeline for a PR by number use pull/&lt;number&gt;/head for the PR ref or pull/&lt;number&gt;/merge for the merge ref (GitHub only)  | [optional] [default to null]
**Tag** | **string** | The tag that should be used to fetch the config file. The commit that this tag points to is used for the pipeline. Note that branch and tag are mutually exclusive.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

