# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLabel**](LabelsApi.md#GetLabel) | **Get** /labels/{id} | Gets a label
[**RemoveLabel**](LabelsApi.md#RemoveLabel) | **Delete** /labels/{id} | Deletes a label
[**UpdateLabel**](LabelsApi.md#UpdateLabel) | **Put** /labels/{id} | Updates a label

# **GetLabel**
> LabelEntity GetLabel(ctx, id)
Gets a label

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The label id. | 

### Return type

[**LabelEntity**](LabelEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveLabel**
> LabelEntity RemoveLabel(ctx, id, optional)
Deletes a label

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The label id. | 
 **optional** | ***LabelsApiRemoveLabelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelsApiRemoveLabelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | [**optional.Interface of LongParameter**](.md)| The revision is used to verify the client is working with the latest version of the flow. | 
 **clientId** | [**optional.Interface of ClientIdParameter**](.md)| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 
 **disconnectedNodeAcknowledged** | **optional.bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**LabelEntity**](LabelEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLabel**
> LabelEntity UpdateLabel(ctx, body, id)
Updates a label

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LabelEntity**](LabelEntity.md)| The label configuration details. | 
  **id** | **string**| The label id. | 

### Return type

[**LabelEntity**](LabelEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

