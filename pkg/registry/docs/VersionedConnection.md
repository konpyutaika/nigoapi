# VersionedConnection

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The component&#x27;s unique identifier | [optional] [default to null]
**InstanceIdentifier** | **string** | The instance ID of an existing component that is described by this VersionedComponent, or null if this is not mapped to an instantiated component | [optional] [default to null]
**Name** | **string** | The component&#x27;s name | [optional] [default to null]
**Comments** | **string** | The user-supplied comments for the component | [optional] [default to null]
**Position** | [***Position**](Position.md) |  | [optional] [default to null]
**Source** | [***ConnectableComponent**](ConnectableComponent.md) |  | [optional] [default to null]
**Destination** | [***ConnectableComponent**](ConnectableComponent.md) |  | [optional] [default to null]
**LabelIndex** | **int32** | The index of the bend point where to place the connection label. | [optional] [default to null]
**ZIndex** | **int64** | The z index of the connection. | [optional] [default to null]
**SelectedRelationships** | **[]string** | The selected relationship that comprise the connection. | [optional] [default to null]
**BackPressureObjectThreshold** | **int64** | The object count threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won&#x27;t impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue. | [optional] [default to null]
**BackPressureDataSizeThreshold** | **string** | The object data size threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won&#x27;t impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue. | [optional] [default to null]
**FlowFileExpiration** | **string** | The amount of time a flow file may be in the flow before it will be automatically aged out of the flow. Once a flow file reaches this age it will be terminated from the flow the next time a processor attempts to start work on it. | [optional] [default to null]
**Prioritizers** | **[]string** | The comparators used to prioritize the queue. | [optional] [default to null]
**Bends** | [**[]Position**](Position.md) | The bend points on the connection. | [optional] [default to null]
**LoadBalanceStrategy** | **string** | The Strategy to use for load balancing data across the cluster, or null, if no Load Balance Strategy has been specified. | [optional] [default to null]
**PartitioningAttribute** | **string** | The attribute to use for partitioning data as it is load balanced across the cluster. If the Load Balance Strategy is configured to use PARTITION_BY_ATTRIBUTE, the value returned by this method is the name of the FlowFile Attribute that will be used to determine which node in the cluster should receive a given FlowFile. If the Load Balance Strategy is unset or is set to any other value, the Partitioning Attribute has no effect. | [optional] [default to null]
**LoadBalanceCompression** | **string** | Whether or not compression should be used when transferring FlowFiles between nodes | [optional] [default to null]
**ComponentType** | **string** |  | [optional] [default to null]
**GroupIdentifier** | **string** | The ID of the Process Group that this component belongs to | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

