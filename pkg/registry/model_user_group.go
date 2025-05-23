/*
 * Apache NiFi Registry REST API
 *
 * REST API definition for Apache NiFi Registry web services
 *
 * API version: 2.4.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package registry

type UserGroup struct {
	// The access policies granted to this tenant.
	AccessPolicies []AccessPolicySummary `json:"accessPolicies,omitempty"`
	// Indicates if this tenant is configurable, based on which UserGroupProvider has been configured to manage it.
	Configurable bool `json:"configurable,omitempty"`
	// The computer-generated identifier of the tenant.
	Identifier string `json:"identifier,omitempty"`
	// The human-facing identity of the tenant. This can only be changed if the tenant is configurable.
	Identity string `json:"identity,omitempty"`
	ResourcePermissions *ResourcePermissions `json:"resourcePermissions,omitempty"`
	Revision *RevisionInfo `json:"revision,omitempty"`
	// The users that belong to this user group. This can only be changed if this group is configurable.
	Users []Tenant `json:"users,omitempty"`
}
