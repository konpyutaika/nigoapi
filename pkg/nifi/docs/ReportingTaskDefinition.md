# ReportingTaskDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | **string** | The group name of the bundle that provides the referenced type. | [optional] [default to null]
**Artifact** | **string** | The artifact name of the bundle that provides the referenced type. | [optional] [default to null]
**Version** | **string** | The version of the bundle that provides the referenced type. | [optional] [default to null]
**Type_** | **string** | The fully-qualified class type | [default to null]
**TypeDescription** | **string** | The description of the type. | [optional] [default to null]
**BuildInfo** | [***BuildInfo**](BuildInfo.md) | The build metadata for this component | [optional] [default to null]
**ProvidedApiImplementations** | [**[]DefinedType**](DefinedType.md) | If this type represents a provider for an interface, this lists the APIs it implements | [optional] [default to null]
**Tags** | **[]string** | The tags associated with this type | [optional] [default to null]
**Deprecated** | **bool** | Whether or not the component has been deprecated | [optional] [default to null]
**DeprecationReason** | **string** | If this component has been deprecated, this optional field can be used to provide an explanation | [optional] [default to null]
**PropertyDescriptors** | [**map[string]PropertyDescriptor**](PropertyDescriptor.md) | Descriptions of configuration properties applicable to this component. | [optional] [default to null]
**SupportsDynamicProperties** | **bool** | Whether or not this component makes use of dynamic (user-set) properties. | [optional] [default to null]
**SupportedSchedulingStrategies** | **[]string** | The supported scheduling strategies, such as TIME_DRIVER or CRON. | [optional] [default to null]
**DefaultSchedulingStrategy** | **string** | The default scheduling strategy for the reporting task. | [optional] [default to null]
**DefaultSchedulingPeriodBySchedulingStrategy** | **map[string]string** | The default scheduling period for each scheduling strategy. The scheduling period is expected to be a time period, such as \&quot;30 sec\&quot;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


