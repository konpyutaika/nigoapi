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

type ClusterSummaryDto struct {
	// Whether this NiFi instance is clustered.
	Clustered bool `json:"clustered,omitempty"`
	// The number of nodes that are currently connected to the cluster
	ConnectedNodeCount int32 `json:"connectedNodeCount,omitempty"`
	// When clustered, reports the number of nodes connected vs the number of nodes in the cluster.
	ConnectedNodes string `json:"connectedNodes,omitempty"`
	// Whether this NiFi instance is connected to a cluster.
	ConnectedToCluster bool `json:"connectedToCluster,omitempty"`
	// The number of nodes in the cluster, regardless of whether or not they are connected
	TotalNodeCount int32 `json:"totalNodeCount,omitempty"`
}
