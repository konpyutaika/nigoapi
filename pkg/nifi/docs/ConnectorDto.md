# ConnectorDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveConfiguration** | [***ConnectorConfigurationDto**](ConnectorConfigurationDTO.md) |  | [optional] [default to null]
**AvailableActions** | [**[]ConnectorActionDto**](ConnectorActionDTO.md) | The available actions that can be performed on this Connector. | [optional] [default to null]
**Bundle** | [***BundleDto**](BundleDTO.md) |  | [optional] [default to null]
**ConfigurationUrl** | **string** | The URL for this connector&#x27;s configuration/wizard custom UI, if applicable. | [optional] [default to null]
**DetailsUrl** | **string** | The URL for this connector&#x27;s details custom UI, if applicable. | [optional] [default to null]
**ExtensionMissing** | **bool** | Whether the extension for this connector is missing. | [optional] [default to null]
**Id** | **string** | The id of the component. | [optional] [default to null]
**ManagedProcessGroupId** | **string** | The identifier of the root Process Group managed by this Connector. | [optional] [default to null]
**MultipleVersionsAvailable** | **bool** | Whether multiple versions of this connector are available. | [optional] [default to null]
**Name** | **string** | The name of the Connector. | [optional] [default to null]
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] [default to null]
**Position** | [***PositionDto**](PositionDTO.md) |  | [optional] [default to null]
**State** | **string** | The state of the Connector. | [optional] [default to null]
**Type_** | **string** | The fully qualified type of the Connector. | [optional] [default to null]
**ValidationErrors** | **[]string** | The validation errors for the connector. | [optional] [default to null]
**ValidationStatus** | **string** | Indicates whether the Connector is valid, invalid, or still in the process of validating | [optional] [default to null]
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] [default to null]
**WorkingConfiguration** | [***ConnectorConfigurationDto**](ConnectorConfigurationDTO.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

