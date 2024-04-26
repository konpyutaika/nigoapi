# ConnectionDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] [default to null]
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] [default to null]
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] [default to null]
**Position** | [***PositionDto**](PositionDTO.md) |  | [optional] [default to null]
**Source** | [***ConnectableDto**](ConnectableDTO.md) |  | [optional] [default to null]
**Destination** | [***ConnectableDto**](ConnectableDTO.md) |  | [optional] [default to null]
**Name** | **string** | The name of the connection. | [optional] [default to null]
**LabelIndex** | **int32** | The index of the bend point where to place the connection label. | [optional] [default to null]
**GetzIndex** | **int64** | The z index of the connection. | [optional] [default to null]
**SelectedRelationships** | **[]string** | The selected relationship that comprise the connection. | [optional] [default to null]
**AvailableRelationships** | **[]string** | The relationships that the source of the connection currently supports. | [optional] [default to null]
**BackPressureObjectThreshold** | **int64** | The object count threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won&#x27;t impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue. | [optional] [default to null]
**BackPressureDataSizeThreshold** | **string** | The object data size threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won&#x27;t impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue. | [optional] [default to null]
**FlowFileExpiration** | **string** | The amount of time a flow file may be in the flow before it will be automatically aged out of the flow. Once a flow file reaches this age it will be terminated from the flow the next time a processor attempts to start work on it. | [optional] [default to null]
**Prioritizers** | **[]string** | The comparators used to prioritize the queue. | [optional] [default to null]
**Bends** | [**[]PositionDto**](PositionDTO.md) | The bend points on the connection. | [optional] [default to null]
**LoadBalanceStrategy** | **string** | How to load balance the data in this Connection across the nodes in the cluster. | [optional] [default to null]
**LoadBalancePartitionAttribute** | **string** | The FlowFile Attribute to use for determining which node a FlowFile will go to if the Load Balancing Strategy is set to PARTITION_BY_ATTRIBUTE | [optional] [default to null]
**LoadBalanceCompression** | **string** | Whether or not data should be compressed when being transferred between nodes in the cluster. | [optional] [default to null]
**LoadBalanceStatus** | **string** | The current status of the Connection&#x27;s Load Balancing Activities. Status can indicate that Load Balancing is not configured for the connection, that Load Balancing is configured but inactive (not currently transferring data to another node), or that Load Balancing is configured and actively transferring data to another node. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

