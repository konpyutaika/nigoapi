# ComponentSearchResultDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component that matched the search. | [optional] [default to null]
**GroupId** | **string** | The group id of the component that matched the search. | [optional] [default to null]
**ParentGroup** | [***SearchResultGroupDto**](SearchResultGroupDTO.md) | The parent group of the component that matched the search. | [optional] [default to null]
**VersionedGroup** | [***SearchResultGroupDto**](SearchResultGroupDTO.md) | The nearest versioned ancestor group of the component that matched the search. | [optional] [default to null]
**Name** | **string** | The name of the component that matched the search. | [optional] [default to null]
**Matches** | **[]string** | What matched the search from the component. | [optional] [default to null]

[[Back to Model list]](../pkg/nifi/README.md#documentation-for-models) [[Back to API list]](../pkg/nifi/README.md#documentation-for-api-endpoints) [[Back to README]](../pkg/nifi/README.md)


