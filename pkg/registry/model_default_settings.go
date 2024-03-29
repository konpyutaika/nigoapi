/*
 * Apache NiFi Registry REST API
 *
 * The REST API provides an interface to a registry with operations for saving, versioning, reading NiFi flows and components.
 *
 * API version: 1.24.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package registry

type DefaultSettings struct {
	// The default yield duration
	YieldDuration string `json:"yieldDuration,omitempty"`
	// The default penalty duration
	PenaltyDuration string `json:"penaltyDuration,omitempty"`
	// The default bulletin level
	BulletinLevel string `json:"bulletinLevel,omitempty"`
}
