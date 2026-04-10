# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyConnectorUpdate**](ConnectorsApi.md#ApplyConnectorUpdate) | **Post** /connectors/{id}/apply-update | Applies an update to a connector
[**CancelDrain**](ConnectorsApi.md#CancelDrain) | **Delete** /connectors/{id}/drain | Cancels the draining of FlowFiles for a connector
[**ClearConnectorControllerServiceState**](ConnectorsApi.md#ClearConnectorControllerServiceState) | **Post** /connectors/{id}/controller-services/{controllerServiceId}/state/clear-requests | Clears the state for a controller service within a connector
[**ClearConnectorProcessorState**](ConnectorsApi.md#ClearConnectorProcessorState) | **Post** /connectors/{id}/processors/{processorId}/state/clear-requests | Clears the state for a processor within a connector
[**CreateAsset**](ConnectorsApi.md#CreateAsset) | **Post** /connectors/{id}/assets | Creates a new Asset in the given Connector
[**CreateConnector**](ConnectorsApi.md#CreateConnector) | **Post** /connectors | Creates a new connector
[**CreatePurgeRequest**](ConnectorsApi.md#CreatePurgeRequest) | **Post** /connectors/{id}/purge-requests | Creates a request to purge the FlowFiles for this connector
[**DeleteConfigurationStepVerificationRequest**](ConnectorsApi.md#DeleteConfigurationStepVerificationRequest) | **Delete** /connectors/{id}/configuration-steps/{configurationStepName}/verify-config/{requestId} | Deletes the Verification Request with the given ID
[**DeleteConnector**](ConnectorsApi.md#DeleteConnector) | **Delete** /connectors/{id} | Deletes a connector
[**DiscardConnectorUpdate**](ConnectorsApi.md#DiscardConnectorUpdate) | **Delete** /connectors/{id}/working-configuration | Discards the working configuration of a connector
[**GetAssetContent**](ConnectorsApi.md#GetAssetContent) | **Get** /connectors/{id}/assets/{assetId} | Retrieves the content of the asset with the given id for a connector
[**GetAssets**](ConnectorsApi.md#GetAssets) | **Get** /connectors/{id}/assets | Lists the assets that belong to the Connector with the given ID
[**GetConfigurationStepVerificationRequest**](ConnectorsApi.md#GetConfigurationStepVerificationRequest) | **Get** /connectors/{id}/configuration-steps/{configurationStepName}/verify-config/{requestId} | Returns the Verification Request with the given ID
[**GetConnector**](ConnectorsApi.md#GetConnector) | **Get** /connectors/{id} | Gets a connector
[**GetConnectorConfigurationStep**](ConnectorsApi.md#GetConnectorConfigurationStep) | **Get** /connectors/{id}/configuration-steps/{configurationStepName} | Gets a specific configuration step by name for a connector
[**GetConnectorConfigurationSteps**](ConnectorsApi.md#GetConnectorConfigurationSteps) | **Get** /connectors/{id}/configuration-steps | Gets all configuration step names for a connector
[**GetConnectorControllerServiceState**](ConnectorsApi.md#GetConnectorControllerServiceState) | **Get** /connectors/{id}/controller-services/{controllerServiceId}/state | Gets the state for a controller service within a connector
[**GetConnectorProcessorState**](ConnectorsApi.md#GetConnectorProcessorState) | **Get** /connectors/{id}/processors/{processorId}/state | Gets the state for a processor within a connector
[**GetConnectorPropertyAllowableValues**](ConnectorsApi.md#GetConnectorPropertyAllowableValues) | **Get** /connectors/{id}/configuration-steps/{configurationStepName}/property-groups/{propertyGroupName}/properties/{propertyName}/allowable-values | Gets the allowable values for a specific property in a connector&#x27;s configuration step
[**GetConnectorStatus**](ConnectorsApi.md#GetConnectorStatus) | **Get** /connectors/{id}/status | Gets the status for the process group managed by a connector
[**GetControllerServicesFromConnectorProcessGroup**](ConnectorsApi.md#GetControllerServicesFromConnectorProcessGroup) | **Get** /connectors/{connectorId}/flow/process-groups/{processGroupId}/controller-services | Gets all controller services for a process group within a connector
[**GetFlow**](ConnectorsApi.md#GetFlow) | **Get** /connectors/{connectorId}/flow/process-groups/{processGroupId} | Gets the flow for a process group within a connector
[**GetPurgeRequest**](ConnectorsApi.md#GetPurgeRequest) | **Get** /connectors/{id}/purge-requests/{purge-request-id} | Gets the current status of a purge request for the specified connector
[**GetSecrets**](ConnectorsApi.md#GetSecrets) | **Get** /connectors/{id}/secrets | Gets all secrets available for configuring a connector
[**InitiateDrain**](ConnectorsApi.md#InitiateDrain) | **Post** /connectors/{id}/drain | Initiates draining of FlowFiles for a connector
[**RemovePurgeRequest**](ConnectorsApi.md#RemovePurgeRequest) | **Delete** /connectors/{id}/purge-requests/{purge-request-id} | Cancels and/or removes a request to purge the FlowFiles for this connector
[**SearchConnector**](ConnectorsApi.md#SearchConnector) | **Get** /connectors/{id}/search-results | Performs a search against the encapsulated process group of this connector using the specified search term
[**SubmitConfigurationStepVerificationRequest**](ConnectorsApi.md#SubmitConfigurationStepVerificationRequest) | **Post** /connectors/{id}/configuration-steps/{configurationStepName}/verify-config | Performs verification of a configuration step for a connector
[**UpdateConnector**](ConnectorsApi.md#UpdateConnector) | **Put** /connectors/{id} | Updates a connector
[**UpdateConnectorConfigurationStep**](ConnectorsApi.md#UpdateConnectorConfigurationStep) | **Put** /connectors/{id}/configuration-steps/{configurationStepName} | Updates a specific configuration step by name for a connector
[**UpdateRunStatus**](ConnectorsApi.md#UpdateRunStatus) | **Put** /connectors/{id}/run-status | Updates run status of a connector

# **ApplyConnectorUpdate**
> ConnectorEntity ApplyConnectorUpdate(ctx, body, id)
Applies an update to a connector

This will apply any pending configuration changes to the connector. The client can poll the connector endpoint to check when the update is complete.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectorEntity**](ConnectorEntity.md)| The connector configuration with revision. | 
  **id** | **string**| The connector id. | 

### Return type

[**ConnectorEntity**](ConnectorEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelDrain**
> ConnectorEntity CancelDrain(ctx, id, optional)
Cancels the draining of FlowFiles for a connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 
 **optional** | ***ConnectorsApiCancelDrainOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiCancelDrainOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | [**optional.Interface of LongParameter**](.md)| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | [**optional.Interface of ClientIdParameter**](.md)| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ConnectorEntity**](ConnectorEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearConnectorControllerServiceState**
> ComponentStateEntity ClearConnectorControllerServiceState(ctx, id, controllerServiceId, optional)
Clears the state for a controller service within a connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 
  **controllerServiceId** | **string**| The controller service id. | 
 **optional** | ***ConnectorsApiClearConnectorControllerServiceStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiClearConnectorControllerServiceStateOpts struct
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

# **ClearConnectorProcessorState**
> ComponentStateEntity ClearConnectorProcessorState(ctx, id, processorId, optional)
Clears the state for a processor within a connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 
  **processorId** | **string**| The processor id. | 
 **optional** | ***ConnectorsApiClearConnectorProcessorStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiClearConnectorProcessorStateOpts struct
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

# **CreateAsset**
> AssetEntity CreateAsset(ctx, body, id, optional)
Creates a new Asset in the given Connector

This endpoint will create a new Asset in the Connector. The Asset will be created with the given name and the contents of the file that is uploaded.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)| The contents of the asset | 
  **id** | **string**|  | 
 **optional** | ***ConnectorsApiCreateAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiCreateAssetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filename** | **optional.string**|  | 

### Return type

[**AssetEntity**](AssetEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConnector**
> ConnectorEntity CreateConnector(ctx, body)
Creates a new connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectorEntity**](ConnectorEntity.md)| The connector configuration details. | 

### Return type

[**ConnectorEntity**](ConnectorEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePurgeRequest**
> DropRequestEntity CreatePurgeRequest(ctx, id)
Creates a request to purge the FlowFiles for this connector

This will create a request to purge all FlowFiles from the connector. The connector must be in a STOPPED state before purging can begin. This is an asynchronous operation. The client should poll the returned URI to get the status of the purge request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 

### Return type

[**DropRequestEntity**](DropRequestEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConfigurationStepVerificationRequest**
> VerifyConnectorConfigStepRequestEntity DeleteConfigurationStepVerificationRequest(ctx, id, configurationStepName, requestId)
Deletes the Verification Request with the given ID

Deletes the Verification Request with the given ID. After a request is created, it is expected that the client will properly clean up the request by DELETE'ing it, once the Verification process has completed. If the request is deleted before the request completes, then the Verification request will finish the step that it is currently performing and then will cancel any subsequent steps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 
  **configurationStepName** | **string**| The configuration step name. | 
  **requestId** | **string**| The ID of the Verification Request | 

### Return type

[**VerifyConnectorConfigStepRequestEntity**](VerifyConnectorConfigStepRequestEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConnector**
> ConnectorEntity DeleteConnector(ctx, id, optional)
Deletes a connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 
 **optional** | ***ConnectorsApiDeleteConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiDeleteConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | [**optional.Interface of LongParameter**](.md)| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | [**optional.Interface of ClientIdParameter**](.md)| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ConnectorEntity**](ConnectorEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscardConnectorUpdate**
> ConnectorEntity DiscardConnectorUpdate(ctx, id, optional)
Discards the working configuration of a connector

This will discard any pending configuration changes for the connector and revert to the last applied configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 
 **optional** | ***ConnectorsApiDiscardConnectorUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiDiscardConnectorUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | [**optional.Interface of LongParameter**](.md)| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | [**optional.Interface of ClientIdParameter**](.md)| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**ConnectorEntity**](ConnectorEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAssetContent**
> string GetAssetContent(ctx, id, assetId)
Retrieves the content of the asset with the given id for a connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 
  **assetId** | **string**| The asset id. | 

### Return type

**string**

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAssets**
> AssetsEntity GetAssets(ctx, id)
Lists the assets that belong to the Connector with the given ID

Lists the assets that belong to the Connector with the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 

### Return type

[**AssetsEntity**](AssetsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigurationStepVerificationRequest**
> VerifyConnectorConfigStepRequestEntity GetConfigurationStepVerificationRequest(ctx, id, configurationStepName, requestId)
Returns the Verification Request with the given ID

Returns the Verification Request with the given ID. Once a Verification Request has been created, that request can subsequently be retrieved via this endpoint, and the request that is fetched will contain the updated state, such as percent complete, the current state of the request, and any failures.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 
  **configurationStepName** | **string**| The configuration step name. | 
  **requestId** | **string**| The ID of the Verification Request | 

### Return type

[**VerifyConnectorConfigStepRequestEntity**](VerifyConnectorConfigStepRequestEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnector**
> ConnectorEntity GetConnector(ctx, id)
Gets a connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 

### Return type

[**ConnectorEntity**](ConnectorEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorConfigurationStep**
> ConfigurationStepEntity GetConnectorConfigurationStep(ctx, id, configurationStepName)
Gets a specific configuration step by name for a connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 
  **configurationStepName** | **string**| The configuration step name. | 

### Return type

[**ConfigurationStepEntity**](ConfigurationStepEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorConfigurationSteps**
> ConfigurationStepNamesEntity GetConnectorConfigurationSteps(ctx, id)
Gets all configuration step names for a connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 

### Return type

[**ConfigurationStepNamesEntity**](ConfigurationStepNamesEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorControllerServiceState**
> ComponentStateEntity GetConnectorControllerServiceState(ctx, id, controllerServiceId)
Gets the state for a controller service within a connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 
  **controllerServiceId** | **string**| The controller service id. | 

### Return type

[**ComponentStateEntity**](ComponentStateEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorProcessorState**
> ComponentStateEntity GetConnectorProcessorState(ctx, id, processorId)
Gets the state for a processor within a connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 
  **processorId** | **string**| The processor id. | 

### Return type

[**ComponentStateEntity**](ComponentStateEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorPropertyAllowableValues**
> ConnectorPropertyAllowableValuesEntity GetConnectorPropertyAllowableValues(ctx, id, configurationStepName, propertyGroupName, propertyName, optional)
Gets the allowable values for a specific property in a connector's configuration step

Gets the allowable values for a specific property that supports dynamic fetching of allowable values. The filter parameter can be used to narrow down the results based on the property's filtering logic.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 
  **configurationStepName** | **string**| The configuration step name. | 
  **propertyGroupName** | **string**| The property group name. | 
  **propertyName** | **string**| The property name. | 
 **optional** | ***ConnectorsApiGetConnectorPropertyAllowableValuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetConnectorPropertyAllowableValuesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **filter** | **optional.string**| Optional filter to narrow down the allowable values. | 

### Return type

[**ConnectorPropertyAllowableValuesEntity**](ConnectorPropertyAllowableValuesEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectorStatus**
> ProcessGroupStatusEntity GetConnectorStatus(ctx, id, optional)
Gets the status for the process group managed by a connector

Returns the status for the process group managed by the specified connector. The status includes status for all descendent components. When invoked with recursive set to true, it will return the current status of every component in the connector's encapsulated flow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 
 **optional** | ***ConnectorsApiGetConnectorStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetConnectorStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recursive** | **optional.bool**| Whether all descendant groups and the status of their content will be included. Optional, defaults to false | [default to false]
 **nodewise** | **optional.bool**| Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **optional.string**| The id of the node where to get the status. | 

### Return type

[**ProcessGroupStatusEntity**](ProcessGroupStatusEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetControllerServicesFromConnectorProcessGroup**
> ControllerServicesEntity GetControllerServicesFromConnectorProcessGroup(ctx, connectorId, processGroupId, optional)
Gets all controller services for a process group within a connector

Returns the controller services for the specified process group within the connector's hierarchy. The processGroupId can be obtained from the managedProcessGroupId field of the ConnectorDTO for the root process group, or from child process groups within the flow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectorId** | **string**| The connector id. | 
  **processGroupId** | **string**| The process group id. | 
 **optional** | ***ConnectorsApiGetControllerServicesFromConnectorProcessGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetControllerServicesFromConnectorProcessGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeAncestorGroups** | **optional.bool**| Whether or not to include parent/ancestor process groups | [default to true]
 **includeDescendantGroups** | **optional.bool**| Whether or not to include descendant process groups | [default to false]
 **includeReferencingComponents** | **optional.bool**| Whether or not to include services&#x27; referencing components in the response | [default to true]

### Return type

[**ControllerServicesEntity**](ControllerServicesEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlow**
> ProcessGroupFlowEntity GetFlow(ctx, connectorId, processGroupId, optional)
Gets the flow for a process group within a connector

Returns the flow for the specified process group within the connector's hierarchy. The processGroupId can be obtained from the managedProcessGroupId field of the ConnectorDTO for the root process group, or from child process groups within the flow. If the uiOnly query parameter is provided with a value of true, the returned entity may only contain fields that are necessary for rendering the NiFi User Interface. As such, the selected fields may change at any time, even during incremental releases, without warning. As a result, this parameter should not be provided by any client other than the UI.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectorId** | **string**| The connector id. | 
  **processGroupId** | **string**| The process group id. | 
 **optional** | ***ConnectorsApiGetFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiGetFlowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uiOnly** | **optional.bool**| Whether to return only UI-specific fields | [default to false]

### Return type

[**ProcessGroupFlowEntity**](ProcessGroupFlowEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPurgeRequest**
> DropRequestEntity GetPurgeRequest(ctx, id, purgeRequestId)
Gets the current status of a purge request for the specified connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 
  **purgeRequestId** | **string**| The purge request id. | 

### Return type

[**DropRequestEntity**](DropRequestEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecrets**
> SecretsEntity GetSecrets(ctx, id)
Gets all secrets available for configuring a connector

Returns metadata for all secrets available from all secret providers. This endpoint is used when configuring a connector to discover available secrets. Note: Actual secret values are not included in the response for security reasons.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 

### Return type

[**SecretsEntity**](SecretsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InitiateDrain**
> ConnectorEntity InitiateDrain(ctx, body, id)
Initiates draining of FlowFiles for a connector

This will initiate draining of FlowFiles for a stopped connector. Draining allows the connector to process data that is currently in the flow but does not ingest any additional data. The connector must be in a STOPPED state before draining can begin. Once initiated, the connector will transition to a DRAINING state. Use the DELETE method on this endpoint to cancel an ongoing drain operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectorEntity**](ConnectorEntity.md)| The connector entity with revision. | 
  **id** | **string**| The connector id. | 

### Return type

[**ConnectorEntity**](ConnectorEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemovePurgeRequest**
> DropRequestEntity RemovePurgeRequest(ctx, id, purgeRequestId)
Cancels and/or removes a request to purge the FlowFiles for this connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 
  **purgeRequestId** | **string**| The purge request id. | 

### Return type

[**DropRequestEntity**](DropRequestEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchConnector**
> SearchResultsEntity SearchConnector(ctx, id, optional)
Performs a search against the encapsulated process group of this connector using the specified search term

Only search results from authorized components will be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connector id. | 
 **optional** | ***ConnectorsApiSearchConnectorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConnectorsApiSearchConnectorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **optional.string**| The search term. | 

### Return type

[**SearchResultsEntity**](SearchResultsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitConfigurationStepVerificationRequest**
> VerifyConnectorConfigStepRequestEntity SubmitConfigurationStepVerificationRequest(ctx, body, id, configurationStepName)
Performs verification of a configuration step for a connector

This will initiate the process of verifying a given Connector Configuration Step. This may be a long-running task. As a result, this endpoint will immediately return a VerifyConnectorConfigStepRequestEntity, and the process of performing the verification will occur asynchronously in the background. The client may then periodically poll the status of the request by issuing a GET request to /connectors/{connectorId}/configuration-steps/{stepName}/verify-config/{requestId}. Once the request is completed, the client is expected to issue a DELETE request to /connectors/{connectorId}/configuration-steps/{stepName}/verify-config/{requestId}.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VerifyConnectorConfigStepRequestEntity**](VerifyConnectorConfigStepRequestEntity.md)| The verify config request entity containing the configuration step to verify. | 
  **id** | **string**| The connector id. | 
  **configurationStepName** | **string**| The configuration step name. | 

### Return type

[**VerifyConnectorConfigStepRequestEntity**](VerifyConnectorConfigStepRequestEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConnector**
> ConnectorEntity UpdateConnector(ctx, body, id)
Updates a connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectorEntity**](ConnectorEntity.md)| The connector configuration details. | 
  **id** | **string**| The connector id. | 

### Return type

[**ConnectorEntity**](ConnectorEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConnectorConfigurationStep**
> ConfigurationStepEntity UpdateConnectorConfigurationStep(ctx, body, id, configurationStepName)
Updates a specific configuration step by name for a connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConfigurationStepEntity**](ConfigurationStepEntity.md)| The configuration step configuration. | 
  **id** | **string**| The connector id. | 
  **configurationStepName** | **string**| The configuration step name. | 

### Return type

[**ConfigurationStepEntity**](ConfigurationStepEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRunStatus**
> ConnectorEntity UpdateRunStatus(ctx, body, id)
Updates run status of a connector

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectorRunStatusEntity**](ConnectorRunStatusEntity.md)| The connector run status. | 
  **id** | **string**| The connector id. | 

### Return type

[**ConnectorEntity**](ConnectorEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

