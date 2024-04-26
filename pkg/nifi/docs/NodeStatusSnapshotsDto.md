# NodeStatusSnapshotsDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | **string** | The id of the node. | [optional] [default to null]
**Address** | **string** | The node&#x27;s host/ip address. | [optional] [default to null]
**ApiPort** | **int32** | The port the node is listening for API requests. | [optional] [default to null]
**StatusSnapshots** | [**[]StatusSnapshotDto**](StatusSnapshotDTO.md) | A list of StatusSnapshotDTO objects that provide the actual metric values for the component for this node. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

