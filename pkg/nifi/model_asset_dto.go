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

// The Asset.
type AssetDto struct {
	// The digest of the asset, will be null if the asset content is missing.
	Digest string `json:"digest,omitempty"`
	// The identifier of the asset.
	Id string `json:"id,omitempty"`
	// Indicates if the content of the asset is missing.
	MissingContent bool `json:"missingContent,omitempty"`
	// The name of the asset.
	Name string `json:"name,omitempty"`
}
