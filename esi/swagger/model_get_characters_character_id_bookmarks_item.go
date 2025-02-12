/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.21
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Optional object that is returned if a bookmark was made on a particular item.
type GetCharactersCharacterIdBookmarksItem struct {
	// item_id integer
	ItemId int64 `json:"item_id"`
	// type_id integer
	TypeId int32 `json:"type_id"`
}
