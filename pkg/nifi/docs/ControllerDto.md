# ControllerDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the NiFi. | [optional] [default to null]
**Name** | **string** | The name of the NiFi. | [optional] [default to null]
**Comments** | **string** | The comments for the NiFi. | [optional] [default to null]
**RunningCount** | **int32** | The number of running components in the NiFi. | [optional] [default to null]
**StoppedCount** | **int32** | The number of stopped components in the NiFi. | [optional] [default to null]
**InvalidCount** | **int32** | The number of invalid components in the NiFi. | [optional] [default to null]
**DisabledCount** | **int32** | The number of disabled components in the NiFi. | [optional] [default to null]
**ActiveRemotePortCount** | **int32** | The number of active remote ports contained in the NiFi. | [optional] [default to null]
**InactiveRemotePortCount** | **int32** | The number of inactive remote ports contained in the NiFi. | [optional] [default to null]
**InputPortCount** | **int32** | The number of input ports contained in the NiFi. | [optional] [default to null]
**OutputPortCount** | **int32** | The number of output ports in the NiFi. | [optional] [default to null]
**RemoteSiteListeningPort** | **int32** | The Socket Port on which this instance is listening for Remote Transfers of Flow Files. If this instance is not configured to receive Flow Files from remote instances, this will be null. | [optional] [default to null]
**RemoteSiteHttpListeningPort** | **int32** | The HTTP(S) Port on which this instance is listening for Remote Transfers of Flow Files. If this instance is not configured to receive Flow Files from remote instances, this will be null. | [optional] [default to null]
**SiteToSiteSecure** | **bool** | Indicates whether or not Site-to-Site communications with this instance is secure (2-way authentication). | [optional] [default to null]
**InstanceId** | **string** | If clustered, the id of the Cluster Manager, otherwise the id of the NiFi. | [optional] [default to null]
**InputPorts** | [**[]PortDto**](PortDTO.md) | The input ports available to send data to for the NiFi. | [optional] [default to null]
**OutputPorts** | [**[]PortDto**](PortDTO.md) | The output ports available to received data from the NiFi. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

