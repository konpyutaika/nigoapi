# ReportingTaskDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] [default to null]
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] [default to null]
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] [default to null]
**Position** | [***PositionDto**](PositionDTO.md) | The position of this component in the UI if applicable. | [optional] [default to null]
**Name** | **string** | The name of the reporting task. | [optional] [default to null]
**Type_** | **string** | The fully qualified type of the reporting task. | [optional] [default to null]
**Bundle** | [***BundleDto**](BundleDTO.md) | The details of the artifact that bundled this processor type. | [optional] [default to null]
**State** | **string** | The state of the reporting task. | [optional] [default to null]
**Comments** | **string** | The comments of the reporting task. | [optional] [default to null]
**PersistsState** | **bool** | Whether the reporting task persists state. | [optional] [default to null]
**Restricted** | **bool** | Whether the reporting task requires elevated privileges. | [optional] [default to null]
**Deprecated** | **bool** | Whether the reporting task has been deprecated. | [optional] [default to null]
**MultipleVersionsAvailable** | **bool** | Whether the reporting task has multiple versions available. | [optional] [default to null]
**SchedulingPeriod** | **string** | The frequency with which to schedule the reporting task. The format of the value willd epend on the valud of the schedulingStrategy. | [optional] [default to null]
**SchedulingStrategy** | **string** | The scheduling strategy that determines how the schedulingPeriod value should be interpreted. | [optional] [default to null]
**DefaultSchedulingPeriod** | **map[string]string** | The default scheduling period for the different scheduling strategies. | [optional] [default to null]
**Properties** | **map[string]string** | The properties of the reporting task. | [optional] [default to null]
**Descriptors** | [**map[string]PropertyDescriptorDto**](PropertyDescriptorDTO.md) | The descriptors for the reporting tasks properties. | [optional] [default to null]
**CustomUiUrl** | **string** | The URL for the custom configuration UI for the reporting task. | [optional] [default to null]
**AnnotationData** | **string** | The annotation data for the repoting task. This is how the custom UI relays configuration to the reporting task. | [optional] [default to null]
**ValidationErrors** | **[]string** | Gets the validation errors from the reporting task. These validation errors represent the problems with the reporting task that must be resolved before it can be scheduled to run. | [optional] [default to null]
**ValidationStatus** | **string** | Indicates whether the Processor is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the Processor is valid) | [optional] [default to null]
**ActiveThreadCount** | **int32** | The number of active threads for the reporting task. | [optional] [default to null]
**ExtensionMissing** | **bool** | Whether the underlying extension is missing. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


