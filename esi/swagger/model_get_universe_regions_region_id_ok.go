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
type GetUniverseRegionsRegionIdOk struct {
	// constellations array
	Constellations []int32 `json:"constellations"`
	// description string
	Description string `json:"description,omitempty"`
	// name string
	Name string `json:"name"`
	// region_id integer
	RegionId int32 `json:"region_id"`
}
