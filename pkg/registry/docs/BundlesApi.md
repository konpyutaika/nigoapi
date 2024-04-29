# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBundle**](BundlesApi.md#DeleteBundle) | **Delete** /bundles/{bundleId} | Delete bundle
[**DeleteBundleVersion**](BundlesApi.md#DeleteBundleVersion) | **Delete** /bundles/{bundleId}/versions/{version} | Delete bundle version
[**GetBundle**](BundlesApi.md#GetBundle) | **Get** /bundles/{bundleId} | Get bundle
[**GetBundleVersion**](BundlesApi.md#GetBundleVersion) | **Get** /bundles/{bundleId}/versions/{version} | Get bundle version
[**GetBundleVersionContent**](BundlesApi.md#GetBundleVersionContent) | **Get** /bundles/{bundleId}/versions/{version}/content | Get bundle version content
[**GetBundleVersionExtension**](BundlesApi.md#GetBundleVersionExtension) | **Get** /bundles/{bundleId}/versions/{version}/extensions/{name} | Get bundle version extension
[**GetBundleVersionExtensionAdditionalDetailsDocs**](BundlesApi.md#GetBundleVersionExtensionAdditionalDetailsDocs) | **Get** /bundles/{bundleId}/versions/{version}/extensions/{name}/docs/additional-details | Get bundle version extension docs details
[**GetBundleVersionExtensionDocs**](BundlesApi.md#GetBundleVersionExtensionDocs) | **Get** /bundles/{bundleId}/versions/{version}/extensions/{name}/docs | Get bundle version extension docs
[**GetBundleVersionExtensions**](BundlesApi.md#GetBundleVersionExtensions) | **Get** /bundles/{bundleId}/versions/{version}/extensions | Get bundle version extensions
[**GetBundleVersions**](BundlesApi.md#GetBundleVersions) | **Get** /bundles/{bundleId}/versions | Get bundle versions
[**GetBundleVersions1**](BundlesApi.md#GetBundleVersions1) | **Get** /bundles/versions | Get all bundle versions
[**GetBundles**](BundlesApi.md#GetBundles) | **Get** /bundles | Get all bundles

# **DeleteBundle**
> Bundle DeleteBundle(ctx, bundleId)
Delete bundle

Deletes the given extension bundle and all of it's versions.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bundleId** | **string**| The extension bundle identifier | 

### Return type

[**Bundle**](Bundle.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBundleVersion**
> BundleVersion DeleteBundleVersion(ctx, bundleId, version)
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

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundle**
> Bundle GetBundle(ctx, bundleId)
Get bundle

Gets the metadata about an extension bundle.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bundleId** | **string**| The extension bundle identifier | 

### Return type

[**Bundle**](Bundle.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundleVersion**
> BundleVersion GetBundleVersion(ctx, bundleId, version)
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

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundleVersionContent**
> string GetBundleVersionContent(ctx, bundleId, version)
Get bundle version content

Gets the binary content for the given version of the given extension bundle.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bundleId** | **string**| The extension bundle identifier | 
  **version** | **string**| The version of the bundle | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundleVersionExtension**
> []Extension GetBundleVersionExtension(ctx, bundleId, version, name)
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

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundleVersionExtensionAdditionalDetailsDocs**
> GetBundleVersionExtensionAdditionalDetailsDocs(ctx, bundleId, version, name)
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

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundleVersionExtensionDocs**
> GetBundleVersionExtensionDocs(ctx, bundleId, version, name)
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

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundleVersionExtensions**
> []ExtensionMetadata GetBundleVersionExtensions(ctx, bundleId, version)
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

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundleVersions**
> []BundleVersionMetadata GetBundleVersions(ctx, bundleId)
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

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundleVersions1**
> []BundleVersionMetadata GetBundleVersions1(ctx, optional)
Get all bundle versions

Gets the metadata about extension bundle versions across all authorized buckets with optional filters applied. If the user is not authorized to any buckets, an empty list will be returned.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BundlesApiGetBundleVersions1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BundlesApiGetBundleVersions1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **optional.string**| Optional groupId to filter results. The value may be an exact match, or a wildcard, such as &#x27;com.%&#x27; to select all bundle versions where the groupId starts with &#x27;com.&#x27;. | 
 **artifactId** | **optional.string**| Optional artifactId to filter results. The value may be an exact match, or a wildcard, such as &#x27;nifi-%&#x27; to select all bundle versions where the artifactId starts with &#x27;nifi-&#x27;. | 
 **version** | **optional.string**| Optional version to filter results. The value maye be an exact match, or a wildcard, such as &#x27;1.0.%&#x27; to select all bundle versions where the version starts with &#x27;1.0.&#x27;. | 

### Return type

[**[]BundleVersionMetadata**](BundleVersionMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundles**
> []Bundle GetBundles(ctx, optional)
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
 **bucketName** | **optional.string**| Optional bucket name to filter results. The value may be an exact match, or a wildcard, such as &#x27;My Bucket%&#x27; to select all bundles where the bucket name starts with &#x27;My Bucket&#x27;. | 
 **groupId** | **optional.string**| Optional groupId to filter results. The value may be an exact match, or a wildcard, such as &#x27;com.%&#x27; to select all bundles where the groupId starts with &#x27;com.&#x27;. | 
 **artifactId** | **optional.string**| Optional artifactId to filter results. The value may be an exact match, or a wildcard, such as &#x27;nifi-%&#x27; to select all bundles where the artifactId starts with &#x27;nifi-&#x27;. | 

### Return type

[**[]Bundle**](Bundle.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

