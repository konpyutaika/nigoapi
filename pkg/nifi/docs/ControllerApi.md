# \ControllerApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBulletin**](ControllerApi.md#CreateBulletin) | **Post** /controller/bulletin | Creates a new bulletin
[**CreateControllerService**](ControllerApi.md#CreateControllerService) | **Post** /controller/controller-services | Creates a new controller service
[**CreateFlowRegistryClient**](ControllerApi.md#CreateFlowRegistryClient) | **Post** /controller/registry-clients | Creates a new flow registry client
[**CreateParameterProvider**](ControllerApi.md#CreateParameterProvider) | **Post** /controller/parameter-providers | Creates a new parameter provider
[**CreateReportingTask**](ControllerApi.md#CreateReportingTask) | **Post** /controller/reporting-tasks | Creates a new reporting task
[**DeleteFlowRegistryClient**](ControllerApi.md#DeleteFlowRegistryClient) | **Delete** /controller/registry-clients/{id} | Deletes a flow registry client
[**DeleteHistory**](ControllerApi.md#DeleteHistory) | **Delete** /controller/history | Purges history
[**DeleteNode**](ControllerApi.md#DeleteNode) | **Delete** /controller/cluster/nodes/{id} | Removes a node from the cluster
[**GetCluster**](ControllerApi.md#GetCluster) | **Get** /controller/cluster | Gets the contents of the cluster
[**GetControllerConfig**](ControllerApi.md#GetControllerConfig) | **Get** /controller/config | Retrieves the configuration for this NiFi Controller
[**GetFlowRegistryClient**](ControllerApi.md#GetFlowRegistryClient) | **Get** /controller/registry-clients/{id} | Gets a flow registry client
[**GetFlowRegistryClients**](ControllerApi.md#GetFlowRegistryClients) | **Get** /controller/registry-clients | Gets the listing of available flow registry clients
[**GetNode**](ControllerApi.md#GetNode) | **Get** /controller/cluster/nodes/{id} | Gets a node in the cluster
[**GetNodeStatusHistory**](ControllerApi.md#GetNodeStatusHistory) | **Get** /controller/status/history | Gets status history for the node
[**GetPropertyDescriptor**](ControllerApi.md#GetPropertyDescriptor) | **Get** /controller/registry-clients/{id}/descriptors | Gets a flow registry client property descriptor
[**GetRegistryClientTypes**](ControllerApi.md#GetRegistryClientTypes) | **Get** /controller/registry-types | Retrieves the types of flow  that this NiFi supports
[**UpdateControllerConfig**](ControllerApi.md#UpdateControllerConfig) | **Put** /controller/config | Retrieves the configuration for this NiFi
[**UpdateFlowRegistryClient**](ControllerApi.md#UpdateFlowRegistryClient) | **Put** /controller/registry-clients/{id} | Updates a flow registry client
[**UpdateNode**](ControllerApi.md#UpdateNode) | **Put** /controller/cluster/nodes/{id} | Updates a node in the cluster


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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

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

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
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

 **version** | **optional.String**| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | **optional.String**| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.Bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**FlowRegistryClientEntity**](FlowRegistryClientEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHistory**
> HistoryEntity DeleteHistory(ctx, endDate)
Purges history



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endDate** | **string**| Purge actions before this date/time. | 

### Return type

[**HistoryEntity**](HistoryEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
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

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

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

No authorization required

### HTTP request headers

 - **Content-Type**: */*
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

No authorization required

### HTTP request headers

 - **Content-Type**: */*
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

No authorization required

### HTTP request headers

 - **Content-Type**: */*
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

No authorization required

### HTTP request headers

 - **Content-Type**: */*
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

No authorization required

### HTTP request headers

 - **Content-Type**: */*
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

No authorization required

### HTTP request headers

 - **Content-Type**: */*
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


 **sensitive** | **optional.Bool**| Property Descriptor requested sensitive status | [default to false]

### Return type

[**PropertyDescriptorEntity**](PropertyDescriptorEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
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

No authorization required

### HTTP request headers

 - **Content-Type**: */*
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

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFlowRegistryClient**
> FlowRegistryClientEntity UpdateFlowRegistryClient(ctx, id, body)
Updates a flow registry client



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The flow registry client id. | 
  **body** | [**FlowRegistryClientEntity**](FlowRegistryClientEntity.md)| The flow registry client configuration details. | 

### Return type

[**FlowRegistryClientEntity**](FlowRegistryClientEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNode**
> NodeEntity UpdateNode(ctx, id, body)
Updates a node in the cluster



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The node id. | 
  **body** | [**NodeEntity**](NodeEntity.md)| The node configuration. The only configuration that will be honored at this endpoint is the status. | 

### Return type

[**NodeEntity**](NodeEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

