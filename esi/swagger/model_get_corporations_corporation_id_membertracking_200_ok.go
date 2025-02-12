/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.21
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"time"
)

// 200 ok object
type GetCorporationsCorporationIdMembertracking200Ok struct {
	// base_id integer
	BaseId int32 `json:"base_id,omitempty"`
	// character_id integer
	CharacterId int32 `json:"character_id"`
	// location_id integer
	LocationId int64 `json:"location_id,omitempty"`
	// logoff_date string
	LogoffDate time.Time `json:"logoff_date,omitempty"`
	// logon_date string
	LogonDate time.Time `json:"logon_date,omitempty"`
	// ship_type_id integer
	ShipTypeId int32 `json:"ship_type_id,omitempty"`
	// start_date string
	StartDate time.Time `json:"start_date,omitempty"`
}
