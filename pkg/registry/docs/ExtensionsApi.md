# \ExtensionsApi

All URIs are relative to *http://localhost/nifi-registry-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExtensions**](ExtensionsApi.md#GetExtensions) | **Get** /extensions | Get all extensions
[**GetExtensionsProvidingServiceAPI**](ExtensionsApi.md#GetExtensionsProvidingServiceAPI) | **Get** /extensions/provided-service-api | Get extensions providing service API
[**GetTags**](ExtensionsApi.md#GetTags) | **Get** /extensions/tags | Get extension tags


# **GetExtensions**
> ExtensionMetadataContainer GetExtensions(ctx, optional)
Get all extensions

Gets the metadata for all extensions that match the filter params and are part of bundles located in buckets the current user is authorized for. If the user is not authorized to any buckets, an empty result set will be returned.  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExtensionsApiGetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExtensionsApiGetExtensionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleType** | **optional.String**| The type of bundles to return | 
 **extensionType** | **optional.String**| The type of extensions to return | 
 **tag** | [**optional.Interface of []string**](string.md)| The tags to filter on, will be used in an OR statement | 

### Return type

[**ExtensionMetadataContainer**](ExtensionMetadataContainer.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtensionsProvidingServiceAPI**
> ExtensionMetadataContainer GetExtensionsProvidingServiceAPI(ctx, className, groupId, artifactId, version)
Get extensions providing service API

Gets the metadata for extensions that provide the specified API and are part of bundles located in buckets the current user is authorized for. If the user is not authorized to any buckets, an empty result set will be returned.  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **className** | **string**| The name of the service API class | 
  **groupId** | **string**| The groupId of the bundle containing the service API class | 
  **artifactId** | **string**| The artifactId of the bundle containing the service API class | 
  **version** | **string**| The version of the bundle containing the service API class | 

### Return type

[**ExtensionMetadataContainer**](ExtensionMetadataContainer.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTags**
> []TagCount GetTags(ctx, )
Get extension tags

Gets all the extension tags known to this NiFi Registry instance, along with the number of extensions that have the given tag.  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]TagCount**](TagCount.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

