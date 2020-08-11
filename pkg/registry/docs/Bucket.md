# Bucket

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | [***JaxbLink**](JaxbLink.md) | An WebLink to this entity. | [optional] [default to null]
**Identifier** | **string** | An ID to uniquely identify this object. | [optional] [default to null]
**Name** | **string** | The name of the bucket. | [default to null]
**CreatedTimestamp** | **int64** | The timestamp of when the bucket was first created. This is set by the server at creation time. | [optional] [default to null]
**Description** | **string** | A description of the bucket. | [optional] [default to null]
**AllowBundleRedeploy** | **bool** | Indicates if this bucket allows the same version of an extension bundle to be redeployed and thus overwrite the existing artifact. By default this is false. | [optional] [default to null]
**AllowPublicRead** | **bool** | Indicates if this bucket allows read access to unauthenticated anonymous users | [optional] [default to null]
**Permissions** | [***Permissions**](Permissions.md) | The access that the current user has to this bucket. | [optional] [default to null]
**Revision** | [***RevisionInfo**](RevisionInfo.md) | The revision of this entity used for optimistic-locking during updates. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


