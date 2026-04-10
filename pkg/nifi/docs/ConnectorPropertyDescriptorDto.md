# ConnectorPropertyDescriptorDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowableValues** | [**[]AllowableValueEntity**](AllowableValueEntity.md) | Allowable values for the property. If empty then the allowed values are not constrained. | [optional] [default to null]
**AllowableValuesFetchable** | **bool** | Whether the allowable values are dynamically fetchable based on other property values. | [optional] [default to null]
**DefaultValue** | **string** | The default value for the property. | [optional] [default to null]
**Dependencies** | [**[]ConnectorPropertyDependencyDto**](ConnectorPropertyDependencyDTO.md) | The dependencies this property has on other properties. | [optional] [default to null]
**Description** | **string** | The description of the property. | [optional] [default to null]
**Name** | **string** | The name of the property. | [optional] [default to null]
**Required** | **bool** | Whether the property is required. | [optional] [default to null]
**Type_** | **string** | The type of the property (STRING, INTEGER, BOOLEAN, DOUBLE, STRING_LIST). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

