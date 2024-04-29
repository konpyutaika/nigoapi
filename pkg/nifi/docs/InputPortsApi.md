# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInputPort**](InputPortsApi.md#GetInputPort) | **Get** /input-ports/{id} | Gets an input port
[**RemoveInputPort**](InputPortsApi.md#RemoveInputPort) | **Delete** /input-ports/{id} | Deletes an input port
[**UpdateInputPort**](InputPortsApi.md#UpdateInputPort) | **Put** /input-ports/{id} | Updates an input port
[**UpdateRunStatus2**](InputPortsApi.md#UpdateRunStatus2) | **Put** /input-ports/{id}/run-status | Updates run status of an input-port

# **GetInputPort**
> PortEntity GetInputPort(ctx, id)
Gets an input port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The input port id. | 

### Return type

[**PortEntity**](PortEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveInputPort**
> PortEntity RemoveInputPort(ctx, id, optional)
Deletes an input port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The input port id. | 
 **optional** | ***InputPortsApiRemoveInputPortOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InputPortsApiRemoveInputPortOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | [**optional.Interface of LongParameter**](.md)| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | [**optional.Interface of ClientIdParameter**](.md)| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**PortEntity**](PortEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInputPort**
> PortEntity UpdateInputPort(ctx, body, id)
Updates an input port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PortEntity**](PortEntity.md)| The input port configuration details. | 
  **id** | **string**| The input port id. | 

### Return type

[**PortEntity**](PortEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRunStatus2**
> ProcessorEntity UpdateRunStatus2(ctx, body, id)
Updates run status of an input-port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PortRunStatusEntity**](PortRunStatusEntity.md)| The port run status. | 
  **id** | **string**| The port id. | 

### Return type

[**ProcessorEntity**](ProcessorEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

