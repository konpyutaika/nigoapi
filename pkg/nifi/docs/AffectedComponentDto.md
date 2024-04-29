# AffectedComponentDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessGroupId** | **string** | The UUID of the Process Group that this component is in | [optional] [default to null]
**Id** | **string** | The UUID of this component | [optional] [default to null]
**ReferenceType** | **string** | The type of this component | [optional] [default to null]
**Name** | **string** | The name of this component. | [optional] [default to null]
**State** | **string** | The scheduled state of a processor or reporting task referencing a controller service. If this component is another controller service, this field represents the controller service state. | [optional] [default to null]
**ActiveThreadCount** | **int32** | The number of active threads for the referencing component. | [optional] [default to null]
**ValidationErrors** | **[]string** | The validation errors for the component. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

