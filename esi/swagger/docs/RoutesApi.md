# {{classname}}

All URIs are relative to *https://esi.evetech.net/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRouteOriginDestination**](RoutesApi.md#GetRouteOriginDestination) | **Get** /route/{origin}/{destination}/ | Get route

# **GetRouteOriginDestination**
> []int32 GetRouteOriginDestination(ctx, destination, origin, optional)
Get route

Get the systems between origin and destination  --- Alternate route: `/dev/route/{origin}/{destination}/`  Alternate route: `/legacy/route/{origin}/{destination}/`  Alternate route: `/v1/route/{origin}/{destination}/`  --- This route is cached for up to 86400 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destination** | **int32**| destination solar system ID | 
  **origin** | **int32**| origin solar system ID | 
 **optional** | ***RoutesApiGetRouteOriginDestinationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoutesApiGetRouteOriginDestinationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **avoid** | [**optional.Interface of []int32**](int32.md)| avoid solar system ID(s) | 
 **connections** | [**optional.Interface of [][]int32**]([]int32.md)| connected solar system pairs | 
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **flag** | **optional.String**| route security preference | [default to shortest]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 

### Return type

**[]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

