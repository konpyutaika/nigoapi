# ParameterProviderDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] [default to null]
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] [default to null]
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] [default to null]
**Position** | [***PositionDto**](PositionDTO.md) | The position of this component in the UI if applicable. | [optional] [default to null]
**Name** | **string** | The name of the parameter provider. | [optional] [default to null]
**Type_** | **string** | The fully qualified type of the parameter provider. | [optional] [default to null]
**Bundle** | [***BundleDto**](BundleDTO.md) | The details of the artifact that bundled this parameter provider type. | [optional] [default to null]
**Comments** | **string** | The comments of the parameter provider. | [optional] [default to null]
**PersistsState** | **bool** | Whether the parameter provider persists state. | [optional] [default to null]
**Restricted** | **bool** | Whether the parameter provider requires elevated privileges. | [optional] [default to null]
**Deprecated** | **bool** | Whether the parameter provider has been deprecated. | [optional] [default to null]
**MultipleVersionsAvailable** | **bool** | Whether the parameter provider has multiple versions available. | [optional] [default to null]
**Properties** | **map[string]string** | The properties of the parameter provider. | [optional] [default to null]
**Descriptors** | [**map[string]PropertyDescriptorDto**](PropertyDescriptorDTO.md) | The descriptors for the parameter providers properties. | [optional] [default to null]
**ParameterGroupConfigurations** | [**[]ParameterGroupConfigurationEntity**](ParameterGroupConfigurationEntity.md) | Configuration for any fetched parameter groups. | [optional] [default to null]
**AffectedComponents** | [**[]AffectedComponentEntity**](AffectedComponentEntity.md) | The set of all components in the flow that are referencing Parameters provided by this provider | [optional] [default to null]
**ParameterStatus** | [**[]ParameterStatusDto**](ParameterStatusDTO.md) | The status of all provided parameters for this parameter provider | [optional] [default to null]
**ReferencingParameterContexts** | [**[]ParameterProviderReferencingComponentEntity**](ParameterProviderReferencingComponentEntity.md) | The Parameter Contexts that reference this Parameter Provider | [optional] [default to null]
**CustomUiUrl** | **string** | The URL for the custom configuration UI for the parameter provider. | [optional] [default to null]
**AnnotationData** | **string** | The annotation data for the parameter provider. This is how the custom UI relays configuration to the parameter provider. | [optional] [default to null]
**ValidationErrors** | **[]string** | Gets the validation errors from the parameter provider. These validation errors represent the problems with the parameter provider that must be resolved before it can be scheduled to run. | [optional] [default to null]
**ValidationStatus** | **string** | Indicates whether the Parameter Provider is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the Parameter Provider is valid) | [optional] [default to null]
**ExtensionMissing** | **bool** | Whether the underlying extension is missing. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


