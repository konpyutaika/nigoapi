# \BundlesApi

All URIs are relative to *http://localhost/nifi-registry-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBundleVersionExtensionAdditionalDetailsDocs**](BundlesApi.md#GetBundleVersionExtensionAdditionalDetailsDocs) | **Get** /bundles/{bundleId}/versions/{version}/extensions/{name}/docs/additional-details | Get bundle version extension docs details
[**GetBundleVersionExtensionDocs**](BundlesApi.md#GetBundleVersionExtensionDocs) | **Get** /bundles/{bundleId}/versions/{version}/extensions/{name}/docs | Get bundle version extension docs
[**GetBundleVersions**](BundlesApi.md#GetBundleVersions) | **Get** /bundles/versions | Get all bundle versions
[**GetBundles**](BundlesApi.md#GetBundles) | **Get** /bundles | Get all bundles
[**GlobalDeleteBundleVersion**](BundlesApi.md#GlobalDeleteBundleVersion) | **Delete** /bundles/{bundleId}/versions/{version} | Delete bundle version
[**GlobalDeleteExtensionBundle**](BundlesApi.md#GlobalDeleteExtensionBundle) | **Delete** /bundles/{bundleId} | Delete bundle
[**GlobalGetBundleVersion**](BundlesApi.md#GlobalGetBundleVersion) | **Get** /bundles/{bundleId}/versions/{version} | Get bundle version
[**GlobalGetBundleVersionContent**](BundlesApi.md#GlobalGetBundleVersionContent) | **Get** /bundles/{bundleId}/versions/{version}/content | Get bundle version content
[**GlobalGetBundleVersionExtension**](BundlesApi.md#GlobalGetBundleVersionExtension) | **Get** /bundles/{bundleId}/versions/{version}/extensions/{name} | Get bundle version extension
[**GlobalGetBundleVersionExtensions**](BundlesApi.md#GlobalGetBundleVersionExtensions) | **Get** /bundles/{bundleId}/versions/{version}/extensions | Get bundle version extensions
[**GlobalGetBundleVersions**](BundlesApi.md#GlobalGetBundleVersions) | **Get** /bundles/{bundleId}/versions | Get bundle versions
[**GlobalGetExtensionBundle**](BundlesApi.md#GlobalGetExtensionBundle) | **Get** /bundles/{bundleId} | Get bundle


# **GetBundleVersionExtensionAdditionalDetailsDocs**
> string GetBundleVersionExtensionAdditionalDetailsDocs(ctx, bundleId, version, name)
Get bundle version extension docs details

Gets the additional details documentation for the given extension in the given extension bundle version.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bundleId** | **string**| The extension bundle identifier | 
  **version** | **string**| The version of the bundle | 
  **name** | **string**| The fully qualified name of the extension | 

### Return type

**string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundleVersionExtensionDocs**
> string GetBundleVersionExtensionDocs(ctx, bundleId, version, name)
Get bundle version extension docs

Gets the documentation for the given extension in the given extension bundle version.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bundleId** | **string**| The extension bundle identifier | 
  **version** | **string**| The version of the bundle | 
  **name** | **string**| The fully qualified name of the extension | 

### Return type

**string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundleVersions**
> []BundleVersionMetadata GetBundleVersions(ctx, optional)
Get all bundle versions

Gets the metadata about extension bundle versions across all authorized buckets with optional filters applied. If the user is not authorized to any buckets, an empty list will be returned.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BundlesApiGetBundleVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BundlesApiGetBundleVersionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **optional.String**| Optional groupId to filter results. The value may be an exact match, or a wildcard, such as &#39;com.%&#39; to select all bundle versions where the groupId starts with &#39;com.&#39;. | 
 **artifactId** | **optional.String**| Optional artifactId to filter results. The value may be an exact match, or a wildcard, such as &#39;nifi-%&#39; to select all bundle versions where the artifactId starts with &#39;nifi-&#39;. | 
 **version** | **optional.String**| Optional version to filter results. The value maye be an exact match, or a wildcard, such as &#39;1.0.%&#39; to select all bundle versions where the version starts with &#39;1.0.&#39;. | 

### Return type

[**[]BundleVersionMetadata**](BundleVersionMetadata.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundles**
> []ExtensionBundle GetBundles(ctx, optional)
Get all bundles

Gets the metadata for all bundles across all authorized buckets with optional filters applied. The returned results will include only items from buckets for which the user is authorized. If the user is not authorized to any buckets, an empty list will be returned.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BundlesApiGetBundlesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BundlesApiGetBundlesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketName** | **optional.String**| Optional bucket name to filter results. The value may be an exact match, or a wildcard, such as &#39;My Bucket%&#39; to select all bundles where the bucket name starts with &#39;My Bucket&#39;. | 
 **groupId** | **optional.String**| Optional groupId to filter results. The value may be an exact match, or a wildcard, such as &#39;com.%&#39; to select all bundles where the groupId starts with &#39;com.&#39;. | 
 **artifactId** | **optional.String**| Optional artifactId to filter results. The value may be an exact match, or a wildcard, such as &#39;nifi-%&#39; to select all bundles where the artifactId starts with &#39;nifi-&#39;. | 

### Return type

[**[]ExtensionBundle**](ExtensionBundle.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalDeleteBundleVersion**
> BundleVersion GlobalDeleteBundleVersion(ctx, bundleId, version)
Delete bundle version

Deletes the given extension bundle version and it's associated binary content.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bundleId** | **string**| The extension bundle identifier | 
  **version** | **string**| The version of the bundle | 

### Return type

[**BundleVersion**](BundleVersion.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalDeleteExtensionBundle**
> ExtensionBundle GlobalDeleteExtensionBundle(ctx, bundleId)
Delete bundle

Deletes the given extension bundle and all of it's versions.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bundleId** | **string**| The extension bundle identifier | 

### Return type

[**ExtensionBundle**](ExtensionBundle.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalGetBundleVersion**
> BundleVersion GlobalGetBundleVersion(ctx, bundleId, version)
Get bundle version

Gets the descriptor for the given version of the given extension bundle.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bundleId** | **string**| The extension bundle identifier | 
  **version** | **string**| The version of the bundle | 

### Return type

[**BundleVersion**](BundleVersion.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalGetBundleVersionContent**
> []string GlobalGetBundleVersionContent(ctx, bundleId, version)
Get bundle version content

Gets the binary content for the given version of the given extension bundle.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bundleId** | **string**| The extension bundle identifier | 
  **version** | **string**| The version of the bundle | 

### Return type

**[]string**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalGetBundleVersionExtension**
> []Extension GlobalGetBundleVersionExtension(ctx, bundleId, version, name)
Get bundle version extension

Gets the metadata about the extension with the given name in the given extension bundle version.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bundleId** | **string**| The extension bundle identifier | 
  **version** | **string**| The version of the bundle | 
  **name** | **string**| The fully qualified name of the extension | 

### Return type

[**[]Extension**](Extension.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalGetBundleVersionExtensions**
> []ExtensionMetadata GlobalGetBundleVersionExtensions(ctx, bundleId, version)
Get bundle version extensions

Gets the metadata about the extensions in the given extension bundle version.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bundleId** | **string**| The extension bundle identifier | 
  **version** | **string**| The version of the bundle | 

### Return type

[**[]ExtensionMetadata**](ExtensionMetadata.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalGetBundleVersions**
> []BundleVersionMetadata GlobalGetBundleVersions(ctx, bundleId)
Get bundle versions

Gets the metadata for the versions of the given extension bundle.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bundleId** | **string**| The extension bundle identifier | 

### Return type

[**[]BundleVersionMetadata**](BundleVersionMetadata.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalGetExtensionBundle**
> ExtensionBundle GlobalGetExtensionBundle(ctx, bundleId)
Get bundle

Gets the metadata about an extension bundle.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bundleId** | **string**| The extension bundle identifier | 

### Return type

[**ExtensionBundle**](ExtensionBundle.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

