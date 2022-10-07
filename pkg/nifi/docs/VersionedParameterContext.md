# VersionedParameterContext

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The component&#39;s unique identifier | [optional] [default to null]
**InstanceIdentifier** | **string** | The instance ID of an existing component that is described by this VersionedComponent, or null if this is not mapped to an instantiated component | [optional] [default to null]
**Name** | **string** | The component&#39;s name | [optional] [default to null]
**Comments** | **string** | The user-supplied comments for the component | [optional] [default to null]
**Position** | [***Position**](Position.md) | The component&#39;s position on the graph | [optional] [default to null]
**Parameters** | [**[]VersionedParameter**](VersionedParameter.md) | The parameters in the context | [optional] [default to null]
**InheritedParameterContexts** | **[]string** | The names of additional parameter contexts from which to inherit parameters | [optional] [default to null]
**Description** | **string** | The description of the parameter context | [optional] [default to null]
**ParameterProvider** | **string** | The identifier of an optional parameter provider | [optional] [default to null]
**ParameterGroupName** | **string** | The corresponding parameter group name fetched from the parameter provider, if applicable | [optional] [default to null]
**Synchronized** | **bool** | True if the parameter provider is set and the context should receive updates when its parameters are next fetched | [optional] [default to null]
**ComponentType** | **string** |  | [optional] [default to null]
**GroupIdentifier** | **string** | The ID of the Process Group that this component belongs to | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


