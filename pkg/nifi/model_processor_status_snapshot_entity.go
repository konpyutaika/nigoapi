/*
 * Apache NiFi REST API
 *
 * REST API definition for Apache NiFi web services
 *
 * API version: 2.0.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nifi

// The status of all processors in the process group.
type ProcessorStatusSnapshotEntity struct {
	// The id of the processor.
	Id string `json:"id,omitempty"`
	ProcessorStatusSnapshot *ProcessorStatusSnapshotDto `json:"processorStatusSnapshot,omitempty"`
	// Indicates whether the user can read a given resource.
	CanRead bool `json:"canRead,omitempty"`
}
