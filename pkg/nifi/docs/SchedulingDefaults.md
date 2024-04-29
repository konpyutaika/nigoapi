# SchedulingDefaults

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSchedulingStrategy** | **string** | The name of the default scheduling strategy | [optional] [default to null]
**DefaultSchedulingPeriodMillis** | **int64** | The default scheduling period in milliseconds | [optional] [default to null]
**PenalizationPeriodMillis** | **int64** | The default penalization period in milliseconds | [optional] [default to null]
**YieldDurationMillis** | **int64** | The default yield duration in milliseconds | [optional] [default to null]
**DefaultRunDurationNanos** | **int64** | The default run duration in nano-seconds | [optional] [default to null]
**DefaultMaxConcurrentTasks** | **string** | The default concurrent tasks | [optional] [default to null]
**DefaultConcurrentTasksBySchedulingStrategy** | **map[string]int32** | The default concurrent tasks for each scheduling strategy | [optional] [default to null]
**DefaultSchedulingPeriodsBySchedulingStrategy** | **map[string]string** | The default scheduling period for each scheduling strategy | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

