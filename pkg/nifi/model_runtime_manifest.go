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

type RuntimeManifest struct {
	// A unique identifier for the manifest
	Identifier string `json:"identifier,omitempty"`
	// The type of the runtime binary, e.g., 'minifi-java' or 'minifi-cpp'
	AgentType string `json:"agentType,omitempty"`
	// The version of the runtime binary, e.g., '1.0.1'
	Version string `json:"version,omitempty"`
	BuildInfo *BuildInfo `json:"buildInfo,omitempty"`
	// All extension bundles included with this runtime
	Bundles []Bundle `json:"bundles,omitempty"`
	SchedulingDefaults *SchedulingDefaults `json:"schedulingDefaults,omitempty"`
}
