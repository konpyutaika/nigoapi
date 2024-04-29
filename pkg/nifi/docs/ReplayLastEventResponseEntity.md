# ReplayLastEventResponseEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentId** | **string** | The UUID of the component whose last event should be replayed. | [optional] [default to null]
**Nodes** | **string** | Which nodes were requested to replay their last provenance event. | [optional] [default to null]
**AggregateSnapshot** | [***ReplayLastEventSnapshotDto**](ReplayLastEventSnapshotDTO.md) |  | [optional] [default to null]
**NodeSnapshots** | [**[]NodeReplayLastEventSnapshotDto**](NodeReplayLastEventSnapshotDTO.md) | The node-wise results | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

