# SystemDiagnosticsDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregateSnapshot** | [***SystemDiagnosticsSnapshotDto**](SystemDiagnosticsSnapshotDTO.md) |  | [optional] [default to null]
**NodeSnapshots** | [**[]NodeSystemDiagnosticsSnapshotDto**](NodeSystemDiagnosticsSnapshotDTO.md) | A systems diagnostics snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

