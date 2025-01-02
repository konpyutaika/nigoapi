# Extension

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSchedule** | [***DefaultSchedule**](DefaultSchedule.md) |  | [optional] [default to null]
**DefaultSettings** | [***DefaultSettings**](DefaultSettings.md) |  | [optional] [default to null]
**DeprecationNotice** | [***DeprecationNotice**](DeprecationNotice.md) |  | [optional] [default to null]
**Description** | **string** | The description of the extension | [optional] [default to null]
**DynamicProperties** | [**[]DynamicProperty**](DynamicProperty.md) | The dynamic properties of the extension | [optional] [default to null]
**DynamicRelationship** | [***DynamicRelationship**](DynamicRelationship.md) |  | [optional] [default to null]
**InputRequirement** | **string** | The input requirement of the extension | [optional] [default to null]
**MultiProcessorUseCases** | [**[]MultiProcessorUseCase**](MultiProcessorUseCase.md) | Zero or more documented use cases for how the processor may be used in conjunction with other processors | [optional] [default to null]
**Name** | **string** | The name of the extension | [default to null]
**PrimaryNodeOnly** | **bool** | Indicates that a processor should be scheduled only on the primary node | [optional] [default to null]
**Properties** | [**[]Property**](Property.md) | The properties of the extension | [optional] [default to null]
**ProvidedServiceAPIs** | [**[]ProvidedServiceApi**](ProvidedServiceAPI.md) | The service APIs provided by this extension | [optional] [default to null]
**ReadsAttributes** | [**[]Attribute**](Attribute.md) | The attributes read from flow files by the extension | [optional] [default to null]
**Relationships** | [**[]Relationship**](Relationship.md) | The relationships of the extension | [optional] [default to null]
**Restricted** | [***Restricted**](Restricted.md) |  | [optional] [default to null]
**SeeAlso** | **[]string** | The names of other extensions to see | [optional] [default to null]
**SideEffectFree** | **bool** | Indicates that a processor is side effect free | [optional] [default to null]
**Stateful** | [***Stateful**](Stateful.md) |  | [optional] [default to null]
**SupportsBatching** | **bool** | Indicates that a processor supports batching | [optional] [default to null]
**SupportsSensitiveDynamicProperties** | **bool** |  | [optional] [default to null]
**SystemResourceConsiderations** | [**[]SystemResourceConsideration**](SystemResourceConsideration.md) | The resource considerations of the extension | [optional] [default to null]
**Tags** | **[]string** | The tags of the extension | [optional] [default to null]
**TriggerSerially** | **bool** | Indicates that a processor should be triggered serially | [optional] [default to null]
**TriggerWhenAnyDestinationAvailable** | **bool** | Indicates that a processor should be triggered when any destinations have space for flow files | [optional] [default to null]
**TriggerWhenEmpty** | **bool** | Indicates that a processor should be triggered when the incoming queues are empty | [optional] [default to null]
**Type_** | **string** | The type of the extension | [default to null]
**UseCases** | [**[]UseCase**](UseCase.md) | Zero or more documented use cases for how the extension may be used | [optional] [default to null]
**WritesAttributes** | [**[]Attribute**](Attribute.md) | The attributes written to flow files by the extension | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

