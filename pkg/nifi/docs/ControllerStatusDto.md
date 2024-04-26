# ControllerStatusDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveThreadCount** | **int32** | The number of active threads in the NiFi. | [optional] [default to null]
**TerminatedThreadCount** | **int32** | The number of terminated threads in the NiFi. | [optional] [default to null]
**Queued** | **string** | The number of flowfiles queued in the NiFi. | [optional] [default to null]
**FlowFilesQueued** | **int32** | The number of FlowFiles queued across the entire flow | [optional] [default to null]
**BytesQueued** | **int64** | The size of the FlowFiles queued across the entire flow | [optional] [default to null]
**RunningCount** | **int32** | The number of running components in the NiFi. | [optional] [default to null]
**StoppedCount** | **int32** | The number of stopped components in the NiFi. | [optional] [default to null]
**InvalidCount** | **int32** | The number of invalid components in the NiFi. | [optional] [default to null]
**DisabledCount** | **int32** | The number of disabled components in the NiFi. | [optional] [default to null]
**ActiveRemotePortCount** | **int32** | The number of active remote ports in the NiFi. | [optional] [default to null]
**InactiveRemotePortCount** | **int32** | The number of inactive remote ports in the NiFi. | [optional] [default to null]
**UpToDateCount** | **int32** | The number of up to date versioned process groups in the NiFi. | [optional] [default to null]
**LocallyModifiedCount** | **int32** | The number of locally modified versioned process groups in the NiFi. | [optional] [default to null]
**StaleCount** | **int32** | The number of stale versioned process groups in the NiFi. | [optional] [default to null]
**LocallyModifiedAndStaleCount** | **int32** | The number of locally modified and stale versioned process groups in the NiFi. | [optional] [default to null]
**SyncFailureCount** | **int32** | The number of versioned process groups in the NiFi that are unable to sync to a registry. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

