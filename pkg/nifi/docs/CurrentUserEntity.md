# CurrentUserEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | **string** | The user identity being serialized. | [optional] [default to null]
**Anonymous** | **bool** | Whether the current user is anonymous. | [optional] [default to null]
**ProvenancePermissions** | [***PermissionsDto**](PermissionsDTO.md) |  | [optional] [default to null]
**CountersPermissions** | [***PermissionsDto**](PermissionsDTO.md) |  | [optional] [default to null]
**TenantsPermissions** | [***PermissionsDto**](PermissionsDTO.md) |  | [optional] [default to null]
**ControllerPermissions** | [***PermissionsDto**](PermissionsDTO.md) |  | [optional] [default to null]
**PoliciesPermissions** | [***PermissionsDto**](PermissionsDTO.md) |  | [optional] [default to null]
**SystemPermissions** | [***PermissionsDto**](PermissionsDTO.md) |  | [optional] [default to null]
**ParameterContextPermissions** | [***PermissionsDto**](PermissionsDTO.md) |  | [optional] [default to null]
**RestrictedComponentsPermissions** | [***PermissionsDto**](PermissionsDTO.md) |  | [optional] [default to null]
**ComponentRestrictionPermissions** | [**[]ComponentRestrictionPermissionDto**](ComponentRestrictionPermissionDTO.md) | Permissions for specific component restrictions. | [optional] [default to null]
**CanVersionFlows** | **bool** | Whether the current user can version flows. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

