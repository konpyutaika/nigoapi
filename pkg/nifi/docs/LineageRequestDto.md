# LineageRequestDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterNodeId** | **string** | The id of the node where this lineage originated if clustered. | [optional] [default to null]
**EventId** | **int64** | The event id that was used to generate this lineage, if applicable. The event id is allowed for any type of lineageRequestType. If the lineageRequestType is FLOWFILE and the flowfile uuid is also included in the request, the event id will be ignored.  | [optional] [default to null]
**LineageRequestType** | **string** | The type of lineage request. PARENTS will return the lineage for the flowfiles that are parents of the specified event. CHILDREN will return the lineage for the flowfiles that are children of the specified event. FLOWFILE will return the lineage for the specified flowfile. | [optional] [default to null]
**Uuid** | **string** | The flowfile uuid that was used to generate the lineage. The flowfile uuid is only allowed when the lineageRequestType is FLOWFILE and will take precedence over event id. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

