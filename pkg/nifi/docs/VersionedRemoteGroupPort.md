# VersionedRemoteGroupPort

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The component&#39;s unique identifier | [optional] [default to null]
**Name** | **string** | The component&#39;s name | [optional] [default to null]
**Comments** | **string** | The user-supplied comments for the component | [optional] [default to null]
**Position** | [***Position**](Position.md) | The component&#39;s position on the graph | [optional] [default to null]
**RemoteGroupId** | **string** | The id of the remote process group that the port resides in. | [optional] [default to null]
**ConcurrentlySchedulableTaskCount** | **int32** | The number of task that may transmit flowfiles to the target port concurrently. | [optional] [default to null]
**UseCompression** | **bool** | Whether the flowfiles are compressed when sent to the target port. | [optional] [default to null]
**BatchSize** | [***BatchSize**](BatchSize.md) | The batch settings for data transmission. | [optional] [default to null]
**ComponentType** | **string** |  | [optional] [default to null]
**TargetId** | **string** | The ID of the port on the target NiFi instance | [optional] [default to null]
**ScheduledState** | **string** | The scheduled state of the component | [optional] [default to null]
**GroupIdentifier** | **string** | The ID of the Process Group that this component belongs to | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


