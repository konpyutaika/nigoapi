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

type AboutDto struct {
	// The title to be used on the page and in the about dialog.
	Title string `json:"title,omitempty"`
	// The version of this NiFi.
	Version string `json:"version,omitempty"`
	// The URI for the NiFi.
	Uri string `json:"uri,omitempty"`
	// The URL for the content viewer if configured.
	ContentViewerUrl string `json:"contentViewerUrl,omitempty"`
	// The timezone of the NiFi instance.
	Timezone string `json:"timezone,omitempty"`
	// Build tag
	BuildTag string `json:"buildTag,omitempty"`
	// Build revision or commit hash
	BuildRevision string `json:"buildRevision,omitempty"`
	// Build branch
	BuildBranch string `json:"buildBranch,omitempty"`
	// Build timestamp
	BuildTimestamp string `json:"buildTimestamp,omitempty"`
}
