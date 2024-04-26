# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvailableBucketItemFields**](ItemsApi.md#GetAvailableBucketItemFields) | **Get** /items/fields | Get item fields
[**GetItems**](ItemsApi.md#GetItems) | **Get** /items | Get all items
[**GetItems1**](ItemsApi.md#GetItems1) | **Get** /items/{bucketId} | Get bucket items

# **GetAvailableBucketItemFields**
> Fields GetAvailableBucketItemFields(ctx, )
Get item fields

Retrieves the item field names for searching or sorting on bucket items.

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

# **GetItems**
> []BucketItem GetItems(ctx, )
Get all items

Get items across all buckets. The returned items will include only items from buckets for which the user is authorized. If the user is not authorized to any buckets, an empty list will be returned.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]BucketItem**](BucketItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetItems1**
> []BucketItem GetItems1(ctx, bucketId)
Get bucket items

Gets the items located in the given bucket.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketId** | **string**| The bucket identifier | 

### Return type

[**[]BucketItem**](BucketItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

