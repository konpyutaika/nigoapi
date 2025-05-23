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

type ProcessGroupEntity struct {
	// The number of active remote ports in the process group.
	ActiveRemotePortCount int32 `json:"activeRemotePortCount,omitempty"`
	// The bulletins for this component.
	Bulletins []BulletinEntity `json:"bulletins,omitempty"`
	Component *ProcessGroupDto `json:"component,omitempty"`
	// The number of disabled components in the process group.
	DisabledCount int32 `json:"disabledCount,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The number of inactive remote ports in the process group.
	InactiveRemotePortCount int32 `json:"inactiveRemotePortCount,omitempty"`
	// The number of input ports in the process group.
	InputPortCount int32 `json:"inputPortCount,omitempty"`
	// The number of invalid components in the process group.
	InvalidCount int32 `json:"invalidCount,omitempty"`
	// The number of local input ports in the process group.
	LocalInputPortCount int32 `json:"localInputPortCount,omitempty"`
	// The number of local output ports in the process group.
	LocalOutputPortCount int32 `json:"localOutputPortCount,omitempty"`
	// The number of locally modified and stale versioned process groups in the process group.
	LocallyModifiedAndStaleCount int32 `json:"locallyModifiedAndStaleCount,omitempty"`
	// The number of locally modified versioned process groups in the process group.
	LocallyModifiedCount int32 `json:"locallyModifiedCount,omitempty"`
	// The number of output ports in the process group.
	OutputPortCount int32 `json:"outputPortCount,omitempty"`
	ParameterContext *ParameterContextReferenceEntity `json:"parameterContext,omitempty"`
	Permissions *PermissionsDto `json:"permissions,omitempty"`
	Position *PositionDto `json:"position,omitempty"`
	// Determines the process group update strategy
	ProcessGroupUpdateStrategy string `json:"processGroupUpdateStrategy,omitempty"`
	// The number of public input ports in the process group.
	PublicInputPortCount int32 `json:"publicInputPortCount,omitempty"`
	// The number of public output ports in the process group.
	PublicOutputPortCount int32 `json:"publicOutputPortCount,omitempty"`
	Revision *RevisionDto `json:"revision,omitempty"`
	// The number of running components in this process group.
	RunningCount int32 `json:"runningCount,omitempty"`
	// The number of stale versioned process groups in the process group.
	StaleCount int32 `json:"staleCount,omitempty"`
	Status *ProcessGroupStatusDto `json:"status,omitempty"`
	// The number of stopped components in the process group.
	StoppedCount int32 `json:"stoppedCount,omitempty"`
	// The number of versioned process groups in the process group that are unable to sync to a registry.
	SyncFailureCount int32 `json:"syncFailureCount,omitempty"`
	// The number of up to date versioned process groups in the process group.
	UpToDateCount int32 `json:"upToDateCount,omitempty"`
	// The URI for futures requests to the component.
	Uri string `json:"uri,omitempty"`
	VersionedFlowSnapshot *RegisteredFlowSnapshot `json:"versionedFlowSnapshot,omitempty"`
	// The current state of the Process Group, as it relates to the Versioned Flow
	VersionedFlowState string `json:"versionedFlowState,omitempty"`
}
