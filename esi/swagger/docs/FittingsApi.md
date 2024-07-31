# {{classname}}

All URIs are relative to *https://esi.evetech.net/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCharactersCharacterIdFittingsFittingId**](FittingsApi.md#DeleteCharactersCharacterIdFittingsFittingId) | **Delete** /characters/{character_id}/fittings/{fitting_id}/ | Delete fitting
[**GetCharactersCharacterIdFittings**](FittingsApi.md#GetCharactersCharacterIdFittings) | **Get** /characters/{character_id}/fittings/ | Get fittings
[**PostCharactersCharacterIdFittings**](FittingsApi.md#PostCharactersCharacterIdFittings) | **Post** /characters/{character_id}/fittings/ | Create fitting

# **DeleteCharactersCharacterIdFittingsFittingId**
> DeleteCharactersCharacterIdFittingsFittingId(ctx, characterId, fittingId, optional)
Delete fitting

Delete a fitting from a character  --- Alternate route: `/dev/characters/{character_id}/fittings/{fitting_id}/`  Alternate route: `/legacy/characters/{character_id}/fittings/{fitting_id}/`  Alternate route: `/v1/characters/{character_id}/fittings/{fitting_id}/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
  **fittingId** | **int32**| ID for a fitting of this character | 
 **optional** | ***FittingsApiDeleteCharactersCharacterIdFittingsFittingIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FittingsApiDeleteCharactersCharacterIdFittingsFittingIdOpts struct
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

# **GetCharactersCharacterIdFittings**
> []GetCharactersCharacterIdFittings200Ok GetCharactersCharacterIdFittings(ctx, characterId, optional)
Get fittings

Return fittings of a character  --- Alternate route: `/dev/characters/{character_id}/fittings/`  Alternate route: `/v2/characters/{character_id}/fittings/`  --- This route is cached for up to 300 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***FittingsApiGetCharactersCharacterIdFittingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FittingsApiGetCharactersCharacterIdFittingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdFittings200Ok**](get_characters_character_id_fittings_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCharactersCharacterIdFittings**
> PostCharactersCharacterIdFittingsCreated PostCharactersCharacterIdFittings(ctx, body, characterId, optional)
Create fitting

Save a new fitting for a character  --- Alternate route: `/dev/characters/{character_id}/fittings/`  Alternate route: `/v2/characters/{character_id}/fittings/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostCharactersCharacterIdFittingsFitting**](PostCharactersCharacterIdFittingsFitting.md)| Details about the new fitting | 
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***FittingsApiPostCharactersCharacterIdFittingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FittingsApiPostCharactersCharacterIdFittingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datasource** | **optional.**| The server name you would like data from | [default to tranquility]
 **token** | **optional.**| Access token to use if unable to set a header | 

### Return type

[**PostCharactersCharacterIdFittingsCreated**](post_characters_character_id_fittings_created.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

