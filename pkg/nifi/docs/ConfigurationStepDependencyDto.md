# ConfigurationStepDependencyDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DependentValues** | **[]string** | The values of the dependent property that satisfy this dependency. If null, any non-null value satisfies the dependency. | [optional] [default to null]
**PropertyName** | **string** | The name of the property within the dependent step that must have a value. | [optional] [default to null]
**StepName** | **string** | The name of the configuration step that this step depends on. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

