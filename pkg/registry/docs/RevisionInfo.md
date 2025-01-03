# RevisionInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | A client identifier used to make a request. By including a client identifier, the API can allow multiple requests without needing the current revision. Due to the asynchronous nature of requests/responses this was implemented to allow the client to make numerous requests without having to wait for the previous response to come back. | [optional] [default to null]
**LastModifier** | **string** | The user that last modified the entity. | [optional] [default to null]
**Version** | **int64** | NiFi Registry employs an optimistic locking strategy where the client must include a revision in their request when performing an update. In a response to a mutable flow request, this field represents the updated base version. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

