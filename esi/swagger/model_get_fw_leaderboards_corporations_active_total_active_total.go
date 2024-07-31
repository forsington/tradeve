/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.21
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// active_total object
type GetFwLeaderboardsCorporationsActiveTotalActiveTotal struct {
	// Amount of kills
	Amount int32 `json:"amount,omitempty"`
	// corporation_id integer
	CorporationId int32 `json:"corporation_id,omitempty"`
}
