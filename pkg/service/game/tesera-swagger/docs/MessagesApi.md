# \MessagesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetByUser**](MessagesApi.md#GetByUser) | **Get** /messages/{сompanionteseraid} | List of messages by user
[**GetByUser_0**](MessagesApi.md#GetByUser_0) | **Get** /v{version}/messages/{сompanionteseraid} | List of messages by user
[**GetList**](MessagesApi.md#GetList) | **Get** /messages/{objecttype}/{objectid} | List of conversation
[**GetList_0**](MessagesApi.md#GetList_0) | **Get** /v{version}/messages/{objecttype}/{objectid} | List of conversation
[**GetUnread**](MessagesApi.md#GetUnread) | **Get** /messages/unread | Get total of unread messages
[**GetUnread_0**](MessagesApi.md#GetUnread_0) | **Get** /v{version}/messages/unread | Get total of unread messages
[**Post**](MessagesApi.md#Post) | **Post** /messages/{useridto}/{replymessageid} | Post new message or reply
[**Post_0**](MessagesApi.md#Post_0) | **Post** /v{version}/messages/{useridto}/{replymessageid} | Post new message or reply


# **GetByUser**
> MessagesByUserResponse GetByUser(ctx, ompanionTeseraId, optional)
List of messages by user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ompanionTeseraId** | **int64**| Сompanion TeseraId in converstion | 
 **optional** | ***MessagesApiGetByUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiGetByUserOpts struct

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

[**MessagesByUserResponse**](MessagesByUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetByUser_0**
> MessagesByUserResponse GetByUser_0(ctx, ompanionTeseraId, version, optional)
List of messages by user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ompanionTeseraId** | **int64**| Сompanion TeseraId in converstion | 
  **version** | **string**|  | 
 **optional** | ***MessagesApiGetByUser_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiGetByUser_1Opts struct

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

[**MessagesByUserResponse**](MessagesByUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetList**
> []MessageInfo GetList(ctx, optional)
List of conversation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagesApiGetListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiGetListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 
 **objectType** | **optional.String**|  | [default to unknown]
 **objectId** | **optional.Int64**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]MessageInfo**](MessageInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetList_0**
> []MessageInfo GetList_0(ctx, version, optional)
List of conversation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***MessagesApiGetList_2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiGetList_2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 
 **objectType** | **optional.String**|  | [default to unknown]
 **objectId** | **optional.Int64**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]MessageInfo**](MessageInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUnread**
> int64 GetUnread(ctx, optional)
Get total of unread messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagesApiGetUnreadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiGetUnreadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUnread_0**
> int64 GetUnread_0(ctx, version, optional)
Get total of unread messages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***MessagesApiGetUnread_3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiGetUnread_3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post**
> Post(ctx, userIdTo, optional)
Post new message or reply

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userIdTo** | **int64**|  | 
 **optional** | ***MessagesApiPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of PostMessageModel**](PostMessageModel.md)|  | 
 **replyMessageId** | **optional.Int64**|  | 
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
> Post_0(ctx, userIdTo, version, optional)
Post new message or reply

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userIdTo** | **int64**|  | 
  **version** | **string**|  | 
 **optional** | ***MessagesApiPost_4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagesApiPost_4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of PostMessageModel**](PostMessageModel.md)|  | 
 **replyMessageId** | **optional.Int64**|  | 
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

