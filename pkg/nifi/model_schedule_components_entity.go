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

type ScheduleComponentsEntity struct {
	// Optional components to schedule. If not specified, all authorized descendant components will be used.
	Components map[string]RevisionDto `json:"components,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`
	// The id of the ProcessGroup
	Id string `json:"id,omitempty"`
	// The desired state of the descendant components
	State string `json:"state,omitempty"`
}
