/*
 * NiFi Rest API
 *
 * The Rest API provides programmatic access to command and control a NiFi instance in real time. Start and                                             stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.23.2
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nifi

type PropertyDependencyDto struct {
	// The name of the property that is being depended upon
	PropertyName string `json:"propertyName,omitempty"`
	// The values for the property that satisfies the dependency, or null if the dependency is satisfied by the presence of any value for the associated property name
	DependentValues []string `json:"dependentValues,omitempty"`
}