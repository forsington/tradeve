# {{classname}}

All URIs are relative to *https://esi.evetech.net/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharactersCharacterIdCalendar**](CalendarApi.md#GetCharactersCharacterIdCalendar) | **Get** /characters/{character_id}/calendar/ | List calendar event summaries
[**GetCharactersCharacterIdCalendarEventId**](CalendarApi.md#GetCharactersCharacterIdCalendarEventId) | **Get** /characters/{character_id}/calendar/{event_id}/ | Get an event
[**GetCharactersCharacterIdCalendarEventIdAttendees**](CalendarApi.md#GetCharactersCharacterIdCalendarEventIdAttendees) | **Get** /characters/{character_id}/calendar/{event_id}/attendees/ | Get attendees
[**PutCharactersCharacterIdCalendarEventId**](CalendarApi.md#PutCharactersCharacterIdCalendarEventId) | **Put** /characters/{character_id}/calendar/{event_id}/ | Respond to an event

# **GetCharactersCharacterIdCalendar**
> []GetCharactersCharacterIdCalendar200Ok GetCharactersCharacterIdCalendar(ctx, characterId, optional)
List calendar event summaries

Get 50 event summaries from the calendar. If no from_event ID is given, the resource will return the next 50 chronological event summaries from now. If a from_event ID is specified, it will return the next 50 chronological event summaries from after that event  --- Alternate route: `/dev/characters/{character_id}/calendar/`  Alternate route: `/legacy/characters/{character_id}/calendar/`  Alternate route: `/v1/characters/{character_id}/calendar/`  Alternate route: `/v2/characters/{character_id}/calendar/`  --- This route is cached for up to 5 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
 **optional** | ***CalendarApiGetCharactersCharacterIdCalendarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CalendarApiGetCharactersCharacterIdCalendarOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **fromEvent** | **optional.Int32**| The event ID to retrieve events from | 
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdCalendar200Ok**](get_characters_character_id_calendar_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdCalendarEventId**
> GetCharactersCharacterIdCalendarEventIdOk GetCharactersCharacterIdCalendarEventId(ctx, characterId, eventId, optional)
Get an event

Get all the information for a specific event  --- Alternate route: `/dev/characters/{character_id}/calendar/{event_id}/`  Alternate route: `/legacy/characters/{character_id}/calendar/{event_id}/`  Alternate route: `/v3/characters/{character_id}/calendar/{event_id}/`  Alternate route: `/v4/characters/{character_id}/calendar/{event_id}/`  --- This route is cached for up to 5 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
  **eventId** | **int32**| The id of the event requested | 
 **optional** | ***CalendarApiGetCharactersCharacterIdCalendarEventIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CalendarApiGetCharactersCharacterIdCalendarEventIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**GetCharactersCharacterIdCalendarEventIdOk**](get_characters_character_id_calendar_event_id_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCharactersCharacterIdCalendarEventIdAttendees**
> []GetCharactersCharacterIdCalendarEventIdAttendees200Ok GetCharactersCharacterIdCalendarEventIdAttendees(ctx, characterId, eventId, optional)
Get attendees

Get all invited attendees for a given event  --- Alternate route: `/dev/characters/{character_id}/calendar/{event_id}/attendees/`  Alternate route: `/legacy/characters/{character_id}/calendar/{event_id}/attendees/`  Alternate route: `/v1/characters/{character_id}/calendar/{event_id}/attendees/`  Alternate route: `/v2/characters/{character_id}/calendar/{event_id}/attendees/`  --- This route is cached for up to 600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int32**| An EVE character ID | 
  **eventId** | **int32**| The id of the event requested | 
 **optional** | ***CalendarApiGetCharactersCharacterIdCalendarEventIdAttendeesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CalendarApiGetCharactersCharacterIdCalendarEventIdAttendeesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **token** | **optional.String**| Access token to use if unable to set a header | 

### Return type

[**[]GetCharactersCharacterIdCalendarEventIdAttendees200Ok**](get_characters_character_id_calendar_event_id_attendees_200_ok.md)

### Authorization

[evesso](../README.md#evesso)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutCharactersCharacterIdCalendarEventId**
> PutCharactersCharacterIdCalendarEventId(ctx, body, characterId, eventId, optional)
Respond to an event

Set your response status to an event  --- Alternate route: `/dev/characters/{character_id}/calendar/{event_id}/`  Alternate route: `/legacy/characters/{character_id}/calendar/{event_id}/`  Alternate route: `/v3/characters/{character_id}/calendar/{event_id}/`  Alternate route: `/v4/characters/{character_id}/calendar/{event_id}/`  --- This route is cached for up to 5 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutCharactersCharacterIdCalendarEventIdResponse**](PutCharactersCharacterIdCalendarEventIdResponse.md)| The response value to set, overriding current value | 
  **characterId** | **int32**| An EVE character ID | 
  **eventId** | **int32**| The ID of the event requested | 
 **optional** | ***CalendarApiPutCharactersCharacterIdCalendarEventIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CalendarApiPutCharactersCharacterIdCalendarEventIdOpts struct
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

