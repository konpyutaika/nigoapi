# ConnectionStatisticsSnapshotDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the connection. | [optional] [default to null]
**PredictedMillisUntilCountBackpressure** | **int64** | The predicted number of milliseconds before the connection will have backpressure applied, based on the queued count. | [optional] [default to null]
**PredictedMillisUntilBytesBackpressure** | **int64** | The predicted number of milliseconds before the connection will have backpressure applied, based on the total number of bytes in the queue. | [optional] [default to null]
**PredictedCountAtNextInterval** | **int32** | The predicted number of queued objects at the next configured interval. | [optional] [default to null]
**PredictedBytesAtNextInterval** | **int64** | The predicted total number of bytes in the queue at the next configured interval. | [optional] [default to null]
**PredictedPercentCount** | **int32** | The predicted percentage of queued objects at the next configured interval. | [optional] [default to null]
**PredictedPercentBytes** | **int32** | The predicted percentage of bytes in the queue against current threshold at the next configured interval. | [optional] [default to null]
**PredictionIntervalMillis** | **int64** | The prediction interval in seconds | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

