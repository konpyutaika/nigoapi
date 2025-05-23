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
import (
	"time"
)

type FormDataContentDisposition struct {
	CreationDate time.Time `json:"creationDate,omitempty"`
	FileName string `json:"fileName,omitempty"`
	ModificationDate time.Time `json:"modificationDate,omitempty"`
	Name string `json:"name,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
	ReadDate time.Time `json:"readDate,omitempty"`
	Size int64 `json:"size,omitempty"`
	Type_ string `json:"type,omitempty"`
}
