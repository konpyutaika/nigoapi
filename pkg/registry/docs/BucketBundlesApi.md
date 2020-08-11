# \BucketBundlesApi

All URIs are relative to *http://localhost/nifi-registry-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExtensionBundleVersion**](BucketBundlesApi.md#CreateExtensionBundleVersion) | **Post** /buckets/{bucketId}/bundles/{bundleType} | Create extension bundle version
[**GetExtensionBundles**](BucketBundlesApi.md#GetExtensionBundles) | **Get** /buckets/{bucketId}/bundles | Get extension bundles by bucket


# **CreateExtensionBundleVersion**
> BundleVersion CreateExtensionBundleVersion(ctx, bucketId, bundleType)
Create extension bundle version

Creates a version of an extension bundle by uploading a binary artifact. If an extension bundle already exists in the given bucket with the same group id and artifact id as that of the bundle being uploaded, then it will be added as a new version to the existing bundle. If an extension bundle does not already exist in the given bucket with the same group id and artifact id, then a new extension bundle will be created and this version will be added to the new bundle. Client's may optionally supply a SHA-256 in hex format through the multi-part form field 'sha256'. If supplied, then this value will be compared against the SHA-256 computed by the server, and the bundle will be rejected if the values do not match. If not supplied, the bundle will be accepted, but will be marked to indicate that the client did not supply a SHA-256 during creation.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketId** | **string**| The bucket identifier | 
  **bundleType** | **string**| The type of the bundle | 

### Return type

[**BundleVersion**](BundleVersion.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtensionBundles**
> []ExtensionBundle GetExtensionBundles(ctx, bucketId)
Get extension bundles by bucket

  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketId** | **string**| The bucket identifier | 

### Return type

[**[]ExtensionBundle**](ExtensionBundle.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

