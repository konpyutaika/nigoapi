# VersionedReportingTask

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnotationData** | **string** | The annotation for the reporting task. This is how the custom UI relays configuration to the reporting task. | [optional] [default to null]
**Bundle** | [***Bundle**](Bundle.md) |  | [optional] [default to null]
**Comments** | **string** | The user-supplied comments for the component | [optional] [default to null]
**ComponentType** | **string** |  | [optional] [default to null]
**GroupIdentifier** | **string** | The ID of the Process Group that this component belongs to | [optional] [default to null]
**Identifier** | **string** | The component&#x27;s unique identifier | [optional] [default to null]
**InstanceIdentifier** | **string** | The instance ID of an existing component that is described by this VersionedComponent, or null if this is not mapped to an instantiated component | [optional] [default to null]
**Name** | **string** | The component&#x27;s name | [optional] [default to null]
**Position** | [***Position**](Position.md) |  | [optional] [default to null]
**Properties** | **map[string]string** | The properties for the component. Properties whose value is not set will only contain the property name. | [optional] [default to null]
**PropertyDescriptors** | [**map[string]VersionedPropertyDescriptor**](VersionedPropertyDescriptor.md) | The property descriptors for the component. | [optional] [default to null]
**ScheduledState** | **string** | Indicates the scheduled state for the Reporting Task | [optional] [default to null]
**SchedulingPeriod** | **string** | The frequency with which to schedule the reporting task. The format of the value will depend on the value of schedulingStrategy. | [optional] [default to null]
**SchedulingStrategy** | **string** | Indicates scheduling strategy that should dictate how the reporting task is triggered. | [optional] [default to null]
**Type_** | **string** | The type of the extension component | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

