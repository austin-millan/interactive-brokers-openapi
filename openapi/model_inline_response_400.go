/*
 * Client Portal Web API
 *
 * Production version of the Client Portal Web API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineResponse400 struct for InlineResponse400
type InlineResponse400 struct {
	// for example-order not confirmed
	Error string `json:"error,omitempty"`
	StatusCode int32 `json:"statusCode,omitempty"`
}
