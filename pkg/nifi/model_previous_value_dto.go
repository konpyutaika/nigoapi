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

// Previous values for a given property.
type PreviousValueDto struct {
	// The previous value.
	PreviousValue string `json:"previousValue,omitempty"`
	// The timestamp when the value was modified.
	Timestamp string `json:"timestamp,omitempty"`
	// The user who changed the previous value.
	UserIdentity string `json:"userIdentity,omitempty"`
}
