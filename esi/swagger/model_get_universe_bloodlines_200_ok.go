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
type GetUniverseBloodlines200Ok struct {
	// bloodline_id integer
	BloodlineId int32 `json:"bloodline_id"`
	// charisma integer
	Charisma int32 `json:"charisma"`
	// corporation_id integer
	CorporationId int32 `json:"corporation_id"`
	// description string
	Description string `json:"description"`
	// intelligence integer
	Intelligence int32 `json:"intelligence"`
	// memory integer
	Memory int32 `json:"memory"`
	// name string
	Name string `json:"name"`
	// perception integer
	Perception int32 `json:"perception"`
	// race_id integer
	RaceId int32 `json:"race_id"`
	// ship_type_id integer
	ShipTypeId int32 `json:"ship_type_id"`
	// willpower integer
	Willpower int32 `json:"willpower"`
}
