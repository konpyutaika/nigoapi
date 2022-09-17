# LocalQueuePartitionDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalFlowFileCount** | **int32** | Total number of FlowFiles owned by the Connection | [optional] [default to null]
**TotalByteCount** | **int64** | Total number of bytes that make up the content for the FlowFiles owned by this Connection | [optional] [default to null]
**ActiveQueueFlowFileCount** | **int32** | Total number of FlowFiles that exist in the Connection&#39;s Active Queue, immediately available to be offered up to a component | [optional] [default to null]
**ActiveQueueByteCount** | **int64** | Total number of bytes that make up the content for the FlowFiles that are present in the Connection&#39;s Active Queue | [optional] [default to null]
**SwapFlowFileCount** | **int32** | The total number of FlowFiles that are swapped out for this Connection | [optional] [default to null]
**SwapByteCount** | **int64** | Total number of bytes that make up the content for the FlowFiles that are swapped out to disk for the Connection | [optional] [default to null]
**SwapFiles** | **int32** | The number of Swap Files that exist for this Connection | [optional] [default to null]
**InFlightFlowFileCount** | **int32** | The number of In-Flight FlowFiles for this Connection. These are FlowFiles that belong to the connection but are currently being operated on by a Processor, Port, etc. | [optional] [default to null]
**InFlightByteCount** | **int64** | The number bytes that make up the content of the FlowFiles that are In-Flight | [optional] [default to null]
**AllActiveQueueFlowFilesPenalized** | **bool** | Whether or not all of the FlowFiles in the Active Queue are penalized | [optional] [default to null]
**AnyActiveQueueFlowFilesPenalized** | **bool** | Whether or not any of the FlowFiles in the Active Queue are penalized | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


