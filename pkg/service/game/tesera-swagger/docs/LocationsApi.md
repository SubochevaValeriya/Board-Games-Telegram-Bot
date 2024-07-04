# \LocationsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSale**](LocationsApi.md#AddSale) | **Post** /locations/city | Add city
[**AddSale_0**](LocationsApi.md#AddSale_0) | **Post** /v{version}/locations/city | Add city
[**GetCityByAlias**](LocationsApi.md#GetCityByAlias) | **Get** /locations/city/{cityalias} | City location by alias
[**GetCityByAlias_0**](LocationsApi.md#GetCityByAlias_0) | **Get** /v{version}/locations/city/{cityalias} | City location by alias
[**GetCountryByAlias**](LocationsApi.md#GetCountryByAlias) | **Get** /locations/country/{countryalias} | Country location by alias
[**GetCountryByAlias_0**](LocationsApi.md#GetCountryByAlias_0) | **Get** /v{version}/locations/country/{countryalias} | Country location by alias


# **AddSale**
> CityResponse AddSale(ctx, optional)
Add city

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocationsApiAddSaleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationsApiAddSaleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | [**optional.Interface of AddCityModel**](AddCityModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CityResponse**](CityResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSale_0**
> CityResponse AddSale_0(ctx, version, optional)
Add city

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***LocationsApiAddSale_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationsApiAddSale_1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of AddCityModel**](AddCityModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CityResponse**](CityResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCityByAlias**
> CityInfo GetCityByAlias(ctx, cityAlias, optional)
City location by alias

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cityAlias** | **string**|  | 
 **optional** | ***LocationsApiGetCityByAliasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationsApiGetCityByAliasOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CityInfo**](CityInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCityByAlias_0**
> CityInfo GetCityByAlias_0(ctx, cityAlias, version, optional)
City location by alias

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cityAlias** | **string**|  | 
  **version** | **string**|  | 
 **optional** | ***LocationsApiGetCityByAlias_2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationsApiGetCityByAlias_2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CityInfo**](CityInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCountryByAlias**
> CountryInfo GetCountryByAlias(ctx, countryAlias, optional)
Country location by alias

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryAlias** | **string**|  | 
 **optional** | ***LocationsApiGetCountryByAliasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationsApiGetCountryByAliasOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CountryInfo**](CountryInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCountryByAlias_0**
> CountryInfo GetCountryByAlias_0(ctx, countryAlias, version, optional)
Country location by alias

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryAlias** | **string**|  | 
  **version** | **string**|  | 
 **optional** | ***LocationsApiGetCountryByAlias_3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationsApiGetCountryByAlias_3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CountryInfo**](CountryInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

