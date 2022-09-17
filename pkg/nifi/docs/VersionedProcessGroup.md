# VersionedProcessGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The component&#39;s unique identifier | [optional] [default to null]
**Name** | **string** | The component&#39;s name | [optional] [default to null]
**Comments** | **string** | The user-supplied comments for the component | [optional] [default to null]
**Position** | [***Position**](Position.md) | The component&#39;s position on the graph | [optional] [default to null]
**ProcessGroups** | [**[]VersionedProcessGroup**](VersionedProcessGroup.md) | The child Process Groups | [optional] [default to null]
**RemoteProcessGroups** | [**[]VersionedRemoteProcessGroup**](VersionedRemoteProcessGroup.md) | The Remote Process Groups | [optional] [default to null]
**Processors** | [**[]VersionedProcessor**](VersionedProcessor.md) | The Processors | [optional] [default to null]
**InputPorts** | [**[]VersionedPort**](VersionedPort.md) | The Input Ports | [optional] [default to null]
**OutputPorts** | [**[]VersionedPort**](VersionedPort.md) | The Output Ports | [optional] [default to null]
**Connections** | [**[]VersionedConnection**](VersionedConnection.md) | The Connections | [optional] [default to null]
**Labels** | [**[]VersionedLabel**](VersionedLabel.md) | The Labels | [optional] [default to null]
**Funnels** | [**[]VersionedFunnel**](VersionedFunnel.md) | The Funnels | [optional] [default to null]
**ControllerServices** | [**[]VersionedControllerService**](VersionedControllerService.md) | The Controller Services | [optional] [default to null]
**VersionedFlowCoordinates** | [***VersionedFlowCoordinates**](VersionedFlowCoordinates.md) | The coordinates where the remote flow is stored, or null if the Process Group is not directly under Version Control | [optional] [default to null]
**Variables** | **map[string]string** | The Variables in the Variable Registry for this Process Group (not including any ancestor or descendant Process Groups) | [optional] [default to null]
**ParameterContextName** | **string** | The name of the parameter context used by this process group | [optional] [default to null]
**FlowFileConcurrency** | **string** | The configured FlowFile Concurrency for the Process Group | [optional] [default to null]
**FlowFileOutboundPolicy** | **string** | The FlowFile Outbound Policy for the Process Group | [optional] [default to null]
**ComponentType** | **string** |  | [optional] [default to null]
**GroupIdentifier** | **string** | The ID of the Process Group that this component belongs to | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


