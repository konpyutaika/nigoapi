# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCounters**](CountersApi.md#GetCounters) | **Get** /counters | Gets the current counters for this NiFi
[**UpdateAllCounters**](CountersApi.md#UpdateAllCounters) | **Put** /counters | Updates all counters. This will reset all counter values to 0
[**UpdateCounter**](CountersApi.md#UpdateCounter) | **Put** /counters/{id} | Updates the specified counter. This will reset the counter value to 0

# **GetCounters**
> CountersEntity GetCounters(ctx, optional)
Gets the current counters for this NiFi

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CountersApiGetCountersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CountersApiGetCountersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodewise** | **optional.bool**| Whether or not to include the breakdown per node. Optional, defaults to false | [default to false]
 **clusterNodeId** | **optional.string**| The id of the node where to get the status. | 

### Return type

[**CountersEntity**](CountersEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAllCounters**
> CountersEntity UpdateAllCounters(ctx, )
Updates all counters. This will reset all counter values to 0

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CountersEntity**](CountersEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCounter**
> CounterEntity UpdateCounter(ctx, id)
Updates the specified counter. This will reset the counter value to 0

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The id of the counter. | 

### Return type

[**CounterEntity**](CounterEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

