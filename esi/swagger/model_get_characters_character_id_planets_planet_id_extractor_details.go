/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.21
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// extractor_details object
type GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails struct {
	// in seconds
	CycleTime int32 `json:"cycle_time,omitempty"`
	// head_radius number
	HeadRadius float32 `json:"head_radius,omitempty"`
	// heads array
	Heads []GetCharactersCharacterIdPlanetsPlanetIdHead `json:"heads"`
	// product_type_id integer
	ProductTypeId int32 `json:"product_type_id,omitempty"`
	// qty_per_cycle integer
	QtyPerCycle int32 `json:"qty_per_cycle,omitempty"`
}
