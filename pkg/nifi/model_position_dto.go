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

// The position of this component in the UI if applicable.
type PositionDto struct {
	// The x coordinate.
	X float64 `json:"x"`
	// The y coordinate.
	Y float64 `json:"y"`
}
