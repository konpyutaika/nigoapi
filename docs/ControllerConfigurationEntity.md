# ControllerConfigurationEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | [***RevisionDto**](RevisionDTO.md) | The revision for this request/response. The revision is required for any mutable flow requests and is included in all responses. | [optional] [default to null]
**Permissions** | [***PermissionsDto**](PermissionsDTO.md) | The permissions for this component. | [optional] [default to null]
**DisconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] [default to null]
**Component** | [***ControllerConfigurationDto**](ControllerConfigurationDTO.md) | The controller configuration. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


