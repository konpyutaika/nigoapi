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

type RunStatusDetailsRequestEntity struct {
	// The IDs of all processors whose run status details should be provided
	ProcessorIds []string `json:"processorIds,omitempty"`
}
