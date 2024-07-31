# {{classname}}

All URIs are relative to *https://esi.evetech.net/latest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInsurancePrices**](InsuranceApi.md#GetInsurancePrices) | **Get** /insurance/prices/ | List insurance levels

# **GetInsurancePrices**
> []GetInsurancePrices200Ok GetInsurancePrices(ctx, optional)
List insurance levels

Return available insurance levels for all ship types  --- Alternate route: `/dev/insurance/prices/`  Alternate route: `/legacy/insurance/prices/`  Alternate route: `/v1/insurance/prices/`  --- This route is cached for up to 3600 seconds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InsuranceApiGetInsurancePricesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsuranceApiGetInsurancePricesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **optional.String**| Language to use in the response | [default to en]
 **datasource** | **optional.String**| The server name you would like data from | [default to tranquility]
 **ifNoneMatch** | **optional.String**| ETag from a previous request. A 304 will be returned if this matches the current ETag | 
 **language** | **optional.String**| Language to use in the response, takes precedence over Accept-Language | [default to en]

### Return type

[**[]GetInsurancePrices200Ok**](get_insurance_prices_200_ok.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

