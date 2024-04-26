# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessToken**](AccessApi.md#CreateAccessToken) | **Post** /access/token | Creates a token for accessing the REST API via username/password
[**CreateAccessTokenFromTicket**](AccessApi.md#CreateAccessTokenFromTicket) | **Post** /access/kerberos | Creates a token for accessing the REST API via Kerberos ticket exchange / SPNEGO negotiation
[**GetAccessStatus**](AccessApi.md#GetAccessStatus) | **Get** /access | Gets the status the client&#x27;s access
[**GetAccessTokenExpiration**](AccessApi.md#GetAccessTokenExpiration) | **Get** /access/token/expiration | Get expiration for current Access Token
[**GetLoginConfig**](AccessApi.md#GetLoginConfig) | **Get** /access/config | Retrieves the access configuration for this NiFi
[**KnoxCallback**](AccessApi.md#KnoxCallback) | **Get** /access/knox/callback | Redirect/callback URI for processing the result of the Apache Knox login sequence.
[**KnoxLogout**](AccessApi.md#KnoxLogout) | **Get** /access/knox/logout | Performs a logout in the Apache Knox.
[**KnoxRequest**](AccessApi.md#KnoxRequest) | **Get** /access/knox/request | Initiates a request to authenticate through Apache Knox.
[**LogOut**](AccessApi.md#LogOut) | **Delete** /access/logout | Performs a logout for other providers that have been issued a JWT.
[**LogOutComplete**](AccessApi.md#LogOutComplete) | **Get** /access/logout/complete | Completes the logout sequence by removing the cached Logout Request and Cookie if they existed and redirects to /nifi/login.

# **CreateAccessToken**
> string CreateAccessToken(ctx, optional)
Creates a token for accessing the REST API via username/password

The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts. The header, the body, and the signature. The expiration of the token is a contained within the body. It is stored in the browser as a cookie, but also returned inthe response body to be stored/used by third party client scripts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccessApiCreateAccessTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccessApiCreateAccessTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.string**|  | 
 **password** | **optional.string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAccessTokenFromTicket**
> string CreateAccessTokenFromTicket(ctx, )
Creates a token for accessing the REST API via Kerberos ticket exchange / SPNEGO negotiation

The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts. The header, the body, and the signature. The expiration of the token is a contained within the body. The token can be used in the Authorization header in the format 'Authorization: Bearer <token>'. It is also stored in the browser as a cookie.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccessStatus**
> AccessStatusEntity GetAccessStatus(ctx, )
Gets the status the client's access

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AccessStatusEntity**](AccessStatusEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccessTokenExpiration**
> GetAccessTokenExpiration(ctx, )
Get expiration for current Access Token

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoginConfig**
> AccessConfigurationEntity GetLoginConfig(ctx, )
Retrieves the access configuration for this NiFi

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AccessConfigurationEntity**](AccessConfigurationEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KnoxCallback**
> KnoxCallback(ctx, )
Redirect/callback URI for processing the result of the Apache Knox login sequence.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KnoxLogout**
> KnoxLogout(ctx, )
Performs a logout in the Apache Knox.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KnoxRequest**
> KnoxRequest(ctx, )
Initiates a request to authenticate through Apache Knox.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogOut**
> LogOut(ctx, )
Performs a logout for other providers that have been issued a JWT.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogOutComplete**
> LogOutComplete(ctx, )
Completes the logout sequence by removing the cached Logout Request and Cookie if they existed and redirects to /nifi/login.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

