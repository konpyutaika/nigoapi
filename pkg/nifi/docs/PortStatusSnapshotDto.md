# PortStatusSnapshotDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the port. | [optional] [default to null]
**GroupId** | **string** | The id of the parent process group of the port. | [optional] [default to null]
**Name** | **string** | The name of the port. | [optional] [default to null]
**ActiveThreadCount** | **int32** | The active thread count for the port. | [optional] [default to null]
**FlowFilesIn** | **int32** | The number of FlowFiles that have been accepted in the last 5 minutes. | [optional] [default to null]
**BytesIn** | **int64** | The size of hte FlowFiles that have been accepted in the last 5 minutes. | [optional] [default to null]
**Input** | **string** | The count/size of flowfiles that have been accepted in the last 5 minutes. | [optional] [default to null]
**FlowFilesOut** | **int32** | The number of FlowFiles that have been processed in the last 5 minutes. | [optional] [default to null]
**BytesOut** | **int64** | The number of bytes that have been processed in the last 5 minutes. | [optional] [default to null]
**Output** | **string** | The count/size of flowfiles that have been processed in the last 5 minutes. | [optional] [default to null]
**Transmitting** | **bool** | Whether the port has incoming or outgoing connections to a remote NiFi. | [optional] [default to null]
**RunStatus** | **string** | The run status of the port. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

