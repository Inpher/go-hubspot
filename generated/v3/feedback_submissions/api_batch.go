/*
Feedback Submissions

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package feedback_submissions

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/inpher/go-hubspot"
	"net/url"
)

// BatchAPIService BatchAPI service
type BatchAPIService service

type ApiPostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchiveRequest struct {
	ctx                            context.Context
	ApiService                     *BatchAPIService
	batchInputSimplePublicObjectId *BatchInputSimplePublicObjectId
}

func (r ApiPostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchiveRequest) BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId BatchInputSimplePublicObjectId) ApiPostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchiveRequest {
	r.batchInputSimplePublicObjectId = &batchInputSimplePublicObjectId
	return r
}

func (r ApiPostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchiveExecute(r)
}

/*
PostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchive Archive a batch of feedback submissions by ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchiveRequest
*/
func (a *BatchAPIService) PostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchive(ctx context.Context) ApiPostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchiveRequest {
	return ApiPostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchiveRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *BatchAPIService) PostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchiveExecute(r ApiPostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchAPIService.PostCrmV3ObjectsFeedbackSubmissionsBatchArchiveArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/feedback_submissions/batch/archive"

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreateRequest struct {
	ctx                                        context.Context
	ApiService                                 *BatchAPIService
	batchInputSimplePublicObjectInputForCreate *BatchInputSimplePublicObjectInputForCreate
}

func (r ApiPostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreateRequest) BatchInputSimplePublicObjectInputForCreate(batchInputSimplePublicObjectInputForCreate BatchInputSimplePublicObjectInputForCreate) ApiPostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreateRequest {
	r.batchInputSimplePublicObjectInputForCreate = &batchInputSimplePublicObjectInputForCreate
	return r
}

func (r ApiPostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreateRequest) Execute() (*BatchResponseSimplePublicObject, *http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreateExecute(r)
}

/*
PostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreate Create a batch of feedback submissions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreateRequest
*/
func (a *BatchAPIService) PostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreate(ctx context.Context) ApiPostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreateRequest {
	return ApiPostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BatchResponseSimplePublicObject
func (a *BatchAPIService) PostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreateExecute(r ApiPostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreateRequest) (*BatchResponseSimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseSimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchAPIService.PostCrmV3ObjectsFeedbackSubmissionsBatchCreateCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/feedback_submissions/batch/create"

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

type ApiPostCrmV3ObjectsFeedbackSubmissionsBatchReadReadRequest struct {
	ctx                                context.Context
	ApiService                         *BatchAPIService
	batchReadInputSimplePublicObjectId *BatchReadInputSimplePublicObjectId
	archived                           *bool
}

func (r ApiPostCrmV3ObjectsFeedbackSubmissionsBatchReadReadRequest) BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId BatchReadInputSimplePublicObjectId) ApiPostCrmV3ObjectsFeedbackSubmissionsBatchReadReadRequest {
	r.batchReadInputSimplePublicObjectId = &batchReadInputSimplePublicObjectId
	return r
}

// Whether to return only results that have been archived.
func (r ApiPostCrmV3ObjectsFeedbackSubmissionsBatchReadReadRequest) Archived(archived bool) ApiPostCrmV3ObjectsFeedbackSubmissionsBatchReadReadRequest {
	r.archived = &archived
	return r
}

func (r ApiPostCrmV3ObjectsFeedbackSubmissionsBatchReadReadRequest) Execute() (*BatchResponseSimplePublicObject, *http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsFeedbackSubmissionsBatchReadReadExecute(r)
}

/*
PostCrmV3ObjectsFeedbackSubmissionsBatchReadRead Read a batch of feedback submissions by internal ID, or unique property values

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostCrmV3ObjectsFeedbackSubmissionsBatchReadReadRequest
*/
func (a *BatchAPIService) PostCrmV3ObjectsFeedbackSubmissionsBatchReadRead(ctx context.Context) ApiPostCrmV3ObjectsFeedbackSubmissionsBatchReadReadRequest {
	return ApiPostCrmV3ObjectsFeedbackSubmissionsBatchReadReadRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BatchResponseSimplePublicObject
func (a *BatchAPIService) PostCrmV3ObjectsFeedbackSubmissionsBatchReadReadExecute(r ApiPostCrmV3ObjectsFeedbackSubmissionsBatchReadReadRequest) (*BatchResponseSimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseSimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchAPIService.PostCrmV3ObjectsFeedbackSubmissionsBatchReadRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/feedback_submissions/batch/read"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchReadInputSimplePublicObjectId == nil {
		return localVarReturnValue, nil, reportError("batchReadInputSimplePublicObjectId is required and must be specified")
	}

	if r.archived != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "archived", r.archived, "")
	} else {
		var defaultValue bool = false
		r.archived = &defaultValue
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

type ApiPostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdateRequest struct {
	ctx                                    context.Context
	ApiService                             *BatchAPIService
	batchInputSimplePublicObjectBatchInput *BatchInputSimplePublicObjectBatchInput
}

func (r ApiPostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdateRequest) BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput BatchInputSimplePublicObjectBatchInput) ApiPostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdateRequest {
	r.batchInputSimplePublicObjectBatchInput = &batchInputSimplePublicObjectBatchInput
	return r
}

func (r ApiPostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdateRequest) Execute() (*BatchResponseSimplePublicObject, *http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdateExecute(r)
}

/*
PostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdate Update a batch of feedback submissions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdateRequest
*/
func (a *BatchAPIService) PostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdate(ctx context.Context) ApiPostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdateRequest {
	return ApiPostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BatchResponseSimplePublicObject
func (a *BatchAPIService) PostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdateExecute(r ApiPostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdateRequest) (*BatchResponseSimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseSimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchAPIService.PostCrmV3ObjectsFeedbackSubmissionsBatchUpdateUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/feedback_submissions/batch/update"

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
