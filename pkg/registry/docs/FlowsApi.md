# \FlowsApi

All URIs are relative to *http://localhost/nifi-registry-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvailableFlowFields**](FlowsApi.md#GetAvailableFlowFields) | **Get** /flows/fields | Get flow fields
[**GlobalGetFlow**](FlowsApi.md#GlobalGetFlow) | **Get** /flows/{flowId} | Get flow
[**GlobalGetFlowVersion**](FlowsApi.md#GlobalGetFlowVersion) | **Get** /flows/{flowId}/versions/{versionNumber} | Get flow version
[**GlobalGetFlowVersions**](FlowsApi.md#GlobalGetFlowVersions) | **Get** /flows/{flowId}/versions | Get flow versions
[**GlobalGetLatestFlowVersion**](FlowsApi.md#GlobalGetLatestFlowVersion) | **Get** /flows/{flowId}/versions/latest | Get latest flow version
[**GlobalGetLatestFlowVersionMetadata**](FlowsApi.md#GlobalGetLatestFlowVersionMetadata) | **Get** /flows/{flowId}/versions/latest/metadata | Get latest flow version metadata


# **GetAvailableFlowFields**
> Fields GetAvailableFlowFields(ctx, )
Get flow fields

Retrieves the flow field names that can be used for searching or sorting on flows.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Fields**](Fields.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalGetFlow**
> VersionedFlow GlobalGetFlow(ctx, flowId)
Get flow

Gets a flow by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **flowId** | **string**| The flow identifier | 

### Return type

[**VersionedFlow**](VersionedFlow.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalGetFlowVersion**
> VersionedFlowSnapshot GlobalGetFlowVersion(ctx, flowId, versionNumber)
Get flow version

Gets the given version of a flow, including metadata and flow content.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

# **GlobalGetFlowVersions**
> []VersionedFlowSnapshotMetadata GlobalGetFlowVersions(ctx, flowId)
Get flow versions

Gets summary information for all versions of a given flow. Versions are ordered newest->oldest.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **flowId** | **string**| The flow identifier | 

### Return type

[**[]VersionedFlowSnapshotMetadata**](VersionedFlowSnapshotMetadata.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalGetLatestFlowVersion**
> VersionedFlowSnapshot GlobalGetLatestFlowVersion(ctx, flowId)
Get latest flow version

Gets the latest version of a flow, including metadata and flow content.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **flowId** | **string**| The flow identifier | 

### Return type

[**VersionedFlowSnapshot**](VersionedFlowSnapshot.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalGetLatestFlowVersionMetadata**
> VersionedFlowSnapshotMetadata GlobalGetLatestFlowVersionMetadata(ctx, flowId)
Get latest flow version metadata

Gets the metadata for the latest version of a flow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **flowId** | **string**| The flow identifier | 

### Return type

[**VersionedFlowSnapshotMetadata**](VersionedFlowSnapshotMetadata.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

