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
type GetCorporationsCorporationIdMedalsIssued200Ok struct {
	// ID of the character who was rewarded this medal
	CharacterId int32 `json:"character_id"`
	// issued_at string
	IssuedAt time.Time `json:"issued_at"`
	// ID of the character who issued the medal
	IssuerId int32 `json:"issuer_id"`
	// medal_id integer
	MedalId int32 `json:"medal_id"`
	// reason string
	Reason string `json:"reason"`
	// status string
	Status string `json:"status"`
}
