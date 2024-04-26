# ConnectionStatusPredictionsSnapshotDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PredictedMillisUntilCountBackpressure** | **int64** | The predicted number of milliseconds before the connection will have backpressure applied, based on the queued count. | [optional] [default to null]
**PredictedMillisUntilBytesBackpressure** | **int64** | The predicted number of milliseconds before the connection will have backpressure applied, based on the total number of bytes in the queue. | [optional] [default to null]
**PredictionIntervalSeconds** | **int32** | The configured interval (in seconds) for predicting connection queue count and size (and percent usage). | [optional] [default to null]
**PredictedCountAtNextInterval** | **int32** | The predicted number of queued objects at the next configured interval. | [optional] [default to null]
**PredictedBytesAtNextInterval** | **int64** | The predicted total number of bytes in the queue at the next configured interval. | [optional] [default to null]
**PredictedPercentCount** | **int32** | Predicted connection percent use regarding queued flow files count and backpressure threshold if configured. | [optional] [default to null]
**PredictedPercentBytes** | **int32** | Predicted connection percent use regarding queued flow files size and backpressure threshold if configured. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

