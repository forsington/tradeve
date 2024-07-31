# {{classname}}

All URIs are relative to *https://esi.evetech.net/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWars**](WarsApi.md#GetWars) | **Get** /wars/ | List wars
[**GetWarsWarId**](WarsApi.md#GetWarsWarId) | **Get** /wars/{war_id}/ | Get war information
[**GetWarsWarIdKillmails**](WarsApi.md#GetWarsWarIdKillmails) | **Get** /wars/{war_id}/killmails/ | List kills for a war

# **GetWars**
> []int32 GetWars(ctx, optional)
List wars

Return a list of wars  --- Alternate route: `/dev/wars/`  Alternate route: `/legacy/wars/`  Alternate route: `/v1/wars/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WarsApiGetWarsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WarsApiGetWarsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **maxWarId** | **optional.Int32**| Only return wars with ID smaller than this | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWarsWarId**
> GetWarsWarIdOk GetWarsWarId(ctx, warId, optional)
Get war information

Return details about a war  --- Alternate route: `/dev/wars/{war_id}/`  Alternate route: `/legacy/wars/{war_id}/`  Alternate route: `/v1/wars/{war_id}/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **warId** | **int32**| ID for a war | 
 **optional** | ***WarsApiGetWarsWarIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WarsApiGetWarsWarIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

[**GetWarsWarIdOk**](get_wars_war_id_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWarsWarIdKillmails**
> []GetWarsWarIdKillmails200Ok GetWarsWarIdKillmails(ctx, warId, optional)
List kills for a war

Return a list of kills related to a war  --- Alternate route: `/dev/wars/{war_id}/killmails/`  Alternate route: `/legacy/wars/{war_id}/killmails/`  Alternate route: `/v1/wars/{war_id}/killmails/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **warId** | **int32**| A valid war ID | 
 **optional** | ***WarsApiGetWarsWarIdKillmailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WarsApiGetWarsWarIdKillmailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **page** | **optional.Int32**| Which page of results to return | [default to 1]

### Return type

[**[]GetWarsWarIdKillmails200Ok**](get_wars_war_id_killmails_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

