# {{classname}}

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSchedule**](ScheduleApi.md#CreateSchedule) | **Post** /project/{project-slug}/schedule | Create a schedule
[**DeleteScheduleById**](ScheduleApi.md#DeleteScheduleById) | **Delete** /schedule/{schedule-id} | Delete a schedule
[**GetScheduleById**](ScheduleApi.md#GetScheduleById) | **Get** /schedule/{schedule-id} | Get a schedule
[**ListSchedulesForProject**](ScheduleApi.md#ListSchedulesForProject) | **Get** /project/{project-slug}/schedule | Get all schedules
[**UpdateSchedule**](ScheduleApi.md#UpdateSchedule) | **Patch** /schedule/{schedule-id} | Update a schedule

# **CreateSchedule**
> Schedule CreateSchedule(ctx, projectSlug, optional)
Create a schedule

Not yet available to projects that use GitLab or GitHub App. Creates a schedule and returns the created schedule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
 **optional** | ***ScheduleApiCreateScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScheduleApiCreateScheduleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateScheduleParameters**](CreateScheduleParameters.md)|  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteScheduleById**
> MessageResponse DeleteScheduleById(ctx, scheduleId)
Delete a schedule

Not yet available to projects that use GitLab or GitHub App. Deletes the schedule by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scheduleId** | [**string**](.md)| The unique ID of the schedule. | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScheduleById**
> Schedule GetScheduleById(ctx, scheduleId)
Get a schedule

Get a schedule by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scheduleId** | [**string**](.md)| The unique ID of the schedule. | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSchedulesForProject**
> InlineResponse20013 ListSchedulesForProject(ctx, projectSlug, optional)
Get all schedules

Returns all schedules for this project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
 **optional** | ***ScheduleApiListSchedulesForProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScheduleApiListSchedulesForProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSchedule**
> Schedule UpdateSchedule(ctx, scheduleId, optional)
Update a schedule

Not yet available to projects that use GitLab or GitHub App. Updates a schedule and returns the updated schedule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scheduleId** | [**string**](.md)| The unique ID of the schedule. | 
 **optional** | ***ScheduleApiUpdateScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScheduleApiUpdateScheduleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateScheduleParameters**](UpdateScheduleParameters.md)|  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

