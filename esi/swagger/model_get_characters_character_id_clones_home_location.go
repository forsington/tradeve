/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.21
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// home_location object
type GetCharactersCharacterIdClonesHomeLocation struct {
	// location_id integer
	LocationId int64 `json:"location_id,omitempty"`
	// location_type string
	LocationType string `json:"location_type,omitempty"`
}
