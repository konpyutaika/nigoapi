# ConnectorStatusSnapshotDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveThreadCount** | **int32** | The number of active threads for the connector. | [optional] [default to null]
**BytesQueued** | **int64** | The number of bytes queued in this connector&#x27;s managed process group. | [optional] [default to null]
**BytesRead** | **int64** | The number of bytes read by processors in this connector&#x27;s managed process group. | [optional] [default to null]
**BytesReceived** | **int64** | The number of bytes received by this connector&#x27;s managed process group. | [optional] [default to null]
**BytesSent** | **int64** | The number of bytes sent by this connector&#x27;s managed process group. | [optional] [default to null]
**BytesWritten** | **int64** | The number of bytes written by processors in this connector&#x27;s managed process group. | [optional] [default to null]
**FlowFilesQueued** | **int32** | The number of FlowFiles queued in this connector&#x27;s managed process group. | [optional] [default to null]
**FlowFilesReceived** | **int32** | The number of FlowFiles received by this connector&#x27;s managed process group. | [optional] [default to null]
**FlowFilesSent** | **int32** | The number of FlowFiles sent by this connector&#x27;s managed process group. | [optional] [default to null]
**GroupId** | **string** | The id of the parent process group to which the connector belongs. | [optional] [default to null]
**Id** | **string** | The id of the connector. | [optional] [default to null]
**Idle** | **bool** | Whether or not the connector is currently idle (no FlowFiles queued and no FlowFiles processed recently). | [optional] [default to null]
**IdleDuration** | **string** | A human-readable representation of how long the connector has been idle, or null if the connector is not idle. | [optional] [default to null]
**IdleDurationMillis** | **int64** | The number of milliseconds the connector has been idle, or null if the connector is not idle. | [optional] [default to null]
**Name** | **string** | The name of the connector. | [optional] [default to null]
**ProcessingPerformanceStatus** | [***ProcessingPerformanceStatusDto**](ProcessingPerformanceStatusDTO.md) |  | [optional] [default to null]
**Queued** | **string** | The count/size of queued data, pretty-printed. | [optional] [default to null]
**QueuedCount** | **string** | The count of queued FlowFiles, pretty-printed. | [optional] [default to null]
**QueuedSize** | **string** | The size of queued data, pretty-printed. | [optional] [default to null]
**Read** | **string** | The number of bytes read, pretty-printed. | [optional] [default to null]
**Received** | **string** | The count/size of data that has been received by this connector, pretty-printed. | [optional] [default to null]
**RunStatus** | **string** | The run status of the connector. | [optional] [default to null]
**Sent** | **string** | The count/size of data that has been sent by this connector, pretty-printed. | [optional] [default to null]
**Type_** | **string** | The type of the connector. | [optional] [default to null]
**Written** | **string** | The number of bytes written, pretty-printed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

