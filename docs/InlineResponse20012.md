# InlineResponse20012

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageTestCount** | **int64** | The average number of tests executed per run | [default to null]
**MostFailedTests** | [**[]InlineResponse20012MostFailedTests**](inline_response_200_12_most_failed_tests.md) | Metrics for the most frequently failing tests | [default to null]
**MostFailedTestsExtra** | **int64** | The number of tests with the same success rate being omitted from most_failed_tests | [default to null]
**SlowestTests** | [**[]InlineResponse20012MostFailedTests**](inline_response_200_12_most_failed_tests.md) | Metrics for the slowest running tests | [default to null]
**SlowestTestsExtra** | **int64** | The number of tests with the same duration rate being omitted from slowest_tests | [default to null]
**TotalTestRuns** | **int64** | The total number of test runs | [default to null]
**TestRuns** | [**[]InlineResponse20012TestRuns**](inline_response_200_12_test_runs.md) | Test counts grouped by pipeline number and workflow id | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

