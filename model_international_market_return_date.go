/*
 * simply-wallst
 *
 * simply-wallst API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// InternationalMarketReturnDate struct for InternationalMarketReturnDate
type InternationalMarketReturnDate struct {
	CountryId    int64   `json:"country_id,omitempty"`
	Country      string  `json:"country,omitempty"`
	CountryIso   string  `json:"country_iso,omitempty"`
	Return7d     float32 `json:"return_7d,omitempty"`
	Return30d    float32 `json:"return_30d,omitempty"`
	Return90d    float32 `json:"return_90d,omitempty"`
	Return1yrAbs float32 `json:"return_1yr_abs,omitempty"`
	Return3yrAbs float32 `json:"return_3yr_abs,omitempty"`
	Return5yrAbs float32 `json:"return_5yr_abs,omitempty"`
}
