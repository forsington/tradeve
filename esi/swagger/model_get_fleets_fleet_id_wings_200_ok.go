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
type GetFleetsFleetIdWings200Ok struct {
	// id integer
	Id int64 `json:"id"`
	// name string
	Name string `json:"name"`
	// squads array
	Squads []GetFleetsFleetIdWingsSquad `json:"squads"`
}
