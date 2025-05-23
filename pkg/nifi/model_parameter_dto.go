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

// The parameter information
type ParameterDto struct {
	// The description of the Parameter
	Description *string `json:"description,omitempty"`
	// Whether or not the Parameter is inherited from another context
	Inherited bool `json:"inherited,omitempty"`
	// The name of the Parameter
	Name             string                           `json:"name,omitempty"`
	ParameterContext *ParameterContextReferenceEntity `json:"parameterContext,omitempty"`
	// Whether or not the Parameter is provided by a ParameterProvider
	Provided bool `json:"provided,omitempty"`
	// A list of identifiers of the assets that are referenced by the parameter
	ReferencedAssets []AssetReferenceDto `json:"referencedAssets,omitempty"`
	// The set of all components in the flow that are referencing this Parameter
	ReferencingComponents []AffectedComponentEntity `json:"referencingComponents,omitempty"`
	// Whether or not the Parameter is sensitive
	Sensitive bool `json:"sensitive,omitempty"`
	// The value of the Parameter
	Value *string `json:"value"`
	// Whether or not the value of the Parameter was removed. When a request is made to change a parameter, the value may be null. The absence of the value may be used either to indicate that the value is not to be changed, or that the value is to be set to null (i.e., removed). This denotes which of the two scenarios is being encountered.
	ValueRemoved bool `json:"valueRemoved,omitempty"`
}
