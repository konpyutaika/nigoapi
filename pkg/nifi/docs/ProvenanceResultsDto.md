# ProvenanceResultsDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProvenanceEvents** | [**[]ProvenanceEventDto**](ProvenanceEventDTO.md) | The provenance events that matched the search criteria. | [optional] [default to null]
**Total** | **string** | The total number of results formatted. | [optional] [default to null]
**TotalCount** | **int64** | The total number of results. | [optional] [default to null]
**Generated** | **string** | Then the search was performed. | [optional] [default to null]
**OldestEvent** | **string** | The oldest event available in the provenance repository. | [optional] [default to null]
**TimeOffset** | **int32** | The time offset of the server that&#x27;s used for event time. | [optional] [default to null]
**Errors** | **[]string** | Any errors that occurred while performing the provenance request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

