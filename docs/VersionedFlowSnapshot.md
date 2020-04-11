# VersionedFlowSnapshot

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotMetadata** | [***VersionedFlowSnapshotMetadata**](VersionedFlowSnapshotMetadata.md) | The metadata for this snapshot | [default to null]
**FlowContents** | [***VersionedProcessGroup**](VersionedProcessGroup.md) | The contents of the versioned flow | [default to null]
**ExternalControllerServices** | [**map[string]ExternalControllerServiceReference**](ExternalControllerServiceReference.md) | The information about controller services that exist outside this versioned flow, but are referenced by components within the versioned flow. | [optional] [default to null]
**ParameterContexts** | [**map[string]VersionedParameterContext**](VersionedParameterContext.md) | The parameter contexts referenced by process groups in the flow contents. The mapping is from the name of the context to the context instance, and it is expected that any context in this map is referenced by at least one process group in this flow. | [optional] [default to null]
**FlowEncodingVersion** | **string** | The optional encoding version of the flow contents. | [optional] [default to null]
**Flow** | [***VersionedFlow**](VersionedFlow.md) | The flow this snapshot is for | [optional] [default to null]
**Bucket** | [***Bucket**](Bucket.md) | The bucket where the flow is located | [optional] [default to null]
**Latest** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


