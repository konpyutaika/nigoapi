# ProcessorDiagnosticsDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Processor** | [***ProcessorDto**](ProcessorDTO.md) | Information about the Processor for which the Diagnostic Report is generated | [optional] [default to null]
**ProcessorStatus** | [***ProcessorStatusDto**](ProcessorStatusDTO.md) | The Status for the Processor for which the Diagnostic Report is generated | [optional] [default to null]
**ReferencedControllerServices** | [**[]ControllerServiceDiagnosticsDto**](ControllerServiceDiagnosticsDTO.md) | Diagnostic Information about all Controller Services that the Processor is referencing | [optional] [default to null]
**IncomingConnections** | [**[]ConnectionDiagnosticsDto**](ConnectionDiagnosticsDTO.md) | Diagnostic Information about all incoming Connections | [optional] [default to null]
**OutgoingConnections** | [**[]ConnectionDiagnosticsDto**](ConnectionDiagnosticsDTO.md) | Diagnostic Information about all outgoing Connections | [optional] [default to null]
**JvmDiagnostics** | [***JvmDiagnosticsDto**](JVMDiagnosticsDTO.md) | Diagnostic Information about the JVM and system-level diagnostics | [optional] [default to null]
**ThreadDumps** | [**[]ThreadDumpDto**](ThreadDumpDTO.md) | Thread Dumps that were taken of the threads that are active in the Processor | [optional] [default to null]
**ClassLoaderDiagnostics** | [***ClassLoaderDiagnosticsDto**](ClassLoaderDiagnosticsDTO.md) | Information about the Controller Service&#39;s Class Loader | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


