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

// The properties of the extension
type Property struct {
	// The allowable values for this property
	AllowableValues []AllowableValue `json:"allowableValues,omitempty"`
	ControllerServiceDefinition *ControllerServiceDefinition `json:"controllerServiceDefinition,omitempty"`
	// The default value
	DefaultValue string `json:"defaultValue,omitempty"`
	// The properties that this property depends on
	Dependencies []Dependency `json:"dependencies,omitempty"`
	// The description
	Description string `json:"description,omitempty"`
	// The display name
	DisplayName string `json:"displayName,omitempty"`
	// Whether or not the processor is dynamic
	Dynamic bool `json:"dynamic,omitempty"`
	// Whether or not the processor dynamically modifies the classpath
	DynamicallyModifiesClasspath bool `json:"dynamicallyModifiesClasspath,omitempty"`
	// The scope of expression language support
	ExpressionLanguageScope string `json:"expressionLanguageScope,omitempty"`
	// Whether or not expression language is supported
	ExpressionLanguageSupported bool `json:"expressionLanguageSupported,omitempty"`
	// The name of the property
	Name string `json:"name,omitempty"`
	// Whether or not the property is required
	Required bool `json:"required,omitempty"`
	ResourceDefinition *ResourceDefinition `json:"resourceDefinition,omitempty"`
	// Whether or not the property is sensitive
	Sensitive bool `json:"sensitive,omitempty"`
}
