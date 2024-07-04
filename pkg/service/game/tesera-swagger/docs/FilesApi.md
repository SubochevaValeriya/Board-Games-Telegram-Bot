# \FilesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFileToObject**](FilesApi.md#AddFileToObject) | **Post** /files/{objecttype}/{objectid} | Add file to object
[**AddFileToObject_0**](FilesApi.md#AddFileToObject_0) | **Post** /v{version}/files/{objecttype}/{objectid} | Add file to object
[**DelReportFileById**](FilesApi.md#DelReportFileById) | **Delete** /files/{objecttype}/{objectid}/{fileid} | Del file from report
[**DelReportFileById_0**](FilesApi.md#DelReportFileById_0) | **Delete** /v{version}/files/{objecttype}/{objectid}/{fileid} | Del file from report
[**GetFiles**](FilesApi.md#GetFiles) | **Get** /files | List of files
[**GetFiles_0**](FilesApi.md#GetFiles_0) | **Get** /v{version}/files | List of files
[**GetReportFiles**](FilesApi.md#GetReportFiles) | **Get** /files/{objecttype}/{objectid} | Get files list by reportId
[**GetReportFiles_0**](FilesApi.md#GetReportFiles_0) | **Get** /v{version}/files/{objecttype}/{objectid} | Get files list by reportId


# **AddFileToObject**
> FileToObjectResponse AddFileToObject(ctx, objectType, objectId, optional)
Add file to object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**| Object Type | 
  **objectId** | **int64**| Object Id | 
 **optional** | ***FilesApiAddFileToObjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilesApiAddFileToObjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **title** | **optional.String**|  | 
 **content** | **optional.Interface of *os.File**|  | 
 **type_** | **optional.Int64**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**FileToObjectResponse**](FileToObjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddFileToObject_0**
> FileToObjectResponse AddFileToObject_0(ctx, objectType, objectId, version, optional)
Add file to object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**| Object Type | 
  **objectId** | **int64**| Object Id | 
  **version** | **string**|  | 
 **optional** | ***FilesApiAddFileToObject_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilesApiAddFileToObject_1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **title** | **optional.String**|  | 
 **content** | **optional.Interface of *os.File**|  | 
 **type_** | **optional.Int64**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**FileToObjectResponse**](FileToObjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DelReportFileById**
> DelReportFileById(ctx, fileId, objectType, objectId, optional)
Del file from report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **int64**| File Id | 
  **objectType** | **string**| Object Type | 
  **objectId** | **int64**| Object Id | 
 **optional** | ***FilesApiDelReportFileByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilesApiDelReportFileByIdOpts struct

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

# **DelReportFileById_0**
> DelReportFileById_0(ctx, fileId, objectType, objectId, version, optional)
Del file from report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileId** | **int64**| File Id | 
  **objectType** | **string**| Object Type | 
  **objectId** | **int64**| Object Id | 
  **version** | **string**|  | 
 **optional** | ***FilesApiDelReportFileById_2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilesApiDelReportFileById_2Opts struct

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

# **GetFiles**
> []FileInfo GetFiles(ctx, optional)
List of files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FilesApiGetFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilesApiGetFilesOpts struct

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

[**[]FileInfo**](FileInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFiles_0**
> []FileInfo GetFiles_0(ctx, version, optional)
List of files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***FilesApiGetFiles_3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilesApiGetFiles_3Opts struct

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

[**[]FileInfo**](FileInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportFiles**
> ObjectFilesResponse GetReportFiles(ctx, objectType, objectId, optional)
Get files list by reportId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**| Object Type | 
  **objectId** | **int64**| Object Id | 
 **optional** | ***FilesApiGetReportFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilesApiGetReportFilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ObjectFilesResponse**](ObjectFilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportFiles_0**
> ObjectFilesResponse GetReportFiles_0(ctx, objectType, objectId, version, optional)
Get files list by reportId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**| Object Type | 
  **objectId** | **int64**| Object Id | 
  **version** | **string**|  | 
 **optional** | ***FilesApiGetReportFiles_4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilesApiGetReportFiles_4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ObjectFilesResponse**](ObjectFilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

