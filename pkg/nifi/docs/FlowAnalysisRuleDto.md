# FlowAnalysisRuleDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] [default to null]
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] [default to null]
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] [default to null]
**Position** | [***PositionDto**](PositionDTO.md) |  | [optional] [default to null]
**Name** | **string** | The name of the flow analysis rule. | [optional] [default to null]
**Type_** | **string** | The fully qualified type of the flow analysis rule. | [optional] [default to null]
**Bundle** | [***BundleDto**](BundleDTO.md) |  | [optional] [default to null]
**State** | **string** | The state of the flow analysis rule. | [optional] [default to null]
**Comments** | **string** | The comments of the flow analysis rule. | [optional] [default to null]
**PersistsState** | **bool** | Whether the flow analysis rule persists state. | [optional] [default to null]
**Restricted** | **bool** | Whether the flow analysis rule requires elevated privileges. | [optional] [default to null]
**Deprecated** | **bool** | Whether the flow analysis rule has been deprecated. | [optional] [default to null]
**MultipleVersionsAvailable** | **bool** | Whether the flow analysis rule has multiple versions available. | [optional] [default to null]
**SupportsSensitiveDynamicProperties** | **bool** | Whether the flow analysis rule supports sensitive dynamic properties. | [optional] [default to null]
**EnforcementPolicy** | **string** | Enforcement Policy. | [optional] [default to null]
**Properties** | **map[string]string** | The properties of the flow analysis rule. | [optional] [default to null]
**Descriptors** | [**map[string]PropertyDescriptorDto**](PropertyDescriptorDTO.md) | The descriptors for the flow analysis rules properties. | [optional] [default to null]
**SensitiveDynamicPropertyNames** | **[]string** | Set of sensitive dynamic property names | [optional] [default to null]
**ValidationErrors** | **[]string** | Gets the validation errors from the flow analysis rule. These validation errors represent the problems with the flow analysis rule that must be resolved before it can be scheduled to run. | [optional] [default to null]
**ValidationStatus** | **string** | Indicates whether the Flow Analysis Rule is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the Flow Analysis Rule is valid) | [optional] [default to null]
**ExtensionMissing** | **bool** | Whether the underlying extension is missing. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

