# UpdateControllerServiceReferenceRequestEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier of the Controller Service. | [optional] [default to null]
**State** | **string** | The new state of the references for the controller service. | [optional] [default to null]
**ReferencingComponentRevisions** | [**map[string]RevisionDto**](RevisionDTO.md) | The revisions for all referencing components. | [optional] [default to null]
**DisconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


