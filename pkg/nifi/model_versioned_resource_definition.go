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

// Returns the Resource Definition that defines which type(s) of resource(s) this property references, if any
type VersionedResourceDefinition struct {
	// The cardinality of the resource
	Cardinality string `json:"cardinality,omitempty"`
	// The types of resource that the Property Descriptor is allowed to reference
	ResourceTypes []string `json:"resourceTypes,omitempty"`
}
