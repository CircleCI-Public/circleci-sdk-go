# InlineResponse2009Items

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the workflow. | [default to null]
**Branch** | **string** | The VCS branch of a Workflow&#x27;s trigger. | [default to null]
**Duration** | **int64** | The duration in seconds of a run. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time the workflow was created. | [default to null]
**StoppedAt** | [**time.Time**](time.Time.md) | The date and time the workflow stopped. | [default to null]
**CreditsUsed** | **int64** | The number of credits used during execution. Note that Insights is not a real time financial reporting tool and should not be used for credit reporting. | [default to null]
**Status** | **string** | Workflow status. | [default to null]
**IsApproval** | **bool** | Describes if the job is an approval job or not. Approval jobs are intermediary jobs that are created to pause the workflow until approved. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

