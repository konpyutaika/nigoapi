# SnippetDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the snippet. | [optional] [default to null]
**Uri** | **string** | The URI of the snippet. | [optional] [default to null]
**ParentGroupId** | **string** | The group id for the components in the snippet. | [optional] [default to null]
**ProcessGroups** | [**map[string]RevisionDto**](RevisionDTO.md) | The ids of the process groups in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests). | [optional] [default to null]
**RemoteProcessGroups** | [**map[string]RevisionDto**](RevisionDTO.md) | The ids of the remote process groups in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests). | [optional] [default to null]
**Processors** | [**map[string]RevisionDto**](RevisionDTO.md) | The ids of the processors in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests). | [optional] [default to null]
**InputPorts** | [**map[string]RevisionDto**](RevisionDTO.md) | The ids of the input ports in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests). | [optional] [default to null]
**OutputPorts** | [**map[string]RevisionDto**](RevisionDTO.md) | The ids of the output ports in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests). | [optional] [default to null]
**Connections** | [**map[string]RevisionDto**](RevisionDTO.md) | The ids of the connections in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests). | [optional] [default to null]
**Labels** | [**map[string]RevisionDto**](RevisionDTO.md) | The ids of the labels in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests). | [optional] [default to null]
**Funnels** | [**map[string]RevisionDto**](RevisionDTO.md) | The ids of the funnels in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


