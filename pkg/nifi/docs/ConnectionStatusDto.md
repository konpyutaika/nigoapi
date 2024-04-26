# ConnectionStatusDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the connection | [optional] [default to null]
**GroupId** | **string** | The ID of the Process Group that the connection belongs to | [optional] [default to null]
**Name** | **string** | The name of the connection | [optional] [default to null]
**StatsLastRefreshed** | **string** | The timestamp of when the stats were last refreshed | [optional] [default to null]
**SourceId** | **string** | The ID of the source component | [optional] [default to null]
**SourceName** | **string** | The name of the source component | [optional] [default to null]
**DestinationId** | **string** | The ID of the destination component | [optional] [default to null]
**DestinationName** | **string** | The name of the destination component | [optional] [default to null]
**AggregateSnapshot** | [***ConnectionStatusSnapshotDto**](ConnectionStatusSnapshotDTO.md) |  | [optional] [default to null]
**NodeSnapshots** | [**[]NodeConnectionStatusSnapshotDto**](NodeConnectionStatusSnapshotDTO.md) | A list of status snapshots for each node | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

