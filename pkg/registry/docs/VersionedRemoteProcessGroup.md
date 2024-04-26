# VersionedRemoteProcessGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The component&#x27;s unique identifier | [optional] [default to null]
**InstanceIdentifier** | **string** | The instance ID of an existing component that is described by this VersionedComponent, or null if this is not mapped to an instantiated component | [optional] [default to null]
**Name** | **string** | The component&#x27;s name | [optional] [default to null]
**Comments** | **string** | The user-supplied comments for the component | [optional] [default to null]
**Position** | [***Position**](Position.md) |  | [optional] [default to null]
**TargetUris** | **string** | The target URIs of the remote process group. If target uris is not set but target uri is set, then returns the single target uri. If neither target uris nor target uri is set, then returns null. | [optional] [default to null]
**CommunicationsTimeout** | **string** | The time period used for the timeout when communicating with the target. | [optional] [default to null]
**YieldDuration** | **string** | When yielding, this amount of time must elapse before the remote process group is scheduled again. | [optional] [default to null]
**TransportProtocol** | **string** | The Transport Protocol that is used for Site-to-Site communications | [optional] [default to null]
**LocalNetworkInterface** | **string** | The local network interface to send/receive data. If not specified, any local address is used. If clustered, all nodes must have an interface with this identifier. | [optional] [default to null]
**ProxyHost** | **string** |  | [optional] [default to null]
**ProxyPort** | **int32** |  | [optional] [default to null]
**ProxyUser** | **string** |  | [optional] [default to null]
**ProxyPassword** | **string** |  | [optional] [default to null]
**InputPorts** | [**[]VersionedRemoteGroupPort**](VersionedRemoteGroupPort.md) | A Set of Input Ports that can be connected to, in order to send data to the remote NiFi instance | [optional] [default to null]
**OutputPorts** | [**[]VersionedRemoteGroupPort**](VersionedRemoteGroupPort.md) | A Set of Output Ports that can be connected to, in order to pull data from the remote NiFi instance | [optional] [default to null]
**ComponentType** | **string** |  | [optional] [default to null]
**GroupIdentifier** | **string** | The ID of the Process Group that this component belongs to | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

