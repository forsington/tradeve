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
type GetCorporationsCorporationIdMedals200Ok struct {
	// created_at string
	CreatedAt time.Time `json:"created_at"`
	// ID of the character who created this medal
	CreatorId int32 `json:"creator_id"`
	// description string
	Description string `json:"description"`
	// medal_id integer
	MedalId int32 `json:"medal_id"`
	// title string
	Title string `json:"title"`
}
