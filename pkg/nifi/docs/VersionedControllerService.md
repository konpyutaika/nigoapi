# VersionedControllerService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The component&#39;s unique identifier | [optional] [default to null]
**Name** | **string** | The component&#39;s name | [optional] [default to null]
**Comments** | **string** | The user-supplied comments for the component | [optional] [default to null]
**Position** | [***Position**](Position.md) | The component&#39;s position on the graph | [optional] [default to null]
**Type_** | **string** | The type of the controller service. | [optional] [default to null]
**Bundle** | [***Bundle**](Bundle.md) | The details of the artifact that bundled this processor type. | [optional] [default to null]
**ControllerServiceApis** | [**[]ControllerServiceApi**](ControllerServiceAPI.md) | Lists the APIs this Controller Service implements. | [optional] [default to null]
**Properties** | **map[string]string** | The properties of the controller service. | [optional] [default to null]
**PropertyDescriptors** | [**map[string]VersionedPropertyDescriptor**](VersionedPropertyDescriptor.md) | The property descriptors for the processor. | [optional] [default to null]
**AnnotationData** | **string** | The annotation for the controller service. This is how the custom UI relays configuration to the controller service. | [optional] [default to null]
**ComponentType** | **string** |  | [optional] [default to null]
**GroupIdentifier** | **string** | The ID of the Process Group that this component belongs to | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


