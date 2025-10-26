# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateControllerServices**](FlowApi.md#ActivateControllerServices) | **Put** /flow/process-groups/{id}/controller-services | Enable or disable Controller Services in the specified Process Group.
[**DownloadReportingTaskSnapshot**](FlowApi.md#DownloadReportingTaskSnapshot) | **Get** /flow/reporting-tasks/download | Download a snapshot of the given reporting tasks and any controller services they use
[**GenerateClientId**](FlowApi.md#GenerateClientId) | **Get** /flow/client-id | Generates a client id.
[**GetAboutInfo**](FlowApi.md#GetAboutInfo) | **Get** /flow/about | Retrieves details about this NiFi to put in the About dialog
[**GetAction**](FlowApi.md#GetAction) | **Get** /flow/history/{id} | Gets an action
[**GetAdditionalDetails**](FlowApi.md#GetAdditionalDetails) | **Get** /flow/additional-details/{group}/{artifact}/{version}/{type} | Retrieves the additional details for the specified component type.
[**GetAllFlowAnalysisResults**](FlowApi.md#GetAllFlowAnalysisResults) | **Get** /flow/flow-analysis/results | Returns all flow analysis results currently in effect
[**GetBanners**](FlowApi.md#GetBanners) | **Get** /flow/banners | Retrieves the banners for this NiFi
[**GetBranches**](FlowApi.md#GetBranches) | **Get** /flow/registries/{id}/branches | Gets the branches from the specified registry for the current user
[**GetBreadcrumbs**](FlowApi.md#GetBreadcrumbs) | **Get** /flow/process-groups/{id}/breadcrumbs | Gets the breadcrumbs for a process group
[**GetBuckets**](FlowApi.md#GetBuckets) | **Get** /flow/registries/{id}/buckets | Gets the buckets from the specified registry for the current user
[**GetBulletinBoard**](FlowApi.md#GetBulletinBoard) | **Get** /flow/bulletin-board | Gets current bulletins
[**GetBulletins**](FlowApi.md#GetBulletins) | **Get** /flow/controller/bulletins | Retrieves Controller level bulletins
[**GetClusterSummary**](FlowApi.md#GetClusterSummary) | **Get** /flow/cluster/summary | The cluster summary for this NiFi
[**GetComponentHistory**](FlowApi.md#GetComponentHistory) | **Get** /flow/history/components/{componentId} | Gets configuration history for a component
[**GetConnectionStatistics**](FlowApi.md#GetConnectionStatistics) | **Get** /flow/connections/{id}/statistics | Gets statistics for a connection
[**GetConnectionStatus**](FlowApi.md#GetConnectionStatus) | **Get** /flow/connections/{id}/status | Gets status for a connection
[**GetConnectionStatusHistory**](FlowApi.md#GetConnectionStatusHistory) | **Get** /flow/connections/{id}/status/history | Gets the status history for a connection
[**GetContentViewers**](FlowApi.md#GetContentViewers) | **Get** /flow/content-viewers | Retrieves the registered content viewers
[**GetControllerServiceDefinition**](FlowApi.md#GetControllerServiceDefinition) | **Get** /flow/controller-service-definition/{group}/{artifact}/{version}/{type} | Retrieves the Controller Service Definition for the specified component type.
[**GetControllerServiceTypes**](FlowApi.md#GetControllerServiceTypes) | **Get** /flow/controller-service-types | Retrieves the types of controller services that this NiFi supports
[**GetControllerServicesFromController**](FlowApi.md#GetControllerServicesFromController) | **Get** /flow/controller/controller-services | Gets controller services for reporting tasks
[**GetControllerServicesFromGroup**](FlowApi.md#GetControllerServicesFromGroup) | **Get** /flow/process-groups/{id}/controller-services | Gets all controller services
[**GetControllerStatus**](FlowApi.md#GetControllerStatus) | **Get** /flow/status | Gets the current status of this NiFi
[**GetCurrentUser**](FlowApi.md#GetCurrentUser) | **Get** /flow/current-user | Retrieves the user identity of the user making the request
[**GetDetails**](FlowApi.md#GetDetails) | **Get** /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/details | Gets the details of a flow from the specified registry and bucket for the specified flow for the current user
[**GetFlow**](FlowApi.md#GetFlow) | **Get** /flow/process-groups/{id} | Gets a process group
[**GetFlowAnalysisResults**](FlowApi.md#GetFlowAnalysisResults) | **Get** /flow/flow-analysis/results/{processGroupId} | Returns flow analysis results produced by the analysis of a given process group
[**GetFlowAnalysisRuleDefinition**](FlowApi.md#GetFlowAnalysisRuleDefinition) | **Get** /flow/flow-analysis-rule-definition/{group}/{artifact}/{version}/{type} | Retrieves the Flow Analysis Rule Definition for the specified component type.
[**GetFlowAnalysisRuleTypes**](FlowApi.md#GetFlowAnalysisRuleTypes) | **Get** /flow/flow-analysis-rule-types | Retrieves the types of available Flow Analysis Rules
[**GetFlowConfig**](FlowApi.md#GetFlowConfig) | **Get** /flow/config | Retrieves the configuration for this NiFi flow
[**GetFlowMetrics**](FlowApi.md#GetFlowMetrics) | **Get** /flow/metrics/{producer} | Gets all metrics for the flow from a particular node
[**GetFlows**](FlowApi.md#GetFlows) | **Get** /flow/registries/{registry-id}/buckets/{bucket-id}/flows | Gets the flows from the specified registry and bucket for the current user
[**GetInputPortStatus**](FlowApi.md#GetInputPortStatus) | **Get** /flow/input-ports/{id}/status | Gets status for an input port
[**GetOutputPortStatus**](FlowApi.md#GetOutputPortStatus) | **Get** /flow/output-ports/{id}/status | Gets status for an output port
[**GetParameterContexts**](FlowApi.md#GetParameterContexts) | **Get** /flow/parameter-contexts | Gets all Parameter Contexts
[**GetParameterProviderDefinition**](FlowApi.md#GetParameterProviderDefinition) | **Get** /flow/parameter-provider-definition/{group}/{artifact}/{version}/{type} | Retrieves the Parameter Provider Definition for the specified component type.
[**GetParameterProviderTypes**](FlowApi.md#GetParameterProviderTypes) | **Get** /flow/parameter-provider-types | Retrieves the types of parameter providers that this NiFi supports
[**GetParameterProviders**](FlowApi.md#GetParameterProviders) | **Get** /flow/parameter-providers | Gets all parameter providers
[**GetPrioritizers**](FlowApi.md#GetPrioritizers) | **Get** /flow/prioritizers | Retrieves the types of prioritizers that this NiFi supports
[**GetProcessGroupStatus**](FlowApi.md#GetProcessGroupStatus) | **Get** /flow/process-groups/{id}/status | Gets the status for a process group
[**GetProcessGroupStatusHistory**](FlowApi.md#GetProcessGroupStatusHistory) | **Get** /flow/process-groups/{id}/status/history | Gets status history for a remote process group
[**GetProcessorDefinition**](FlowApi.md#GetProcessorDefinition) | **Get** /flow/processor-definition/{group}/{artifact}/{version}/{type} | Retrieves the Processor Definition for the specified component type.
[**GetProcessorStatus**](FlowApi.md#GetProcessorStatus) | **Get** /flow/processors/{id}/status | Gets status for a processor
[**GetProcessorStatusHistory**](FlowApi.md#GetProcessorStatusHistory) | **Get** /flow/processors/{id}/status/history | Gets status history for a processor
[**GetProcessorTypes**](FlowApi.md#GetProcessorTypes) | **Get** /flow/processor-types | Retrieves the types of processors that this NiFi supports
[**GetRegistryClients**](FlowApi.md#GetRegistryClients) | **Get** /flow/registries | Gets the listing of available flow registry clients
[**GetRemoteProcessGroupStatus**](FlowApi.md#GetRemoteProcessGroupStatus) | **Get** /flow/remote-process-groups/{id}/status | Gets status for a remote process group
[**GetRemoteProcessGroupStatusHistory**](FlowApi.md#GetRemoteProcessGroupStatusHistory) | **Get** /flow/remote-process-groups/{id}/status/history | Gets the status history
[**GetReportingTaskDefinition**](FlowApi.md#GetReportingTaskDefinition) | **Get** /flow/reporting-task-definition/{group}/{artifact}/{version}/{type} | Retrieves the Reporting Task Definition for the specified component type.
[**GetReportingTaskSnapshot**](FlowApi.md#GetReportingTaskSnapshot) | **Get** /flow/reporting-tasks/snapshot | Get a snapshot of the given reporting tasks and any controller services they use
[**GetReportingTaskTypes**](FlowApi.md#GetReportingTaskTypes) | **Get** /flow/reporting-task-types | Retrieves the types of reporting tasks that this NiFi supports
[**GetReportingTasks**](FlowApi.md#GetReportingTasks) | **Get** /flow/reporting-tasks | Gets all reporting tasks
[**GetRuntimeManifest**](FlowApi.md#GetRuntimeManifest) | **Get** /flow/runtime-manifest | Retrieves the runtime manifest for this NiFi instance.
[**GetVersionDifferences**](FlowApi.md#GetVersionDifferences) | **Get** /flow/registries/{registry-id}/branches/{branch-id-a}/buckets/{bucket-id-a}/flows/{flow-id-a}/{version-a}/diff/branches/{branch-id-b}/buckets/{bucket-id-b}/flows/{flow-id-b}/{version-b} | Gets the differences between two versions of the same versioned flow, the basis of the comparison will be the first version
[**GetVersions**](FlowApi.md#GetVersions) | **Get** /flow/registries/{registry-id}/buckets/{bucket-id}/flows/{flow-id}/versions | Gets the flow versions from the specified registry and bucket for the specified flow for the current user
[**QueryHistory**](FlowApi.md#QueryHistory) | **Get** /flow/history | Gets configuration history
[**ScheduleComponents**](FlowApi.md#ScheduleComponents) | **Put** /flow/process-groups/{id} | Schedule or unschedule components in the specified Process Group.
[**SearchCluster**](FlowApi.md#SearchCluster) | **Get** /flow/cluster/search-results | Searches the cluster for a node with the specified address
[**SearchFlow**](FlowApi.md#SearchFlow) | **Get** /flow/search-results | Performs a search against this NiFi using the specified search term

# **ActivateControllerServices**
> ActivateControllerServicesEntity ActivateControllerServices(ctx, body, id)
Enable or disable Controller Services in the specified Process Group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ActivateControllerServicesEntity**](ActivateControllerServicesEntity.md)| The request to schedule or unschedule. If the comopnents in the request are not specified, all authorized components will be considered. | 
  **id** | **string**| The process group id. | 

### Return type

[**ActivateControllerServicesEntity**](ActivateControllerServicesEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadReportingTaskSnapshot**
> string DownloadReportingTaskSnapshot(ctx, optional)
Download a snapshot of the given reporting tasks and any controller services they use

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FlowApiDownloadReportingTaskSnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiDownloadReportingTaskSnapshotOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportingTaskId** | **optional.string**| Specifies a reporting task id to export. If not specified, all reporting tasks will be exported. | 

### Return type

**string**

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateClientId**
> string GenerateClientId(ctx, )
Generates a client id.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAboutInfo**
> AboutEntity GetAboutInfo(ctx, )
Retrieves details about this NiFi to put in the About dialog

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AboutEntity**](AboutEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAction**
> ActionEntity GetAction(ctx, id)
Gets an action

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**IntegerParameter**](.md)| The action id. | 

### Return type

[**ActionEntity**](ActionEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdditionalDetails**
> AdditionalDetailsEntity GetAdditionalDetails(ctx, group, artifact, version, type_)
Retrieves the additional details for the specified component type.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| The bundle group | 
  **artifact** | **string**| The bundle artifact | 
  **version** | **string**| The bundle version | 
  **type_** | **string**| The processor type | 

### Return type

[**AdditionalDetailsEntity**](AdditionalDetailsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllFlowAnalysisResults**
> FlowAnalysisResultEntity GetAllFlowAnalysisResults(ctx, )
Returns all flow analysis results currently in effect

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FlowAnalysisResultEntity**](FlowAnalysisResultEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBanners**
> BannerEntity GetBanners(ctx, )
Retrieves the banners for this NiFi

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BannerEntity**](BannerEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBranches**
> FlowRegistryBranchesEntity GetBranches(ctx, id)
Gets the branches from the specified registry for the current user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The registry id. | 

### Return type

[**FlowRegistryBranchesEntity**](FlowRegistryBranchesEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBreadcrumbs**
> FlowBreadcrumbEntity GetBreadcrumbs(ctx, id)
Gets the breadcrumbs for a process group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 

### Return type

[**FlowBreadcrumbEntity**](FlowBreadcrumbEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuckets**
> FlowRegistryBucketsEntity GetBuckets(ctx, id, optional)
Gets the buckets from the specified registry for the current user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The registry id. | 
 **optional** | ***FlowApiGetBucketsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetBucketsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branch** | **optional.string**| The name of a branch to get the buckets from. If not specified the default branch of the registry client will be used. | 

### Return type

[**FlowRegistryBucketsEntity**](FlowRegistryBucketsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBulletinBoard**
> BulletinBoardEntity GetBulletinBoard(ctx, optional)
Gets current bulletins

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FlowApiGetBulletinBoardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetBulletinBoardOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | [**optional.Interface of LongParameter**](.md)| Includes bulletins with an id after this value. | 
 **sourceName** | [**optional.Interface of BulletinBoardPatternParameter**](.md)| Includes bulletins originating from this sources whose name match this regular expression. | 
 **message** | [**optional.Interface of BulletinBoardPatternParameter**](.md)| Includes bulletins whose message that match this regular expression. | 
 **sourceId** | [**optional.Interface of BulletinBoardPatternParameter**](.md)| Includes bulletins originating from this sources whose id match this regular expression. | 
 **groupId** | [**optional.Interface of BulletinBoardPatternParameter**](.md)| Includes bulletins originating from this sources whose group id match this regular expression. | 
 **limit** | [**optional.Interface of IntegerParameter**](.md)| The number of bulletins to limit the response to. | 

### Return type

[**BulletinBoardEntity**](BulletinBoardEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBulletins**
> ControllerBulletinsEntity GetBulletins(ctx, )
Retrieves Controller level bulletins

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ControllerBulletinsEntity**](ControllerBulletinsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterSummary**
> ClusterSummaryEntity GetClusterSummary(ctx, )
The cluster summary for this NiFi

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterSummaryEntity**](ClusterSummaryEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComponentHistory**
> ComponentHistoryEntity GetComponentHistory(ctx, componentId)
Gets configuration history for a component

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **componentId** | **string**| The component id. | 

### Return type

[**ComponentHistoryEntity**](ComponentHistoryEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectionStatistics**
> ConnectionStatisticsEntity GetConnectionStatistics(ctx, id, optional)
Gets statistics for a connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connection id. | 
 **optional** | ***FlowApiGetConnectionStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetConnectionStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodewise** | **optional.bool**| Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **optional.string**| The id of the node where to get the statistics. | 

### Return type

[**ConnectionStatisticsEntity**](ConnectionStatisticsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectionStatus**
> ConnectionStatusEntity GetConnectionStatus(ctx, id, optional)
Gets status for a connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connection id. | 
 **optional** | ***FlowApiGetConnectionStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetConnectionStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodewise** | **optional.bool**| Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **optional.string**| The id of the node where to get the status. | 

### Return type

[**ConnectionStatusEntity**](ConnectionStatusEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectionStatusHistory**
> StatusHistoryEntity GetConnectionStatusHistory(ctx, id)
Gets the status history for a connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The connection id. | 

### Return type

[**StatusHistoryEntity**](StatusHistoryEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContentViewers**
> ContentViewerEntity GetContentViewers(ctx, )
Retrieves the registered content viewers

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ContentViewerEntity**](ContentViewerEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetControllerServiceDefinition**
> ControllerServiceDefinition GetControllerServiceDefinition(ctx, group, artifact, version, type_)
Retrieves the Controller Service Definition for the specified component type.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| The bundle group | 
  **artifact** | **string**| The bundle artifact | 
  **version** | **string**| The bundle version | 
  **type_** | **string**| The controller service type | 

### Return type

[**ControllerServiceDefinition**](ControllerServiceDefinition.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetControllerServiceTypes**
> ControllerServiceTypesEntity GetControllerServiceTypes(ctx, optional)
Retrieves the types of controller services that this NiFi supports

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FlowApiGetControllerServiceTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetControllerServiceTypesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceType** | **optional.string**| If specified, will only return controller services that are compatible with this type of service. | 
 **serviceBundleGroup** | **optional.string**| If serviceType specified, is the bundle group of the serviceType. | 
 **serviceBundleArtifact** | **optional.string**| If serviceType specified, is the bundle artifact of the serviceType. | 
 **serviceBundleVersion** | **optional.string**| If serviceType specified, is the bundle version of the serviceType. | 
 **bundleGroupFilter** | **optional.string**| If specified, will only return types that are a member of this bundle group. | 
 **bundleArtifactFilter** | **optional.string**| If specified, will only return types that are a member of this bundle artifact. | 
 **typeFilter** | **optional.string**| If specified, will only return types whose fully qualified classname matches. | 

### Return type

[**ControllerServiceTypesEntity**](ControllerServiceTypesEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetControllerServicesFromController**
> ControllerServicesEntity GetControllerServicesFromController(ctx, optional)
Gets controller services for reporting tasks

If the uiOnly query parameter is provided with a value of true, the returned entity may only contain fields that are necessary for rendering the NiFi User Interface. As such, the selected fields may change at any time, even during incremental releases, without warning. As a result, this parameter should not be provided by any client other than the UI.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FlowApiGetControllerServicesFromControllerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetControllerServicesFromControllerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uiOnly** | **optional.bool**|  | [default to false]
 **includeReferencingComponents** | **optional.bool**| Whether or not to include services&#x27; referencing components in the response | [default to true]

### Return type

[**ControllerServicesEntity**](ControllerServicesEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetControllerServicesFromGroup**
> ControllerServicesEntity GetControllerServicesFromGroup(ctx, id, optional)
Gets all controller services

If the uiOnly query parameter is provided with a value of true, the returned entity may only contain fields that are necessary for rendering the NiFi User Interface. As such, the selected fields may change at any time, even during incremental releases, without warning. As a result, this parameter should not be provided by any client other than the UI.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 
 **optional** | ***FlowApiGetControllerServicesFromGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetControllerServicesFromGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAncestorGroups** | **optional.bool**| Whether or not to include parent/ancestor process groups | [default to true]
 **includeDescendantGroups** | **optional.bool**| Whether or not to include descendant process groups | [default to false]
 **includeReferencingComponents** | **optional.bool**| Whether or not to include services&#x27; referencing components in the response | [default to true]
 **uiOnly** | **optional.bool**|  | [default to false]

### Return type

[**ControllerServicesEntity**](ControllerServicesEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetControllerStatus**
> ControllerStatusEntity GetControllerStatus(ctx, )
Gets the current status of this NiFi

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ControllerStatusEntity**](ControllerStatusEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentUser**
> CurrentUserEntity GetCurrentUser(ctx, )
Retrieves the user identity of the user making the request

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CurrentUserEntity**](CurrentUserEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetails**
> VersionedFlowEntity GetDetails(ctx, registryId, bucketId, flowId, optional)
Gets the details of a flow from the specified registry and bucket for the specified flow for the current user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryId** | **string**| The registry client id. | 
  **bucketId** | **string**| The bucket id. | 
  **flowId** | **string**| The flow id. | 
 **optional** | ***FlowApiGetDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **branch** | **optional.string**| The name of a branch to get the flow from. If not specified the default branch of the registry client will be used. | 

### Return type

[**VersionedFlowEntity**](VersionedFlowEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlow**
> ProcessGroupFlowEntity GetFlow(ctx, id, optional)
Gets a process group

If the uiOnly query parameter is provided with a value of true, the returned entity may only contain fields that are necessary for rendering the NiFi User Interface. As such, the selected fields may change at any time, even during incremental releases, without warning. As a result, this parameter should not be provided by any client other than the UI.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 
 **optional** | ***FlowApiGetFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetFlowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uiOnly** | **optional.bool**|  | [default to false]

### Return type

[**ProcessGroupFlowEntity**](ProcessGroupFlowEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlowAnalysisResults**
> FlowAnalysisResultEntity GetFlowAnalysisResults(ctx, processGroupId)
Returns flow analysis results produced by the analysis of a given process group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processGroupId** | **string**| The id of the process group representing (a part of) the flow to be analyzed. | 

### Return type

[**FlowAnalysisResultEntity**](FlowAnalysisResultEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlowAnalysisRuleDefinition**
> FlowAnalysisRuleDefinition GetFlowAnalysisRuleDefinition(ctx, group, artifact, version, type_)
Retrieves the Flow Analysis Rule Definition for the specified component type.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| The bundle group | 
  **artifact** | **string**| The bundle artifact | 
  **version** | **string**| The bundle version | 
  **type_** | **string**| The flow analysis rule type | 

### Return type

[**FlowAnalysisRuleDefinition**](FlowAnalysisRuleDefinition.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlowAnalysisRuleTypes**
> FlowAnalysisRuleTypesEntity GetFlowAnalysisRuleTypes(ctx, optional)
Retrieves the types of available Flow Analysis Rules

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FlowApiGetFlowAnalysisRuleTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetFlowAnalysisRuleTypesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleGroupFilter** | **optional.string**| If specified, will only return types that are a member of this bundle group. | 
 **bundleArtifactFilter** | **optional.string**| If specified, will only return types that are a member of this bundle artifact. | 
 **type_** | **optional.string**| If specified, will only return types whose fully qualified classname matches. | 

### Return type

[**FlowAnalysisRuleTypesEntity**](FlowAnalysisRuleTypesEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlowConfig**
> FlowConfigurationEntity GetFlowConfig(ctx, )
Retrieves the configuration for this NiFi flow

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FlowConfigurationEntity**](FlowConfigurationEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlowMetrics**
> StreamingOutput GetFlowMetrics(ctx, producer, optional)
Gets all metrics for the flow from a particular node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **producer** | **string**| The producer for flow file metrics. Each producer may have its own output format. | 
 **optional** | ***FlowApiGetFlowMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetFlowMetricsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includedRegistries** | **optional.string**| Set of included metrics registries. Duplicate the parameter to include multiple registries. All registries are included by default. | 
 **sampleName** | **optional.string**| Regular Expression Pattern to be applied against the sample name field | 
 **sampleLabelValue** | **optional.string**| Regular Expression Pattern to be applied against the sample label value field | 
 **rootFieldName** | **optional.string**| Name of the first field of JSON object. Applicable for JSON producer only. | 

### Return type

[**StreamingOutput**](StreamingOutput.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlows**
> VersionedFlowsEntity GetFlows(ctx, registryId, bucketId, optional)
Gets the flows from the specified registry and bucket for the current user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryId** | **string**| The registry client id. | 
  **bucketId** | **string**| The bucket id. | 
 **optional** | ***FlowApiGetFlowsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetFlowsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branch** | **optional.string**| The name of a branch to get the flows from. If not specified the default branch of the registry client will be used. | 

### Return type

[**VersionedFlowsEntity**](VersionedFlowsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInputPortStatus**
> PortStatusEntity GetInputPortStatus(ctx, id, optional)
Gets status for an input port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The input port id. | 
 **optional** | ***FlowApiGetInputPortStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetInputPortStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodewise** | **optional.bool**| Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **optional.string**| The id of the node where to get the status. | 

### Return type

[**PortStatusEntity**](PortStatusEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOutputPortStatus**
> PortStatusEntity GetOutputPortStatus(ctx, id, optional)
Gets status for an output port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The output port id. | 
 **optional** | ***FlowApiGetOutputPortStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetOutputPortStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodewise** | **optional.bool**| Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **optional.string**| The id of the node where to get the status. | 

### Return type

[**PortStatusEntity**](PortStatusEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetParameterContexts**
> ParameterContextsEntity GetParameterContexts(ctx, )
Gets all Parameter Contexts

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ParameterContextsEntity**](ParameterContextsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetParameterProviderDefinition**
> ParameterProviderDefinition GetParameterProviderDefinition(ctx, group, artifact, version, type_)
Retrieves the Parameter Provider Definition for the specified component type.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| The bundle group | 
  **artifact** | **string**| The bundle artifact | 
  **version** | **string**| The bundle version | 
  **type_** | **string**| The parameter provider type | 

### Return type

[**ParameterProviderDefinition**](ParameterProviderDefinition.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetParameterProviderTypes**
> ParameterProviderTypesEntity GetParameterProviderTypes(ctx, optional)
Retrieves the types of parameter providers that this NiFi supports

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FlowApiGetParameterProviderTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetParameterProviderTypesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleGroupFilter** | **optional.string**| If specified, will only return types that are a member of this bundle group. | 
 **bundleArtifactFilter** | **optional.string**| If specified, will only return types that are a member of this bundle artifact. | 
 **type_** | **optional.string**| If specified, will only return types whose fully qualified classname matches. | 

### Return type

[**ParameterProviderTypesEntity**](ParameterProviderTypesEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetParameterProviders**
> ParameterProvidersEntity GetParameterProviders(ctx, )
Gets all parameter providers

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ParameterProvidersEntity**](ParameterProvidersEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPrioritizers**
> PrioritizerTypesEntity GetPrioritizers(ctx, )
Retrieves the types of prioritizers that this NiFi supports

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PrioritizerTypesEntity**](PrioritizerTypesEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcessGroupStatus**
> ProcessGroupStatusEntity GetProcessGroupStatus(ctx, id, optional)
Gets the status for a process group

The status for a process group includes status for all descendent components. When invoked on the root group with recursive set to true, it will return the current status of every component in the flow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 
 **optional** | ***FlowApiGetProcessGroupStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetProcessGroupStatusOpts struct
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

# **GetProcessGroupStatusHistory**
> StatusHistoryEntity GetProcessGroupStatusHistory(ctx, id)
Gets status history for a remote process group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The process group id. | 

### Return type

[**StatusHistoryEntity**](StatusHistoryEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcessorDefinition**
> ProcessorDefinition GetProcessorDefinition(ctx, group, artifact, version, type_)
Retrieves the Processor Definition for the specified component type.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| The bundle group | 
  **artifact** | **string**| The bundle artifact | 
  **version** | **string**| The bundle version | 
  **type_** | **string**| The processor type | 

### Return type

[**ProcessorDefinition**](ProcessorDefinition.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcessorStatus**
> ProcessorStatusEntity GetProcessorStatus(ctx, id, optional)
Gets status for a processor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The processor id. | 
 **optional** | ***FlowApiGetProcessorStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetProcessorStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodewise** | **optional.bool**| Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **optional.string**| The id of the node where to get the status. | 

### Return type

[**ProcessorStatusEntity**](ProcessorStatusEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcessorStatusHistory**
> StatusHistoryEntity GetProcessorStatusHistory(ctx, id)
Gets status history for a processor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The processor id. | 

### Return type

[**StatusHistoryEntity**](StatusHistoryEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProcessorTypes**
> ProcessorTypesEntity GetProcessorTypes(ctx, optional)
Retrieves the types of processors that this NiFi supports

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FlowApiGetProcessorTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetProcessorTypesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleGroupFilter** | **optional.string**| If specified, will only return types that are a member of this bundle group. | 
 **bundleArtifactFilter** | **optional.string**| If specified, will only return types that are a member of this bundle artifact. | 
 **type_** | **optional.string**| If specified, will only return types whose fully qualified classname matches. | 

### Return type

[**ProcessorTypesEntity**](ProcessorTypesEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRegistryClients**
> FlowRegistryClientsEntity GetRegistryClients(ctx, )
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

# **GetRemoteProcessGroupStatus**
> RemoteProcessGroupStatusEntity GetRemoteProcessGroupStatus(ctx, id, optional)
Gets status for a remote process group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The remote process group id. | 
 **optional** | ***FlowApiGetRemoteProcessGroupStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetRemoteProcessGroupStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodewise** | **optional.bool**| Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **optional.string**| The id of the node where to get the status. | 

### Return type

[**RemoteProcessGroupStatusEntity**](RemoteProcessGroupStatusEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRemoteProcessGroupStatusHistory**
> StatusHistoryEntity GetRemoteProcessGroupStatusHistory(ctx, id)
Gets the status history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The remote process group id. | 

### Return type

[**StatusHistoryEntity**](StatusHistoryEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportingTaskDefinition**
> ReportingTaskDefinition GetReportingTaskDefinition(ctx, group, artifact, version, type_)
Retrieves the Reporting Task Definition for the specified component type.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **group** | **string**| The bundle group | 
  **artifact** | **string**| The bundle artifact | 
  **version** | **string**| The bundle version | 
  **type_** | **string**| The reporting task type | 

### Return type

[**ReportingTaskDefinition**](ReportingTaskDefinition.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportingTaskSnapshot**
> VersionedReportingTaskSnapshot GetReportingTaskSnapshot(ctx, optional)
Get a snapshot of the given reporting tasks and any controller services they use

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FlowApiGetReportingTaskSnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetReportingTaskSnapshotOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportingTaskId** | **optional.string**| Specifies a reporting task id to export. If not specified, all reporting tasks will be exported. | 

### Return type

[**VersionedReportingTaskSnapshot**](VersionedReportingTaskSnapshot.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportingTaskTypes**
> ReportingTaskTypesEntity GetReportingTaskTypes(ctx, optional)
Retrieves the types of reporting tasks that this NiFi supports

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FlowApiGetReportingTaskTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetReportingTaskTypesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleGroupFilter** | **optional.string**| If specified, will only return types that are a member of this bundle group. | 
 **bundleArtifactFilter** | **optional.string**| If specified, will only return types that are a member of this bundle artifact. | 
 **type_** | **optional.string**| If specified, will only return types whose fully qualified classname matches. | 

### Return type

[**ReportingTaskTypesEntity**](ReportingTaskTypesEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportingTasks**
> ReportingTasksEntity GetReportingTasks(ctx, )
Gets all reporting tasks

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ReportingTasksEntity**](ReportingTasksEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRuntimeManifest**
> RuntimeManifestEntity GetRuntimeManifest(ctx, )
Retrieves the runtime manifest for this NiFi instance.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RuntimeManifestEntity**](RuntimeManifestEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVersionDifferences**
> FlowComparisonEntity GetVersionDifferences(ctx, registryId, branchIdA, bucketIdA, flowIdA, versionA, branchIdB, bucketIdB, flowIdB, versionB, optional)
Gets the differences between two versions of the same versioned flow, the basis of the comparison will be the first version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryId** | **string**| The registry client id. | 
  **branchIdA** | **string**| The branch id for the base version. | 
  **bucketIdA** | **string**| The bucket id for the base version. | 
  **flowIdA** | **string**| The flow id for the base version. | 
  **versionA** | **string**| The base version. | 
  **branchIdB** | **string**| The branch id for the compared version. | 
  **bucketIdB** | **string**| The bucket id for the compared version. | 
  **flowIdB** | **string**| The flow id for the compared version. | 
  **versionB** | **string**| The compared version. | 
 **optional** | ***FlowApiGetVersionDifferencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetVersionDifferencesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------









 **offset** | **optional.int32**| Must be a non-negative number. Specifies the starting point of the listing. 0 means start from the beginning. | [default to 0]
 **limit** | **optional.int32**| Limits the number of differences listed. This might lead to partial result. 0 means no limitation is applied. | [default to 1000]

### Return type

[**FlowComparisonEntity**](FlowComparisonEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVersions**
> VersionedFlowSnapshotMetadataSetEntity GetVersions(ctx, registryId, bucketId, flowId, optional)
Gets the flow versions from the specified registry and bucket for the specified flow for the current user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registryId** | **string**| The registry client id. | 
  **bucketId** | **string**| The bucket id. | 
  **flowId** | **string**| The flow id. | 
 **optional** | ***FlowApiGetVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiGetVersionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **branch** | **optional.string**| The name of a branch to get the flow versions from. If not specified the default branch of the registry client will be used. | 

### Return type

[**VersionedFlowSnapshotMetadataSetEntity**](VersionedFlowSnapshotMetadataSetEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryHistory**
> HistoryEntity QueryHistory(ctx, offset, count, optional)
Gets configuration history

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **offset** | [**IntegerParameter**](.md)| The offset into the result set. | 
  **count** | [**IntegerParameter**](.md)| The number of actions to return. | 
 **optional** | ***FlowApiQueryHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiQueryHistoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sortColumn** | **optional.string**| The field to sort on. | 
 **sortOrder** | **optional.string**| The direction to sort. | 
 **startDate** | [**optional.Interface of DateTimeParameter**](.md)| Include actions after this date. | 
 **endDate** | [**optional.Interface of DateTimeParameter**](.md)| Include actions before this date. | 
 **userIdentity** | **optional.string**| Include actions performed by this user. | 
 **sourceId** | **optional.string**| Include actions on this component. | 

### Return type

[**HistoryEntity**](HistoryEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScheduleComponents**
> ScheduleComponentsEntity ScheduleComponents(ctx, body, id)
Schedule or unschedule components in the specified Process Group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScheduleComponentsEntity**](ScheduleComponentsEntity.md)| The request to schedule or unschedule. If the components in the request are not specified, all authorized components will be considered. | 
  **id** | **string**| The process group id. | 

### Return type

[**ScheduleComponentsEntity**](ScheduleComponentsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchCluster**
> ClusterSearchResultsEntity SearchCluster(ctx, q)
Searches the cluster for a node with the specified address

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **q** | **string**| Node address to search for. | 

### Return type

[**ClusterSearchResultsEntity**](ClusterSearchResultsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchFlow**
> SearchResultsEntity SearchFlow(ctx, optional)
Performs a search against this NiFi using the specified search term

Only search results from authorized components will be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FlowApiSearchFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FlowApiSearchFlowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.string**|  | 
 **a** | **optional.string**|  | 

### Return type

[**SearchResultsEntity**](SearchResultsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

