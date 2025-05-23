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

type ParameterProviderReferencingComponentDto struct {
	// The id of the component referencing a parameter provider.
	Id string `json:"id,omitempty"`
	// The name of the component referencing a parameter provider.
	Name string `json:"name,omitempty"`
}
