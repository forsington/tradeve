# {{classname}}

All URIs are relative to *https://esi.evetech.net/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUniverseAncestries**](UniverseApi.md#GetUniverseAncestries) | **Get** /universe/ancestries/ | Get ancestries
[**GetUniverseAsteroidBeltsAsteroidBeltId**](UniverseApi.md#GetUniverseAsteroidBeltsAsteroidBeltId) | **Get** /universe/asteroid_belts/{asteroid_belt_id}/ | Get asteroid belt information
[**GetUniverseBloodlines**](UniverseApi.md#GetUniverseBloodlines) | **Get** /universe/bloodlines/ | Get bloodlines
[**GetUniverseCategories**](UniverseApi.md#GetUniverseCategories) | **Get** /universe/categories/ | Get item categories
[**GetUniverseCategoriesCategoryId**](UniverseApi.md#GetUniverseCategoriesCategoryId) | **Get** /universe/categories/{category_id}/ | Get item category information
[**GetUniverseConstellations**](UniverseApi.md#GetUniverseConstellations) | **Get** /universe/constellations/ | Get constellations
[**GetUniverseConstellationsConstellationId**](UniverseApi.md#GetUniverseConstellationsConstellationId) | **Get** /universe/constellations/{constellation_id}/ | Get constellation information
[**GetUniverseFactions**](UniverseApi.md#GetUniverseFactions) | **Get** /universe/factions/ | Get factions
[**GetUniverseGraphics**](UniverseApi.md#GetUniverseGraphics) | **Get** /universe/graphics/ | Get graphics
[**GetUniverseGraphicsGraphicId**](UniverseApi.md#GetUniverseGraphicsGraphicId) | **Get** /universe/graphics/{graphic_id}/ | Get graphic information
[**GetUniverseGroups**](UniverseApi.md#GetUniverseGroups) | **Get** /universe/groups/ | Get item groups
[**GetUniverseGroupsGroupId**](UniverseApi.md#GetUniverseGroupsGroupId) | **Get** /universe/groups/{group_id}/ | Get item group information
[**GetUniverseMoonsMoonId**](UniverseApi.md#GetUniverseMoonsMoonId) | **Get** /universe/moons/{moon_id}/ | Get moon information
[**GetUniversePlanetsPlanetId**](UniverseApi.md#GetUniversePlanetsPlanetId) | **Get** /universe/planets/{planet_id}/ | Get planet information
[**GetUniverseRaces**](UniverseApi.md#GetUniverseRaces) | **Get** /universe/races/ | Get character races
[**GetUniverseRegions**](UniverseApi.md#GetUniverseRegions) | **Get** /universe/regions/ | Get regions
[**GetUniverseRegionsRegionId**](UniverseApi.md#GetUniverseRegionsRegionId) | **Get** /universe/regions/{region_id}/ | Get region information
[**GetUniverseStargatesStargateId**](UniverseApi.md#GetUniverseStargatesStargateId) | **Get** /universe/stargates/{stargate_id}/ | Get stargate information
[**GetUniverseStarsStarId**](UniverseApi.md#GetUniverseStarsStarId) | **Get** /universe/stars/{star_id}/ | Get star information
[**GetUniverseStationsStationId**](UniverseApi.md#GetUniverseStationsStationId) | **Get** /universe/stations/{station_id}/ | Get station information
[**GetUniverseStructures**](UniverseApi.md#GetUniverseStructures) | **Get** /universe/structures/ | List all public structures
[**GetUniverseStructuresStructureId**](UniverseApi.md#GetUniverseStructuresStructureId) | **Get** /universe/structures/{structure_id}/ | Get structure information
[**GetUniverseSystemJumps**](UniverseApi.md#GetUniverseSystemJumps) | **Get** /universe/system_jumps/ | Get system jumps
[**GetUniverseSystemKills**](UniverseApi.md#GetUniverseSystemKills) | **Get** /universe/system_kills/ | Get system kills
[**GetUniverseSystems**](UniverseApi.md#GetUniverseSystems) | **Get** /universe/systems/ | Get solar systems
[**GetUniverseSystemsSystemId**](UniverseApi.md#GetUniverseSystemsSystemId) | **Get** /universe/systems/{system_id}/ | Get solar system information
[**GetUniverseTypes**](UniverseApi.md#GetUniverseTypes) | **Get** /universe/types/ | Get types
[**GetUniverseTypesTypeId**](UniverseApi.md#GetUniverseTypesTypeId) | **Get** /universe/types/{type_id}/ | Get type information
[**PostUniverseIds**](UniverseApi.md#PostUniverseIds) | **Post** /universe/ids/ | Bulk names to IDs
[**PostUniverseNames**](UniverseApi.md#PostUniverseNames) | **Post** /universe/names/ | Get names and categories for a set of IDs

# **GetUniverseAncestries**
> []GetUniverseAncestries200Ok GetUniverseAncestries(ctx, optional)
Get ancestries

Get all character ancestries  --- Alternate route: `/legacy/universe/ancestries/`  Alternate route: `/v1/universe/ancestries/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UniverseApiGetUniverseAncestriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseAncestriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **optional.String**| Language to use in the response | [default to en]
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **optional.String**| Language to use in the response, takes precedence over Accept-Language | [default to en]

### Return type

[**[]GetUniverseAncestries200Ok**](get_universe_ancestries_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseAsteroidBeltsAsteroidBeltId**
> GetUniverseAsteroidBeltsAsteroidBeltIdOk GetUniverseAsteroidBeltsAsteroidBeltId(ctx, asteroidBeltId, optional)
Get asteroid belt information

Get information on an asteroid belt  --- Alternate route: `/legacy/universe/asteroid_belts/{asteroid_belt_id}/`  Alternate route: `/v1/universe/asteroid_belts/{asteroid_belt_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **asteroidBeltId** | **int32**| asteroid_belt_id integer | 
 **optional** | ***UniverseApiGetUniverseAsteroidBeltsAsteroidBeltIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseAsteroidBeltsAsteroidBeltIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetUniverseAsteroidBeltsAsteroidBeltIdOk**](get_universe_asteroid_belts_asteroid_belt_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseBloodlines**
> []GetUniverseBloodlines200Ok GetUniverseBloodlines(ctx, optional)
Get bloodlines

Get a list of bloodlines  --- Alternate route: `/legacy/universe/bloodlines/`  Alternate route: `/v1/universe/bloodlines/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UniverseApiGetUniverseBloodlinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseBloodlinesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **optional.String**| Language to use in the response | [default to en]
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **optional.String**| Language to use in the response, takes precedence over Accept-Language | [default to en]

### Return type

[**[]GetUniverseBloodlines200Ok**](get_universe_bloodlines_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseCategories**
> []int32 GetUniverseCategories(ctx, optional)
Get item categories

Get a list of item categories  --- Alternate route: `/legacy/universe/categories/`  Alternate route: `/v1/universe/categories/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UniverseApiGetUniverseCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseCategoriesOpts struct
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

# **GetUniverseCategoriesCategoryId**
> GetUniverseCategoriesCategoryIdOk GetUniverseCategoriesCategoryId(ctx, categoryId, optional)
Get item category information

Get information of an item category  --- Alternate route: `/legacy/universe/categories/{category_id}/`  Alternate route: `/v1/universe/categories/{category_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryId** | **int32**| An Eve item category ID | 
 **optional** | ***UniverseApiGetUniverseCategoriesCategoryIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseCategoriesCategoryIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **optional.String**| Language to use in the response | [default to en]
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **optional.String**| Language to use in the response, takes precedence over Accept-Language | [default to en]

### Return type

[**GetUniverseCategoriesCategoryIdOk**](get_universe_categories_category_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseConstellations**
> []int32 GetUniverseConstellations(ctx, optional)
Get constellations

Get a list of constellations  --- Alternate route: `/legacy/universe/constellations/`  Alternate route: `/v1/universe/constellations/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UniverseApiGetUniverseConstellationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseConstellationsOpts struct
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

# **GetUniverseConstellationsConstellationId**
> GetUniverseConstellationsConstellationIdOk GetUniverseConstellationsConstellationId(ctx, constellationId, optional)
Get constellation information

Get information on a constellation  --- Alternate route: `/legacy/universe/constellations/{constellation_id}/`  Alternate route: `/v1/universe/constellations/{constellation_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **constellationId** | **int32**| constellation_id integer | 
 **optional** | ***UniverseApiGetUniverseConstellationsConstellationIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseConstellationsConstellationIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **optional.String**| Language to use in the response | [default to en]
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **optional.String**| Language to use in the response, takes precedence over Accept-Language | [default to en]

### Return type

[**GetUniverseConstellationsConstellationIdOk**](get_universe_constellations_constellation_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseFactions**
> []GetUniverseFactions200Ok GetUniverseFactions(ctx, optional)
Get factions

Get a list of factions  --- Alternate route: `/dev/universe/factions/`  Alternate route: `/v2/universe/factions/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UniverseApiGetUniverseFactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseFactionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **optional.String**| Language to use in the response | [default to en]
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **optional.String**| Language to use in the response, takes precedence over Accept-Language | [default to en]

### Return type

[**[]GetUniverseFactions200Ok**](get_universe_factions_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseGraphics**
> []int32 GetUniverseGraphics(ctx, optional)
Get graphics

Get a list of graphics  --- Alternate route: `/legacy/universe/graphics/`  Alternate route: `/v1/universe/graphics/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UniverseApiGetUniverseGraphicsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseGraphicsOpts struct
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

# **GetUniverseGraphicsGraphicId**
> GetUniverseGraphicsGraphicIdOk GetUniverseGraphicsGraphicId(ctx, graphicId, optional)
Get graphic information

Get information on a graphic  --- Alternate route: `/dev/universe/graphics/{graphic_id}/`  Alternate route: `/legacy/universe/graphics/{graphic_id}/`  Alternate route: `/v1/universe/graphics/{graphic_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **graphicId** | **int32**| graphic_id integer | 
 **optional** | ***UniverseApiGetUniverseGraphicsGraphicIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseGraphicsGraphicIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetUniverseGraphicsGraphicIdOk**](get_universe_graphics_graphic_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseGroups**
> []int32 GetUniverseGroups(ctx, optional)
Get item groups

Get a list of item groups  --- Alternate route: `/legacy/universe/groups/`  Alternate route: `/v1/universe/groups/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UniverseApiGetUniverseGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **optional.Int32**| Which page of results to return | [default to 1]

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseGroupsGroupId**
> GetUniverseGroupsGroupIdOk GetUniverseGroupsGroupId(ctx, groupId, optional)
Get item group information

Get information on an item group  --- Alternate route: `/dev/universe/groups/{group_id}/`  Alternate route: `/legacy/universe/groups/{group_id}/`  Alternate route: `/v1/universe/groups/{group_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int32**| An Eve item group ID | 
 **optional** | ***UniverseApiGetUniverseGroupsGroupIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseGroupsGroupIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **optional.String**| Language to use in the response | [default to en]
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **optional.String**| Language to use in the response, takes precedence over Accept-Language | [default to en]

### Return type

[**GetUniverseGroupsGroupIdOk**](get_universe_groups_group_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseMoonsMoonId**
> GetUniverseMoonsMoonIdOk GetUniverseMoonsMoonId(ctx, moonId, optional)
Get moon information

Get information on a moon  --- Alternate route: `/legacy/universe/moons/{moon_id}/`  Alternate route: `/v1/universe/moons/{moon_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moonId** | **int32**| moon_id integer | 
 **optional** | ***UniverseApiGetUniverseMoonsMoonIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseMoonsMoonIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetUniverseMoonsMoonIdOk**](get_universe_moons_moon_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniversePlanetsPlanetId**
> GetUniversePlanetsPlanetIdOk GetUniversePlanetsPlanetId(ctx, planetId, optional)
Get planet information

Get information on a planet  --- Alternate route: `/legacy/universe/planets/{planet_id}/`  Alternate route: `/v1/universe/planets/{planet_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **planetId** | **int32**| planet_id integer | 
 **optional** | ***UniverseApiGetUniversePlanetsPlanetIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniversePlanetsPlanetIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetUniversePlanetsPlanetIdOk**](get_universe_planets_planet_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseRaces**
> []GetUniverseRaces200Ok GetUniverseRaces(ctx, optional)
Get character races

Get a list of character races  --- Alternate route: `/dev/universe/races/`  Alternate route: `/legacy/universe/races/`  Alternate route: `/v1/universe/races/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UniverseApiGetUniverseRacesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseRacesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **optional.String**| Language to use in the response | [default to en]
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **optional.String**| Language to use in the response, takes precedence over Accept-Language | [default to en]

### Return type

[**[]GetUniverseRaces200Ok**](get_universe_races_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseRegions**
> []int32 GetUniverseRegions(ctx, optional)
Get regions

Get a list of regions  --- Alternate route: `/legacy/universe/regions/`  Alternate route: `/v1/universe/regions/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UniverseApiGetUniverseRegionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseRegionsOpts struct
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

# **GetUniverseRegionsRegionId**
> GetUniverseRegionsRegionIdOk GetUniverseRegionsRegionId(ctx, regionId, optional)
Get region information

Get information on a region  --- Alternate route: `/legacy/universe/regions/{region_id}/`  Alternate route: `/v1/universe/regions/{region_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **regionId** | **int32**| region_id integer | 
 **optional** | ***UniverseApiGetUniverseRegionsRegionIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseRegionsRegionIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **optional.String**| Language to use in the response | [default to en]
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **optional.String**| Language to use in the response, takes precedence over Accept-Language | [default to en]

### Return type

[**GetUniverseRegionsRegionIdOk**](get_universe_regions_region_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseStargatesStargateId**
> GetUniverseStargatesStargateIdOk GetUniverseStargatesStargateId(ctx, stargateId, optional)
Get stargate information

Get information on a stargate  --- Alternate route: `/legacy/universe/stargates/{stargate_id}/`  Alternate route: `/v1/universe/stargates/{stargate_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stargateId** | **int32**| stargate_id integer | 
 **optional** | ***UniverseApiGetUniverseStargatesStargateIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseStargatesStargateIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetUniverseStargatesStargateIdOk**](get_universe_stargates_stargate_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseStarsStarId**
> GetUniverseStarsStarIdOk GetUniverseStarsStarId(ctx, starId, optional)
Get star information

Get information on a star  --- Alternate route: `/legacy/universe/stars/{star_id}/`  Alternate route: `/v1/universe/stars/{star_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **starId** | **int32**| star_id integer | 
 **optional** | ***UniverseApiGetUniverseStarsStarIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseStarsStarIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetUniverseStarsStarIdOk**](get_universe_stars_star_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseStationsStationId**
> GetUniverseStationsStationIdOk GetUniverseStationsStationId(ctx, stationId, optional)
Get station information

Get information on a station  --- Alternate route: `/dev/universe/stations/{station_id}/`  Alternate route: `/v2/universe/stations/{station_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **stationId** | **int32**| station_id integer | 
 **optional** | ***UniverseApiGetUniverseStationsStationIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseStationsStationIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetUniverseStationsStationIdOk**](get_universe_stations_station_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseStructures**
> []int64 GetUniverseStructures(ctx, optional)
List all public structures

List all public structures  --- Alternate route: `/dev/universe/structures/`  Alternate route: `/legacy/universe/structures/`  Alternate route: `/v1/universe/structures/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UniverseApiGetUniverseStructuresOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseStructuresOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **filter** | **optional.String**| Only list public structures that have this service online | 
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

**[]int64**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseStructuresStructureId**
> GetUniverseStructuresStructureIdOk GetUniverseStructuresStructureId(ctx, structureId, optional)
Get structure information

Returns information on requested structure if you are on the ACL. Otherwise, returns \"Forbidden\" for all inputs.  --- Alternate route: `/v2/universe/structures/{structure_id}/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **structureId** | **int64**| An Eve structure ID | 
 **optional** | ***UniverseApiGetUniverseStructuresStructureIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseStructuresStructureIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**GetUniverseStructuresStructureIdOk**](get_universe_structures_structure_id_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseSystemJumps**
> []GetUniverseSystemJumps200Ok GetUniverseSystemJumps(ctx, optional)
Get system jumps

Get the number of jumps in solar systems within the last hour ending at the timestamp of the Last-Modified header, excluding wormhole space. Only systems with jumps will be listed  --- Alternate route: `/legacy/universe/system_jumps/`  Alternate route: `/v1/universe/system_jumps/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UniverseApiGetUniverseSystemJumpsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseSystemJumpsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**[]GetUniverseSystemJumps200Ok**](get_universe_system_jumps_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseSystemKills**
> []GetUniverseSystemKills200Ok GetUniverseSystemKills(ctx, optional)
Get system kills

Get the number of ship, pod and NPC kills per solar system within the last hour ending at the timestamp of the Last-Modified header, excluding wormhole space. Only systems with kills will be listed  --- Alternate route: `/v2/universe/system_kills/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UniverseApiGetUniverseSystemKillsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseSystemKillsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**[]GetUniverseSystemKills200Ok**](get_universe_system_kills_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseSystems**
> []int32 GetUniverseSystems(ctx, optional)
Get solar systems

Get a list of solar systems  --- Alternate route: `/dev/universe/systems/`  Alternate route: `/legacy/universe/systems/`  Alternate route: `/v1/universe/systems/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UniverseApiGetUniverseSystemsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseSystemsOpts struct
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

# **GetUniverseSystemsSystemId**
> GetUniverseSystemsSystemIdOk GetUniverseSystemsSystemId(ctx, systemId, optional)
Get solar system information

Get information on a solar system.  --- Alternate route: `/dev/universe/systems/{system_id}/`  Alternate route: `/v4/universe/systems/{system_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **systemId** | **int32**| system_id integer | 
 **optional** | ***UniverseApiGetUniverseSystemsSystemIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseSystemsSystemIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **optional.String**| Language to use in the response | [default to en]
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **optional.String**| Language to use in the response, takes precedence over Accept-Language | [default to en]

### Return type

[**GetUniverseSystemsSystemIdOk**](get_universe_systems_system_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseTypes**
> []int32 GetUniverseTypes(ctx, optional)
Get types

Get a list of type ids  --- Alternate route: `/legacy/universe/types/`  Alternate route: `/v1/universe/types/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UniverseApiGetUniverseTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseTypesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **optional.Int32**| Which page of results to return | [default to 1]

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUniverseTypesTypeId**
> GetUniverseTypesTypeIdOk GetUniverseTypesTypeId(ctx, typeId, optional)
Get type information

Get information on a type  --- Alternate route: `/dev/universe/types/{type_id}/`  Alternate route: `/v3/universe/types/{type_id}/`  --- This route expires daily at 11:05

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **typeId** | **int32**| An Eve item type ID | 
 **optional** | ***UniverseApiGetUniverseTypesTypeIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiGetUniverseTypesTypeIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **optional.String**| Language to use in the response | [default to en]
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **optional.String**| Language to use in the response, takes precedence over Accept-Language | [default to en]

### Return type

[**GetUniverseTypesTypeIdOk**](get_universe_types_type_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUniverseIds**
> PostUniverseIdsOk PostUniverseIds(ctx, body, optional)
Bulk names to IDs

Resolve a set of names to IDs in the following categories: agents, alliances, characters, constellations, corporations factions, inventory_types, regions, stations, and systems. Only exact matches will be returned. All names searched for are cached for 12 hours  --- Alternate route: `/dev/universe/ids/`  Alternate route: `/legacy/universe/ids/`  Alternate route: `/v1/universe/ids/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]string**](string.md)| The names to resolve | 
 **optional** | ***UniverseApiPostUniverseIdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiPostUniverseIdsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **optional.**| Language to use in the response | [default to en]
 **datasource** | **optional.**| The server name you would like data from | [default to tranquility]
 **language** | **optional.**| Language to use in the response, takes precedence over Accept-Language | [default to en]

### Return type

[**PostUniverseIdsOk**](post_universe_ids_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostUniverseNames**
> []PostUniverseNames200Ok PostUniverseNames(ctx, body, optional)
Get names and categories for a set of IDs

Resolve a set of IDs to names and categories. Supported ID's for resolving are: Characters, Corporations, Alliances, Stations, Solar Systems, Constellations, Regions, Types, Factions  --- Alternate route: `/dev/universe/names/`  Alternate route: `/v3/universe/names/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]int32**](int32.md)| The ids to resolve | 
 **optional** | ***UniverseApiPostUniverseNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UniverseApiPostUniverseNamesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.**| The server name you would like data from | [default to tranquility]

### Return type

[**[]PostUniverseNames200Ok**](post_universe_names_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

