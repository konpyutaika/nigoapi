# UserGroupDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] [default to null]
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] [default to null]
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] [default to null]
**Position** | [***PositionDto**](PositionDTO.md) |  | [optional] [default to null]
**Identity** | **string** | The identity of the tenant. | [optional] [default to null]
**Configurable** | **bool** | Whether this tenant is configurable. | [optional] [default to null]
**Users** | [**[]TenantEntity**](TenantEntity.md) | The users that belong to the user group. | [optional] [default to null]
**AccessPolicies** | [**[]AccessPolicyEntity**](AccessPolicyEntity.md) | The access policies this user group belongs to. This field was incorrectly defined as an AccessPolicyEntity. For compatibility reasons the field will remain of this type, however only the fields that are present in the AccessPolicySummaryEntity will be populated here. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

