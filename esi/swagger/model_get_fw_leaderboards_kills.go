/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.21
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Top 4 rankings of factions by number of kills from yesterday, last week and in total
type GetFwLeaderboardsKills struct {
	// Top 4 ranking of factions active in faction warfare by total kills. A faction is considered \"active\" if they have participated in faction warfare in the past 14 days
	ActiveTotal []GetFwLeaderboardsActiveTotalActiveTotal `json:"active_total"`
	// Top 4 ranking of factions by kills in the past week
	LastWeek []GetFwLeaderboardsLastWeekLastWeek `json:"last_week"`
	// Top 4 ranking of factions by kills in the past day
	Yesterday []GetFwLeaderboardsYesterdayYesterday `json:"yesterday"`
}
