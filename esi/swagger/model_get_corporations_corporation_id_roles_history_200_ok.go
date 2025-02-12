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
type GetCorporationsCorporationIdRolesHistory200Ok struct {
	// changed_at string
	ChangedAt time.Time `json:"changed_at"`
	// The character whose roles are changed
	CharacterId int32 `json:"character_id"`
	// ID of the character who issued this change
	IssuerId int32 `json:"issuer_id"`
	// new_roles array
	NewRoles []string `json:"new_roles"`
	// old_roles array
	OldRoles []string `json:"old_roles"`
	// role_type string
	RoleType string `json:"role_type"`
}
