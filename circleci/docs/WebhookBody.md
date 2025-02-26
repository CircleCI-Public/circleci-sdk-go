# WebhookBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the webhook | [default to null]
**Events** | **[]string** | Events that will trigger the webhook | [default to null]
**Url** | **string** | URL to deliver the webhook to. Note: protocol must be included as well (only https is supported) | [default to null]
**VerifyTls** | **bool** | Whether to enforce TLS certificate verification when delivering the webhook | [default to null]
**SigningSecret** | **string** | Secret used to build an HMAC hash of the payload and passed as a header in the webhook request | [default to null]
**Scope** | [***WebhookScope**](webhook_scope.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

