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

type Extension struct {
	DefaultSchedule *DefaultSchedule `json:"defaultSchedule,omitempty"`
	DefaultSettings *DefaultSettings `json:"defaultSettings,omitempty"`
	DeprecationNotice *DeprecationNotice `json:"deprecationNotice,omitempty"`
	// The description of the extension
	Description string `json:"description,omitempty"`
	// The dynamic properties of the extension
	DynamicProperties []DynamicProperty `json:"dynamicProperties,omitempty"`
	DynamicRelationship *DynamicRelationship `json:"dynamicRelationship,omitempty"`
	// The input requirement of the extension
	InputRequirement string `json:"inputRequirement,omitempty"`
	// Zero or more documented use cases for how the processor may be used in conjunction with other processors
	MultiProcessorUseCases []MultiProcessorUseCase `json:"multiProcessorUseCases,omitempty"`
	// The name of the extension
	Name string `json:"name"`
	// Indicates that a processor should be scheduled only on the primary node
	PrimaryNodeOnly bool `json:"primaryNodeOnly,omitempty"`
	// The properties of the extension
	Properties []Property `json:"properties,omitempty"`
	// The service APIs provided by this extension
	ProvidedServiceAPIs []ProvidedServiceApi `json:"providedServiceAPIs,omitempty"`
	// The attributes read from flow files by the extension
	ReadsAttributes []Attribute `json:"readsAttributes,omitempty"`
	// The relationships of the extension
	Relationships []Relationship `json:"relationships,omitempty"`
	Restricted *Restricted `json:"restricted,omitempty"`
	// The names of other extensions to see
	SeeAlso []string `json:"seeAlso,omitempty"`
	// Indicates that a processor is side effect free
	SideEffectFree bool `json:"sideEffectFree,omitempty"`
	Stateful *Stateful `json:"stateful,omitempty"`
	// Indicates that a processor supports batching
	SupportsBatching bool `json:"supportsBatching,omitempty"`
	SupportsSensitiveDynamicProperties bool `json:"supportsSensitiveDynamicProperties,omitempty"`
	// The resource considerations of the extension
	SystemResourceConsiderations []SystemResourceConsideration `json:"systemResourceConsiderations,omitempty"`
	// The tags of the extension
	Tags []string `json:"tags,omitempty"`
	// Indicates that a processor should be triggered serially
	TriggerSerially bool `json:"triggerSerially,omitempty"`
	// Indicates that a processor should be triggered when any destinations have space for flow files
	TriggerWhenAnyDestinationAvailable bool `json:"triggerWhenAnyDestinationAvailable,omitempty"`
	// Indicates that a processor should be triggered when the incoming queues are empty
	TriggerWhenEmpty bool `json:"triggerWhenEmpty,omitempty"`
	// The type of the extension
	Type_ string `json:"type"`
	// Zero or more documented use cases for how the extension may be used
	UseCases []UseCase `json:"useCases,omitempty"`
	// The attributes written to flow files by the extension
	WritesAttributes []Attribute `json:"writesAttributes,omitempty"`
}
