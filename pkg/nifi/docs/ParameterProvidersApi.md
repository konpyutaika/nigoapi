# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyzeConfiguration1**](ParameterProvidersApi.md#AnalyzeConfiguration1) | **Post** /parameter-providers/{id}/config/analysis | Performs analysis of the component&#x27;s configuration, providing information about which attributes are referenced.
[**ClearState2**](ParameterProvidersApi.md#ClearState2) | **Post** /parameter-providers/{id}/state/clear-requests | Clears the state for a parameter provider
[**DeleteApplyParametersRequest**](ParameterProvidersApi.md#DeleteApplyParametersRequest) | **Delete** /parameter-providers/{providerId}/apply-parameters-requests/{requestId} | Deletes the Apply Parameters Request with the given ID
[**DeleteVerificationRequest1**](ParameterProvidersApi.md#DeleteVerificationRequest1) | **Delete** /parameter-providers/{id}/config/verification-requests/{requestId} | Deletes the Verification Request with the given ID
[**FetchParameters**](ParameterProvidersApi.md#FetchParameters) | **Post** /parameter-providers/{id}/parameters/fetch-requests | Fetches and temporarily caches the parameters for a provider
[**GetParameterProvider**](ParameterProvidersApi.md#GetParameterProvider) | **Get** /parameter-providers/{id} | Gets a parameter provider
[**GetParameterProviderApplyParametersRequest**](ParameterProvidersApi.md#GetParameterProviderApplyParametersRequest) | **Get** /parameter-providers/{providerId}/apply-parameters-requests/{requestId} | Returns the Apply Parameters Request with the given ID
[**GetParameterProviderReferences**](ParameterProvidersApi.md#GetParameterProviderReferences) | **Get** /parameter-providers/{id}/references | Gets all references to a parameter provider
[**GetPropertyDescriptor2**](ParameterProvidersApi.md#GetPropertyDescriptor2) | **Get** /parameter-providers/{id}/descriptors | Gets a parameter provider property descriptor
[**GetState1**](ParameterProvidersApi.md#GetState1) | **Get** /parameter-providers/{id}/state | Gets the state for a parameter provider
[**GetVerificationRequest1**](ParameterProvidersApi.md#GetVerificationRequest1) | **Get** /parameter-providers/{id}/config/verification-requests/{requestId} | Returns the Verification Request with the given ID
[**RemoveParameterProvider**](ParameterProvidersApi.md#RemoveParameterProvider) | **Delete** /parameter-providers/{id} | Deletes a parameter provider
[**SubmitApplyParameters**](ParameterProvidersApi.md#SubmitApplyParameters) | **Post** /parameter-providers/{providerId}/apply-parameters-requests | Initiate a request to apply the fetched parameters of a Parameter Provider
[**SubmitConfigVerificationRequest1**](ParameterProvidersApi.md#SubmitConfigVerificationRequest1) | **Post** /parameter-providers/{id}/config/verification-requests | Performs verification of the Parameter Provider&#x27;s configuration
[**UpdateParameterProvider**](ParameterProvidersApi.md#UpdateParameterProvider) | **Put** /parameter-providers/{id} | Updates a parameter provider

# **AnalyzeConfiguration1**
> ConfigurationAnalysisEntity AnalyzeConfiguration1(ctx, body, id)
Performs analysis of the component's configuration, providing information about which attributes are referenced.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConfigurationAnalysisEntity**](ConfigurationAnalysisEntity.md)| The configuration analysis request. | 
  **id** | **string**| The parameter provider id. | 

### Return type

[**ConfigurationAnalysisEntity**](ConfigurationAnalysisEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearState2**
> ComponentStateEntity ClearState2(ctx, id, optional)
Clears the state for a parameter provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The parameter provider id. | 
 **optional** | ***ParameterProvidersApiClearState2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ParameterProvidersApiClearState2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ComponentStateEntity**](ComponentStateEntity.md)| Optional component state to perform a selective key removal. If omitted, clears all state. | 

### Return type

[**ComponentStateEntity**](ComponentStateEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: */*, application/json
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


 **disconnectedNodeAcknowledged** | **optional.bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ParameterProviderApplyParametersRequestEntity**](ParameterProviderApplyParametersRequestEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVerificationRequest1**
> VerifyConfigRequestEntity DeleteVerificationRequest1(ctx, id, requestId)
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

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchParameters**
> ParameterProviderEntity FetchParameters(ctx, body, id)
Fetches and temporarily caches the parameters for a provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ParameterProviderParameterFetchEntity**](ParameterProviderParameterFetchEntity.md)| The parameter fetch request. | 
  **id** | **string**| The parameter provider id. | 

### Return type

[**ParameterProviderEntity**](ParameterProviderEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

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

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
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

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
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

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPropertyDescriptor2**
> PropertyDescriptorEntity GetPropertyDescriptor2(ctx, id, propertyName)
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

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetState1**
> ComponentStateEntity GetState1(ctx, id)
Gets the state for a parameter provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The parameter provider id. | 

### Return type

[**ComponentStateEntity**](ComponentStateEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVerificationRequest1**
> VerifyConfigRequestEntity GetVerificationRequest1(ctx, id, requestId)
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

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
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

 **version** | [**optional.Interface of LongParameter**](.md)| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | [**optional.Interface of ClientIdParameter**](.md)| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ParameterProviderEntity**](ParameterProviderEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitApplyParameters**
> ParameterProviderApplyParametersRequestEntity SubmitApplyParameters(ctx, body, providerId)
Initiate a request to apply the fetched parameters of a Parameter Provider

This will initiate the process of applying fetched parameters to all referencing Parameter Contexts. Changing the value of a Parameter may require that one or more components be stopped and restarted, so this action may take significantly more time than many other REST API actions. As a result, this endpoint will immediately return a ParameterProviderApplyParametersRequestEntity, and the process of updating the necessary components will occur asynchronously in the background. The client may then periodically poll the status of the request by issuing a GET request to /parameter-providers/apply-parameters-requests/{requestId}. Once the request is completed, the client is expected to issue a DELETE request to /parameter-providers/apply-parameters-requests/{requestId}.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ParameterProviderParameterApplicationEntity**](ParameterProviderParameterApplicationEntity.md)| The apply parameters request. | 
  **providerId** | **string**|  | 

### Return type

[**ParameterProviderApplyParametersRequestEntity**](ParameterProviderApplyParametersRequestEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitConfigVerificationRequest1**
> VerifyConfigRequestEntity SubmitConfigVerificationRequest1(ctx, body, id)
Performs verification of the Parameter Provider's configuration

This will initiate the process of verifying a given Parameter Provider configuration. This may be a long-running task. As a result, this endpoint will immediately return a ParameterProviderConfigVerificationRequestEntity, and the process of performing the verification will occur asynchronously in the background. The client may then periodically poll the status of the request by issuing a GET request to /parameter-providers/{serviceId}/verification-requests/{requestId}. Once the request is completed, the client is expected to issue a DELETE request to /parameter-providers/{providerId}/verification-requests/{requestId}.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VerifyConfigRequestEntity**](VerifyConfigRequestEntity.md)| The parameter provider configuration verification request. | 
  **id** | **string**| The parameter provider id. | 

### Return type

[**VerifyConfigRequestEntity**](VerifyConfigRequestEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateParameterProvider**
> ParameterProviderEntity UpdateParameterProvider(ctx, body, id)
Updates a parameter provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ParameterProviderEntity**](ParameterProviderEntity.md)| The parameter provider configuration details. | 
  **id** | **string**| The parameter provider id. | 

### Return type

[**ParameterProviderEntity**](ParameterProviderEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

