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

// Explicit restrictions that indicate a require permission to use the component
type Restriction struct {
	// The explanation of this restriction
	Explanation string `json:"explanation,omitempty"`
	// The permission required for this restriction
	RequiredPermission string `json:"requiredPermission,omitempty"`
}
