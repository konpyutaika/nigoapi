# ProcessorStatusSnapshotDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the processor. | [optional] [default to null]
**GroupId** | **string** | The id of the parent process group to which the processor belongs. | [optional] [default to null]
**Name** | **string** | The name of the prcessor. | [optional] [default to null]
**Type_** | **string** | The type of the processor. | [optional] [default to null]
**RunStatus** | **string** | The state of the processor. | [optional] [default to null]
**ExecutionNode** | **string** | Indicates the node where the process will execute. | [optional] [default to null]
**BytesRead** | **int64** | The number of bytes read by this Processor in the last 5 mintues | [optional] [default to null]
**BytesWritten** | **int64** | The number of bytes written by this Processor in the last 5 minutes | [optional] [default to null]
**Read** | **string** | The number of bytes read in the last 5 minutes. | [optional] [default to null]
**Written** | **string** | The number of bytes written in the last 5 minutes. | [optional] [default to null]
**FlowFilesIn** | **int32** | The number of FlowFiles that have been accepted in the last 5 minutes | [optional] [default to null]
**BytesIn** | **int64** | The size of the FlowFiles that have been accepted in the last 5 minutes | [optional] [default to null]
**Input** | **string** | The count/size of flowfiles that have been accepted in the last 5 minutes. | [optional] [default to null]
**FlowFilesOut** | **int32** | The number of FlowFiles transferred to a Connection in the last 5 minutes | [optional] [default to null]
**BytesOut** | **int64** | The size of the FlowFiles transferred to a Connection in the last 5 minutes | [optional] [default to null]
**Output** | **string** | The count/size of flowfiles that have been processed in the last 5 minutes. | [optional] [default to null]
**TaskCount** | **int32** | The number of times this Processor has run in the last 5 minutes | [optional] [default to null]
**TasksDurationNanos** | **int64** | The number of nanoseconds that this Processor has spent running in the last 5 minutes | [optional] [default to null]
**Tasks** | **string** | The total number of task this connectable has completed over the last 5 minutes. | [optional] [default to null]
**TasksDuration** | **string** | The total duration of all tasks for this connectable over the last 5 minutes. | [optional] [default to null]
**ActiveThreadCount** | **int32** | The number of threads currently executing in the processor. | [optional] [default to null]
**TerminatedThreadCount** | **int32** | The number of threads currently terminated for the processor. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

