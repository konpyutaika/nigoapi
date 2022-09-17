# JvmFlowDiagnosticsSnapshotDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uptime** | **string** | How long this node has been running, formatted as hours:minutes:seconds.milliseconds | [optional] [default to null]
**TimeZone** | **string** | The name of the Time Zone that is configured, if available | [optional] [default to null]
**ActiveTimerDrivenThreads** | **int32** | The number of timer-driven threads that are active | [optional] [default to null]
**ActiveEventDrivenThreads** | **int32** | The number of event-driven threads that are active | [optional] [default to null]
**BundlesLoaded** | [**[]BundleDto**](BundleDTO.md) | The NiFi Bundles (NARs) that are loaded by NiFi | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


