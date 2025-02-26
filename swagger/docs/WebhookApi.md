# {{classname}}

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhookApi.md#CreateWebhook) | **Post** /webhook | Create an outbound webhook
[**DeleteWebhook**](WebhookApi.md#DeleteWebhook) | **Delete** /webhook/{webhook-id} | Delete an outbound webhook
[**GetWebhookById**](WebhookApi.md#GetWebhookById) | **Get** /webhook/{webhook-id} | Get a webhook
[**GetWebhooks**](WebhookApi.md#GetWebhooks) | **Get** /webhook | List webhooks
[**UpdateWebhook**](WebhookApi.md#UpdateWebhook) | **Put** /webhook/{webhook-id} | Update an outbound webhook

# **CreateWebhook**
> Webhook CreateWebhook(ctx, optional)
Create an outbound webhook

Creates an outbound webhook.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhookApiCreateWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiCreateWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of WebhookBody**](WebhookBody.md)|  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWebhook**
> MessageResponse DeleteWebhook(ctx, webhookId)
Delete an outbound webhook

Deletes an outbound webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | [**string**](.md)| ID of the webhook (UUID) | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhookById**
> Webhook GetWebhookById(ctx, webhookId)
Get a webhook

Get an outbound webhook by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | [**string**](.md)| ID of the webhook (UUID) | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhooks**
> InlineResponse20014 GetWebhooks(ctx, scopeId, scopeType)
List webhooks

Get a list of outbound webhooks that match the given scope-type and scope-id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scopeId** | [**string**](.md)| ID of the scope being used (at the moment, only project ID is supported) | 
  **scopeType** | **string**| Type of the scope being used | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWebhook**
> Webhook UpdateWebhook(ctx, webhookId, optional)
Update an outbound webhook

Updates an outbound webhook.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | [**string**](.md)| ID of the webhook (UUID) | 
 **optional** | ***WebhookApiUpdateWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookApiUpdateWebhookOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of WebhookWebhookidBody**](WebhookWebhookidBody.md)|  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

