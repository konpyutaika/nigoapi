# ThreadDumpDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | **string** | The ID of the node in the cluster | [optional] [default to null]
**NodeAddress** | **string** | The address of the node in the cluster | [optional] [default to null]
**ApiPort** | **int32** | The port the node is listening for API requests. | [optional] [default to null]
**StackTrace** | **string** | The stack trace for the thread | [optional] [default to null]
**ThreadName** | **string** | The name of the thread | [optional] [default to null]
**ThreadActiveMillis** | **int64** | The number of milliseconds that the thread has been executing in the Processor | [optional] [default to null]
**TaskTerminated** | **bool** | Indicates whether or not the user has requested that the task be terminated. If this is true, it may indicate that the thread is in a state where it will continue running indefinitely without returning. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


