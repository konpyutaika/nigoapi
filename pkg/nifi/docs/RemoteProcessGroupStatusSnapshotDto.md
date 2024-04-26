# RemoteProcessGroupStatusSnapshotDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the remote process group. | [optional] [default to null]
**GroupId** | **string** | The id of the parent process group the remote process group resides in. | [optional] [default to null]
**Name** | **string** | The name of the remote process group. | [optional] [default to null]
**TargetUri** | **string** | The URI of the target system. | [optional] [default to null]
**TransmissionStatus** | **string** | The transmission status of the remote process group. | [optional] [default to null]
**ActiveThreadCount** | **int32** | The number of active threads for the remote process group. | [optional] [default to null]
**FlowFilesSent** | **int32** | The number of FlowFiles sent to the remote process group in the last 5 minutes. | [optional] [default to null]
**BytesSent** | **int64** | The size of the FlowFiles sent to the remote process group in the last 5 minutes. | [optional] [default to null]
**Sent** | **string** | The count/size of the flowfiles sent to the remote process group in the last 5 minutes. | [optional] [default to null]
**FlowFilesReceived** | **int32** | The number of FlowFiles received from the remote process group in the last 5 minutes. | [optional] [default to null]
**BytesReceived** | **int64** | The size of the FlowFiles received from the remote process group in the last 5 minutes. | [optional] [default to null]
**Received** | **string** | The count/size of the flowfiles received from the remote process group in the last 5 minutes. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

