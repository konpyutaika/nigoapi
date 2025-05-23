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

// The connections in this flow.
type ConnectionEntity struct {
	// The bend points on the connection.
	Bends []PositionDto `json:"bends,omitempty"`
	// The bulletins for this component.
	Bulletins []BulletinEntity `json:"bulletins,omitempty"`
	Component *ConnectionDto   `json:"component,omitempty"`
	// The identifier of the group of the destination of this connection.
	DestinationGroupId string `json:"destinationGroupId,omitempty"`
	// The identifier of the destination of this connection.
	DestinationId string `json:"destinationId,omitempty"`
	// The type of component the destination connectable is.
	DestinationType string `json:"destinationType"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`
	// The z index of the connection.
	GetzIndex int64 `json:"getzIndex,omitempty"`
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The index of the bend point where to place the connection label.
	LabelIndex  int32           `json:"labelIndex"`
	Permissions *PermissionsDto `json:"permissions,omitempty"`
	Position    *PositionDto    `json:"position,omitempty"`
	Revision    *RevisionDto    `json:"revision,omitempty"`
	// The identifier of the group of the source of this connection.
	SourceGroupId string `json:"sourceGroupId,omitempty"`
	// The identifier of the source of this connection.
	SourceId string `json:"sourceId,omitempty"`
	// The type of component the source connectable is.
	SourceType string               `json:"sourceType"`
	Status     *ConnectionStatusDto `json:"status,omitempty"`
	// The URI for futures requests to the component.
	Uri string `json:"uri,omitempty"`
}
