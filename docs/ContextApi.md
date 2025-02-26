# {{classname}}

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEnvironmentVariableToContext**](ContextApi.md#AddEnvironmentVariableToContext) | **Put** /context/{context-id}/environment-variable/{env-var-name} | Add or update an environment variable
[**CreateContext**](ContextApi.md#CreateContext) | **Post** /context | Create a new context
[**CreateContextRestriction**](ContextApi.md#CreateContextRestriction) | **Post** /context/{context_id}/restrictions | 🧪 Create context restriction
[**DeleteContext**](ContextApi.md#DeleteContext) | **Delete** /context/{context-id} | Delete a context
[**DeleteContextRestriction**](ContextApi.md#DeleteContextRestriction) | **Delete** /context/{context_id}/restrictions/{restriction_id} | 🧪 Delete context restriction
[**DeleteEnvironmentVariableFromContext**](ContextApi.md#DeleteEnvironmentVariableFromContext) | **Delete** /context/{context-id}/environment-variable/{env-var-name} | Remove an environment variable
[**GetContext**](ContextApi.md#GetContext) | **Get** /context/{context-id} | Get a context
[**GetContextRestrictions**](ContextApi.md#GetContextRestrictions) | **Get** /context/{context_id}/restrictions | 🧪 Get context restrictions
[**ListContexts**](ContextApi.md#ListContexts) | **Get** /context | List contexts
[**ListEnvironmentVariablesFromContext**](ContextApi.md#ListEnvironmentVariablesFromContext) | **Get** /context/{context-id}/environment-variable | List environment variables

# **AddEnvironmentVariableToContext**
> InlineResponse2002 AddEnvironmentVariableToContext(ctx, contextId, envVarName, optional)
Add or update an environment variable

Create or update an environment variable within a context. Returns information about the environment variable, not including its value.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | [**string**](.md)| ID of the context (UUID) | 
  **envVarName** | **string**| The name of the environment variable | 
 **optional** | ***ContextApiAddEnvironmentVariableToContextOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContextApiAddEnvironmentVariableToContextOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EnvironmentvariableEnvvarnameBody**](EnvironmentvariableEnvvarnameBody.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateContext**
> Context CreateContext(ctx, optional)
Create a new context

Creates a new context.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContextApiCreateContextOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContextApiCreateContextOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ContextBody**](ContextBody.md)|  | 

### Return type

[**Context**](Context.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateContextRestriction**
> RestrictionCreated CreateContextRestriction(ctx, body, contextId)
🧪 Create context restriction

[__EXPERIMENTAL__] Creates project restriction on a context.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ContextIdRestrictionsBody**](ContextIdRestrictionsBody.md)|  | 
  **contextId** | **string**| An opaque identifier of a context. | 

### Return type

[**RestrictionCreated**](restriction_created.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContext**
> MessageResponse DeleteContext(ctx, contextId)
Delete a context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | [**string**](.md)| ID of the context (UUID) | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContextRestriction**
> RestrictionDeleted DeleteContextRestriction(ctx, contextId, restrictionId)
🧪 Delete context restriction

[__EXPERIMENTAL__] Deletes a project restriction on a context.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**| An opaque identifier of a context. | 
  **restrictionId** | **string**| An opaque identifier of a context restriction. | 

### Return type

[**RestrictionDeleted**](restriction_deleted.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnvironmentVariableFromContext**
> MessageResponse DeleteEnvironmentVariableFromContext(ctx, envVarName, contextId)
Remove an environment variable

Delete an environment variable from a context.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **envVarName** | **string**| The name of the environment variable | 
  **contextId** | [**string**](.md)| ID of the context (UUID) | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContext**
> Context GetContext(ctx, contextId)
Get a context

Returns basic information about a context.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | [**string**](.md)| ID of the context (UUID) | 

### Return type

[**Context**](Context.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContextRestrictions**
> ContextProjectRestrictionsList GetContextRestrictions(ctx, contextId)
🧪 Get context restrictions

[__EXPERIMENTAL__] Gets a list of project restrictions associated with a context.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | **string**| An opaque identifier of a context. | 

### Return type

[**ContextProjectRestrictionsList**](context_project_restrictions_list.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListContexts**
> InlineResponse200 ListContexts(ctx, optional)
List contexts

List all contexts for an owner.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContextApiListContextsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContextApiListContextsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerId** | [**optional.Interface of string**](.md)| The unique ID of the owner of the context. Specify either this or owner-slug. | 
 **ownerSlug** | **optional.String**| A string that represents an organization. Specify either this or owner-id. Cannot be used for accounts. | 
 **ownerType** | **optional.String**| The type of the owner. Defaults to \&quot;organization\&quot;. Accounts are only used as context owners in server. | 
 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEnvironmentVariablesFromContext**
> InlineResponse2001 ListEnvironmentVariablesFromContext(ctx, contextId, optional)
List environment variables

List information about environment variables in a context, not including their values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextId** | [**string**](.md)| ID of the context (UUID) | 
 **optional** | ***ContextApiListEnvironmentVariablesFromContextOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContextApiListEnvironmentVariablesFromContextOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

