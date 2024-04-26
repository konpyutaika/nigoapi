# FlowConfigurationDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportsManagedAuthorizer** | **bool** | Whether this NiFi supports a managed authorizer. Managed authorizers can visualize users, groups, and policies in the UI. | [optional] [default to null]
**SupportsConfigurableAuthorizer** | **bool** | Whether this NiFi supports a configurable authorizer. | [optional] [default to null]
**SupportsConfigurableUsersAndGroups** | **bool** | Whether this NiFi supports configurable users and groups. | [optional] [default to null]
**AutoRefreshIntervalSeconds** | **int64** | The interval in seconds between the automatic NiFi refresh requests. | [optional] [default to null]
**CurrentTime** | **string** | The current time on the system. | [optional] [default to null]
**TimeOffset** | **int32** | The time offset of the system. | [optional] [default to null]
**DefaultBackPressureObjectThreshold** | **int64** | The default back pressure object threshold. | [optional] [default to null]
**DefaultBackPressureDataSizeThreshold** | **string** | The default back pressure data size threshold. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

