# ProcessGroupDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] [default to null]
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] [default to null]
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] [default to null]
**Position** | [***PositionDto**](PositionDTO.md) |  | [optional] [default to null]
**Name** | **string** | The name of the process group. | [optional] [default to null]
**Comments** | **string** | The comments for the process group. | [optional] [default to null]
**VersionControlInformation** | [***VersionControlInformationDto**](VersionControlInformationDTO.md) |  | [optional] [default to null]
**ParameterContext** | [***ParameterContextReferenceEntity**](ParameterContextReferenceEntity.md) |  | [optional] [default to null]
**FlowfileConcurrency** | **string** | The FlowFile Concurrency for this Process Group. | [optional] [default to null]
**FlowfileOutboundPolicy** | **string** | The Outbound Policy that is used for determining how FlowFiles should be transferred out of the Process Group. | [optional] [default to null]
**DefaultFlowFileExpiration** | **string** | The default FlowFile Expiration for this Process Group. | [optional] [default to null]
**DefaultBackPressureObjectThreshold** | **int64** | Default value used in this Process Group for the maximum number of objects that can be queued before back pressure is applied. | [optional] [default to null]
**DefaultBackPressureDataSizeThreshold** | **string** | Default value used in this Process Group for the maximum data size of objects that can be queued before back pressure is applied. | [optional] [default to null]
**LogFileSuffix** | **string** | The log file suffix for this Process Group for dedicated logging. | [optional] [default to null]
**ExecutionEngine** | **string** | The Execution Engine that should be used to run the flow represented by this Process Group. | [optional] [default to null]
**MaxConcurrentTasks** | **int32** | The maximum number of concurrent tasks to use when running the flow using the Stateless Engine | [optional] [default to null]
**StatelessFlowTimeout** | **string** | The maximum amount of time that the flow can be run using the Stateless Engine before the flow times out | [optional] [default to null]
**RunningCount** | **int32** | The number of running components in this process group. | [optional] [default to null]
**StoppedCount** | **int32** | The number of stopped components in the process group. | [optional] [default to null]
**InvalidCount** | **int32** | The number of invalid components in the process group. | [optional] [default to null]
**DisabledCount** | **int32** | The number of disabled components in the process group. | [optional] [default to null]
**ActiveRemotePortCount** | **int32** | The number of active remote ports in the process group. | [optional] [default to null]
**InactiveRemotePortCount** | **int32** | The number of inactive remote ports in the process group. | [optional] [default to null]
**UpToDateCount** | **int32** | The number of up to date versioned process groups in the process group. | [optional] [default to null]
**LocallyModifiedCount** | **int32** | The number of locally modified versioned process groups in the process group. | [optional] [default to null]
**StaleCount** | **int32** | The number of stale versioned process groups in the process group. | [optional] [default to null]
**LocallyModifiedAndStaleCount** | **int32** | The number of locally modified and stale versioned process groups in the process group. | [optional] [default to null]
**SyncFailureCount** | **int32** | The number of versioned process groups in the process group that are unable to sync to a registry. | [optional] [default to null]
**LocalInputPortCount** | **int32** | The number of local input ports in the process group. | [optional] [default to null]
**LocalOutputPortCount** | **int32** | The number of local output ports in the process group. | [optional] [default to null]
**PublicInputPortCount** | **int32** | The number of public input ports in the process group. | [optional] [default to null]
**PublicOutputPortCount** | **int32** | The number of public output ports in the process group. | [optional] [default to null]
**StatelessGroupScheduledState** | **string** | If the Process Group is configured to run in using the Stateless Engine, represents the current state. Otherwise, will be STOPPED. | [optional] [default to null]
**Contents** | [***FlowSnippetDto**](FlowSnippetDTO.md) |  | [optional] [default to null]
**OutputPortCount** | **int32** | The number of output ports in the process group. | [optional] [default to null]
**InputPortCount** | **int32** | The number of input ports in the process group. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

