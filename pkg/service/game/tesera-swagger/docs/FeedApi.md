# \FeedApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckSubscribe**](FeedApi.md#CheckSubscribe) | **Get** /feed/subscribe/{objecttype}/{alias} | Check user subscribe for item object and alias
[**CheckSubscribe_0**](FeedApi.md#CheckSubscribe_0) | **Get** /v{version}/feed/subscribe/{objecttype}/{alias} | Check user subscribe for item object and alias
[**DrawEvents**](FeedApi.md#DrawEvents) | **Get** /feed/drawevents | List of draw event in user subscribes
[**DrawEvents_0**](FeedApi.md#DrawEvents_0) | **Get** /v{version}/feed/drawevents | List of draw event in user subscribes
[**Events**](FeedApi.md#Events) | **Get** /feed/events | List of event in user subscribes
[**Events_0**](FeedApi.md#Events_0) | **Get** /v{version}/feed/events | List of event in user subscribes
[**Subscribe**](FeedApi.md#Subscribe) | **Post** /feed/subscribe | User subscribe on feed
[**Subscribe_0**](FeedApi.md#Subscribe_0) | **Post** /v{version}/feed/subscribe | User subscribe on feed
[**Unsubscribe**](FeedApi.md#Unsubscribe) | **Delete** /feed/unsubscribe/{objecttype}/{alias} | User unsubscribe of feed
[**Unsubscribe_0**](FeedApi.md#Unsubscribe_0) | **Delete** /v{version}/feed/unsubscribe/{objecttype}/{alias} | User unsubscribe of feed
[**UserSubscribes**](FeedApi.md#UserSubscribes) | **Get** /feed | List of user subscribes
[**UserSubscribes_0**](FeedApi.md#UserSubscribes_0) | **Get** /v{version}/feed | List of user subscribes


# **CheckSubscribe**
> SubscribeInfo CheckSubscribe(ctx, objectType, alias, optional)
Check user subscribe for item object and alias

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **alias** | **string**|  | 
 **optional** | ***FeedApiCheckSubscribeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiCheckSubscribeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**SubscribeInfo**](SubscribeInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckSubscribe_0**
> SubscribeInfo CheckSubscribe_0(ctx, objectType, alias, version, optional)
Check user subscribe for item object and alias

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **alias** | **string**|  | 
  **version** | **string**|  | 
 **optional** | ***FeedApiCheckSubscribe_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiCheckSubscribe_1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**SubscribeInfo**](SubscribeInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DrawEvents**
> EventsDrawListResponse DrawEvents(ctx, optional)
List of draw event in user subscribes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FeedApiDrawEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiDrawEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscribeId** | **optional.Int64**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**EventsDrawListResponse**](EventsDrawListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DrawEvents_0**
> EventsDrawListResponse DrawEvents_0(ctx, version, optional)
List of draw event in user subscribes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***FeedApiDrawEvents_2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiDrawEvents_2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscribeId** | **optional.Int64**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**EventsDrawListResponse**](EventsDrawListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Events**
> []FeedEventInfo Events(ctx, optional)
List of event in user subscribes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FeedApiEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscribeId** | **optional.Int64**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]FeedEventInfo**](FeedEventInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Events_0**
> []FeedEventInfo Events_0(ctx, version, optional)
List of event in user subscribes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***FeedApiEvents_3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiEvents_3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscribeId** | **optional.Int64**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]FeedEventInfo**](FeedEventInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Subscribe**
> UserSubscribeItemInfo Subscribe(ctx, optional)
User subscribe on feed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FeedApiSubscribeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiSubscribeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | [**optional.Interface of FeedSubscribeModel**](FeedSubscribeModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**UserSubscribeItemInfo**](UserSubscribeItemInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Subscribe_0**
> UserSubscribeItemInfo Subscribe_0(ctx, version, optional)
User subscribe on feed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***FeedApiSubscribe_4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiSubscribe_4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of FeedSubscribeModel**](FeedSubscribeModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**UserSubscribeItemInfo**](UserSubscribeItemInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Unsubscribe**
> Unsubscribe(ctx, objectType, alias, optional)
User unsubscribe of feed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**| Object type | 
  **alias** | **string**| Object alias | 
 **optional** | ***FeedApiUnsubscribeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiUnsubscribeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Unsubscribe_0**
> Unsubscribe_0(ctx, objectType, alias, version, optional)
User unsubscribe of feed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**| Object type | 
  **alias** | **string**| Object alias | 
  **version** | **string**|  | 
 **optional** | ***FeedApiUnsubscribe_5Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiUnsubscribe_5Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserSubscribes**
> []UserSubscribeItemInfo UserSubscribes(ctx, optional)
List of user subscribes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FeedApiUserSubscribesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiUserSubscribesOpts struct

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

[**[]UserSubscribeItemInfo**](UserSubscribeItemInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserSubscribes_0**
> []UserSubscribeItemInfo UserSubscribes_0(ctx, version, optional)
List of user subscribes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***FeedApiUserSubscribes_6Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiUserSubscribes_6Opts struct

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

[**[]UserSubscribeItemInfo**](UserSubscribeItemInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

