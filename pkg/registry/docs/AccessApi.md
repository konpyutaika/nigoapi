# \AccessApi

All URIs are relative to *http://localhost/nifi-registry-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessTokenByTryingAllProviders**](AccessApi.md#CreateAccessTokenByTryingAllProviders) | **Post** /access/token | Create token trying all providers
[**CreateAccessTokenUsingBasicAuthCredentials**](AccessApi.md#CreateAccessTokenUsingBasicAuthCredentials) | **Post** /access/token/login | Create token using basic auth
[**CreateAccessTokenUsingIdentityProviderCredentials**](AccessApi.md#CreateAccessTokenUsingIdentityProviderCredentials) | **Post** /access/token/identity-provider | Create token using identity provider
[**CreateAccessTokenUsingKerberosTicket**](AccessApi.md#CreateAccessTokenUsingKerberosTicket) | **Post** /access/token/kerberos | Create token using kerberos
[**GetAccessStatus**](AccessApi.md#GetAccessStatus) | **Get** /access | Get access status
[**GetIdentityProviderUsageInstructions**](AccessApi.md#GetIdentityProviderUsageInstructions) | **Get** /access/token/identity-provider/usage | Get identity provider usage
[**LogOut**](AccessApi.md#LogOut) | **Delete** /access/logout | Performs a logout for other providers that have been issued a JWT.
[**OidcCallback**](AccessApi.md#OidcCallback) | **Get** /access/oidc/callback | Redirect/callback URI for processing the result of the OpenId Connect login sequence.
[**OidcExchange**](AccessApi.md#OidcExchange) | **Post** /access/oidc/exchange | Retrieves a JWT following a successful login sequence using the configured OpenId Connect provider.
[**OidcLogout**](AccessApi.md#OidcLogout) | **Delete** /access/oidc/logout | Performs a logout in the OpenId Provider.
[**OidcRequest**](AccessApi.md#OidcRequest) | **Get** /access/oidc/request | Initiates a request to authenticate through the configured OpenId Connect provider.
[**TestIdentityProviderRecognizesCredentialsFormat**](AccessApi.md#TestIdentityProviderRecognizesCredentialsFormat) | **Post** /access/token/identity-provider/test | Test identity provider


# **CreateAccessTokenByTryingAllProviders**
> string CreateAccessTokenByTryingAllProviders(ctx, )
Create token trying all providers

Creates a token for accessing the REST API via auto-detected method of verifying client identity claim credentials. The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts. The header, the body, and the signature. The expiration of the token is a contained within the body. The token can be used in the Authorization header in the format 'Authorization: Bearer <token>'.

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

# **CreateAccessTokenUsingBasicAuthCredentials**
> string CreateAccessTokenUsingBasicAuthCredentials(ctx, )
Create token using basic auth

Creates a token for accessing the REST API via username/password. The user credentials must be passed in standard HTTP Basic Auth format. That is: 'Authorization: Basic <credentials>', where <credentials> is the base64 encoded value of '<username>:<password>'. The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts. The header, the body, and the signature. The expiration of the token is a contained within the body. The token can be used in the Authorization header in the format 'Authorization: Bearer <token>'.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAccessTokenUsingIdentityProviderCredentials**
> string CreateAccessTokenUsingIdentityProviderCredentials(ctx, )
Create token using identity provider

Creates a token for accessing the REST API via a custom identity provider. The user credentials must be passed in a format understood by the custom identity provider, e.g., a third-party auth token in an HTTP header. The exact format of the user credentials expected by the custom identity provider can be discovered by 'GET /access/token/identity-provider/usage'. The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts. The header, the body, and the signature. The expiration of the token is a contained within the body. The token can be used in the Authorization header in the format 'Authorization: Bearer <token>'.

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

# **CreateAccessTokenUsingKerberosTicket**
> string CreateAccessTokenUsingKerberosTicket(ctx, )
Create token using kerberos

Creates a token for accessing the REST API via Kerberos Service Tickets or SPNEGO Tokens (which includes Kerberos Service Tickets). The token returned is formatted as a JSON Web Token (JWT). The token is base64 encoded and comprised of three parts. The header, the body, and the signature. The expiration of the token is a contained within the body. The token can be used in the Authorization header in the format 'Authorization: Bearer <token>'.

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

# **GetAccessStatus**
> CurrentUser GetAccessStatus(ctx, )
Get access status

Returns the current client's authenticated identity and permissions to top-level resources

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CurrentUser**](CurrentUser.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdentityProviderUsageInstructions**
> string GetIdentityProviderUsageInstructions(ctx, )
Get identity provider usage

Provides a description of how the currently configured identity provider expects credentials to be passed to POST /access/token/identity-provider

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

# **LogOut**
> LogOut(ctx, )
Performs a logout for other providers that have been issued a JWT.

  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

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

# **OidcCallback**
> OidcCallback(ctx, )
Redirect/callback URI for processing the result of the OpenId Connect login sequence.

  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

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

  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

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

  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

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

  NOTE: This endpoint is subject to change as NiFi Registry and its REST API evolve.

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

# **TestIdentityProviderRecognizesCredentialsFormat**
> string TestIdentityProviderRecognizesCredentialsFormat(ctx, )
Test identity provider

Tests the format of the credentials against this identity provider without preforming authentication on the credentials to validate them. The user credentials should be passed in a format understood by the custom identity provider as defined by 'GET /access/token/identity-provider/usage'.

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

