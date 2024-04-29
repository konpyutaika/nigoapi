# ConnectionStatisticsDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the connection | [optional] [default to null]
**StatsLastRefreshed** | **string** | The timestamp of when the stats were last refreshed | [optional] [default to null]
**AggregateSnapshot** | [***ConnectionStatisticsSnapshotDto**](ConnectionStatisticsSnapshotDTO.md) |  | [optional] [default to null]
**NodeSnapshots** | [**[]NodeConnectionStatisticsSnapshotDto**](NodeConnectionStatisticsSnapshotDTO.md) | A list of status snapshots for each node | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

