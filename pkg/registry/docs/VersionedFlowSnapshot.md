# VersionedFlowSnapshot

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | [***Bucket**](Bucket.md) |  | [optional] [default to null]
**ExternalControllerServices** | [**map[string]ExternalControllerServiceReference**](ExternalControllerServiceReference.md) | The information about controller services that exist outside this versioned flow, but are referenced by components within the versioned flow. | [optional] [default to null]
**Flow** | [***VersionedFlow**](VersionedFlow.md) |  | [optional] [default to null]
**FlowContents** | [***VersionedProcessGroup**](VersionedProcessGroup.md) |  | [default to null]
**FlowEncodingVersion** | **string** | The optional encoding version of the flow contents. | [optional] [default to null]
**Latest** | **bool** |  | [optional] [default to null]
**ParameterContexts** | [**map[string]VersionedParameterContext**](VersionedParameterContext.md) | The parameter contexts referenced by process groups in the flow contents. The mapping is from the name of the context to the context instance, and it is expected that any context in this map is referenced by at least one process group in this flow. | [optional] [default to null]
**ParameterProviders** | [**map[string]ParameterProviderReference**](ParameterProviderReference.md) | Contains basic information about parameter providers referenced in the versioned flow. | [optional] [default to null]
**SnapshotMetadata** | [***VersionedFlowSnapshotMetadata**](VersionedFlowSnapshotMetadata.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

