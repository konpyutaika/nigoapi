# ConnectionDiagnosticsDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connection** | [***ConnectionDto**](ConnectionDTO.md) | Details about the connection | [optional] [default to null]
**AggregateSnapshot** | [***ConnectionDiagnosticsSnapshotDto**](ConnectionDiagnosticsSnapshotDTO.md) | Aggregate values for all nodes in the cluster, or for this instance if not clustered | [optional] [default to null]
**NodeSnapshots** | [**[]ConnectionDiagnosticsSnapshotDto**](ConnectionDiagnosticsSnapshotDTO.md) | A list of values for each node in the cluster, if clustered. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


