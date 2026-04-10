# ConnectorValueReferenceDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetReferences** | [**[]AssetReferenceDto**](AssetReferenceDTO.md) | The asset references. Applicable when valueType is ASSET_REFERENCE. | [optional] [default to null]
**FullyQualifiedSecretName** | **string** | The fully qualified secret name. Applicable when valueType is SECRET_REFERENCE. | [optional] [default to null]
**SecretName** | **string** | The secret name. Applicable when valueType is SECRET_REFERENCE. | [optional] [default to null]
**SecretProviderId** | **string** | The secret provider identifier. Applicable when valueType is SECRET_REFERENCE. | [optional] [default to null]
**SecretProviderName** | **string** | The secret provider name. Applicable when valueType is SECRET_REFERENCE. | [optional] [default to null]
**Value** | **string** | The string literal value. Applicable when valueType is STRING_LITERAL. | [optional] [default to null]
**ValueType** | **string** | The type of value (STRING_LITERAL, ASSET_REFERENCE, or SECRET_REFERENCE). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

