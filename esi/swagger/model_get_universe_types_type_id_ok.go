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
type GetUniverseTypesTypeIdOk struct {
	// capacity number
	Capacity float32 `json:"capacity,omitempty"`
	// description string
	Description string `json:"description"`
	// dogma_attributes array
	DogmaAttributes []GetUniverseTypesTypeIdDogmaAttribute `json:"dogma_attributes,omitempty"`
	// dogma_effects array
	DogmaEffects []GetUniverseTypesTypeIdDogmaEffect `json:"dogma_effects,omitempty"`
	// graphic_id integer
	GraphicId int32 `json:"graphic_id,omitempty"`
	// group_id integer
	GroupId int32 `json:"group_id"`
	// icon_id integer
	IconId int32 `json:"icon_id,omitempty"`
	// This only exists for types that can be put on the market
	MarketGroupId int32 `json:"market_group_id,omitempty"`
	// mass number
	Mass float32 `json:"mass,omitempty"`
	// name string
	Name string `json:"name"`
	// packaged_volume number
	PackagedVolume float32 `json:"packaged_volume,omitempty"`
	// portion_size integer
	PortionSize int32 `json:"portion_size,omitempty"`
	// published boolean
	Published bool `json:"published"`
	// radius number
	Radius float32 `json:"radius,omitempty"`
	// type_id integer
	TypeId int32 `json:"type_id"`
	// volume number
	Volume float32 `json:"volume,omitempty"`
}
