# BundleVersionMetadata

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | [***Link**](Link.md) |  | [optional] [default to null]
**Id** | **string** | The id of this version of the extension bundle | [default to null]
**BundleId** | **string** | The id of the extension bundle this version is for | [default to null]
**BucketId** | **string** | The id of the bucket the extension bundle belongs to | [default to null]
**GroupId** | **string** |  | [optional] [default to null]
**ArtifactId** | **string** |  | [optional] [default to null]
**Version** | **string** | The version of the extension bundle | [default to null]
**Timestamp** | **int64** | The timestamp of the create date of this version | [optional] [default to null]
**Author** | **string** | The identity that created this version | [default to null]
**Description** | **string** | The description for this version | [optional] [default to null]
**Sha256** | **string** | The hex representation of the SHA-256 digest of the binary content for this version | [default to null]
**Sha256Supplied** | **bool** | Whether or not the client supplied a SHA-256 when uploading the bundle | [default to null]
**ContentSize** | **int64** | The size of the binary content for this version in bytes | [default to null]
**SystemApiVersion** | **string** | The version of the system API that this bundle version was built against | [default to null]
**BuildInfo** | [***BuildInfo**](BuildInfo.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

