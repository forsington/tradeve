# {{classname}}

All URIs are relative to *https://esi.evetech.net/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdPlanets**](PlanetaryInteractionApi.md#GetCharactersCharacterIdPlanets) | **Get** /characters/{character_id}/planets/ | Get colonies
[**GetCharactersCharacterIdPlanetsPlanetId**](PlanetaryInteractionApi.md#GetCharactersCharacterIdPlanetsPlanetId) | **Get** /characters/{character_id}/planets/{planet_id}/ | Get colony layout
[**GetCorporationsCorporationIdCustomsOffices**](PlanetaryInteractionApi.md#GetCorporationsCorporationIdCustomsOffices) | **Get** /corporations/{corporation_id}/customs_offices/ | List corporation customs offices
[**GetUniverseSchematicsSchematicId**](PlanetaryInteractionApi.md#GetUniverseSchematicsSchematicId) | **Get** /universe/schematics/{schematic_id}/ | Get schematic information

# **GetCharactersCharacterIdPlanets**
> []GetCharactersCharacterIdPlanets200Ok GetCharactersCharacterIdPlanets(ctx, characterId, optional)
Get colonies

Returns a list of all planetary colonies owned by a character.  --- Alternate route: `/dev/characters/{character_id}/planets/`  Alternate route: `/legacy/characters/{character_id}/planets/`  Alternate route: `/v1/characters/{character_id}/planets/`  --- This route is cached for up to 600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***PlanetaryInteractionApiGetCharactersCharacterIdPlanetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanetaryInteractionApiGetCharactersCharacterIdPlanetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdPlanets200Ok**](get_characters_character_id_planets_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdPlanetsPlanetId**
> GetCharactersCharacterIdPlanetsPlanetIdOk GetCharactersCharacterIdPlanetsPlanetId(ctx, characterId, planetId, optional)
Get colony layout

Returns full details on the layout of a single planetary colony, including links, pins and routes. Note: Planetary information is only recalculated when the colony is viewed through the client. Information will not update until this criteria is met.  --- Alternate route: `/dev/characters/{character_id}/planets/{planet_id}/`  Alternate route: `/v3/characters/{character_id}/planets/{planet_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
  **planetId** | **int32**| Planet id of the target planet | 
 **optional** | ***PlanetaryInteractionApiGetCharactersCharacterIdPlanetsPlanetIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanetaryInteractionApiGetCharactersCharacterIdPlanetsPlanetIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**GetCharactersCharacterIdPlanetsPlanetIdOk**](get_characters_character_id_planets_planet_id_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdCustomsOffices**
> []GetCorporationsCorporationIdCustomsOffices200Ok GetCorporationsCorporationIdCustomsOffices(ctx, corporationId, optional)
List corporation customs offices

List customs offices owned by a corporation  --- Alternate route: `/dev/corporations/{corporation_id}/customs_offices/`  Alternate route: `/legacy/corporations/{corporation_id}/customs_offices/`  Alternate route: `/v1/corporations/{corporation_id}/customs_offices/`  --- This route is cached for up to 3600 seconds  --- Requires one of the following EVE corporation role(s): Director 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | ***PlanetaryInteractionApiGetCorporationsCorporationIdCustomsOfficesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanetaryInteractionApiGetCorporationsCorporationIdCustomsOfficesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **optional.Int32**| Which page of results to return | [default to 1]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCorporationsCorporationIdCustomsOffices200Ok**](get_corporations_corporation_id_customs_offices_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseSchematicsSchematicId**
> GetUniverseSchematicsSchematicIdOk GetUniverseSchematicsSchematicId(ctx, schematicId, optional)
Get schematic information

Get information on a planetary factory schematic  --- Alternate route: `/dev/universe/schematics/{schematic_id}/`  Alternate route: `/legacy/universe/schematics/{schematic_id}/`  Alternate route: `/v1/universe/schematics/{schematic_id}/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **schematicId** | **int32**| A PI schematic ID | 
 **optional** | ***PlanetaryInteractionApiGetUniverseSchematicsSchematicIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanetaryInteractionApiGetUniverseSchematicsSchematicIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetUniverseSchematicsSchematicIdOk**](get_universe_schematics_schematic_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

