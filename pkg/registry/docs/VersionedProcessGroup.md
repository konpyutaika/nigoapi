# VersionedProcessGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | **string** | The user-supplied comments for the component | [optional] [default to null]
**ComponentType** | **string** |  | [optional] [default to null]
**Connections** | [**[]VersionedConnection**](VersionedConnection.md) | The Connections | [optional] [default to null]
**ControllerServices** | [**[]VersionedControllerService**](VersionedControllerService.md) | The Controller Services | [optional] [default to null]
**DefaultBackPressureDataSizeThreshold** | **string** | Default value used in this Process Group for the maximum data size of objects that can be queued before back pressure is applied. | [optional] [default to null]
**DefaultBackPressureObjectThreshold** | **int64** | Default value used in this Process Group for the maximum number of objects that can be queued before back pressure is applied. | [optional] [default to null]
**DefaultFlowFileExpiration** | **string** | The default FlowFile Expiration for this Process Group. | [optional] [default to null]
**ExecutionEngine** | **string** | The Execution Engine that should be used to run the components within the group. | [optional] [default to null]
**FlowFileConcurrency** | **string** | The configured FlowFile Concurrency for the Process Group | [optional] [default to null]
**FlowFileOutboundPolicy** | **string** | The FlowFile Outbound Policy for the Process Group | [optional] [default to null]
**Funnels** | [**[]VersionedFunnel**](VersionedFunnel.md) | The Funnels | [optional] [default to null]
**GroupIdentifier** | **string** | The ID of the Process Group that this component belongs to | [optional] [default to null]
**Identifier** | **string** | The component&#x27;s unique identifier | [optional] [default to null]
**InputPorts** | [**[]VersionedPort**](VersionedPort.md) | The Input Ports | [optional] [default to null]
**InstanceIdentifier** | **string** | The instance ID of an existing component that is described by this VersionedComponent, or null if this is not mapped to an instantiated component | [optional] [default to null]
**Labels** | [**[]VersionedLabel**](VersionedLabel.md) | The Labels | [optional] [default to null]
**LogFileSuffix** | **string** | The log file suffix for this Process Group for dedicated logging. | [optional] [default to null]
**MaxConcurrentTasks** | **int32** | The maximum number of concurrent tasks that should be scheduled for this Process Group when using the Stateless Engine | [optional] [default to null]
**Name** | **string** | The component&#x27;s name | [optional] [default to null]
**OutputPorts** | [**[]VersionedPort**](VersionedPort.md) | The Output Ports | [optional] [default to null]
**ParameterContextName** | **string** | The name of the parameter context used by this process group | [optional] [default to null]
**Position** | [***Position**](Position.md) |  | [optional] [default to null]
**ProcessGroups** | [**[]VersionedProcessGroup**](VersionedProcessGroup.md) | The child Process Groups | [optional] [default to null]
**Processors** | [**[]VersionedProcessor**](VersionedProcessor.md) | The Processors | [optional] [default to null]
**RemoteProcessGroups** | [**[]VersionedRemoteProcessGroup**](VersionedRemoteProcessGroup.md) | The Remote Process Groups | [optional] [default to null]
**ScheduledState** | **string** | The Scheduled State of the Process Group, if the group is configured to use the Stateless Execution Engine. Otherwise, this value has no relevance. | [optional] [default to null]
**StatelessFlowTimeout** | **string** | The maximum amount of time that the flow is allows to run using the Stateless engine before it times out and is considered a failure | [optional] [default to null]
**VersionedFlowCoordinates** | [***VersionedFlowCoordinates**](VersionedFlowCoordinates.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

