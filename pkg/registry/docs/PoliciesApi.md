# \PoliciesApi

All URIs are relative to *http://localhost/nifi-registry-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessPolicy**](PoliciesApi.md#CreateAccessPolicy) | **Post** /policies | Create access policy
[**GetAccessPolicies**](PoliciesApi.md#GetAccessPolicies) | **Get** /policies | Get all access policies
[**GetAccessPolicy**](PoliciesApi.md#GetAccessPolicy) | **Get** /policies/{id} | Get access policy
[**GetAccessPolicyForResource**](PoliciesApi.md#GetAccessPolicyForResource) | **Get** /policies/{action}/{resource} | Get access policy for resource
[**GetResources**](PoliciesApi.md#GetResources) | **Get** /policies/resources | Get available resources
[**RemoveAccessPolicy**](PoliciesApi.md#RemoveAccessPolicy) | **Delete** /policies/{id} | Delete access policy
[**UpdateAccessPolicy**](PoliciesApi.md#UpdateAccessPolicy) | **Put** /policies/{id} | Update access policy


# **CreateAccessPolicy**
> AccessPolicy CreateAccessPolicy(ctx, body)
Create access policy



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccessPolicy**](AccessPolicy.md)| The access policy configuration details. | 

### Return type

[**AccessPolicy**](AccessPolicy.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccessPolicies**
> []AccessPolicy GetAccessPolicies(ctx, )
Get all access policies



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AccessPolicy**](AccessPolicy.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccessPolicy**
> AccessPolicy GetAccessPolicy(ctx, id)
Get access policy



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The access policy id. | 

### Return type

[**AccessPolicy**](AccessPolicy.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccessPolicyForResource**
> AccessPolicy GetAccessPolicyForResource(ctx, action, resource)
Get access policy for resource

Gets an access policy for the specified action and resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **action** | **string**| The request action. | 
  **resource** | **string**| The resource of the policy. | 

### Return type

[**AccessPolicy**](AccessPolicy.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResources**
> []Resource GetResources(ctx, )
Get available resources

Gets the available resources that support access/authorization policies

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Resource**](Resource.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveAccessPolicy**
> AccessPolicy RemoveAccessPolicy(ctx, version, id, optional)
Delete access policy



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**| The version is used to verify the client is working with the latest version of the entity. | 
  **id** | **string**| The access policy id. | 
 **optional** | ***PoliciesApiRemoveAccessPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoliciesApiRemoveAccessPolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **optional.String**| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 

### Return type

[**AccessPolicy**](AccessPolicy.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccessPolicy**
> AccessPolicy UpdateAccessPolicy(ctx, id, body)
Update access policy



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The access policy id. | 
  **body** | [**AccessPolicy**](AccessPolicy.md)| The access policy configuration details. | 

### Return type

[**AccessPolicy**](AccessPolicy.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

