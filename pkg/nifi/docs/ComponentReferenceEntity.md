# ComponentReferenceEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | [***RevisionDto**](RevisionDTO.md) |  | [optional] [default to null]
**Id** | **string** | The id of the component. | [optional] [default to null]
**Uri** | **string** | The URI for futures requests to the component. | [optional] [default to null]
**Position** | [***PositionDto**](PositionDTO.md) |  | [optional] [default to null]
**Permissions** | [***PermissionsDto**](PermissionsDTO.md) |  | [optional] [default to null]
**Bulletins** | [**[]BulletinEntity**](BulletinEntity.md) | The bulletins for this component. | [optional] [default to null]
**DisconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] [default to null]
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] [default to null]
**Component** | [***ComponentReferenceDto**](ComponentReferenceDTO.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

