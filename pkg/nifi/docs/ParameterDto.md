# ParameterDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Parameter | [optional] [default to null]
**Description** | **string** | The description of the Parameter | [optional] [default to null]
**Sensitive** | **bool** | Whether or not the Parameter is sensitive | [optional] [default to null]
**Value** | **string** | The value of the Parameter | [required] [default to null]
**ValueRemoved** | **bool** | Whether or not the value of the Parameter was removed. When a request is made to change a parameter, the value may be null. The absence of the value may be used either to indicate that the value is not to be changed, or that the value is to be set to null (i.e., removed). This denotes which of the two scenarios is being encountered.  | [optional] [default to null]
**Provided** | **bool** | Whether or not the Parameter is provided by a ParameterProvider | [optional] [default to null]
**ReferencedAssets** | [**[]AssetReferenceDto**](AssetReferenceDTO.md) | A list of identifiers of the assets that are referenced by the parameter | [optional] [default to null]
**ReferencingComponents** | [**[]AffectedComponentEntity**](AffectedComponentEntity.md) | The set of all components in the flow that are referencing this Parameter | [optional] [default to null]
**ParameterContext** | [***ParameterContextReferenceEntity**](ParameterContextReferenceEntity.md) |  | [optional] [default to null]
**Inherited** | **bool** | Whether or not the Parameter is inherited from another context | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

