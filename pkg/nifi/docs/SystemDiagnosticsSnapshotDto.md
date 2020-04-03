# SystemDiagnosticsSnapshotDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNonHeap** | **string** | Total size of non heap. | [optional] [default to null]
**TotalNonHeapBytes** | **int64** | Total number of bytes allocated to the JVM not used for heap | [optional] [default to null]
**UsedNonHeap** | **string** | Amount of use non heap. | [optional] [default to null]
**UsedNonHeapBytes** | **int64** | Total number of bytes used by the JVM not in the heap space | [optional] [default to null]
**FreeNonHeap** | **string** | Amount of free non heap. | [optional] [default to null]
**FreeNonHeapBytes** | **int64** | Total number of free non-heap bytes available to the JVM | [optional] [default to null]
**MaxNonHeap** | **string** | Maximum size of non heap. | [optional] [default to null]
**MaxNonHeapBytes** | **int64** | The maximum number of bytes that the JVM can use for non-heap purposes | [optional] [default to null]
**NonHeapUtilization** | **string** | Utilization of non heap. | [optional] [default to null]
**TotalHeap** | **string** | Total size of heap. | [optional] [default to null]
**TotalHeapBytes** | **int64** | The total number of bytes that are available for the JVM heap to use | [optional] [default to null]
**UsedHeap** | **string** | Amount of used heap. | [optional] [default to null]
**UsedHeapBytes** | **int64** | The number of bytes of JVM heap that are currently being used | [optional] [default to null]
**FreeHeap** | **string** | Amount of free heap. | [optional] [default to null]
**FreeHeapBytes** | **int64** | The number of bytes that are allocated to the JVM heap but not currently being used | [optional] [default to null]
**MaxHeap** | **string** | Maximum size of heap. | [optional] [default to null]
**MaxHeapBytes** | **int64** | The maximum number of bytes that can be used by the JVM | [optional] [default to null]
**HeapUtilization** | **string** | Utilization of heap. | [optional] [default to null]
**AvailableProcessors** | **int32** | Number of available processors if supported by the underlying system. | [optional] [default to null]
**ProcessorLoadAverage** | **float64** | The processor load average if supported by the underlying system. | [optional] [default to null]
**TotalThreads** | **int32** | Total number of threads. | [optional] [default to null]
**DaemonThreads** | **int32** | Number of daemon threads. | [optional] [default to null]
**Uptime** | **string** | The uptime of the Java virtual machine | [optional] [default to null]
**FlowFileRepositoryStorageUsage** | [***StorageUsageDto**](StorageUsageDTO.md) | The flowfile repository storage usage. | [optional] [default to null]
**ContentRepositoryStorageUsage** | [**[]StorageUsageDto**](StorageUsageDTO.md) | The content repository storage usage. | [optional] [default to null]
**ProvenanceRepositoryStorageUsage** | [**[]StorageUsageDto**](StorageUsageDTO.md) | The provenance repository storage usage. | [optional] [default to null]
**GarbageCollection** | [**[]GarbageCollectionDto**](GarbageCollectionDTO.md) | The garbage collection details. | [optional] [default to null]
**StatsLastRefreshed** | **string** | When the diagnostics were generated. | [optional] [default to null]
**VersionInfo** | [***VersionInfoDto**](VersionInfoDTO.md) | The nifi, os, java, and build version information | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


