# ReplayLastEventSnapshotDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventAvailable** | **bool** | Whether or not an event was available. This may not be populated if there was a failure. | [optional] [default to null]
**EventsReplayed** | **[]int64** | The IDs of the events that were successfully replayed | [optional] [default to null]
**FailureExplanation** | **string** | If unable to replay an event, specifies why the event could not be replayed | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

