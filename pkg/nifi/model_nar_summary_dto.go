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

// The NAR summary
type NarSummaryDto struct {
	// The identifier of the NAR.
	Identifier string `json:"identifier,omitempty"`
	Coordinate *NarCoordinateDto `json:"coordinate,omitempty"`
	DependencyCoordinate *NarCoordinateDto `json:"dependencyCoordinate,omitempty"`
	// The time the NAR was built according to it's MANIFEST
	BuildTime string `json:"buildTime,omitempty"`
	// The plugin that created the NAR according to it's MANIFEST
	CreatedBy string `json:"createdBy,omitempty"`
	// The hex digest of the NAR contents
	Digest string `json:"digest,omitempty"`
	// The source of this NAR
	SourceType string `json:"sourceType,omitempty"`
	// The identifier of the source of this NAR
	SourceIdentifier string `json:"sourceIdentifier,omitempty"`
	// The number of extensions contained in this NAR
	ExtensionCount int32 `json:"extensionCount,omitempty"`
	// The state of the NAR (i.e. Installed, or not)
	State string `json:"state,omitempty"`
	// Information about why the installation failed, only populated when the state is failed
	FailureMessage string `json:"failureMessage,omitempty"`
	// Indicates if the install task has completed
	InstallComplete bool `json:"installComplete,omitempty"`
}
