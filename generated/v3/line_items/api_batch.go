/*
Line Items

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package line_items

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/clarkmcc/go-hubspot"
	"net/url"
)

// BatchApiService BatchApi service
type BatchApiService service

type ApiPostCrmV3ObjectsLineItemsBatchArchiveArchiveRequest struct {
	ctx                            context.Context
	ApiService                     *BatchApiService
	batchInputSimplePublicObjectId *BatchInputSimplePublicObjectId
}

func (r ApiPostCrmV3ObjectsLineItemsBatchArchiveArchiveRequest) BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId BatchInputSimplePublicObjectId) ApiPostCrmV3ObjectsLineItemsBatchArchiveArchiveRequest {
	r.batchInputSimplePublicObjectId = &batchInputSimplePublicObjectId
	return r
}

func (r ApiPostCrmV3ObjectsLineItemsBatchArchiveArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsLineItemsBatchArchiveArchiveExecute(r)
}

/*
PostCrmV3ObjectsLineItemsBatchArchiveArchive Archive a batch of line items by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsLineItemsBatchArchiveArchiveRequest
*/
func (a *BatchApiService) PostCrmV3ObjectsLineItemsBatchArchiveArchive(ctx context.Context) ApiPostCrmV3ObjectsLineItemsBatchArchiveArchiveRequest {
	return ApiPostCrmV3ObjectsLineItemsBatchArchiveArchiveRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *BatchApiService) PostCrmV3ObjectsLineItemsBatchArchiveArchiveExecute(r ApiPostCrmV3ObjectsLineItemsBatchArchiveArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostCrmV3ObjectsLineItemsBatchArchiveArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/line_items/batch/archive"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputSimplePublicObjectId == nil {
		return nil, reportError("batchInputSimplePublicObjectId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchInputSimplePublicObjectId
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPostCrmV3ObjectsLineItemsBatchCreateCreateRequest struct {
	ctx                                        context.Context
	ApiService                                 *BatchApiService
	batchInputSimplePublicObjectInputForCreate *BatchInputSimplePublicObjectInputForCreate
}

func (r ApiPostCrmV3ObjectsLineItemsBatchCreateCreateRequest) BatchInputSimplePublicObjectInputForCreate(batchInputSimplePublicObjectInputForCreate BatchInputSimplePublicObjectInputForCreate) ApiPostCrmV3ObjectsLineItemsBatchCreateCreateRequest {
	r.batchInputSimplePublicObjectInputForCreate = &batchInputSimplePublicObjectInputForCreate
	return r
}

func (r ApiPostCrmV3ObjectsLineItemsBatchCreateCreateRequest) Execute() (*BatchResponseSimplePublicObject, *http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsLineItemsBatchCreateCreateExecute(r)
}

/*
PostCrmV3ObjectsLineItemsBatchCreateCreate Create a batch of line items

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsLineItemsBatchCreateCreateRequest
*/
func (a *BatchApiService) PostCrmV3ObjectsLineItemsBatchCreateCreate(ctx context.Context) ApiPostCrmV3ObjectsLineItemsBatchCreateCreateRequest {
	return ApiPostCrmV3ObjectsLineItemsBatchCreateCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BatchResponseSimplePublicObject
func (a *BatchApiService) PostCrmV3ObjectsLineItemsBatchCreateCreateExecute(r ApiPostCrmV3ObjectsLineItemsBatchCreateCreateRequest) (*BatchResponseSimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseSimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostCrmV3ObjectsLineItemsBatchCreateCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/line_items/batch/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputSimplePublicObjectInputForCreate == nil {
		return localVarReturnValue, nil, reportError("batchInputSimplePublicObjectInputForCreate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchInputSimplePublicObjectInputForCreate
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostCrmV3ObjectsLineItemsBatchReadReadRequest struct {
	ctx                                context.Context
	ApiService                         *BatchApiService
	batchReadInputSimplePublicObjectId *BatchReadInputSimplePublicObjectId
	archived                           *bool
}

func (r ApiPostCrmV3ObjectsLineItemsBatchReadReadRequest) BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId BatchReadInputSimplePublicObjectId) ApiPostCrmV3ObjectsLineItemsBatchReadReadRequest {
	r.batchReadInputSimplePublicObjectId = &batchReadInputSimplePublicObjectId
	return r
}

// Whether to return only results that have been archived.
func (r ApiPostCrmV3ObjectsLineItemsBatchReadReadRequest) Archived(archived bool) ApiPostCrmV3ObjectsLineItemsBatchReadReadRequest {
	r.archived = &archived
	return r
}

func (r ApiPostCrmV3ObjectsLineItemsBatchReadReadRequest) Execute() (*BatchResponseSimplePublicObject, *http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsLineItemsBatchReadReadExecute(r)
}

/*
PostCrmV3ObjectsLineItemsBatchReadRead Read a batch of line items by internal ID, or unique property values

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsLineItemsBatchReadReadRequest
*/
func (a *BatchApiService) PostCrmV3ObjectsLineItemsBatchReadRead(ctx context.Context) ApiPostCrmV3ObjectsLineItemsBatchReadReadRequest {
	return ApiPostCrmV3ObjectsLineItemsBatchReadReadRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BatchResponseSimplePublicObject
func (a *BatchApiService) PostCrmV3ObjectsLineItemsBatchReadReadExecute(r ApiPostCrmV3ObjectsLineItemsBatchReadReadRequest) (*BatchResponseSimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseSimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostCrmV3ObjectsLineItemsBatchReadRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/line_items/batch/read"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchReadInputSimplePublicObjectId == nil {
		return localVarReturnValue, nil, reportError("batchReadInputSimplePublicObjectId is required and must be specified")
	}

	if r.archived != nil {
		localVarQueryParams.Add("archived", parameterToString(*r.archived, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchReadInputSimplePublicObjectId
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostCrmV3ObjectsLineItemsBatchUpdateUpdateRequest struct {
	ctx                                    context.Context
	ApiService                             *BatchApiService
	batchInputSimplePublicObjectBatchInput *BatchInputSimplePublicObjectBatchInput
}

func (r ApiPostCrmV3ObjectsLineItemsBatchUpdateUpdateRequest) BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput BatchInputSimplePublicObjectBatchInput) ApiPostCrmV3ObjectsLineItemsBatchUpdateUpdateRequest {
	r.batchInputSimplePublicObjectBatchInput = &batchInputSimplePublicObjectBatchInput
	return r
}

func (r ApiPostCrmV3ObjectsLineItemsBatchUpdateUpdateRequest) Execute() (*BatchResponseSimplePublicObject, *http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsLineItemsBatchUpdateUpdateExecute(r)
}

/*
PostCrmV3ObjectsLineItemsBatchUpdateUpdate Update a batch of line items

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsLineItemsBatchUpdateUpdateRequest
*/
func (a *BatchApiService) PostCrmV3ObjectsLineItemsBatchUpdateUpdate(ctx context.Context) ApiPostCrmV3ObjectsLineItemsBatchUpdateUpdateRequest {
	return ApiPostCrmV3ObjectsLineItemsBatchUpdateUpdateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BatchResponseSimplePublicObject
func (a *BatchApiService) PostCrmV3ObjectsLineItemsBatchUpdateUpdateExecute(r ApiPostCrmV3ObjectsLineItemsBatchUpdateUpdateRequest) (*BatchResponseSimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseSimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostCrmV3ObjectsLineItemsBatchUpdateUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/line_items/batch/update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputSimplePublicObjectBatchInput == nil {
		return localVarReturnValue, nil, reportError("batchInputSimplePublicObjectBatchInput is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchInputSimplePublicObjectBatchInput
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
