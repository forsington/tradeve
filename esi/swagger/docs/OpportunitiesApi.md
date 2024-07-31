# {{classname}}

All URIs are relative to *https://esi.evetech.net/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdOpportunities**](OpportunitiesApi.md#GetCharactersCharacterIdOpportunities) | **Get** /characters/{character_id}/opportunities/ | Get a character&#x27;s completed tasks
[**GetOpportunitiesGroups**](OpportunitiesApi.md#GetOpportunitiesGroups) | **Get** /opportunities/groups/ | Get opportunities groups
[**GetOpportunitiesGroupsGroupId**](OpportunitiesApi.md#GetOpportunitiesGroupsGroupId) | **Get** /opportunities/groups/{group_id}/ | Get opportunities group
[**GetOpportunitiesTasks**](OpportunitiesApi.md#GetOpportunitiesTasks) | **Get** /opportunities/tasks/ | Get opportunities tasks
[**GetOpportunitiesTasksTaskId**](OpportunitiesApi.md#GetOpportunitiesTasksTaskId) | **Get** /opportunities/tasks/{task_id}/ | Get opportunities task

# **GetCharactersCharacterIdOpportunities**
> []GetCharactersCharacterIdOpportunities200Ok GetCharactersCharacterIdOpportunities(ctx, characterId, optional)
Get a character's completed tasks

Return a list of tasks finished by a character  --- Alternate route: `/dev/characters/{character_id}/opportunities/`  Alternate route: `/legacy/characters/{character_id}/opportunities/`  Alternate route: `/v1/characters/{character_id}/opportunities/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***OpportunitiesApiGetCharactersCharacterIdOpportunitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpportunitiesApiGetCharactersCharacterIdOpportunitiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdOpportunities200Ok**](get_characters_character_id_opportunities_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOpportunitiesGroups**
> []int32 GetOpportunitiesGroups(ctx, optional)
Get opportunities groups

Return a list of opportunities groups  --- Alternate route: `/dev/opportunities/groups/`  Alternate route: `/legacy/opportunities/groups/`  Alternate route: `/v1/opportunities/groups/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OpportunitiesApiGetOpportunitiesGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpportunitiesApiGetOpportunitiesGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOpportunitiesGroupsGroupId**
> GetOpportunitiesGroupsGroupIdOk GetOpportunitiesGroupsGroupId(ctx, groupId, optional)
Get opportunities group

Return information of an opportunities group  --- Alternate route: `/dev/opportunities/groups/{group_id}/`  Alternate route: `/legacy/opportunities/groups/{group_id}/`  Alternate route: `/v1/opportunities/groups/{group_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int32**| ID of an opportunities group | 
 **optional** | ***OpportunitiesApiGetOpportunitiesGroupsGroupIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpportunitiesApiGetOpportunitiesGroupsGroupIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **optional.String**| Language to use in the response | [default to en]
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **optional.String**| Language to use in the response, takes precedence over Accept-Language | [default to en]

### Return type

[**GetOpportunitiesGroupsGroupIdOk**](get_opportunities_groups_group_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOpportunitiesTasks**
> []int32 GetOpportunitiesTasks(ctx, optional)
Get opportunities tasks

Return a list of opportunities tasks  --- Alternate route: `/dev/opportunities/tasks/`  Alternate route: `/legacy/opportunities/tasks/`  Alternate route: `/v1/opportunities/tasks/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OpportunitiesApiGetOpportunitiesTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpportunitiesApiGetOpportunitiesTasksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOpportunitiesTasksTaskId**
> GetOpportunitiesTasksTaskIdOk GetOpportunitiesTasksTaskId(ctx, taskId, optional)
Get opportunities task

Return information of an opportunities task  --- Alternate route: `/dev/opportunities/tasks/{task_id}/`  Alternate route: `/legacy/opportunities/tasks/{task_id}/`  Alternate route: `/v1/opportunities/tasks/{task_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **int32**| ID of an opportunities task | 
 **optional** | ***OpportunitiesApiGetOpportunitiesTasksTaskIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpportunitiesApiGetOpportunitiesTasksTaskIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetOpportunitiesTasksTaskIdOk**](get_opportunities_tasks_task_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

