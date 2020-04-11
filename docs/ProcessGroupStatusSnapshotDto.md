# ProcessGroupStatusSnapshotDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the process group. | [optional] [default to null]
**Name** | **string** | The name of this process group. | [optional] [default to null]
**ConnectionStatusSnapshots** | [**[]ConnectionStatusSnapshotEntity**](ConnectionStatusSnapshotEntity.md) | The status of all connections in the process group. | [optional] [default to null]
**ProcessorStatusSnapshots** | [**[]ProcessorStatusSnapshotEntity**](ProcessorStatusSnapshotEntity.md) | The status of all processors in the process group. | [optional] [default to null]
**ProcessGroupStatusSnapshots** | [**[]ProcessGroupStatusSnapshotEntity**](ProcessGroupStatusSnapshotEntity.md) | The status of all process groups in the process group. | [optional] [default to null]
**RemoteProcessGroupStatusSnapshots** | [**[]RemoteProcessGroupStatusSnapshotEntity**](RemoteProcessGroupStatusSnapshotEntity.md) | The status of all remote process groups in the process group. | [optional] [default to null]
**InputPortStatusSnapshots** | [**[]PortStatusSnapshotEntity**](PortStatusSnapshotEntity.md) | The status of all input ports in the process group. | [optional] [default to null]
**OutputPortStatusSnapshots** | [**[]PortStatusSnapshotEntity**](PortStatusSnapshotEntity.md) | The status of all output ports in the process group. | [optional] [default to null]
**VersionedFlowState** | **string** | The current state of the Process Group, as it relates to the Versioned Flow | [optional] [default to null]
**FlowFilesIn** | **int32** | The number of FlowFiles that have come into this ProcessGroup in the last 5 minutes | [optional] [default to null]
**BytesIn** | **int64** | The number of bytes that have come into this ProcessGroup in the last 5 minutes | [optional] [default to null]
**Input** | **string** | The input count/size for the process group in the last 5 minutes (pretty printed). | [optional] [default to null]
**FlowFilesQueued** | **int32** | The number of FlowFiles that are queued up in this ProcessGroup right now | [optional] [default to null]
**BytesQueued** | **int64** | The number of bytes that are queued up in this ProcessGroup right now | [optional] [default to null]
**Queued** | **string** | The count/size that is queued in the the process group. | [optional] [default to null]
**QueuedCount** | **string** | The count that is queued for the process group. | [optional] [default to null]
**QueuedSize** | **string** | The size that is queued for the process group. | [optional] [default to null]
**BytesRead** | **int64** | The number of bytes read by components in this ProcessGroup in the last 5 minutes | [optional] [default to null]
**Read** | **string** | The number of bytes read in the last 5 minutes. | [optional] [default to null]
**BytesWritten** | **int64** | The number of bytes written by components in this ProcessGroup in the last 5 minutes | [optional] [default to null]
**Written** | **string** | The number of bytes written in the last 5 minutes. | [optional] [default to null]
**FlowFilesOut** | **int32** | The number of FlowFiles transferred out of this ProcessGroup in the last 5 minutes | [optional] [default to null]
**BytesOut** | **int64** | The number of bytes transferred out of this ProcessGroup in the last 5 minutes | [optional] [default to null]
**Output** | **string** | The output count/size for the process group in the last 5 minutes. | [optional] [default to null]
**FlowFilesTransferred** | **int32** | The number of FlowFiles transferred in this ProcessGroup in the last 5 minutes | [optional] [default to null]
**BytesTransferred** | **int64** | The number of bytes transferred in this ProcessGroup in the last 5 minutes | [optional] [default to null]
**Transferred** | **string** | The count/size transferred to/from queues in the process group in the last 5 minutes. | [optional] [default to null]
**BytesReceived** | **int64** | The number of bytes received from external sources by components within this ProcessGroup in the last 5 minutes | [optional] [default to null]
**FlowFilesReceived** | **int32** | The number of FlowFiles received from external sources by components within this ProcessGroup in the last 5 minutes | [optional] [default to null]
**Received** | **string** | The count/size sent to the process group in the last 5 minutes. | [optional] [default to null]
**BytesSent** | **int64** | The number of bytes sent to an external sink by components within this ProcessGroup in the last 5 minutes | [optional] [default to null]
**FlowFilesSent** | **int32** | The number of FlowFiles sent to an external sink by components within this ProcessGroup in the last 5 minutes | [optional] [default to null]
**Sent** | **string** | The count/size sent from this process group in the last 5 minutes. | [optional] [default to null]
**ActiveThreadCount** | **int32** | The active thread count for this process group. | [optional] [default to null]
**TerminatedThreadCount** | **int32** | The number of threads currently terminated for the process group. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


