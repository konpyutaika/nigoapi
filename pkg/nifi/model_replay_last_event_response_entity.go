/*
 * NiFi Rest API
 *
 * The Rest API provides programmatic access to command and control a NiFi instance in real time. Start and                                             stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.23.2
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nifi

type ReplayLastEventResponseEntity struct {
	// The UUID of the component whose last event should be replayed.
	ComponentId string `json:"componentId,omitempty"`
	// Which nodes were requested to replay their last provenance event.
	Nodes string `json:"nodes,omitempty"`
	// The aggregate result of all nodes' responses
	AggregateSnapshot *ReplayLastEventSnapshotDto `json:"aggregateSnapshot,omitempty"`
	// The node-wise results
	NodeSnapshots []NodeReplayLastEventSnapshotDto `json:"nodeSnapshots,omitempty"`
}