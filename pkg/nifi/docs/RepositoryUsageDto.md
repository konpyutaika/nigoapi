# RepositoryUsageDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the repository | [optional] [default to null]
**FileStoreHash** | **string** | A SHA-256 hash of the File Store name/path that is used to store the repository&#39;s data. This information is exposed as a hash in order to avoid exposing potentially sensitive information that is not generally relevant. What is typically relevant is whether or not multiple repositories on the same node are using the same File Store, as this indicates that the repositories are competing for the resources of the backing disk/storage mechanism. | [optional] [default to null]
**FreeSpace** | **string** | Amount of free space. | [optional] [default to null]
**TotalSpace** | **string** | Amount of total space. | [optional] [default to null]
**FreeSpaceBytes** | **int64** | The number of bytes of free space. | [optional] [default to null]
**TotalSpaceBytes** | **int64** | The number of bytes of total space. | [optional] [default to null]
**Utilization** | **string** | Utilization of this storage location. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


