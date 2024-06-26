/*
 * Apache NiFi Registry REST API
 *
 * REST API definition for Apache NiFi Registry web services
 *
 * API version: 2.0.0-M2
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package registry

// An WebLink to this entity.
type Link struct {
	Uri string `json:"uri,omitempty"`
	Title string `json:"title,omitempty"`
	UriBuilder *UriBuilder `json:"uriBuilder,omitempty"`
	Rel string `json:"rel,omitempty"`
	Rels []string `json:"rels,omitempty"`
	Type_ string `json:"type,omitempty"`
	Params map[string]string `json:"params,omitempty"`
}
