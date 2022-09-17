# PropertyDescriptor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the property key | [default to null]
**DisplayName** | **string** | The display name of the property key, if different from the name | [optional] [default to null]
**Description** | **string** | The description of what the property does | [optional] [default to null]
**AllowableValues** | [**[]PropertyAllowableValue**](PropertyAllowableValue.md) | A list of the allowable values for the property | [optional] [default to null]
**DefaultValue** | **string** | The default value if a user-set value is not specified | [optional] [default to null]
**Required** | **bool** | Whether or not  the property is required for the component | [optional] [default to null]
**Sensitive** | **bool** | Whether or not  the value of the property is considered sensitive (e.g., passwords and keys) | [optional] [default to null]
**ExpressionLanguageScope** | **string** | The scope of expression language supported by this property | [optional] [default to null]
**ExpressionLanguageScopeDescription** | **string** | The description of the expression language scope supported by this property | [optional] [default to null]
**TypeProvidedByValue** | [***DefinedType**](DefinedType.md) | Indicates that this property is for selecting a controller service of the specified type | [optional] [default to null]
**ValidRegex** | **string** | A regular expression that can be used to validate the value of this property | [optional] [default to null]
**Validator** | **string** | Name of the validator used for this property descriptor | [optional] [default to null]
**Dynamic** | **bool** | Whether or not the descriptor is for a dynamically added property | [optional] [default to null]
**ResourceDefinition** | [***PropertyResourceDefinition**](PropertyResourceDefinition.md) | Indicates that this property references external resources | [optional] [default to null]
**Dependencies** | [**[]PropertyDependency**](PropertyDependency.md) | The dependencies that this property has on other properties | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


