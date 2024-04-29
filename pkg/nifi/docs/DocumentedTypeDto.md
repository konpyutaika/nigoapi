# DocumentedTypeDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | The fully qualified name of the type. | [optional] [default to null]
**Bundle** | [***BundleDto**](BundleDTO.md) |  | [optional] [default to null]
**ControllerServiceApis** | [**[]ControllerServiceApiDto**](ControllerServiceApiDTO.md) | If this type represents a ControllerService, this lists the APIs it implements. | [optional] [default to null]
**Description** | **string** | The description of the type. | [optional] [default to null]
**Restricted** | **bool** | Whether this type is restricted. | [optional] [default to null]
**UsageRestriction** | **string** | The optional description of why the usage of this component is restricted. | [optional] [default to null]
**ExplicitRestrictions** | [**[]ExplicitRestrictionDto**](ExplicitRestrictionDTO.md) | An optional collection of explicit restrictions. If specified, these explicit restrictions will be enfored. | [optional] [default to null]
**DeprecationReason** | **string** | The description of why the usage of this component is restricted. | [optional] [default to null]
**Tags** | **[]string** | The tags associated with this type. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

