# ProvenanceRequestDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchTerms** | [**map[string]ProvenanceSearchValueDto**](ProvenanceSearchValueDTO.md) | The search terms used to perform the search. | [optional] [default to null]
**ClusterNodeId** | **string** | The id of the node in the cluster where this provenance originated. | [optional] [default to null]
**StartDate** | **string** | The earliest event time to include in the query. | [optional] [default to null]
**EndDate** | **string** | The latest event time to include in the query. | [optional] [default to null]
**MinimumFileSize** | **string** | The minimum file size to include in the query. | [optional] [default to null]
**MaximumFileSize** | **string** | The maximum file size to include in the query. | [optional] [default to null]
**MaxResults** | **int32** | The maximum number of results to include. | [optional] [default to null]
**Summarize** | **bool** | Whether or not to summarize provenance events returned. This property is false by default. | [optional] [default to null]
**IncrementalResults** | **bool** | Whether or not incremental results are returned. If false, provenance events are only returned once the query completes. This property is true by default. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

