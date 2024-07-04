# \PublicationsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](PublicationsApi.md#Get) | **Get** /publications | List of publications
[**GetRatingByUser**](PublicationsApi.md#GetRatingByUser) | **Get** /publications/{objecttype}/{alias}/ratings/{userid} | Get publication rating by user
[**GetRatingByUser_0**](PublicationsApi.md#GetRatingByUser_0) | **Get** /v{version}/publications/{objecttype}/{alias}/ratings/{userid} | Get publication rating by user
[**GetRatings**](PublicationsApi.md#GetRatings) | **Get** /publications/{objecttype}/{alias}/ratings | Get ratings
[**GetRatings_0**](PublicationsApi.md#GetRatings_0) | **Get** /v{version}/publications/{objecttype}/{alias}/ratings | Get ratings
[**Get_0**](PublicationsApi.md#Get_0) | **Get** /v{version}/publications | List of publications
[**Post**](PublicationsApi.md#Post) | **Post** /publications/{objecttype}/{alias}/ratings | Post new rating
[**Post_0**](PublicationsApi.md#Post_0) | **Post** /v{version}/publications/{objecttype}/{alias}/ratings | Post new rating


# **Get**
> []PublicationInfo Get(ctx, optional)
List of publications

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PublicationsApiGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicationsApiGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diary** | **optional.Bool**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]PublicationInfo**](PublicationInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRatingByUser**
> PublicationInfoRating GetRatingByUser(ctx, objectType, alias, userId, optional)
Get publication rating by user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**| Object type publication | 
  **alias** | **string**| Publication alias | 
  **userId** | **int64**| User id | 
 **optional** | ***PublicationsApiGetRatingByUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicationsApiGetRatingByUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**PublicationInfoRating**](PublicationInfoRating.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRatingByUser_0**
> PublicationInfoRating GetRatingByUser_0(ctx, objectType, alias, userId, version, optional)
Get publication rating by user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**| Object type publication | 
  **alias** | **string**| Publication alias | 
  **userId** | **int64**| User id | 
  **version** | **string**|  | 
 **optional** | ***PublicationsApiGetRatingByUser_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicationsApiGetRatingByUser_1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**PublicationInfoRating**](PublicationInfoRating.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRatings**
> []PublicationInfoRating GetRatings(ctx, objectType, alias, optional)
Get ratings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**| Object type publication | 
  **alias** | **string**| Publication alias | 
 **optional** | ***PublicationsApiGetRatingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicationsApiGetRatingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]PublicationInfoRating**](PublicationInfoRating.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRatings_0**
> []PublicationInfoRating GetRatings_0(ctx, objectType, alias, version, optional)
Get ratings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**| Object type publication | 
  **alias** | **string**| Publication alias | 
  **version** | **string**|  | 
 **optional** | ***PublicationsApiGetRatings_2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicationsApiGetRatings_2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]PublicationInfoRating**](PublicationInfoRating.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_0**
> []PublicationInfo Get_0(ctx, version, optional)
List of publications

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***PublicationsApiGet_3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicationsApiGet_3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **diary** | **optional.Bool**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]PublicationInfo**](PublicationInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post**
> Post(ctx, objectType, alias, optional)
Post new rating

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**| Object type publication | 
  **alias** | **string**| Publication alias | 
 **optional** | ***PublicationsApiPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicationsApiPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of PostRatingModel**](PostRatingModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post_0**
> Post_0(ctx, objectType, alias, version, optional)
Post new rating

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**| Object type publication | 
  **alias** | **string**| Publication alias | 
  **version** | **string**|  | 
 **optional** | ***PublicationsApiPost_4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PublicationsApiPost_4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **model** | [**optional.Interface of PostRatingModel**](PostRatingModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

