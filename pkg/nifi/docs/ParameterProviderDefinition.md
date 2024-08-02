# ParameterProviderDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | **string** | The group name of the bundle that provides the referenced type. | [optional] [default to null]
**Artifact** | **string** | The artifact name of the bundle that provides the referenced type. | [optional] [default to null]
**Version** | **string** | The version of the bundle that provides the referenced type. | [optional] [default to null]
**Type_** | **string** | The fully-qualified class type | [optional] [default to null]
**TypeDescription** | **string** | The description of the type. | [optional] [default to null]
**BuildInfo** | [***BuildInfo**](BuildInfo.md) |  | [optional] [default to null]
**ProvidedApiImplementations** | [**[]DefinedType**](DefinedType.md) | If this type represents a provider for an interface, this lists the APIs it implements | [optional] [default to null]
**Tags** | **[]string** | The tags associated with this type | [optional] [default to null]
**SeeAlso** | **[]string** | The names of other component types that may be related | [optional] [default to null]
**Deprecated** | **bool** | Whether or not the component has been deprecated | [optional] [default to null]
**DeprecationReason** | **string** | If this component has been deprecated, this optional field can be used to provide an explanation | [optional] [default to null]
**DeprecationAlternatives** | **[]string** | If this component has been deprecated, this optional field provides alternatives to use | [optional] [default to null]
**Restricted** | **bool** | Whether or not the component has a general restriction | [optional] [default to null]
**RestrictedExplanation** | **string** | An optional description of the general restriction | [optional] [default to null]
**ExplicitRestrictions** | [**[]Restriction**](Restriction.md) | Explicit restrictions that indicate a require permission to use the component | [optional] [default to null]
**Stateful** | [***Stateful**](Stateful.md) |  | [optional] [default to null]
**SystemResourceConsiderations** | [**[]SystemResourceConsideration**](SystemResourceConsideration.md) | The system resource considerations for the given component | [optional] [default to null]
**AdditionalDetails** | **bool** | Indicates if the component has additional details documentation | [optional] [default to null]
**PropertyDescriptors** | [**map[string]PropertyDescriptor**](PropertyDescriptor.md) | Descriptions of configuration properties applicable to this component. | [optional] [default to null]
**SupportsDynamicProperties** | **bool** | Whether or not this component makes use of dynamic (user-set) properties. | [optional] [default to null]
**SupportsSensitiveDynamicProperties** | **bool** | Whether or not this component makes use of sensitive dynamic (user-set) properties. | [optional] [default to null]
**DynamicProperties** | [**[]DynamicProperty**](DynamicProperty.md) | Describes the dynamic properties supported by this component | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

