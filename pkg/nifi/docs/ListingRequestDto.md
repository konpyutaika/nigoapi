# ListingRequestDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id for this listing request. | [optional] [default to null]
**Uri** | **string** | The URI for future requests to this listing request. | [optional] [default to null]
**SubmissionTime** | **string** | The timestamp when the query was submitted. | [optional] [default to null]
**LastUpdated** | **string** | The last time this listing request was updated. | [optional] [default to null]
**PercentCompleted** | **int32** | The current percent complete. | [optional] [default to null]
**Finished** | **bool** | Whether the query has finished. | [optional] [default to null]
**FailureReason** | **string** | The reason, if any, that this listing request failed. | [optional] [default to null]
**MaxResults** | **int32** | The maximum number of FlowFileSummary objects to return | [optional] [default to null]
**State** | **string** | The current state of the listing request. | [optional] [default to null]
**QueueSize** | [***QueueSizeDto**](QueueSizeDTO.md) |  | [optional] [default to null]
**FlowFileSummaries** | [**[]FlowFileSummaryDto**](FlowFileSummaryDTO.md) | The FlowFile summaries. The summaries will be populated once the request has completed. | [optional] [default to null]
**SourceRunning** | **bool** | Whether the source of the connection is running | [optional] [default to null]
**DestinationRunning** | **bool** | Whether the destination of the connection is running | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

