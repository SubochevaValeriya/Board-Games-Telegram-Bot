# \CommentsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](CommentsApi.md#Get) | **Get** /comments/{objecttype}/{alias}/{lastcommentid} | List of comments for specific object
[**Get_0**](CommentsApi.md#Get_0) | **Get** /v{version}/comments/{objecttype}/{alias}/{lastcommentid} | List of comments for specific object
[**Latest**](CommentsApi.md#Latest) | **Get** /comments/latest | List of latest comments
[**LatestCount**](CommentsApi.md#LatestCount) | **Get** /comments/latest/count | Get comments posted on last week count
[**LatestCount_0**](CommentsApi.md#LatestCount_0) | **Get** /v{version}/comments/latest/count | Get comments posted on last week count
[**Latest_0**](CommentsApi.md#Latest_0) | **Get** /v{version}/comments/latest | List of latest comments
[**Post**](CommentsApi.md#Post) | **Post** /comments/{objecttype}/{alias}/{replycommentid} | Post new comment or reply
[**Post_0**](CommentsApi.md#Post_0) | **Post** /v{version}/comments/{objecttype}/{alias}/{replycommentid} | Post new comment or reply


# **Get**
> []CommentInfoShort Get(ctx, objectType, alias, optional)
List of comments for specific object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **alias** | **string**|  | 
 **optional** | ***CommentsApiGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommentsApiGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lastCommentId** | **optional.Int64**|  | 
 **limit** | **optional.Int32**|  | [default to 5]
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]CommentInfoShort**](CommentInfoShort.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_0**
> []CommentInfoShort Get_0(ctx, objectType, alias, version, optional)
List of comments for specific object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **alias** | **string**|  | 
  **version** | **string**|  | 
 **optional** | ***CommentsApiGet_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommentsApiGet_1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lastCommentId** | **optional.Int64**|  | 
 **limit** | **optional.Int32**|  | [default to 5]
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]CommentInfoShort**](CommentInfoShort.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Latest**
> []CommentInfoShort Latest(ctx, optional)
List of latest comments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommentsApiLatestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommentsApiLatestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectType** | **optional.String**| Filter by object type | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]CommentInfoShort**](CommentInfoShort.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LatestCount**
> CommentsLatestCountResponse LatestCount(ctx, optional)
Get comments posted on last week count

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CommentsApiLatestCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommentsApiLatestCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CommentsLatestCountResponse**](CommentsLatestCountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LatestCount_0**
> CommentsLatestCountResponse LatestCount_0(ctx, version, optional)
Get comments posted on last week count

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***CommentsApiLatestCount_2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommentsApiLatestCount_2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CommentsLatestCountResponse**](CommentsLatestCountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Latest_0**
> []CommentInfoShort Latest_0(ctx, version, optional)
List of latest comments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***CommentsApiLatest_3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommentsApiLatest_3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectType** | **optional.String**| Filter by object type | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]CommentInfoShort**](CommentInfoShort.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post**
> Post(ctx, objectType, alias, optional)
Post new comment or reply

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **alias** | **string**|  | 
 **optional** | ***CommentsApiPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommentsApiPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of PostCommentModel**](PostCommentModel.md)|  | 
 **replyCommentId** | **optional.Int64**|  | 
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
Post new comment or reply

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**|  | 
  **alias** | **string**|  | 
  **version** | **string**|  | 
 **optional** | ***CommentsApiPost_4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommentsApiPost_4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **model** | [**optional.Interface of PostCommentModel**](PostCommentModel.md)|  | 
 **replyCommentId** | **optional.Int64**|  | 
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

