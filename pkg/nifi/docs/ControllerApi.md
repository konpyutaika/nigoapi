# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyzeFlowAnalysisRuleConfiguration**](ControllerApi.md#AnalyzeFlowAnalysisRuleConfiguration) | **Post** /controller/flow-analysis-rules/{id}/config/analysis | Performs analysis of the component&#x27;s configuration, providing information about which attributes are referenced.
[**ClearState**](ControllerApi.md#ClearState) | **Post** /controller/flow-analysis-rules/{id}/state/clear-requests | Clears the state for a flow analysis rule
[**CreateBulletin**](ControllerApi.md#CreateBulletin) | **Post** /controller/bulletin | Creates a new bulletin
[**CreateControllerService**](ControllerApi.md#CreateControllerService) | **Post** /controller/controller-services | Creates a new controller service
[**CreateFlowAnalysisRule**](ControllerApi.md#CreateFlowAnalysisRule) | **Post** /controller/flow-analysis-rules | Creates a new flow analysis rule
[**CreateFlowRegistryClient**](ControllerApi.md#CreateFlowRegistryClient) | **Post** /controller/registry-clients | Creates a new flow registry client
[**CreateParameterProvider**](ControllerApi.md#CreateParameterProvider) | **Post** /controller/parameter-providers | Creates a new parameter provider
[**CreateReportingTask**](ControllerApi.md#CreateReportingTask) | **Post** /controller/reporting-tasks | Creates a new reporting task
[**DeleteFlowAnalysisRuleVerificationRequest**](ControllerApi.md#DeleteFlowAnalysisRuleVerificationRequest) | **Delete** /controller/flow-analysis-rules/{id}/config/verification-requests/{requestId} | Deletes the Verification Request with the given ID
[**DeleteFlowRegistryClient**](ControllerApi.md#DeleteFlowRegistryClient) | **Delete** /controller/registry-clients/{id} | Deletes a flow registry client
[**DeleteHistory**](ControllerApi.md#DeleteHistory) | **Delete** /controller/history | Purges history
[**DeleteNar**](ControllerApi.md#DeleteNar) | **Delete** /controller/nar-manager/nars/{id} | Deletes an installed NAR
[**DeleteNode**](ControllerApi.md#DeleteNode) | **Delete** /controller/cluster/nodes/{id} | Removes a node from the cluster
[**DownloadNar**](ControllerApi.md#DownloadNar) | **Get** /controller/nar-manager/nars/{id}/content | Retrieves the content of the NAR with the given id
[**GetCluster**](ControllerApi.md#GetCluster) | **Get** /controller/cluster | Gets the contents of the cluster
[**GetControllerConfig**](ControllerApi.md#GetControllerConfig) | **Get** /controller/config | Retrieves the configuration for this NiFi Controller
[**GetFlowAnalysisRule**](ControllerApi.md#GetFlowAnalysisRule) | **Get** /controller/flow-analysis-rules/{id} | Gets a flow analysis rule
[**GetFlowAnalysisRulePropertyDescriptor**](ControllerApi.md#GetFlowAnalysisRulePropertyDescriptor) | **Get** /controller/flow-analysis-rules/{id}/descriptors | Gets a flow analysis rule property descriptor
[**GetFlowAnalysisRuleState**](ControllerApi.md#GetFlowAnalysisRuleState) | **Get** /controller/flow-analysis-rules/{id}/state | Gets the state for a flow analysis rule
[**GetFlowAnalysisRuleVerificationRequest**](ControllerApi.md#GetFlowAnalysisRuleVerificationRequest) | **Get** /controller/flow-analysis-rules/{id}/config/verification-requests/{requestId} | Returns the Verification Request with the given ID
[**GetFlowAnalysisRules**](ControllerApi.md#GetFlowAnalysisRules) | **Get** /controller/flow-analysis-rules | Gets all flow analysis rules
[**GetFlowRegistryClient**](ControllerApi.md#GetFlowRegistryClient) | **Get** /controller/registry-clients/{id} | Gets a flow registry client
[**GetFlowRegistryClients**](ControllerApi.md#GetFlowRegistryClients) | **Get** /controller/registry-clients | Gets the listing of available flow registry clients
[**GetNarDetails**](ControllerApi.md#GetNarDetails) | **Get** /controller/nar-manager/nars/{id}/details | Retrieves the component types available from the installed NARs
[**GetNarSummaries**](ControllerApi.md#GetNarSummaries) | **Get** /controller/nar-manager/nars | Retrieves summary information for installed NARs
[**GetNarSummary**](ControllerApi.md#GetNarSummary) | **Get** /controller/nar-manager/nars/{id} | Retrieves the summary information for the NAR with the given identifier
[**GetNode**](ControllerApi.md#GetNode) | **Get** /controller/cluster/nodes/{id} | Gets a node in the cluster
[**GetNodeStatusHistory**](ControllerApi.md#GetNodeStatusHistory) | **Get** /controller/status/history | Gets status history for the node
[**GetPropertyDescriptor**](ControllerApi.md#GetPropertyDescriptor) | **Get** /controller/registry-clients/{id}/descriptors | Gets a flow registry client property descriptor
[**GetRegistryClientTypes**](ControllerApi.md#GetRegistryClientTypes) | **Get** /controller/registry-types | Retrieves the types of flow  that this NiFi supports
[**ImportReportingTaskSnapshot**](ControllerApi.md#ImportReportingTaskSnapshot) | **Post** /controller/reporting-tasks/import | Imports a reporting task snapshot
[**RemoveFlowAnalysisRule**](ControllerApi.md#RemoveFlowAnalysisRule) | **Delete** /controller/flow-analysis-rules/{id} | Deletes a flow analysis rule
[**SubmitFlowAnalysisRuleConfigVerificationRequest**](ControllerApi.md#SubmitFlowAnalysisRuleConfigVerificationRequest) | **Post** /controller/flow-analysis-rules/{id}/config/verification-requests | Performs verification of the Flow Analysis Rule&#x27;s configuration
[**UpdateControllerConfig**](ControllerApi.md#UpdateControllerConfig) | **Put** /controller/config | Retrieves the configuration for this NiFi
[**UpdateFlowAnalysisRule**](ControllerApi.md#UpdateFlowAnalysisRule) | **Put** /controller/flow-analysis-rules/{id} | Updates a flow analysis rule
[**UpdateFlowRegistryClient**](ControllerApi.md#UpdateFlowRegistryClient) | **Put** /controller/registry-clients/{id} | Updates a flow registry client
[**UpdateNode**](ControllerApi.md#UpdateNode) | **Put** /controller/cluster/nodes/{id} | Updates a node in the cluster
[**UpdateRunStatus**](ControllerApi.md#UpdateRunStatus) | **Put** /controller/flow-analysis-rules/{id}/run-status | Updates run status of a flow analysis rule
[**UploadNar**](ControllerApi.md#UploadNar) | **Post** /controller/nar-manager/nars/content | Uploads a NAR and requests for it to be installed

# **AnalyzeFlowAnalysisRuleConfiguration**
> ConfigurationAnalysisEntity AnalyzeFlowAnalysisRuleConfiguration(ctx, body, id)
Performs analysis of the component's configuration, providing information about which attributes are referenced.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConfigurationAnalysisEntity**](ConfigurationAnalysisEntity.md)| The configuration analysis request. | 
  **id** | **string**| The flow analysis rules id. | 

### Return type

[**ConfigurationAnalysisEntity**](ConfigurationAnalysisEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearState**
> ComponentStateEntity ClearState(ctx, id, optional)
Clears the state for a flow analysis rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The flow analysis rule id. | 
 **optional** | ***ControllerApiClearStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ControllerApiClearStateOpts struct
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

# **CreateBulletin**
> BulletinEntity CreateBulletin(ctx, body)
Creates a new bulletin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulletinEntity**](BulletinEntity.md)| The reporting task configuration details. | 

### Return type

[**BulletinEntity**](BulletinEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateControllerService**
> ControllerServiceEntity CreateControllerService(ctx, body)
Creates a new controller service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ControllerServiceEntity**](ControllerServiceEntity.md)| The controller service configuration details. | 

### Return type

[**ControllerServiceEntity**](ControllerServiceEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFlowAnalysisRule**
> FlowAnalysisRuleEntity CreateFlowAnalysisRule(ctx, body)
Creates a new flow analysis rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FlowAnalysisRuleEntity**](FlowAnalysisRuleEntity.md)| The flow analysis rule configuration details. | 

### Return type

[**FlowAnalysisRuleEntity**](FlowAnalysisRuleEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFlowRegistryClient**
> FlowRegistryClientEntity CreateFlowRegistryClient(ctx, body)
Creates a new flow registry client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FlowRegistryClientEntity**](FlowRegistryClientEntity.md)| The flow registry client configuration details. | 

### Return type

[**FlowRegistryClientEntity**](FlowRegistryClientEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateParameterProvider**
> ParameterProviderEntity CreateParameterProvider(ctx, body)
Creates a new parameter provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ParameterProviderEntity**](ParameterProviderEntity.md)| The parameter provider configuration details. | 

### Return type

[**ParameterProviderEntity**](ParameterProviderEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateReportingTask**
> ReportingTaskEntity CreateReportingTask(ctx, body)
Creates a new reporting task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReportingTaskEntity**](ReportingTaskEntity.md)| The reporting task configuration details. | 

### Return type

[**ReportingTaskEntity**](ReportingTaskEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFlowAnalysisRuleVerificationRequest**
> VerifyConfigRequestEntity DeleteFlowAnalysisRuleVerificationRequest(ctx, id, requestId)
Deletes the Verification Request with the given ID

Deletes the Verification Request with the given ID. After a request is created, it is expected that the client will properly clean up the request by DELETE'ing it, once the Verification process has completed. If the request is deleted before the request completes, then the Verification request will finish the step that it is currently performing and then will cancel any subsequent steps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the Flow Analysis Rule | 
  **requestId** | **string**| The ID of the Verification Request | 

### Return type

[**VerifyConfigRequestEntity**](VerifyConfigRequestEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFlowRegistryClient**
> FlowRegistryClientEntity DeleteFlowRegistryClient(ctx, id, optional)
Deletes a flow registry client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The flow registry client id. | 
 **optional** | ***ControllerApiDeleteFlowRegistryClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ControllerApiDeleteFlowRegistryClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | [**optional.Interface of LongParameter**](.md)| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | [**optional.Interface of ClientIdParameter**](.md)| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**FlowRegistryClientEntity**](FlowRegistryClientEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHistory**
> HistoryEntity DeleteHistory(ctx, endDate)
Purges history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endDate** | [**DateTimeParameter**](.md)| Purge actions before this date/time. | 

### Return type

[**HistoryEntity**](HistoryEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNar**
> NarSummaryEntity DeleteNar(ctx, id, optional)
Deletes an installed NAR

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the NAR. | 
 **optional** | ***ControllerApiDeleteNarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ControllerApiDeleteNarOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disconnectedNodeAcknowledged** | **optional.bool**|  | [default to false]
 **force** | **optional.bool**|  | [default to false]

### Return type

[**NarSummaryEntity**](NarSummaryEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNode**
> NodeEntity DeleteNode(ctx, id)
Removes a node from the cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The node id. | 

### Return type

[**NodeEntity**](NodeEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadNar**
> string DownloadNar(ctx, id)
Retrieves the content of the NAR with the given id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the NAR. | 

### Return type

**string**

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCluster**
> ClusterEntity GetCluster(ctx, )
Gets the contents of the cluster

Returns the contents of the cluster including all nodes and their status.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterEntity**](ClusterEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetControllerConfig**
> ControllerConfigurationEntity GetControllerConfig(ctx, )
Retrieves the configuration for this NiFi Controller

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ControllerConfigurationEntity**](ControllerConfigurationEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlowAnalysisRule**
> FlowAnalysisRuleEntity GetFlowAnalysisRule(ctx, id)
Gets a flow analysis rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The flow analysis rule id. | 

### Return type

[**FlowAnalysisRuleEntity**](FlowAnalysisRuleEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlowAnalysisRulePropertyDescriptor**
> PropertyDescriptorEntity GetFlowAnalysisRulePropertyDescriptor(ctx, id, propertyName, optional)
Gets a flow analysis rule property descriptor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The flow analysis rule id. | 
  **propertyName** | **string**| The property name. | 
 **optional** | ***ControllerApiGetFlowAnalysisRulePropertyDescriptorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ControllerApiGetFlowAnalysisRulePropertyDescriptorOpts struct
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

# **GetFlowAnalysisRuleState**
> ComponentStateEntity GetFlowAnalysisRuleState(ctx, id)
Gets the state for a flow analysis rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The flow analysis rule id. | 

### Return type

[**ComponentStateEntity**](ComponentStateEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlowAnalysisRuleVerificationRequest**
> VerifyConfigRequestEntity GetFlowAnalysisRuleVerificationRequest(ctx, id, requestId)
Returns the Verification Request with the given ID

Returns the Verification Request with the given ID. Once an Verification Request has been created, that request can subsequently be retrieved via this endpoint, and the request that is fetched will contain the updated state, such as percent complete, the current state of the request, and any failures. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the Flow Analysis Rule | 
  **requestId** | **string**| The ID of the Verification Request | 

### Return type

[**VerifyConfigRequestEntity**](VerifyConfigRequestEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlowAnalysisRules**
> FlowAnalysisRulesEntity GetFlowAnalysisRules(ctx, )
Gets all flow analysis rules

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FlowAnalysisRulesEntity**](FlowAnalysisRulesEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlowRegistryClient**
> FlowRegistryClientEntity GetFlowRegistryClient(ctx, id)
Gets a flow registry client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The flow registry client id. | 

### Return type

[**FlowRegistryClientEntity**](FlowRegistryClientEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlowRegistryClients**
> FlowRegistryClientsEntity GetFlowRegistryClients(ctx, )
Gets the listing of available flow registry clients

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FlowRegistryClientsEntity**](FlowRegistryClientsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNarDetails**
> NarDetailsEntity GetNarDetails(ctx, id)
Retrieves the component types available from the installed NARs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the NAR. | 

### Return type

[**NarDetailsEntity**](NarDetailsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNarSummaries**
> NarSummariesEntity GetNarSummaries(ctx, )
Retrieves summary information for installed NARs

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NarSummariesEntity**](NarSummariesEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNarSummary**
> NarDetailsEntity GetNarSummary(ctx, id)
Retrieves the summary information for the NAR with the given identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the NAR. | 

### Return type

[**NarDetailsEntity**](NarDetailsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNode**
> NodeEntity GetNode(ctx, id)
Gets a node in the cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The node id. | 

### Return type

[**NodeEntity**](NodeEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeStatusHistory**
> ComponentHistoryEntity GetNodeStatusHistory(ctx, )
Gets status history for the node

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ComponentHistoryEntity**](ComponentHistoryEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPropertyDescriptor**
> PropertyDescriptorEntity GetPropertyDescriptor(ctx, id, propertyName, optional)
Gets a flow registry client property descriptor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The flow registry client id. | 
  **propertyName** | **string**| The property name. | 
 **optional** | ***ControllerApiGetPropertyDescriptorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ControllerApiGetPropertyDescriptorOpts struct
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

# **GetRegistryClientTypes**
> FlowRegistryClientTypesEntity GetRegistryClientTypes(ctx, )
Retrieves the types of flow  that this NiFi supports

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FlowRegistryClientTypesEntity**](FlowRegistryClientTypesEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportReportingTaskSnapshot**
> VersionedReportingTaskImportResponseEntity ImportReportingTaskSnapshot(ctx, body)
Imports a reporting task snapshot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VersionedReportingTaskImportRequestEntity**](VersionedReportingTaskImportRequestEntity.md)| The import request containing the reporting task snapshot to import. | 

### Return type

[**VersionedReportingTaskImportResponseEntity**](VersionedReportingTaskImportResponseEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveFlowAnalysisRule**
> FlowAnalysisRuleEntity RemoveFlowAnalysisRule(ctx, id, optional)
Deletes a flow analysis rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The flow analysis rule id. | 
 **optional** | ***ControllerApiRemoveFlowAnalysisRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ControllerApiRemoveFlowAnalysisRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | [**optional.Interface of LongParameter**](.md)| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | [**optional.Interface of ClientIdParameter**](.md)| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**FlowAnalysisRuleEntity**](FlowAnalysisRuleEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitFlowAnalysisRuleConfigVerificationRequest**
> VerifyConfigRequestEntity SubmitFlowAnalysisRuleConfigVerificationRequest(ctx, body, id)
Performs verification of the Flow Analysis Rule's configuration

This will initiate the process of verifying a given Flow Analysis Rule configuration. This may be a long-running task. As a result, this endpoint will immediately return a FlowAnalysisRuleConfigVerificationRequestEntity, and the process of performing the verification will occur asynchronously in the background. The client may then periodically poll the status of the request by issuing a GET request to /flow-analysis-rules/{taskId}/verification-requests/{requestId}. Once the request is completed, the client is expected to issue a DELETE request to /flow-analysis-rules/{serviceId}/verification-requests/{requestId}.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VerifyConfigRequestEntity**](VerifyConfigRequestEntity.md)| The flow analysis rules configuration verification request. | 
  **id** | **string**| The flow analysis rules id. | 

### Return type

[**VerifyConfigRequestEntity**](VerifyConfigRequestEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateControllerConfig**
> ControllerConfigurationEntity UpdateControllerConfig(ctx, body)
Retrieves the configuration for this NiFi

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ControllerConfigurationEntity**](ControllerConfigurationEntity.md)| The controller configuration. | 

### Return type

[**ControllerConfigurationEntity**](ControllerConfigurationEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFlowAnalysisRule**
> FlowAnalysisRuleEntity UpdateFlowAnalysisRule(ctx, body, id)
Updates a flow analysis rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FlowAnalysisRuleEntity**](FlowAnalysisRuleEntity.md)| The flow analysis rule configuration details. | 
  **id** | **string**| The flow analysis rule id. | 

### Return type

[**FlowAnalysisRuleEntity**](FlowAnalysisRuleEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFlowRegistryClient**
> FlowRegistryClientEntity UpdateFlowRegistryClient(ctx, body, id)
Updates a flow registry client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FlowRegistryClientEntity**](FlowRegistryClientEntity.md)| The flow registry client configuration details. | 
  **id** | **string**| The flow registry client id. | 

### Return type

[**FlowRegistryClientEntity**](FlowRegistryClientEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNode**
> NodeEntity UpdateNode(ctx, body, id)
Updates a node in the cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodeEntity**](NodeEntity.md)| The node configuration. The only configuration that will be honored at this endpoint is the status. | 
  **id** | **string**| The node id. | 

### Return type

[**NodeEntity**](NodeEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRunStatus**
> FlowAnalysisRuleEntity UpdateRunStatus(ctx, body, id)
Updates run status of a flow analysis rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FlowAnalysisRuleRunStatusEntity**](FlowAnalysisRuleRunStatusEntity.md)| The flow analysis rule run status. | 
  **id** | **string**| The flow analysis rule id. | 

### Return type

[**FlowAnalysisRuleEntity**](FlowAnalysisRuleEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadNar**
> NarSummaryEntity UploadNar(ctx, body, optional)
Uploads a NAR and requests for it to be installed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)| The contents of the NAR file. | 
 **optional** | ***ControllerApiUploadNarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ControllerApiUploadNarOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filename** | **optional.string**|  | 

### Return type

[**NarSummaryEntity**](NarSummaryEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

