# {{classname}}

All URIs are relative to *https://esi.evetech.net/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFleetsFleetIdMembersMemberId**](FleetsApi.md#DeleteFleetsFleetIdMembersMemberId) | **Delete** /fleets/{fleet_id}/members/{member_id}/ | Kick fleet member
[**DeleteFleetsFleetIdSquadsSquadId**](FleetsApi.md#DeleteFleetsFleetIdSquadsSquadId) | **Delete** /fleets/{fleet_id}/squads/{squad_id}/ | Delete fleet squad
[**DeleteFleetsFleetIdWingsWingId**](FleetsApi.md#DeleteFleetsFleetIdWingsWingId) | **Delete** /fleets/{fleet_id}/wings/{wing_id}/ | Delete fleet wing
[**GetCharactersCharacterIdFleet**](FleetsApi.md#GetCharactersCharacterIdFleet) | **Get** /characters/{character_id}/fleet/ | Get character fleet info
[**GetFleetsFleetId**](FleetsApi.md#GetFleetsFleetId) | **Get** /fleets/{fleet_id}/ | Get fleet information
[**GetFleetsFleetIdMembers**](FleetsApi.md#GetFleetsFleetIdMembers) | **Get** /fleets/{fleet_id}/members/ | Get fleet members
[**GetFleetsFleetIdWings**](FleetsApi.md#GetFleetsFleetIdWings) | **Get** /fleets/{fleet_id}/wings/ | Get fleet wings
[**PostFleetsFleetIdMembers**](FleetsApi.md#PostFleetsFleetIdMembers) | **Post** /fleets/{fleet_id}/members/ | Create fleet invitation
[**PostFleetsFleetIdWings**](FleetsApi.md#PostFleetsFleetIdWings) | **Post** /fleets/{fleet_id}/wings/ | Create fleet wing
[**PostFleetsFleetIdWingsWingIdSquads**](FleetsApi.md#PostFleetsFleetIdWingsWingIdSquads) | **Post** /fleets/{fleet_id}/wings/{wing_id}/squads/ | Create fleet squad
[**PutFleetsFleetId**](FleetsApi.md#PutFleetsFleetId) | **Put** /fleets/{fleet_id}/ | Update fleet
[**PutFleetsFleetIdMembersMemberId**](FleetsApi.md#PutFleetsFleetIdMembersMemberId) | **Put** /fleets/{fleet_id}/members/{member_id}/ | Move fleet member
[**PutFleetsFleetIdSquadsSquadId**](FleetsApi.md#PutFleetsFleetIdSquadsSquadId) | **Put** /fleets/{fleet_id}/squads/{squad_id}/ | Rename fleet squad
[**PutFleetsFleetIdWingsWingId**](FleetsApi.md#PutFleetsFleetIdWingsWingId) | **Put** /fleets/{fleet_id}/wings/{wing_id}/ | Rename fleet wing

# **DeleteFleetsFleetIdMembersMemberId**
> DeleteFleetsFleetIdMembersMemberId(ctx, fleetId, memberId, optional)
Kick fleet member

Kick a fleet member  --- Alternate route: `/dev/fleets/{fleet_id}/members/{member_id}/`  Alternate route: `/legacy/fleets/{fleet_id}/members/{member_id}/`  Alternate route: `/v1/fleets/{fleet_id}/members/{member_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fleetId** | **int64**| ID for a fleet | 
  **memberId** | **int32**| The character ID of a member in this fleet | 
 **optional** | ***FleetsApiDeleteFleetsFleetIdMembersMemberIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FleetsApiDeleteFleetsFleetIdMembersMemberIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFleetsFleetIdSquadsSquadId**
> DeleteFleetsFleetIdSquadsSquadId(ctx, fleetId, squadId, optional)
Delete fleet squad

Delete a fleet squad, only empty squads can be deleted  --- Alternate route: `/dev/fleets/{fleet_id}/squads/{squad_id}/`  Alternate route: `/legacy/fleets/{fleet_id}/squads/{squad_id}/`  Alternate route: `/v1/fleets/{fleet_id}/squads/{squad_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fleetId** | **int64**| ID for a fleet | 
  **squadId** | **int64**| The squad to delete | 
 **optional** | ***FleetsApiDeleteFleetsFleetIdSquadsSquadIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FleetsApiDeleteFleetsFleetIdSquadsSquadIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFleetsFleetIdWingsWingId**
> DeleteFleetsFleetIdWingsWingId(ctx, fleetId, wingId, optional)
Delete fleet wing

Delete a fleet wing, only empty wings can be deleted. The wing may contain squads, but the squads must be empty  --- Alternate route: `/dev/fleets/{fleet_id}/wings/{wing_id}/`  Alternate route: `/legacy/fleets/{fleet_id}/wings/{wing_id}/`  Alternate route: `/v1/fleets/{fleet_id}/wings/{wing_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fleetId** | **int64**| ID for a fleet | 
  **wingId** | **int64**| The wing to delete | 
 **optional** | ***FleetsApiDeleteFleetsFleetIdWingsWingIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FleetsApiDeleteFleetsFleetIdWingsWingIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdFleet**
> GetCharactersCharacterIdFleetOk GetCharactersCharacterIdFleet(ctx, characterId, optional)
Get character fleet info

Return the fleet ID the character is in, if any.  --- Alternate route: `/legacy/characters/{character_id}/fleet/`  Alternate route: `/v1/characters/{character_id}/fleet/`  --- This route is cached for up to 60 seconds  --- Warning: This route has an upgrade available  --- [Diff of the upcoming changes](https://esi.evetech.net/diff/latest/dev/#GET-/characters/{character_id}/fleet/)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***FleetsApiGetCharactersCharacterIdFleetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FleetsApiGetCharactersCharacterIdFleetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**GetCharactersCharacterIdFleetOk**](get_characters_character_id_fleet_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFleetsFleetId**
> GetFleetsFleetIdOk GetFleetsFleetId(ctx, fleetId, optional)
Get fleet information

Return details about a fleet  --- Alternate route: `/dev/fleets/{fleet_id}/`  Alternate route: `/legacy/fleets/{fleet_id}/`  Alternate route: `/v1/fleets/{fleet_id}/`  --- This route is cached for up to 5 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fleetId** | **int64**| ID for a fleet | 
 **optional** | ***FleetsApiGetFleetsFleetIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FleetsApiGetFleetsFleetIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**GetFleetsFleetIdOk**](get_fleets_fleet_id_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFleetsFleetIdMembers**
> []GetFleetsFleetIdMembers200Ok GetFleetsFleetIdMembers(ctx, fleetId, optional)
Get fleet members

Return information about fleet members  --- Alternate route: `/dev/fleets/{fleet_id}/members/`  Alternate route: `/legacy/fleets/{fleet_id}/members/`  Alternate route: `/v1/fleets/{fleet_id}/members/`  --- This route is cached for up to 5 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fleetId** | **int64**| ID for a fleet | 
 **optional** | ***FleetsApiGetFleetsFleetIdMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FleetsApiGetFleetsFleetIdMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **optional.String**| Language to use in the response | [default to en]
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **optional.String**| Language to use in the response, takes precedence over Accept-Language | [default to en]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetFleetsFleetIdMembers200Ok**](get_fleets_fleet_id_members_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFleetsFleetIdWings**
> []GetFleetsFleetIdWings200Ok GetFleetsFleetIdWings(ctx, fleetId, optional)
Get fleet wings

Return information about wings in a fleet  --- Alternate route: `/dev/fleets/{fleet_id}/wings/`  Alternate route: `/legacy/fleets/{fleet_id}/wings/`  Alternate route: `/v1/fleets/{fleet_id}/wings/`  --- This route is cached for up to 5 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fleetId** | **int64**| ID for a fleet | 
 **optional** | ***FleetsApiGetFleetsFleetIdWingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FleetsApiGetFleetsFleetIdWingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **optional.String**| Language to use in the response | [default to en]
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **optional.String**| Language to use in the response, takes precedence over Accept-Language | [default to en]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetFleetsFleetIdWings200Ok**](get_fleets_fleet_id_wings_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFleetsFleetIdMembers**
> PostFleetsFleetIdMembers(ctx, body, fleetId, optional)
Create fleet invitation

Invite a character into the fleet. If a character has a CSPA charge set it is not possible to invite them to the fleet using ESI  --- Alternate route: `/dev/fleets/{fleet_id}/members/`  Alternate route: `/legacy/fleets/{fleet_id}/members/`  Alternate route: `/v1/fleets/{fleet_id}/members/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostFleetsFleetIdMembersInvitation**](PostFleetsFleetIdMembersInvitation.md)| Details of the invitation | 
  **fleetId** | **int64**| ID for a fleet | 
 **optional** | ***FleetsApiPostFleetsFleetIdMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FleetsApiPostFleetsFleetIdMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datasource** | **optional.**| The server name you would like data from | [default to tranquility]
 **token** | **optional.**| Access token to use if unable to set a header | 

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFleetsFleetIdWings**
> PostFleetsFleetIdWingsCreated PostFleetsFleetIdWings(ctx, fleetId, optional)
Create fleet wing

Create a new wing in a fleet  --- Alternate route: `/dev/fleets/{fleet_id}/wings/`  Alternate route: `/legacy/fleets/{fleet_id}/wings/`  Alternate route: `/v1/fleets/{fleet_id}/wings/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fleetId** | **int64**| ID for a fleet | 
 **optional** | ***FleetsApiPostFleetsFleetIdWingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FleetsApiPostFleetsFleetIdWingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**PostFleetsFleetIdWingsCreated**](post_fleets_fleet_id_wings_created.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFleetsFleetIdWingsWingIdSquads**
> PostFleetsFleetIdWingsWingIdSquadsCreated PostFleetsFleetIdWingsWingIdSquads(ctx, fleetId, wingId, optional)
Create fleet squad

Create a new squad in a fleet  --- Alternate route: `/dev/fleets/{fleet_id}/wings/{wing_id}/squads/`  Alternate route: `/legacy/fleets/{fleet_id}/wings/{wing_id}/squads/`  Alternate route: `/v1/fleets/{fleet_id}/wings/{wing_id}/squads/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fleetId** | **int64**| ID for a fleet | 
  **wingId** | **int64**| The wing_id to create squad in | 
 **optional** | ***FleetsApiPostFleetsFleetIdWingsWingIdSquadsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FleetsApiPostFleetsFleetIdWingsWingIdSquadsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**PostFleetsFleetIdWingsWingIdSquadsCreated**](post_fleets_fleet_id_wings_wing_id_squads_created.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFleetsFleetId**
> PutFleetsFleetId(ctx, body, fleetId, optional)
Update fleet

Update settings about a fleet  --- Alternate route: `/dev/fleets/{fleet_id}/`  Alternate route: `/legacy/fleets/{fleet_id}/`  Alternate route: `/v1/fleets/{fleet_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutFleetsFleetIdNewSettings**](PutFleetsFleetIdNewSettings.md)| What to update for this fleet | 
  **fleetId** | **int64**| ID for a fleet | 
 **optional** | ***FleetsApiPutFleetsFleetIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FleetsApiPutFleetsFleetIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datasource** | **optional.**| The server name you would like data from | [default to tranquility]
 **token** | **optional.**| Access token to use if unable to set a header | 

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFleetsFleetIdMembersMemberId**
> PutFleetsFleetIdMembersMemberId(ctx, body, fleetId, memberId, optional)
Move fleet member

Move a fleet member around  --- Alternate route: `/dev/fleets/{fleet_id}/members/{member_id}/`  Alternate route: `/legacy/fleets/{fleet_id}/members/{member_id}/`  Alternate route: `/v1/fleets/{fleet_id}/members/{member_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutFleetsFleetIdMembersMemberIdMovement**](PutFleetsFleetIdMembersMemberIdMovement.md)| Details of the invitation | 
  **fleetId** | **int64**| ID for a fleet | 
  **memberId** | **int32**| The character ID of a member in this fleet | 
 **optional** | ***FleetsApiPutFleetsFleetIdMembersMemberIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FleetsApiPutFleetsFleetIdMembersMemberIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **datasource** | **optional.**| The server name you would like data from | [default to tranquility]
 **token** | **optional.**| Access token to use if unable to set a header | 

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFleetsFleetIdSquadsSquadId**
> PutFleetsFleetIdSquadsSquadId(ctx, body, fleetId, squadId, optional)
Rename fleet squad

Rename a fleet squad  --- Alternate route: `/dev/fleets/{fleet_id}/squads/{squad_id}/`  Alternate route: `/legacy/fleets/{fleet_id}/squads/{squad_id}/`  Alternate route: `/v1/fleets/{fleet_id}/squads/{squad_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutFleetsFleetIdSquadsSquadIdNaming**](PutFleetsFleetIdSquadsSquadIdNaming.md)| New name of the squad | 
  **fleetId** | **int64**| ID for a fleet | 
  **squadId** | **int64**| The squad to rename | 
 **optional** | ***FleetsApiPutFleetsFleetIdSquadsSquadIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FleetsApiPutFleetsFleetIdSquadsSquadIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **datasource** | **optional.**| The server name you would like data from | [default to tranquility]
 **token** | **optional.**| Access token to use if unable to set a header | 

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutFleetsFleetIdWingsWingId**
> PutFleetsFleetIdWingsWingId(ctx, body, fleetId, wingId, optional)
Rename fleet wing

Rename a fleet wing  --- Alternate route: `/dev/fleets/{fleet_id}/wings/{wing_id}/`  Alternate route: `/legacy/fleets/{fleet_id}/wings/{wing_id}/`  Alternate route: `/v1/fleets/{fleet_id}/wings/{wing_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutFleetsFleetIdWingsWingIdNaming**](PutFleetsFleetIdWingsWingIdNaming.md)| New name of the wing | 
  **fleetId** | **int64**| ID for a fleet | 
  **wingId** | **int64**| The wing to rename | 
 **optional** | ***FleetsApiPutFleetsFleetIdWingsWingIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FleetsApiPutFleetsFleetIdWingsWingIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **datasource** | **optional.**| The server name you would like data from | [default to tranquility]
 **token** | **optional.**| Access token to use if unable to set a header | 

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

