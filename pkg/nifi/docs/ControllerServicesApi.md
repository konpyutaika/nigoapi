# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyzeConfiguration**](ControllerServicesApi.md#AnalyzeConfiguration) | **Post** /controller-services/{id}/config/analysis | Performs analysis of the component&#x27;s configuration, providing information about which attributes are referenced.
[**ClearState1**](ControllerServicesApi.md#ClearState1) | **Post** /controller-services/{id}/state/clear-requests | Clears the state for a controller service
[**DeleteVerificationRequest**](ControllerServicesApi.md#DeleteVerificationRequest) | **Delete** /controller-services/{id}/config/verification-requests/{requestId} | Deletes the Verification Request with the given ID
[**GetControllerService**](ControllerServicesApi.md#GetControllerService) | **Get** /controller-services/{id} | Gets a controller service
[**GetControllerServiceReferences**](ControllerServicesApi.md#GetControllerServiceReferences) | **Get** /controller-services/{id}/references | Gets a controller service
[**GetPropertyDescriptor1**](ControllerServicesApi.md#GetPropertyDescriptor1) | **Get** /controller-services/{id}/descriptors | Gets a controller service property descriptor
[**GetState**](ControllerServicesApi.md#GetState) | **Get** /controller-services/{id}/state | Gets the state for a controller service
[**GetVerificationRequest**](ControllerServicesApi.md#GetVerificationRequest) | **Get** /controller-services/{id}/config/verification-requests/{requestId} | Returns the Verification Request with the given ID
[**RemoveControllerService**](ControllerServicesApi.md#RemoveControllerService) | **Delete** /controller-services/{id} | Deletes a controller service
[**SubmitConfigVerificationRequest**](ControllerServicesApi.md#SubmitConfigVerificationRequest) | **Post** /controller-services/{id}/config/verification-requests | Performs verification of the Controller Service&#x27;s configuration
[**UpdateControllerService**](ControllerServicesApi.md#UpdateControllerService) | **Put** /controller-services/{id} | Updates a controller service
[**UpdateControllerServiceReferences**](ControllerServicesApi.md#UpdateControllerServiceReferences) | **Put** /controller-services/{id}/references | Updates a controller services references
[**UpdateRunStatus1**](ControllerServicesApi.md#UpdateRunStatus1) | **Put** /controller-services/{id}/run-status | Updates run status of a controller service

# **AnalyzeConfiguration**
> ConfigurationAnalysisEntity AnalyzeConfiguration(ctx, body, id)
Performs analysis of the component's configuration, providing information about which attributes are referenced.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConfigurationAnalysisEntity**](ConfigurationAnalysisEntity.md)| The configuration analysis request. | 
  **id** | **string**| The controller service id. | 

### Return type

[**ConfigurationAnalysisEntity**](ConfigurationAnalysisEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearState1**
> ComponentStateEntity ClearState1(ctx, id, optional)
Clears the state for a controller service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The controller service id. | 
 **optional** | ***ControllerServicesApiClearState1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ControllerServicesApiClearState1Opts struct
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

# **DeleteVerificationRequest**
> VerifyConfigRequestEntity DeleteVerificationRequest(ctx, id, requestId)
Deletes the Verification Request with the given ID

Deletes the Verification Request with the given ID. After a request is created, it is expected that the client will properly clean up the request by DELETE'ing it, once the Verification process has completed. If the request is deleted before the request completes, then the Verification request will finish the step that it is currently performing and then will cancel any subsequent steps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the Controller Service | 
  **requestId** | **string**| The ID of the Verification Request | 

### Return type

[**VerifyConfigRequestEntity**](VerifyConfigRequestEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetControllerService**
> ControllerServiceEntity GetControllerService(ctx, id, optional)
Gets a controller service

If the uiOnly query parameter is provided with a value of true, the returned entity may only contain fields that are necessary for rendering the NiFi User Interface. As such, the selected fields may change at any time, even during incremental releases, without warning. As a result, this parameter should not be provided by any client other than the UI.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The controller service id. | 
 **optional** | ***ControllerServicesApiGetControllerServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ControllerServicesApiGetControllerServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uiOnly** | **optional.bool**|  | [default to false]

### Return type

[**ControllerServiceEntity**](ControllerServiceEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetControllerServiceReferences**
> ControllerServiceReferencingComponentsEntity GetControllerServiceReferences(ctx, id)
Gets a controller service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The controller service id. | 

### Return type

[**ControllerServiceReferencingComponentsEntity**](ControllerServiceReferencingComponentsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPropertyDescriptor1**
> PropertyDescriptorEntity GetPropertyDescriptor1(ctx, id, propertyName, optional)
Gets a controller service property descriptor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The controller service id. | 
  **propertyName** | **string**| The property name to return the descriptor for. | 
 **optional** | ***ControllerServicesApiGetPropertyDescriptor1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ControllerServicesApiGetPropertyDescriptor1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sensitive** | **optional.bool**| Property Descriptor requested sensitive status | 

### Return type

[**PropertyDescriptorEntity**](PropertyDescriptorEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetState**
> ComponentStateEntity GetState(ctx, id)
Gets the state for a controller service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The controller service id. | 

### Return type

[**ComponentStateEntity**](ComponentStateEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
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
  **id** | **string**| The ID of the Controller Service | 
  **requestId** | **string**| The ID of the Verification Request | 

### Return type

[**VerifyConfigRequestEntity**](VerifyConfigRequestEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveControllerService**
> ControllerServiceEntity RemoveControllerService(ctx, id, optional)
Deletes a controller service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The controller service id. | 
 **optional** | ***ControllerServicesApiRemoveControllerServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ControllerServicesApiRemoveControllerServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | [**optional.Interface of LongParameter**](.md)| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | [**optional.Interface of ClientIdParameter**](.md)| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ControllerServiceEntity**](ControllerServiceEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitConfigVerificationRequest**
> VerifyConfigRequestEntity SubmitConfigVerificationRequest(ctx, body, id)
Performs verification of the Controller Service's configuration

This will initiate the process of verifying a given Controller Service configuration. This may be a long-running task. As a result, this endpoint will immediately return a ControllerServiceConfigVerificationRequestEntity, and the process of performing the verification will occur asynchronously in the background. The client may then periodically poll the status of the request by issuing a GET request to /controller-services/{serviceId}/verification-requests/{requestId}. Once the request is completed, the client is expected to issue a DELETE request to /controller-services/{serviceId}/verification-requests/{requestId}.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VerifyConfigRequestEntity**](VerifyConfigRequestEntity.md)| The controller service configuration verification request. | 
  **id** | **string**| The controller service id. | 

### Return type

[**VerifyConfigRequestEntity**](VerifyConfigRequestEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateControllerService**
> ControllerServiceEntity UpdateControllerService(ctx, body, id)
Updates a controller service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ControllerServiceEntity**](ControllerServiceEntity.md)| The controller service configuration details. | 
  **id** | **string**| The controller service id. | 

### Return type

[**ControllerServiceEntity**](ControllerServiceEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateControllerServiceReferences**
> ControllerServiceReferencingComponentsEntity UpdateControllerServiceReferences(ctx, body, id)
Updates a controller services references

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateControllerServiceReferenceRequestEntity**](UpdateControllerServiceReferenceRequestEntity.md)| The controller service request update request. | 
  **id** | **string**| The controller service id. | 

### Return type

[**ControllerServiceReferencingComponentsEntity**](ControllerServiceReferencingComponentsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRunStatus1**
> ControllerServiceEntity UpdateRunStatus1(ctx, body, id)
Updates run status of a controller service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ControllerServiceRunStatusEntity**](ControllerServiceRunStatusEntity.md)| The controller service run status. | 
  **id** | **string**| The controller service id. | 

### Return type

[**ControllerServiceEntity**](ControllerServiceEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

