/*
 * NiFi Rest API
 *
 * The Rest API provides programmatic access to command and control a NiFi instance in real time. Start and                                             stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.15.3
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nifi

type JvmFlowDiagnosticsSnapshotDto struct {
	// How long this node has been running, formatted as hours:minutes:seconds.milliseconds
	Uptime string `json:"uptime,omitempty"`
	// The name of the Time Zone that is configured, if available
	TimeZone string `json:"timeZone,omitempty"`
	// The number of timer-driven threads that are active
	ActiveTimerDrivenThreads int32 `json:"activeTimerDrivenThreads,omitempty"`
	// The number of event-driven threads that are active
	ActiveEventDrivenThreads int32 `json:"activeEventDrivenThreads,omitempty"`
	// The NiFi Bundles (NARs) that are loaded by NiFi
	BundlesLoaded []BundleDto `json:"bundlesLoaded,omitempty"`
}