# UserDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessPolicies** | [**[]AccessPolicySummaryEntity**](AccessPolicySummaryEntity.md) | The access policies this user belongs to. | [optional] [default to null]
**Configurable** | **bool** | Whether this tenant is configurable. | [optional] [default to null]
**Id** | **string** | The id of the component. | [optional] [default to null]
**Identity** | **string** | The identity of the tenant. | [optional] [default to null]
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] [default to null]
**Position** | [***PositionDto**](PositionDTO.md) |  | [optional] [default to null]
**UserGroups** | [**[]TenantEntity**](TenantEntity.md) | The groups to which the user belongs. This field is read only and it provided for convenience. | [optional] [default to null]
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

