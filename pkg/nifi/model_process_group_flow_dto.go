/*
 * Apache NiFi REST API
 *
 * REST API definition for Apache NiFi web services
 *
 * API version: 2.0.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nifi

type ProcessGroupFlowDto struct {
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The URI for futures requests to the component.
	Uri string `json:"uri,omitempty"`
	// The id of parent process group of this component if applicable.
	ParentGroupId string `json:"parentGroupId,omitempty"`
	ParameterContext *ParameterContextReferenceEntity `json:"parameterContext,omitempty"`
	Breadcrumb *FlowBreadcrumbEntity `json:"breadcrumb,omitempty"`
	Flow *FlowDto `json:"flow,omitempty"`
	// The time the flow for the process group was last refreshed.
	LastRefreshed string `json:"lastRefreshed,omitempty"`
}
