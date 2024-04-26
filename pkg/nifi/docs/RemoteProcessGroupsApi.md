# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRemoteProcessGroup**](RemoteProcessGroupsApi.md#GetRemoteProcessGroup) | **Get** /remote-process-groups/{id} | Gets a remote process group
[**GetState3**](RemoteProcessGroupsApi.md#GetState3) | **Get** /remote-process-groups/{id}/state | Gets the state for a RemoteProcessGroup
[**RemoveRemoteProcessGroup**](RemoteProcessGroupsApi.md#RemoveRemoteProcessGroup) | **Delete** /remote-process-groups/{id} | Deletes a remote process group
[**UpdateRemoteProcessGroup**](RemoteProcessGroupsApi.md#UpdateRemoteProcessGroup) | **Put** /remote-process-groups/{id} | Updates a remote process group
[**UpdateRemoteProcessGroupInputPort**](RemoteProcessGroupsApi.md#UpdateRemoteProcessGroupInputPort) | **Put** /remote-process-groups/{id}/input-ports/{port-id} | Updates a remote port
[**UpdateRemoteProcessGroupInputPortRunStatus**](RemoteProcessGroupsApi.md#UpdateRemoteProcessGroupInputPortRunStatus) | **Put** /remote-process-groups/{id}/input-ports/{port-id}/run-status | Updates run status of a remote port
[**UpdateRemoteProcessGroupOutputPort**](RemoteProcessGroupsApi.md#UpdateRemoteProcessGroupOutputPort) | **Put** /remote-process-groups/{id}/output-ports/{port-id} | Updates a remote port
[**UpdateRemoteProcessGroupOutputPortRunStatus**](RemoteProcessGroupsApi.md#UpdateRemoteProcessGroupOutputPortRunStatus) | **Put** /remote-process-groups/{id}/output-ports/{port-id}/run-status | Updates run status of a remote port
[**UpdateRemoteProcessGroupRunStatus**](RemoteProcessGroupsApi.md#UpdateRemoteProcessGroupRunStatus) | **Put** /remote-process-groups/{id}/run-status | Updates run status of a remote process group
[**UpdateRemoteProcessGroupRunStatuses**](RemoteProcessGroupsApi.md#UpdateRemoteProcessGroupRunStatuses) | **Put** /remote-process-groups/process-group/{id}/run-status | Updates run status of all remote process groups in a process group (recursively)

# **GetRemoteProcessGroup**
> RemoteProcessGroupEntity GetRemoteProcessGroup(ctx, id)
Gets a remote process group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The remote process group id. | 

### Return type

[**RemoteProcessGroupEntity**](RemoteProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetState3**
> ComponentStateEntity GetState3(ctx, id)
Gets the state for a RemoteProcessGroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The processor id. | 

### Return type

[**ComponentStateEntity**](ComponentStateEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveRemoteProcessGroup**
> RemoteProcessGroupEntity RemoveRemoteProcessGroup(ctx, id, optional)
Deletes a remote process group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The remote process group id. | 
 **optional** | ***RemoteProcessGroupsApiRemoveRemoteProcessGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RemoteProcessGroupsApiRemoveRemoteProcessGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | [**optional.Interface of LongParameter**](.md)| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | [**optional.Interface of ClientIdParameter**](.md)| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.Bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**RemoteProcessGroupEntity**](RemoteProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRemoteProcessGroup**
> RemoteProcessGroupEntity UpdateRemoteProcessGroup(ctx, body, id)
Updates a remote process group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoteProcessGroupEntity**](RemoteProcessGroupEntity.md)| The remote process group. | 
  **id** | **string**| The remote process group id. | 

### Return type

[**RemoteProcessGroupEntity**](RemoteProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRemoteProcessGroupInputPort**
> RemoteProcessGroupPortEntity UpdateRemoteProcessGroupInputPort(ctx, body, id, portId)
Updates a remote port

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoteProcessGroupPortEntity**](RemoteProcessGroupPortEntity.md)| The remote process group port. | 
  **id** | **string**| The remote process group id. | 
  **portId** | **string**| The remote process group port id. | 

### Return type

[**RemoteProcessGroupPortEntity**](RemoteProcessGroupPortEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRemoteProcessGroupInputPortRunStatus**
> RemoteProcessGroupPortEntity UpdateRemoteProcessGroupInputPortRunStatus(ctx, body, id, portId)
Updates run status of a remote port

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemotePortRunStatusEntity**](RemotePortRunStatusEntity.md)| The remote process group port. | 
  **id** | **string**| The remote process group id. | 
  **portId** | **string**| The remote process group port id. | 

### Return type

[**RemoteProcessGroupPortEntity**](RemoteProcessGroupPortEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRemoteProcessGroupOutputPort**
> RemoteProcessGroupPortEntity UpdateRemoteProcessGroupOutputPort(ctx, body, id, portId)
Updates a remote port

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoteProcessGroupPortEntity**](RemoteProcessGroupPortEntity.md)| The remote process group port. | 
  **id** | **string**| The remote process group id. | 
  **portId** | **string**| The remote process group port id. | 

### Return type

[**RemoteProcessGroupPortEntity**](RemoteProcessGroupPortEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRemoteProcessGroupOutputPortRunStatus**
> RemoteProcessGroupPortEntity UpdateRemoteProcessGroupOutputPortRunStatus(ctx, body, id, portId)
Updates run status of a remote port

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemotePortRunStatusEntity**](RemotePortRunStatusEntity.md)| The remote process group port. | 
  **id** | **string**| The remote process group id. | 
  **portId** | **string**| The remote process group port id. | 

### Return type

[**RemoteProcessGroupPortEntity**](RemoteProcessGroupPortEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRemoteProcessGroupRunStatus**
> RemoteProcessGroupEntity UpdateRemoteProcessGroupRunStatus(ctx, body, id)
Updates run status of a remote process group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemotePortRunStatusEntity**](RemotePortRunStatusEntity.md)| The remote process group run status. | 
  **id** | **string**| The remote process group id. | 

### Return type

[**RemoteProcessGroupEntity**](RemoteProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRemoteProcessGroupRunStatuses**
> RemoteProcessGroupEntity UpdateRemoteProcessGroupRunStatuses(ctx, body, id)
Updates run status of all remote process groups in a process group (recursively)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemotePortRunStatusEntity**](RemotePortRunStatusEntity.md)| The remote process groups run status. | 
  **id** | **string**| The process group id. | 

### Return type

[**RemoteProcessGroupEntity**](RemoteProcessGroupEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

