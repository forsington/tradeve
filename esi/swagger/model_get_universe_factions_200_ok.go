/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.21
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 200 ok object
type GetUniverseFactions200Ok struct {
	// corporation_id integer
	CorporationId int32 `json:"corporation_id,omitempty"`
	// description string
	Description string `json:"description"`
	// faction_id integer
	FactionId int32 `json:"faction_id"`
	// is_unique boolean
	IsUnique bool `json:"is_unique"`
	// militia_corporation_id integer
	MilitiaCorporationId int32 `json:"militia_corporation_id,omitempty"`
	// name string
	Name string `json:"name"`
	// size_factor number
	SizeFactor float32 `json:"size_factor"`
	// solar_system_id integer
	SolarSystemId int32 `json:"solar_system_id,omitempty"`
	// station_count integer
	StationCount int32 `json:"station_count"`
	// station_system_count integer
	StationSystemCount int32 `json:"station_system_count"`
}
