# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommitInputPortTransaction**](DataTransferApi.md#CommitInputPortTransaction) | **Delete** /data-transfer/input-ports/{portId}/transactions/{transactionId} | Commit or cancel the specified transaction
[**CommitOutputPortTransaction**](DataTransferApi.md#CommitOutputPortTransaction) | **Delete** /data-transfer/output-ports/{portId}/transactions/{transactionId} | Commit or cancel the specified transaction
[**CreatePortTransaction**](DataTransferApi.md#CreatePortTransaction) | **Post** /data-transfer/{portType}/{portId}/transactions | Create a transaction to the specified output port or input port
[**ExtendInputPortTransactionTTL**](DataTransferApi.md#ExtendInputPortTransactionTTL) | **Put** /data-transfer/input-ports/{portId}/transactions/{transactionId} | Extend transaction TTL
[**ExtendOutputPortTransactionTTL**](DataTransferApi.md#ExtendOutputPortTransactionTTL) | **Put** /data-transfer/output-ports/{portId}/transactions/{transactionId} | Extend transaction TTL
[**ReceiveFlowFiles**](DataTransferApi.md#ReceiveFlowFiles) | **Post** /data-transfer/input-ports/{portId}/transactions/{transactionId}/flow-files | Transfer flow files to the input port
[**TransferFlowFiles**](DataTransferApi.md#TransferFlowFiles) | **Get** /data-transfer/output-ports/{portId}/transactions/{transactionId}/flow-files | Transfer flow files from the output port

# **CommitInputPortTransaction**
> TransactionResultEntity CommitInputPortTransaction(ctx, responseCode, portId, transactionId, optional)
Commit or cancel the specified transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **responseCode** | **int32**| The response code. Available values are BAD_CHECKSUM(19), CONFIRM_TRANSACTION(12) or CANCEL_TRANSACTION(15). | 
  **portId** | **string**| The input port id. | 
  **transactionId** | **string**| The transaction id. | 
 **optional** | ***DataTransferApiCommitInputPortTransactionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataTransferApiCommitInputPortTransactionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 

### Return type

[**TransactionResultEntity**](TransactionResultEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommitOutputPortTransaction**
> TransactionResultEntity CommitOutputPortTransaction(ctx, responseCode, checksum, portId, transactionId, optional)
Commit or cancel the specified transaction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **responseCode** | **int32**| The response code. Available values are CONFIRM_TRANSACTION(12) or CANCEL_TRANSACTION(15). | 
  **checksum** | **string**| A checksum calculated at client side using CRC32 to check flow file content integrity. It must match with the value calculated at server side. | 
  **portId** | **string**| The output port id. | 
  **transactionId** | **string**| The transaction id. | 
 **optional** | ***DataTransferApiCommitOutputPortTransactionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataTransferApiCommitOutputPortTransactionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 

### Return type

[**TransactionResultEntity**](TransactionResultEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePortTransaction**
> TransactionResultEntity CreatePortTransaction(ctx, portType, portId, optional)
Create a transaction to the specified output port or input port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portType** | **string**| The port type. | 
  **portId** | **string**|  | 
 **optional** | ***DataTransferApiCreatePortTransactionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataTransferApiCreatePortTransactionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 

### Return type

[**TransactionResultEntity**](TransactionResultEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtendInputPortTransactionTTL**
> TransactionResultEntity ExtendInputPortTransactionTTL(ctx, portId, transactionId, optional)
Extend transaction TTL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portId** | **string**|  | 
  **transactionId** | **string**|  | 
 **optional** | ***DataTransferApiExtendInputPortTransactionTTLOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataTransferApiExtendInputPortTransactionTTLOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 

### Return type

[**TransactionResultEntity**](TransactionResultEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtendOutputPortTransactionTTL**
> TransactionResultEntity ExtendOutputPortTransactionTTL(ctx, portId, transactionId, optional)
Extend transaction TTL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portId** | **string**|  | 
  **transactionId** | **string**|  | 
 **optional** | ***DataTransferApiExtendOutputPortTransactionTTLOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataTransferApiExtendOutputPortTransactionTTLOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 

### Return type

[**TransactionResultEntity**](TransactionResultEntity.md)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReceiveFlowFiles**
> string ReceiveFlowFiles(ctx, portId, transactionId, optional)
Transfer flow files to the input port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portId** | **string**| The input port id. | 
  **transactionId** | **string**|  | 
 **optional** | ***DataTransferApiReceiveFlowFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataTransferApiReceiveFlowFilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 

### Return type

**string**

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransferFlowFiles**
> TransferFlowFiles(ctx, portId, transactionId, optional)
Transfer flow files from the output port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **portId** | **string**| The output port id. | 
  **transactionId** | **string**|  | 
 **optional** | ***DataTransferApiTransferFlowFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataTransferApiTransferFlowFilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 

### Return type

 (empty response body)

### Authorization

[CookieSecureAuthorizationBearer](../README.md#CookieSecureAuthorizationBearer), [HTTPBearerJWT](../README.md#HTTPBearerJWT)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

