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

// The status for this FlowAnalysisRule.
type FlowAnalysisRuleStatusDto struct {
	// The number of active threads for the component.
	ActiveThreadCount int32 `json:"activeThreadCount,omitempty"`
	// The run status of this FlowAnalysisRule
	RunStatus string `json:"runStatus,omitempty"`
	// Indicates whether the component is valid, invalid, or still in the process of validating (i.e., it is unknown whether or not the component is valid)
	ValidationStatus string `json:"validationStatus,omitempty"`
}
