# {{classname}}

All URIs are relative to *https://esi.evetech.net/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdSearch**](SearchApi.md#GetCharactersCharacterIdSearch) | **Get** /characters/{character_id}/search/ | Search on a string

# **GetCharactersCharacterIdSearch**
> GetCharactersCharacterIdSearchOk GetCharactersCharacterIdSearch(ctx, categories, characterId, search, optional)
Search on a string

Search for entities that match a given sub-string.  --- Alternate route: `/dev/characters/{character_id}/search/`  Alternate route: `/legacy/characters/{character_id}/search/`  Alternate route: `/v3/characters/{character_id}/search/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categories** | [**[]string**](string.md)| Type of entities to search for | 
  **characterId** | **int32**| An EVE character ID | 
  **search** | **string**| The string to search on | 
 **optional** | ***SearchApiGetCharactersCharacterIdSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiGetCharactersCharacterIdSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **acceptLanguage** | **optional.String**| Language to use in the response | [default to en]
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **optional.String**| Language to use in the response, takes precedence over Accept-Language | [default to en]
 **strict** | **optional.Bool**| Whether the search should be a strict match | [default to false]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**GetCharactersCharacterIdSearchOk**](get_characters_character_id_search_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

