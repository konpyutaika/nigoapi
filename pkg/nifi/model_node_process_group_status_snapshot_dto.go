/*
 * Apache NiFi REST API
 *
 * REST API definition for Apache NiFi web services
 *
 * API version: 2.4.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nifi

// The status reported by each node in the cluster. If the NiFi instance is a standalone instance, rather than a clustered instance, this value may be null.
type NodeProcessGroupStatusSnapshotDto struct {
	// The API address of the node
	Address string `json:"address,omitempty"`
	// The API port used to communicate with the node
	ApiPort int32 `json:"apiPort,omitempty"`
	// The unique ID that identifies the node
	NodeId string `json:"nodeId,omitempty"`
	StatusSnapshot *ProcessGroupStatusSnapshotDto `json:"statusSnapshot,omitempty"`
}
