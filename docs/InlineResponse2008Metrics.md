# InlineResponse2008Metrics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalRuns** | **int64** | The total number of runs, including runs that are still on-hold or running. | [default to null]
**SuccessfulRuns** | **int64** | The number of successful runs. | [default to null]
**Mttr** | **int64** | The mean time to recovery (mean time between failures and their next success) in seconds. | [default to null]
**TotalCreditsUsed** | **int64** | The total credits consumed by the workflow in the aggregation window. Note that Insights is not a real time financial reporting tool and should not be used for credit reporting. | [default to null]
**FailedRuns** | **int64** | The number of failed runs. | [default to null]
**SuccessRate** | **float32** |  | [default to null]
**DurationMetrics** | [***InlineResponse2008MetricsDurationMetrics**](inline_response_200_8_metrics_duration_metrics.md) |  | [default to null]
**TotalRecoveries** | **int64** | The number of recovered workflow executions per day. | [default to null]
**Throughput** | **float32** | The average number of runs per day. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

