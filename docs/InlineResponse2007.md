# InlineResponse2007

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlakyTests** | [**[]InlineResponse2007Flakytests**](inline_response_200_7_flakytests.md) | A list of all instances of flakes. Note that a test is no longer considered flaky after 2 weeks have passed without a flake. Each flake resets this timer. | [default to null]
**TotalFlakyTests** | **float64** | A count of unique tests that have failed. If your project has N tests that have flaked multiple times each, this will be equal to N. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

