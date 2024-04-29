# ProvenanceNodeDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the node. | [optional] [default to null]
**FlowFileUuid** | **string** | The uuid of the flowfile associated with the provenance event. | [optional] [default to null]
**ParentUuids** | **[]string** | The uuid of the parent flowfiles of the provenance event. | [optional] [default to null]
**ChildUuids** | **[]string** | The uuid of the childrent flowfiles of the provenance event. | [optional] [default to null]
**ClusterNodeIdentifier** | **string** | The identifier of the node that this event/flowfile originated from. | [optional] [default to null]
**Type_** | **string** | The type of the node. | [optional] [default to null]
**EventType** | **string** | If the type is EVENT, this is the type of event. | [optional] [default to null]
**Millis** | **int64** | The timestamp of the node in milliseconds. | [optional] [default to null]
**Timestamp** | **string** | The timestamp of the node formatted. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

