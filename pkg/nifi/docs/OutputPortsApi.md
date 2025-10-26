# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOutputPort**](OutputPortsApi.md#GetOutputPort) | **Get** /output-ports/{id} | Gets an output port
[**RemoveOutputPort**](OutputPortsApi.md#RemoveOutputPort) | **Delete** /output-ports/{id} | Deletes an output port
[**UpdateOutputPort**](OutputPortsApi.md#UpdateOutputPort) | **Put** /output-ports/{id} | Updates an output port
[**UpdateRunStatus3**](OutputPortsApi.md#UpdateRunStatus3) | **Put** /output-ports/{id}/run-status | Updates run status of an output-port

# **GetOutputPort**
> PortEntity GetOutputPort(ctx, id)
Gets an output port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The output port id. | 

### Return type

[**PortEntity**](PortEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveOutputPort**
> PortEntity RemoveOutputPort(ctx, id, optional)
Deletes an output port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The output port id. | 
 **optional** | ***OutputPortsApiRemoveOutputPortOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OutputPortsApiRemoveOutputPortOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | [**optional.Interface of LongParameter**](.md)| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | [**optional.Interface of ClientIdParameter**](.md)| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**PortEntity**](PortEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOutputPort**
> PortEntity UpdateOutputPort(ctx, body, id)
Updates an output port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PortEntity**](PortEntity.md)| The output port configuration details. | 
  **id** | **string**| The output port id. | 

### Return type

[**PortEntity**](PortEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRunStatus3**
> ProcessorEntity UpdateRunStatus3(ctx, body, id)
Updates run status of an output-port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PortRunStatusEntity**](PortRunStatusEntity.md)| The port run status. | 
  **id** | **string**| The port id. | 

### Return type

[**ProcessorEntity**](ProcessorEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

