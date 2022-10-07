# ParameterProviderParameterApplicationEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the parameter provider. | [optional] [default to null]
**Revision** | [***RevisionDto**](RevisionDTO.md) | The revision for this request/response. The revision is required for any mutable flow requests and is included in all responses. | [optional] [default to null]
**DisconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] [default to null]
**ParameterGroupConfigurations** | [**[]ParameterGroupConfigurationEntity**](ParameterGroupConfigurationEntity.md) | Configuration for the fetched Parameter Groups | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


