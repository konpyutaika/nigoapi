# ComponentStateDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentId** | **string** | The component identifier. | [optional] [default to null]
**StateDescription** | **string** | Description of the state this component persists. | [optional] [default to null]
**ClusterState** | [***StateMapDto**](StateMapDTO.md) | The cluster state for this component, or null if this NiFi is a standalone instance. | [optional] [default to null]
**LocalState** | [***StateMapDto**](StateMapDTO.md) | The local state for this component. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


