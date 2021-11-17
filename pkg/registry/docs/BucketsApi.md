# \BucketsApi

All URIs are relative to *http://localhost/nifi-registry-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBucket**](BucketsApi.md#CreateBucket) | **Post** /buckets | Create bucket
[**DeleteBucket**](BucketsApi.md#DeleteBucket) | **Delete** /buckets/{bucketId} | Delete bucket
[**GetAvailableBucketFields**](BucketsApi.md#GetAvailableBucketFields) | **Get** /buckets/fields | Get bucket fields
[**GetBucket**](BucketsApi.md#GetBucket) | **Get** /buckets/{bucketId} | Get bucket
[**GetBuckets**](BucketsApi.md#GetBuckets) | **Get** /buckets | Get all buckets
[**UpdateBucket**](BucketsApi.md#UpdateBucket) | **Put** /buckets/{bucketId} | Update bucket


# **CreateBucket**
> Bucket CreateBucket(ctx, body)
Create bucket



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Bucket**](Bucket.md)| The bucket to create | 

### Return type

[**Bucket**](Bucket.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBucket**
> Bucket DeleteBucket(ctx, version, bucketId, optional)
Delete bucket

Deletes the bucket with the given id, along with all objects stored in the bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**| The version is used to verify the client is working with the latest version of the entity. | 
  **bucketId** | **string**| The bucket identifier | 
 **optional** | ***BucketsApiDeleteBucketOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BucketsApiDeleteBucketOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **optional.String**| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 

### Return type

[**Bucket**](Bucket.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvailableBucketFields**
> Fields GetAvailableBucketFields(ctx, )
Get bucket fields

Retrieves bucket field names for searching or sorting on buckets.

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

# **GetBucket**
> Bucket GetBucket(ctx, bucketId)
Get bucket

Gets the bucket with the given id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketId** | **string**| The bucket identifier | 

### Return type

[**Bucket**](Bucket.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuckets**
> []Bucket GetBuckets(ctx, )
Get all buckets

The returned list will include only buckets for which the user is authorized.If the user is not authorized for any buckets, this returns an empty list.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Bucket**](Bucket.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBucket**
> Bucket UpdateBucket(ctx, bucketId, body)
Update bucket

Updates the bucket with the given id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketId** | **string**| The bucket identifier | 
  **body** | [**Bucket**](Bucket.md)| The updated bucket | 

### Return type

[**Bucket**](Bucket.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

