# RemoteProcessGroupDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] [default to null]
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] [default to null]
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] [default to null]
**Position** | [***PositionDto**](PositionDTO.md) | The position of this component in the UI if applicable. | [optional] [default to null]
**TargetUri** | **string** | The target URI of the remote process group. If target uri is not set, but uris are set, then returns the first url in the urls. If neither target uri nor uris are set, then returns null. | [optional] [default to null]
**TargetUris** | **string** | The target URI of the remote process group. If target uris is not set but target uri is set, then returns a collection containing the single target uri. If neither target uris nor uris are set, then returns null. | [optional] [default to null]
**TargetSecure** | **bool** | Whether the target is running securely. | [optional] [default to null]
**Name** | **string** | The name of the remote process group. | [optional] [default to null]
**Comments** | **string** | The comments for the remote process group. | [optional] [default to null]
**CommunicationsTimeout** | **string** | The time period used for the timeout when communicating with the target. | [optional] [default to null]
**YieldDuration** | **string** | When yielding, this amount of time must elapse before the remote process group is scheduled again. | [optional] [default to null]
**TransportProtocol** | **string** |  | [optional] [default to null]
**LocalNetworkInterface** | **string** | The local network interface to send/receive data. If not specified, any local address is used. If clustered, all nodes must have an interface with this identifier. | [optional] [default to null]
**ProxyHost** | **string** |  | [optional] [default to null]
**ProxyPort** | **int32** |  | [optional] [default to null]
**ProxyUser** | **string** |  | [optional] [default to null]
**ProxyPassword** | **string** |  | [optional] [default to null]
**AuthorizationIssues** | **[]string** | Any remote authorization issues for the remote process group. | [optional] [default to null]
**ValidationErrors** | **[]string** | The validation errors for the remote process group. These validation errors represent the problems with the remote process group that must be resolved before it can transmit. | [optional] [default to null]
**Transmitting** | **bool** | Whether the remote process group is actively transmitting. | [optional] [default to null]
**InputPortCount** | **int32** | The number of remote input ports currently available on the target. | [optional] [default to null]
**OutputPortCount** | **int32** | The number of remote output ports currently available on the target. | [optional] [default to null]
**ActiveRemoteInputPortCount** | **int32** | The number of active remote input ports. | [optional] [default to null]
**InactiveRemoteInputPortCount** | **int32** | The number of inactive remote input ports. | [optional] [default to null]
**ActiveRemoteOutputPortCount** | **int32** | The number of active remote output ports. | [optional] [default to null]
**InactiveRemoteOutputPortCount** | **int32** | The number of inactive remote output ports. | [optional] [default to null]
**FlowRefreshed** | **string** | The timestamp when this remote process group was last refreshed. | [optional] [default to null]
**Contents** | [***RemoteProcessGroupContentsDto**](RemoteProcessGroupContentsDTO.md) | The contents of the remote process group. Will contain available input/output ports. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


