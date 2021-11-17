# ControllerServiceDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] [default to null]
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] [default to null]
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] [default to null]
**Position** | [***PositionDto**](PositionDTO.md) | The position of this component in the UI if applicable. | [optional] [default to null]
**Name** | **string** | The name of the controller service. | [optional] [default to null]
**Type_** | **string** | The type of the controller service. | [optional] [default to null]
**Bundle** | [***BundleDto**](BundleDTO.md) | The details of the artifact that bundled this processor type. | [optional] [default to null]
**ControllerServiceApis** | [**[]ControllerServiceApiDto**](ControllerServiceApiDTO.md) | Lists the APIs this Controller Service implements. | [optional] [default to null]
**Comments** | **string** | The comments for the controller service. | [optional] [default to null]
**State** | **string** | The state of the controller service. | [optional] [default to null]
**PersistsState** | **bool** | Whether the controller service persists state. | [optional] [default to null]
**Restricted** | **bool** | Whether the controller service requires elevated privileges. | [optional] [default to null]
**Deprecated** | **bool** | Whether the ontroller service has been deprecated. | [optional] [default to null]
**MultipleVersionsAvailable** | **bool** | Whether the controller service has multiple versions available. | [optional] [default to null]
**Properties** | **map[string]string** | The properties of the controller service. | [optional] [default to null]
**Descriptors** | [**map[string]PropertyDescriptorDto**](PropertyDescriptorDTO.md) | The descriptors for the controller service properties. | [optional] [default to null]
**CustomUiUrl** | **string** | The URL for the controller services custom configuration UI if applicable. | [optional] [default to null]
**AnnotationData** | **string** | The annotation for the controller service. This is how the custom UI relays configuration to the controller service. | [optional] [default to null]
**ReferencingComponents** | [**[]ControllerServiceReferencingComponentEntity**](ControllerServiceReferencingComponentEntity.md) | All components referencing this controller service. | [optional] [default to null]
**ValidationErrors** | **[]string** | The validation errors from the controller service. These validation errors represent the problems with the controller service that must be resolved before it can be enabled. | [optional] [default to null]
**ValidationStatus** | **string** | Indicates whether the ControllerService is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the ControllerService is valid) | [optional] [default to null]
**ExtensionMissing** | **bool** | Whether the underlying extension is missing. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


