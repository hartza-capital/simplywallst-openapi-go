/*
 * simply-wallst
 *
 * simply-wallst API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// GridsMeta struct for GridsMeta
type GridsMeta struct {
	TotalRecords     int64  `json:"total_records,omitempty"`
	RealTotalRecords int64  `json:"real_total_records,omitempty"`
	State            string `json:"state,omitempty"`
	NoResultIfLimit  bool   `json:"noResultIfLimit,omitempty"`
}
