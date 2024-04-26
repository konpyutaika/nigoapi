# DropRequestDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id for this drop request. | [optional] [default to null]
**Uri** | **string** | The URI for future requests to this drop request. | [optional] [default to null]
**SubmissionTime** | **string** | The timestamp when the query was submitted. | [optional] [default to null]
**LastUpdated** | **string** | The last time this drop request was updated. | [optional] [default to null]
**PercentCompleted** | **int32** | The current percent complete. | [optional] [default to null]
**Finished** | **bool** | Whether the query has finished. | [optional] [default to null]
**FailureReason** | **string** | The reason, if any, that this drop request failed. | [optional] [default to null]
**CurrentCount** | **int32** | The number of flow files currently queued. | [optional] [default to null]
**CurrentSize** | **int64** | The size of flow files currently queued in bytes. | [optional] [default to null]
**Current** | **string** | The count and size of flow files currently queued. | [optional] [default to null]
**OriginalCount** | **int32** | The number of flow files to be dropped as a result of this request. | [optional] [default to null]
**OriginalSize** | **int64** | The size of flow files to be dropped as a result of this request in bytes. | [optional] [default to null]
**Original** | **string** | The count and size of flow files to be dropped as a result of this request. | [optional] [default to null]
**DroppedCount** | **int32** | The number of flow files that have been dropped thus far. | [optional] [default to null]
**DroppedSize** | **int64** | The size of flow files that have been dropped thus far in bytes. | [optional] [default to null]
**Dropped** | **string** | The count and size of flow files that have been dropped thus far. | [optional] [default to null]
**State** | **string** | The current state of the drop request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

