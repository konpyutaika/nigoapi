# RemoteProcessGroupPortDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the port. | [optional] [default to null]
**TargetId** | **string** | The id of the target port. | [optional] [default to null]
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] [default to null]
**GroupId** | **string** | The id of the remote process group that the port resides in. | [optional] [default to null]
**Name** | **string** | The name of the target port. | [optional] [default to null]
**Comments** | **string** | The comments as configured on the target port. | [optional] [default to null]
**ConcurrentlySchedulableTaskCount** | **int32** | The number of task that may transmit flowfiles to the target port concurrently. | [optional] [default to null]
**Transmitting** | **bool** | Whether the remote port is configured for transmission. | [optional] [default to null]
**UseCompression** | **bool** | Whether the flowfiles are compressed when sent to the target port. | [optional] [default to null]
**Exists** | **bool** | Whether the target port exists. | [optional] [default to null]
**TargetRunning** | **bool** | Whether the target port is running. | [optional] [default to null]
**Connected** | **bool** | Whether the port has either an incoming or outgoing connection. | [optional] [default to null]
**BatchSettings** | [***BatchSettingsDto**](BatchSettingsDTO.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

