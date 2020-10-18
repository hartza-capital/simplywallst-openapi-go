/*
 * simply-wallst
 *
 * simply-wallst API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// DevelopmentEvent struct for DevelopmentEvent
type DevelopmentEvent struct {
	CompanyId string                `json:"company_id,omitempty"`
	Id        int64                 `json:"id,omitempty"`
	Priority  bool                  `json:"priority,omitempty"`
	Types     DevelopmentEventTypes `json:"types,omitempty"`
}
