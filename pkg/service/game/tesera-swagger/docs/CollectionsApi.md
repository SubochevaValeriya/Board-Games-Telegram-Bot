# \CollectionsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGame**](CollectionsApi.md#AddGame) | **Post** /collections/custom/{id} | Add game to custom collection
[**AddGame_0**](CollectionsApi.md#AddGame_0) | **Post** /v{version}/collections/custom/{id} | Add game to custom collection
[**AddItem**](CollectionsApi.md#AddItem) | **Post** /collections/base/{collectiontype} | Add game to collection
[**AddItem_0**](CollectionsApi.md#AddItem_0) | **Post** /v{version}/collections/base/{collectiontype} | Add game to collection
[**BaseCollectionGameCount**](CollectionsApi.md#BaseCollectionGameCount) | **Get** /collections/base/{collectiontype}/gamecount | Base collection&#39;s game/addition count
[**BaseCollectionGameCountByAlias**](CollectionsApi.md#BaseCollectionGameCountByAlias) | **Get** /collections/base/{collectiontype}/{alias}/gamecount | Base collection&#39;s game/addition count by user alias
[**BaseCollectionGameCountByAlias_0**](CollectionsApi.md#BaseCollectionGameCountByAlias_0) | **Get** /v{version}/collections/base/{collectiontype}/{alias}/gamecount | Base collection&#39;s game/addition count by user alias
[**BaseCollectionGameCount_0**](CollectionsApi.md#BaseCollectionGameCount_0) | **Get** /v{version}/collections/base/{collectiontype}/gamecount | Base collection&#39;s game/addition count
[**Bgg**](CollectionsApi.md#Bgg) | **Get** /collections/bgg/{username} | Import bgg collection
[**Bgg_0**](CollectionsApi.md#Bgg_0) | **Get** /v{version}/collections/bgg/{username} | Import bgg collection
[**ClearGames**](CollectionsApi.md#ClearGames) | **Get** /collections/custom/{id}/gamesclear | Games in custom collection clear
[**ClearGames_0**](CollectionsApi.md#ClearGames_0) | **Get** /v{version}/collections/custom/{id}/gamesclear | Games in custom collection clear
[**Collections**](CollectionsApi.md#Collections) | **Get** /collections/custom | Custom collections
[**CollectionsByUserId**](CollectionsApi.md#CollectionsByUserId) | **Get** /collections/custom/user/{userid} | Custom collections by userId
[**CollectionsByUserId_0**](CollectionsApi.md#CollectionsByUserId_0) | **Get** /v{version}/collections/custom/user/{userid} | Custom collections by userId
[**Collections_0**](CollectionsApi.md#Collections_0) | **Get** /v{version}/collections/custom | Custom collections
[**CreateCollection**](CollectionsApi.md#CreateCollection) | **Post** /collections/custom | Create new custom collection
[**CreateCollection_0**](CollectionsApi.md#CreateCollection_0) | **Post** /v{version}/collections/custom | Create new custom collection
[**CustomCollectionGameCount**](CollectionsApi.md#CustomCollectionGameCount) | **Get** /collections/custom/{collectionid}/gamecount | Custom collection&#39;s game/addition count
[**CustomCollectionGameCount_0**](CollectionsApi.md#CustomCollectionGameCount_0) | **Get** /v{version}/collections/custom/{collectionid}/gamecount | Custom collection&#39;s game/addition count
[**DelItem**](CollectionsApi.md#DelItem) | **Delete** /collections/base/{collectiontype}/{alias} | Delete game from collection
[**DelItem_0**](CollectionsApi.md#DelItem_0) | **Delete** /v{version}/collections/base/{collectiontype}/{alias} | Delete game from collection
[**DeleteCollection**](CollectionsApi.md#DeleteCollection) | **Delete** /collections/custom/{id} | Delete custom collection
[**DeleteCollection_0**](CollectionsApi.md#DeleteCollection_0) | **Delete** /v{version}/collections/custom/{id} | Delete custom collection
[**DeleteGame**](CollectionsApi.md#DeleteGame) | **Delete** /collections/custom/{id}/{alias} | Delete game from custom collection
[**DeleteGame_0**](CollectionsApi.md#DeleteGame_0) | **Delete** /v{version}/collections/custom/{id}/{alias} | Delete game from custom collection
[**GameInCollections**](CollectionsApi.md#GameInCollections) | **Get** /collections/{gamealias} | Game in collections
[**GameInCollections_0**](CollectionsApi.md#GameInCollections_0) | **Get** /v{version}/collections/{gamealias} | Game in collections
[**Games**](CollectionsApi.md#Games) | **Get** /collections/custom/{id}/games | Games in custom collection
[**Games_0**](CollectionsApi.md#Games_0) | **Get** /v{version}/collections/custom/{id}/games | Games in custom collection
[**Get**](CollectionsApi.md#Get) | **Get** /collections/base/{collectiontype} | Games in collection
[**GetByUser**](CollectionsApi.md#GetByUser) | **Get** /collections/base/{collectiontype}/{alias} | Games in collection by user
[**GetByUser_0**](CollectionsApi.md#GetByUser_0) | **Get** /v{version}/collections/base/{collectiontype}/{alias} | Games in collection by user
[**GetCollections**](CollectionsApi.md#GetCollections) | **Get** /collections | Based collections list
[**GetCollectionsByLogin**](CollectionsApi.md#GetCollectionsByLogin) | **Get** /collections/user/{alias} | Based collections list by user login
[**GetCollectionsByLogin_0**](CollectionsApi.md#GetCollectionsByLogin_0) | **Get** /v{version}/collections/user/{alias} | Based collections list by user login
[**GetCollections_0**](CollectionsApi.md#GetCollections_0) | **Get** /v{version}/collections | Based collections list
[**GetCustom**](CollectionsApi.md#GetCustom) | **Get** /collections/custom/{id} | Custom collection info
[**GetCustom_0**](CollectionsApi.md#GetCustom_0) | **Get** /v{version}/collections/custom/{id} | Custom collection info
[**Get_0**](CollectionsApi.md#Get_0) | **Get** /v{version}/collections/base/{collectiontype} | Games in collection
[**TradeAddItem**](CollectionsApi.md#TradeAddItem) | **Post** /collections/trade/{collectiontype} | Add game to trade collection
[**TradeAddItem_0**](CollectionsApi.md#TradeAddItem_0) | **Post** /v{version}/collections/trade/{collectiontype} | Add game to trade collection
[**TradeCollectionGameCount**](CollectionsApi.md#TradeCollectionGameCount) | **Get** /collections/trade/{collectiontype}/gamecount | Trade collection&#39;s game/addition count
[**TradeCollectionGameCountByAlias**](CollectionsApi.md#TradeCollectionGameCountByAlias) | **Get** /collections/trade/{collectiontype}/{alias}/gamecount | Trade collection&#39;s game/addition count by user alias
[**TradeCollectionGameCountByAlias_0**](CollectionsApi.md#TradeCollectionGameCountByAlias_0) | **Get** /v{version}/collections/trade/{collectiontype}/{alias}/gamecount | Trade collection&#39;s game/addition count by user alias
[**TradeCollectionGameCount_0**](CollectionsApi.md#TradeCollectionGameCount_0) | **Get** /v{version}/collections/trade/{collectiontype}/gamecount | Trade collection&#39;s game/addition count
[**TradeDelItem**](CollectionsApi.md#TradeDelItem) | **Delete** /collections/trade/{collectiontype}/{itemid} | Delete game from trade collection
[**TradeDelItem_0**](CollectionsApi.md#TradeDelItem_0) | **Delete** /v{version}/collections/trade/{collectiontype}/{itemid} | Delete game from trade collection
[**TradeGet**](CollectionsApi.md#TradeGet) | **Get** /collections/trade/{collectiontype}/{userid} | Games in trade collection
[**TradeGet_0**](CollectionsApi.md#TradeGet_0) | **Get** /v{version}/collections/trade/{collectiontype}/{userid} | Games in trade collection
[**UpdItem**](CollectionsApi.md#UpdItem) | **Put** /collections/base/{collectiontype}/{alias} | Update game in collection
[**UpdItem_0**](CollectionsApi.md#UpdItem_0) | **Put** /v{version}/collections/base/{collectiontype}/{alias} | Update game in collection
[**UpdateCollection**](CollectionsApi.md#UpdateCollection) | **Put** /collections/custom/{id} | Update custom collection
[**UpdateCollection_0**](CollectionsApi.md#UpdateCollection_0) | **Put** /v{version}/collections/custom/{id} | Update custom collection
[**UpdateGame**](CollectionsApi.md#UpdateGame) | **Put** /collections/custom/{id}/{alias} | Update game in custom collection
[**UpdateGame_0**](CollectionsApi.md#UpdateGame_0) | **Put** /v{version}/collections/custom/{id}/{alias} | Update game in custom collection


# **AddGame**
> CustomCollectionGameInfo AddGame(ctx, id, optional)
Add game to custom collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Collection Id | 
 **optional** | ***CollectionsApiAddGameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiAddGameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of AddGameToCustomCollectionModel**](AddGameToCustomCollectionModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CustomCollectionGameInfo**](CustomCollectionGameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddGame_0**
> CustomCollectionGameInfo AddGame_0(ctx, id, version, optional)
Add game to custom collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Collection Id | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiAddGame_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiAddGame_1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of AddGameToCustomCollectionModel**](AddGameToCustomCollectionModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CustomCollectionGameInfo**](CustomCollectionGameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddItem**
> CollectionGameInfo AddItem(ctx, collectionType, optional)
Add game to collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (own, played, toplay, wishlist) | 
 **optional** | ***CollectionsApiAddItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiAddItemOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of AddGameToCollectionModel**](AddGameToCollectionModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CollectionGameInfo**](CollectionGameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddItem_0**
> CollectionGameInfo AddItem_0(ctx, collectionType, version, optional)
Add game to collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (own, played, toplay, wishlist) | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiAddItem_2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiAddItem_2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of AddGameToCollectionModel**](AddGameToCollectionModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CollectionGameInfo**](CollectionGameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BaseCollectionGameCount**
> CollectionGameCountResponse BaseCollectionGameCount(ctx, collectionType, optional)
Base collection's game/addition count

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (own, played, toplay, wishlist) | 
 **optional** | ***CollectionsApiBaseCollectionGameCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiBaseCollectionGameCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **playersMin** | **optional.Int64**|  | 
 **playersMax** | **optional.Int64**|  | 
 **playtimeMin** | **optional.Int64**|  | 
 **playtimeMax** | **optional.Int64**|  | 
 **playersAgeMin** | **optional.Int64**|  | 
 **typeTag** | **optional.Int64**|  | 
 **settingTag** | **optional.Int64**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CollectionGameCountResponse**](CollectionGameCountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BaseCollectionGameCountByAlias**
> CollectionGameCountResponse BaseCollectionGameCountByAlias(ctx, collectionType, alias, optional)
Base collection's game/addition count by user alias

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (own, played, toplay, wishlist) | 
  **alias** | **string**| User alias | 
 **optional** | ***CollectionsApiBaseCollectionGameCountByAliasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiBaseCollectionGameCountByAliasOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **playersMin** | **optional.Int64**|  | 
 **playersMax** | **optional.Int64**|  | 
 **playtimeMin** | **optional.Int64**|  | 
 **playtimeMax** | **optional.Int64**|  | 
 **playersAgeMin** | **optional.Int64**|  | 
 **typeTag** | **optional.Int64**|  | 
 **settingTag** | **optional.Int64**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CollectionGameCountResponse**](CollectionGameCountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BaseCollectionGameCountByAlias_0**
> CollectionGameCountResponse BaseCollectionGameCountByAlias_0(ctx, collectionType, alias, version, optional)
Base collection's game/addition count by user alias

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (own, played, toplay, wishlist) | 
  **alias** | **string**| User alias | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiBaseCollectionGameCountByAlias_3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiBaseCollectionGameCountByAlias_3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **playersMin** | **optional.Int64**|  | 
 **playersMax** | **optional.Int64**|  | 
 **playtimeMin** | **optional.Int64**|  | 
 **playtimeMax** | **optional.Int64**|  | 
 **playersAgeMin** | **optional.Int64**|  | 
 **typeTag** | **optional.Int64**|  | 
 **settingTag** | **optional.Int64**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CollectionGameCountResponse**](CollectionGameCountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BaseCollectionGameCount_0**
> CollectionGameCountResponse BaseCollectionGameCount_0(ctx, collectionType, version, optional)
Base collection's game/addition count

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (own, played, toplay, wishlist) | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiBaseCollectionGameCount_4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiBaseCollectionGameCount_4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **playersMin** | **optional.Int64**|  | 
 **playersMax** | **optional.Int64**|  | 
 **playtimeMin** | **optional.Int64**|  | 
 **playtimeMax** | **optional.Int64**|  | 
 **playersAgeMin** | **optional.Int64**|  | 
 **typeTag** | **optional.Int64**|  | 
 **settingTag** | **optional.Int64**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CollectionGameCountResponse**](CollectionGameCountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Bgg**
> BggCollectionImportResponse Bgg(ctx, username, optional)
Import bgg collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| UserName on bgg site | 
 **optional** | ***CollectionsApiBggOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiBggOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**BggCollectionImportResponse**](BggCollectionImportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Bgg_0**
> BggCollectionImportResponse Bgg_0(ctx, username, version, optional)
Import bgg collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| UserName on bgg site | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiBgg_5Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiBgg_5Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**BggCollectionImportResponse**](BggCollectionImportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearGames**
> []CustomCollectionGameInfo ClearGames(ctx, id, optional)
Games in custom collection clear

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Collection Id | 
 **optional** | ***CollectionsApiClearGamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiClearGamesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gamesType** | **optional.String**|  | 
 **sort** | **optional.String**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]CustomCollectionGameInfo**](CustomCollectionGameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearGames_0**
> []CustomCollectionGameInfo ClearGames_0(ctx, id, version, optional)
Games in custom collection clear

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Collection Id | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiClearGames_6Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiClearGames_6Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gamesType** | **optional.String**|  | 
 **sort** | **optional.String**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]CustomCollectionGameInfo**](CustomCollectionGameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Collections**
> []CustomCollectionInfo Collections(ctx, optional)
Custom collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionsApiCollectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCollectionsOpts struct

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

[**[]CustomCollectionInfo**](CustomCollectionInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionsByUserId**
> []CustomCollectionInfo CollectionsByUserId(ctx, optional)
Custom collections by userId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionsApiCollectionsByUserIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCollectionsByUserIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **optional.Int64**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]CustomCollectionInfo**](CustomCollectionInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectionsByUserId_0**
> []CustomCollectionInfo CollectionsByUserId_0(ctx, version, optional)
Custom collections by userId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***CollectionsApiCollectionsByUserId_7Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCollectionsByUserId_7Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **optional.Int64**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]CustomCollectionInfo**](CustomCollectionInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Collections_0**
> []CustomCollectionInfo Collections_0(ctx, version, optional)
Custom collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***CollectionsApiCollections_8Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCollections_8Opts struct

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

[**[]CustomCollectionInfo**](CustomCollectionInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCollection**
> CustomCollectionInfo CreateCollection(ctx, optional)
Create new custom collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionsApiCreateCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCreateCollectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | [**optional.Interface of CollectionModel**](CollectionModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CustomCollectionInfo**](CustomCollectionInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCollection_0**
> CustomCollectionInfo CreateCollection_0(ctx, version, optional)
Create new custom collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***CollectionsApiCreateCollection_9Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCreateCollection_9Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of CollectionModel**](CollectionModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CustomCollectionInfo**](CustomCollectionInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomCollectionGameCount**
> CollectionGameCountResponse CustomCollectionGameCount(ctx, collectionId, optional)
Custom collection's game/addition count

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Custom collection id | 
 **optional** | ***CollectionsApiCustomCollectionGameCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCustomCollectionGameCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CollectionGameCountResponse**](CollectionGameCountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomCollectionGameCount_0**
> CollectionGameCountResponse CustomCollectionGameCount_0(ctx, collectionId, version, optional)
Custom collection's game/addition count

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **int64**| Custom collection id | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiCustomCollectionGameCount_10Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiCustomCollectionGameCount_10Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CollectionGameCountResponse**](CollectionGameCountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DelItem**
> DelItem(ctx, collectionType, alias, optional)
Delete game from collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (own, played, toplay, wishlist) | 
  **alias** | **string**| Game alias | 
 **optional** | ***CollectionsApiDelItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiDelItemOpts struct

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

# **DelItem_0**
> DelItem_0(ctx, collectionType, alias, version, optional)
Delete game from collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (own, played, toplay, wishlist) | 
  **alias** | **string**| Game alias | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiDelItem_11Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiDelItem_11Opts struct

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

# **DeleteCollection**
> DeleteCollection(ctx, id, optional)
Delete custom collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Collection Id | 
 **optional** | ***CollectionsApiDeleteCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiDeleteCollectionOpts struct

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

# **DeleteCollection_0**
> DeleteCollection_0(ctx, id, version, optional)
Delete custom collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Collection Id | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiDeleteCollection_12Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiDeleteCollection_12Opts struct

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

# **DeleteGame**
> DeleteGame(ctx, id, alias, optional)
Delete game from custom collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Collection Id | 
  **alias** | **string**| Game alias | 
 **optional** | ***CollectionsApiDeleteGameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiDeleteGameOpts struct

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

# **DeleteGame_0**
> DeleteGame_0(ctx, id, alias, version, optional)
Delete game from custom collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Collection Id | 
  **alias** | **string**| Game alias | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiDeleteGame_13Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiDeleteGame_13Opts struct

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

# **GameInCollections**
> GameInCollectionsResponse GameInCollections(ctx, gameAlias, optional)
Game in collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameAlias** | **string**|  | 
 **optional** | ***CollectionsApiGameInCollectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiGameInCollectionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**GameInCollectionsResponse**](GameInCollectionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GameInCollections_0**
> GameInCollectionsResponse GameInCollections_0(ctx, gameAlias, version, optional)
Game in collections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameAlias** | **string**|  | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiGameInCollections_14Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiGameInCollections_14Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**GameInCollectionsResponse**](GameInCollectionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Games**
> CustomCollectionGamesResponse Games(ctx, id, optional)
Games in custom collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Collection Id | 
 **optional** | ***CollectionsApiGamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiGamesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gamesType** | **optional.String**|  | 
 **sort** | **optional.String**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CustomCollectionGamesResponse**](CustomCollectionGamesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Games_0**
> CustomCollectionGamesResponse Games_0(ctx, id, version, optional)
Games in custom collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Collection Id | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiGames_15Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiGames_15Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gamesType** | **optional.String**|  | 
 **sort** | **optional.String**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CustomCollectionGamesResponse**](CustomCollectionGamesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> []CollectionGameInfo Get(ctx, collectionType, optional)
Games in collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (own, played, toplay, wishlist) | 
 **optional** | ***CollectionsApiGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gamesType** | **optional.String**|  | 
 **sort** | **optional.String**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **playersMin** | **optional.Int64**|  | 
 **playersMax** | **optional.Int64**|  | 
 **playtimeMin** | **optional.Int64**|  | 
 **playtimeMax** | **optional.Int64**|  | 
 **playersAgeMin** | **optional.Int64**|  | 
 **typeTag** | **optional.Int64**|  | 
 **settingTag** | **optional.Int64**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]CollectionGameInfo**](CollectionGameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetByUser**
> []CollectionGameInfo GetByUser(ctx, collectionType, alias, optional)
Games in collection by user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (own, played, toplay, wishlist) | 
  **alias** | **string**| User alias (teseraId or login) | 
 **optional** | ***CollectionsApiGetByUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiGetByUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gamesType** | **optional.String**|  | 
 **sort** | **optional.String**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **playersMin** | **optional.Int64**|  | 
 **playersMax** | **optional.Int64**|  | 
 **playtimeMin** | **optional.Int64**|  | 
 **playtimeMax** | **optional.Int64**|  | 
 **playersAgeMin** | **optional.Int64**|  | 
 **typeTag** | **optional.Int64**|  | 
 **settingTag** | **optional.Int64**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]CollectionGameInfo**](CollectionGameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetByUser_0**
> []CollectionGameInfo GetByUser_0(ctx, collectionType, alias, version, optional)
Games in collection by user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (own, played, toplay, wishlist) | 
  **alias** | **string**| User alias (teseraId or login) | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiGetByUser_16Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiGetByUser_16Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **gamesType** | **optional.String**|  | 
 **sort** | **optional.String**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **playersMin** | **optional.Int64**|  | 
 **playersMax** | **optional.Int64**|  | 
 **playtimeMin** | **optional.Int64**|  | 
 **playtimeMax** | **optional.Int64**|  | 
 **playersAgeMin** | **optional.Int64**|  | 
 **typeTag** | **optional.Int64**|  | 
 **settingTag** | **optional.Int64**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]CollectionGameInfo**](CollectionGameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCollections**
> CollectionsInfoResponse GetCollections(ctx, optional)
Based collections list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectionsApiGetCollectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiGetCollectionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CollectionsInfoResponse**](CollectionsInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCollectionsByLogin**
> CollectionsInfoResponse GetCollectionsByLogin(ctx, alias, optional)
Based collections list by user login

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| User alias (teseraId or login) | 
 **optional** | ***CollectionsApiGetCollectionsByLoginOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiGetCollectionsByLoginOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CollectionsInfoResponse**](CollectionsInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCollectionsByLogin_0**
> CollectionsInfoResponse GetCollectionsByLogin_0(ctx, alias, version, optional)
Based collections list by user login

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| User alias (teseraId or login) | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiGetCollectionsByLogin_17Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiGetCollectionsByLogin_17Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CollectionsInfoResponse**](CollectionsInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCollections_0**
> CollectionsInfoResponse GetCollections_0(ctx, version, optional)
Based collections list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***CollectionsApiGetCollections_18Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiGetCollections_18Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CollectionsInfoResponse**](CollectionsInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustom**
> CustomCollectionInfo GetCustom(ctx, id, optional)
Custom collection info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Collection Id | 
 **optional** | ***CollectionsApiGetCustomOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiGetCustomOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CustomCollectionInfo**](CustomCollectionInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustom_0**
> CustomCollectionInfo GetCustom_0(ctx, id, version, optional)
Custom collection info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Collection Id | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiGetCustom_19Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiGetCustom_19Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CustomCollectionInfo**](CustomCollectionInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get_0**
> []CollectionGameInfo Get_0(ctx, collectionType, version, optional)
Games in collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (own, played, toplay, wishlist) | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiGet_20Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiGet_20Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gamesType** | **optional.String**|  | 
 **sort** | **optional.String**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **playersMin** | **optional.Int64**|  | 
 **playersMax** | **optional.Int64**|  | 
 **playtimeMin** | **optional.Int64**|  | 
 **playtimeMax** | **optional.Int64**|  | 
 **playersAgeMin** | **optional.Int64**|  | 
 **typeTag** | **optional.Int64**|  | 
 **settingTag** | **optional.Int64**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]CollectionGameInfo**](CollectionGameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TradeAddItem**
> TradeCollectionGameInfo TradeAddItem(ctx, collectionType, optional)
Add game to trade collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (sell, buy) | 
 **optional** | ***CollectionsApiTradeAddItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiTradeAddItemOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of AddAdvertModel**](AddAdvertModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**TradeCollectionGameInfo**](TradeCollectionGameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TradeAddItem_0**
> TradeCollectionGameInfo TradeAddItem_0(ctx, collectionType, version, optional)
Add game to trade collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (sell, buy) | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiTradeAddItem_21Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiTradeAddItem_21Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of AddAdvertModel**](AddAdvertModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**TradeCollectionGameInfo**](TradeCollectionGameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TradeCollectionGameCount**
> TradeCollectionGameCountResponse TradeCollectionGameCount(ctx, collectionType, optional)
Trade collection's game/addition count

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (sell, buy) | 
 **optional** | ***CollectionsApiTradeCollectionGameCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiTradeCollectionGameCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**TradeCollectionGameCountResponse**](TradeCollectionGameCountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TradeCollectionGameCountByAlias**
> TradeCollectionGameCountResponse TradeCollectionGameCountByAlias(ctx, collectionType, alias, optional)
Trade collection's game/addition count by user alias

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (sell, buy) | 
  **alias** | **string**| User alias | 
 **optional** | ***CollectionsApiTradeCollectionGameCountByAliasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiTradeCollectionGameCountByAliasOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**TradeCollectionGameCountResponse**](TradeCollectionGameCountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TradeCollectionGameCountByAlias_0**
> TradeCollectionGameCountResponse TradeCollectionGameCountByAlias_0(ctx, collectionType, alias, version, optional)
Trade collection's game/addition count by user alias

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (sell, buy) | 
  **alias** | **string**| User alias | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiTradeCollectionGameCountByAlias_22Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiTradeCollectionGameCountByAlias_22Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**TradeCollectionGameCountResponse**](TradeCollectionGameCountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TradeCollectionGameCount_0**
> TradeCollectionGameCountResponse TradeCollectionGameCount_0(ctx, collectionType, version, optional)
Trade collection's game/addition count

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (sell, buy) | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiTradeCollectionGameCount_23Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiTradeCollectionGameCount_23Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**TradeCollectionGameCountResponse**](TradeCollectionGameCountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TradeDelItem**
> TradeDelItem(ctx, collectionType, itemId, optional)
Delete game from trade collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (sell, buy) | 
  **itemId** | **int64**| Trade item id | 
 **optional** | ***CollectionsApiTradeDelItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiTradeDelItemOpts struct

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

# **TradeDelItem_0**
> TradeDelItem_0(ctx, collectionType, itemId, version, optional)
Delete game from trade collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (sell, buy) | 
  **itemId** | **int64**| Trade item id | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiTradeDelItem_24Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiTradeDelItem_24Opts struct

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

# **TradeGet**
> []TradeCollectionGameInfo TradeGet(ctx, collectionType, userId, optional)
Games in trade collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type list (sell, buy) | 
  **userId** | **int64**| User id | 
 **optional** | ***CollectionsApiTradeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiTradeGetOpts struct

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

[**[]TradeCollectionGameInfo**](TradeCollectionGameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TradeGet_0**
> []TradeCollectionGameInfo TradeGet_0(ctx, collectionType, userId, version, optional)
Games in trade collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type list (sell, buy) | 
  **userId** | **int64**| User id | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiTradeGet_25Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiTradeGet_25Opts struct

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

[**[]TradeCollectionGameInfo**](TradeCollectionGameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdItem**
> CollectionGameInfo UpdItem(ctx, collectionType, alias, optional)
Update game in collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (own, played, toplay, wishlist) | 
  **alias** | **string**| Game alias | 
 **optional** | ***CollectionsApiUpdItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiUpdItemOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of UpdateGameInCollectionModel**](UpdateGameInCollectionModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CollectionGameInfo**](CollectionGameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdItem_0**
> CollectionGameInfo UpdItem_0(ctx, collectionType, alias, version, optional)
Update game in collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionType** | **string**| Type (own, played, toplay, wishlist) | 
  **alias** | **string**| Game alias | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiUpdItem_26Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiUpdItem_26Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **model** | [**optional.Interface of UpdateGameInCollectionModel**](UpdateGameInCollectionModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CollectionGameInfo**](CollectionGameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCollection**
> CustomCollectionInfo UpdateCollection(ctx, id, optional)
Update custom collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Collection Id | 
 **optional** | ***CollectionsApiUpdateCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiUpdateCollectionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of CollectionModel**](CollectionModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CustomCollectionInfo**](CustomCollectionInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCollection_0**
> CustomCollectionInfo UpdateCollection_0(ctx, id, version, optional)
Update custom collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Collection Id | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiUpdateCollection_27Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiUpdateCollection_27Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of CollectionModel**](CollectionModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CustomCollectionInfo**](CustomCollectionInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGame**
> CustomCollectionGameInfo UpdateGame(ctx, id, alias, optional)
Update game in custom collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| List Id | 
  **alias** | **string**| Game alias | 
 **optional** | ***CollectionsApiUpdateGameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiUpdateGameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of UpdateGameInCustomCollectionModel**](UpdateGameInCustomCollectionModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CustomCollectionGameInfo**](CustomCollectionGameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGame_0**
> CustomCollectionGameInfo UpdateGame_0(ctx, id, alias, version, optional)
Update game in custom collection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| List Id | 
  **alias** | **string**| Game alias | 
  **version** | **string**|  | 
 **optional** | ***CollectionsApiUpdateGame_28Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionsApiUpdateGame_28Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **model** | [**optional.Interface of UpdateGameInCustomCollectionModel**](UpdateGameInCustomCollectionModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**CustomCollectionGameInfo**](CustomCollectionGameInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

