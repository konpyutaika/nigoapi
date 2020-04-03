# LabelDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] [default to null]
**VersionedComponentId** | **string** | The ID of the corresponding component that is under version control | [optional] [default to null]
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] [default to null]
**Position** | [***PositionDto**](PositionDTO.md) | The position of this component in the UI if applicable. | [optional] [default to null]
**Label** | **string** | The text that appears in the label. | [optional] [default to null]
**Width** | **float64** | The width of the label in pixels when at a 1:1 scale. | [optional] [default to null]
**Height** | **float64** | The height of the label in pixels when at a 1:1 scale. | [optional] [default to null]
**Style** | **map[string]string** | The styles for this label (font-size : 12px, background-color : #eee, etc). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


