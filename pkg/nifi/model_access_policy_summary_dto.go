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

type AccessPolicySummaryDto struct {
	// The action associated with this access policy.
	Action string `json:"action,omitempty"`
	ComponentReference *ComponentReferenceEntity `json:"componentReference,omitempty"`
	// Whether this policy is configurable.
	Configurable bool `json:"configurable,omitempty"`
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The id of parent process group of this component if applicable.
	ParentGroupId string `json:"parentGroupId,omitempty"`
	Position *PositionDto `json:"position,omitempty"`
	// The resource for this access policy.
	Resource string `json:"resource,omitempty"`
	// The ID of the corresponding component that is under version control
	VersionedComponentId string `json:"versionedComponentId,omitempty"`
}
