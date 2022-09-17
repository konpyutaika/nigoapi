# JvmDiagnosticsDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clustered** | **bool** | Whether or not the NiFi instance is clustered | [optional] [default to null]
**Connected** | **bool** | Whether or not the node is connected to the cluster | [optional] [default to null]
**AggregateSnapshot** | [***JvmDiagnosticsSnapshotDto**](JVMDiagnosticsSnapshotDTO.md) | Aggregate JVM diagnostic information about the entire cluster | [optional] [default to null]
**NodeSnapshots** | [**[]NodeJvmDiagnosticsSnapshotDto**](NodeJVMDiagnosticsSnapshotDTO.md) | Node-wise breakdown of JVM diagnostic information | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


