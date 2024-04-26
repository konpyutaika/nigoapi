# User

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The computer-generated identifier of the tenant. | [optional] [default to null]
**Identity** | **string** | The human-facing identity of the tenant. This can only be changed if the tenant is configurable. | [optional] [default to null]
**Configurable** | **bool** | Indicates if this tenant is configurable, based on which UserGroupProvider has been configured to manage it. | [optional] [default to null]
**ResourcePermissions** | [***ResourcePermissions**](ResourcePermissions.md) |  | [optional] [default to null]
**AccessPolicies** | [**[]AccessPolicySummary**](AccessPolicySummary.md) | The access policies granted to this tenant. | [optional] [default to null]
**Revision** | [***RevisionInfo**](RevisionInfo.md) |  | [optional] [default to null]
**UserGroups** | [**[]Tenant**](Tenant.md) | The groups to which the user belongs. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

