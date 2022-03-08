# \ProcessorsApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyzeConfiguration**](ProcessorsApi.md#AnalyzeConfiguration) | **Post** /processors/{id}/config/analysis | Performs analysis of the component&#39;s configuration, providing information about which attributes are referenced.
[**ClearState**](ProcessorsApi.md#ClearState) | **Post** /processors/{id}/state/clear-requests | Clears the state for a processor
[**DeleteProcessor**](ProcessorsApi.md#DeleteProcessor) | **Delete** /processors/{id} | Deletes a processor
[**DeleteVerificationRequest**](ProcessorsApi.md#DeleteVerificationRequest) | **Delete** /processors/{id}/config/verification-requests/{requestId} | Deletes the Verification Request with the given ID
[**GetProcessor**](ProcessorsApi.md#GetProcessor) | **Get** /processors/{id} | Gets a processor
[**GetProcessorDiagnostics**](ProcessorsApi.md#GetProcessorDiagnostics) | **Get** /processors/{id}/diagnostics | Gets diagnostics information about a processor
[**GetProcessorRunStatusDetails**](ProcessorsApi.md#GetProcessorRunStatusDetails) | **Post** /processors/run-status-details/queries | Submits a query to retrieve the run status details of all processors that are in the given list of Processor IDs
[**GetPropertyDescriptor**](ProcessorsApi.md#GetPropertyDescriptor) | **Get** /processors/{id}/descriptors | Gets the descriptor for a processor property
[**GetState**](ProcessorsApi.md#GetState) | **Get** /processors/{id}/state | Gets the state for a processor
[**GetVerificationRequest**](ProcessorsApi.md#GetVerificationRequest) | **Get** /processors/{id}/config/verification-requests/{requestId} | Returns the Verification Request with the given ID
[**SubmitProcessorVerificationRequest**](ProcessorsApi.md#SubmitProcessorVerificationRequest) | **Post** /processors/{id}/config/verification-requests | Performs verification of the Processor&#39;s configuration
[**TerminateProcessor**](ProcessorsApi.md#TerminateProcessor) | **Delete** /processors/{id}/threads | Terminates a processor, essentially \&quot;deleting\&quot; its threads and any active tasks
[**UpdateProcessor**](ProcessorsApi.md#UpdateProcessor) | **Put** /processors/{id} | Updates a processor
[**UpdateRunStatus**](ProcessorsApi.md#UpdateRunStatus) | **Put** /processors/{id}/run-status | Updates run status of a processor


# **AnalyzeConfiguration**
> ConfigurationAnalysisEntity AnalyzeConfiguration(ctx, id, body)
Performs analysis of the component's configuration, providing information about which attributes are referenced.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The processor id. | 
  **body** | [**ConfigurationAnalysisEntity**](ConfigurationAnalysisEntity.md)| The processor configuration analysis request. | 

### Return type

[**ConfigurationAnalysisEntity**](ConfigurationAnalysisEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearState**
> ComponentStateEntity ClearState(ctx, id)
Clears the state for a processor



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

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProcessor**
> ProcessorEntity DeleteProcessor(ctx, id, optional)
Deletes a processor



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The processor id. | 
 **optional** | ***ProcessorsApiDeleteProcessorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessorsApiDeleteProcessorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.String**| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | **optional.String**| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.Bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ProcessorEntity**](ProcessorEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVerificationRequest**
> VerifyConfigRequestEntity DeleteVerificationRequest(ctx, id, requestId)
Deletes the Verification Request with the given ID

Deletes the Verification Request with the given ID. After a request is created, it is expected that the client will properly clean up the request by DELETE'ing it, once the Verification process has completed. If the request is deleted before the request completes, then the Verification request will finish the step that it is currently performing and then will cancel any subsequent steps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the Processor | 
  **requestId** | **string**| The ID of the Verification Request | 

### Return type

[**VerifyConfigRequestEntity**](VerifyConfigRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcessor**
> ProcessorEntity GetProcessor(ctx, id)
Gets a processor



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The processor id. | 

### Return type

[**ProcessorEntity**](ProcessorEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcessorDiagnostics**
> ProcessorEntity GetProcessorDiagnostics(ctx, id)
Gets diagnostics information about a processor

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The processor id. | 

### Return type

[**ProcessorEntity**](ProcessorEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcessorRunStatusDetails**
> ProcessorsRunStatusDetailsEntity GetProcessorRunStatusDetails(ctx, optional)
Submits a query to retrieve the run status details of all processors that are in the given list of Processor IDs



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProcessorsApiGetProcessorRunStatusDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessorsApiGetProcessorRunStatusDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RunStatusDetailsRequestEntity**](RunStatusDetailsRequestEntity.md)| The request for the processors that should be included in the results | 

### Return type

[**ProcessorsRunStatusDetailsEntity**](ProcessorsRunStatusDetailsEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPropertyDescriptor**
> PropertyDescriptorEntity GetPropertyDescriptor(ctx, id, propertyName, optional)
Gets the descriptor for a processor property



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The processor id. | 
  **propertyName** | **string**| The property name. | 
 **optional** | ***ProcessorsApiGetPropertyDescriptorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProcessorsApiGetPropertyDescriptorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **optional.String**| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 

### Return type

[**PropertyDescriptorEntity**](PropertyDescriptorEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetState**
> ComponentStateEntity GetState(ctx, id)
Gets the state for a processor



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

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVerificationRequest**
> VerifyConfigRequestEntity GetVerificationRequest(ctx, id, requestId)
Returns the Verification Request with the given ID

Returns the Verification Request with the given ID. Once an Verification Request has been created, that request can subsequently be retrieved via this endpoint, and the request that is fetched will contain the updated state, such as percent complete, the current state of the request, and any failures. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the Processor | 
  **requestId** | **string**| The ID of the Verification Request | 

### Return type

[**VerifyConfigRequestEntity**](VerifyConfigRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitProcessorVerificationRequest**
> VerifyConfigRequestEntity SubmitProcessorVerificationRequest(ctx, id, body)
Performs verification of the Processor's configuration

This will initiate the process of verifying a given Processor configuration. This may be a long-running task. As a result, this endpoint will immediately return a ProcessorConfigVerificationRequestEntity, and the process of performing the verification will occur asynchronously in the background. The client may then periodically poll the status of the request by issuing a GET request to /processors/{processorId}/verification-requests/{requestId}. Once the request is completed, the client is expected to issue a DELETE request to /processors/{processorId}/verification-requests/{requestId}.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The processor id. | 
  **body** | [**VerifyConfigRequestEntity**](VerifyConfigRequestEntity.md)| The processor configuration verification request. | 

### Return type

[**VerifyConfigRequestEntity**](VerifyConfigRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminateProcessor**
> ProcessorEntity TerminateProcessor(ctx, id)
Terminates a processor, essentially \"deleting\" its threads and any active tasks



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The processor id. | 

### Return type

[**ProcessorEntity**](ProcessorEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProcessor**
> ProcessorEntity UpdateProcessor(ctx, id, body)
Updates a processor



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The processor id. | 
  **body** | [**ProcessorEntity**](ProcessorEntity.md)| The processor configuration details. | 

### Return type

[**ProcessorEntity**](ProcessorEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRunStatus**
> ProcessorEntity UpdateRunStatus(ctx, id, body)
Updates run status of a processor



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The processor id. | 
  **body** | [**ProcessorRunStatusEntity**](ProcessorRunStatusEntity.md)| The processor run status. | 

### Return type

[**ProcessorEntity**](ProcessorEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

