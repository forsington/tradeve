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
type GetFleetsFleetIdOk struct {
	// Is free-move enabled
	IsFreeMove bool `json:"is_free_move"`
	// Does the fleet have an active fleet advertisement
	IsRegistered bool `json:"is_registered"`
	// Is EVE Voice enabled
	IsVoiceEnabled bool `json:"is_voice_enabled"`
	// Fleet MOTD in CCP flavoured HTML
	Motd string `json:"motd"`
}
