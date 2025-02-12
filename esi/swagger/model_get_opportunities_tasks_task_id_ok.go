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
type GetOpportunitiesTasksTaskIdOk struct {
	// description string
	Description string `json:"description"`
	// name string
	Name string `json:"name"`
	// notification string
	Notification string `json:"notification"`
	// task_id integer
	TaskId int32 `json:"task_id"`
}
