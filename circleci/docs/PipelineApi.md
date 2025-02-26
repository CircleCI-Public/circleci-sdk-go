# {{classname}}

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContinuePipeline**](PipelineApi.md#ContinuePipeline) | **Post** /pipeline/continue | Continue a pipeline
[**GetPipelineById**](PipelineApi.md#GetPipelineById) | **Get** /pipeline/{pipeline-id} | Get a pipeline by ID
[**GetPipelineByNumber**](PipelineApi.md#GetPipelineByNumber) | **Get** /project/{project-slug}/pipeline/{pipeline-number} | Get a pipeline by pipeline number
[**GetPipelineConfigById**](PipelineApi.md#GetPipelineConfigById) | **Get** /pipeline/{pipeline-id}/config | Get a pipeline&#x27;s configuration
[**GetPipelineValuesById**](PipelineApi.md#GetPipelineValuesById) | **Get** /pipeline/{pipeline-id}/values | Get pipeline values for a pipeline
[**ListMyPipelines**](PipelineApi.md#ListMyPipelines) | **Get** /project/{project-slug}/pipeline/mine | Get your pipelines
[**ListPipelines**](PipelineApi.md#ListPipelines) | **Get** /pipeline | Get a list of pipelines
[**ListPipelinesForProject**](PipelineApi.md#ListPipelinesForProject) | **Get** /project/{project-slug}/pipeline | Get all pipelines
[**ListWorkflowsByPipelineId**](PipelineApi.md#ListWorkflowsByPipelineId) | **Get** /pipeline/{pipeline-id}/workflow | Get a pipeline&#x27;s workflows
[**TriggerPipeline**](PipelineApi.md#TriggerPipeline) | **Post** /project/{project-slug}/pipeline | Trigger a new pipeline
[**TriggerPipelineRun**](PipelineApi.md#TriggerPipelineRun) | **Post** /project/{provider}/{organization}/{project}/pipeline/run | [Recommended] Trigger a new pipeline

# **ContinuePipeline**
> MessageResponse ContinuePipeline(ctx, optional)
Continue a pipeline

Continue a pipeline from the setup phase. For information on using pipeline parameters with dynamic configuration, see the [Pipeline values and parameters](https://circleci.com/docs/pipeline-variables/#pipeline-parameters-and-dynamic-configuration) docs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PipelineApiContinuePipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiContinuePipelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PipelineContinueBody**](PipelineContinueBody.md)|  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineById**
> Pipeline GetPipelineById(ctx, pipelineId)
Get a pipeline by ID

Returns a pipeline by the pipeline ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pipelineId** | [**string**](.md)| The unique ID of the pipeline. | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineByNumber**
> Pipeline GetPipelineByNumber(ctx, projectSlug, pipelineNumber)
Get a pipeline by pipeline number

Returns a pipeline by the pipeline number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
  **pipelineNumber** | [**Object**](.md)| The number of the pipeline. | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineConfigById**
> PipelineConfig GetPipelineConfigById(ctx, pipelineId)
Get a pipeline's configuration

Returns a pipeline's configuration by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pipelineId** | [**string**](.md)| The unique ID of the pipeline. | 

### Return type

[**PipelineConfig**](PipelineConfig.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineValuesById**
> map[string]Object GetPipelineValuesById(ctx, pipelineId)
Get pipeline values for a pipeline

Returns a map of pipeline values by pipeline ID. For more information see the [pipeline values reference page](https://circleci.com/docs/variables/#pipeline-values).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pipelineId** | [**string**](.md)| The unique ID of the pipeline. | 

### Return type

**map[string]Object**

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMyPipelines**
> PipelineListResponse ListMyPipelines(ctx, projectSlug, optional)
Get your pipelines

Returns a sequence of all pipelines for this project triggered by the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
 **optional** | ***PipelineApiListMyPipelinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiListMyPipelinesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 

### Return type

[**PipelineListResponse**](PipelineListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPipelines**
> PipelineListResponse ListPipelines(ctx, optional)
Get a list of pipelines

Returns all pipelines for the most recently built projects (max 250) you follow in an organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PipelineApiListPipelinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiListPipelinesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgSlug** | **optional.String**| Org slug in the form &#x60;vcs-slug/org-name&#x60;. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60; and replace the &#x60;org-name&#x60; with the organization ID (found in Organization Settings). | 
 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 
 **mine** | **optional.Bool**| Only include entries created by your user. | 

### Return type

[**PipelineListResponse**](PipelineListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPipelinesForProject**
> PipelineListResponse ListPipelinesForProject(ctx, projectSlug, optional)
Get all pipelines

Returns all pipelines for this project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
 **optional** | ***PipelineApiListPipelinesForProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiListPipelinesForProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branch** | **optional.String**| The name of a vcs branch. | 
 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 

### Return type

[**PipelineListResponse**](PipelineListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWorkflowsByPipelineId**
> WorkflowListResponse ListWorkflowsByPipelineId(ctx, pipelineId, optional)
Get a pipeline's workflows

Returns a paginated list of workflows by pipeline ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pipelineId** | [**string**](.md)| The unique ID of the pipeline. | 
 **optional** | ***PipelineApiListWorkflowsByPipelineIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiListWorkflowsByPipelineIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 

### Return type

[**WorkflowListResponse**](WorkflowListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerPipeline**
> PipelineCreation TriggerPipeline(ctx, projectSlug, optional)
Trigger a new pipeline

Not available to projects that use GitLab or GitHub App. Triggers a new pipeline on the project. **GitHub App users should use the [new Trigger Pipeline API](#tag/Pipeline/operation/triggerPipelineRun)**.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
 **optional** | ***PipelineApiTriggerPipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiTriggerPipelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TriggerPipelineParameters**](TriggerPipelineParameters.md)|  | 

### Return type

[**PipelineCreation**](PipelineCreation.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerPipelineRun**
> Pipeline TriggerPipelineRun(ctx, provider, organization, project, optional)
[Recommended] Trigger a new pipeline

Trigger a pipeline given a pipeline definition ID. Supports all integrations except GitLab.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **provider** | **string**| The &#x60;provider&#x60; segment of a project or org slug, the first of the three. This may be a VCS. For projects that use GitHub App, use &#x60;circleci&#x60;. | 
  **organization** | **string**| The &#x60;organization&#x60; segment of a project or org slug, the second of the three. For GitHub OAuth or Bitbucket projects, this is the organization name. For projects that use GitLab or GitHub App, use the organization ID (found in Organization Settings). | 
  **project** | **string**| The &#x60;project&#x60; segment of a project slug, the third of the three. For GitHub OAuth or Bitbucket projects, this is the repository name. For projects that use GitLab or GitHub App, use the project ID (found in Project Settings). | 
 **optional** | ***PipelineApiTriggerPipelineRunOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiTriggerPipelineRunOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of PipelineRequest**](PipelineRequest.md)|  | 

### Return type

[**Pipeline**](pipeline.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

