# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInputContent**](ProvenanceEventsApi.md#GetInputContent) | **Get** /provenance-events/{id}/content/input | Gets the input content for a provenance event
[**GetLatestProvenanceEvents**](ProvenanceEventsApi.md#GetLatestProvenanceEvents) | **Get** /provenance-events/latest/{componentId} | Retrieves the latest cached Provenance Events for the specified component
[**GetOutputContent**](ProvenanceEventsApi.md#GetOutputContent) | **Get** /provenance-events/{id}/content/output | Gets the output content for a provenance event
[**GetProvenanceEvent**](ProvenanceEventsApi.md#GetProvenanceEvent) | **Get** /provenance-events/{id} | Gets a provenance event
[**SubmitReplay**](ProvenanceEventsApi.md#SubmitReplay) | **Post** /provenance-events/replays | Replays content from a provenance event
[**SubmitReplayLatestEvent**](ProvenanceEventsApi.md#SubmitReplayLatestEvent) | **Post** /provenance-events/latest/replays | Replays content from a provenance event

# **GetInputContent**
> GetInputContent(ctx, id, optional)
Gets the input content for a provenance event

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**LongParameter**](.md)| The provenance event id. | 
 **optional** | ***ProvenanceEventsApiGetInputContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvenanceEventsApiGetInputContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **range_** | **optional.string**| Range of bytes requested | 
 **clusterNodeId** | **optional.string**| The id of the node where the content exists if clustered. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLatestProvenanceEvents**
> LatestProvenanceEventsEntity GetLatestProvenanceEvents(ctx, componentId, optional)
Retrieves the latest cached Provenance Events for the specified component

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **componentId** | **string**| The ID of the component to retrieve the latest Provenance Events for. | 
 **optional** | ***ProvenanceEventsApiGetLatestProvenanceEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvenanceEventsApiGetLatestProvenanceEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.int32**| The number of events to limit the response to. Defaults to 10. | [default to 10]

### Return type

[**LatestProvenanceEventsEntity**](LatestProvenanceEventsEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOutputContent**
> GetOutputContent(ctx, id, optional)
Gets the output content for a provenance event

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**LongParameter**](.md)| The provenance event id. | 
 **optional** | ***ProvenanceEventsApiGetOutputContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvenanceEventsApiGetOutputContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **range_** | **optional.string**| Range of bytes requested | 
 **clusterNodeId** | **optional.string**| The id of the node where the content exists if clustered. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProvenanceEvent**
> ProvenanceEventEntity GetProvenanceEvent(ctx, id, optional)
Gets a provenance event

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**LongParameter**](.md)| The provenance event id. | 
 **optional** | ***ProvenanceEventsApiGetProvenanceEventOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvenanceEventsApiGetProvenanceEventOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterNodeId** | **optional.string**| The id of the node where this event exists if clustered. | 

### Return type

[**ProvenanceEventEntity**](ProvenanceEventEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
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

# **SubmitReplayLatestEvent**
> ReplayLastEventResponseEntity SubmitReplayLatestEvent(ctx, body)
Replays content from a provenance event

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReplayLastEventRequestEntity**](ReplayLastEventRequestEntity.md)| The replay request. | 

### Return type

[**ReplayLastEventResponseEntity**](ReplayLastEventResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

