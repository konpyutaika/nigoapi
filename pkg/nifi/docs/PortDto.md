# PortDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] [default to null]
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] [default to null]
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] [default to null]
**Position** | [***PositionDto**](PositionDTO.md) | The position of this component in the UI if applicable. | [optional] [default to null]
**Name** | **string** | The name of the port. | [optional] [default to null]
**Comments** | **string** | The comments for the port. | [optional] [default to null]
**State** | **string** | The state of the port. | [optional] [default to null]
**Type_** | **string** | The type of port. | [optional] [default to null]
**Transmitting** | **bool** | Whether the port has incoming or output connections to a remote NiFi. This is only applicable when the port is allowed to be accessed remotely. | [optional] [default to null]
**ConcurrentlySchedulableTaskCount** | **int32** | The number of tasks that should be concurrently scheduled for the port. | [optional] [default to null]
**UserAccessControl** | **[]string** | The users that are allowed to access the port. | [optional] [default to null]
**GroupAccessControl** | **[]string** | The user groups that are allowed to access the port. | [optional] [default to null]
**AllowRemoteAccess** | **bool** | Whether this port can be accessed remotely via Site-to-Site protocol. | [optional] [default to null]
**ValidationErrors** | **[]string** | Gets the validation errors from this port. These validation errors represent the problems with the port that must be resolved before it can be started. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


