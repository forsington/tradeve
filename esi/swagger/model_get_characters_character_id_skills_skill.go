/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.21
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// skill object
type GetCharactersCharacterIdSkillsSkill struct {
	// active_skill_level integer
	ActiveSkillLevel int32 `json:"active_skill_level"`
	// skill_id integer
	SkillId int32 `json:"skill_id"`
	// skillpoints_in_skill integer
	SkillpointsInSkill int64 `json:"skillpoints_in_skill"`
	// trained_skill_level integer
	TrainedSkillLevel int32 `json:"trained_skill_level"`
}
