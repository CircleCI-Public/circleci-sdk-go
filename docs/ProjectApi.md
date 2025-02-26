# {{classname}}

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCheckoutKey**](ProjectApi.md#CreateCheckoutKey) | **Post** /project/{project-slug}/checkout-key | Create a new checkout key
[**CreateEnvVar**](ProjectApi.md#CreateEnvVar) | **Post** /project/{project-slug}/envvar | Create an environment variable
[**CreateProject**](ProjectApi.md#CreateProject) | **Post** /project/{provider}/{organization}/{project} | 🧪 Create a project
[**DeleteCheckoutKey**](ProjectApi.md#DeleteCheckoutKey) | **Delete** /project/{project-slug}/checkout-key/{fingerprint} | Delete a checkout key
[**DeleteEnvVar**](ProjectApi.md#DeleteEnvVar) | **Delete** /project/{project-slug}/envvar/{name} | Delete an environment variable
[**GetCheckoutKey**](ProjectApi.md#GetCheckoutKey) | **Get** /project/{project-slug}/checkout-key/{fingerprint} | Get a checkout key
[**GetEnvVar**](ProjectApi.md#GetEnvVar) | **Get** /project/{project-slug}/envvar/{name} | Get a masked environment variable
[**GetProjectBySlug**](ProjectApi.md#GetProjectBySlug) | **Get** /project/{project-slug} | Get a project
[**GetProjectSettings**](ProjectApi.md#GetProjectSettings) | **Get** /project/{provider}/{organization}/{project}/settings | 🧪 Get project settings
[**ListCheckoutKeys**](ProjectApi.md#ListCheckoutKeys) | **Get** /project/{project-slug}/checkout-key | Get all checkout keys
[**ListEnvVars**](ProjectApi.md#ListEnvVars) | **Get** /project/{project-slug}/envvar | List all environment variables
[**PatchProjectSettings**](ProjectApi.md#PatchProjectSettings) | **Patch** /project/{provider}/{organization}/{project}/settings | 🧪 Update project settings

# **CreateCheckoutKey**
> CheckoutKey CreateCheckoutKey(ctx, projectSlug, optional)
Create a new checkout key

Not available to projects that use GitLab or GitHub App. Creates a new checkout key. This API request is only usable with a user API token.                            Please ensure that you have authorized your account with GitHub before creating user keys.                            This is necessary to give CircleCI the permission to create a user key associated with                            your GitHub user account. You can find this page by visiting Project Settings > Checkout SSH Keys

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
 **optional** | ***ProjectApiCreateCheckoutKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiCreateCheckoutKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CheckoutKeyInput**](CheckoutKeyInput.md)|  | 

### Return type

[**CheckoutKey**](CheckoutKey.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEnvVar**
> EnvironmentVariable1 CreateEnvVar(ctx, projectSlug, optional)
Create an environment variable

Creates a new environment variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
 **optional** | ***ProjectApiCreateEnvVarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiCreateEnvVarOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EnvironmentVariable**](EnvironmentVariable.md)|  | 

### Return type

[**EnvironmentVariable1**](EnvironmentVariable_1.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProject**
> ProjectSettings CreateProject(ctx, provider, organization, project)
🧪 Create a project

[__EXPERIMENTAL__]  Creates a new CircleCI project, and returns a list of the default advanced settings. Can only be called on a repo with a main branch and an existing config.yml file. Not yet available to projects that use GitLab or GitHub App.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **provider** | **string**| The &#x60;provider&#x60; segment of a project or org slug, the first of the three. This may be a VCS. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60;. | 
  **organization** | **string**| The &#x60;organization&#x60; segment of a project or org slug, the second of the three. For GitHub OAuth or Bitbucket projects, this is the organization name. For projects that use GitLab or GitHub App, use the organization ID (found in Organization Settings). | 
  **project** | **string**| The &#x60;project&#x60; segment of a project slug, the third of the three. For GitHub OAuth or Bitbucket projects, this is the repository name. For projects that use GitLab or GitHub App, use the project ID (found in Project Settings). | 

### Return type

[**ProjectSettings**](project_settings.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCheckoutKey**
> MessageResponse DeleteCheckoutKey(ctx, projectSlug, fingerprint)
Delete a checkout key

Deletes the checkout key via md5 or sha256 fingerprint. sha256 keys should be url-encoded.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
  **fingerprint** | **string**| An SSH key fingerprint. | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnvVar**
> MessageResponse DeleteEnvVar(ctx, projectSlug, name)
Delete an environment variable

Deletes the environment variable named :name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
  **name** | **string**| The name of the environment variable. | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCheckoutKey**
> CheckoutKey GetCheckoutKey(ctx, projectSlug, fingerprint)
Get a checkout key

Returns an individual checkout key via md5 or sha256 fingerprint. sha256 keys should be url-encoded.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
  **fingerprint** | **string**| An SSH key fingerprint. | 

### Return type

[**CheckoutKey**](CheckoutKey.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvVar**
> EnvironmentVariable1 GetEnvVar(ctx, projectSlug, name)
Get a masked environment variable

Returns the masked value of environment variable :name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
  **name** | **string**| The name of the environment variable. | 

### Return type

[**EnvironmentVariable1**](EnvironmentVariable_1.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectBySlug**
> Project GetProjectBySlug(ctx, projectSlug)
Get a project

Retrieves a project by project slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 

### Return type

[**Project**](Project.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectSettings**
> ProjectSettings GetProjectSettings(ctx, provider, organization, project)
🧪 Get project settings

[__EXPERIMENTAL__] Returns a list of the advanced settings for a CircleCI project, whether enabled (true) or not (false).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **provider** | **string**| The &#x60;provider&#x60; segment of a project or org slug, the first of the three. This may be a VCS. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60;. | 
  **organization** | **string**| The &#x60;organization&#x60; segment of a project or org slug, the second of the three. For GitHub OAuth or Bitbucket projects, this is the organization name. For projects that use GitLab or GitHub App, use the organization ID (found in Organization Settings). | 
  **project** | **string**| The &#x60;project&#x60; segment of a project slug, the third of the three. For GitHub OAuth or Bitbucket projects, this is the repository name. For projects that use GitLab or GitHub App, use the project ID (found in Project Settings). | 

### Return type

[**ProjectSettings**](project_settings.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCheckoutKeys**
> CheckoutKeyListResponse ListCheckoutKeys(ctx, projectSlug, optional)
Get all checkout keys

Returns a sequence of checkout keys for `:project`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 
 **optional** | ***ProjectApiListCheckoutKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiListCheckoutKeysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **digest** | **optional.String**| The fingerprint digest type to return. This may be either &#x60;md5&#x60; or &#x60;sha256&#x60;. If not passed, defaults to &#x60;md5&#x60;. | 

### Return type

[**CheckoutKeyListResponse**](CheckoutKeyListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEnvVars**
> EnvironmentVariableListResponse ListEnvVars(ctx, projectSlug)
List all environment variables

Returns four 'x' characters, in addition to the last four ASCII characters of the value, consistent with the display of environment variable values on the CircleCI website.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60; as the &#x60;vcs-slug&#x60;, replace &#x60;org-name&#x60; with the organization ID (found in Organization Settings), and replace &#x60;repo-name&#x60; with the project ID (found in Project Settings). | 

### Return type

[**EnvironmentVariableListResponse**](EnvironmentVariableListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchProjectSettings**
> ProjectSettings PatchProjectSettings(ctx, body, provider, organization, project)
🧪 Update project settings

[__EXPERIMENTAL__] Updates one or more of the advanced settings for a CircleCI project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectSettings**](ProjectSettings.md)| The setting(s) to update, including one or more fields in the JSON object. Note that &#x60;oss: true&#x60; will only be set on projects whose underlying repositories are actually open source. | 
  **provider** | **string**| The &#x60;provider&#x60; segment of a project or org slug, the first of the three. This may be a VCS. For projects that use GitLab or GitHub App, use &#x60;circleci&#x60;. | 
  **organization** | **string**| The &#x60;organization&#x60; segment of a project or org slug, the second of the three. For GitHub OAuth or Bitbucket projects, this is the organization name. For projects that use GitLab or GitHub App, use the organization ID (found in Organization Settings). | 
  **project** | **string**| The &#x60;project&#x60; segment of a project slug, the third of the three. For GitHub OAuth or Bitbucket projects, this is the repository name. For projects that use GitLab or GitHub App, use the project ID (found in Project Settings). | 

### Return type

[**ProjectSettings**](project_settings.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

