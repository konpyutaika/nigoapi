# FlowRegistryClientDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnotationData** | **string** | The annotation data for the registry client. This is how the custom UI relays configuration to the registry client. | [optional] [default to null]
**Bundle** | [***BundleDto**](BundleDTO.md) |  | [optional] [default to null]
**Deprecated** | **bool** | Whether the registry client has been deprecated. | [optional] [default to null]
**Description** | **string** | The registry description | [optional] [default to null]
**Descriptors** | [**map[string]PropertyDescriptorDto**](PropertyDescriptorDTO.md) | The descriptors for the registry client properties. | [optional] [default to null]
**ExtensionMissing** | **bool** | Whether the underlying extension is missing. | [optional] [default to null]
**Id** | **string** | The registry identifier | [optional] [default to null]
**MultipleVersionsAvailable** | **bool** | Whether the flow registry client has multiple versions available. | [optional] [default to null]
**Name** | **string** | The registry name | [optional] [default to null]
**Properties** | **map[string]string** | The properties of the registry client. | [optional] [default to null]
**Restricted** | **bool** | Whether the registry client requires elevated privileges. | [optional] [default to null]
**SensitiveDynamicPropertyNames** | **[]string** | Set of sensitive dynamic property names | [optional] [default to null]
**SupportsBranching** | **bool** | Whether the registry client supports branching. | [optional] [default to null]
**SupportsSensitiveDynamicProperties** | **bool** | Whether the registry client supports sensitive dynamic properties. | [optional] [default to null]
**Type_** | **string** | The type of the registry client. | [optional] [default to null]
**ValidationErrors** | **[]string** | Gets the validation errors from the registry client. These validation errors represent the problems with the registry client that must be resolved before it can be used for interacting with the flow registry. | [optional] [default to null]
**ValidationStatus** | **string** | Indicates whether the Registry Client is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the Registry Client is valid) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

