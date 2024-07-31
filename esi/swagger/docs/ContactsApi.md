# {{classname}}

All URIs are relative to *https://esi.evetech.net/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCharactersCharacterIdContacts**](ContactsApi.md#DeleteCharactersCharacterIdContacts) | **Delete** /characters/{character_id}/contacts/ | Delete contacts
[**GetAlliancesAllianceIdContacts**](ContactsApi.md#GetAlliancesAllianceIdContacts) | **Get** /alliances/{alliance_id}/contacts/ | Get alliance contacts
[**GetAlliancesAllianceIdContactsLabels**](ContactsApi.md#GetAlliancesAllianceIdContactsLabels) | **Get** /alliances/{alliance_id}/contacts/labels/ | Get alliance contact labels
[**GetCharactersCharacterIdContacts**](ContactsApi.md#GetCharactersCharacterIdContacts) | **Get** /characters/{character_id}/contacts/ | Get contacts
[**GetCharactersCharacterIdContactsLabels**](ContactsApi.md#GetCharactersCharacterIdContactsLabels) | **Get** /characters/{character_id}/contacts/labels/ | Get contact labels
[**GetCorporationsCorporationIdContacts**](ContactsApi.md#GetCorporationsCorporationIdContacts) | **Get** /corporations/{corporation_id}/contacts/ | Get corporation contacts
[**GetCorporationsCorporationIdContactsLabels**](ContactsApi.md#GetCorporationsCorporationIdContactsLabels) | **Get** /corporations/{corporation_id}/contacts/labels/ | Get corporation contact labels
[**PostCharactersCharacterIdContacts**](ContactsApi.md#PostCharactersCharacterIdContacts) | **Post** /characters/{character_id}/contacts/ | Add contacts
[**PutCharactersCharacterIdContacts**](ContactsApi.md#PutCharactersCharacterIdContacts) | **Put** /characters/{character_id}/contacts/ | Edit contacts

# **DeleteCharactersCharacterIdContacts**
> DeleteCharactersCharacterIdContacts(ctx, characterId, contactIds, optional)
Delete contacts

Bulk delete contacts  --- Alternate route: `/dev/characters/{character_id}/contacts/`  Alternate route: `/v2/characters/{character_id}/contacts/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
  **contactIds** | [**[]int32**](int32.md)| A list of contacts to delete | 
 **optional** | ***ContactsApiDeleteCharactersCharacterIdContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContactsApiDeleteCharactersCharacterIdContactsOpts struct
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

# **GetAlliancesAllianceIdContacts**
> []GetAlliancesAllianceIdContacts200Ok GetAlliancesAllianceIdContacts(ctx, allianceId, optional)
Get alliance contacts

Return contacts of an alliance  --- Alternate route: `/dev/alliances/{alliance_id}/contacts/`  Alternate route: `/v2/alliances/{alliance_id}/contacts/`  --- This route is cached for up to 300 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **allianceId** | **int32**| An EVE alliance ID | 
 **optional** | ***ContactsApiGetAlliancesAllianceIdContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContactsApiGetAlliancesAllianceIdContactsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **optional.Int32**| Which page of results to return | [default to 1]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetAlliancesAllianceIdContacts200Ok**](get_alliances_alliance_id_contacts_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlliancesAllianceIdContactsLabels**
> []GetAlliancesAllianceIdContactsLabels200Ok GetAlliancesAllianceIdContactsLabels(ctx, allianceId, optional)
Get alliance contact labels

Return custom labels for an alliance's contacts  --- Alternate route: `/dev/alliances/{alliance_id}/contacts/labels/`  Alternate route: `/legacy/alliances/{alliance_id}/contacts/labels/`  Alternate route: `/v1/alliances/{alliance_id}/contacts/labels/`  --- This route is cached for up to 300 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **allianceId** | **int32**| An EVE alliance ID | 
 **optional** | ***ContactsApiGetAlliancesAllianceIdContactsLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContactsApiGetAlliancesAllianceIdContactsLabelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetAlliancesAllianceIdContactsLabels200Ok**](get_alliances_alliance_id_contacts_labels_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdContacts**
> []GetCharactersCharacterIdContacts200Ok GetCharactersCharacterIdContacts(ctx, characterId, optional)
Get contacts

Return contacts of a character  --- Alternate route: `/dev/characters/{character_id}/contacts/`  Alternate route: `/v2/characters/{character_id}/contacts/`  --- This route is cached for up to 300 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***ContactsApiGetCharactersCharacterIdContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContactsApiGetCharactersCharacterIdContactsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **optional.Int32**| Which page of results to return | [default to 1]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdContacts200Ok**](get_characters_character_id_contacts_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdContactsLabels**
> []GetCharactersCharacterIdContactsLabels200Ok GetCharactersCharacterIdContactsLabels(ctx, characterId, optional)
Get contact labels

Return custom labels for a character's contacts  --- Alternate route: `/dev/characters/{character_id}/contacts/labels/`  Alternate route: `/legacy/characters/{character_id}/contacts/labels/`  Alternate route: `/v1/characters/{character_id}/contacts/labels/`  --- This route is cached for up to 300 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***ContactsApiGetCharactersCharacterIdContactsLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContactsApiGetCharactersCharacterIdContactsLabelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdContactsLabels200Ok**](get_characters_character_id_contacts_labels_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdContacts**
> []GetCorporationsCorporationIdContacts200Ok GetCorporationsCorporationIdContacts(ctx, corporationId, optional)
Get corporation contacts

Return contacts of a corporation  --- Alternate route: `/dev/corporations/{corporation_id}/contacts/`  Alternate route: `/v2/corporations/{corporation_id}/contacts/`  --- This route is cached for up to 300 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | ***ContactsApiGetCorporationsCorporationIdContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContactsApiGetCorporationsCorporationIdContactsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **optional.Int32**| Which page of results to return | [default to 1]
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCorporationsCorporationIdContacts200Ok**](get_corporations_corporation_id_contacts_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporationsCorporationIdContactsLabels**
> []GetCorporationsCorporationIdContactsLabels200Ok GetCorporationsCorporationIdContactsLabels(ctx, corporationId, optional)
Get corporation contact labels

Return custom labels for a corporation's contacts  --- Alternate route: `/dev/corporations/{corporation_id}/contacts/labels/`  Alternate route: `/legacy/corporations/{corporation_id}/contacts/labels/`  Alternate route: `/v1/corporations/{corporation_id}/contacts/labels/`  --- This route is cached for up to 300 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **corporationId** | **int32**| An EVE corporation ID | 
 **optional** | ***ContactsApiGetCorporationsCorporationIdContactsLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContactsApiGetCorporationsCorporationIdContactsLabelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCorporationsCorporationIdContactsLabels200Ok**](get_corporations_corporation_id_contacts_labels_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCharactersCharacterIdContacts**
> []int32 PostCharactersCharacterIdContacts(ctx, body, characterId, standing, optional)
Add contacts

Bulk add contacts with same settings  --- Alternate route: `/dev/characters/{character_id}/contacts/`  Alternate route: `/v2/characters/{character_id}/contacts/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]int32**](int32.md)| A list of contacts | 
  **characterId** | **int32**| An EVE character ID | 
  **standing** | **float32**| Standing for the contact | 
 **optional** | ***ContactsApiPostCharactersCharacterIdContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContactsApiPostCharactersCharacterIdContactsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **datasource** | **optional.**| The server name you would like data from | [default to tranquility]
 **labelIds** | [**optional.Interface of []int64**](int64.md)| Add custom labels to the new contact | 
 **token** | **optional.**| Access token to use if unable to set a header | 
 **watched** | **optional.**| Whether the contact should be watched, note this is only effective on characters | [default to false]

### Return type

**[]int32**

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCharactersCharacterIdContacts**
> PutCharactersCharacterIdContacts(ctx, body, characterId, standing, optional)
Edit contacts

Bulk edit contacts with same settings  --- Alternate route: `/dev/characters/{character_id}/contacts/`  Alternate route: `/v2/characters/{character_id}/contacts/` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]int32**](int32.md)| A list of contacts | 
  **characterId** | **int32**| An EVE character ID | 
  **standing** | **float32**| Standing for the contact | 
 **optional** | ***ContactsApiPutCharactersCharacterIdContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContactsApiPutCharactersCharacterIdContactsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **datasource** | **optional.**| The server name you would like data from | [default to tranquility]
 **labelIds** | [**optional.Interface of []int64**](int64.md)| Add custom labels to the contact | 
 **token** | **optional.**| Access token to use if unable to set a header | 
 **watched** | **optional.**| Whether the contact should be watched, note this is only effective on characters | [default to false]

### Return type

 (empty response body)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

