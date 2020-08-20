# VersionedFlow

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | [***JaxbLink**](JaxbLink.md) | An WebLink to this entity. | [optional] [default to null]
**Identifier** | **string** | An ID to uniquely identify this object. | [optional] [default to null]
**Name** | **string** | The name of the item. | [default to null]
**Description** | **string** | A description of the item. | [optional] [default to null]
**BucketIdentifier** | **string** | The identifier of the bucket this items belongs to. This cannot be changed after the item is created. | [default to null]
**BucketName** | **string** | The name of the bucket this items belongs to. | [optional] [default to null]
**CreatedTimestamp** | **int64** | The timestamp of when the item was created, as milliseconds since epoch. | [optional] [default to null]
**ModifiedTimestamp** | **int64** | The timestamp of when the item was last modified, as milliseconds since epoch. | [optional] [default to null]
**Type_** | **string** | The type of item. | [default to null]
**Permissions** | [***Permissions**](Permissions.md) | The access that the current user has to the bucket containing this item. | [optional] [default to null]
**VersionCount** | **int64** | The number of versions of this flow. | [optional] [default to null]
**Revision** | [***RevisionInfo**](RevisionInfo.md) | The revision of this entity used for optimistic-locking during updates. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


