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

type PortStatusSnapshotDto struct {
	// The active thread count for the port.
	ActiveThreadCount int32 `json:"activeThreadCount,omitempty"`
	// The size of hte FlowFiles that have been accepted in the last 5 minutes.
	BytesIn int64 `json:"bytesIn,omitempty"`
	// The number of bytes that have been processed in the last 5 minutes.
	BytesOut int64 `json:"bytesOut,omitempty"`
	// The number of FlowFiles that have been accepted in the last 5 minutes.
	FlowFilesIn int32 `json:"flowFilesIn,omitempty"`
	// The number of FlowFiles that have been processed in the last 5 minutes.
	FlowFilesOut int32 `json:"flowFilesOut,omitempty"`
	// The id of the parent process group of the port.
	GroupId string `json:"groupId,omitempty"`
	// The id of the port.
	Id string `json:"id,omitempty"`
	// The count/size of flowfiles that have been accepted in the last 5 minutes.
	Input string `json:"input,omitempty"`
	// The name of the port.
	Name string `json:"name,omitempty"`
	// The count/size of flowfiles that have been processed in the last 5 minutes.
	Output string `json:"output,omitempty"`
	// The run status of the port.
	RunStatus string `json:"runStatus,omitempty"`
	// Whether the port has incoming or outgoing connections to a remote NiFi.
	Transmitting bool `json:"transmitting,omitempty"`
}
