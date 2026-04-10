# ListenPortDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationProtocols** | **[]string** | Supported application protocols, if applicable | [optional] [default to null]
**ComponentClass** | **string** | The class type of the component providing the listen port | [optional] [default to null]
**ComponentId** | **string** | The id of the component providing the listen port | [optional] [default to null]
**ComponentName** | **string** | The name of the component providing the listen port | [optional] [default to null]
**ComponentType** | **string** | The type of component providing the listen port (e.g., Processor, ControllerService) | [optional] [default to null]
**ParentGroupId** | **string** | The id of the process group containing the component providing the listen port, if applicable | [optional] [default to null]
**ParentGroupName** | **string** | The name of the process group containing the component providing the listen port, if applicable | [optional] [default to null]
**PortName** | **string** | The name of the the listen port. Useful context for components that provide multiple ports. | [optional] [default to null]
**PortNumber** | **int32** | The ingress port number | [optional] [default to null]
**TransportProtocol** | **string** | The ingress transport protocol (TCP or UDP) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

