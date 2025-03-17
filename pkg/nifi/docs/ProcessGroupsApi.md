# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Copy**](ProcessGroupsApi.md#Copy) | **Post** /process-groups/{id}/copy | Generates a copy response for the given copy request
[**CopySnippet**](ProcessGroupsApi.md#CopySnippet) | **Post** /process-groups/{id}/snippet-instance | Copies a snippet and discards it.
[**CreateConnection**](ProcessGroupsApi.md#CreateConnection) | **Post** /process-groups/{id}/connections | Creates a connection
[**CreateControllerService1**](ProcessGroupsApi.md#CreateControllerService1) | **Post** /process-groups/{id}/controller-services | Creates a new controller service
[**CreateEmptyAllConnectionsRequest**](ProcessGroupsApi.md#CreateEmptyAllConnectionsRequest) | **Post** /process-groups/{id}/empty-all-connections-requests | Creates a request to drop all flowfiles of all connection queues in this process group.
[**CreateFunnel**](ProcessGroupsApi.md#CreateFunnel) | **Post** /process-groups/{id}/funnels | Creates a funnel
[**CreateInputPort**](ProcessGroupsApi.md#CreateInputPort) | **Post** /process-groups/{id}/input-ports | Creates an input port
[**CreateLabel**](ProcessGroupsApi.md#CreateLabel) | **Post** /process-groups/{id}/labels | Creates a label
[**CreateOutputPort**](ProcessGroupsApi.md#CreateOutputPort) | **Post** /process-groups/{id}/output-ports | Creates an output port
[**CreateProcessGroup**](ProcessGroupsApi.md#CreateProcessGroup) | **Post** /process-groups/{id}/process-groups | Creates a process group
[**CreateProcessor**](ProcessGroupsApi.md#CreateProcessor) | **Post** /process-groups/{id}/processors | Creates a new processor
[**CreateRemoteProcessGroup**](ProcessGroupsApi.md#CreateRemoteProcessGroup) | **Post** /process-groups/{id}/remote-process-groups | Creates a new process group
[**DeleteReplaceProcessGroupRequest**](ProcessGroupsApi.md#DeleteReplaceProcessGroupRequest) | **Delete** /process-groups/replace-requests/{id} | Deletes the Replace Request with the given ID
[**ExportProcessGroup**](ProcessGroupsApi.md#ExportProcessGroup) | **Get** /process-groups/{id}/download | Gets a process group for download
[**GetConnections**](ProcessGroupsApi.md#GetConnections) | **Get** /process-groups/{id}/connections | Gets all connections
[**GetDropAllFlowfilesRequest**](ProcessGroupsApi.md#GetDropAllFlowfilesRequest) | **Get** /process-groups/{id}/empty-all-connections-requests/{drop-request-id} | Gets the current status of a drop all flowfiles request.
[**GetFunnels**](ProcessGroupsApi.md#GetFunnels) | **Get** /process-groups/{id}/funnels | Gets all funnels
[**GetInputPorts**](ProcessGroupsApi.md#GetInputPorts) | **Get** /process-groups/{id}/input-ports | Gets all input ports
[**GetLabels**](ProcessGroupsApi.md#GetLabels) | **Get** /process-groups/{id}/labels | Gets all labels
[**GetLocalModifications**](ProcessGroupsApi.md#GetLocalModifications) | **Get** /process-groups/{id}/local-modifications | Gets a list of local modifications to the Process Group since it was last synchronized with the Flow Registry
[**GetOutputPorts**](ProcessGroupsApi.md#GetOutputPorts) | **Get** /process-groups/{id}/output-ports | Gets all output ports
[**GetProcessGroup**](ProcessGroupsApi.md#GetProcessGroup) | **Get** /process-groups/{id} | Gets a process group
[**GetProcessGroups**](ProcessGroupsApi.md#GetProcessGroups) | **Get** /process-groups/{id}/process-groups | Gets all process groups
[**GetProcessors**](ProcessGroupsApi.md#GetProcessors) | **Get** /process-groups/{id}/processors | Gets all processors
[**GetRemoteProcessGroups**](ProcessGroupsApi.md#GetRemoteProcessGroups) | **Get** /process-groups/{id}/remote-process-groups | Gets all remote process groups
[**GetReplaceProcessGroupRequest**](ProcessGroupsApi.md#GetReplaceProcessGroupRequest) | **Get** /process-groups/replace-requests/{id} | Returns the Replace Request with the given ID
[**ImportProcessGroup**](ProcessGroupsApi.md#ImportProcessGroup) | **Post** /process-groups/{id}/process-groups/import | Imports a specified process group
[**InitiateReplaceProcessGroup**](ProcessGroupsApi.md#InitiateReplaceProcessGroup) | **Post** /process-groups/{id}/replace-requests | Initiate the Replace Request of a Process Group with the given ID
[**Paste**](ProcessGroupsApi.md#Paste) | **Put** /process-groups/{id}/paste | Pastes into the specified process group
[**RemoveDropRequest1**](ProcessGroupsApi.md#RemoveDropRequest1) | **Delete** /process-groups/{id}/empty-all-connections-requests/{drop-request-id} | Cancels and/or removes a request to drop all flowfiles.
[**RemoveProcessGroup**](ProcessGroupsApi.md#RemoveProcessGroup) | **Delete** /process-groups/{id} | Deletes a process group
[**ReplaceProcessGroup**](ProcessGroupsApi.md#ReplaceProcessGroup) | **Put** /process-groups/{id}/flow-contents | Replace Process Group contents with the given ID with the specified Process Group contents
[**UpdateProcessGroup**](ProcessGroupsApi.md#UpdateProcessGroup) | **Put** /process-groups/{id} | Updates a process group
[**UploadProcessGroup**](ProcessGroupsApi.md#UploadProcessGroup) | **Post** /process-groups/{id}/process-groups/upload | Uploads a versioned flow definition and creates a process group

# **Copy**
> CopyResponseEntity Copy(ctx, body, id)
Generates a copy response for the given copy request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CopyRequestEntity**](CopyRequestEntity.md)| The request including the components to be copied from the specified Process Group. | 
  **id** | **string**| The process group id. | 

### Return type

[**CopyResponseEntity**](CopyResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CopySnippet**
> FlowEntity CopySnippet(ctx, body, id)
Copies a snippet and discards it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CopySnippetRequestEntity**](CopySnippetRequestEntity.md)| The copy snippet request. | 
  **id** | **string**| The process group id. | 

### Return type

[**FlowEntity**](FlowEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConnection**
> ConnectionEntity CreateConnection(ctx, body, id)
Creates a connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectionEntity**](ConnectionEntity.md)| The connection configuration details. | 
  **id** | **string**| The process group id. | 

### Return type

[**ConnectionEntity**](ConnectionEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateControllerService1**
> ControllerServiceEntity CreateControllerService1(ctx, body, id)
Creates a new controller service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ControllerServiceEntity**](ControllerServiceEntity.md)| The controller service configuration details. | 
  **id** | **string**| The process group id. | 

### Return type

[**ControllerServiceEntity**](ControllerServiceEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEmptyAllConnectionsRequest**
> DropRequestEntity CreateEmptyAllConnectionsRequest(ctx, id)
Creates a request to drop all flowfiles of all connection queues in this process group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 

### Return type

[**DropRequestEntity**](DropRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFunnel**
> FunnelEntity CreateFunnel(ctx, body, id)
Creates a funnel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FunnelEntity**](FunnelEntity.md)| The funnel configuration details. | 
  **id** | **string**| The process group id. | 

### Return type

[**FunnelEntity**](FunnelEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateInputPort**
> PortEntity CreateInputPort(ctx, body, id)
Creates an input port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PortEntity**](PortEntity.md)| The input port configuration details. | 
  **id** | **string**| The process group id. | 

### Return type

[**PortEntity**](PortEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLabel**
> LabelEntity CreateLabel(ctx, body, id)
Creates a label

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LabelEntity**](LabelEntity.md)| The label configuration details. | 
  **id** | **string**| The process group id. | 

### Return type

[**LabelEntity**](LabelEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOutputPort**
> PortEntity CreateOutputPort(ctx, body, id)
Creates an output port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PortEntity**](PortEntity.md)| The output port configuration. | 
  **id** | **string**| The process group id. | 

### Return type

[**PortEntity**](PortEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProcessGroup**
> ProcessGroupEntity CreateProcessGroup(ctx, body, id, optional)
Creates a process group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProcessGroupEntity**](ProcessGroupEntity.md)| The process group configuration details. | 
  **id** | **string**| The process group id. | 
 **optional** | ***ProcessGroupsApiCreateProcessGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessGroupsApiCreateProcessGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parameterContextHandlingStrategy** | **optional.string**| Handling Strategy controls whether to keep or replace Parameter Contexts | [default to KEEP_EXISTING]

### Return type

[**ProcessGroupEntity**](ProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProcessor**
> ProcessorEntity CreateProcessor(ctx, body, id)
Creates a new processor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProcessorEntity**](ProcessorEntity.md)| The processor configuration details. | 
  **id** | **string**| The process group id. | 

### Return type

[**ProcessorEntity**](ProcessorEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRemoteProcessGroup**
> RemoteProcessGroupEntity CreateRemoteProcessGroup(ctx, body, id)
Creates a new process group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoteProcessGroupEntity**](RemoteProcessGroupEntity.md)| The remote process group configuration details. | 
  **id** | **string**| The process group id. | 

### Return type

[**RemoteProcessGroupEntity**](RemoteProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteReplaceProcessGroupRequest**
> ProcessGroupReplaceRequestEntity DeleteReplaceProcessGroupRequest(ctx, id, optional)
Deletes the Replace Request with the given ID

Deletes the Replace Request with the given ID. After a request is created via a POST to /process-groups/{id}/replace-requests, it is expected that the client will properly clean up the request by DELETE'ing it, once the Replace process has completed. If the request is deleted before the request completes, then the Replace request will finish the step that it is currently performing and then will cancel any subsequent steps. Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the Update Request | 
 **optional** | ***ProcessGroupsApiDeleteReplaceProcessGroupRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessGroupsApiDeleteReplaceProcessGroupRequestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disconnectedNodeAcknowledged** | **optional.bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ProcessGroupReplaceRequestEntity**](ProcessGroupReplaceRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportProcessGroup**
> string ExportProcessGroup(ctx, id, optional)
Gets a process group for download

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 
 **optional** | ***ProcessGroupsApiExportProcessGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessGroupsApiExportProcessGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeReferencedServices** | **optional.bool**| If referenced services from outside the target group should be included | [default to false]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnections**
> ConnectionsEntity GetConnections(ctx, id)
Gets all connections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 

### Return type

[**ConnectionsEntity**](ConnectionsEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDropAllFlowfilesRequest**
> DropRequestEntity GetDropAllFlowfilesRequest(ctx, id, dropRequestId)
Gets the current status of a drop all flowfiles request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 
  **dropRequestId** | **string**| The drop request id. | 

### Return type

[**DropRequestEntity**](DropRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFunnels**
> FunnelsEntity GetFunnels(ctx, id)
Gets all funnels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 

### Return type

[**FunnelsEntity**](FunnelsEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInputPorts**
> InputPortsEntity GetInputPorts(ctx, id)
Gets all input ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 

### Return type

[**InputPortsEntity**](InputPortsEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLabels**
> LabelsEntity GetLabels(ctx, id)
Gets all labels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 

### Return type

[**LabelsEntity**](LabelsEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLocalModifications**
> FlowComparisonEntity GetLocalModifications(ctx, id)
Gets a list of local modifications to the Process Group since it was last synchronized with the Flow Registry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 

### Return type

[**FlowComparisonEntity**](FlowComparisonEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOutputPorts**
> OutputPortsEntity GetOutputPorts(ctx, id)
Gets all output ports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 

### Return type

[**OutputPortsEntity**](OutputPortsEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcessGroup**
> ProcessGroupEntity GetProcessGroup(ctx, id)
Gets a process group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 

### Return type

[**ProcessGroupEntity**](ProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcessGroups**
> ProcessGroupsEntity GetProcessGroups(ctx, id)
Gets all process groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 

### Return type

[**ProcessGroupsEntity**](ProcessGroupsEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcessors**
> ProcessorsEntity GetProcessors(ctx, id, optional)
Gets all processors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 
 **optional** | ***ProcessGroupsApiGetProcessorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessGroupsApiGetProcessorsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDescendantGroups** | **optional.bool**| Whether or not to include processors from descendant process groups | [default to false]

### Return type

[**ProcessorsEntity**](ProcessorsEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRemoteProcessGroups**
> RemoteProcessGroupsEntity GetRemoteProcessGroups(ctx, id)
Gets all remote process groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 

### Return type

[**RemoteProcessGroupsEntity**](RemoteProcessGroupsEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReplaceProcessGroupRequest**
> ProcessGroupReplaceRequestEntity GetReplaceProcessGroupRequest(ctx, id)
Returns the Replace Request with the given ID

Returns the Replace Request with the given ID. Once a Replace Request has been created by performing a POST to /process-groups/{id}/replace-requests, that request can subsequently be retrieved via this endpoint, and the request that is fetched will contain the updated state, such as percent complete, the current state of the request, and any failures. Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the Replace Request | 

### Return type

[**ProcessGroupReplaceRequestEntity**](ProcessGroupReplaceRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportProcessGroup**
> ProcessGroupEntity ImportProcessGroup(ctx, id, optional)
Imports a specified process group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 
 **optional** | ***ProcessGroupsApiImportProcessGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessGroupsApiImportProcessGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ProcessGroupUploadEntity**](ProcessGroupUploadEntity.md)|  | 

### Return type

[**ProcessGroupEntity**](ProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InitiateReplaceProcessGroup**
> ProcessGroupReplaceRequestEntity InitiateReplaceProcessGroup(ctx, body, id)
Initiate the Replace Request of a Process Group with the given ID

This will initiate the action of replacing a process group with the given process group. This can be a lengthy process, as it will stop any Processors and disable any Controller Services necessary to perform the action and then restart them. As a result, the endpoint will immediately return a ProcessGroupReplaceRequestEntity, and the process of replacing the flow will occur asynchronously in the background. The client may then periodically poll the status of the request by issuing a GET request to /process-groups/replace-requests/{requestId}. Once the request is completed, the client is expected to issue a DELETE request to /process-groups/replace-requests/{requestId}. Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProcessGroupImportEntity**](ProcessGroupImportEntity.md)| The process group replace request entity | 
  **id** | **string**| The process group id. | 

### Return type

[**ProcessGroupReplaceRequestEntity**](ProcessGroupReplaceRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Paste**
> PasteResponseEntity Paste(ctx, body, id)
Pastes into the specified process group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PasteRequestEntity**](PasteRequestEntity.md)| The request including the components to be pasted into the specified Process Group. | 
  **id** | **string**| The process group id. | 

### Return type

[**PasteResponseEntity**](PasteResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveDropRequest1**
> DropRequestEntity RemoveDropRequest1(ctx, id, dropRequestId)
Cancels and/or removes a request to drop all flowfiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 
  **dropRequestId** | **string**| The drop request id. | 

### Return type

[**DropRequestEntity**](DropRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveProcessGroup**
> ProcessGroupEntity RemoveProcessGroup(ctx, id, optional)
Deletes a process group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 
 **optional** | ***ProcessGroupsApiRemoveProcessGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessGroupsApiRemoveProcessGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | [**optional.Interface of LongParameter**](.md)| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | [**optional.Interface of ClientIdParameter**](.md)| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ProcessGroupEntity**](ProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceProcessGroup**
> ProcessGroupImportEntity ReplaceProcessGroup(ctx, body, id)
Replace Process Group contents with the given ID with the specified Process Group contents

This endpoint is used for replication within a cluster, when replacing a flow with a new flow. It expects that the flow beingreplaced is not under version control and that the given snapshot will not modify any Processor that is currently running or any Controller Service that is enabled. Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProcessGroupImportEntity**](ProcessGroupImportEntity.md)| The process group replace request entity. | 
  **id** | **string**| The process group id. | 

### Return type

[**ProcessGroupImportEntity**](ProcessGroupImportEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProcessGroup**
> ProcessGroupEntity UpdateProcessGroup(ctx, body, id)
Updates a process group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProcessGroupEntity**](ProcessGroupEntity.md)| The process group configuration details. | 
  **id** | **string**| The process group id. | 

### Return type

[**ProcessGroupEntity**](ProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadProcessGroup**
> ProcessGroupEntity UploadProcessGroup(ctx, id, optional)
Uploads a versioned flow definition and creates a process group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 
 **optional** | ***ProcessGroupsApiUploadProcessGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessGroupsApiUploadProcessGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **optional.string**|  | 
 **disconnectedNodeAcknowledged** | **optional.bool**|  | 
 **file** | [**optional.Interface of interface{}**](.md)|  | 
 **groupName** | **optional.string**|  | 
 **positionX** | **optional.float64**|  | 
 **positionY** | **optional.float64**|  | 

### Return type

[**ProcessGroupEntity**](ProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

