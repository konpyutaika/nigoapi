# ConnectorStatusDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregateSnapshot** | [***ConnectorStatusSnapshotDto**](ConnectorStatusSnapshotDTO.md) |  | [optional] [default to null]
**GroupId** | **string** | The id of the group this connector belongs to. | [optional] [default to null]
**Id** | **string** | The id of the connector. | [optional] [default to null]
**Name** | **string** | The name of the connector. | [optional] [default to null]
**NodeSnapshots** | [**[]NodeConnectorStatusSnapshotDto**](NodeConnectorStatusSnapshotDTO.md) | A status snapshot for each node in the cluster. If the NiFi instance is a standalone instance, rather than a cluster, this may be null. | [optional] [default to null]
**RunStatus** | **string** | The run status of the connector. | [optional] [default to null]
**StatsLastRefreshed** | **string** | The timestamp of when the stats were last refreshed. | [optional] [default to null]
**Type_** | **string** | The type of the connector. | [optional] [default to null]
**ValidationStatus** | **string** | The validation status of the connector. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

