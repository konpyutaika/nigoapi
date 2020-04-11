# ProcessGroupFlowDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the component. | [optional] [default to null]
**Uri** | **string** | The URI for futures requests to the component. | [optional] [default to null]
**ParentGroupId** | **string** | The id of parent process group of this component if applicable. | [optional] [default to null]
**ParameterContext** | [***ParameterContextReferenceEntity**](ParameterContextReferenceEntity.md) | The Parameter Context, or null if no Parameter Context has been bound to the Process Group | [optional] [default to null]
**Breadcrumb** | [***FlowBreadcrumbEntity**](FlowBreadcrumbEntity.md) | The breadcrumb of the process group. | [optional] [default to null]
**Flow** | [***FlowDto**](FlowDTO.md) | The flow structure starting at this Process Group. | [optional] [default to null]
**LastRefreshed** | **string** | The time the flow for the process group was last refreshed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


