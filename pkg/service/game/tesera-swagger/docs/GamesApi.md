# \GamesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DelRatingByUser**](GamesApi.md#DelRatingByUser) | **Delete** /games/{alias}/ratings | Del game rating by user
[**DelRatingByUser_0**](GamesApi.md#DelRatingByUser_0) | **Delete** /v{version}/games/{alias}/ratings | Del game rating by user
[**Get**](GamesApi.md#Get) | **Get** /games | List of games
[**GetFiles**](GamesApi.md#GetFiles) | **Get** /games/{alias}/files | Get game&#39;s files
[**GetFiles_0**](GamesApi.md#GetFiles_0) | **Get** /v{version}/games/{alias}/files | Get game&#39;s files
[**GetLinks**](GamesApi.md#GetLinks) | **Get** /games/{alias}/links | Get game&#39;s links
[**GetLinks_0**](GamesApi.md#GetLinks_0) | **Get** /v{version}/games/{alias}/links | Get game&#39;s links
[**GetNews**](GamesApi.md#GetNews) | **Get** /games/{alias}/news | List of news
[**GetNews_0**](GamesApi.md#GetNews_0) | **Get** /v{version}/games/{alias}/news | List of news
[**GetOffers**](GamesApi.md#GetOffers) | **Get** /games/{alias}/offers | Get available offers from stores
[**GetOffers_0**](GamesApi.md#GetOffers_0) | **Get** /v{version}/games/{alias}/offers | Get available offers from stores
[**GetOwns**](GamesApi.md#GetOwns) | **Get** /games/{alias}/owns | Get users own game
[**GetOwns_0**](GamesApi.md#GetOwns_0) | **Get** /v{version}/games/{alias}/owns | Get users own game
[**GetPhotos**](GamesApi.md#GetPhotos) | **Get** /games/{alias}/photos | Get game&#39;s photos
[**GetPhotos_0**](GamesApi.md#GetPhotos_0) | **Get** /v{version}/games/{alias}/photos | Get game&#39;s photos
[**GetRatingByUser**](GamesApi.md#GetRatingByUser) | **Get** /games/{alias}/ratings/{userid} | Get game&#39;s ratings by user
[**GetRatingByUser_0**](GamesApi.md#GetRatingByUser_0) | **Get** /v{version}/games/{alias}/ratings/{userid} | Get game&#39;s ratings by user
[**GetRatings**](GamesApi.md#GetRatings) | **Get** /games/{alias}/ratings | Get game&#39;s ratings
[**GetRatings_0**](GamesApi.md#GetRatings_0) | **Get** /v{version}/games/{alias}/ratings | Get game&#39;s ratings
[**GetRelated**](GamesApi.md#GetRelated) | **Get** /games/{alias}/related/{type} | List of related games
[**GetRelated_0**](GamesApi.md#GetRelated_0) | **Get** /v{version}/games/{alias}/related/{type} | List of related games
[**GetSimilar**](GamesApi.md#GetSimilar) | **Get** /games/{alias}/similar | List of similar games
[**GetSimilar_0**](GamesApi.md#GetSimilar_0) | **Get** /v{version}/games/{alias}/similar | List of similar games
[**Get_0**](GamesApi.md#Get_0) | **Get** /v{version}/games | List of games
[**Get_1**](GamesApi.md#Get_1) | **Get** /games/{alias} | Get a specific game by alias
[**Get_2**](GamesApi.md#Get_2) | **Get** /v{version}/games/{alias} | Get a specific game by alias
[**Post**](GamesApi.md#Post) | **Post** /games/{alias}/ratings | Post new rating
[**PostRatingTesera**](GamesApi.md#PostRatingTesera) | **Post** /games/{alias}/ratingstesera | Post new rating T1 format, only support T1
[**PostRatingTesera_0**](GamesApi.md#PostRatingTesera_0) | **Post** /v{version}/games/{alias}/ratingstesera | Post new rating T1 format, only support T1
[**Post_0**](GamesApi.md#Post_0) | **Post** /v{version}/games/{alias}/ratings | Post new rating


# **DelRatingByUser**
> DelRatingByUser(ctx, alias, optional)
Del game rating by user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
 **optional** | ***GamesApiDelRatingByUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiDelRatingByUserOpts struct

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

# **DelRatingByUser_0**
> DelRatingByUser_0(ctx, alias, version, optional)
Del game rating by user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
  **version** | **string**|  | 
 **optional** | ***GamesApiDelRatingByUser_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiDelRatingByUser_1Opts struct

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

# **Get**
> []GameInfo Get(ctx, optional)
List of games

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GamesApiGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**|  | 
 **playersMin** | **optional.Int64**|  | 
 **playersMax** | **optional.Int64**|  | 
 **playTimeMin** | **optional.Int64**|  | 
 **playTimeMax** | **optional.Int64**|  | 
 **playersAgeMin** | **optional.Int64**|  | 
 **tags** | **optional.String**|  | 
 **emptyPicture** | **optional.Bool**|  | 
 **withAdditions** | **optional.Bool**|  | 
 **sort** | **optional.String**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]GameInfo**](GameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFiles**
> []FileInfo GetFiles(ctx, alias, optional)
Get game's files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
 **optional** | ***GamesApiGetFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetFilesOpts struct

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
> []FileInfo GetFiles_0(ctx, alias, version, optional)
Get game's files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
  **version** | **string**|  | 
 **optional** | ***GamesApiGetFiles_2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetFiles_2Opts struct

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

# **GetLinks**
> []LinkInfo GetLinks(ctx, alias, optional)
Get game's links

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
 **optional** | ***GamesApiGetLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetLinksOpts struct

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

[**[]LinkInfo**](LinkInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLinks_0**
> []LinkInfo GetLinks_0(ctx, alias, version, optional)
Get game's links

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
  **version** | **string**|  | 
 **optional** | ***GamesApiGetLinks_3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetLinks_3Opts struct

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

[**[]LinkInfo**](LinkInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNews**
> []NewsInfoShort GetNews(ctx, alias, optional)
List of news

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**|  | 
 **optional** | ***GamesApiGetNewsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetNewsOpts struct

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

[**[]NewsInfoShort**](NewsInfoShort.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNews_0**
> []NewsInfoShort GetNews_0(ctx, alias, version, optional)
List of news

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**|  | 
  **version** | **string**|  | 
 **optional** | ***GamesApiGetNews_4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetNews_4Opts struct

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

[**[]NewsInfoShort**](NewsInfoShort.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOffers**
> []GameOffer GetOffers(ctx, alias, optional)
Get available offers from stores

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
 **optional** | ***GamesApiGetOffersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetOffersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]GameOffer**](GameOffer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOffers_0**
> []GameOffer GetOffers_0(ctx, alias, version, optional)
Get available offers from stores

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
  **version** | **string**|  | 
 **optional** | ***GamesApiGetOffers_5Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetOffers_5Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]GameOffer**](GameOffer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOwns**
> []UserInfoOwner GetOwns(ctx, alias, optional)
Get users own game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
 **optional** | ***GamesApiGetOwnsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetOwnsOpts struct

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

[**[]UserInfoOwner**](UserInfoOwner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOwns_0**
> []UserInfoOwner GetOwns_0(ctx, alias, version, optional)
Get users own game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
  **version** | **string**|  | 
 **optional** | ***GamesApiGetOwns_6Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetOwns_6Opts struct

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

[**[]UserInfoOwner**](UserInfoOwner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPhotos**
> []PhotoInfo GetPhotos(ctx, alias, optional)
Get game's photos

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
 **optional** | ***GamesApiGetPhotosOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetPhotosOpts struct

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

[**[]PhotoInfo**](PhotoInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPhotos_0**
> []PhotoInfo GetPhotos_0(ctx, alias, version, optional)
Get game's photos

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
  **version** | **string**|  | 
 **optional** | ***GamesApiGetPhotos_7Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetPhotos_7Opts struct

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

[**[]PhotoInfo**](PhotoInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRatingByUser**
> GameInfoRating GetRatingByUser(ctx, alias, userId, optional)
Get game's ratings by user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
  **userId** | **int64**| User id | 
 **optional** | ***GamesApiGetRatingByUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetRatingByUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**GameInfoRating**](GameInfoRating.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRatingByUser_0**
> GameInfoRating GetRatingByUser_0(ctx, alias, userId, version, optional)
Get game's ratings by user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
  **userId** | **int64**| User id | 
  **version** | **string**|  | 
 **optional** | ***GamesApiGetRatingByUser_8Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetRatingByUser_8Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**GameInfoRating**](GameInfoRating.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRatings**
> []GameInfoRating GetRatings(ctx, alias, optional)
Get game's ratings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
 **optional** | ***GamesApiGetRatingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetRatingsOpts struct

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

[**[]GameInfoRating**](GameInfoRating.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRatings_0**
> []GameInfoRating GetRatings_0(ctx, alias, version, optional)
Get game's ratings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
  **version** | **string**|  | 
 **optional** | ***GamesApiGetRatings_9Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetRatings_9Opts struct

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

[**[]GameInfoRating**](GameInfoRating.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRelated**
> []GameInfoShort GetRelated(ctx, alias, type_, optional)
List of related games

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**|  | 
  **type_** | **string**|  | 
 **optional** | ***GamesApiGetRelatedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetRelatedOpts struct

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

[**[]GameInfoShort**](GameInfoShort.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRelated_0**
> []GameInfoShort GetRelated_0(ctx, alias, type_, version, optional)
List of related games

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**|  | 
  **type_** | **string**|  | 
  **version** | **string**|  | 
 **optional** | ***GamesApiGetRelated_10Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetRelated_10Opts struct

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

[**[]GameInfoShort**](GameInfoShort.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSimilar**
> []GameInfoShort GetSimilar(ctx, alias, optional)
List of similar games

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**|  | 
 **optional** | ***GamesApiGetSimilarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetSimilarOpts struct

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

[**[]GameInfoShort**](GameInfoShort.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSimilar_0**
> []GameInfoShort GetSimilar_0(ctx, alias, version, optional)
List of similar games

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**|  | 
  **version** | **string**|  | 
 **optional** | ***GamesApiGetSimilar_11Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGetSimilar_11Opts struct

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

[**[]GameInfoShort**](GameInfoShort.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_0**
> []GameInfo Get_0(ctx, version, optional)
List of games

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***GamesApiGet_12Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGet_12Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**|  | 
 **playersMin** | **optional.Int64**|  | 
 **playersMax** | **optional.Int64**|  | 
 **playTimeMin** | **optional.Int64**|  | 
 **playTimeMax** | **optional.Int64**|  | 
 **playersAgeMin** | **optional.Int64**|  | 
 **tags** | **optional.String**|  | 
 **emptyPicture** | **optional.Bool**|  | 
 **withAdditions** | **optional.Bool**|  | 
 **sort** | **optional.String**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]GameInfo**](GameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_1**
> GameInfoResponse Get_1(ctx, alias, optional)
Get a specific game by alias

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
 **optional** | ***GamesApiGet_13Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGet_13Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**GameInfoResponse**](GameInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_2**
> GameInfoResponse Get_2(ctx, alias, version, optional)
Get a specific game by alias

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
  **version** | **string**|  | 
 **optional** | ***GamesApiGet_14Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiGet_14Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**GameInfoResponse**](GameInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Post**
> Post(ctx, alias, optional)
Post new rating

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**|  | 
 **optional** | ***GamesApiPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiPostOpts struct

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

# **PostRatingTesera**
> PostRatingTesera(ctx, alias, optional)
Post new rating T1 format, only support T1

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**|  | 
 **optional** | ***GamesApiPostRatingTeseraOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiPostRatingTeseraOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of PostRatingTeseraModel**](PostRatingTeseraModel.md)|  | 
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

# **PostRatingTesera_0**
> PostRatingTesera_0(ctx, alias, version, optional)
Post new rating T1 format, only support T1

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**|  | 
  **version** | **string**|  | 
 **optional** | ***GamesApiPostRatingTesera_15Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiPostRatingTesera_15Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of PostRatingTeseraModel**](PostRatingTeseraModel.md)|  | 
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
> Post_0(ctx, alias, version, optional)
Post new rating

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**|  | 
  **version** | **string**|  | 
 **optional** | ***GamesApiPost_16Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GamesApiPost_16Opts struct

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

