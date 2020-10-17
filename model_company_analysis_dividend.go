/*
 * simply-wallst
 *
 * simply-wallst API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CompanyAnalysisDividend struct for CompanyAnalysisDividend
type CompanyAnalysisDividend struct {
	Current  float32 `json:"current,omitempty"`
	Future   float32 `json:"future,omitempty"`
	Upcoming bool    `json:"upcoming,omitempty"`
	ExDate   int32   `json:"ex_date,omitempty"`
}