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
type GetCorporationsCorporationIdWallets200Ok struct {
	// balance number
	Balance float64 `json:"balance"`
	// division integer
	Division int32 `json:"division"`
}
