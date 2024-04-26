# VersionedFlowUpdateRequestDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The unique ID of this request. | [optional] [default to null]
**ProcessGroupId** | **string** | The unique ID of the Process Group being updated | [optional] [default to null]
**Uri** | **string** | The URI for future requests to this drop request. | [optional] [default to null]
**LastUpdated** | **string** | The last time this request was updated. | [optional] [default to null]
**Complete** | **bool** | Whether or not this request has completed | [optional] [default to null]
**FailureReason** | **string** | An explanation of why this request failed, or null if this request has not failed | [optional] [default to null]
**PercentCompleted** | **int32** | The percentage complete for the request, between 0 and 100 | [optional] [default to null]
**State** | **string** | The state of the request | [optional] [default to null]
**VersionControlInformation** | [***VersionControlInformationDto**](VersionControlInformationDTO.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

