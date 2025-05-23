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

// The Parameter Context that is being operated on. This may not be populated until the request has successfully completed.
type ParameterContextDto struct {
	// The Process Groups that are bound to this Parameter Context
	BoundProcessGroups []ProcessGroupEntity `json:"boundProcessGroups"`
	// The Description of the Parameter Context.
	Description string `json:"description,omitempty"`
	// The ID the Parameter Context.
	Id string `json:"id,omitempty"`
	// A list of references of Parameter Contexts from which this one inherits parameters
	InheritedParameterContexts []ParameterContextReferenceEntity `json:"inheritedParameterContexts"`
	// The Name of the Parameter Context.
	Name                           string                                `json:"name,omitempty"`
	ParameterProviderConfiguration *ParameterProviderConfigurationEntity `json:"parameterProviderConfiguration,omitempty"`
	// The Parameters for the Parameter Context
	Parameters []ParameterEntity `json:"parameters"`
}
