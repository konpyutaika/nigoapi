# ProcessingPerformanceStatusDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The unique ID of the process group that the Processor belongs to | [optional] [default to null]
**CpuDuration** | **int64** | The number of nanoseconds has spent on CPU usage in the last 5 minutes. | [optional] [default to null]
**ContentReadDuration** | **int64** | The number of nanoseconds has spent to read content in the last 5 minutes. | [optional] [default to null]
**ContentWriteDuration** | **int64** | The number of nanoseconds has spent to write content in the last 5 minutes. | [optional] [default to null]
**SessionCommitDuration** | **int64** | The number of nanoseconds has spent running to commit sessions the last 5 minutes. | [optional] [default to null]
**GarbageCollectionDuration** | **int64** | The number of nanoseconds has spent running garbage collection in the last 5 minutes. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

