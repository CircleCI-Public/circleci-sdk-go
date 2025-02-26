# {{classname}}

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsageExport**](UsageApi.md#CreateUsageExport) | **Post** /organizations/{org_id}/usage_export_job | Create a usage export
[**GetUsageExport**](UsageApi.md#GetUsageExport) | **Get** /organizations/{org_id}/usage_export_job/{usage_export_job_id} | Get a usage export

# **CreateUsageExport**
> UsageExportJob CreateUsageExport(ctx, body, orgId)
Create a usage export

Submits a request to create a usage export for an organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrgIdUsageExportJobBody**](OrgIdUsageExportJobBody.md)|  | 
  **orgId** | **string**| An opaque identifier of an organization. | 

### Return type

[**UsageExportJob**](usage_export_job.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsageExport**
> GetUsageExportJobStatus GetUsageExport(ctx, orgId, usageExportJobId)
Get a usage export

Gets a usage export for an organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | **string**| An opaque identifier of an organization. | 
  **usageExportJobId** | [**string**](.md)| An opaque identifier of a usage export job. | 

### Return type

[**GetUsageExportJobStatus**](get_usage_export_job_status.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

