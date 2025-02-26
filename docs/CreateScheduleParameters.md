# CreateScheduleParameters

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the schedule. | [default to null]
**Timetable** | [***AnyOfCreateScheduleParametersTimetable**](AnyOfCreateScheduleParametersTimetable.md) | Timetable that specifies when a schedule triggers. | [default to null]
**AttributionActor** | **string** | The attribution-actor of the scheduled pipeline. | [default to null]
**Parameters** | [**map[string]Object**](.md) | Pipeline parameters represented as key-value pairs. Must contain branch or tag. | [default to null]
**Description** | **string** | Description of the schedule. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

