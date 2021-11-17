# ProcessorDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] [default to null]
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] [default to null]
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] [default to null]
**Position** | [***PositionDto**](PositionDTO.md) | The position of this component in the UI if applicable. | [optional] [default to null]
**Name** | **string** | The name of the processor. | [optional] [default to null]
**Type_** | **string** | The type of the processor. | [optional] [default to null]
**Bundle** | [***BundleDto**](BundleDTO.md) | The details of the artifact that bundled this processor type. | [optional] [default to null]
**State** | **string** | The state of the processor | [optional] [default to null]
**Style** | **map[string]string** | Styles for the processor (background-color : #eee). | [optional] [default to null]
**Relationships** | [**[]RelationshipDto**](RelationshipDTO.md) | The available relationships that the processor currently supports. | [optional] [default to null]
**Description** | **string** | The description of the processor. | [optional] [default to null]
**SupportsParallelProcessing** | **bool** | Whether the processor supports parallel processing. | [optional] [default to null]
**SupportsEventDriven** | **bool** | Whether the processor supports event driven scheduling. | [optional] [default to null]
**SupportsBatching** | **bool** | Whether the processor supports batching. This makes the run duration settings available. | [optional] [default to null]
**PersistsState** | **bool** | Whether the processor persists state. | [optional] [default to null]
**Restricted** | **bool** | Whether the processor requires elevated privileges. | [optional] [default to null]
**Deprecated** | **bool** | Whether the processor has been deprecated. | [optional] [default to null]
**ExecutionNodeRestricted** | **bool** | Indicates if the execution node of a processor is restricted to run only on the primary node | [optional] [default to null]
**MultipleVersionsAvailable** | **bool** | Whether the processor has multiple versions available. | [optional] [default to null]
**InputRequirement** | **string** | The input requirement for this processor. | [optional] [default to null]
**Config** | [***ProcessorConfigDto**](ProcessorConfigDTO.md) | The configuration details for the processor. These details will be included in a response if the verbose flag is included in a request. | [optional] [default to null]
**ValidationErrors** | **[]string** | The validation errors for the processor. These validation errors represent the problems with the processor that must be resolved before it can be started. | [optional] [default to null]
**ValidationStatus** | **string** | Indicates whether the Processor is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the Processor is valid) | [optional] [default to null]
**ExtensionMissing** | **bool** | Whether the underlying extension is missing. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


