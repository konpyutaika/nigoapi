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

type ExtensionRepoBucket struct {
	// The name of the bucket
	BucketName string `json:"bucketName,omitempty"`
	Link *Link `json:"link,omitempty"`
}
