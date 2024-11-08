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

type ConnectionStatisticsDto struct {
	// The ID of the connection
	Id string `json:"id,omitempty"`
	// The timestamp of when the stats were last refreshed
	StatsLastRefreshed string `json:"statsLastRefreshed,omitempty"`
	AggregateSnapshot *ConnectionStatisticsSnapshotDto `json:"aggregateSnapshot,omitempty"`
	// A list of status snapshots for each node
	NodeSnapshots []NodeConnectionStatisticsSnapshotDto `json:"nodeSnapshots,omitempty"`
}
