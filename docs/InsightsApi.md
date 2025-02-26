# {{classname}}

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllInsightsBranches**](InsightsApi.md#GetAllInsightsBranches) | **Get** /insights/{project-slug}/branches | Get all branches for a project
[**GetFlakyTests**](InsightsApi.md#GetFlakyTests) | **Get** /insights/{project-slug}/flaky-tests | Get flaky tests for a project
[**GetJobTimeseries**](InsightsApi.md#GetJobTimeseries) | **Get** /insights/time-series/{project-slug}/workflows/{workflow-name}/jobs | Job timeseries data
[**GetOrgSummaryData**](InsightsApi.md#GetOrgSummaryData) | **Get** /insights/{org-slug}/summary | Get summary metrics with trends for the entire org, and for each project.
[**GetProjectWorkflowJobMetrics**](InsightsApi.md#GetProjectWorkflowJobMetrics) | **Get** /insights/{project-slug}/workflows/{workflow-name}/jobs | Get summary metrics for a project workflow&#x27;s jobs.
[**GetProjectWorkflowMetrics**](InsightsApi.md#GetProjectWorkflowMetrics) | **Get** /insights/{project-slug}/workflows | Get summary metrics for a project&#x27;s workflows
[**GetProjectWorkflowRuns**](InsightsApi.md#GetProjectWorkflowRuns) | **Get** /insights/{project-slug}/workflows/{workflow-name} | Get recent runs of a workflow
[**GetProjectWorkflowTestMetrics**](InsightsApi.md#GetProjectWorkflowTestMetrics) | **Get** /insights/{project-slug}/workflows/{workflow-name}/test-metrics | Get test metrics for a project&#x27;s workflows
[**GetProjectWorkflowsPageData**](InsightsApi.md#GetProjectWorkflowsPageData) | **Get** /insights/pages/{project-slug}/summary | Get summary metrics and trends for a project across it&#x27;s workflows and branches
[**GetWorkflowSummary**](InsightsApi.md#GetWorkflowSummary) | **Get** /insights/{project-slug}/workflows/{workflow-name}/summary | Get metrics and trends for workflows

# **GetAllInsightsBranches**
> InlineResponse2006 GetAllInsightsBranches(ctx, projectSlug, optional)
Get all branches for a project

Get a list of all branches for a specified project. The list will only contain branches currently available within Insights. The maximum number of branches returned by this endpoint is 5,000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
 **optional** | ***InsightsApiGetAllInsightsBranchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightsApiGetAllInsightsBranchesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowName** | **optional.String**| The name of a workflow. If not passed we will scope the API call to the project. | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlakyTests**
> InlineResponse2007 GetFlakyTests(ctx, projectSlug)
Get flaky tests for a project

Get a list of flaky tests for a given project. Flaky tests are branch agnostic.              A flaky test is a test that passed and failed in the same commit.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobTimeseries**
> InlineResponse2004 GetJobTimeseries(ctx, projectSlug, workflowName, optional)
Job timeseries data

Get timeseries data for all jobs within a workflow. Hourly granularity data is only retained for 48 hours while daily granularity data is retained for 90 days.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
  **workflowName** | **string**| The name of the workflow. | 
 **optional** | ***InsightsApiGetJobTimeseriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightsApiGetJobTimeseriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branch** | **optional.String**| The name of a vcs branch. If not passed we will scope the API call to the default branch. | 
 **granularity** | **optional.String**| The granularity for which to query timeseries data. | 
 **startDate** | **optional.Time**| Include only executions that started at or after this date. This must be specified if an end-date is provided. | 
 **endDate** | **optional.Time**| Include only executions that started before this date. This date can be at most 90 days after the start-date. | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgSummaryData**
> InlineResponse2005 GetOrgSummaryData(ctx, orgSlug, optional)
Get summary metrics with trends for the entire org, and for each project.

Gets aggregated summary metrics with trends for the entire org.               Also gets aggregated metrics and trends for each project belonging to the org.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgSlug** | **string**| Org slug in the form &#x60;vcs-slug/org-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
 **optional** | ***InsightsApiGetOrgSummaryDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightsApiGetOrgSummaryDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportingWindow** | **optional.String**| The time window used to calculate summary metrics. If not provided, defaults to last-90-days | 
 **projectNames** | [**optional.Interface of interface{}**](.md)| List of project names. | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectWorkflowJobMetrics**
> InlineResponse20010 GetProjectWorkflowJobMetrics(ctx, projectSlug, workflowName, optional)
Get summary metrics for a project workflow's jobs.

Get summary metrics for a project workflow's jobs. Job runs going back at most 90 days are included in the aggregation window. Metrics are refreshed daily, and thus may not include executions from the last 24 hours. Please note that Insights is not a financial reporting tool and should not be used for precise credit reporting.  Credit reporting from Insights does not use the same source of truth as the billing information that is found in the Plan Overview page in the CircleCI UI, nor does the underlying data have the same data accuracy guarantees as the billing information in the CircleCI UI.  This may lead to discrepancies between credits reported from Insights and the billing information in the Plan Overview page of the CircleCI UI.  For precise credit reporting, always use the Plan Overview page in the CircleCI UI.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
  **workflowName** | **string**| The name of the workflow. | 
 **optional** | ***InsightsApiGetProjectWorkflowJobMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightsApiGetProjectWorkflowJobMetricsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 
 **allBranches** | **optional.Bool**| Whether to retrieve data for all branches combined. Use either this parameter OR the branch name parameter. | 
 **branch** | **optional.String**| The name of a vcs branch. If not passed we will scope the API call to the default branch. | 
 **reportingWindow** | **optional.String**| The time window used to calculate summary metrics. If not provided, defaults to last-90-days | 
 **jobName** | **optional.String**| The name of the jobs you would like to filter from your workflow. If not specified, all workflow jobs will be returned. The job name can either be the full job name or just a substring of the job name. | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectWorkflowMetrics**
> InlineResponse2008 GetProjectWorkflowMetrics(ctx, projectSlug, optional)
Get summary metrics for a project's workflows

Get summary metrics for a project's workflows.  Workflow runs going back at most 90 days are included in the aggregation window. Metrics are refreshed daily, and thus may not include executions from the last 24 hours.  Please note that Insights is not a financial reporting tool and should not be used for precise credit reporting.  Credit reporting from Insights does not use the same source of truth as the billing information that is found in the Plan Overview page in the CircleCI UI, nor does the underlying data have the same data accuracy guarantees as the billing information in the CircleCI UI.  This may lead to discrepancies between credits reported from Insights and the billing information in the Plan Overview page of the CircleCI UI.  For precise credit reporting, always use the Plan Overview page in the CircleCI UI.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
 **optional** | ***InsightsApiGetProjectWorkflowMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightsApiGetProjectWorkflowMetricsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 
 **allBranches** | **optional.Bool**| Whether to retrieve data for all branches combined. Use either this parameter OR the branch name parameter. | 
 **branch** | **optional.String**| The name of a vcs branch. If not passed we will scope the API call to the default branch. | 
 **reportingWindow** | **optional.String**| The time window used to calculate summary metrics. If not provided, defaults to last-90-days | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectWorkflowRuns**
> InlineResponse2009 GetProjectWorkflowRuns(ctx, projectSlug, workflowName, optional)
Get recent runs of a workflow

Get recent runs of a workflow. Runs going back at most 90 days are returned. Please note that Insights is not a financial reporting tool and should not be used for precise credit reporting.  Credit reporting from Insights does not use the same source of truth as the billing information that is found in the Plan Overview page in the CircleCI UI, nor does the underlying data have the same data accuracy guarantees as the billing information in the CircleCI UI.  This may lead to discrepancies between credits reported from Insights and the billing information in the Plan Overview page of the CircleCI UI.  For precise credit reporting, always use the Plan Overview page in the CircleCI UI.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
  **workflowName** | **string**| The name of the workflow. | 
 **optional** | ***InsightsApiGetProjectWorkflowRunsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightsApiGetProjectWorkflowRunsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **allBranches** | **optional.Bool**| Whether to retrieve data for all branches combined. Use either this parameter OR the branch name parameter. | 
 **branch** | **optional.String**| The name of a vcs branch. If not passed we will scope the API call to the default branch. | 
 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 
 **startDate** | **optional.Time**| Include only executions that started at or after this date. This must be specified if an end-date is provided. | 
 **endDate** | **optional.Time**| Include only executions that started before this date. This date can be at most 90 days after the start-date. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectWorkflowTestMetrics**
> InlineResponse20012 GetProjectWorkflowTestMetrics(ctx, projectSlug, workflowName, optional)
Get test metrics for a project's workflows

Get test metrics for a project's workflows. Currently tests metrics are calculated based on 10 most recent workflow runs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
  **workflowName** | **string**| The name of the workflow. | 
 **optional** | ***InsightsApiGetProjectWorkflowTestMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightsApiGetProjectWorkflowTestMetricsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branch** | **optional.String**| The name of a vcs branch. If not passed we will scope the API call to the default branch. | 
 **allBranches** | **optional.Bool**| Whether to retrieve data for all branches combined. Use either this parameter OR the branch name parameter. | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectWorkflowsPageData**
> InlineResponse2003 GetProjectWorkflowsPageData(ctx, projectSlug, optional)
Get summary metrics and trends for a project across it's workflows and branches

Get summary metrics and trends for a project at workflow and branch level.              Workflow runs going back at most 90 days are included in the aggregation window.              Trends are only supported upto last 30 days.              Please note that Insights is not a financial reporting tool and should not be used for precise credit reporting.  Credit reporting from Insights does not use the same source of truth as the billing information that is found in the Plan Overview page in the CircleCI UI, nor does the underlying data have the same data accuracy guarantees as the billing information in the CircleCI UI.  This may lead to discrepancies between credits reported from Insights and the billing information in the Plan Overview page of the CircleCI UI.  For precise credit reporting, always use the Plan Overview page in the CircleCI UI.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
 **optional** | ***InsightsApiGetProjectWorkflowsPageDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightsApiGetProjectWorkflowsPageDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportingWindow** | **optional.String**| The time window used to calculate summary metrics. If not provided, defaults to last-90-days | 
 **branches** | [**optional.Interface of interface{}**](.md)| The names of VCS branches to include in branch-level workflow metrics. | 
 **workflowNames** | [**optional.Interface of interface{}**](.md)| The names of workflows to include in workflow-level metrics. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorkflowSummary**
> InlineResponse20011 GetWorkflowSummary(ctx, projectSlug, workflowName, optional)
Get metrics and trends for workflows

Get the metrics and trends for a particular workflow on a single branch or all branches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
  **workflowName** | **string**| The name of the workflow. | 
 **optional** | ***InsightsApiGetWorkflowSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightsApiGetWorkflowSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **allBranches** | **optional.Bool**| Whether to retrieve data for all branches combined. Use either this parameter OR the branch name parameter. | 
 **branch** | **optional.String**| The name of a vcs branch. If not passed we will scope the API call to the default branch. | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

