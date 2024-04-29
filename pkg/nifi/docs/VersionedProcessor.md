# VersionedProcessor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The component&#x27;s unique identifier | [optional] [default to null]
**InstanceIdentifier** | **string** | The instance ID of an existing component that is described by this VersionedComponent, or null if this is not mapped to an instantiated component | [optional] [default to null]
**Name** | **string** | The component&#x27;s name | [optional] [default to null]
**Comments** | **string** | The user-supplied comments for the component | [optional] [default to null]
**Position** | [***Position**](Position.md) |  | [optional] [default to null]
**Type_** | **string** | The type of the extension component | [optional] [default to null]
**Bundle** | [***Bundle**](Bundle.md) |  | [optional] [default to null]
**Properties** | **map[string]string** | The properties for the component. Properties whose value is not set will only contain the property name. | [optional] [default to null]
**PropertyDescriptors** | [**map[string]VersionedPropertyDescriptor**](VersionedPropertyDescriptor.md) | The property descriptors for the component. | [optional] [default to null]
**Style** | **map[string]string** | Stylistic data for rendering in a UI | [optional] [default to null]
**AnnotationData** | **string** | The annotation data for the processor used to relay configuration between a custom UI and the procesosr. | [optional] [default to null]
**SchedulingPeriod** | **string** | The frequency with which to schedule the processor. The format of the value will depend on th value of schedulingStrategy. | [optional] [default to null]
**SchedulingStrategy** | **string** | Indicates how the processor should be scheduled to run. | [optional] [default to null]
**ExecutionNode** | **string** | Indicates the node where the process will execute. | [optional] [default to null]
**PenaltyDuration** | **string** | The amout of time that is used when the process penalizes a flowfile. | [optional] [default to null]
**YieldDuration** | **string** | The amount of time that must elapse before this processor is scheduled again after yielding. | [optional] [default to null]
**BulletinLevel** | **string** | The level at which the processor will report bulletins. | [optional] [default to null]
**RunDurationMillis** | **int64** | The run duration for the processor in milliseconds. | [optional] [default to null]
**ConcurrentlySchedulableTaskCount** | **int32** | The number of tasks that should be concurrently schedule for the processor. If the processor doesn&#x27;t allow parallol processing then any positive input will be ignored. | [optional] [default to null]
**AutoTerminatedRelationships** | **[]string** | The names of all relationships that cause a flow file to be terminated if the relationship is not connected elsewhere. This property differs from the &#x27;isAutoTerminate&#x27; property of the RelationshipDTO in that the RelationshipDTO is meant to depict the current configuration, whereas this property can be set in a DTO when updating a Processor in order to change which Relationships should be auto-terminated. | [optional] [default to null]
**ScheduledState** | **string** | The scheduled state of the component | [optional] [default to null]
**RetryCount** | **int32** | Overall number of retries. | [optional] [default to null]
**RetriedRelationships** | **[]string** | All the relationships should be retried. | [optional] [default to null]
**BackoffMechanism** | **string** | Determines whether the FlowFile should be penalized or the processor should be yielded between retries. | [optional] [default to null]
**MaxBackoffPeriod** | **string** | Maximum amount of time to be waited during a retry period. | [optional] [default to null]
**ComponentType** | **string** |  | [optional] [default to null]
**GroupIdentifier** | **string** | The ID of the Process Group that this component belongs to | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

