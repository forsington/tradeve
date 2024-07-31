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
type GetCharactersCharacterIdClonesOk struct {
	HomeLocation *GetCharactersCharacterIdClonesHomeLocation `json:"home_location,omitempty"`
	// jump_clones array
	JumpClones []GetCharactersCharacterIdClonesJumpClone `json:"jump_clones"`
	// last_clone_jump_date string
	LastCloneJumpDate time.Time `json:"last_clone_jump_date,omitempty"`
	// last_station_change_date string
	LastStationChangeDate time.Time `json:"last_station_change_date,omitempty"`
}
