# OrgIdUsageExportJobBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | [**time.Time**](time.Time.md) | The start date &amp; time (inclusive) of the range from which data will be pulled. Must be no more than one year ago. | [default to null]
**End** | [**time.Time**](time.Time.md) | The end date &amp; time (inclusive) of the range from which data will be pulled. Must be no more than 31 days after &#x60;start&#x60;. | [default to null]
**SharedOrgIds** | **[]string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

