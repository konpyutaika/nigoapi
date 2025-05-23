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

// The information about controller services that exist outside this versioned flow, but are referenced by components within the versioned flow.
type ExternalControllerServiceReference struct {
	// The identifier of the controller service
	Identifier string `json:"identifier,omitempty"`
	// The name of the controller service
	Name string `json:"name,omitempty"`
}
