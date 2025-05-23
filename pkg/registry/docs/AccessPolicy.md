# AccessPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action associated with this access policy. | [optional] [default to null]
**Configurable** | **bool** | Indicates if this access policy is configurable, based on which Authorizer has been configured to manage it. | [optional] [default to null]
**Identifier** | **string** | The id of the policy. Set by server at creation time. | [optional] [default to null]
**Resource** | **string** | The resource for this access policy. | [optional] [default to null]
**Revision** | [***RevisionInfo**](RevisionInfo.md) |  | [optional] [default to null]
**UserGroups** | [**[]Tenant**](Tenant.md) | The set of user group IDs associated with this access policy. | [optional] [default to null]
**Users** | [**[]Tenant**](Tenant.md) | The set of user IDs associated with this access policy. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

