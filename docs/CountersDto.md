# CountersDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregateSnapshot** | [***CountersSnapshotDto**](CountersSnapshotDTO.md) | A Counters snapshot that represents the aggregate values of all nodes in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this represents the stats of the single instance. | [optional] [default to null]
**NodeSnapshots** | [**[]NodeCountersSnapshotDto**](NodeCountersSnapshotDTO.md) | A Counters snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


