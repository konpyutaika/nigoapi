# ParameterGroupConfigurationEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | **string** | The name of the external parameter group to which the provided parameter names apply. | [optional] [default to null]
**ParameterContextName** | **string** | The name of the ParameterContext that receives the parameters in this group | [optional] [default to null]
**ParameterSensitivities** | **map[string]string** | All fetched parameter names that should be applied. | [optional] [default to null]
**Synchronized** | **bool** | True if this group should be synchronized to a ParameterContext, including creating one if it does not exist. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

