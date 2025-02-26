# {{classname}}

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyBundle**](PolicyManagementApi.md#CreatePolicyBundle) | **Post** /owner/{ownerID}/context/{context}/policy-bundle | Creates policy bundle for the context
[**GetDecisionLog**](PolicyManagementApi.md#GetDecisionLog) | **Get** /owner/{ownerID}/context/{context}/decision/{decisionID} | Retrieves the owner&#x27;s decision audit log by given decisionID
[**GetDecisionLogPolicyBundle**](PolicyManagementApi.md#GetDecisionLogPolicyBundle) | **Get** /owner/{ownerID}/context/{context}/decision/{decisionID}/policy-bundle | Retrieves Policy Bundle for a given decision log ID
[**GetDecisionLogs**](PolicyManagementApi.md#GetDecisionLogs) | **Get** /owner/{ownerID}/context/{context}/decision | Retrieves the owner&#x27;s decision audit logs.
[**GetDecisionSettings**](PolicyManagementApi.md#GetDecisionSettings) | **Get** /owner/{ownerID}/context/{context}/decision/settings | Get the decision settings
[**GetPolicyBundle**](PolicyManagementApi.md#GetPolicyBundle) | **Get** /owner/{ownerID}/context/{context}/policy-bundle | Retrieves Policy Bundle
[**GetPolicyDocument**](PolicyManagementApi.md#GetPolicyDocument) | **Get** /owner/{ownerID}/context/{context}/policy-bundle/{policyName} | Retrieves a policy document
[**MakeDecision**](PolicyManagementApi.md#MakeDecision) | **Post** /owner/{ownerID}/context/{context}/decision | Makes a decision
[**SetDecisionSettings**](PolicyManagementApi.md#SetDecisionSettings) | **Patch** /owner/{ownerID}/context/{context}/decision/settings | Set the decision settings

# **CreatePolicyBundle**
> BundleDiff CreatePolicyBundle(ctx, ownerID, context, optional)
Creates policy bundle for the context

This endpoint replaces the current policy bundle with the provided policy bundle

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerID** | **string**|  | 
  **context** | **string**|  | 
 **optional** | ***PolicyManagementApiCreatePolicyBundleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyManagementApiCreatePolicyBundleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of BundlePayload**](BundlePayload.md)|  | 
 **dry** | **optional.**|  | 

### Return type

[**BundleDiff**](BundleDiff.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDecisionLog**
> DecisionLog GetDecisionLog(ctx, ownerID, context, decisionID)
Retrieves the owner's decision audit log by given decisionID

This endpoint will retrieve a decision for a given decision log ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerID** | **string**|  | 
  **context** | **string**|  | 
  **decisionID** | **string**|  | 

### Return type

[**DecisionLog**](DecisionLog.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDecisionLogPolicyBundle**
> map[string][]Policy GetDecisionLogPolicyBundle(ctx, ownerID, context, decisionID)
Retrieves Policy Bundle for a given decision log ID

This endpoint will retrieve a policy bundle for a given decision log ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerID** | **string**|  | 
  **context** | **string**|  | 
  **decisionID** | **string**|  | 

### Return type

[**map[string][]Policy**](map.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDecisionLogs**
> []DecisionLog GetDecisionLogs(ctx, ownerID, context, optional)
Retrieves the owner's decision audit logs.

This endpoint will return a list of decision audit logs that were made using this owner's policies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerID** | **string**|  | 
  **context** | **string**|  | 
 **optional** | ***PolicyManagementApiGetDecisionLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyManagementApiGetDecisionLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **optional.String**| Return decisions matching this decision status. | 
 **after** | **optional.Time**| Return decisions made after this date. | 
 **before** | **optional.Time**| Return decisions made before this date. | 
 **branch** | **optional.String**| Return decisions made on this branch. | 
 **projectId** | **optional.String**| Return decisions made for this project. | 
 **buildNumber** | **optional.String**| Return decisions made for this build number. | 
 **offset** | **optional.Int32**| Sets the offset when retrieving the decisions, for paging. | 

### Return type

[**[]DecisionLog**](DecisionLog.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDecisionSettings**
> DecisionSettings GetDecisionSettings(ctx, ownerID, context)
Get the decision settings

This endpoint retrieves the current decision settings (eg enable/disable policy evaluation)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerID** | **string**|  | 
  **context** | **string**|  | 

### Return type

[**DecisionSettings**](DecisionSettings.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyBundle**
> map[string][]Policy GetPolicyBundle(ctx, ownerID, context)
Retrieves Policy Bundle

This endpoint will retrieve a policy bundle

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerID** | **string**|  | 
  **context** | **string**|  | 

### Return type

[**map[string][]Policy**](map.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyDocument**
> Policy GetPolicyDocument(ctx, ownerID, context, policyName)
Retrieves a policy document

This endpoint will retrieve a policy document.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerID** | **string**|  | 
  **context** | **string**|  | 
  **policyName** | **string**| the policy name set by the rego policy_name rule | 

### Return type

[**Policy**](Policy.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MakeDecision**
> Decision MakeDecision(ctx, ownerID, context, optional)
Makes a decision

This endpoint will evaluate input data (config+metadata) against owner's stored policies and return a decision.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerID** | **string**|  | 
  **context** | **string**|  | 
 **optional** | ***PolicyManagementApiMakeDecisionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyManagementApiMakeDecisionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ContextDecisionBody**](ContextDecisionBody.md)|  | 

### Return type

[**Decision**](Decision.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetDecisionSettings**
> DecisionSettings SetDecisionSettings(ctx, ownerID, context, optional)
Set the decision settings

This endpoint allows modifying decision settings (eg enable/disable policy evaluation)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ownerID** | **string**|  | 
  **context** | **string**|  | 
 **optional** | ***PolicyManagementApiSetDecisionSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyManagementApiSetDecisionSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DecisionSettings**](DecisionSettings.md)|  | 

### Return type

[**DecisionSettings**](DecisionSettings.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

