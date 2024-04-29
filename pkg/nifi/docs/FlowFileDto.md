# FlowFileDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | The URI that can be used to access this FlowFile. | [optional] [default to null]
**Uuid** | **string** | The FlowFile UUID. | [optional] [default to null]
**Filename** | **string** | The FlowFile filename. | [optional] [default to null]
**Position** | **int32** | The FlowFile&#x27;s position in the queue. | [optional] [default to null]
**Size** | **int64** | The FlowFile file size. | [optional] [default to null]
**QueuedDuration** | **int64** | How long this FlowFile has been enqueued. | [optional] [default to null]
**LineageDuration** | **int64** | Duration since the FlowFile&#x27;s greatest ancestor entered the flow. | [optional] [default to null]
**PenaltyExpiresIn** | **int64** | How long in milliseconds until the FlowFile penalty expires. | [optional] [default to null]
**ClusterNodeId** | **string** | The id of the node where this FlowFile resides. | [optional] [default to null]
**ClusterNodeAddress** | **string** | The label for the node where this FlowFile resides. | [optional] [default to null]
**Attributes** | **map[string]string** | The FlowFile attributes. | [optional] [default to null]
**ContentClaimSection** | **string** | The section in which the content claim lives. | [optional] [default to null]
**ContentClaimContainer** | **string** | The container in which the content claim lives. | [optional] [default to null]
**ContentClaimIdentifier** | **string** | The identifier of the content claim. | [optional] [default to null]
**ContentClaimOffset** | **int64** | The offset into the content claim where the flowfile&#x27;s content begins. | [optional] [default to null]
**ContentClaimFileSize** | **string** | The file size of the content claim formatted. | [optional] [default to null]
**ContentClaimFileSizeBytes** | **int64** | The file size of the content claim in bytes. | [optional] [default to null]
**Penalized** | **bool** | If the FlowFile is penalized. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

