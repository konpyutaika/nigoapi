# \AccesssamlApi

All URIs are relative to *http://localhost/nifi-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SamlLocalLogout**](AccesssamlApi.md#SamlLocalLogout) | **Get** /access/saml/local-logout | Local logout when SAML is enabled, does not communicate with the IDP.
[**SamlLoginExchange**](AccesssamlApi.md#SamlLoginExchange) | **Post** /access/saml/login/exchange | Retrieves a JWT following a successful login sequence using the configured SAML identity provider.
[**SamlLoginHttpPostConsumer**](AccesssamlApi.md#SamlLoginHttpPostConsumer) | **Post** /access/saml/login/consumer | Processes the SSO response from the SAML identity provider for HTTP-POST binding.
[**SamlLoginHttpRedirectConsumer**](AccesssamlApi.md#SamlLoginHttpRedirectConsumer) | **Get** /access/saml/login/consumer | Processes the SSO response from the SAML identity provider for HTTP-REDIRECT binding.
[**SamlLoginRequest**](AccesssamlApi.md#SamlLoginRequest) | **Get** /access/saml/login/request | Initiates an SSO request to the configured SAML identity provider.
[**SamlMetadata**](AccesssamlApi.md#SamlMetadata) | **Get** /access/saml/metadata | Retrieves the service provider metadata.
[**SamlSingleLogoutHttpPostConsumer**](AccesssamlApi.md#SamlSingleLogoutHttpPostConsumer) | **Post** /access/saml/single-logout/consumer | Processes a SingleLogout message from the configured SAML identity provider using the HTTP-POST binding.
[**SamlSingleLogoutHttpRedirectConsumer**](AccesssamlApi.md#SamlSingleLogoutHttpRedirectConsumer) | **Get** /access/saml/single-logout/consumer | Processes a SingleLogout message from the configured SAML identity provider using the HTTP-REDIRECT binding.
[**SamlSingleLogoutRequest**](AccesssamlApi.md#SamlSingleLogoutRequest) | **Get** /access/saml/single-logout/request | Initiates a logout request using the SingleLogout service of the configured SAML identity provider.


# **SamlLocalLogout**
> SamlLocalLogout(ctx, )
Local logout when SAML is enabled, does not communicate with the IDP.

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

# **SamlLoginExchange**
> string SamlLoginExchange(ctx, )
Retrieves a JWT following a successful login sequence using the configured SAML identity provider.

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

# **SamlLoginHttpPostConsumer**
> SamlLoginHttpPostConsumer(ctx, )
Processes the SSO response from the SAML identity provider for HTTP-POST binding.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SamlLoginHttpRedirectConsumer**
> SamlLoginHttpRedirectConsumer(ctx, )
Processes the SSO response from the SAML identity provider for HTTP-REDIRECT binding.

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

# **SamlLoginRequest**
> SamlLoginRequest(ctx, )
Initiates an SSO request to the configured SAML identity provider.

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

# **SamlMetadata**
> SamlMetadata(ctx, )
Retrieves the service provider metadata.

Note: This endpoint is subject to change as NiFi and it's REST API evolve.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/samlmetadata+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SamlSingleLogoutHttpPostConsumer**
> SamlSingleLogoutHttpPostConsumer(ctx, )
Processes a SingleLogout message from the configured SAML identity provider using the HTTP-POST binding.

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

# **SamlSingleLogoutHttpRedirectConsumer**
> SamlSingleLogoutHttpRedirectConsumer(ctx, )
Processes a SingleLogout message from the configured SAML identity provider using the HTTP-REDIRECT binding.

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

# **SamlSingleLogoutRequest**
> SamlSingleLogoutRequest(ctx, )
Initiates a logout request using the SingleLogout service of the configured SAML identity provider.

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

