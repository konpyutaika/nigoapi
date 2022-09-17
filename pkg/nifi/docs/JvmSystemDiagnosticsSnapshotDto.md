# JvmSystemDiagnosticsSnapshotDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowFileRepositoryStorageUsage** | [***RepositoryUsageDto**](RepositoryUsageDTO.md) | Information about the FlowFile Repository&#39;s usage | [optional] [default to null]
**ContentRepositoryStorageUsage** | [**[]RepositoryUsageDto**](RepositoryUsageDTO.md) | Information about the Content Repository&#39;s usage | [optional] [default to null]
**ProvenanceRepositoryStorageUsage** | [**[]RepositoryUsageDto**](RepositoryUsageDTO.md) | Information about the Provenance Repository&#39;s usage | [optional] [default to null]
**MaxHeapBytes** | **int64** | The maximum number of bytes that the JVM heap is configured to use for heap | [optional] [default to null]
**MaxHeap** | **string** | The maximum number of bytes that the JVM heap is configured to use, as a human-readable value | [optional] [default to null]
**GarbageCollectionDiagnostics** | [**[]GarbageCollectionDiagnosticsDto**](GarbageCollectionDiagnosticsDTO.md) | Diagnostic information about the JVM&#39;s garbage collections | [optional] [default to null]
**CpuCores** | **int32** | The number of CPU Cores available on the system | [optional] [default to null]
**CpuLoadAverage** | **float64** | The 1-minute CPU Load Average | [optional] [default to null]
**PhysicalMemoryBytes** | **int64** | The number of bytes of RAM available on the system | [optional] [default to null]
**PhysicalMemory** | **string** | The number of bytes of RAM available on the system as a human-readable value | [optional] [default to null]
**OpenFileDescriptors** | **int64** | The number of files that are open by the NiFi process | [optional] [default to null]
**MaxOpenFileDescriptors** | **int64** | The maximum number of open file descriptors that are available to each process | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


