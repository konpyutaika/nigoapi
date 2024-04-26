# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExtensionRepoArtifacts**](ExtensionRepositoryApi.md#GetExtensionRepoArtifacts) | **Get** /extension-repository/{bucketName}/{groupId} | Get extension repo artifacts
[**GetExtensionRepoBuckets**](ExtensionRepositoryApi.md#GetExtensionRepoBuckets) | **Get** /extension-repository | Get extension repo buckets
[**GetExtensionRepoGroups**](ExtensionRepositoryApi.md#GetExtensionRepoGroups) | **Get** /extension-repository/{bucketName} | Get extension repo groups
[**GetExtensionRepoVersion**](ExtensionRepositoryApi.md#GetExtensionRepoVersion) | **Get** /extension-repository/{bucketName}/{groupId}/{artifactId}/{version} | Get extension repo version
[**GetExtensionRepoVersionContent**](ExtensionRepositoryApi.md#GetExtensionRepoVersionContent) | **Get** /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/content | Get extension repo version content
[**GetExtensionRepoVersionExtension**](ExtensionRepositoryApi.md#GetExtensionRepoVersionExtension) | **Get** /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name} | Get extension repo extension
[**GetExtensionRepoVersionExtensionAdditionalDetailsDocs**](ExtensionRepositoryApi.md#GetExtensionRepoVersionExtensionAdditionalDetailsDocs) | **Get** /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs/additional-details | Get extension repo extension details
[**GetExtensionRepoVersionExtensionDocs**](ExtensionRepositoryApi.md#GetExtensionRepoVersionExtensionDocs) | **Get** /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions/{name}/docs | Get extension repo extension docs
[**GetExtensionRepoVersionExtensions**](ExtensionRepositoryApi.md#GetExtensionRepoVersionExtensions) | **Get** /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/extensions | Get extension repo extensions
[**GetExtensionRepoVersionSha256**](ExtensionRepositoryApi.md#GetExtensionRepoVersionSha256) | **Get** /extension-repository/{bucketName}/{groupId}/{artifactId}/{version}/sha256 | Get extension repo version checksum
[**GetExtensionRepoVersions**](ExtensionRepositoryApi.md#GetExtensionRepoVersions) | **Get** /extension-repository/{bucketName}/{groupId}/{artifactId} | Get extension repo versions
[**GetGlobalExtensionRepoVersionSha256**](ExtensionRepositoryApi.md#GetGlobalExtensionRepoVersionSha256) | **Get** /extension-repository/{groupId}/{artifactId}/{version}/sha256 | Get global extension repo version checksum

# **GetExtensionRepoArtifacts**
> []ExtensionRepoArtifact GetExtensionRepoArtifacts(ctx, bucketName, groupId)
Get extension repo artifacts

Gets the artifacts in the extension repository in the given bucket and group.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketName** | **string**| The bucket name | 
  **groupId** | **string**| The group id | 

### Return type

[**[]ExtensionRepoArtifact**](ExtensionRepoArtifact.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtensionRepoBuckets**
> []ExtensionRepoBucket GetExtensionRepoBuckets(ctx, )
Get extension repo buckets

Gets the names of the buckets the current user is authorized for in order to browse the repo by bucket.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ExtensionRepoBucket**](ExtensionRepoBucket.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtensionRepoGroups**
> []ExtensionRepoGroup GetExtensionRepoGroups(ctx, bucketName)
Get extension repo groups

Gets the groups in the extension repository in the given bucket.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketName** | **string**| The bucket name | 

### Return type

[**[]ExtensionRepoGroup**](ExtensionRepoGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtensionRepoVersion**
> ExtensionRepoVersion GetExtensionRepoVersion(ctx, bucketName, groupId, artifactId, version)
Get extension repo version

Gets information about the version in the given bucket, group, and artifact.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketName** | **string**| The bucket name | 
  **groupId** | **string**| The group identifier | 
  **artifactId** | **string**| The artifact identifier | 
  **version** | **string**| The version | 

### Return type

[**ExtensionRepoVersion**](ExtensionRepoVersion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtensionRepoVersionContent**
> string GetExtensionRepoVersionContent(ctx, bucketName, groupId, artifactId, version)
Get extension repo version content

Gets the binary content of the bundle with the given bucket, group, artifact, and version.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketName** | **string**| The bucket name | 
  **groupId** | **string**| The group identifier | 
  **artifactId** | **string**| The artifact identifier | 
  **version** | **string**| The version | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtensionRepoVersionExtension**
> Extension GetExtensionRepoVersionExtension(ctx, bucketName, groupId, artifactId, version, name)
Get extension repo extension

Gets information about the extension with the given name in the given bucket, group, artifact, and version.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketName** | **string**| The bucket name | 
  **groupId** | **string**| The group identifier | 
  **artifactId** | **string**| The artifact identifier | 
  **version** | **string**| The version | 
  **name** | **string**| The fully qualified name of the extension | 

### Return type

[**Extension**](Extension.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtensionRepoVersionExtensionAdditionalDetailsDocs**
> GetExtensionRepoVersionExtensionAdditionalDetailsDocs(ctx, bucketName, groupId, artifactId, version, name)
Get extension repo extension details

Gets the additional details documentation for the extension with the given name in the given bucket, group, artifact, and version.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketName** | **string**| The bucket name | 
  **groupId** | **string**| The group identifier | 
  **artifactId** | **string**| The artifact identifier | 
  **version** | **string**| The version | 
  **name** | **string**| The fully qualified name of the extension | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtensionRepoVersionExtensionDocs**
> GetExtensionRepoVersionExtensionDocs(ctx, bucketName, groupId, artifactId, version, name)
Get extension repo extension docs

Gets the documentation for the extension with the given name in the given bucket, group, artifact, and version.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketName** | **string**| The bucket name | 
  **groupId** | **string**| The group identifier | 
  **artifactId** | **string**| The artifact identifier | 
  **version** | **string**| The version | 
  **name** | **string**| The fully qualified name of the extension | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtensionRepoVersionExtensions**
> []ExtensionMetadata GetExtensionRepoVersionExtensions(ctx, bucketName, groupId, artifactId, version)
Get extension repo extensions

Gets information about the extensions in the given bucket, group, artifact, and version.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketName** | **string**| The bucket name | 
  **groupId** | **string**| The group identifier | 
  **artifactId** | **string**| The artifact identifier | 
  **version** | **string**| The version | 

### Return type

[**[]ExtensionMetadata**](ExtensionMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtensionRepoVersionSha256**
> GetExtensionRepoVersionSha256(ctx, bucketName, groupId, artifactId, version)
Get extension repo version checksum

Gets the hex representation of the SHA-256 digest for the binary content of the bundle with the given bucket, group, artifact, and version.  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketName** | **string**| The bucket name | 
  **groupId** | **string**| The group identifier | 
  **artifactId** | **string**| The artifact identifier | 
  **version** | **string**| The version | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtensionRepoVersions**
> []ExtensionRepoVersionSummary GetExtensionRepoVersions(ctx, bucketName, groupId, artifactId)
Get extension repo versions

Gets the versions in the extension repository for the given bucket, group, and artifact.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketName** | **string**| The bucket name | 
  **groupId** | **string**| The group identifier | 
  **artifactId** | **string**| The artifact identifier | 

### Return type

[**[]ExtensionRepoVersionSummary**](ExtensionRepoVersionSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGlobalExtensionRepoVersionSha256**
> GetGlobalExtensionRepoVersionSha256(ctx, groupId, artifactId, version)
Get global extension repo version checksum

Gets the hex representation of the SHA-256 digest for the binary content with the given bucket, group, artifact, and version. Since the same group-artifact-version can exist in multiple buckets, this will return the checksum of the first one returned. This will be consistent since the checksum must be the same when existing in multiple buckets.   NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**| The group identifier | 
  **artifactId** | **string**| The artifact identifier | 
  **version** | **string**| The version | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

