# AccessPolicyDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] [default to null]
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] [default to null]
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] [default to null]
**Position** | [***PositionDto**](PositionDTO.md) |  | [optional] [default to null]
**Resource** | **string** | The resource for this access policy. | [optional] [default to null]
**Action** | **string** | The action associated with this access policy. | [optional] [default to null]
**ComponentReference** | [***ComponentReferenceEntity**](ComponentReferenceEntity.md) |  | [optional] [default to null]
**Configurable** | **bool** | Whether this policy is configurable. | [optional] [default to null]
**Users** | [**[]TenantEntity**](TenantEntity.md) | The set of user IDs associated with this access policy. | [optional] [default to null]
**UserGroups** | [**[]TenantEntity**](TenantEntity.md) | The set of user group IDs associated with this access policy. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

