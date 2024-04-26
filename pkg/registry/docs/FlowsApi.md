# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvailableFlowFields**](FlowsApi.md#GetAvailableFlowFields) | **Get** /flows/fields | Get flow fields
[**GetFlow1**](FlowsApi.md#GetFlow1) | **Get** /flows/{flowId} | Get flow
[**GetFlowVersion1**](FlowsApi.md#GetFlowVersion1) | **Get** /flows/{flowId}/versions/{versionNumber} | Get flow version
[**GetFlowVersions1**](FlowsApi.md#GetFlowVersions1) | **Get** /flows/{flowId}/versions | Get flow versions
[**GetLatestFlowVersion1**](FlowsApi.md#GetLatestFlowVersion1) | **Get** /flows/{flowId}/versions/latest | Get latest flow version
[**GetLatestFlowVersionMetadata1**](FlowsApi.md#GetLatestFlowVersionMetadata1) | **Get** /flows/{flowId}/versions/latest/metadata | Get latest flow version metadata

# **GetAvailableFlowFields**
> Fields GetAvailableFlowFields(ctx, )
Get flow fields

Retrieves the flow field names that can be used for searching or sorting on flows.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Fields**](Fields.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlow1**
> VersionedFlow GetFlow1(ctx, flowId)
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

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlowVersion1**
> VersionedFlowSnapshot GetFlowVersion1(ctx, flowId, versionNumber)
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

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFlowVersions1**
> []VersionedFlowSnapshotMetadata GetFlowVersions1(ctx, flowId)
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

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestFlowVersion1**
> VersionedFlowSnapshot GetLatestFlowVersion1(ctx, flowId)
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

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestFlowVersionMetadata1**
> VersionedFlowSnapshotMetadata GetLatestFlowVersionMetadata1(ctx, flowId)
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

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

