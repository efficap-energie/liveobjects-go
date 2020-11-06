/*
 * Live Objects REST API Guide v2.13.3
 *
 * API description for Live Objects service
 *
 * API version: 2.13.3
 * Contact: liveobjects.support@orange.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package liveobjects

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// UserAuthenticationApiService UserAuthenticationApi service
type UserAuthenticationApiService service

type ApiAuthenticateUserUsingPOSTRequest struct {
	ctx _context.Context
	ApiService *UserAuthenticationApiService
	cookie *bool
	authenticationRequest *AuthReqWeb
}

func (r ApiAuthenticateUserUsingPOSTRequest) Cookie(cookie bool) ApiAuthenticateUserUsingPOSTRequest {
	r.cookie = &cookie
	return r
}
func (r ApiAuthenticateUserUsingPOSTRequest) AuthenticationRequest(authenticationRequest AuthReqWeb) ApiAuthenticateUserUsingPOSTRequest {
	r.authenticationRequest = &authenticationRequest
	return r
}

func (r ApiAuthenticateUserUsingPOSTRequest) Execute() (AuthResWeb, *_nethttp.Response, error) {
	return r.ApiService.AuthenticateUserUsingPOSTExecute(r)
}

/*
 * AuthenticateUserUsingPOST Authenticate a user
 * Usage of this API will be reported in your access log under 'authentication' category.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiAuthenticateUserUsingPOSTRequest
 */
func (a *UserAuthenticationApiService) AuthenticateUserUsingPOST(ctx _context.Context) ApiAuthenticateUserUsingPOSTRequest {
	return ApiAuthenticateUserUsingPOSTRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return AuthResWeb
 */
func (a *UserAuthenticationApiService) AuthenticateUserUsingPOSTExecute(r ApiAuthenticateUserUsingPOSTRequest) (AuthResWeb, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AuthResWeb
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAuthenticationApiService.AuthenticateUserUsingPOST")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/auth"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.cookie != nil {
		localVarQueryParams.Add("cookie", parameterToString(*r.cookie, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.authenticationRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCookiesDeleteUsingDELETERequest struct {
	ctx _context.Context
	ApiService *UserAuthenticationApiService
	xAPIKEY *string
}

func (r ApiCookiesDeleteUsingDELETERequest) XAPIKEY(xAPIKEY string) ApiCookiesDeleteUsingDELETERequest {
	r.xAPIKEY = &xAPIKEY
	return r
}

func (r ApiCookiesDeleteUsingDELETERequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.CookiesDeleteUsingDELETEExecute(r)
}

/*
 * CookiesDeleteUsingDELETE cookiesDelete
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCookiesDeleteUsingDELETERequest
 */
func (a *UserAuthenticationApiService) CookiesDeleteUsingDELETE(ctx _context.Context) ApiCookiesDeleteUsingDELETERequest {
	return ApiCookiesDeleteUsingDELETERequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *UserAuthenticationApiService) CookiesDeleteUsingDELETEExecute(r ApiCookiesDeleteUsingDELETERequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAuthenticationApiService.CookiesDeleteUsingDELETE")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/cookies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return nil, reportError("xAPIKEY is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-API-KEY"] = parameterToString(*r.xAPIKEY, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetTenantIdUsingGETRequest struct {
	ctx _context.Context
	ApiService *UserAuthenticationApiService
	xAPIKEY *string
}

func (r ApiGetTenantIdUsingGETRequest) XAPIKEY(xAPIKEY string) ApiGetTenantIdUsingGETRequest {
	r.xAPIKEY = &xAPIKEY
	return r
}

func (r ApiGetTenantIdUsingGETRequest) Execute() (SimpleStringWeb, *_nethttp.Response, error) {
	return r.ApiService.GetTenantIdUsingGETExecute(r)
}

/*
 * GetTenantIdUsingGET Get your tenant id
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetTenantIdUsingGETRequest
 */
func (a *UserAuthenticationApiService) GetTenantIdUsingGET(ctx _context.Context) ApiGetTenantIdUsingGETRequest {
	return ApiGetTenantIdUsingGETRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return SimpleStringWeb
 */
func (a *UserAuthenticationApiService) GetTenantIdUsingGETExecute(r ApiGetTenantIdUsingGETRequest) (SimpleStringWeb, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SimpleStringWeb
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAuthenticationApiService.GetTenantIdUsingGET")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/whoami"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return localVarReturnValue, nil, reportError("xAPIKEY is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-API-KEY"] = parameterToString(*r.xAPIKEY, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiLogoutUsingPOSTRequest struct {
	ctx _context.Context
	ApiService *UserAuthenticationApiService
	xAPIKEY *string
	cookie *bool
}

func (r ApiLogoutUsingPOSTRequest) XAPIKEY(xAPIKEY string) ApiLogoutUsingPOSTRequest {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiLogoutUsingPOSTRequest) Cookie(cookie bool) ApiLogoutUsingPOSTRequest {
	r.cookie = &cookie
	return r
}

func (r ApiLogoutUsingPOSTRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.LogoutUsingPOSTExecute(r)
}

/*
 * LogoutUsingPOST Log out of the current session
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiLogoutUsingPOSTRequest
 */
func (a *UserAuthenticationApiService) LogoutUsingPOST(ctx _context.Context) ApiLogoutUsingPOSTRequest {
	return ApiLogoutUsingPOSTRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *UserAuthenticationApiService) LogoutUsingPOSTExecute(r ApiLogoutUsingPOSTRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAuthenticationApiService.LogoutUsingPOST")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/logout"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return nil, reportError("xAPIKEY is required and must be specified")
	}

	if r.cookie != nil {
		localVarQueryParams.Add("cookie", parameterToString(*r.cookie, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-API-KEY"] = parameterToString(*r.xAPIKEY, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiResetUserPasswordUsingPOSTRequest struct {
	ctx _context.Context
	ApiService *UserAuthenticationApiService
	userResetPasswordRequest *ResetPasswordReqWeb
}

func (r ApiResetUserPasswordUsingPOSTRequest) UserResetPasswordRequest(userResetPasswordRequest ResetPasswordReqWeb) ApiResetUserPasswordUsingPOSTRequest {
	r.userResetPasswordRequest = &userResetPasswordRequest
	return r
}

func (r ApiResetUserPasswordUsingPOSTRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ResetUserPasswordUsingPOSTExecute(r)
}

/*
 * ResetUserPasswordUsingPOST Reset user's password
 * Usage of this API will be reported in your access log under 'authentication' category.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiResetUserPasswordUsingPOSTRequest
 */
func (a *UserAuthenticationApiService) ResetUserPasswordUsingPOST(ctx _context.Context) ApiResetUserPasswordUsingPOSTRequest {
	return ApiResetUserPasswordUsingPOSTRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *UserAuthenticationApiService) ResetUserPasswordUsingPOSTExecute(r ApiResetUserPasswordUsingPOSTRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAuthenticationApiService.ResetUserPasswordUsingPOST")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/resetpwd"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.userResetPasswordRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateUserEmailWithTokenUsingPOSTRequest struct {
	ctx _context.Context
	ApiService *UserAuthenticationApiService
	xAPIKEY *string
	userUpdateEmailRequest *UpdateEmailReqWeb
}

func (r ApiUpdateUserEmailWithTokenUsingPOSTRequest) XAPIKEY(xAPIKEY string) ApiUpdateUserEmailWithTokenUsingPOSTRequest {
	r.xAPIKEY = &xAPIKEY
	return r
}
func (r ApiUpdateUserEmailWithTokenUsingPOSTRequest) UserUpdateEmailRequest(userUpdateEmailRequest UpdateEmailReqWeb) ApiUpdateUserEmailWithTokenUsingPOSTRequest {
	r.userUpdateEmailRequest = &userUpdateEmailRequest
	return r
}

func (r ApiUpdateUserEmailWithTokenUsingPOSTRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UpdateUserEmailWithTokenUsingPOSTExecute(r)
}

/*
 * UpdateUserEmailWithTokenUsingPOST Update user's email
 * Usage of this API will be reported in your access log under 'authentication' category.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiUpdateUserEmailWithTokenUsingPOSTRequest
 */
func (a *UserAuthenticationApiService) UpdateUserEmailWithTokenUsingPOST(ctx _context.Context) ApiUpdateUserEmailWithTokenUsingPOSTRequest {
	return ApiUpdateUserEmailWithTokenUsingPOSTRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *UserAuthenticationApiService) UpdateUserEmailWithTokenUsingPOSTExecute(r ApiUpdateUserEmailWithTokenUsingPOSTRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAuthenticationApiService.UpdateUserEmailWithTokenUsingPOST")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/updateEmail"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAPIKEY == nil {
		return nil, reportError("xAPIKEY is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-API-KEY"] = parameterToString(*r.xAPIKEY, "")
	// body params
	localVarPostBody = r.userUpdateEmailRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateUserPasswordWithTokenUsingPOSTRequest struct {
	ctx _context.Context
	ApiService *UserAuthenticationApiService
	userUpdatePasswordRequest *UpdatePasswordWithTokenReqWeb
}

func (r ApiUpdateUserPasswordWithTokenUsingPOSTRequest) UserUpdatePasswordRequest(userUpdatePasswordRequest UpdatePasswordWithTokenReqWeb) ApiUpdateUserPasswordWithTokenUsingPOSTRequest {
	r.userUpdatePasswordRequest = &userUpdatePasswordRequest
	return r
}

func (r ApiUpdateUserPasswordWithTokenUsingPOSTRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UpdateUserPasswordWithTokenUsingPOSTExecute(r)
}

/*
 * UpdateUserPasswordWithTokenUsingPOST Update user's password
 * Usage of this API will be reported in your access log under 'authentication' category.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiUpdateUserPasswordWithTokenUsingPOSTRequest
 */
func (a *UserAuthenticationApiService) UpdateUserPasswordWithTokenUsingPOST(ctx _context.Context) ApiUpdateUserPasswordWithTokenUsingPOSTRequest {
	return ApiUpdateUserPasswordWithTokenUsingPOSTRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *UserAuthenticationApiService) UpdateUserPasswordWithTokenUsingPOSTExecute(r ApiUpdateUserPasswordWithTokenUsingPOSTRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserAuthenticationApiService.UpdateUserPasswordWithTokenUsingPOST")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v0/setpwd"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.userUpdatePasswordRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
