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

type BulletinBoardPatternParameter struct {
	Pattern *interface{} `json:"pattern,omitempty"`
	RawPattern string `json:"rawPattern,omitempty"`
}
