# \VoiceApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Webhook**](VoiceApi.md#Webhook) | **Get** /voice/{key}/webhook/{app} | Process web hook
[**Webhook_0**](VoiceApi.md#Webhook_0) | **Post** /voice/{key}/webhook/{app} | Process web hook
[**Webhook_1**](VoiceApi.md#Webhook_1) | **Get** /v{version}/voice/{key}/webhook/{app} | Process web hook
[**Webhook_2**](VoiceApi.md#Webhook_2) | **Post** /v{version}/voice/{key}/webhook/{app} | Process web hook


# **Webhook**
> Webhook(ctx, key, app, optional)
Process web hook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**|  | 
  **app** | **string**|  | 
 **optional** | ***VoiceApiWebhookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VoiceApiWebhookOpts struct

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

# **Webhook_0**
> Webhook_0(ctx, key, app, optional)
Process web hook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**|  | 
  **app** | **string**|  | 
 **optional** | ***VoiceApiWebhook_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VoiceApiWebhook_1Opts struct

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

# **Webhook_1**
> Webhook_1(ctx, key, app, version, optional)
Process web hook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**|  | 
  **app** | **string**|  | 
  **version** | **string**|  | 
 **optional** | ***VoiceApiWebhook_2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VoiceApiWebhook_2Opts struct

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

# **Webhook_2**
> Webhook_2(ctx, key, app, version, optional)
Process web hook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**|  | 
  **app** | **string**|  | 
  **version** | **string**|  | 
 **optional** | ***VoiceApiWebhook_3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VoiceApiWebhook_3Opts struct

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

