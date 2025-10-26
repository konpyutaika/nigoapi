# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSnippet**](SnippetsApi.md#CreateSnippet) | **Post** /snippets | Creates a snippet. The snippet will be automatically discarded if not used in a subsequent request after 1 minute.
[**DeleteSnippet**](SnippetsApi.md#DeleteSnippet) | **Delete** /snippets/{id} | Deletes the components in a snippet and discards the snippet
[**UpdateSnippet**](SnippetsApi.md#UpdateSnippet) | **Put** /snippets/{id} | Move&#x27;s the components in this Snippet into a new Process Group and discards the snippet

# **CreateSnippet**
> SnippetEntity CreateSnippet(ctx, body)
Creates a snippet. The snippet will be automatically discarded if not used in a subsequent request after 1 minute.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SnippetEntity**](SnippetEntity.md)| The snippet configuration details. | 

### Return type

[**SnippetEntity**](SnippetEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSnippet**
> SnippetEntity DeleteSnippet(ctx, id, optional)
Deletes the components in a snippet and discards the snippet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The snippet id. | 
 **optional** | ***SnippetsApiDeleteSnippetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnippetsApiDeleteSnippetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disconnectedNodeAcknowledged** | **optional.bool**| Acknowledges that this node is disconnected to allow for mutable requests to proceed. | [default to false]

### Return type

[**SnippetEntity**](SnippetEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSnippet**
> SnippetEntity UpdateSnippet(ctx, body, id)
Move's the components in this Snippet into a new Process Group and discards the snippet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SnippetEntity**](SnippetEntity.md)| The snippet configuration details. | 
  **id** | **string**| The snippet id. | 

### Return type

[**SnippetEntity**](SnippetEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

