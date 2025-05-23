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

// A summary top-level resource access policies granted to this tenant.
type ResourcePermissions struct {
	AnyTopLevelResource *Permissions `json:"anyTopLevelResource,omitempty"`
	Buckets *Permissions `json:"buckets,omitempty"`
	Policies *Permissions `json:"policies,omitempty"`
	Proxy *Permissions `json:"proxy,omitempty"`
	Tenants *Permissions `json:"tenants,omitempty"`
}
