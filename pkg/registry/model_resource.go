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

type Resource struct {
	// The identifier of the resource.
	Identifier string `json:"identifier,omitempty"`
	// The name of the resource.
	Name string `json:"name,omitempty"`
}
