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

type FlowRegistryPermissions struct {
	CanDelete bool `json:"canDelete,omitempty"`
	CanRead bool `json:"canRead,omitempty"`
	CanWrite bool `json:"canWrite,omitempty"`
}
