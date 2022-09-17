# \BucketFlowsApi

All URIs are relative to *http://localhost/nifi-registry-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFlow**](BucketFlowsApi.md#CreateFlow) | **Post** /buckets/{bucketId}/flows | Create flow
[**CreateFlowVersion**](BucketFlowsApi.md#CreateFlowVersion) | **Post** /buckets/{bucketId}/flows/{flowId}/versions | Create flow version
[**DeleteFlow**](BucketFlowsApi.md#DeleteFlow) | **Delete** /buckets/{bucketId}/flows/{flowId} | Delete bucket flow
[**GetFlow**](BucketFlowsApi.md#GetFlow) | **Get** /buckets/{bucketId}/flows/{flowId} | Get bucket flow
[**GetFlowDiff**](BucketFlowsApi.md#GetFlowDiff) | **Get** /buckets/{bucketId}/flows/{flowId}/diff/{versionA}/{versionB} | Get bucket flow diff
[**GetFlowVersion**](BucketFlowsApi.md#GetFlowVersion) | **Get** /buckets/{bucketId}/flows/{flowId}/versions/{versionNumber} | Get bucket flow version
[**GetFlowVersions**](BucketFlowsApi.md#GetFlowVersions) | **Get** /buckets/{bucketId}/flows/{flowId}/versions | Get bucket flow versions
[**GetFlows**](BucketFlowsApi.md#GetFlows) | **Get** /buckets/{bucketId}/flows | Get bucket flows
[**GetLatestFlowVersion**](BucketFlowsApi.md#GetLatestFlowVersion) | **Get** /buckets/{bucketId}/flows/{flowId}/versions/latest | Get latest bucket flow version content
[**GetLatestFlowVersionMetadata**](BucketFlowsApi.md#GetLatestFlowVersionMetadata) | **Get** /buckets/{bucketId}/flows/{flowId}/versions/latest/metadata | Get latest bucket flow version metadata
[**UpdateFlow**](BucketFlowsApi.md#UpdateFlow) | **Put** /buckets/{bucketId}/flows/{flowId} | Update bucket flow


# **CreateFlow**
> VersionedFlow CreateFlow(ctx, bucketId, body)
Create flow

Creates a flow in the given bucket. The flow id is created by the server and populated in the returned entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketId** | **string**| The bucket identifier | 
  **body** | [**VersionedFlow**](VersionedFlow.md)| The details of the flow to create. | 

### Return type

[**VersionedFlow**](VersionedFlow.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFlowVersion**
> VersionedFlowSnapshot CreateFlowVersion(ctx, bucketId, flowId, body)
Create flow version

Creates the next version of a flow. The version number of the object being created must be the next available version integer. Flow versions are immutable after they are created.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketId** | **string**| The bucket identifier | 
  **flowId** | **string**| The flow identifier | 
  **body** | [**VersionedFlowSnapshot**](VersionedFlowSnapshot.md)| The new versioned flow snapshot. | 

### Return type

[**VersionedFlowSnapshot**](VersionedFlowSnapshot.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFlow**
> VersionedFlow DeleteFlow(ctx, version, bucketId, flowId, optional)
Delete bucket flow

Deletes a flow, including all saved versions of that flow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**| The version is used to verify the client is working with the latest version of the entity. | 
  **bucketId** | **string**| The bucket identifier | 
  **flowId** | **string**| The flow identifier | 
 **optional** | ***BucketFlowsApiDeleteFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BucketFlowsApiDeleteFlowOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientId** | **optional.String**| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 

### Return type

[**VersionedFlow**](VersionedFlow.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlow**
> VersionedFlow GetFlow(ctx, bucketId, flowId)
Get bucket flow

Retrieves the flow with the given id in the given bucket.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketId** | **string**| The bucket identifier | 
  **flowId** | **string**| The flow identifier | 

### Return type

[**VersionedFlow**](VersionedFlow.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlowDiff**
> VersionedFlowDifference GetFlowDiff(ctx, bucketId, flowId, versionA, versionB)
Get bucket flow diff

Computes the differences between two given versions of a flow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketId** | **string**| The bucket identifier | 
  **flowId** | **string**| The flow identifier | 
  **versionA** | **int32**| The first version number | 
  **versionB** | **int32**| The second version number | 

### Return type

[**VersionedFlowDifference**](VersionedFlowDifference.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlowVersion**
> VersionedFlowSnapshot GetFlowVersion(ctx, bucketId, flowId, versionNumber)
Get bucket flow version

Gets the given version of a flow, including the metadata and content for the version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketId** | **string**| The bucket identifier | 
  **flowId** | **string**| The flow identifier | 
  **versionNumber** | **int32**| The version number | 

### Return type

[**VersionedFlowSnapshot**](VersionedFlowSnapshot.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlowVersions**
> []VersionedFlowSnapshotMetadata GetFlowVersions(ctx, bucketId, flowId)
Get bucket flow versions

Gets summary information for all versions of a flow. Versions are ordered newest->oldest.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketId** | **string**| The bucket identifier | 
  **flowId** | **string**| The flow identifier | 

### Return type

[**[]VersionedFlowSnapshotMetadata**](VersionedFlowSnapshotMetadata.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlows**
> []VersionedFlow GetFlows(ctx, bucketId)
Get bucket flows

Retrieves all flows in the given bucket.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketId** | **string**| The bucket identifier | 

### Return type

[**[]VersionedFlow**](VersionedFlow.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestFlowVersion**
> VersionedFlowSnapshot GetLatestFlowVersion(ctx, bucketId, flowId)
Get latest bucket flow version content

Gets the latest version of a flow, including the metadata and content of the flow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketId** | **string**| The bucket identifier | 
  **flowId** | **string**| The flow identifier | 

### Return type

[**VersionedFlowSnapshot**](VersionedFlowSnapshot.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestFlowVersionMetadata**
> VersionedFlowSnapshotMetadata GetLatestFlowVersionMetadata(ctx, bucketId, flowId)
Get latest bucket flow version metadata

Gets the metadata for the latest version of a flow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketId** | **string**| The bucket identifier | 
  **flowId** | **string**| The flow identifier | 

### Return type

[**VersionedFlowSnapshotMetadata**](VersionedFlowSnapshotMetadata.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFlow**
> VersionedFlow UpdateFlow(ctx, bucketId, flowId, body)
Update bucket flow

Updates the flow with the given id in the given bucket.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketId** | **string**| The bucket identifier | 
  **flowId** | **string**| The flow identifier | 
  **body** | [**VersionedFlow**](VersionedFlow.md)| The updated flow | 

### Return type

[**VersionedFlow**](VersionedFlow.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

