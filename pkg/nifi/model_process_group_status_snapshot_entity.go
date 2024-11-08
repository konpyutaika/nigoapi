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

// The status of all process groups in the process group.
type ProcessGroupStatusSnapshotEntity struct {
	// The id of the process group.
	Id string `json:"id,omitempty"`
	ProcessGroupStatusSnapshot *ProcessGroupStatusSnapshotDto `json:"processGroupStatusSnapshot,omitempty"`
	// Indicates whether the user can read a given resource.
	CanRead bool `json:"canRead,omitempty"`
}
