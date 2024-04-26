# ClusterSummaryDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectedNodes** | **string** | When clustered, reports the number of nodes connected vs the number of nodes in the cluster. | [optional] [default to null]
**ConnectedNodeCount** | **int32** | The number of nodes that are currently connected to the cluster | [optional] [default to null]
**TotalNodeCount** | **int32** | The number of nodes in the cluster, regardless of whether or not they are connected | [optional] [default to null]
**Clustered** | **bool** | Whether this NiFi instance is clustered. | [optional] [default to null]
**ConnectedToCluster** | **bool** | Whether this NiFi instance is connected to a cluster. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

