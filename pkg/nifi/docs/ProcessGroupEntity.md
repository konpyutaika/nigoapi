# ProcessGroupEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | [***RevisionDto**](RevisionDTO.md) | The revision for this request/response. The revision is required for any mutable flow requests and is included in all responses. | [optional] [default to null]
**Id** | **string** | The id of the component. | [optional] [default to null]
**Uri** | **string** | The URI for futures requests to the component. | [optional] [default to null]
**Position** | [***PositionDto**](PositionDTO.md) | The position of this component in the UI if applicable. | [optional] [default to null]
**Permissions** | [***PermissionsDto**](PermissionsDTO.md) | The permissions for this component. | [optional] [default to null]
**Bulletins** | [**[]BulletinEntity**](BulletinEntity.md) | The bulletins for this component. | [optional] [default to null]
**DisconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] [default to null]
**Component** | [***ProcessGroupDto**](ProcessGroupDTO.md) |  | [optional] [default to null]
**Status** | [***ProcessGroupStatusDto**](ProcessGroupStatusDTO.md) | The status of the process group. | [optional] [default to null]
**VersionedFlowSnapshot** | [***VersionedFlowSnapshot**](VersionedFlowSnapshot.md) | Returns the Versioned Flow that describes the contents of the Versioned Flow to be imported | [optional] [default to null]
**RunningCount** | **int32** | The number of running components in this process group. | [optional] [default to null]
**StoppedCount** | **int32** | The number of stopped components in the process group. | [optional] [default to null]
**InvalidCount** | **int32** | The number of invalid components in the process group. | [optional] [default to null]
**DisabledCount** | **int32** | The number of disabled components in the process group. | [optional] [default to null]
**ActiveRemotePortCount** | **int32** | The number of active remote ports in the process group. | [optional] [default to null]
**InactiveRemotePortCount** | **int32** | The number of inactive remote ports in the process group. | [optional] [default to null]
**VersionedFlowState** | **string** | The current state of the Process Group, as it relates to the Versioned Flow | [optional] [default to null]
**UpToDateCount** | **int32** | The number of up to date versioned process groups in the process group. | [optional] [default to null]
**LocallyModifiedCount** | **int32** | The number of locally modified versioned process groups in the process group. | [optional] [default to null]
**StaleCount** | **int32** | The number of stale versioned process groups in the process group. | [optional] [default to null]
**LocallyModifiedAndStaleCount** | **int32** | The number of locally modified and stale versioned process groups in the process group. | [optional] [default to null]
**SyncFailureCount** | **int32** | The number of versioned process groups in the process group that are unable to sync to a registry. | [optional] [default to null]
**LocalInputPortCount** | **int32** | The number of local input ports in the process group. | [optional] [default to null]
**LocalOutputPortCount** | **int32** | The number of local output ports in the process group. | [optional] [default to null]
**PublicInputPortCount** | **int32** | The number of public input ports in the process group. | [optional] [default to null]
**PublicOutputPortCount** | **int32** | The number of public output ports in the process group. | [optional] [default to null]
**ParameterContext** | [***ParameterContextReferenceEntity**](ParameterContextReferenceEntity.md) | The Parameter Context, or null if no Parameter Context has been bound to the Process Group | [optional] [default to null]
**InputPortCount** | **int32** | The number of input ports in the process group. | [optional] [default to null]
**OutputPortCount** | **int32** | The number of output ports in the process group. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


