# {{classname}}

All URIs are relative to *https://esi.evetech.net/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdAssets**](AssetsApi.md#GetCharactersCharacterIdAssets) | **Get** /characters/{character_id}/assets/ | Get character assets
[**GetCorporationsCorporationIdAssets**](AssetsApi.md#GetCorporationsCorporationIdAssets) | **Get** /corporations/{corporation_id}/assets/ | Get corporation assets
[**PostCharactersCharacterIdAssetsLocations**](AssetsApi.md#PostCharactersCharacterIdAssetsLocations) | **Post** /characters/{character_id}/assets/locations/ | Get character asset locations
[**PostCharactersCharacterIdAssetsNames**](AssetsApi.md#PostCharactersCharacterIdAssetsNames) | **Post** /characters/{character_id}/assets/names/ | Get character asset names
[**PostCorporationsCorporationIdAssetsLocations**](AssetsApi.md#PostCorporationsCorporationIdAssetsLocations) | **Post** /corporations/{corporation_id}/assets/locations/ | Get corporation asset locations
[**PostCorporationsCorporationIdAssetsNames**](AssetsApi.md#PostCorporationsCorporationIdAssetsNames) | **Post** /corporations/{corporation_id}/assets/names/ | Get corporation asset names

# **GetCharactersCharacterIdAssets**
> []GetCharactersCharacterIdAssets200Ok GetCharactersCharacterIdAssets(ctx, characterId, optional)
Get character assets

Return a list of the characters assets  --- Alternate route: `/dev/characters/{character_id}/assets/`  Alternate route: `/v5/characters/{character_id}/assets/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***AssetsApiGetCharactersCharacterIdAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiGetCharactersCharacterIdAssetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **optional.Int32**| Which page of results to return | [default to 1]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdAssets200Ok**](get_characters_character_id_assets_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdAssets**
> []GetCorporationsCorporationIdAssets200Ok GetCorporationsCorporationIdAssets(ctx, corporationId, optional)
Get corporation assets

Return a list of the corporation assets  --- Alternate route: `/dev/corporations/{corporation_id}/assets/`  Alternate route: `/v5/corporations/{corporation_id}/assets/`  --- This route is cached for up to 3600 seconds  --- Requires one of the following EVE corporation role(s): Director 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | ***AssetsApiGetCorporationsCorporationIdAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiGetCorporationsCorporationIdAssetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **optional.Int32**| Which page of results to return | [default to 1]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCorporationsCorporationIdAssets200Ok**](get_corporations_corporation_id_assets_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCharactersCharacterIdAssetsLocations**
> []PostCharactersCharacterIdAssetsLocations200Ok PostCharactersCharacterIdAssetsLocations(ctx, body, characterId, optional)
Get character asset locations

Return locations for a set of item ids, which you can get from character assets endpoint. Coordinates for items in hangars or stations are set to (0,0,0)  --- Alternate route: `/dev/characters/{character_id}/assets/locations/`  Alternate route: `/v2/characters/{character_id}/assets/locations/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]int64**](int64.md)| A list of item ids | 
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***AssetsApiPostCharactersCharacterIdAssetsLocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiPostCharactersCharacterIdAssetsLocationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datasource** | **optional.**| The server name you would like data from | [default to tranquility]
 **token** | **optional.**| Access token to use if unable to set a header | 

### Return type

[**[]PostCharactersCharacterIdAssetsLocations200Ok**](post_characters_character_id_assets_locations_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCharactersCharacterIdAssetsNames**
> []PostCharactersCharacterIdAssetsNames200Ok PostCharactersCharacterIdAssetsNames(ctx, body, characterId, optional)
Get character asset names

Return names for a set of item ids, which you can get from character assets endpoint. Typically used for items that can customize names, like containers or ships.  --- Alternate route: `/dev/characters/{character_id}/assets/names/`  Alternate route: `/legacy/characters/{character_id}/assets/names/`  Alternate route: `/v1/characters/{character_id}/assets/names/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]int64**](int64.md)| A list of item ids | 
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***AssetsApiPostCharactersCharacterIdAssetsNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiPostCharactersCharacterIdAssetsNamesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datasource** | **optional.**| The server name you would like data from | [default to tranquility]
 **token** | **optional.**| Access token to use if unable to set a header | 

### Return type

[**[]PostCharactersCharacterIdAssetsNames200Ok**](post_characters_character_id_assets_names_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCorporationsCorporationIdAssetsLocations**
> []PostCorporationsCorporationIdAssetsLocations200Ok PostCorporationsCorporationIdAssetsLocations(ctx, body, corporationId, optional)
Get corporation asset locations

Return locations for a set of item ids, which you can get from corporation assets endpoint. Coordinates for items in hangars or stations are set to (0,0,0)  --- Alternate route: `/dev/corporations/{corporation_id}/assets/locations/`  Alternate route: `/v2/corporations/{corporation_id}/assets/locations/`   --- Requires one of the following EVE corporation role(s): Director 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]int64**](int64.md)| A list of item ids | 
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | ***AssetsApiPostCorporationsCorporationIdAssetsLocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiPostCorporationsCorporationIdAssetsLocationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datasource** | **optional.**| The server name you would like data from | [default to tranquility]
 **token** | **optional.**| Access token to use if unable to set a header | 

### Return type

[**[]PostCorporationsCorporationIdAssetsLocations200Ok**](post_corporations_corporation_id_assets_locations_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCorporationsCorporationIdAssetsNames**
> []PostCorporationsCorporationIdAssetsNames200Ok PostCorporationsCorporationIdAssetsNames(ctx, body, corporationId, optional)
Get corporation asset names

Return names for a set of item ids, which you can get from corporation assets endpoint. Only valid for items that can customize names, like containers or ships  --- Alternate route: `/dev/corporations/{corporation_id}/assets/names/`  Alternate route: `/legacy/corporations/{corporation_id}/assets/names/`  Alternate route: `/v1/corporations/{corporation_id}/assets/names/`   --- Requires one of the following EVE corporation role(s): Director 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]int64**](int64.md)| A list of item ids | 
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | ***AssetsApiPostCorporationsCorporationIdAssetsNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AssetsApiPostCorporationsCorporationIdAssetsNamesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datasource** | **optional.**| The server name you would like data from | [default to tranquility]
 **token** | **optional.**| Access token to use if unable to set a header | 

### Return type

[**[]PostCorporationsCorporationIdAssetsNames200Ok**](post_corporations_corporation_id_assets_names_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

