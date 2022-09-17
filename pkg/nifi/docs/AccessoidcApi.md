# \AccessoidcApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OidcCallback**](AccessoidcApi.md#OidcCallback) | **Get** /access/oidc/callback | Redirect/callback URI for processing the result of the OpenId Connect login sequence.
[**OidcExchange**](AccessoidcApi.md#OidcExchange) | **Post** /access/oidc/exchange | Retrieves a JWT following a successful login sequence using the configured OpenId Connect provider.
[**OidcLogout**](AccessoidcApi.md#OidcLogout) | **Get** /access/oidc/logout | Performs a logout in the OpenId Provider.
[**OidcLogoutCallback**](AccessoidcApi.md#OidcLogoutCallback) | **Get** /access/oidc/logoutCallback | Redirect/callback URI for processing the result of the OpenId Connect logout sequence.
[**OidcRequest**](AccessoidcApi.md#OidcRequest) | **Get** /access/oidc/request | Initiates a request to authenticate through the configured OpenId Connect provider.


# **OidcCallback**
> OidcCallback(ctx, )
Redirect/callback URI for processing the result of the OpenId Connect login sequence.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OidcExchange**
> string OidcExchange(ctx, )
Retrieves a JWT following a successful login sequence using the configured OpenId Connect provider.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OidcLogout**
> OidcLogout(ctx, )
Performs a logout in the OpenId Provider.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OidcLogoutCallback**
> OidcLogoutCallback(ctx, )
Redirect/callback URI for processing the result of the OpenId Connect logout sequence.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OidcRequest**
> OidcRequest(ctx, )
Initiates a request to authenticate through the configured OpenId Connect provider.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

