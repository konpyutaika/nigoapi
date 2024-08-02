# NodeDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | **string** | The id of the node. | [optional] [default to null]
**Address** | **string** | The node&#x27;s host/ip address. | [optional] [default to null]
**ApiPort** | **int32** | The port the node is listening for API requests. | [optional] [default to null]
**Status** | **string** | The node&#x27;s status. | [optional] [default to null]
**Heartbeat** | **string** | the time of the nodes&#x27;s last heartbeat. | [optional] [default to null]
**ConnectionRequested** | **string** | The time of the node&#x27;s last connection request. | [optional] [default to null]
**Roles** | **[]string** | The roles of this node. | [optional] [default to null]
**ActiveThreadCount** | **int32** | The active threads for the NiFi on the node. | [optional] [default to null]
**Queued** | **string** | The queue the NiFi on the node. | [optional] [default to null]
**Events** | [**[]NodeEventDto**](NodeEventDTO.md) | The node&#x27;s events. | [optional] [default to null]
**NodeStartTime** | **string** | The time at which this Node was last refreshed. | [optional] [default to null]
**FlowFilesQueued** | **int32** | The number of FlowFiles that are queued up on the node | [optional] [default to null]
**BytesQueued** | **int64** | The total size of all FlowFiles that are queued up on the node | [optional] [default to null]
**FlowFileBytes** | **int64** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

