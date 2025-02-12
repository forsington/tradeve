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
type GetFwStats200Ok struct {
	// faction_id integer
	FactionId int32            `json:"faction_id"`
	Kills     *GetFwStatsKills `json:"kills"`
	// How many pilots fight for the given faction
	Pilots int32 `json:"pilots"`
	// The number of solar systems controlled by the given faction
	SystemsControlled int32                    `json:"systems_controlled"`
	VictoryPoints     *GetFwStatsVictoryPoints `json:"victory_points"`
}
