# VersionedComponentState

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterState** | **map[string]string** | The cluster-scoped state of the component, or null if not exported | [optional] [default to null]
**LocalNodeStates** | [**[]VersionedNodeState**](VersionedNodeState.md) | The local-scoped state of the component ordered by node ordinal index, or null if not exported | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

