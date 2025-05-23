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

type ExtensionRepoVersionSummary struct {
	// The artifact id
	ArtifactId string `json:"artifactId,omitempty"`
	// The identity of the user that created this version
	Author string `json:"author,omitempty"`
	// The bucket name
	BucketName string `json:"bucketName,omitempty"`
	// The group id
	GroupId string `json:"groupId,omitempty"`
	Link *Link `json:"link,omitempty"`
	// The timestamp of when this version was created
	Timestamp int64 `json:"timestamp,omitempty"`
	// The version
	Version string `json:"version,omitempty"`
}
