# {{classname}}

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgClaims**](OIDCTokenManagementApi.md#DeleteOrgClaims) | **Delete** /org/{orgID}/oidc-custom-claims | Delete org-level claims
[**DeleteProjectClaims**](OIDCTokenManagementApi.md#DeleteProjectClaims) | **Delete** /org/{orgID}/project/{projectID}/oidc-custom-claims | Delete project-level claims
[**GetOrgClaims**](OIDCTokenManagementApi.md#GetOrgClaims) | **Get** /org/{orgID}/oidc-custom-claims | Get org-level claims
[**GetProjectClaims**](OIDCTokenManagementApi.md#GetProjectClaims) | **Get** /org/{orgID}/project/{projectID}/oidc-custom-claims | Get project-level claims
[**PatchOrgClaims**](OIDCTokenManagementApi.md#PatchOrgClaims) | **Patch** /org/{orgID}/oidc-custom-claims | Patch org-level claims
[**PatchProjectClaims**](OIDCTokenManagementApi.md#PatchProjectClaims) | **Patch** /org/{orgID}/project/{projectID}/oidc-custom-claims | Patch project-level claims

# **DeleteOrgClaims**
> ClaimResponse DeleteOrgClaims(ctx, orgID, claims)
Delete org-level claims

Deletes org-level custom claims of OIDC identity tokens

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgID** | [**string**](.md)|  | 
  **claims** | **string**| comma separated list of claims to delete. Valid values are \&quot;audience\&quot; and \&quot;ttl\&quot;. | 

### Return type

[**ClaimResponse**](ClaimResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProjectClaims**
> ClaimResponse DeleteProjectClaims(ctx, orgID, projectID, claims)
Delete project-level claims

Deletes project-level custom claims of OIDC identity tokens

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgID** | [**string**](.md)|  | 
  **projectID** | [**string**](.md)|  | 
  **claims** | **string**| comma separated list of claims to delete. Valid values are \&quot;audience\&quot; and \&quot;ttl\&quot;. | 

### Return type

[**ClaimResponse**](ClaimResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgClaims**
> ClaimResponse GetOrgClaims(ctx, orgID)
Get org-level claims

Fetches org-level custom claims of OIDC identity tokens

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgID** | [**string**](.md)|  | 

### Return type

[**ClaimResponse**](ClaimResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectClaims**
> ClaimResponse GetProjectClaims(ctx, orgID, projectID)
Get project-level claims

Fetches project-level custom claims of OIDC identity tokens

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgID** | [**string**](.md)|  | 
  **projectID** | [**string**](.md)|  | 

### Return type

[**ClaimResponse**](ClaimResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchOrgClaims**
> ClaimResponse PatchOrgClaims(ctx, orgID, optional)
Patch org-level claims

Creates/Updates org-level custom claims of OIDC identity tokens

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgID** | [**string**](.md)|  | 
 **optional** | ***OIDCTokenManagementApiPatchOrgClaimsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OIDCTokenManagementApiPatchOrgClaimsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PatchClaimsRequest**](PatchClaimsRequest.md)|  | 

### Return type

[**ClaimResponse**](ClaimResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchProjectClaims**
> ClaimResponse PatchProjectClaims(ctx, orgID, projectID, optional)
Patch project-level claims

Creates/Updates project-level custom claims of OIDC identity tokens

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgID** | [**string**](.md)|  | 
  **projectID** | [**string**](.md)|  | 
 **optional** | ***OIDCTokenManagementApiPatchProjectClaimsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OIDCTokenManagementApiPatchProjectClaimsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PatchClaimsRequest**](PatchClaimsRequest.md)|  | 

### Return type

[**ClaimResponse**](ClaimResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

