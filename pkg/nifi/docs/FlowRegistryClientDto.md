# FlowRegistryClientDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The registry identifier | [optional] [default to null]
**Name** | **string** | The registry name | [optional] [default to null]
**Description** | **string** | The registry description | [optional] [default to null]
**Uri** | **string** |  | [optional] [default to null]
**Type_** | **string** | The type of the controller service. | [optional] [default to null]
**Bundle** | [***BundleDto**](BundleDTO.md) | The details of the artifact that bundled this processor type. | [optional] [default to null]
**Properties** | **map[string]string** | The properties of the controller service. | [optional] [default to null]
**Descriptors** | [**map[string]PropertyDescriptorDto**](PropertyDescriptorDTO.md) | The descriptors for the controller service properties. | [optional] [default to null]
**SensitiveDynamicPropertyNames** | **[]string** | Set of sensitive dynamic property names | [optional] [default to null]
**SupportsSensitiveDynamicProperties** | **bool** | Whether the reporting task supports sensitive dynamic properties. | [optional] [default to null]
**Restricted** | **bool** | Whether the reporting task requires elevated privileges. | [optional] [default to null]
**Deprecated** | **bool** | Whether the reporting task has been deprecated. | [optional] [default to null]
**ValidationErrors** | **[]string** | Gets the validation errors from the reporting task. These validation errors represent the problems with the reporting task that must be resolved before it can be scheduled to run. | [optional] [default to null]
**ValidationStatus** | **string** | Indicates whether the Processor is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the Processor is valid) | [optional] [default to null]
**AnnotationData** | **string** | The annotation data for the repoting task. This is how the custom UI relays configuration to the reporting task. | [optional] [default to null]
**MultipleVersionsAvailable** | **bool** | Whether the flow registry client has multiple versions available. | [optional] [default to null]
**ExtensionMissing** | **bool** | Whether the underlying extension is missing. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


