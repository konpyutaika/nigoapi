# ProvenanceDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the provenance query. | [optional] [default to null]
**Uri** | **string** | The URI for this query. Used for obtaining/deleting the request at a later time | [optional] [default to null]
**SubmissionTime** | **string** | The timestamp when the query was submitted. | [optional] [default to null]
**Expiration** | **string** | The timestamp when the query will expire. | [optional] [default to null]
**PercentCompleted** | **int32** | The current percent complete. | [optional] [default to null]
**Finished** | **bool** | Whether the query has finished. | [optional] [default to null]
**Request** | [***ProvenanceRequestDto**](ProvenanceRequestDTO.md) |  | [optional] [default to null]
**Results** | [***ProvenanceResultsDto**](ProvenanceResultsDTO.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

