# {{classname}}

All URIs are relative to *https://esi.evetech.net/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterId**](CharacterApi.md#GetCharactersCharacterId) | **Get** /characters/{character_id}/ | Get character&#x27;s public information
[**GetCharactersCharacterIdAgentsResearch**](CharacterApi.md#GetCharactersCharacterIdAgentsResearch) | **Get** /characters/{character_id}/agents_research/ | Get agents research
[**GetCharactersCharacterIdBlueprints**](CharacterApi.md#GetCharactersCharacterIdBlueprints) | **Get** /characters/{character_id}/blueprints/ | Get blueprints
[**GetCharactersCharacterIdCorporationhistory**](CharacterApi.md#GetCharactersCharacterIdCorporationhistory) | **Get** /characters/{character_id}/corporationhistory/ | Get corporation history
[**GetCharactersCharacterIdFatigue**](CharacterApi.md#GetCharactersCharacterIdFatigue) | **Get** /characters/{character_id}/fatigue/ | Get jump fatigue
[**GetCharactersCharacterIdMedals**](CharacterApi.md#GetCharactersCharacterIdMedals) | **Get** /characters/{character_id}/medals/ | Get medals
[**GetCharactersCharacterIdNotifications**](CharacterApi.md#GetCharactersCharacterIdNotifications) | **Get** /characters/{character_id}/notifications/ | Get character notifications
[**GetCharactersCharacterIdNotificationsContacts**](CharacterApi.md#GetCharactersCharacterIdNotificationsContacts) | **Get** /characters/{character_id}/notifications/contacts/ | Get new contact notifications
[**GetCharactersCharacterIdPortrait**](CharacterApi.md#GetCharactersCharacterIdPortrait) | **Get** /characters/{character_id}/portrait/ | Get character portraits
[**GetCharactersCharacterIdRoles**](CharacterApi.md#GetCharactersCharacterIdRoles) | **Get** /characters/{character_id}/roles/ | Get character corporation roles
[**GetCharactersCharacterIdStandings**](CharacterApi.md#GetCharactersCharacterIdStandings) | **Get** /characters/{character_id}/standings/ | Get standings
[**GetCharactersCharacterIdTitles**](CharacterApi.md#GetCharactersCharacterIdTitles) | **Get** /characters/{character_id}/titles/ | Get character corporation titles
[**PostCharactersAffiliation**](CharacterApi.md#PostCharactersAffiliation) | **Post** /characters/affiliation/ | Character affiliation
[**PostCharactersCharacterIdCspa**](CharacterApi.md#PostCharactersCharacterIdCspa) | **Post** /characters/{character_id}/cspa/ | Calculate a CSPA charge cost

# **GetCharactersCharacterId**
> GetCharactersCharacterIdOk GetCharactersCharacterId(ctx, characterId, optional)
Get character's public information

Public information about a character  --- Alternate route: `/dev/characters/{character_id}/`  Alternate route: `/legacy/characters/{character_id}/`  Alternate route: `/v5/characters/{character_id}/`  --- This route is cached for up to 604800 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***CharacterApiGetCharactersCharacterIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CharacterApiGetCharactersCharacterIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetCharactersCharacterIdOk**](get_characters_character_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdAgentsResearch**
> []GetCharactersCharacterIdAgentsResearch200Ok GetCharactersCharacterIdAgentsResearch(ctx, characterId, optional)
Get agents research

Return a list of agents research information for a character. The formula for finding the current research points with an agent is: currentPoints = remainderPoints + pointsPerDay * days(currentTime - researchStartDate)  --- Alternate route: `/dev/characters/{character_id}/agents_research/`  Alternate route: `/v2/characters/{character_id}/agents_research/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***CharacterApiGetCharactersCharacterIdAgentsResearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CharacterApiGetCharactersCharacterIdAgentsResearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdAgentsResearch200Ok**](get_characters_character_id_agents_research_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdBlueprints**
> []GetCharactersCharacterIdBlueprints200Ok GetCharactersCharacterIdBlueprints(ctx, characterId, optional)
Get blueprints

Return a list of blueprints the character owns  --- Alternate route: `/dev/characters/{character_id}/blueprints/`  Alternate route: `/v3/characters/{character_id}/blueprints/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***CharacterApiGetCharactersCharacterIdBlueprintsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CharacterApiGetCharactersCharacterIdBlueprintsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **optional.Int32**| Which page of results to return | [default to 1]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdBlueprints200Ok**](get_characters_character_id_blueprints_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdCorporationhistory**
> []GetCharactersCharacterIdCorporationhistory200Ok GetCharactersCharacterIdCorporationhistory(ctx, characterId, optional)
Get corporation history

Get a list of all the corporations a character has been a member of  --- Alternate route: `/dev/characters/{character_id}/corporationhistory/`  Alternate route: `/v2/characters/{character_id}/corporationhistory/`  --- This route is cached for up to 86400 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***CharacterApiGetCharactersCharacterIdCorporationhistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CharacterApiGetCharactersCharacterIdCorporationhistoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**[]GetCharactersCharacterIdCorporationhistory200Ok**](get_characters_character_id_corporationhistory_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdFatigue**
> GetCharactersCharacterIdFatigueOk GetCharactersCharacterIdFatigue(ctx, characterId, optional)
Get jump fatigue

Return a character's jump activation and fatigue information  --- Alternate route: `/dev/characters/{character_id}/fatigue/`  Alternate route: `/v2/characters/{character_id}/fatigue/`  --- This route is cached for up to 300 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***CharacterApiGetCharactersCharacterIdFatigueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CharacterApiGetCharactersCharacterIdFatigueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**GetCharactersCharacterIdFatigueOk**](get_characters_character_id_fatigue_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdMedals**
> []GetCharactersCharacterIdMedals200Ok GetCharactersCharacterIdMedals(ctx, characterId, optional)
Get medals

Return a list of medals the character has  --- Alternate route: `/dev/characters/{character_id}/medals/`  Alternate route: `/v2/characters/{character_id}/medals/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***CharacterApiGetCharactersCharacterIdMedalsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CharacterApiGetCharactersCharacterIdMedalsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdMedals200Ok**](get_characters_character_id_medals_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdNotifications**
> []GetCharactersCharacterIdNotifications200Ok GetCharactersCharacterIdNotifications(ctx, characterId, optional)
Get character notifications

Return character notifications  --- Alternate route: `/dev/characters/{character_id}/notifications/`  Alternate route: `/v5/characters/{character_id}/notifications/`  Alternate route: `/v6/characters/{character_id}/notifications/`  --- This route is cached for up to 600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***CharacterApiGetCharactersCharacterIdNotificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CharacterApiGetCharactersCharacterIdNotificationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdNotifications200Ok**](get_characters_character_id_notifications_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdNotificationsContacts**
> []GetCharactersCharacterIdNotificationsContacts200Ok GetCharactersCharacterIdNotificationsContacts(ctx, characterId, optional)
Get new contact notifications

Return notifications about having been added to someone's contact list  --- Alternate route: `/dev/characters/{character_id}/notifications/contacts/`  Alternate route: `/v2/characters/{character_id}/notifications/contacts/`  --- This route is cached for up to 600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***CharacterApiGetCharactersCharacterIdNotificationsContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CharacterApiGetCharactersCharacterIdNotificationsContactsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdNotificationsContacts200Ok**](get_characters_character_id_notifications_contacts_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdPortrait**
> GetCharactersCharacterIdPortraitOk GetCharactersCharacterIdPortrait(ctx, characterId, optional)
Get character portraits

Get portrait urls for a character  --- Alternate route: `/dev/characters/{character_id}/portrait/`  Alternate route: `/v2/characters/{character_id}/portrait/`  Alternate route: `/v3/characters/{character_id}/portrait/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***CharacterApiGetCharactersCharacterIdPortraitOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CharacterApiGetCharactersCharacterIdPortraitOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetCharactersCharacterIdPortraitOk**](get_characters_character_id_portrait_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdRoles**
> GetCharactersCharacterIdRolesOk GetCharactersCharacterIdRoles(ctx, characterId, optional)
Get character corporation roles

Returns a character's corporation roles  --- Alternate route: `/dev/characters/{character_id}/roles/`  Alternate route: `/v3/characters/{character_id}/roles/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***CharacterApiGetCharactersCharacterIdRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CharacterApiGetCharactersCharacterIdRolesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**GetCharactersCharacterIdRolesOk**](get_characters_character_id_roles_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdStandings**
> []GetCharactersCharacterIdStandings200Ok GetCharactersCharacterIdStandings(ctx, characterId, optional)
Get standings

Return character standings from agents, NPC corporations, and factions  --- Alternate route: `/dev/characters/{character_id}/standings/`  Alternate route: `/v2/characters/{character_id}/standings/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***CharacterApiGetCharactersCharacterIdStandingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CharacterApiGetCharactersCharacterIdStandingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdStandings200Ok**](get_characters_character_id_standings_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdTitles**
> []GetCharactersCharacterIdTitles200Ok GetCharactersCharacterIdTitles(ctx, characterId, optional)
Get character corporation titles

Returns a character's titles  --- Alternate route: `/dev/characters/{character_id}/titles/`  Alternate route: `/v2/characters/{character_id}/titles/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***CharacterApiGetCharactersCharacterIdTitlesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CharacterApiGetCharactersCharacterIdTitlesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdTitles200Ok**](get_characters_character_id_titles_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCharactersAffiliation**
> []PostCharactersAffiliation200Ok PostCharactersAffiliation(ctx, body, optional)
Character affiliation

Bulk lookup of character IDs to corporation, alliance and faction  --- Alternate route: `/dev/characters/affiliation/`  Alternate route: `/v2/characters/affiliation/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]int32**](int32.md)| The character IDs to fetch affiliations for. All characters must exist, or none will be returned | 
 **optional** | ***CharacterApiPostCharactersAffiliationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CharacterApiPostCharactersAffiliationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.**| The server name you would like data from | [default to tranquility]

### Return type

[**[]PostCharactersAffiliation200Ok**](post_characters_affiliation_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCharactersCharacterIdCspa**
> float32 PostCharactersCharacterIdCspa(ctx, body, characterId, optional)
Calculate a CSPA charge cost

Takes a source character ID in the url and a set of target character ID's in the body, returns a CSPA charge cost  --- Alternate route: `/dev/characters/{character_id}/cspa/`  Alternate route: `/v5/characters/{character_id}/cspa/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]int32**](int32.md)| The target characters to calculate the charge for | 
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***CharacterApiPostCharactersCharacterIdCspaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CharacterApiPostCharactersCharacterIdCspaOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datasource** | **optional.**| The server name you would like data from | [default to tranquility]
 **token** | **optional.**| Access token to use if unable to set a header | 

### Return type

**float32**

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

