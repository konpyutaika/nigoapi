# VerifyConfigRequestDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The ID of the request | [optional] [default to null]
**Uri** | **string** | The URI for the request | [optional] [default to null]
**SubmissionTime** | [**time.Time**](time.Time.md) | The timestamp of when the request was submitted | [optional] [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) | The timestamp of when the request was last updated | [optional] [default to null]
**Complete** | **bool** | Whether or not the request is completed | [optional] [default to null]
**FailureReason** | **string** | The reason for the request failing, or null if the request has not failed | [optional] [default to null]
**PercentCompleted** | **int32** | A value between 0 and 100 (inclusive) indicating how close the request is to completion | [optional] [default to null]
**State** | **string** | A description of the current state of the request | [optional] [default to null]
**UpdateSteps** | [**[]VerifyConfigUpdateStepDto**](VerifyConfigUpdateStepDTO.md) | The steps that are required in order to complete the request, along with the status of each | [optional] [default to null]
**ComponentId** | **string** | The ID of the component whose configuration was verified | [optional] [default to null]
**Properties** | **map[string]string** | The configured component properties | [optional] [default to null]
**Attributes** | **map[string]string** | FlowFile Attributes that should be used to evaluate Expression Language for resolving property values | [optional] [default to null]
**Results** | [**[]ConfigVerificationResultDto**](ConfigVerificationResultDTO.md) | The Results of the verification | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

