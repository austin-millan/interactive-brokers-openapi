/*
 * Client Portal Web API
 *
 * Production version of the Client Portal Web API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PerformanceTpps Time period performance data
type PerformanceTpps struct {
	// array of dates, the length should be same as the length of returns inside data.
	Dates []string `json:"dates,omitempty"`
	// M means Month
	Freq string `json:"freq,omitempty"`
	Data []PerformanceCpsData `json:"data,omitempty"`
}
