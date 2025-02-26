# Schedule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the schedule. | [default to null]
**Timetable** | [***AnyOfScheduleTimetable**](AnyOfScheduleTimetable.md) | Timetable that specifies when a schedule triggers. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date and time the pipeline was last updated. | [default to null]
**Name** | **string** | Name of the schedule. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time the pipeline was created. | [default to null]
**ProjectSlug** | **string** | The project-slug for the schedule | [default to null]
**Parameters** | [**map[string]Object**](.md) | Pipeline parameters represented as key-value pairs. Must contain branch or tag. | [default to null]
**Actor** | [***User1**](User_1.md) |  | [default to null]
**Description** | **string** | Description of the schedule. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

