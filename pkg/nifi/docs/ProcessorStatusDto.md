# ProcessorStatusDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The unique ID of the process group that the Processor belongs to | [optional] [default to null]
**Id** | **string** | The unique ID of the Processor | [optional] [default to null]
**Name** | **string** | The name of the Processor | [optional] [default to null]
**Type_** | **string** | The type of the Processor | [optional] [default to null]
**RunStatus** | **string** | The run status of the Processor | [optional] [default to null]
**StatsLastRefreshed** | **string** | The timestamp of when the stats were last refreshed | [optional] [default to null]
**AggregateSnapshot** | [***ProcessorStatusSnapshotDto**](ProcessorStatusSnapshotDTO.md) |  | [optional] [default to null]
**NodeSnapshots** | [**[]NodeProcessorStatusSnapshotDto**](NodeProcessorStatusSnapshotDTO.md) | A status snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

