# PropertyDescriptorDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name for the property. | [optional] [default to null]
**DisplayName** | **string** | The human readable name for the property. | [optional] [default to null]
**Description** | **string** | The description for the property. Used to relay additional details to a user or provide a mechanism of documenting intent. | [optional] [default to null]
**DefaultValue** | **string** | The default value for the property. | [optional] [default to null]
**AllowableValues** | [**[]AllowableValueEntity**](AllowableValueEntity.md) | Allowable values for the property. If empty then the allowed values are not constrained. | [optional] [default to null]
**Required** | **bool** | Whether the property is required. | [optional] [default to null]
**Sensitive** | **bool** | Whether the property is sensitive and protected whenever stored or represented. | [optional] [default to null]
**Dynamic** | **bool** | Whether the property is dynamic (user-defined). | [optional] [default to null]
**SupportsEl** | **bool** | Whether the property supports expression language. | [optional] [default to null]
**ExpressionLanguageScope** | **string** | Scope of the Expression Language evaluation for the property. | [optional] [default to null]
**IdentifiesControllerService** | **string** | If the property identifies a controller service this returns the fully qualified type. | [optional] [default to null]
**IdentifiesControllerServiceBundle** | [***BundleDto**](BundleDTO.md) |  | [optional] [default to null]
**Dependencies** | [**[]PropertyDependencyDto**](PropertyDependencyDTO.md) | A list of dependencies that must be met in order for this Property to be relevant. If any of these dependencies is not met, the property described by this Property Descriptor is not relevant. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

