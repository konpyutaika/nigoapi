# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPeers**](SiteToSiteApi.md#GetPeers) | **Get** /site-to-site/peers | Returns the available Peers and its status of this NiFi
[**GetSiteToSiteDetails**](SiteToSiteApi.md#GetSiteToSiteDetails) | **Get** /site-to-site | Returns the details about this NiFi necessary to communicate via site to site

# **GetPeers**
> PeersEntity GetPeers(ctx, )
Returns the available Peers and its status of this NiFi

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PeersEntity**](PeersEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteToSiteDetails**
> ControllerEntity GetSiteToSiteDetails(ctx, )
Returns the details about this NiFi necessary to communicate via site to site

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ControllerEntity**](ControllerEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

