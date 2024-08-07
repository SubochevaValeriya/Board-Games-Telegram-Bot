/*
 * Tesera API
 *
 * API for Tesera
 *
 * API version: v1
 * Contact:
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type CitiesApiService service

/*
CitiesApiService Get cities list
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CitiesApiListOpts - Optional Parameters:
     * @param "Order" (optional.String) -
     * @param "UserCity" (optional.Bool) -
     * @param "Sort" (optional.String) -
     * @param "Offset" (optional.Int32) -
     * @param "Limit" (optional.Int32) -
     * @param "IsCancellationRequested" (optional.Bool) -
     * @param "CanBeCanceled" (optional.Bool) -
     * @param "WaitHandleHandle" (optional.Interface of map[string]string) -
     * @param "WaitHandleSafeWaitHandleIsInvalid" (optional.Bool) -
     * @param "WaitHandleSafeWaitHandleIsClosed" (optional.Bool) -

@return IActionResult
*/

type CitiesApiListOpts struct {
	Order                             optional.String
	UserCity                          optional.Bool
	Sort                              optional.String
	Offset                            optional.Int32
	Limit                             optional.Int32
	IsCancellationRequested           optional.Bool
	CanBeCanceled                     optional.Bool
	WaitHandleHandle                  optional.Interface
	WaitHandleSafeWaitHandleIsInvalid optional.Bool
	WaitHandleSafeWaitHandleIsClosed  optional.Bool
}

func (a *CitiesApiService) List(ctx context.Context, localVarOptionals *CitiesApiListOpts) (IActionResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue IActionResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Order.IsSet() {
		localVarQueryParams.Add("Order", parameterToString(localVarOptionals.Order.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserCity.IsSet() {
		localVarQueryParams.Add("UserCity", parameterToString(localVarOptionals.UserCity.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("Sort", parameterToString(localVarOptionals.Sort.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("Offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("Limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IsCancellationRequested.IsSet() {
		localVarQueryParams.Add("IsCancellationRequested", parameterToString(localVarOptionals.IsCancellationRequested.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CanBeCanceled.IsSet() {
		localVarQueryParams.Add("CanBeCanceled", parameterToString(localVarOptionals.CanBeCanceled.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WaitHandleHandle.IsSet() {
		localVarQueryParams.Add("WaitHandle.Handle", parameterToString(localVarOptionals.WaitHandleHandle.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WaitHandleSafeWaitHandleIsInvalid.IsSet() {
		localVarQueryParams.Add("WaitHandle.SafeWaitHandle.IsInvalid", parameterToString(localVarOptionals.WaitHandleSafeWaitHandleIsInvalid.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WaitHandleSafeWaitHandleIsClosed.IsSet() {
		localVarQueryParams.Add("WaitHandle.SafeWaitHandle.IsClosed", parameterToString(localVarOptionals.WaitHandleSafeWaitHandleIsClosed.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v IActionResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
CitiesApiService Get cities list
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param version
 * @param optional nil or *CitiesApiList_1Opts - Optional Parameters:
     * @param "Order" (optional.String) -
     * @param "UserCity" (optional.Bool) -
     * @param "Sort" (optional.String) -
     * @param "Offset" (optional.Int32) -
     * @param "Limit" (optional.Int32) -
     * @param "IsCancellationRequested" (optional.Bool) -
     * @param "CanBeCanceled" (optional.Bool) -
     * @param "WaitHandleHandle" (optional.Interface of map[string]string) -
     * @param "WaitHandleSafeWaitHandleIsInvalid" (optional.Bool) -
     * @param "WaitHandleSafeWaitHandleIsClosed" (optional.Bool) -

@return IActionResult
*/

type CitiesApiList_1Opts struct {
	Order                             optional.String
	UserCity                          optional.Bool
	Sort                              optional.String
	Offset                            optional.Int32
	Limit                             optional.Int32
	IsCancellationRequested           optional.Bool
	CanBeCanceled                     optional.Bool
	WaitHandleHandle                  optional.Interface
	WaitHandleSafeWaitHandleIsInvalid optional.Bool
	WaitHandleSafeWaitHandleIsClosed  optional.Bool
}

func (a *CitiesApiService) List_1(ctx context.Context, version string, localVarOptionals *CitiesApiList_1Opts) (IActionResult, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue IActionResult
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v{version}/cities"
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", fmt.Sprintf("%v", version), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Order.IsSet() {
		localVarQueryParams.Add("Order", parameterToString(localVarOptionals.Order.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserCity.IsSet() {
		localVarQueryParams.Add("UserCity", parameterToString(localVarOptionals.UserCity.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("Sort", parameterToString(localVarOptionals.Sort.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("Offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("Limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IsCancellationRequested.IsSet() {
		localVarQueryParams.Add("IsCancellationRequested", parameterToString(localVarOptionals.IsCancellationRequested.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CanBeCanceled.IsSet() {
		localVarQueryParams.Add("CanBeCanceled", parameterToString(localVarOptionals.CanBeCanceled.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WaitHandleHandle.IsSet() {
		localVarQueryParams.Add("WaitHandle.Handle", parameterToString(localVarOptionals.WaitHandleHandle.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WaitHandleSafeWaitHandleIsInvalid.IsSet() {
		localVarQueryParams.Add("WaitHandle.SafeWaitHandle.IsInvalid", parameterToString(localVarOptionals.WaitHandleSafeWaitHandleIsInvalid.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.WaitHandleSafeWaitHandleIsClosed.IsSet() {
		localVarQueryParams.Add("WaitHandle.SafeWaitHandle.IsClosed", parameterToString(localVarOptionals.WaitHandleSafeWaitHandleIsClosed.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v IActionResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
