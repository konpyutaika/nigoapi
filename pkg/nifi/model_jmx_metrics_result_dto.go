/*
 * NiFi Rest API
 *
 * The Rest API provides programmatic access to command and control a NiFi instance in real time. Start and                                             stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.23.2
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nifi

type JmxMetricsResultDto struct {
	// The bean name of the metrics bean.
	BeanName string `json:"beanName,omitempty"`
	// The attribute name of the metrics bean's attribute.
	AttributeName string `json:"attributeName,omitempty"`
	// The attribute value of the the metrics bean's attribute
	AttributeValue interface{} `json:"attributeValue,omitempty"`
}