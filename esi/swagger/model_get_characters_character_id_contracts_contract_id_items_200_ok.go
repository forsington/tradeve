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
type GetCharactersCharacterIdContractsContractIdItems200Ok struct {
	// true if the contract issuer has submitted this item with the contract, false if the isser is asking for this item in the contract
	IsIncluded bool `json:"is_included"`
	// is_singleton boolean
	IsSingleton bool `json:"is_singleton"`
	// Number of items in the stack
	Quantity int32 `json:"quantity"`
	// -1 indicates that the item is a singleton (non-stackable). If the item happens to be a Blueprint, -1 is an Original and -2 is a Blueprint Copy
	RawQuantity int32 `json:"raw_quantity,omitempty"`
	// Unique ID for the item
	RecordId int64 `json:"record_id"`
	// Type ID for item
	TypeId int32 `json:"type_id"`
}
