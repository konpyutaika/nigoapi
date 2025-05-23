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

type ControllerServicesEntity struct {
	ControllerServices []ControllerServiceEntity `json:"controllerServices,omitempty"`
	// The current time on the system.
	CurrentTime string `json:"currentTime,omitempty"`
}
