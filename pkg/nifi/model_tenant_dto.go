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

type TenantDto struct {
	// Whether this tenant is configurable.
	Configurable bool `json:"configurable,omitempty"`
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The identity of the tenant.
	Identity string `json:"identity,omitempty"`
	// The id of parent process group of this component if applicable.
	ParentGroupId string `json:"parentGroupId,omitempty"`
	Position *PositionDto `json:"position,omitempty"`
	// The ID of the corresponding component that is under version control
	VersionedComponentId string `json:"versionedComponentId,omitempty"`
}
