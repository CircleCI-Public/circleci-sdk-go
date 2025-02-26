# SchedulescheduleidTimetable

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerHour** | **int32** | Number of times a schedule triggers per hour, value must be between 1 and 60 | [optional] [default to null]
**HoursOfDay** | **[]int32** | Hours in a day in which the schedule triggers. | [optional] [default to null]
**DaysOfWeek** | **[]string** | Days in a week in which the schedule triggers. | [optional] [default to null]
**DaysOfMonth** | **[]int32** | Days in a month in which the schedule triggers. This is mutually exclusive with days in a week. | [optional] [default to null]
**Months** | **[]string** | Months in which the schedule triggers. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

