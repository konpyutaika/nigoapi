# ConnectionStatisticsDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the connection | [optional] [default to null]
**StatsLastRefreshed** | **string** | The timestamp of when the stats were last refreshed | [optional] [default to null]
**AggregateSnapshot** | [***ConnectionStatisticsSnapshotDto**](ConnectionStatisticsSnapshotDTO.md) | The status snapshot that represents the aggregate stats of the cluster | [optional] [default to null]
**NodeSnapshots** | [**[]NodeConnectionStatisticsSnapshotDto**](NodeConnectionStatisticsSnapshotDTO.md) | A list of status snapshots for each node | [optional] [default to null]

[[Back to Model list]](../pkg/nifi/README.md#documentation-for-models) [[Back to API list]](../pkg/nifi/README.md#documentation-for-api-endpoints) [[Back to README]](../pkg/nifi/README.md)


