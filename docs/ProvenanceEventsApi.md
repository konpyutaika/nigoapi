# \ProvenanceEventsApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInputContent**](ProvenanceEventsApi.md#GetInputContent) | **Get** /provenance-events/{id}/content/input | Gets the input content for a provenance event
[**GetOutputContent**](ProvenanceEventsApi.md#GetOutputContent) | **Get** /provenance-events/{id}/content/output | Gets the output content for a provenance event
[**GetProvenanceEvent**](ProvenanceEventsApi.md#GetProvenanceEvent) | **Get** /provenance-events/{id} | Gets a provenance event
[**SubmitReplay**](ProvenanceEventsApi.md#SubmitReplay) | **Post** /provenance-events/replays | Replays content from a provenance event


# **GetInputContent**
> StreamingOutput GetInputContent(ctx, id, optional)
Gets the input content for a provenance event



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The provenance event id. | 
 **optional** | ***ProvenanceEventsApiGetInputContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvenanceEventsApiGetInputContentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterNodeId** | **optional.String**| The id of the node where the content exists if clustered. | 

### Return type

[**StreamingOutput**](StreamingOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOutputContent**
> StreamingOutput GetOutputContent(ctx, id, optional)
Gets the output content for a provenance event



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The provenance event id. | 
 **optional** | ***ProvenanceEventsApiGetOutputContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvenanceEventsApiGetOutputContentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterNodeId** | **optional.String**| The id of the node where the content exists if clustered. | 

### Return type

[**StreamingOutput**](StreamingOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProvenanceEvent**
> ProvenanceEventEntity GetProvenanceEvent(ctx, id, optional)
Gets a provenance event



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The provenance event id. | 
 **optional** | ***ProvenanceEventsApiGetProvenanceEventOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvenanceEventsApiGetProvenanceEventOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterNodeId** | **optional.String**| The id of the node where this event exists if clustered. | 

### Return type

[**ProvenanceEventEntity**](ProvenanceEventEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitReplay**
> ProvenanceEventEntity SubmitReplay(ctx, body)
Replays content from a provenance event



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubmitReplayRequestEntity**](SubmitReplayRequestEntity.md)| The replay request. | 

### Return type

[**ProvenanceEventEntity**](ProvenanceEventEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

