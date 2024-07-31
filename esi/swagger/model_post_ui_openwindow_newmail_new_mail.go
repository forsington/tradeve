/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.21
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// new_mail object
type PostUiOpenwindowNewmailNewMail struct {
	// body string
	Body string `json:"body"`
	// recipients array
	Recipients []int32 `json:"recipients"`
	// subject string
	Subject string `json:"subject"`
	// to_corp_or_alliance_id integer
	ToCorpOrAllianceId int32 `json:"to_corp_or_alliance_id,omitempty"`
	// Corporations, alliances and mailing lists are all types of mailing groups. You may only send to one mailing group, at a time, so you may fill out either this field or the to_corp_or_alliance_ids field
	ToMailingListId int32 `json:"to_mailing_list_id,omitempty"`
}
