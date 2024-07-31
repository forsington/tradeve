# {{classname}}

All URIs are relative to *https://esi.evetech.net/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdLocation**](LocationApi.md#GetCharactersCharacterIdLocation) | **Get** /characters/{character_id}/location/ | Get character location
[**GetCharactersCharacterIdOnline**](LocationApi.md#GetCharactersCharacterIdOnline) | **Get** /characters/{character_id}/online/ | Get character online
[**GetCharactersCharacterIdShip**](LocationApi.md#GetCharactersCharacterIdShip) | **Get** /characters/{character_id}/ship/ | Get current ship

# **GetCharactersCharacterIdLocation**
> GetCharactersCharacterIdLocationOk GetCharactersCharacterIdLocation(ctx, characterId, optional)
Get character location

Information about the characters current location. Returns the current solar system id, and also the current station or structure ID if applicable  --- Alternate route: `/dev/characters/{character_id}/location/`  Alternate route: `/legacy/characters/{character_id}/location/`  Alternate route: `/v1/characters/{character_id}/location/`  Alternate route: `/v2/characters/{character_id}/location/`  --- This route is cached for up to 5 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***LocationApiGetCharactersCharacterIdLocationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiGetCharactersCharacterIdLocationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**GetCharactersCharacterIdLocationOk**](get_characters_character_id_location_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdOnline**
> GetCharactersCharacterIdOnlineOk GetCharactersCharacterIdOnline(ctx, characterId, optional)
Get character online

Checks if the character is currently online  --- Alternate route: `/dev/characters/{character_id}/online/`  Alternate route: `/v2/characters/{character_id}/online/`  Alternate route: `/v3/characters/{character_id}/online/`  --- This route is cached for up to 60 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***LocationApiGetCharactersCharacterIdOnlineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiGetCharactersCharacterIdOnlineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**GetCharactersCharacterIdOnlineOk**](get_characters_character_id_online_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdShip**
> GetCharactersCharacterIdShipOk GetCharactersCharacterIdShip(ctx, characterId, optional)
Get current ship

Get the current ship type, name and id  --- Alternate route: `/dev/characters/{character_id}/ship/`  Alternate route: `/legacy/characters/{character_id}/ship/`  Alternate route: `/v1/characters/{character_id}/ship/`  Alternate route: `/v2/characters/{character_id}/ship/`  --- This route is cached for up to 5 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***LocationApiGetCharactersCharacterIdShipOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationApiGetCharactersCharacterIdShipOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**GetCharactersCharacterIdShipOk**](get_characters_character_id_ship_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

