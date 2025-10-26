# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJmxMetrics**](SystemDiagnosticsApi.md#GetJmxMetrics) | **Get** /system-diagnostics/jmx-metrics | Retrieve available JMX metrics
[**GetSystemDiagnostics**](SystemDiagnosticsApi.md#GetSystemDiagnostics) | **Get** /system-diagnostics | Gets the diagnostics for the system NiFi is running on

# **GetJmxMetrics**
> JmxMetricsResultsEntity GetJmxMetrics(ctx, optional)
Retrieve available JMX metrics

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemDiagnosticsApiGetJmxMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemDiagnosticsApiGetJmxMetricsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **beanNameFilter** | **optional.string**| Regular Expression Pattern to be applied against the ObjectName | 

### Return type

[**JmxMetricsResultsEntity**](JmxMetricsResultsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSystemDiagnostics**
> SystemDiagnosticsEntity GetSystemDiagnostics(ctx, optional)
Gets the diagnostics for the system NiFi is running on

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemDiagnosticsApiGetSystemDiagnosticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemDiagnosticsApiGetSystemDiagnosticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodewise** | **optional.bool**| Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **diagnosticLevel** | **optional.string**| Whether or not to include verbose details. Optional, defaults to false | [default to BASIC]
 **clusterNodeId** | **optional.string**| The id of the node where to get the status. | 

### Return type

[**SystemDiagnosticsEntity**](SystemDiagnosticsEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

