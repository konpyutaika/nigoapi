# VersionedFlowSnapshotEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] [default to null]
**ProcessGroupRevision** | [***RevisionDto**](RevisionDTO.md) |  | [optional] [default to null]
**RegistryId** | **string** | The ID of the Registry that this flow belongs to | [optional] [default to null]
**UpdateDescendantVersionedFlows** | **bool** | If the Process Group to be updated has a child or descendant Process Group that is also under Version Control, this specifies whether or not the contents of that child/descendant Process Group should be updated. | [optional] [default to null]
**VersionedFlow** | [***RegisteredFlowSnapshot**](RegisteredFlowSnapshot.md) |  | [optional] [default to null]
**VersionedFlowSnapshot** | [***RegisteredFlowSnapshot**](RegisteredFlowSnapshot.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

