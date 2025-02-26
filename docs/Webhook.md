# Webhook

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | URL to deliver the webhook to. Note: protocol must be included as well (only https is supported) | [default to null]
**VerifyTls** | **bool** | Whether to enforce TLS certificate verification when delivering the webhook | [default to null]
**Id** | **string** | The unique ID of the webhook | [default to null]
**SigningSecret** | **string** | Masked value of the secret used to build an HMAC hash of the payload and passed as a header in the webhook request | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date and time the webhook was last updated. | [default to null]
**Name** | **string** | Name of the webhook | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time the webhook was created. | [default to null]
**Scope** | [***InlineResponse20014Scope**](inline_response_200_14_scope.md) |  | [default to null]
**Events** | **[]string** | Events that will trigger the webhook | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

