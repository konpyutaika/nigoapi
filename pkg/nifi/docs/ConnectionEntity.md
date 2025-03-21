# ConnectionEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bends** | [**[]PositionDto**](PositionDTO.md) | The bend points on the connection. | [optional] [default to null]
**Bulletins** | [**[]BulletinEntity**](BulletinEntity.md) | The bulletins for this component. | [optional] [default to null]
**Component** | [***ConnectionDto**](ConnectionDTO.md) |  | [optional] [default to null]
**DestinationGroupId** | **string** | The identifier of the group of the destination of this connection. | [optional] [default to null]
**DestinationId** | **string** | The identifier of the destination of this connection. | [optional] [default to null]
**DestinationType** | **string** | The type of component the destination connectable is. | [default to null]
**DisconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] [default to null]
**GetzIndex** | **int64** | The z index of the connection. | [optional] [default to null]
**Id** | **string** | The id of the component. | [optional] [default to null]
**LabelIndex** | **int32** | The index of the bend point where to place the connection label. | [optional] [default to null]
**Permissions** | [***PermissionsDto**](PermissionsDTO.md) |  | [optional] [default to null]
**Position** | [***PositionDto**](PositionDTO.md) |  | [optional] [default to null]
**Revision** | [***RevisionDto**](RevisionDTO.md) |  | [optional] [default to null]
**SourceGroupId** | **string** | The identifier of the group of the source of this connection. | [optional] [default to null]
**SourceId** | **string** | The identifier of the source of this connection. | [optional] [default to null]
**SourceType** | **string** | The type of component the source connectable is. | [default to null]
**Status** | [***ConnectionStatusDto**](ConnectionStatusDTO.md) |  | [optional] [default to null]
**Uri** | **string** | The URI for futures requests to the component. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

