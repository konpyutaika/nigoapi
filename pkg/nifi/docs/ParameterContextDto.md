# ParameterContextDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoundProcessGroups** | [**[]ProcessGroupEntity**](ProcessGroupEntity.md) | The Process Groups that are bound to this Parameter Context | [required] [default to null]
**Description** | **string** | The Description of the Parameter Context. | [optional] [default to null]
**Id** | **string** | The ID the Parameter Context. | [optional] [default to null]
**InheritedParameterContexts** | [**[]ParameterContextReferenceEntity**](ParameterContextReferenceEntity.md) | A list of references of Parameter Contexts from which this one inherits parameters | [required] [default to null]
**Name** | **string** | The Name of the Parameter Context. | [optional] [default to null]
**ParameterProviderConfiguration** | [***ParameterProviderConfigurationEntity**](ParameterProviderConfigurationEntity.md) |  | [optional] [default to null]
**Parameters** | [**[]ParameterEntity**](ParameterEntity.md) | The Parameters for the Parameter Context | [required] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

