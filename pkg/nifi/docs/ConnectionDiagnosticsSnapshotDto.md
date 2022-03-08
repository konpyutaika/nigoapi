# ConnectionDiagnosticsSnapshotDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalFlowFileCount** | **int32** | Total number of FlowFiles owned by the Connection | [optional] [default to null]
**TotalByteCount** | **int64** | Total number of bytes that make up the content for the FlowFiles owned by this Connection | [optional] [default to null]
**NodeIdentifier** | **string** | The Node Identifier that this information pertains to | [optional] [default to null]
**LocalQueuePartition** | [***LocalQueuePartitionDto**](LocalQueuePartitionDTO.md) | The local queue partition, from which components can pull FlowFiles on this node. | [optional] [default to null]
**RemoteQueuePartitions** | [**[]RemoteQueuePartitionDto**](RemoteQueuePartitionDTO.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


