# ScheduleComponentsEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the ProcessGroup | [optional] [default to null]
**State** | **string** | The desired state of the descendant components | [optional] [default to null]
**Components** | [**map[string]RevisionDto**](RevisionDTO.md) | Optional components to schedule. If not specified, all authorized descendant components will be used. | [optional] [default to null]
**DisconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

