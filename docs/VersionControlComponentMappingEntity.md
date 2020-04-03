# VersionControlComponentMappingEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionControlComponentMapping** | **map[string]string** | The mapping of Versioned Component Identifiers to instance ID&#39;s | [optional] [default to null]
**ProcessGroupRevision** | [***RevisionDto**](RevisionDTO.md) | The revision of the Process Group | [optional] [default to null]
**DisconnectedNodeAcknowledged** | **bool** | Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [optional] [default to null]
**VersionControlInformation** | [***VersionControlInformationDto**](VersionControlInformationDTO.md) | The Version Control information | [optional] [default to null]

[[Back to Model list]](../pkg/nifi/README.md#documentation-for-models) [[Back to API list]](../pkg/nifi/README.md#documentation-for-api-endpoints) [[Back to README]](../pkg/nifi/README.md)


