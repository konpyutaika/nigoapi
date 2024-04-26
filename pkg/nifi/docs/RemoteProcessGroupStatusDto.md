# RemoteProcessGroupStatusDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The unique ID of the process group that the Processor belongs to | [optional] [default to null]
**Id** | **string** | The unique ID of the Processor | [optional] [default to null]
**Name** | **string** | The name of the remote process group. | [optional] [default to null]
**TargetUri** | **string** | The URI of the target system. | [optional] [default to null]
**TransmissionStatus** | **string** | The transmission status of the remote process group. | [optional] [default to null]
**StatsLastRefreshed** | **string** | The time the status for the process group was last refreshed. | [optional] [default to null]
**ValidationStatus** | **string** | Indicates whether the component is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the component is valid) | [optional] [default to null]
**AggregateSnapshot** | [***RemoteProcessGroupStatusSnapshotDto**](RemoteProcessGroupStatusSnapshotDTO.md) |  | [optional] [default to null]
**NodeSnapshots** | [**[]NodeRemoteProcessGroupStatusSnapshotDto**](NodeRemoteProcessGroupStatusSnapshotDTO.md) | A status snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

