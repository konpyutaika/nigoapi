/*
 * Apache NiFi REST API
 *
 * REST API definition for Apache NiFi web services
 *
 * API version: 2.1.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nifi

// Represents the processor's processing performance.
type ProcessingPerformanceStatusDto struct {
	// The number of nanoseconds has spent to read content in the last 5 minutes.
	ContentReadDuration int64 `json:"contentReadDuration,omitempty"`
	// The number of nanoseconds has spent to write content in the last 5 minutes.
	ContentWriteDuration int64 `json:"contentWriteDuration,omitempty"`
	// The number of nanoseconds has spent on CPU usage in the last 5 minutes.
	CpuDuration int64 `json:"cpuDuration,omitempty"`
	// The number of nanoseconds has spent running garbage collection in the last 5 minutes.
	GarbageCollectionDuration int64 `json:"garbageCollectionDuration,omitempty"`
	// The unique ID of the process group that the Processor belongs to
	Identifier string `json:"identifier,omitempty"`
	// The number of nanoseconds has spent running to commit sessions the last 5 minutes.
	SessionCommitDuration int64 `json:"sessionCommitDuration,omitempty"`
}