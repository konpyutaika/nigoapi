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

// The Flow Update Request
type VersionedFlowUpdateRequestDto struct {
	// Whether or not this request has completed
	Complete bool `json:"complete,omitempty"`
	// An explanation of why this request failed, or null if this request has not failed
	FailureReason string `json:"failureReason,omitempty"`
	// The last time this request was updated.
	LastUpdated string `json:"lastUpdated,omitempty"`
	// The percentage complete for the request, between 0 and 100
	PercentCompleted int32 `json:"percentCompleted,omitempty"`
	// The unique ID of the Process Group being updated
	ProcessGroupId string `json:"processGroupId,omitempty"`
	// The unique ID of this request.
	RequestId string `json:"requestId,omitempty"`
	// The state of the request
	State string `json:"state,omitempty"`
	// The URI for future requests to this drop request.
	Uri string `json:"uri,omitempty"`
	VersionControlInformation *VersionControlInformationDto `json:"versionControlInformation,omitempty"`
}
