# UpdateControllerServiceReferenceRequestEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] [default to null]
**Id** | **string** | The identifier of the Controller Service. | [optional] [default to null]
**ReferencingComponentRevisions** | [**map[string]RevisionDto**](RevisionDTO.md) | The revisions for all referencing components. | [optional] [default to null]
**State** | **string** | The new state of the references for the controller service. | [optional] [default to null]
**UiOnly** | **bool** | Indicates whether or not the response should only include fields necessary for rendering the NiFi User Interface. As such, when this value is set to true, some fields may be returned as null values, and the selected fields may change at any time without notice. As a result, this value should not be set to true by any client other than the UI.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

