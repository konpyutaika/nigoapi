# NarSummaryDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildTime** | **string** | The time the NAR was built according to it&#x27;s MANIFEST | [optional] [default to null]
**Coordinate** | [***NarCoordinateDto**](NarCoordinateDTO.md) |  | [optional] [default to null]
**CreatedBy** | **string** | The plugin that created the NAR according to it&#x27;s MANIFEST | [optional] [default to null]
**DependencyCoordinate** | [***NarCoordinateDto**](NarCoordinateDTO.md) |  | [optional] [default to null]
**Digest** | **string** | The hex digest of the NAR contents | [optional] [default to null]
**ExtensionCount** | **int32** | The number of extensions contained in this NAR | [optional] [default to null]
**FailureMessage** | **string** | Information about why the installation failed, only populated when the state is failed | [optional] [default to null]
**Identifier** | **string** | The identifier of the NAR. | [optional] [default to null]
**InstallComplete** | **bool** | Indicates if the install task has completed | [optional] [default to null]
**SourceIdentifier** | **string** | The identifier of the source of this NAR | [optional] [default to null]
**SourceType** | **string** | The source of this NAR | [optional] [default to null]
**State** | **string** | The state of the NAR (i.e. Installed, or not) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

