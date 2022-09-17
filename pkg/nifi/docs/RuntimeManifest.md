# RuntimeManifest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | A unique identifier for the manifest | [optional] [default to null]
**AgentType** | **string** | The type of the runtime binary, e.g., &#39;minifi-java&#39; or &#39;minifi-cpp&#39; | [optional] [default to null]
**Version** | **string** | The version of the runtime binary, e.g., &#39;1.0.1&#39; | [optional] [default to null]
**BuildInfo** | [***BuildInfo**](BuildInfo.md) | Build summary for this runtime binary | [optional] [default to null]
**Bundles** | [**[]Bundle**](Bundle.md) | All extension bundles included with this runtime | [optional] [default to null]
**SchedulingDefaults** | [***SchedulingDefaults**](SchedulingDefaults.md) | Scheduling defaults for components defined in this manifest | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


