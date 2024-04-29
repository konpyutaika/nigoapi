# ConnectionStatusSnapshotDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the connection. | [optional] [default to null]
**GroupId** | **string** | The id of the process group the connection belongs to. | [optional] [default to null]
**Name** | **string** | The name of the connection. | [optional] [default to null]
**SourceId** | **string** | The id of the source of the connection. | [optional] [default to null]
**SourceName** | **string** | The name of the source of the connection. | [optional] [default to null]
**DestinationId** | **string** | The id of the destination of the connection. | [optional] [default to null]
**DestinationName** | **string** | The name of the destination of the connection. | [optional] [default to null]
**Predictions** | [***ConnectionStatusPredictionsSnapshotDto**](ConnectionStatusPredictionsSnapshotDTO.md) |  | [optional] [default to null]
**FlowFilesIn** | **int32** | The number of FlowFiles that have come into the connection in the last 5 minutes. | [optional] [default to null]
**BytesIn** | **int64** | The size of the FlowFiles that have come into the connection in the last 5 minutes. | [optional] [default to null]
**Input** | **string** | The input count/size for the connection in the last 5 minutes, pretty printed. | [optional] [default to null]
**FlowFilesOut** | **int32** | The number of FlowFiles that have left the connection in the last 5 minutes. | [optional] [default to null]
**BytesOut** | **int64** | The number of bytes that have left the connection in the last 5 minutes. | [optional] [default to null]
**Output** | **string** | The output count/sie for the connection in the last 5 minutes, pretty printed. | [optional] [default to null]
**FlowFilesQueued** | **int32** | The number of FlowFiles that are currently queued in the connection. | [optional] [default to null]
**BytesQueued** | **int64** | The size of the FlowFiles that are currently queued in the connection. | [optional] [default to null]
**Queued** | **string** | The total count and size of queued flowfiles formatted. | [optional] [default to null]
**QueuedSize** | **string** | The total size of flowfiles that are queued formatted. | [optional] [default to null]
**QueuedCount** | **string** | The number of flowfiles that are queued, pretty printed. | [optional] [default to null]
**PercentUseCount** | **int32** | Connection percent use regarding queued flow files count and backpressure threshold if configured. | [optional] [default to null]
**PercentUseBytes** | **int32** | Connection percent use regarding queued flow files size and backpressure threshold if configured. | [optional] [default to null]
**FlowFileAvailability** | **string** | The availability of FlowFiles in this connection | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

