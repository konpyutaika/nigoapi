/*
 * Apache NiFi REST API
 *
 * REST API definition for Apache NiFi web services
 *
 * API version: 2.1.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nifi

type NarDetailsEntity struct {
	// The ControllerService types contained in the NAR
	ControllerServiceTypes []DocumentedTypeDto `json:"controllerServiceTypes,omitempty"`
	// The coordinates of NARs that depend on this NAR
	DependentCoordinates []NarCoordinateDto `json:"dependentCoordinates,omitempty"`
	// The FlowAnalysisRule types contained in the NAR
	FlowAnalysisRuleTypes []DocumentedTypeDto `json:"flowAnalysisRuleTypes,omitempty"`
	// The FlowRegistryClient types contained in the NAR
	FlowRegistryClientTypes []DocumentedTypeDto `json:"flowRegistryClientTypes,omitempty"`
	NarSummary *NarSummaryDto `json:"narSummary,omitempty"`
	// The ParameterProvider types contained in the NAR
	ParameterProviderTypes []DocumentedTypeDto `json:"parameterProviderTypes,omitempty"`
	// The Processor types contained in the NAR
	ProcessorTypes []DocumentedTypeDto `json:"processorTypes,omitempty"`
	// The ReportingTask types contained in the NAR
	ReportingTaskTypes []DocumentedTypeDto `json:"reportingTaskTypes,omitempty"`
}