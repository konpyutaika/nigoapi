# \ParameterProvidersApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyzeConfiguration**](ParameterProvidersApi.md#AnalyzeConfiguration) | **Post** /parameter-providers/{id}/config/analysis | Performs analysis of the component&#39;s configuration, providing information about which attributes are referenced.
[**ClearState**](ParameterProvidersApi.md#ClearState) | **Post** /parameter-providers/{id}/state/clear-requests | Clears the state for a parameter provider
[**DeleteApplyParametersRequest**](ParameterProvidersApi.md#DeleteApplyParametersRequest) | **Delete** /parameter-providers/{providerId}/apply-parameters-requests/{requestId} | Deletes the Apply Parameters Request with the given ID
[**DeleteVerificationRequest**](ParameterProvidersApi.md#DeleteVerificationRequest) | **Delete** /parameter-providers/{id}/config/verification-requests/{requestId} | Deletes the Verification Request with the given ID
[**FetchParameters**](ParameterProvidersApi.md#FetchParameters) | **Post** /parameter-providers/{id}/parameters/fetch-requests | Fetches and temporarily caches the parameters for a provider
[**GetParameterProvider**](ParameterProvidersApi.md#GetParameterProvider) | **Get** /parameter-providers/{id} | Gets a parameter provider
[**GetParameterProviderApplyParametersRequest**](ParameterProvidersApi.md#GetParameterProviderApplyParametersRequest) | **Get** /parameter-providers/{providerId}/apply-parameters-requests/{requestId} | Returns the Apply Parameters Request with the given ID
[**GetParameterProviderReferences**](ParameterProvidersApi.md#GetParameterProviderReferences) | **Get** /parameter-providers/{id}/references | Gets all references to a parameter provider
[**GetPropertyDescriptor**](ParameterProvidersApi.md#GetPropertyDescriptor) | **Get** /parameter-providers/{id}/descriptors | Gets a parameter provider property descriptor
[**GetState**](ParameterProvidersApi.md#GetState) | **Get** /parameter-providers/{id}/state | Gets the state for a parameter provider
[**GetVerificationRequest**](ParameterProvidersApi.md#GetVerificationRequest) | **Get** /parameter-providers/{id}/config/verification-requests/{requestId} | Returns the Verification Request with the given ID
[**RemoveParameterProvider**](ParameterProvidersApi.md#RemoveParameterProvider) | **Delete** /parameter-providers/{id} | Deletes a parameter provider
[**SubmitApplyParameters**](ParameterProvidersApi.md#SubmitApplyParameters) | **Post** /parameter-providers/{providerId}/apply-parameters-requests | Initiate a request to apply the fetched parameters of a Parameter Provider
[**SubmitConfigVerificationRequest**](ParameterProvidersApi.md#SubmitConfigVerificationRequest) | **Post** /parameter-providers/{id}/config/verification-requests | Performs verification of the Parameter Provider&#39;s configuration
[**UpdateParameterProvider**](ParameterProvidersApi.md#UpdateParameterProvider) | **Put** /parameter-providers/{id} | Updates a parameter provider


# **AnalyzeConfiguration**
> ConfigurationAnalysisEntity AnalyzeConfiguration(ctx, id, body)
Performs analysis of the component's configuration, providing information about which attributes are referenced.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The parameter provider id. | 
  **body** | [**ConfigurationAnalysisEntity**](ConfigurationAnalysisEntity.md)| The configuration analysis request. | 

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
Clears the state for a parameter provider



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The parameter provider id. | 

### Return type

[**ComponentStateEntity**](ComponentStateEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApplyParametersRequest**
> ParameterProviderApplyParametersRequestEntity DeleteApplyParametersRequest(ctx, providerId, requestId, optional)
Deletes the Apply Parameters Request with the given ID

Deletes the Apply Parameters Request with the given ID. After a request is created via a POST to /nifi-api/parameter-providers/apply-parameters-requests, it is expected that the client will properly clean up the request by DELETE'ing it, once the Apply process has completed. If the request is deleted before the request completes, then the Apply Parameters Request will finish the step that it is currently performing and then will cancel any subsequent steps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the Parameter Provider | 
  **requestId** | **string**| The ID of the Apply Parameters Request | 
 **optional** | ***ParameterProvidersApiDeleteApplyParametersRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ParameterProvidersApiDeleteApplyParametersRequestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **disconnectedNodeAcknowledged** | **optional.Bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ParameterProviderApplyParametersRequestEntity**](ParameterProviderApplyParametersRequestEntity.md)

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
  **id** | **string**| The ID of the Parameter Provider | 
  **requestId** | **string**| The ID of the Verification Request | 

### Return type

[**VerifyConfigRequestEntity**](VerifyConfigRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchParameters**
> ParameterProviderEntity FetchParameters(ctx, id, body)
Fetches and temporarily caches the parameters for a provider



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The parameter provider id. | 
  **body** | [**ParameterProviderParameterFetchEntity**](ParameterProviderParameterFetchEntity.md)| The parameter fetch request. | 

### Return type

[**ParameterProviderEntity**](ParameterProviderEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetParameterProvider**
> ParameterProviderEntity GetParameterProvider(ctx, id)
Gets a parameter provider



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The parameter provider id. | 

### Return type

[**ParameterProviderEntity**](ParameterProviderEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetParameterProviderApplyParametersRequest**
> ParameterProviderApplyParametersRequestEntity GetParameterProviderApplyParametersRequest(ctx, providerId, requestId)
Returns the Apply Parameters Request with the given ID

Returns the Apply Parameters Request with the given ID. Once an Apply Parameters Request has been created by performing a POST to /nifi-api/parameter-providers, that request can subsequently be retrieved via this endpoint, and the request that is fetched will contain the state, such as percent complete, the current state of the request, and any failures. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**| The ID of the Parameter Provider | 
  **requestId** | **string**| The ID of the Apply Parameters Request | 

### Return type

[**ParameterProviderApplyParametersRequestEntity**](ParameterProviderApplyParametersRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetParameterProviderReferences**
> ParameterProviderReferencingComponentsEntity GetParameterProviderReferences(ctx, id)
Gets all references to a parameter provider



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The parameter provider id. | 

### Return type

[**ParameterProviderReferencingComponentsEntity**](ParameterProviderReferencingComponentsEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPropertyDescriptor**
> PropertyDescriptorEntity GetPropertyDescriptor(ctx, id, propertyName)
Gets a parameter provider property descriptor



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The parameter provider id. | 
  **propertyName** | **string**| The property name. | 

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
Gets the state for a parameter provider



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The parameter provider id. | 

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
  **id** | **string**| The ID of the Parameter Provider | 
  **requestId** | **string**| The ID of the Verification Request | 

### Return type

[**VerifyConfigRequestEntity**](VerifyConfigRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveParameterProvider**
> ParameterProviderEntity RemoveParameterProvider(ctx, id, optional)
Deletes a parameter provider



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The parameter provider id. | 
 **optional** | ***ParameterProvidersApiRemoveParameterProviderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ParameterProvidersApiRemoveParameterProviderOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.String**| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | **optional.String**| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.Bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ParameterProviderEntity**](ParameterProviderEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitApplyParameters**
> ParameterProviderApplyParametersRequestEntity SubmitApplyParameters(ctx, providerId, body)
Initiate a request to apply the fetched parameters of a Parameter Provider

This will initiate the process of applying fetched parameters to all referencing Parameter Contexts. Changing the value of a Parameter may require that one or more components be stopped and restarted, so this action may take significantly more time than many other REST API actions. As a result, this endpoint will immediately return a ParameterProviderApplyParametersRequestEntity, and the process of updating the necessary components will occur asynchronously in the background. The client may then periodically poll the status of the request by issuing a GET request to /parameter-providers/apply-parameters-requests/{requestId}. Once the request is completed, the client is expected to issue a DELETE request to /parameter-providers/apply-parameters-requests/{requestId}.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **providerId** | **string**|  | 
  **body** | [**ParameterProviderParameterApplicationEntity**](ParameterProviderParameterApplicationEntity.md)| The apply parameters request. | 

### Return type

[**ParameterProviderApplyParametersRequestEntity**](ParameterProviderApplyParametersRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitConfigVerificationRequest**
> VerifyConfigRequestEntity SubmitConfigVerificationRequest(ctx, id, body)
Performs verification of the Parameter Provider's configuration

This will initiate the process of verifying a given Parameter Provider configuration. This may be a long-running task. As a result, this endpoint will immediately return a ParameterProviderConfigVerificationRequestEntity, and the process of performing the verification will occur asynchronously in the background. The client may then periodically poll the status of the request by issuing a GET request to /parameter-providers/{serviceId}/verification-requests/{requestId}. Once the request is completed, the client is expected to issue a DELETE request to /parameter-providers/{providerId}/verification-requests/{requestId}.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The parameter provider id. | 
  **body** | [**VerifyConfigRequestEntity**](VerifyConfigRequestEntity.md)| The parameter provider configuration verification request. | 

### Return type

[**VerifyConfigRequestEntity**](VerifyConfigRequestEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateParameterProvider**
> ParameterProviderEntity UpdateParameterProvider(ctx, id, body)
Updates a parameter provider



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The parameter provider id. | 
  **body** | [**ParameterProviderEntity**](ParameterProviderEntity.md)| The parameter provider configuration details. | 

### Return type

[**ParameterProviderEntity**](ParameterProviderEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

