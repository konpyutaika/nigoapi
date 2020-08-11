# \TenantsApi

All URIs are relative to *http://localhost/nifi-registry-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](TenantsApi.md#CreateUser) | **Post** /tenants/users | Create user
[**CreateUserGroup**](TenantsApi.md#CreateUserGroup) | **Post** /tenants/user-groups | Create user group
[**GetUser**](TenantsApi.md#GetUser) | **Get** /tenants/users/{id} | Get user
[**GetUserGroup**](TenantsApi.md#GetUserGroup) | **Get** /tenants/user-groups/{id} | Get user group
[**GetUserGroups**](TenantsApi.md#GetUserGroups) | **Get** /tenants/user-groups | Get user groups
[**GetUsers**](TenantsApi.md#GetUsers) | **Get** /tenants/users | Get all users
[**RemoveUser**](TenantsApi.md#RemoveUser) | **Delete** /tenants/users/{id} | Delete user
[**RemoveUserGroup**](TenantsApi.md#RemoveUserGroup) | **Delete** /tenants/user-groups/{id} | Delete user group
[**UpdateUser**](TenantsApi.md#UpdateUser) | **Put** /tenants/users/{id} | Update user
[**UpdateUserGroup**](TenantsApi.md#UpdateUserGroup) | **Put** /tenants/user-groups/{id} | Update user group


# **CreateUser**
> User CreateUser(ctx, body)
Create user

  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**User**](User.md)| The user configuration details. | 

### Return type

[**User**](User.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUserGroup**
> UserGroup CreateUserGroup(ctx, body)
Create user group

  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserGroup**](UserGroup.md)| The user group configuration details. | 

### Return type

[**UserGroup**](UserGroup.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> User GetUser(ctx, id)
Get user

  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The user id. | 

### Return type

[**User**](User.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserGroup**
> UserGroup GetUserGroup(ctx, id)
Get user group

  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The user group id. | 

### Return type

[**UserGroup**](UserGroup.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserGroups**
> []UserGroup GetUserGroups(ctx, )
Get user groups

  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]UserGroup**](UserGroup.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsers**
> []User GetUsers(ctx, )
Get all users

  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]User**](User.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveUser**
> User RemoveUser(ctx, version, id, optional)
Delete user

  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**| The version is used to verify the client is working with the latest version of the entity. | 
  **id** | **string**| The user id. | 
 **optional** | ***TenantsApiRemoveUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenantsApiRemoveUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **optional.String**| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 

### Return type

[**User**](User.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveUserGroup**
> UserGroup RemoveUserGroup(ctx, version, id, optional)
Delete user group

  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**| The version is used to verify the client is working with the latest version of the entity. | 
  **id** | **string**| The user group id. | 
 **optional** | ***TenantsApiRemoveUserGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenantsApiRemoveUserGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **optional.String**| If the client id is not specified, new one will be generated. This value (whether specified or generated) is included in the response. | 

### Return type

[**UserGroup**](UserGroup.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> User UpdateUser(ctx, id, body)
Update user

  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The user id. | 
  **body** | [**User**](User.md)| The user configuration details. | 

### Return type

[**User**](User.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserGroup**
> UserGroup UpdateUserGroup(ctx, id, body)
Update user group

  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The user group id. | 
  **body** | [**UserGroup**](UserGroup.md)| The user group configuration details. | 

### Return type

[**UserGroup**](UserGroup.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

