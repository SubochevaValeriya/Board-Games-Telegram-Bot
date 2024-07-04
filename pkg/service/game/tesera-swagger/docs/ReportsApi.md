# \ReportsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAddToReport**](ReportsApi.md#AddAddToReport) | **Post** /reports/{reportid}/additions | Add addition game to report
[**AddAddToReport_0**](ReportsApi.md#AddAddToReport_0) | **Post** /v{version}/reports/{reportid}/additions | Add addition game to report
[**AddFileToReport**](ReportsApi.md#AddFileToReport) | **Post** /reports/{reportid}/files | Add file to report
[**AddFileToReport_0**](ReportsApi.md#AddFileToReport_0) | **Post** /v{version}/reports/{reportid}/files | Add file to report
[**AddPlayersToReport**](ReportsApi.md#AddPlayersToReport) | **Post** /reports/{reportid}/players | Add player to report
[**AddPlayersToReport_0**](ReportsApi.md#AddPlayersToReport_0) | **Post** /v{version}/reports/{reportid}/players | Add player to report
[**AddPtsCateg**](ReportsApi.md#AddPtsCateg) | **Post** /reports/config/{gameid}/points | Add point categories to a game
[**AddPtsCateg_0**](ReportsApi.md#AddPtsCateg_0) | **Post** /v{version}/reports/config/{gameid}/points | Add point categories to a game
[**AddReport**](ReportsApi.md#AddReport) | **Post** /reports | Create report
[**AddReportPlace**](ReportsApi.md#AddReportPlace) | **Post** /reports/places | Create place for reports
[**AddReportPlace_0**](ReportsApi.md#AddReportPlace_0) | **Post** /v{version}/reports/places | Create place for reports
[**AddReportPlayer**](ReportsApi.md#AddReportPlayer) | **Post** /reports/players | 
[**AddReportPlayer_0**](ReportsApi.md#AddReportPlayer_0) | **Post** /v{version}/reports/players | 
[**AddReport_0**](ReportsApi.md#AddReport_0) | **Post** /v{version}/reports | Create report
[**Bgg**](ReportsApi.md#Bgg) | **Get** /reports/bgg/{username} | Import reports from bgg
[**BggNew**](ReportsApi.md#BggNew) | **Post** /reports/bgg | Import reports from bgg
[**BggNew_0**](ReportsApi.md#BggNew_0) | **Post** /v{version}/reports/bgg | Import reports from bgg
[**Bgg_0**](ReportsApi.md#Bgg_0) | **Get** /v{version}/reports/bgg/{username} | Import reports from bgg
[**DelAddFromReport**](ReportsApi.md#DelAddFromReport) | **Delete** /reports/{reportid}/additions/{gameid} | Del addition game from report by linkId
[**DelAddFromReport_0**](ReportsApi.md#DelAddFromReport_0) | **Delete** /v{version}/reports/{reportid}/additions/{gameid} | Del addition game from report by linkId
[**DelPlace**](ReportsApi.md#DelPlace) | **Delete** /reports/places/{placeid} | Del place
[**DelPlace_0**](ReportsApi.md#DelPlace_0) | **Delete** /v{version}/reports/places/{placeid} | Del place
[**DelPlayer**](ReportsApi.md#DelPlayer) | **Delete** /reports/players/{playerid} | Del custom player
[**DelPlayer_0**](ReportsApi.md#DelPlayer_0) | **Delete** /v{version}/reports/players/{playerid} | Del custom player
[**DelPtsCateg**](ReportsApi.md#DelPtsCateg) | **Delete** /reports/config/{gameid}/points/{recordid} | Del game point categories
[**DelPtsCateg_0**](ReportsApi.md#DelPtsCateg_0) | **Delete** /v{version}/reports/config/{gameid}/points/{recordid} | Del game point categories
[**DelReport**](ReportsApi.md#DelReport) | **Delete** /reports/{reportid} | Del report
[**DelReportComment**](ReportsApi.md#DelReportComment) | **Delete** /reports/{reportid}/comments/{commentid} | Delete comment from report
[**DelReportComment_0**](ReportsApi.md#DelReportComment_0) | **Delete** /v{version}/reports/{reportid}/comments/{commentid} | Delete comment from report
[**DelReportFileById**](ReportsApi.md#DelReportFileById) | **Delete** /reports/{reportid}/files/{fileid} | Del file from report
[**DelReportFileById_0**](ReportsApi.md#DelReportFileById_0) | **Delete** /v{version}/reports/{reportid}/files/{fileid} | Del file from report
[**DelReportPlayerByLink**](ReportsApi.md#DelReportPlayerByLink) | **Delete** /reports/{reportid}/players/{linkid} | Del player from report by linkId
[**DelReportPlayerByLink_0**](ReportsApi.md#DelReportPlayerByLink_0) | **Delete** /v{version}/reports/{reportid}/players/{linkid} | Del player from report by linkId
[**DelReportRate**](ReportsApi.md#DelReportRate) | **Delete** /reports/{reportid}/like | 
[**DelReportRate_0**](ReportsApi.md#DelReportRate_0) | **Delete** /v{version}/reports/{reportid}/like | 
[**DelReport_0**](ReportsApi.md#DelReport_0) | **Delete** /v{version}/reports/{reportid} | Del report
[**EditPlace**](ReportsApi.md#EditPlace) | **Put** /reports/places/{placeid} | Edit place
[**EditPlace_0**](ReportsApi.md#EditPlace_0) | **Put** /v{version}/reports/places/{placeid} | Edit place
[**EditReportComment**](ReportsApi.md#EditReportComment) | **Put** /reports/{reportid}/comments/{commentid} | Edit comment from report
[**EditReportComment_0**](ReportsApi.md#EditReportComment_0) | **Put** /v{version}/reports/{reportid}/comments/{commentid} | Edit comment from report
[**GetGameReportsFiles**](ReportsApi.md#GetGameReportsFiles) | **Get** /reports/game/{alias}/files | Get file list by game alias
[**GetGameReportsFiles_0**](ReportsApi.md#GetGameReportsFiles_0) | **Get** /v{version}/reports/game/{alias}/files | Get file list by game alias
[**GetReportById**](ReportsApi.md#GetReportById) | **Get** /reports/{reportid} | Get report by id
[**GetReportById_0**](ReportsApi.md#GetReportById_0) | **Get** /v{version}/reports/{reportid} | Get report by id
[**GetReportComments**](ReportsApi.md#GetReportComments) | **Get** /reports/{reportid}/comments | Get comments from report
[**GetReportComments_0**](ReportsApi.md#GetReportComments_0) | **Get** /v{version}/reports/{reportid}/comments | Get comments from report
[**GetReportFiles**](ReportsApi.md#GetReportFiles) | **Get** /reports/{reportid}/files | 
[**GetReportFiles_0**](ReportsApi.md#GetReportFiles_0) | **Get** /v{version}/reports/{reportid}/files | 
[**GetReportGameConfig**](ReportsApi.md#GetReportGameConfig) | **Get** /reports/config/{gameid} | Get game config
[**GetReportGameConfig_0**](ReportsApi.md#GetReportGameConfig_0) | **Get** /v{version}/reports/config/{gameid} | Get game config
[**GetReportLocations**](ReportsApi.md#GetReportLocations) | **Get** /reports/places | Search places for reports
[**GetReportLocations_0**](ReportsApi.md#GetReportLocations_0) | **Get** /v{version}/reports/places | Search places for reports
[**GetReportPlayers**](ReportsApi.md#GetReportPlayers) | **Get** /reports/{reportid}/players | Get players list by reportId
[**GetReportPlayersLast**](ReportsApi.md#GetReportPlayersLast) | **Get** /reports/players/last | 
[**GetReportPlayersLast_0**](ReportsApi.md#GetReportPlayersLast_0) | **Get** /v{version}/reports/players/last | 
[**GetReportPlayers_0**](ReportsApi.md#GetReportPlayers_0) | **Get** /v{version}/reports/{reportid}/players | Get players list by reportId
[**GetReportUserInfo**](ReportsApi.md#GetReportUserInfo) | **Get** /reports/user/{alias}/info | Get report user info
[**GetReportUserInfo_0**](ReportsApi.md#GetReportUserInfo_0) | **Get** /v{version}/reports/user/{alias}/info | Get report user info
[**GetReports**](ReportsApi.md#GetReports) | **Get** /reports | List of reports
[**GetReportsByUser**](ReportsApi.md#GetReportsByUser) | **Get** /reports/user/{alias} | List of reports by user login
[**GetReportsByUser_0**](ReportsApi.md#GetReportsByUser_0) | **Get** /v{version}/reports/user/{alias} | List of reports by user login
[**GetReports_0**](ReportsApi.md#GetReports_0) | **Get** /v{version}/reports | List of reports
[**PostReportComment**](ReportsApi.md#PostReportComment) | **Post** /reports/{reportid}/comments/{replytoid} | Add new comment to report
[**PostReportComment_0**](ReportsApi.md#PostReportComment_0) | **Post** /v{version}/reports/{reportid}/comments/{replytoid} | Add new comment to report
[**PostReportRate**](ReportsApi.md#PostReportRate) | **Post** /reports/{reportid}/like | 
[**PostReportRate_0**](ReportsApi.md#PostReportRate_0) | **Post** /v{version}/reports/{reportid}/like | 
[**SearchReportPlayers**](ReportsApi.md#SearchReportPlayers) | **Get** /reports/players | 
[**SearchReportPlayers_0**](ReportsApi.md#SearchReportPlayers_0) | **Get** /v{version}/reports/players | 
[**UpdPlayer**](ReportsApi.md#UpdPlayer) | **Put** /reports/players/{playerid} | Upd custom player
[**UpdPlayer_0**](ReportsApi.md#UpdPlayer_0) | **Put** /v{version}/reports/players/{playerid} | Upd custom player
[**UpdPlayersToReport**](ReportsApi.md#UpdPlayersToReport) | **Put** /reports/{reportid}/players/{linkid} | Upd player in report
[**UpdPlayersToReport_0**](ReportsApi.md#UpdPlayersToReport_0) | **Put** /v{version}/reports/{reportid}/players/{linkid} | Upd player in report
[**UpdPtsCateg**](ReportsApi.md#UpdPtsCateg) | **Put** /reports/config/{gameid}/points/{recordid} | Upd game point categories
[**UpdPtsCateg_0**](ReportsApi.md#UpdPtsCateg_0) | **Put** /v{version}/reports/config/{gameid}/points/{recordid} | Upd game point categories
[**UpdReport**](ReportsApi.md#UpdReport) | **Put** /reports/{reportid} | Upd report
[**UpdReportGameConfig**](ReportsApi.md#UpdReportGameConfig) | **Put** /reports/config/{gameid} | Add/upd game config
[**UpdReportGameConfig_0**](ReportsApi.md#UpdReportGameConfig_0) | **Put** /v{version}/reports/config/{gameid} | Add/upd game config
[**UpdReport_0**](ReportsApi.md#UpdReport_0) | **Put** /v{version}/reports/{reportid} | Upd report


# **AddAddToReport**
> ReportAdditionInfoResponse AddAddToReport(ctx, reportId, optional)
Add addition game to report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
 **optional** | ***ReportsApiAddAddToReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiAddAddToReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameId** | **optional.Int64**| Game Id | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportAdditionInfoResponse**](ReportAdditionInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddAddToReport_0**
> ReportAdditionInfoResponse AddAddToReport_0(ctx, reportId, version, optional)
Add addition game to report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiAddAddToReport_1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiAddAddToReport_1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gameId** | **optional.Int64**| Game Id | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportAdditionInfoResponse**](ReportAdditionInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddFileToReport**
> FileToReportResponse AddFileToReport(ctx, reportId, optional)
Add file to report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
 **optional** | ***ReportsApiAddFileToReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiAddFileToReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**|  | 
 **content** | **optional.Interface of *os.File**|  | 
 **type_** | **optional.Int64**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**FileToReportResponse**](FileToReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddFileToReport_0**
> FileToReportResponse AddFileToReport_0(ctx, reportId, version, optional)
Add file to report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiAddFileToReport_2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiAddFileToReport_2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **name** | **optional.String**|  | 
 **content** | **optional.Interface of *os.File**|  | 
 **type_** | **optional.Int64**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**FileToReportResponse**](FileToReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPlayersToReport**
> PlayerToReportResponse AddPlayersToReport(ctx, reportId, optional)
Add player to report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
 **optional** | ***ReportsApiAddPlayersToReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiAddPlayersToReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of AddPlayerToReportModel**](AddPlayerToReportModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**PlayerToReportResponse**](PlayerToReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPlayersToReport_0**
> PlayerToReportResponse AddPlayersToReport_0(ctx, reportId, version, optional)
Add player to report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiAddPlayersToReport_3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiAddPlayersToReport_3Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of AddPlayerToReportModel**](AddPlayerToReportModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**PlayerToReportResponse**](PlayerToReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPtsCateg**
> ReportPointsInfo AddPtsCateg(ctx, gameId, optional)
Add point categories to a game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **int64**| Game id | 
 **optional** | ***ReportsApiAddPtsCategOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiAddPtsCategOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of AddReportPointsModel**](AddReportPointsModel.md)| model | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportPointsInfo**](ReportPointsInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPtsCateg_0**
> ReportPointsInfo AddPtsCateg_0(ctx, gameId, version, optional)
Add point categories to a game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **int64**| Game id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiAddPtsCateg_4Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiAddPtsCateg_4Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of AddReportPointsModel**](AddReportPointsModel.md)| model | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportPointsInfo**](ReportPointsInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddReport**
> ReportInfoResponse AddReport(ctx, optional)
Create report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportsApiAddReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiAddReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | [**optional.Interface of AddReportModel**](AddReportModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportInfoResponse**](ReportInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddReportPlace**
> ReportPlaceInfoResponse AddReportPlace(ctx, optional)
Create place for reports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportsApiAddReportPlaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiAddReportPlaceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | [**optional.Interface of AddReportPlaceModel**](AddReportPlaceModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportPlaceInfoResponse**](ReportPlaceInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddReportPlace_0**
> ReportPlaceInfoResponse AddReportPlace_0(ctx, version, optional)
Create place for reports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***ReportsApiAddReportPlace_5Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiAddReportPlace_5Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of AddReportPlaceModel**](AddReportPlaceModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportPlaceInfoResponse**](ReportPlaceInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddReportPlayer**
> ReportPlayerInfoResponse AddReportPlayer(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportsApiAddReportPlayerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiAddReportPlayerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | [**optional.Interface of AddReportPlayerModel**](AddReportPlayerModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportPlayerInfoResponse**](ReportPlayerInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddReportPlayer_0**
> ReportPlayerInfoResponse AddReportPlayer_0(ctx, version, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***ReportsApiAddReportPlayer_6Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiAddReportPlayer_6Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of AddReportPlayerModel**](AddReportPlayerModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportPlayerInfoResponse**](ReportPlayerInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddReport_0**
> ReportInfoResponse AddReport_0(ctx, version, optional)
Create report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***ReportsApiAddReport_7Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiAddReport_7Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of AddReportModel**](AddReportModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportInfoResponse**](ReportInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Bgg**
> BggReportsImportResponse Bgg(ctx, username, optional)
Import reports from bgg

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| UserName on bgg site | 
 **optional** | ***ReportsApiBggOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiBggOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**BggReportsImportResponse**](BggReportsImportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BggNew**
> BggReportsImportResponse BggNew(ctx, optional)
Import reports from bgg

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportsApiBggNewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiBggNewOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | [**optional.Interface of ReportImportModel**](ReportImportModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**BggReportsImportResponse**](BggReportsImportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BggNew_0**
> BggReportsImportResponse BggNew_0(ctx, version, optional)
Import reports from bgg

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***ReportsApiBggNew_8Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiBggNew_8Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of ReportImportModel**](ReportImportModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**BggReportsImportResponse**](BggReportsImportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Bgg_0**
> BggReportsImportResponse Bgg_0(ctx, username, version, optional)
Import reports from bgg

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| UserName on bgg site | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiBgg_9Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiBgg_9Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**BggReportsImportResponse**](BggReportsImportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DelAddFromReport**
> DelAddFromReport(ctx, reportId, gameId, optional)
Del addition game from report by linkId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
  **gameId** | **int64**| Game Id | 
 **optional** | ***ReportsApiDelAddFromReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiDelAddFromReportOpts struct

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

# **DelAddFromReport_0**
> DelAddFromReport_0(ctx, reportId, gameId, version, optional)
Del addition game from report by linkId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
  **gameId** | **int64**| Game Id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiDelAddFromReport_10Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiDelAddFromReport_10Opts struct

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

# **DelPlace**
> DelPlace(ctx, placeId, optional)
Del place

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **placeId** | **int64**| Place Id | 
 **optional** | ***ReportsApiDelPlaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiDelPlaceOpts struct

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

# **DelPlace_0**
> DelPlace_0(ctx, placeId, version, optional)
Del place

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **placeId** | **int64**| Place Id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiDelPlace_11Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiDelPlace_11Opts struct

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

# **DelPlayer**
> DelPlayer(ctx, playerId, optional)
Del custom player

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playerId** | **int64**| Player Id | 
 **optional** | ***ReportsApiDelPlayerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiDelPlayerOpts struct

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

# **DelPlayer_0**
> DelPlayer_0(ctx, playerId, version, optional)
Del custom player

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playerId** | **int64**| Player Id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiDelPlayer_12Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiDelPlayer_12Opts struct

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

# **DelPtsCateg**
> DelPtsCateg(ctx, gameId, recordId, optional)
Del game point categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **int64**| Game id | 
  **recordId** | **int64**| Points category record id | 
 **optional** | ***ReportsApiDelPtsCategOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiDelPtsCategOpts struct

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

# **DelPtsCateg_0**
> DelPtsCateg_0(ctx, gameId, recordId, version, optional)
Del game point categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **int64**| Game id | 
  **recordId** | **int64**| Points category record id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiDelPtsCateg_13Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiDelPtsCateg_13Opts struct

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

# **DelReport**
> DelReport(ctx, reportId, optional)
Del report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
 **optional** | ***ReportsApiDelReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiDelReportOpts struct

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

# **DelReportComment**
> DelReportComment(ctx, reportId, commentId, optional)
Delete comment from report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report id | 
  **commentId** | **int64**| Comment id | 
 **optional** | ***ReportsApiDelReportCommentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiDelReportCommentOpts struct

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

# **DelReportComment_0**
> DelReportComment_0(ctx, reportId, commentId, version, optional)
Delete comment from report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report id | 
  **commentId** | **int64**| Comment id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiDelReportComment_14Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiDelReportComment_14Opts struct

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

# **DelReportFileById**
> DelReportFileById(ctx, reportId, fileId, optional)
Del file from report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
  **fileId** | **int64**| File Id | 
 **optional** | ***ReportsApiDelReportFileByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiDelReportFileByIdOpts struct

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
> DelReportFileById_0(ctx, reportId, fileId, version, optional)
Del file from report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
  **fileId** | **int64**| File Id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiDelReportFileById_15Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiDelReportFileById_15Opts struct

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

# **DelReportPlayerByLink**
> DelReportPlayerByLink(ctx, reportId, linkId, optional)
Del player from report by linkId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
  **linkId** | **int64**| Link Id | 
 **optional** | ***ReportsApiDelReportPlayerByLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiDelReportPlayerByLinkOpts struct

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

# **DelReportPlayerByLink_0**
> DelReportPlayerByLink_0(ctx, reportId, linkId, version, optional)
Del player from report by linkId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
  **linkId** | **int64**| Link Id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiDelReportPlayerByLink_16Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiDelReportPlayerByLink_16Opts struct

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

# **DelReportRate**
> ReportRatingResponse DelReportRate(ctx, reportId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**|  | 
 **optional** | ***ReportsApiDelReportRateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiDelReportRateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportRatingResponse**](ReportRatingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DelReportRate_0**
> ReportRatingResponse DelReportRate_0(ctx, reportId, version, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**|  | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiDelReportRate_17Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiDelReportRate_17Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportRatingResponse**](ReportRatingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DelReport_0**
> DelReport_0(ctx, reportId, version, optional)
Del report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiDelReport_18Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiDelReport_18Opts struct

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

# **EditPlace**
> ReportPlaceInfo EditPlace(ctx, placeId, optional)
Edit place

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **placeId** | **int64**| Place Id | 
 **optional** | ***ReportsApiEditPlaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiEditPlaceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of AddReportPlaceModel**](AddReportPlaceModel.md)| model | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportPlaceInfo**](ReportPlaceInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditPlace_0**
> ReportPlaceInfo EditPlace_0(ctx, placeId, version, optional)
Edit place

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **placeId** | **int64**| Place Id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiEditPlace_19Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiEditPlace_19Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of AddReportPlaceModel**](AddReportPlaceModel.md)| model | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportPlaceInfo**](ReportPlaceInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditReportComment**
> ReportCommentInfo EditReportComment(ctx, reportId, commentId, optional)
Edit comment from report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report id | 
  **commentId** | **int64**| Comment id | 
 **optional** | ***ReportsApiEditReportCommentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiEditReportCommentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of ReportCommentModel**](ReportCommentModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportCommentInfo**](ReportCommentInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditReportComment_0**
> ReportCommentInfo EditReportComment_0(ctx, reportId, commentId, version, optional)
Edit comment from report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report id | 
  **commentId** | **int64**| Comment id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiEditReportComment_20Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiEditReportComment_20Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **model** | [**optional.Interface of ReportCommentModel**](ReportCommentModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportCommentInfo**](ReportCommentInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGameReportsFiles**
> []FileToObjectInfo GetGameReportsFiles(ctx, alias, optional)
Get file list by game alias

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
 **optional** | ***ReportsApiGetGameReportsFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetGameReportsFilesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]FileToObjectInfo**](FileToObjectInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGameReportsFiles_0**
> []FileToObjectInfo GetGameReportsFiles_0(ctx, alias, version, optional)
Get file list by game alias

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| Game alias | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiGetGameReportsFiles_21Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetGameReportsFiles_21Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **sort** | **optional.String**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]FileToObjectInfo**](FileToObjectInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportById**
> ReportInfoResponse GetReportById(ctx, reportId, optional)
Get report by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report ID | 
 **optional** | ***ReportsApiGetReportByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportInfoResponse**](ReportInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportById_0**
> ReportInfoResponse GetReportById_0(ctx, reportId, version, optional)
Get report by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report ID | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiGetReportById_22Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportById_22Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportInfoResponse**](ReportInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportComments**
> ReportCommentsResponse GetReportComments(ctx, reportId, optional)
Get comments from report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report id | 
 **optional** | ***ReportsApiGetReportCommentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportCommentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportCommentsResponse**](ReportCommentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportComments_0**
> ReportCommentsResponse GetReportComments_0(ctx, reportId, version, optional)
Get comments from report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiGetReportComments_23Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportComments_23Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportCommentsResponse**](ReportCommentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportFiles**
> ObjectFilesResponse GetReportFiles(ctx, reportId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**|  | 
 **optional** | ***ReportsApiGetReportFilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportFilesOpts struct

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
> ObjectFilesResponse GetReportFiles_0(ctx, reportId, version, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**|  | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiGetReportFiles_24Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportFiles_24Opts struct

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

# **GetReportGameConfig**
> ReportConfig GetReportGameConfig(ctx, gameId, optional)
Get game config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **int64**| Game id | 
 **optional** | ***ReportsApiGetReportGameConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportGameConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportConfig**](ReportConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportGameConfig_0**
> ReportConfig GetReportGameConfig_0(ctx, gameId, version, optional)
Get game config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **int64**| Game id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiGetReportGameConfig_25Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportGameConfig_25Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportConfig**](ReportConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportLocations**
> []ReportPlaceInfo GetReportLocations(ctx, optional)
Search places for reports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportsApiGetReportLocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportLocationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]ReportPlaceInfo**](ReportPlaceInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportLocations_0**
> []ReportPlaceInfo GetReportLocations_0(ctx, version, optional)
Search places for reports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***ReportsApiGetReportLocations_26Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportLocations_26Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]ReportPlaceInfo**](ReportPlaceInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportPlayers**
> ReportPlayersResponse GetReportPlayers(ctx, reportId, optional)
Get players list by reportId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
 **optional** | ***ReportsApiGetReportPlayersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportPlayersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportPlayersResponse**](ReportPlayersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportPlayersLast**
> ReportPlayersGroupsResponse GetReportPlayersLast(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportsApiGetReportPlayersLastOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportPlayersLastOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportPlayersGroupsResponse**](ReportPlayersGroupsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportPlayersLast_0**
> ReportPlayersGroupsResponse GetReportPlayersLast_0(ctx, version, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***ReportsApiGetReportPlayersLast_27Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportPlayersLast_27Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportPlayersGroupsResponse**](ReportPlayersGroupsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportPlayers_0**
> ReportPlayersResponse GetReportPlayers_0(ctx, reportId, version, optional)
Get players list by reportId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiGetReportPlayers_28Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportPlayers_28Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportPlayersResponse**](ReportPlayersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportUserInfo**
> ReportUserInfo GetReportUserInfo(ctx, alias, optional)
Get report user info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| User alias | 
 **optional** | ***ReportsApiGetReportUserInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportUserInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportUserInfo**](ReportUserInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportUserInfo_0**
> ReportUserInfo GetReportUserInfo_0(ctx, alias, version, optional)
Get report user info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| User alias | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiGetReportUserInfo_29Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportUserInfo_29Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportUserInfo**](ReportUserInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReports**
> []ReportInfo GetReports(ctx, optional)
List of reports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportsApiGetReportsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewDraft** | **optional.Bool**|  | 
 **game** | **optional.String**|  | 
 **city** | **optional.String**|  | 
 **accessType** | **optional.Int64**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]ReportInfo**](ReportInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsByUser**
> []ReportInfo GetReportsByUser(ctx, alias, optional)
List of reports by user login

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| User alias (teseraId or login) | 
 **optional** | ***ReportsApiGetReportsByUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsByUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **viewDraft** | **optional.Bool**|  | 
 **game** | **optional.String**|  | 
 **city** | **optional.String**|  | 
 **accessType** | **optional.Int64**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]ReportInfo**](ReportInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportsByUser_0**
> []ReportInfo GetReportsByUser_0(ctx, alias, version, optional)
List of reports by user login

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | **string**| User alias (teseraId or login) | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiGetReportsByUser_30Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReportsByUser_30Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewDraft** | **optional.Bool**|  | 
 **game** | **optional.String**|  | 
 **city** | **optional.String**|  | 
 **accessType** | **optional.Int64**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]ReportInfo**](ReportInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReports_0**
> []ReportInfo GetReports_0(ctx, version, optional)
List of reports

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***ReportsApiGetReports_31Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiGetReports_31Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **viewDraft** | **optional.Bool**|  | 
 **game** | **optional.String**|  | 
 **city** | **optional.String**|  | 
 **accessType** | **optional.Int64**|  | 
 **offset** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**[]ReportInfo**](ReportInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostReportComment**
> ReportCommentInfo PostReportComment(ctx, reportId, optional)
Add new comment to report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report id | 
 **optional** | ***ReportsApiPostReportCommentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiPostReportCommentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of ReportCommentModel**](ReportCommentModel.md)|  | 
 **replyToId** | **optional.Int64**| Reply to comment id (optional) | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportCommentInfo**](ReportCommentInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostReportComment_0**
> ReportCommentInfo PostReportComment_0(ctx, reportId, version, optional)
Add new comment to report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiPostReportComment_32Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiPostReportComment_32Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of ReportCommentModel**](ReportCommentModel.md)|  | 
 **replyToId** | **optional.Int64**| Reply to comment id (optional) | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportCommentInfo**](ReportCommentInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostReportRate**
> ReportRatingResponse PostReportRate(ctx, reportId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**|  | 
 **optional** | ***ReportsApiPostReportRateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiPostReportRateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportRatingResponse**](ReportRatingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostReportRate_0**
> ReportRatingResponse PostReportRate_0(ctx, reportId, version, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**|  | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiPostReportRate_33Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiPostReportRate_33Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportRatingResponse**](ReportRatingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchReportPlayers**
> ReportPlayersGroupsResponse SearchReportPlayers(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReportsApiSearchReportPlayersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiSearchReportPlayersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportPlayersGroupsResponse**](ReportPlayersGroupsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchReportPlayers_0**
> ReportPlayersGroupsResponse SearchReportPlayers_0(ctx, version, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **version** | **string**|  | 
 **optional** | ***ReportsApiSearchReportPlayers_34Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiSearchReportPlayers_34Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportPlayersGroupsResponse**](ReportPlayersGroupsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdPlayer**
> ReportPlayerInfoResponse UpdPlayer(ctx, playerId, optional)
Upd custom player

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playerId** | **int64**| Player id | 
 **optional** | ***ReportsApiUpdPlayerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiUpdPlayerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of AddReportPlayerModel**](AddReportPlayerModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportPlayerInfoResponse**](ReportPlayerInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdPlayer_0**
> ReportPlayerInfoResponse UpdPlayer_0(ctx, playerId, version, optional)
Upd custom player

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playerId** | **int64**| Player id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiUpdPlayer_35Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiUpdPlayer_35Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of AddReportPlayerModel**](AddReportPlayerModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportPlayerInfoResponse**](ReportPlayerInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdPlayersToReport**
> PlayerToReportResponse UpdPlayersToReport(ctx, reportId, linkId, optional)
Upd player in report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
  **linkId** | **int64**| Link Id | 
 **optional** | ***ReportsApiUpdPlayersToReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiUpdPlayersToReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of AddPlayerToReportModel**](AddPlayerToReportModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**PlayerToReportResponse**](PlayerToReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdPlayersToReport_0**
> PlayerToReportResponse UpdPlayersToReport_0(ctx, reportId, linkId, version, optional)
Upd player in report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
  **linkId** | **int64**| Link Id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiUpdPlayersToReport_36Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiUpdPlayersToReport_36Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **model** | [**optional.Interface of AddPlayerToReportModel**](AddPlayerToReportModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**PlayerToReportResponse**](PlayerToReportResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdPtsCateg**
> ReportPointsInfo UpdPtsCateg(ctx, gameId, recordId, optional)
Upd game point categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **int64**| Game id | 
  **recordId** | **int64**| Points category record id | 
 **optional** | ***ReportsApiUpdPtsCategOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiUpdPtsCategOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of AddReportPointsModel**](AddReportPointsModel.md)| model | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportPointsInfo**](ReportPointsInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdPtsCateg_0**
> ReportPointsInfo UpdPtsCateg_0(ctx, gameId, recordId, version, optional)
Upd game point categories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **int64**| Game id | 
  **recordId** | **int64**| Points category record id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiUpdPtsCateg_37Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiUpdPtsCateg_37Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **model** | [**optional.Interface of AddReportPointsModel**](AddReportPointsModel.md)| model | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportPointsInfo**](ReportPointsInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdReport**
> ReportInfoResponse UpdReport(ctx, reportId, optional)
Upd report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
 **optional** | ***ReportsApiUpdReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiUpdReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of AddReportModel**](AddReportModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportInfoResponse**](ReportInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdReportGameConfig**
> ReportConfig UpdReportGameConfig(ctx, gameId, optional)
Add/upd game config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **int64**| Game id | 
 **optional** | ***ReportsApiUpdReportGameConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiUpdReportGameConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**optional.Interface of AddReportConfigModel**](AddReportConfigModel.md)| model | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportConfig**](ReportConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdReportGameConfig_0**
> ReportConfig UpdReportGameConfig_0(ctx, gameId, version, optional)
Add/upd game config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gameId** | **int64**| Game id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiUpdReportGameConfig_38Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiUpdReportGameConfig_38Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of AddReportConfigModel**](AddReportConfigModel.md)| model | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportConfig**](ReportConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdReport_0**
> ReportInfoResponse UpdReport_0(ctx, reportId, version, optional)
Upd report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reportId** | **int64**| Report Id | 
  **version** | **string**|  | 
 **optional** | ***ReportsApiUpdReport_39Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportsApiUpdReport_39Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | [**optional.Interface of AddReportModel**](AddReportModel.md)|  | 
 **isCancellationRequested** | **optional.Bool**|  | 
 **canBeCanceled** | **optional.Bool**|  | 
 **waitHandleHandle** | [**optional.Interface of map[string]string**](string.md)|  | 
 **waitHandleSafeWaitHandleIsInvalid** | **optional.Bool**|  | 
 **waitHandleSafeWaitHandleIsClosed** | **optional.Bool**|  | 

### Return type

[**ReportInfoResponse**](ReportInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json, application/xml, text/xml, application/_*+xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

