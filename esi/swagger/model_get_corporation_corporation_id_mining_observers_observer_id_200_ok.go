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
type GetCorporationCorporationIdMiningObserversObserverId200Ok struct {
	// The character that did the mining
	CharacterId int32 `json:"character_id"`
	// last_updated string
	LastUpdated string `json:"last_updated"`
	// quantity integer
	Quantity int64 `json:"quantity"`
	// The corporation id of the character at the time data was recorded.
	RecordedCorporationId int32 `json:"recorded_corporation_id"`
	// type_id integer
	TypeId int32 `json:"type_id"`
}
