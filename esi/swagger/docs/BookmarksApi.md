# {{classname}}

All URIs are relative to *https://esi.evetech.net/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdBookmarks**](BookmarksApi.md#GetCharactersCharacterIdBookmarks) | **Get** /characters/{character_id}/bookmarks/ | List bookmarks
[**GetCharactersCharacterIdBookmarksFolders**](BookmarksApi.md#GetCharactersCharacterIdBookmarksFolders) | **Get** /characters/{character_id}/bookmarks/folders/ | List bookmark folders
[**GetCorporationsCorporationIdBookmarks**](BookmarksApi.md#GetCorporationsCorporationIdBookmarks) | **Get** /corporations/{corporation_id}/bookmarks/ | List corporation bookmarks
[**GetCorporationsCorporationIdBookmarksFolders**](BookmarksApi.md#GetCorporationsCorporationIdBookmarksFolders) | **Get** /corporations/{corporation_id}/bookmarks/folders/ | List corporation bookmark folders

# **GetCharactersCharacterIdBookmarks**
> []GetCharactersCharacterIdBookmarks200Ok GetCharactersCharacterIdBookmarks(ctx, characterId, optional)
List bookmarks

A list of your character's personal bookmarks  --- Alternate route: `/dev/characters/{character_id}/bookmarks/`  Alternate route: `/v2/characters/{character_id}/bookmarks/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***BookmarksApiGetCharactersCharacterIdBookmarksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookmarksApiGetCharactersCharacterIdBookmarksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **optional.Int32**| Which page of results to return | [default to 1]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdBookmarks200Ok**](get_characters_character_id_bookmarks_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdBookmarksFolders**
> []GetCharactersCharacterIdBookmarksFolders200Ok GetCharactersCharacterIdBookmarksFolders(ctx, characterId, optional)
List bookmark folders

A list of your character's personal bookmark folders  --- Alternate route: `/dev/characters/{character_id}/bookmarks/folders/`  Alternate route: `/v2/characters/{character_id}/bookmarks/folders/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***BookmarksApiGetCharactersCharacterIdBookmarksFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookmarksApiGetCharactersCharacterIdBookmarksFoldersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **optional.Int32**| Which page of results to return | [default to 1]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdBookmarksFolders200Ok**](get_characters_character_id_bookmarks_folders_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdBookmarks**
> []GetCorporationsCorporationIdBookmarks200Ok GetCorporationsCorporationIdBookmarks(ctx, corporationId, optional)
List corporation bookmarks

A list of your corporation's bookmarks  --- Alternate route: `/dev/corporations/{corporation_id}/bookmarks/`  Alternate route: `/legacy/corporations/{corporation_id}/bookmarks/`  Alternate route: `/v1/corporations/{corporation_id}/bookmarks/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | ***BookmarksApiGetCorporationsCorporationIdBookmarksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookmarksApiGetCorporationsCorporationIdBookmarksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **optional.Int32**| Which page of results to return | [default to 1]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCorporationsCorporationIdBookmarks200Ok**](get_corporations_corporation_id_bookmarks_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdBookmarksFolders**
> []GetCorporationsCorporationIdBookmarksFolders200Ok GetCorporationsCorporationIdBookmarksFolders(ctx, corporationId, optional)
List corporation bookmark folders

A list of your corporation's bookmark folders  --- Alternate route: `/dev/corporations/{corporation_id}/bookmarks/folders/`  Alternate route: `/legacy/corporations/{corporation_id}/bookmarks/folders/`  Alternate route: `/v1/corporations/{corporation_id}/bookmarks/folders/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | ***BookmarksApiGetCorporationsCorporationIdBookmarksFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookmarksApiGetCorporationsCorporationIdBookmarksFoldersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **optional.Int32**| Which page of results to return | [default to 1]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCorporationsCorporationIdBookmarksFolders200Ok**](get_corporations_corporation_id_bookmarks_folders_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

